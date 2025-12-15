package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"
)

type PageData struct {
	A int
}

type Artist struct {
	Id           int
	Name         string
	Image        string
	Members      []string
	CreationDate int
}

var AllArtists []Artist
var counter int // global variable

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
	data := PageData{A: counter}
	render(w, "index.html", data)
}

// load artist page
func ArtistesHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "artistes.html", AllArtists)
}

func SearchResultsHandler(w http.ResponseWriter, r *http.Request) {
	// Artist search query
	query := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("query")))
	member := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("member")))

	// Filter parameters
	yearMinStr := r.URL.Query().Get("year_min")
	yearMaxStr := r.URL.Query().Get("year_max")
	membersCounts := r.URL.Query()["members_count"] // can be multiple

	var yearMin, yearMax int
	if yearMinStr != "" {
		fmt.Sscanf(yearMinStr, "%d", &yearMin)
	}
	if yearMaxStr != "" {
		fmt.Sscanf(yearMaxStr, "%d", &yearMax)
	}

	// Convert membersCounts to integers
	var membersFilter []int
	for _, s := range membersCounts {
		if s == "5+" {
			membersFilter = append(membersFilter, 5) // we'll treat 5 as "5 or more"
		} else {
			var n int
			fmt.Sscanf(s, "%d", &n)
			membersFilter = append(membersFilter, n)
		}
	}

	var results []Artist
	for _, artist := range AllArtists {
		// Artist/member search
		match := false
		if query != "" {
			if strings.Contains(strings.ToLower(artist.Name), query) {
				match = true
			} else if member != "" {
				for _, m := range artist.Members {
					if strings.Contains(strings.ToLower(m), member) {
						match = true
						break
					}
				}
			}
		} else {
			match = true // no query = match all
		}

		// Filter by year
		if yearMin != 0 && artist.CreationDate < yearMin {
			match = false
		}
		if yearMax != 0 && artist.CreationDate > yearMax {
			match = false
		}

		// Filter by number of members
		if len(membersFilter) > 0 {
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

	render(w, "artistes.html", results)
}

// SearchHandler returns matching artist names as JSON
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("query")))
	searchType := r.URL.Query().Get("type") // "artist" or "member"

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
			render(w, "details.html", artist)
			return
		}
	}

	http.NotFound(w, r)
}

// when submit is sent
func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 303)
		return
	}

	counter++ // increase the value
	http.Redirect(w, r, "/", 303)
}
