package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"bookstore/pkg/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.BookStoreRoutes(r)
	http.Handle("/", r)

	fmt.Println("Server started at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
