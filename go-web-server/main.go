package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PostBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/post", postRequestDetails)

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	fmt.Println("Starting server at port 80")
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go Lovvers! \n")
	println("Received request for:", r.URL.Path)
	println("Request method:", r.Method)
	params := fmt.Sprintf("Request Params: %s\n", r.URL.Query().Get("name"))
	fmt.Fprintf(w, "%s", params)
}

func postRequestDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "Only POST method is allowed")
		return
	}
	defer r.Body.Close()
	var postBody PostBody
	err := json.NewDecoder(r.Body).Decode(&postBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Received POST request with body: %+v \n", postBody)
	fmt.Fprintf(w, "Name: %+v\n", postBody.Name)
	fmt.Fprintf(w, "Email: %+v\n", postBody.Email)
}
