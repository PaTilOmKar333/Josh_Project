package service

import (
	"project/models"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestCreateBook(t *testing.T) {
	mockRepo := new(MockRepositry)
	book := models.Book{BookName: "chava", AuthorName: "omkar patil", AvailableCopies: 10, BookStatusID: 1}

	mockRepo.On("CreateBook", book).Return(book, nil)

	testService := InitBookService(mockRepo)

	result, err := testService.CreateBook(book) //.CreateBook(&book)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, 0, result.BookID)
	assert.Equal(t, "chava", result.BookName)
	assert.Equal(t, "omkar patil", result.AuthorName)
	assert.Equal(t, 10, result.AvailableCopies)
	assert.Equal(t, 1, result.BookStatusID)
	assert.Nil(t, err)

	//	result,_:=

}
