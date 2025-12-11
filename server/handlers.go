package server

import (
	"html/template"
	"net/http"
)

type PageData struct {
	A int
}

var counter int // global variable

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

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{A: counter}
	render(w, "index.html", data)
}

func ArtistesHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "artistes.html", nil)
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 303)
		return
	}

	counter++ // increase the value
	http.Redirect(w, r, "/", 303)
}
