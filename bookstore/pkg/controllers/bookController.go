package controllers

import (
	"bookstore/pkg/models"
	"bookstore/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var newBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := models.GetAllBooks()
	response, err := (&utils.Response{}).PrepareResponse(books, http.StatusOK, "Books retrieved successfully")
	if err != nil {
		http.Error(w, "Failed to prepare response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bookId := mux.Vars(r)["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		response, _ := (&utils.Response{}).PrepareResponse(nil, http.StatusBadRequest, "Invalid book ID")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	bookDetail, _ := models.GetBookById(ID)

	if bookDetail.ID == 0 {
		response, _ := (&utils.Response{}).PrepareResponse(nil, http.StatusNotFound, "Book not found")
		w.WriteHeader(http.StatusNotFound)
		w.Write(response)
		return
	}

	response, err := (&utils.Response{}).PrepareResponse(bookDetail, http.StatusOK, "Book retrieved successfully")
	if err != nil {
		http.Error(w, "Failed to prepare response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {}

func UpdateBook(w http.ResponseWriter, r *http.Request) {}

func DeleteBook(w http.ResponseWriter, r *http.Request) {}
