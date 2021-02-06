package main

import (
	"fmt"
	"net/http"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received " + r.Method + " request from " + r.RemoteAddr + " with host " + r.Host)
	fmt.Fprintf(w, "URI Host: "+r.URL.Host+" Scheme: "+r.URL.Scheme)
	fmt.Fprintf(w, "\nHost: "+r.Host)
}

func main() {

	http.HandleFunc("/", home)

	port := os.Getenv("PORT")
	fmt.Printf("Starting server on port %s...", port)
	http.ListenAndServe(":"+port, nil)
}
