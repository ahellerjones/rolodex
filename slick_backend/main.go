package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func main() {
	// Define the HTTP endpoints and their corresponding handler functions
	http.HandleFunc("/", landingHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/contacts", contactsHandler)

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
	loginOrSignup bool `json:loginOrSignup`
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
			http.Error(w, err.Error(), 410)
			return
		}
		// If we're just trying to login
		if loginInfo.loginOrSetup { 
			err = checkUsernamePassword(loginInfo) // TODO 
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.WriteHeader(http.StatusOK) // 200
			fmt.Fprint(w, "{%s}", loginInfo.user)
		} else { 
			// Else we're trying to setup an account 
			err = checkIfUserExists(loginInfo) // TODO 
			if err != nil { 
				http.Error(w, err.Error(), 500)
				return
			} 
			err = createUser(loginInfo)
			if err != nil { 
				http.Error(w, err.Error(), 500)
				return
			} 
		}
		default:
			w.WriteHeader(http.StatusConflict) // 410
			fmt.Fprint(w, "Illegal request to /login")
	}
}

func contactsHandler(w http.ResponseWriter, r *http.Request) {
switch r.Method {
	case http.MethodGet:
		getContacts	




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
		http.Error(w, err.Error(), 410)
		return
	}
	// If we're just trying to login
	if loginInfo.loginOrSetup { 
		err = checkUsernamePassword(loginInfo) // TODO 
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(http.StatusOK) // 200
		fmt.Fprint(w, "{%s}", loginInfo.user)
	} else { 
		// Else we're trying to setup an account 
		err = checkIfUserExists(loginInfo) // TODO 
		if err != nil { 
			http.Error(w, err.Error(), 500)
			return
		} 
		err = createUser(loginInfo)
		if err != nil { 
			http.Error(w, err.Error(), 500)
			return
		} 
	}
	default:
		w.WriteHeader(http.StatusConflict) // 410
		fmt.Fprint(w, "Illegal request to /login")
}