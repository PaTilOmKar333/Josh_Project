package handler

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	//r := mux.NewRouter()

	//r.HandleFunc("/users", LoginHandler(depnd.Userservice)).Methods("GET")
	router := mux.NewRouter()

	//router.HandleFunc("/login/", LoginHandler(depnd.Userservice)).Methods("GET")
	router.HandleFunc("/users/", AllUsersHandler(depnd.Userservice)).Methods("GET")
	router.HandleFunc("/user/", CreateUsersHandler(depnd.Userservice)).Methods("POST")
	//	router.HandleFunc("/user/{user_id}", GetUsersByIDHandler(depnd.Userservice)).Methods("GET")
	router.HandleFunc("/user", GetUsersByEmailOrIDHandler(depnd.Userservice)).Methods("GET")
	//router.HandleFunc("/user", GetUsersByEmailOrIDHandler(depnd.Userservice)).Methods("GET")
	//router.HandleFunc("/userr/{email_id}/", GetUsersByEmailHandler(depnd.Userservice)).Methods("GET")
	//router.HandleFunc("/userr/{id}/", GetUsersByEmailOrIDHandler(depnd.Userservice)).Methods("GET")
	router.HandleFunc("/user/{email_id}/", UpdateUserHandler(depnd.Userservice)).Methods("PUT")
	router.HandleFunc("/user/{user_id}/", DeleteUserHandler(depnd.Userservice)).Methods("DELETE")
	router.HandleFunc("/book/", AllBooksHandler(depnd.Bookservice)).Methods("GET")
	router.HandleFunc("/book/", CreateBooksHandler(depnd.Bookservice)).Methods("POST")
	router.HandleFunc("/book/{book_id}/", DeleteBookHandler(depnd.Bookservice)).Methods("DELETE")
	router.HandleFunc("/user/{user_id}/book/{book_id}/", IssueBookHandler(depnd.BookReportService)).Methods("POST")
	router.HandleFunc("/user/{user_id}/book/", GetBookReportHandler(depnd.BookReportService)).Methods("GET")
	router.HandleFunc("/user/{user_id}/book/{book_id}/return", ReturnBookHandler(depnd.BookReportService)).Methods("POST")
	//	/user/{user_id}/book/{book_id}/return
	return router
}
