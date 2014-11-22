package main

import "net/http"

func authHandler(w http.ResponseWriter, r *http.Request) {
	//http.Error(w, "Not Implemented.", http.StatusInternalServerError)
	userName := r.FormValue("un")
	password := r.FormValue("pw")
	ProcessAuthentication(userName, password)
}
