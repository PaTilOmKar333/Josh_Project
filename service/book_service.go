//struct and methods

package service

import (
	"project/models"
	"project/repo"
)

type BookServiceInterface interface {
	ListBooks() (bookslist []models.BookList, err error)
	CreateBook(book models.Book) (createdBook models.Book, err error)
	DeleteBook(bid int) (id int, err error)
}

type bookService struct {
	repo repo.BookRepoInterface
}

func InitBookService(r repo.BookRepoInterface) BookServiceInterface {

	//initialies
	//repo.InitUserRepo()
	return &bookService{
		repo: r,
	}
}

func (bs *bookService) ListBooks() (bookslist []models.BookList, err error) {
	bookslist, err = bs.repo.ListBooks()
	//fmt.Println("service layer: ", users)
	if err != nil {
		return
	}
	return
}

func (bs *bookService) CreateBook(book models.Book) (createdBook models.Book, err error) {
	createdBook, err = bs.repo.CreateBook(book)
	if err != nil {
		return
	}
	return
}

func (bs *bookService) DeleteBook(bid int) (id int, err error) {
	id, err = bs.repo.DeleteBook(bid)
	if err != nil {
		return
	}
	return
}
