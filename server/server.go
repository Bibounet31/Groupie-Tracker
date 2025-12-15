package server

import (
	"fmt"
	"net/http"
)

func Start() {
	// Routes
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/submit", SubmitHandler)
	http.HandleFunc("/artistes", ArtistesHandler)
	http.HandleFunc("/details/", DetailsHandler)

	// Static files
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("web/css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("web/img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("web/js"))))
	fmt.Println("✅ Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
