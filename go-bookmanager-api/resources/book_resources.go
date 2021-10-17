package resources

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"matheusmv.com/go-bookmanager-api/models"
	"matheusmv.com/go-bookmanager-api/services"
)

type bookResourceHandler struct {
	service services.BookService
}

func NewBookResourceHandler(service services.BookService) *bookResourceHandler {
	return &bookResourceHandler{service}
}

func (handler *bookResourceHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, isPresent := params["id"]

	if !isPresent {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode([]string{"ID required"})
		return
	}

	bookID, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode([]string{"invalid ID"})
		return
	}

	book, err := handler.service.FindABookById(bookID)

	if err != nil || book.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode([]string{"not found"})
		return
	}

	json.NewEncoder(w).Encode(book)
}

func (handler *bookResourceHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := handler.service.GetAllBooks()

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		json.NewEncoder(w).Encode(books)
		return
	}

	json.NewEncoder(w).Encode(books)
}

func (handler *bookResourceHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book models.Book

	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newBook, err := handler.service.CreateABook(book)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func (handler *bookResourceHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, isPresent := params["id"]

	if !isPresent {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode([]string{"ID required"})
		return
	}

	bookID, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode([]string{"invalid ID"})
		return
	}

	var book models.Book

	err = json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	updatedBook, err := handler.service.UpdateABook(bookID, book)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedBook)
}

func (handler *bookResourceHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, isPresent := params["id"]

	if !isPresent {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode([]string{"ID required"})
		return
	}

	bookID, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode([]string{"invalid ID"})
		return
	}

	deleted, err := handler.service.DeleteABook(bookID)

	if err != nil || !deleted {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode([]string{})
}
