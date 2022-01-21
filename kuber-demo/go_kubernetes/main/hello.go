package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/api/hello", Hello)
	http.ListenAndServe("0.0.0.0:8081", handler)
	// Listen to all request to port 8081.
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Hello world`)
}
