package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About")
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about_handler)
	fmt.Println("Server Starting")
	http.ListenAndServe(":9192", nil)
}
