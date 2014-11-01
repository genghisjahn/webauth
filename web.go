package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Started")
	http.HandleFunc("/auth", authHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
