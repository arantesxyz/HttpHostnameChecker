package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type info struct {
	Message string `json:"message"`
	Source  string `json:"source"`
	Method  string `json:"method"`
	Host    string `json:"host"`
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info{
		Message: "Hello World!!!",
		Source:  "https://github.com/arantesxyz/HttpHostnameChecker",
		Method:  r.Method,
		Host:    r.Host,
	})

	fmt.Println("Received " + r.Method + " request from " + r.RemoteAddr + " with host " + r.Host)
}

func main() {

	http.HandleFunc("/", home)

	port := os.Getenv("PORT")
	fmt.Printf("Starting server on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
