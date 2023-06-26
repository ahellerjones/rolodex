package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type LoginInfo struct {
	Username      string `json:"Username"`
	Pass          string `json:"Password"`
	LoginOrSignup bool   `json:"loginOrSignup"`
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
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
		if loginInfo.LoginOrSignup {
			userId, err := h.Database.InsertUser(loginInfo) // TODO
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.WriteHeader(http.StatusOK) // 200
			userIdResponse := Identification{
				UserID: userId,
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
			exists, err := h.Database.CheckUsernameExists(loginInfo) // TODO
			if exists || err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			userId, err := h.Database.InsertUser(loginInfo)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.WriteHeader(http.StatusOK) // 200
			userIdResponse := Identification{
				UserID: userId,
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
