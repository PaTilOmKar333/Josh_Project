//struct and methods

package repo

import (
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
	var bookreport models.BookReport
	var book models.Book
	sqlStatement1 := `select available_book_copies from books WHERE book_id=$1`
	err = brr.db.Get(&book, sqlStatement1, bid)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in issueing book. we are working on this")
		return
	}
	availabaleCopies := book.AvailableCopies
	if availabaleCopies < 1 {
		err = errors.New("sorry for inconvenience, book is unavailable now")
		return
	}

	if availabaleCopies >= 1 {
		sqlStatement2 := `INSERT INTO book_report(book_id, user_id, issue_date, return_date) VALUES ($1, $2, $3, $4) RETURNING bkreport_id`

		err = brr.db.Get(&bookreport, sqlStatement2, bid, uid, time.Now().Format("01-02-2006"), time.Now().AddDate(0, 0, 7))
		if err != nil {
			log.Println(err)
			err = errors.New("sorry for inconvenience, there is error in issueing book. we are working on this")
			return
		}
		id = bookreport.BookReportID

		sqlStatement3 := `UPDATE books SET available_book_copies=$2 WHERE book_id=$1 RETURNING book_id`
		err = brr.db.Get(&book, sqlStatement3, bid, availabaleCopies-1)
		if err != nil {
			log.Println(err)
			err = errors.New("sorry for inconvenience, there is error in issueing book. we are working on this")
			return
		}

		if availabaleCopies-2 == 0 {
			sqlStatement4 := `UPDATE books SET status_id=2 WHERE book_id=$1 RETURNING book_id`
			err = brr.db.Get(&book, sqlStatement4, bid)
			if err != nil {
				log.Println(err)
				err = errors.New("sorry for inconvenience, there is error in issueing book. we are working on this")
				return
			}
		}
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
			err = errors.New("sorry for inconvenience, there is error in issueing book. we are working on this")
			return
		}

		sqlStatement3 := `select * from users where user_id=$1`
		err = brr.db.Get(&user, sqlStatement3, bookReport.UserID)
		if err != nil {
			log.Println(err)
			err = errors.New("sorry for inconvenience, there is error in issueing book. we are working on this")
			return
		}

		bookReportList := models.BookReportToBookReportList(bookReport, book, user)
		bookReportLists = append(bookReportLists, bookReportList)

	}
	return
}

func (brr *bookreportRepo) ReturnBook(uid, bid int) (ReturnBookReport models.BookReportList, err error) {
	var bookReport models.BookReport
	var bookReport1 models.BookReport
	var book models.Book
	var user models.User

	sqlStatement := `UPDATE book_report SET actual_return_date=$3 WHERE book_id=$1 AND user_id=$2 RETURNING bkreport_id`
	err = brr.db.Get(&bookReport1, sqlStatement, bid, uid, time.Now().Format("01-02-2006"))
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in returning book. we are working on this")
		return
	}
	bookReportID := bookReport1.BookReportID

	sqlStatement1 := `select book_name,available_book_copies from books WHERE book_id=$1`
	err = brr.db.Get(&book, sqlStatement1, bid)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in returning book. we are working on this")
		return
	}
	availabaleCopies := book.AvailableCopies

	sqlStatement2 := `UPDATE books SET available_book_copies=$2 WHERE book_id=$1 RETURNING book_id`
	err = brr.db.Get(&book, sqlStatement2, bid, availabaleCopies+1)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in returning book. we are working on this")
		return
	}

	sqlStatement3 := `select first_name from users WHERE user_id=$1`
	err = brr.db.Get(&user, sqlStatement3, uid)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in returning book. we are working on this")
		return
	}

	sqlStatement4 := `select * from book_report WHERE bkreport_id=$1`
	err = brr.db.Get(&bookReport, sqlStatement4, bookReportID)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in returning book. we are working on this")
		return
	}

	ReturnBookReport = models.ReturnBookReportfunc(user, book, bookReport)
	return
}
