package main

import (
	"html/template"
	"net/http"
)

// Define a struct to hold the data for our HTML template
type PageData struct {
	Message string
}

func main() {
	// Define a handler function to serve the HTML page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Create an instance of PageData with a message
		data := PageData{
			Message: "Hello from Go!",
		}

		// Parse the HTML template
		tmpl := template.Must(template.ParseFiles("index.html"))

		// Execute the template, passing the data
		tmpl.Execute(w, data)
	})

	// Start the web server on port 8080
	http.ListenAndServe(":8080", nil)
}
