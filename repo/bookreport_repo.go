//struct and methods

package repo

import (
	"database/sql"
	"errors"
	"log"
	"project/app"
	"project/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type BookReportRepoInterface interface {
	IssueBook(uid, bid int) (id int, err error)
	GetBookReport(uid int) (bookReportLists []models.BookReportList, err error)
	GetAllBookReport() (bookReportLists []models.BookReportList, err error)
	ReturnBook(uid, bid int) (ReturnBookReport models.BookReportList, err error)
}

type bookreportRepo struct {
	db *sqlx.DB
}

func InitBookReportRepo() BookReportRepoInterface {
	var brr bookreportRepo
	brr.db = app.GetDB()
	return &brr
}

func (brr *bookreportRepo) IssueBook(uid, bid int) (id int, err error) {
	var bookreport, checkbookreport models.BookReport
	var book models.Book
	var selectuser models.User

	sqlStatement1 := `select user_id, first_name, last_name, age, email, role_id, password, address FROM users where user_id=$1 `
	err = brr.db.Get(&selectuser, sqlStatement1, uid)

	if err == sql.ErrNoRows {
		log.Println(err)
		err = errors.New("issuing to wrong user ID.user with provided ID is not present in database")
		return
	} else if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in fetching user. we are working on this")
		return
	}

	sqlStatement2 := `select book_id,available_book_copies from books WHERE book_id=$1`
	err = brr.db.Get(&book, sqlStatement2, bid)
	if err == sql.ErrNoRows {
		log.Println(err)
		err = errors.New("issuing wrong book ID. book with provided ID is not present in database")
		return
	} else if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in issueing book. we are working on this")
		return
	}
	availabaleCopies := book.AvailableCopies
	if availabaleCopies < 1 {
		err = errors.New("sorry for inconvenience, book is unavailable now")
		return
	}
	//select bkreport_id from book_report WHERE book_id=25 and user_id=4 and actual_retun_date IS NULL

	sqlStatement6 := `select bkreport_id from book_report WHERE book_id=$1 and user_id=$2 and actual_retun_date IS NULL`
	err = brr.db.Get(&checkbookreport, sqlStatement6, bid, uid)
	if err == sql.ErrNoRows {
		if availabaleCopies >= 1 {
			sqlStatement3 := `INSERT INTO book_report(book_id, user_id, issue_date, return_date) VALUES ($1, $2, $3, $4) RETURNING bkreport_id`

			err = brr.db.Get(&bookreport, sqlStatement3, bid, uid, time.Now().Format("01-02-2006"), time.Now().AddDate(0, 0, 7))
			if err != nil {
				log.Println(err)
				err = errors.New("sorry for inconvenience, there is error in issueing book. we are working on this")
				return
			}
			id = bookreport.BookReportID

			sqlStatement4 := `UPDATE books SET available_book_copies=$2 WHERE book_id=$1 RETURNING book_id`
			err = brr.db.Get(&book, sqlStatement4, bid, availabaleCopies-1)
			if err != nil {
				log.Println(err)
				err = errors.New("sorry for inconvenience, there is error in issueing book. we are working on this")
				return
			}
			availabaleCopies = availabaleCopies - 1

			if availabaleCopies == 0 {
				sqlStatement5 := `UPDATE books SET status_id=2 WHERE book_id=$1 RETURNING book_id`
				err = brr.db.Get(&book, sqlStatement5, bid)
				if err != nil {
					log.Println(err)
					err = errors.New("sorry for inconvenience, there is error in issueing book. we are working on this")
					return
				}
			}
		}

	} else if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in issueing book. we are working on this")
		return
	} else {
		err = errors.New("same book is already assign to you")
		return
	}

	return
}

func (brr *bookreportRepo) GetAllBookReport() (bookReportLists []models.BookReportList, err error) {

	var bookReports []models.BookReport

	sqlStatement := `select * FROM book_report`
	err = brr.db.Select(&bookReports, sqlStatement)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in fetching book report. we are working on this")
		return
	}

	for _, bookReport := range bookReports {
		var book models.Book
		var user models.User

		sqlStatement2 := `select * from books where book_id=$1`
		err = brr.db.Get(&book, sqlStatement2, bookReport.BookID)
		if err != nil {
			log.Println(err)
			err = errors.New("sorry for inconvenience, there is error in fetching book. we are working on this")
			return
		}

		sqlStatement3 := `select * from users where user_id=$1`
		err = brr.db.Get(&user, sqlStatement3, bookReport.UserID)
		if err != nil {
			log.Println(err)
			err = errors.New("sorry for inconvenience, there is error in fetching book. we are working on this")
			return
		}

		bookReportList := models.BookReportToBookReportList(bookReport, book, user)
		bookReportLists = append(bookReportLists, bookReportList)

	}
	return
}

func (brr *bookreportRepo) GetBookReport(uid int) (bookReportLists []models.BookReportList, err error) {

	var bookReports []models.BookReport

	sqlStatement := `select * FROM book_report where user_id=$1 `
	err = brr.db.Select(&bookReports, sqlStatement, uid)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in fetching book report. we are working on this")
		return
	}
	for _, bookReport := range bookReports {
		var book models.Book
		var user models.User

		sqlStatement2 := `select * from books where book_id=$1`
		err = brr.db.Get(&book, sqlStatement2, bookReport.BookID)
		if err != nil {
			log.Println(err)
			err = errors.New("sorry for inconvenience, there is error in fetching book. we are working on this")
			return
		}

		sqlStatement3 := `select * from users where user_id=$1`
		err = brr.db.Get(&user, sqlStatement3, bookReport.UserID)
		if err != nil {
			log.Println(err)
			err = errors.New("sorry for inconvenience, there is error in fetching book. we are working on this")
			return
		}

		bookReportList := models.BookReportToBookReportList(bookReport, book, user)
		bookReportLists = append(bookReportLists, bookReportList)

	}
	return
}

func (brr *bookreportRepo) ReturnBook(uid, bid int) (ReturnBookReport models.BookReportList, err error) {
	var bookReport models.BookReport
	var bookReport1, bookReport2 models.BookReport
	var book models.Book
	var user models.User
	var selectuser models.User
	var selectbook models.Book

	sqlStatement1 := `select user_id FROM users where user_id=$1 `
	err = brr.db.Get(&selectuser, sqlStatement1, uid)

	if err == sql.ErrNoRows {
		log.Println(err)
		err = errors.New("book returning to wrong user ID.user with provided ID is not present in database")
		return
	} else if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in fetching user. we are working on this")
		return
	}

	sqlStatement2 := `select available_book_copies from books WHERE book_id=$1`
	err = brr.db.Get(&selectbook, sqlStatement2, bid)
	if err == sql.ErrNoRows {
		log.Println(err)
		err = errors.New("returning wrong book. book with provided ID is not present in database")
		return
	} else if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in returning book. we are working on this")
		return
	}

	sqlStatement8 := `select bkreport_id from book_report WHERE book_id=$1 and user_id=$2 and actual_retun_date IS NULL`
	err = brr.db.Get(&bookReport2, sqlStatement8, bid, uid)
	if err == sql.ErrNoRows {
		log.Println(err)
		err = errors.New("there is no book pending from your side")
		return
	} else if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in returning book. we are working on this")
		return
	}

	sqlStatement3 := `UPDATE book_report SET actual_retun_date=$3 WHERE book_id=$1 AND user_id=$2 RETURNING bkreport_id`
	err = brr.db.Get(&bookReport1, sqlStatement3, bid, uid, time.Now().Format("01-02-2006"))
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in returning book. we are working on this")
		return
	}
	bookReportID := bookReport1.BookReportID

	sqlStatement4 := `select book_name,available_book_copies from books WHERE book_id=$1`
	err = brr.db.Get(&book, sqlStatement4, bid)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in returning book. we are working on this")
		return
	}
	availabaleCopies := book.AvailableCopies
	if availabaleCopies == 0 {
		sqlStatement8 := `UPDATE books SET available_book_copies=$2,status_id=$3 WHERE book_id=$1 RETURNING book_id`
		err = brr.db.Get(&book, sqlStatement8, bid, availabaleCopies+1, 1)
		if err != nil {
			log.Println(err)
			err = errors.New("sorry for inconvenience, there is error in returning book. we are working on this")
			return
		}
	} else if availabaleCopies >= 1 {
		sqlStatement5 := `UPDATE books SET available_book_copies=$2 WHERE book_id=$1 RETURNING book_id`
		err = brr.db.Get(&book, sqlStatement5, bid, availabaleCopies+1)
		if err != nil {
			log.Println(err)
			err = errors.New("sorry for inconvenience, there is error in returning book. we are working on this")
			return
		}
	}

	sqlStatement6 := `select first_name from users WHERE user_id=$1`
	err = brr.db.Get(&user, sqlStatement6, uid)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in returning book. we are working on this")
		return
	}

	sqlStatement7 := `select * from book_report WHERE bkreport_id=$1`
	err = brr.db.Get(&bookReport, sqlStatement7, bookReportID)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in returning book. we are working on this")
		return
	}

	ReturnBookReport = models.ReturnBookReportfunc(user, book, bookReport)
	return
}
