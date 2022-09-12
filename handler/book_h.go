package handler

import (
	"encoding/json"
	"fmt"
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
			fmt.Sprintln("error....")
			listBooks.ErrorMsg = "list not render"
			listBooks.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(listBooks)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}
		listBooks.BookList = books
		listBooks.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)
		//	fmt.Println("handler layer", users)

		//json.NewEncoder(w).Encode(users)
		res, _ := json.Marshal(listBooks)
		w.Write(res)

	}
}

func CreateBooksHandler(bookService service.BookServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var cretaeBookResponse models.CretaeBookResponse

		err := json.NewDecoder(r.Body).Decode(&book)

		if err != nil {
			log.Fatalf("unable to decode the request body. %v", err)
		}

		id, err := bookService.CreateBook(book)
		if err != nil {
			fmt.Sprintln("error....")
			cretaeBookResponse.Message = err.Error()
			cretaeBookResponse.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(cretaeBookResponse)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}
		cretaeBookResponse.ID = id
		cretaeBookResponse.Message = "Book Created successfully."
		cretaeBookResponse.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(cretaeBookResponse)
		w.Write(res)
	}
}

func DeleteBookHandler(bookService service.BookServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cretaeBookResponse models.CretaeBookResponse
		params := mux.Vars(r)

		// convert the id type from string to int
		bid, err := strconv.Atoi(params["book_id"])

		if err != nil {
			log.Fatalf("Unable to convert the string into int.  %v", err)
		}
		id, err := bookService.DeleteBook(bid)
		//userService.DeleteUser(uid)

		if err != nil {
			cretaeBookResponse.Message = err.Error()
			cretaeBookResponse.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(cretaeBookResponse)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}
		cretaeBookResponse.ID = id
		cretaeBookResponse.Message = "Book Deleted successfully."
		cretaeBookResponse.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(cretaeBookResponse)
		w.Write(res)
	}
}
