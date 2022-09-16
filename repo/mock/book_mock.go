package mock

import (
	"project/models"

	"github.com/stretchr/testify/mock"
)

type MockRepositry struct {
	mock.Mock
}

func (mock *MockRepositry) CreateBook(book models.Book) (createdBook models.Book, err error) {
	args := mock.Called(book)
	result := args.Get(0).(models.Book)
	return result, args.Error(1)
}
func (mock *MockRepositry) ListBooks() (bookLists []models.BookList, err error) {
	args := mock.Called(bookLists)
	result := args.Get(0).([]models.BookList)

	return result, args.Error(1)
}
func (mock *MockRepositry) DeleteBook(bid int) (id int, err error) {
	args := mock.Called(bid)

	return args.Int(0), args.Error(1)
}
