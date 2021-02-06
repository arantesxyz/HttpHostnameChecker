package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received " + r.Method + " request from " + r.RemoteAddr + " with host " + r.Host)
	fmt.Fprintf(w, "URI Host: "+r.URL.Host+" Scheme: "+r.URL.Scheme)
	fmt.Fprintf(w, "\nHost: "+r.Host)
}

func main() {

	http.HandleFunc("/", home)

	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
