package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"project/models"
	"project/service"
	"strconv"

	"github.com/gorilla/mux"
)

func AllBooksHandler(bookService service.BookServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var listBooks models.BookListResponse

		books, err := bookService.ListBooks()
		if err != nil {
			listBooks.ErrorMsg = err.Error()
			listBooks.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(listBooks)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}
		listBooks.BookList = books
		listBooks.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(listBooks)
		w.Write(res)

	}
}

func CreateBooksHandler(bookService service.BookServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var createBookResponse models.BookResponse

		err := json.NewDecoder(r.Body).Decode(&book)

		if err != nil {
			createBookResponse.Message = "unable to decode the request body."
			createBookResponse.StatusCode = http.StatusBadRequest
			res, _ := json.Marshal(createBookResponse)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)
			return
		}

		createdbook, err := bookService.CreateBook(book)
		if err != nil {
			createBookResponse.Message = err.Error()
			createBookResponse.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(createBookResponse)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)
			return
		}
		createBookResponse.ID = createdbook.BookID
		createBookResponse.Message = "Book Created successfully."
		createBookResponse.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(createBookResponse)
		w.Write(res)
	}
}

func DeleteBookHandler(bookService service.BookServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var DeleteBookResponse models.BookResponse
		params := mux.Vars(r)

		// convert the id type from string to int
		bid, err := strconv.Atoi(params["book_id"])

		if err != nil {
			log.Print(err)
			DeleteBookResponse.Message = "Unable to convert the string bookid into int bookid"
			DeleteBookResponse.StatusCode = http.StatusBadRequest
			res, _ := json.Marshal(DeleteBookResponse)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)
			return
		}
		id, err := bookService.DeleteBook(bid)

		if err != nil {
			DeleteBookResponse.Message = err.Error()
			DeleteBookResponse.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(DeleteBookResponse)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}
		DeleteBookResponse.ID = id
		DeleteBookResponse.Message = "Book Deleted successfully."
		DeleteBookResponse.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(DeleteBookResponse)
		w.Write(res)
	}
}
