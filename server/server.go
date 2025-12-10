package server

import (
	"fmt"
	"html/template"
	"net/http"
)

func render(w http.ResponseWriter, file string, data any) { //render template
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

// -------------------- HANDLERS --------------------

func indexHandler(w http.ResponseWriter, r *http.Request) { //load index.html
	render(w, "index.html", nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) { //executed when a submit request is posted
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 303)
		return
	}
}

// -------------------- PUBLIC FUNCTION --------------------

func Start() {
	// Routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/submit", submitHandler)

	// Static files
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("web/css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("web/img"))))

	fmt.Println("✅ Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
