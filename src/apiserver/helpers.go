package main

import (
	"io"
	"log"
	"net/http"
)

// ReturnStatusBadRequest responds with HTTP 400 status and a custom message
func ReturnStatusBadRequest(response http.ResponseWriter, body string) {
	log.Println("Responding with HTTP status 400 Bad Request")

	response.WriteHeader(http.StatusBadRequest)
	response.Header().Set("Content-Type", "plain/text")
	if body != "" {
		io.WriteString(response, body+"\n")
	}
}

// ReturnStatusNotFound responds with HTTP 404 status
func ReturnStatusNotFound(response http.ResponseWriter) {
	log.Println("Responding with HTTP status 404 Not Found")

	response.WriteHeader(http.StatusNotFound)
}
