package main

import (
	"log"
	"net/http"

	"github.com/Kshitij-Jain99/Golang_Minor_Projects/bookstore-manager/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
