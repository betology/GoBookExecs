package main

import (
	"fmt"
	"log"
	"net/http"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Development")
}

func main() {
	mux := http.NewServeMux()

	// Convert the messageHandler function to HandlerFunc type
	mh := http.HandlerFunc(messageHandler)
	mux.Handle("/welcome", mh)

	log.Println("listening...")
	http.ListenAndServe(":8080", mux)
}
