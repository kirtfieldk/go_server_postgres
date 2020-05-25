package main

import (
	"log"
	"net/http"

	"github.com/keithkfield/pg_api/lp_routes"
)

func main() {
	http.HandleFunc("/", lp_routes.MainPage)
	http.HandleFunc("/search", lp_routes.SearchRoute)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
