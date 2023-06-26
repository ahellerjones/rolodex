package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler contains a reference to our db object
type Handler struct {
	Database *SQLiteHandler
}

func NewHandler(dbPath string) (*Handler, error) {
	db, err := NewSQLiteHandler(dbPath)
	if err != nil {
		fmt.Printf("Failed to create handler: %s\\n", err)
		return nil, err
	}

	h := Handler{
		Database: db,
	}
	return &h, nil
}

func main() {
	handler, err := NewHandler("./rolodex.db")
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	// Define the HTTP endpoints and their corresponding handler functions
	http.HandleFunc("/", landingHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/contacts", handler.ContactsHandler)

	// Start the server on port 8080
	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func landingHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../sexy_frontend/dist/index.html")
}

// rest of handlers are in their own respective files
