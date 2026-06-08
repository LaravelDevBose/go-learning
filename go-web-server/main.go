package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	fmt.Println("Starting server at port 80")
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go Lovvers!")
}
