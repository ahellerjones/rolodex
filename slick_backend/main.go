package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	username      string `json:username`
	pass          string `json:password`
	loginOrSignup bool   `json:loginOrSignup`
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
			userIdResponse := UserId{
				id: userId,
			}
			jsonData, err := json.Marshal(userIdResponse)
			if err != nil {
				http.Error(w, "Failed to marshal JSON during login", http.StatusInternalServerError)
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
			userIdResponse := UserId{
				id: userId,
			}
			jsonData, err := json.Marshal(userIdResponse)
			if err != nil {
				http.Error(w, "Failed to marshal JSON during setup", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		}
	default:
		http.Error(w, "Illedate request to /login", http.StatusConflict)
	}
}

func checkUsernamePassword(loginInfo LoginInfo) (err error) {
	return nil
}

func checkIfUserExists(loginInfo LoginInfo) (err error) {
	return nil
}

func createUser(loginInfo LoginInfo) (err error) {
	return nil
}

// Contact struct contains userID to associate each contact entry with a particular user.
// UserId should probably just be a string TODO:
// I think key should probably be removed too.
type Contact struct {
	UserId      UserId `json:"UserId"` // Also need this!
	Key         int    `json:"Key,omitempty"`
	Name        string `json:"Name"` // Need this!
	Address     string `json:"Address,omitempty"`
	PhoneNumber string `json:"PhoneNumber,omitempty"`
	Email       string `json:"Email,omitempty"`
	Birthday    string `json:"Birthday,omitempty"`
}

type ContactKey struct {
	Key int `json:"Key"`
}

// endpoint /contacts -- for reading a writing contact data.
// POST -- user is storing a conact
// GET -- gets all contacts
// UPDATE -- updates one contact with id
// DELETE -- deletes one contact
func contactsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// POST
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
		key, err := storeContact(contact) // This gets a key for the given contact struct
		if err != nil {
			http.Error(w, err.Error(), 410)
			return
		}
		w.WriteHeader(http.StatusOK) // 200
		contactsKey := ContactKey{
			Key: key,
		}
		jsonData, err := json.Marshal(contactsKey)
		if err != nil {
			http.Error(w, "Failed to marshal JSON during creating a contact", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)

	// GET
	case http.MethodGet:
		contacts, err := getContacts()
		if err != nil {
			http.Error(w, err.Error(), 410)
			return
		}
		// Marshal the struct into JSON
		jsonData, err := json.Marshal(contacts)
		if err != nil {
			http.Error(w, "Failed to marshal JSON during GET contacts", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	// PUT
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body put", http.StatusInternalServerError)
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
		key, err := storeContact(contact)
		if err != nil {
			http.Error(w, err.Error(), 410)
			return
		}
		w.WriteHeader(http.StatusOK) // 200
		contactKey := ContactKey{
			Key: key,
		}
		jsonData, err := json.Marshal(contactKey)
		if err != nil {
			http.Error(w, "Failed to marshal JSON during creating a contact", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	case http.MethodDelete:

	default:
		w.WriteHeader(http.StatusConflict) // 410
		fmt.Fprint(w, "Illegal request to /login")
	}
}
