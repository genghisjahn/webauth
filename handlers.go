package main

import "net/http"

func authHandler(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("un")
	password := r.FormValue("pw")
	authResp, err := ProcessAuthentication(userName, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if !authResp.LoggedIn {
		http.Error(w, "Something went wrong.", http.StatusBadRequest)
	}
}
