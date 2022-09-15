package models

import (
	"time"
)

type BookReport struct {
	BookReportID     int        `json:"book_report_id" db:"bkreport_id"`
	BookID           int        `json:"book_id" db:"book_id"`
	UserID           int        `json:"user_id" db:"user_id"`
	IssueDate        time.Time  `json:"issue_date" db:"issue_date"`
	ReturnDate       time.Time  `json:"return_date" db:"return_date"`
	ActualReturnDate *time.Time `json:"actual_return_date,omitempty" db:"actual_retun_date"`
}

type BookReportResponse struct {
	ID         int    `json:"book_report_id,omitempty"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type BookReportList struct {
	BookReportID     int       `json:"book_report_id,omitempty"`
	BookName         string    `json:"book_name,omitempty"`
	UserName         string    `json:"user_name,omitempty"`
	BookStatus       string    `json:"book_status,omitempty"`
	IssueDate        time.Time `json:"issue_date,omitempty"`
	ReturnDate       time.Time `json:"return_date,omitempty"`
	ActualReturnDate time.Time `json:"actual_return_date,omitempty"`
}

type BookReportListResponse struct {
	BookReportList []BookReportList `json:"book_report_list"`
	StatusCode     int              `json:"status_code"`
	Message        string           `json:"error_msg,omitempty"`
}

type ReturnBookResponse struct {
	BookReportList BookReportList `json:"book_report,omitempty"`
	StatusCode     int            `json:"status_code"`
	Message        string         `json:"error_msg,omitempty"`
}
