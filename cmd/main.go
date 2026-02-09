package main

import (
	"encoding/json"
	"groupie/server"
	"log"
	"net/http"
)

func main() {

	// Load API
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var artists []server.Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		log.Fatal(err)
	}

	server.AllArtists = artists

	// Start server
	server.Start()
}
