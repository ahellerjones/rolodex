package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"json"
)

func main() {
	// Define the HTTP endpoints and their corresponding handler functions
	http.HandleFunc("/", landingHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/user", userHandler)

	// Start the server on port 8080
	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


func landingHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../sexy_frontend/dist/index.html")
}
type LoginInfo struct { 
	user string `json:username`
	pass string `json:password`
	loginOrSetup bool `json:loginOrSetup`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		var loginInfo LoginInfo
		err = json.Unmarshal(body, &loginInfo)
		// This should catch all shitty unmarshalling
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		// If we're just trying to login
		if loginInfo.loginOrSetup { 
			err = checkUsernamePassword(loginInfo)
			if err != nil {
				
			}
		}



		default:
			w.WriteHeader(http.StatusConflict) // 410
			fmt.Fprint(w, "Illegal request to /login")
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// Set the response content type
	w.Header().Set("Content-Type", "application/json")

	// Simulate fetching user data from a database
	user := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

}