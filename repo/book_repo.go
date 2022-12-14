//struct and methods

package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"project/app"
	"project/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

type BookRepoInterface interface {
	ListBooks() (bookLists []models.BookList, err error)
	CreateBook(book models.Book) (createdBook models.Book, err error)
	DeleteBook(bid int) (id int, err error)
}

type bookRepo struct {
	db *sqlx.DB
}

func InitBookRepo() BookRepoInterface {

	var br bookRepo
	br.db = app.GetDB()
	return &br
}

func (br *bookRepo) ListBooks() (bookLists []models.BookList, err error) {
	var books []models.Book
	//var book models.Book

	sqlStatement1 := `select * from books`

	err = br.db.Select(&books, sqlStatement1)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in fetching list of books. we are working on this")
		return
	}
	for _, book := range books {
		var booksStatus models.BookStatus
		sqlStatement2 := `select * from books_status where status_id=$1`
		err = br.db.Get(&booksStatus, sqlStatement2, book.BookStatusID)
		if err != nil {
			log.Println(err)
			err = errors.New("sorry for inconvenience, there is error in fetching list of books. we are working on this")
			return
		}
		bookList := models.BookToBookList(book, booksStatus)
		bookLists = append(bookLists, bookList)
	}

	return
}

func (br *bookRepo) CreateBook(book models.Book) (createdBook models.Book, err error) {
	//var createdBook models.Book

	sqlStatement := `INSERT INTO books( book_name, author_name, available_book_copies,Status_id) VALUES ($1, $2, $3, $4) RETURNING book_id`
	//var id int
	err = br.db.Get(&book, sqlStatement, book.BookName, book.AuthorName, book.AvailableCopies, 1)
	if err != nil {
		if err != nil {

			errorstring := err.Error()

			if strings.Contains(errorstring, "constraint_book_name") {
				fmt.Println("errorstring:", errorstring)
				err = errors.New("book is already exist.please enter new book")
				return
			} else {
				log.Println(err)
				err = errors.New("sorry for inconvenience, there is error in creating new of book. we are working on this")
				return
			}

		}
	}
	id := book.BookID

	sqlStatement1 := `select * from books WHERE book_id=$1`
	err = br.db.Get(&createdBook, sqlStatement1, id)

	if err == sql.ErrNoRows {
		log.Println(err)
		err = errors.New("book with provided ID is not present in database")
		return
	} else if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in fetching user. we are working on this")
		return
	}
	//fmt.Printf("inserted single record %v", id)

	return
}

func (br *bookRepo) DeleteBook(bid int) (id int, err error) {
	var selectbook models.Book
	var deletebook models.Book

	sqlStatement1 := `select * from books WHERE book_id=$1`
	err = br.db.Get(&selectbook, sqlStatement1, bid)

	if err == sql.ErrNoRows {
		log.Println(err)
		err = errors.New("book with provided ID is not present in database")
		return
	} else if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in fetching user. we are working on this")
		return
	}

	sqlStatement2 := `DELETE FROM books WHERE book_id=$1 `
	br.db.Get(&deletebook, sqlStatement2, bid)

	fmt.Printf("Book Deleted Successfully %v", bid)
	id = bid
	return
}
