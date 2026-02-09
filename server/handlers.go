package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"
)

// Artist represents an artist with their information
type Artist struct {
	Id           int
	Name         string
	Image        string
	Members      []string
	CreationDate int
	Relations    string
}

// Relation link dates and locations
type Relation struct {
	Index          int                 `json:"index"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// AllArtists keep all the artists loaded from the API
var AllArtists []Artist

// render sends data to the specified HTML template
func render(w http.ResponseWriter, file string, data any) {
	funcMap := template.FuncMap{
		"pathEscape": url.PathEscape,
	}

	t, err := template.New(file).Funcs(funcMap).ParseFiles("web/html/" + file)
	if err != nil {
		http.Error(w, "Template not found", 500)
		return
	}
	_ = t.Execute(w, data)
}

// IndexHandler loads the index page
func IndexHandler(w http.ResponseWriter, _ *http.Request) {
	render(w, "index.html", nil)
}

// AlbumHandler loads the artist page
func AlbumHandler(w http.ResponseWriter, _ *http.Request) {
	render(w, "albums.html", AllArtists)
}

// SearchResultsHandler handles filtered search results
func SearchResultsHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("query")))
	member := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("member")))

	yearMinStr := r.URL.Query().Get("year_min")
	yearMaxStr := r.URL.Query().Get("year_max")
	membersCounts := r.URL.Query()["members_count"]

	var yearMin, yearMax int
	if yearMinStr != "" {
		_, _ = fmt.Sscanf(yearMinStr, "%d", &yearMin)
	}
	if yearMaxStr != "" {
		_, _ = fmt.Sscanf(yearMaxStr, "%d", &yearMax)
	}

	var membersFilter []int
	for _, s := range membersCounts {
		if s == "5+" {
			membersFilter = append(membersFilter, 5)
		} else {
			var n int
			_, _ = fmt.Sscanf(s, "%d", &n)
			if n > 0 {
				membersFilter = append(membersFilter, n)
			}
		}
	}

	var results []Artist
	for _, artist := range AllArtists {
		match := true

		if query != "" || member != "" {
			searchMatch := false

			if query != "" && strings.Contains(strings.ToLower(artist.Name), query) {
				searchMatch = true
			}

			if member != "" {
				for _, m := range artist.Members {
					if strings.Contains(strings.ToLower(m), member) {
						searchMatch = true
						break
					}
				}
			}

			if !searchMatch {
				match = false
			}
		}

		if match && yearMin != 0 && artist.CreationDate < yearMin {
			match = false
		}

		if match && yearMax != 0 && artist.CreationDate > yearMax {
			match = false
		}

		if match && len(membersFilter) > 0 {
			memberCount := len(artist.Members)
			matchedCount := false
			for _, mc := range membersFilter {
				if mc == 5 && memberCount >= 5 {
					matchedCount = true
					break
				} else if mc == memberCount {
					matchedCount = true
					break
				}
			}
			if !matchedCount {
				match = false
			}
		}

		if match {
			results = append(results, artist)
		}
	}

	render(w, "albums.html", results)
}

// SearchHandler handles search autocomplete
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("query")))
	searchType := r.URL.Query().Get("type")

	var results []string

	if query != "" {
		for _, artist := range AllArtists {
			if searchType == "member" {
				for _, m := range artist.Members {
					if strings.Contains(strings.ToLower(m), query) {
						results = append(results, m)
					}
				}
			} else {
				if strings.Contains(strings.ToLower(artist.Name), query) {
					results = append(results, artist.Name)
				}
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(results)
}

// DetailsHandler send to details page for the specific album
func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/details/")
	name, err := url.PathUnescape(path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	name = strings.TrimSpace(name)

	for _, artist := range AllArtists {
		if strings.EqualFold(artist.Name, name) {
			var relation Relation
			if artist.Relations != "" {
				resp, err := http.Get(artist.Relations)
				if err == nil {
					_ = json.NewDecoder(resp.Body).Decode(&relation)
					_ = resp.Body.Close()
				}
			}

			data := struct {
				Artist
				DatesLocations map[string][]string
			}{
				Artist:         artist,
				DatesLocations: relation.DatesLocations,
			}

			render(w, "details.html", data)
			return
		}
	}

	http.NotFound(w, r)
}

// NotFoundHandler displays error 404 page
func NotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	render(w, "404.html", nil)
}
