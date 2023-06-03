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
type UserId struct { 
	id int `json:id`
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
		if loginInfo.loginOrSignup { 
			userId, err := checkUsernamePassword(loginInfo) // TODO 
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.WriteHeader(http.StatusOK) // 200
			userIdResponse := UserId { 
				id: userId,
			}
			jsonData, err := json.Marshal(userIdResponse)
			if err != nil {
				http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")

			w.Write(jsonData)
		} else { 
			// Else we're trying to setup an account 
			err = checkIfUserExists(loginInfo) // TODO 
			if err != nil { 
				http.Error(w, err.Error(), 500)
				return
			} 
			userId, err := createUser(loginInfo)
			if err != nil { 
				http.Error(w, err.Error(), 500)
				return
			} 
			w.WriteHeader(http.StatusOK) // 200
			fmt.Fprint(w, `{"id":%d}`, userId) // Write back userId 
		}
		default:
			w.WriteHeader(http.StatusConflict) // 410
			fmt.Fprint(w, "Illegal request to /login")
	}
}

func checkUsernamePassword(loginInfo LoginInfo)(err error) { 
	return nil
}

func checkIfUserExists(loginInfo LoginInfo)(err error) { 
	return nil
}

func createUser(loginInfo LoginInfo)(err error) { 
	return nil
}


type Contact struct { 
	name string `json:name`
	address string `json:address`
	phoneNumber string `json:phoneNumber`
	email string `json:email`
	birthday string `json:birthday`
}

func contactsHandler(w http.ResponseWriter, r *http.Request) {
switch r.Method {
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		var contact Contact
		err = json.Unmarshal(body, &contact)
		// This should catch all shitty unmarshalling
		if err != nil {
			http.Error(w, err.Error(), 410)
			return
		}
		id, err := storeContact(contact)
		if err != nil { 
			http.Error(w, err.Error(), 410)
			return
		}
		w.WriteHeader(http.StatusOK) // 200
		fmt.Fprint(w, `{"id":%d}`, id) // Write back userId 


	case http.MethodGet:
	contacts, err := getContacts()
	if err != nil { 
		http.Error(w, err.Error(), 410)
		return
	}
	// Marshal the struct into JSON
	data, err := json.Marshal(contacts)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	default:
		w.WriteHeader(http.StatusConflict) // 410
		fmt.Fprint(w, "Illegal request to /login")
	}
}