package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "go-api-test/database"
	"go-api-test/handlers"
)

func main() {
	fmt.Println("Starting API server...")

	router := mux.NewRouter()

	// Root endpoint
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Go API!")
	}).Methods("GET")

	// Hello endpoint
	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	}).Methods("GET")

	// Items endpoints
	router.HandleFunc("/items", handlers.GetItems).Methods("GET")
	router.HandleFunc("/items", handlers.CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", handlers.GetItem).Methods("GET")
	router.HandleFunc("/items/{id}", handlers.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", handlers.DeleteItem).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
