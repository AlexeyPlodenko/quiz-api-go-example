package main

// @TODO caching of questions

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Main
func main() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/v1").Subrouter()

	subRouter.HandleFunc("/questions/{id:\\d+}", QuestionsHandler).Methods("GET")
	subRouter.HandleFunc("/questions/{id:\\d+}/answers", AnswersHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":888", router))
	log.Println("Started server on port 888")
}
