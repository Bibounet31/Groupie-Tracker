package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"
)

// ---STRUCTS---
type Artist struct {
	Id           int
	Name         string
	Image        string
	Members      []string
	CreationDate int
	Locations    string // URL to locations
	ConcertDates string // URL to dates
	Relations    string // URL to relations
}

type Location struct {
	Index     int      `json:"index"`
	Locations []string `json:"locations"`
}

type Relation struct {
	Index          int                 `json:"index"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

var AllArtists []Artist
var counter int

// send data to page
func render(w http.ResponseWriter, file string, data any) {
	funcMap := template.FuncMap{
		"add":        func(a, b int) int { return a + b },
		"pathEscape": url.PathEscape,
	}

	t, err := template.New(file).Funcs(funcMap).ParseFiles("web/html/" + file)
	if err != nil {
		http.Error(w, "Template not found", 500)
		return
	}
	t.Execute(w, data)
}

// load index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "index.html", nil)
}

// load artist page
func AlbumHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(AllArtists)
	render(w, "albums.html", AllArtists)
}

func SearchResultsHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("query")))
	member := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("member")))

	yearMinStr := r.URL.Query().Get("year_min")
	yearMaxStr := r.URL.Query().Get("year_max")
	membersCounts := r.URL.Query()["members_count"]

	fmt.Printf("DEBUG - Received params:\n")
	fmt.Printf("  query: '%s'\n", query)
	fmt.Printf("  member: '%s'\n", member)
	fmt.Printf("  year_min: '%s'\n", yearMinStr)
	fmt.Printf("  year_max: '%s'\n", yearMaxStr)
	fmt.Printf("  members_count: %v\n", membersCounts)

	var yearMin, yearMax int
	if yearMinStr != "" {
		fmt.Sscanf(yearMinStr, "%d", &yearMin)
	}
	if yearMaxStr != "" {
		fmt.Sscanf(yearMaxStr, "%d", &yearMax)
	}

	fmt.Printf("  Parsed yearMin: %d, yearMax: %d\n", yearMin, yearMax)

	var membersFilter []int
	for _, s := range membersCounts {
		if s == "5+" {
			membersFilter = append(membersFilter, 5)
		} else {
			var n int
			fmt.Sscanf(s, "%d", &n)
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

	fmt.Printf("  Total results: %d\n\n", len(results))

	render(w, "albums.html", results)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("query")))
	searchType := r.URL.Query().Get("type")

	// Initialize as empty array instead of nil
	results := []string{} // Changed from: var results []string

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
	json.NewEncoder(w).Encode(results)
}

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
			// Fetch relations (locations + dates)
			var relation Relation
			if artist.Relations != "" {
				resp, err := http.Get(artist.Relations)
				if err == nil {
					defer resp.Body.Close()
					json.NewDecoder(resp.Body).Decode(&relation)
				}
			}

			// Create data structure for template
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

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 303)
		return
	}

	http.Redirect(w, r, "/", 303)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	render(w, "404.html", nil)
}
