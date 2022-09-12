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

func IssueBookHandler(bookReportService service.BookReportServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var issueBookResponse models.BookReportResponse
		params := mux.Vars(r)

		// convert the id type from string to int
		uid, err := strconv.Atoi(params["user_id"])
		if err != nil {
			log.Fatalf("Unable to convert the string into int.  %v", err)
		}
		bid, err := strconv.Atoi(params["book_id"])
		if err != nil {
			log.Fatalf("Unable to convert the string into int.  %v", err)
		}
		id, err := bookReportService.IssueBook(uid, bid)

		if err != nil {
			fmt.Sprintln("error....")
			issueBookResponse.Message = "error in issueing book"
			issueBookResponse.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(issueBookResponse)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}
		issueBookResponse.ID = id
		issueBookResponse.Message = "Book Issued successfully."
		issueBookResponse.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(issueBookResponse)
		w.Write(res)
	}
}

func GetBookReportHandler(bookReportService service.BookReportServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var bookReportListResponse models.BookReportListResponse
		params := mux.Vars(r)

		// convert the id type from string to int
		uid, err := strconv.Atoi(params["user_id"])
		if err != nil {
			log.Fatalf("Unable to convert the string into int.  %v", err)
		}

		reports, err := bookReportService.GetBookReport(uid)

		if err != nil {
			fmt.Sprintln("error....")
			bookReportListResponse.Message = "error in issueing book"
			bookReportListResponse.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(bookReportListResponse)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}
		bookReportListResponse.BookReportList = reports
		bookReportListResponse.Message = "Book Report fetched successfully."
		bookReportListResponse.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(bookReportListResponse)
		w.Write(res)
	}
}

func ReturnBookHandler(bookReportService service.BookReportServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var returnBookResponse models.ReturnBookResponse
		params := mux.Vars(r)

		// convert the id type from string to int
		uid, err := strconv.Atoi(params["user_id"])
		if err != nil {
			log.Fatalf("Unable to convert the string into int.  %v", err)
		}
		bid, err := strconv.Atoi(params["book_id"])
		if err != nil {
			log.Fatalf("Unable to convert the string into int.  %v", err)
		}
		book, err := bookReportService.ReturnBook(uid, bid)

		if err != nil {
			fmt.Sprintln("error....")
			returnBookResponse.Message = "error in Returning book"
			returnBookResponse.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(returnBookResponse)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}
		returnBookResponse.BookReportList = book
		returnBookResponse.Message = "Book Returned successfully."
		returnBookResponse.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(returnBookResponse)
		w.Write(res)
	}
}
