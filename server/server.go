package server

import (
	"fmt"
	"net/http"
)

type responseWriterWrapper struct {
	http.ResponseWriter
	wroteHeader bool
}

func (rw *responseWriterWrapper) WriteHeader(statusCode int) {
	rw.wroteHeader = true
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriterWrapper) Write(b []byte) (int, error) {
	rw.wroteHeader = true
	return rw.ResponseWriter.Write(b)
}

func Start() {
	// Routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			NotFoundHandler(w, r)
			return
		}
		IndexHandler(w, r)
	})

	http.HandleFunc("/submit", SubmitHandler)
	http.HandleFunc("/album", AlbumHandler)
	http.HandleFunc("/details/", DetailsHandler)
	http.HandleFunc("/search", SearchHandler)
	http.HandleFunc("/artists", SearchResultsHandler)

	// Static files
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("web/css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("web/img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("web/js"))))

	fmt.Println("✅ Server running at http://localhost:8080")

	// Wrap default mux to handle 404
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriterWrapper{ResponseWriter: w}
		http.DefaultServeMux.ServeHTTP(rw, r)

		if !rw.wroteHeader {
			NotFoundHandler(w, r)
		}
	})

	http.ListenAndServe(":8080", handler)
}
