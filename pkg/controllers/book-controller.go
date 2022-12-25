package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/codeshubham/bookstore-management-system/pkg/models"
	"github.com/codeshubham/bookstore-management-system/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, err := json.Marshal(newBooks)
	if err != nil {
		http.Error(w, "Unable to unmarshal", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["bookid"]
	ID, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		http.Error(w, "Error while parsing id from url", http.StatusBadRequest)
		return
	}
	bookDetails, _ := models.GetBookByID(ID)
	res, err := json.Marshal(bookDetails)
	if err != nil {
		http.Error(w, "Unable to unmarshal", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Write(res)
}

func DeleteBook(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookid"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		http.Error(rw, "not able to parse", http.StatusBadRequest)
		return
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(res)
}

func UpdateBook(rw http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	id := vars["bookid"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		http.Error(rw, "not able to parse id", http.StatusBadRequest)
		return
	}
	bookDetails, db := models.GetBookByID(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(res)
}
