package main
import (
	"net/http"
	"fmt"
)


func main() {
	// Create a ner ServeMux
	mux := http.NewServeMux()

	// Register some routes
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/contact", contactHandler)

	// Start the server
	http.ListenAndServe(":8080", mux)
}

func homeHandler(w  http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola about")
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola contact")
}

