package server

import (
	"encoding/json"
	"html/template"
	"net/http"
	"net/url"
	"strings"
)

type Artist struct {
	Id           int
	Name         string
	Image        string
	Members      []string
	CreationDate int
	Locations    string
	ConcertDates string
	Relations    string
}

type Relation struct {
	Index          int                 `json:"index"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

var AllArtists []Artist

// ---------------- TEMPLATE LOADER----------------

func render(w http.ResponseWriter, file string, data any) {
	t, err := template.ParseFiles("web/html/" + file)
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

// ---------------- RENDER PAGES ----------------

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "index.html", nil)
}

func AlbumHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "albums.html", AllArtists)
}

// ---------------- SEARCH RESULTS ----------------

func SearchResultsHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("query")))
	member := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("member")))

	yearMin := r.URL.Query().Get("year_min")
	yearMax := r.URL.Query().Get("year_max")

	var results []Artist

	for _, artist := range AllArtists {

		// search by name
		if query != "" && !strings.Contains(strings.ToLower(artist.Name), query) {
			continue
		}

		// search by member
		if member != "" {
			found := false
			for _, m := range artist.Members {
				if strings.Contains(strings.ToLower(m), member) {
					found = true
				}
			}
			if !found {
				continue
			}
		}

		// filter by min year
		if yearMin != "" && artist.CreationDate < toInt(yearMin) {
			continue
		}

		// filter by max year
		if yearMax != "" && artist.CreationDate > toInt(yearMax) {
			continue
		}

		results = append(results, artist)
	}

	render(w, "albums.html", results)
}

// ---------------- AUTOCOMPLETE SEARCH ----------------

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("query")))
	searchType := r.URL.Query().Get("type")

	results := []string{}

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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// ---------------- DETAILS ----------------

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/details/")
	name, err := url.PathUnescape(path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	for _, artist := range AllArtists {
		if strings.EqualFold(artist.Name, name) {

			var relation Relation
			if artist.Relations != "" {
				resp, err := http.Get(artist.Relations)
				if err == nil {
					defer resp.Body.Close()
					json.NewDecoder(resp.Body).Decode(&relation)
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

// ---------------- FORM ----------------

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// ---------------- 404 ----------------

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	render(w, "404.html", nil)
}

// ---------------- UTILS ----------------
// convert year str to an int
func toInt(s string) int {
	n := 0
	for _, c := range s {
		n = n*10 + int(c-'0')
	}
	return n
}
