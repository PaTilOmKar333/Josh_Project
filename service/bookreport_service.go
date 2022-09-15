//struct and methods

package service

import (
	"context"
	"errors"
	"project/models"
	"project/repo"
)

type BookReportServiceInterface interface {
	IssueBook(uid, bid int) (id int, err error)
	GetBookReport(ctx context.Context, uid int) (bookReportLists []models.BookReportList, err error)
	ReturnBook(uid, bid int) (BookReport models.BookReportList, err error)
}

type bookreportService struct {
	repo repo.BookReportRepoInterface
}

func InitBookReportService(r repo.BookReportRepoInterface) BookReportServiceInterface {

	//initialies
	//repo.InitUserRepo()
	return &bookreportService{
		repo: r,
	}
}

func (brs *bookreportService) IssueBook(uid, bid int) (id int, err error) {
	id, err = brs.repo.IssueBook(uid, bid)
	if err != nil {
		return
	}
	return
}

func (brs *bookreportService) GetBookReport(ctx context.Context, uid int) (bookReportLists []models.BookReportList, err error) {

	val, _ := ctx.Value("ClaimsToVerify").(*models.Claims)
	if val.UserID == uid || val.Role == "admin" || val.Role == "superadmin" {
		bookReportLists, err = brs.repo.GetBookReport(uid)
		if err != nil {
			return
		}
	} else {
		err = errors.New("unauthorized user")

	}

	return
}

func (brs *bookreportService) ReturnBook(uid, bid int) (BookReport models.BookReportList, err error) {
	BookReport, err = brs.repo.ReturnBook(uid, bid)
	if err != nil {
		return
	}
	return
}
