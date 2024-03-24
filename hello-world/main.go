package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Respond with "Hello, World!"
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Serve static files
	http.Handle("/", http.FileServer(http.Dir(".")))

	// Handle AJAX request for /hello
	http.HandleFunc("/hello", helloHandler)

	// Start the server
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
