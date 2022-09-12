package handler

import (
	"project/repo"
	"project/service"
)

type dependancies struct {
	Userservice       service.UserServiceInterface
	Bookservice       service.BookServiceInterface
	BookReportService service.BookReportServiceInterface
}

var depnd dependancies

func InitDependancies() {
	userrepo := repo.InitUserRepo()
	userservice := service.InitUserService(userrepo)
	depnd.Userservice = userservice

	bookrepo := repo.InitBookRepo()
	bookservice := service.InitBookService(bookrepo)
	depnd.Bookservice = bookservice

	bookreportrepo := repo.InitBookReportRepo()
	bookreportservice := service.InitBookReportService(bookreportrepo)
	depnd.BookReportService = bookreportservice
}
