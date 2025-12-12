package server

import (
	"html/template"
	"net/http"
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
		"add": func(a, b int) int { return a + b },
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

// when submit is sent
func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 303)
		return
	}

	counter++ // increase the value
	http.Redirect(w, r, "/", 303)
}
