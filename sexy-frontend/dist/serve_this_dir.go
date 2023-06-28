package main

import (
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	// Define the directory you want to serve
	dir := "."

	// Create a file server handler
	fs := http.FileServer(http.Dir(dir))

	// Register the file server handler for non-root URLs
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Register a custom handler for the root URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file := filepath.Join(dir, "index.html")
		http.ServeFile(w, r, file)
	})

	// Start the server
	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
