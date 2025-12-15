package server

import (
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
