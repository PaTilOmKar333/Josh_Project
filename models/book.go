package models

// User struct declaration
type BookStatus struct {
	StatusID int    `json:"status_id,omitempty" db:"status_id"`
	Status   string `json:"status" db:"status"`
}

type Book struct {
	BookID          int    `json:"book_id" db:"book_id"`
	BookName        string `json:"book_name" db:"book_name"`
	AuthorName      string `json:"author_name" db:"author_name"`
	AvailableCopies int    `json:"no_of_copies_available" db:"available_book_copies"`
	BookStatusID    int    `json:"status_id" db:"status_id"`
}

type BookList struct {
	BookID          int    `json:"book_id"`
	BookName        string `json:"book_name"`
	AuthorName      string `json:"author_name"`
	AvailableCopies int    `json:"no_of_copies_available"`
	Status          string `json:"status"`
}

type BookListResponse struct {
	BookList   []BookList `json:"book_list,omitempty"`
	StatusCode int        `json:"status_code"`
	ErrorMsg   string     `json:"error_msg,omitempty"`
}

type CretaeBookResponse struct {
	ID         int    `json:"book_id,omitempty"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
