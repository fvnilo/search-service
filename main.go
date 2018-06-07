package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nylo-andry/search-service/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/populate", handlers.Populate).
		Queries("number", "{number:[0-9]+}").
		Methods("POST")

	r.HandleFunc("/search", handlers.Search).
		Queries("q", "{q:[A-Za-z0-9]+}", "from", "{from:[0-9]+}", "size", "{size:[0-9]+}").
		Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
