package main

import (
	"log"
	"net/http"

	"github.com/nylo-andry/search-service/handlers"
)

func main() {
	mux := http.NewServeMux()

	// TODO: Add Verbs and queries
	mux.HandleFunc("/populate", handlers.Populate)
	mux.HandleFunc("/search", handlers.Search)

	log.Fatal(http.ListenAndServe(":8000", mux))
}
