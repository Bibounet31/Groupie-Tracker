package server

import (
	"fmt"
	"net/http"
)

func Start() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/submit", submitHandler)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("web/css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("web/img"))))

	fmt.Println("✅ Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
