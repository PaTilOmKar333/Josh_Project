package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func BookToBookList(b Book, bs BookStatus) (bl BookList) {

	bl.BookID = b.BookID
	bl.BookName = b.BookName
	bl.AuthorName = b.AuthorName
	bl.Status = bs.Status
	bl.AvailableCopies = b.AvailableCopies
	return
}

func UserToUserList(u User, ur UserRole) (ul UserList) {
	ul.User_ID = u.User_ID
	ul.FirstName = u.FirstName
	ul.LastName = u.LastName
	ul.Email = u.Email
	ul.Age = u.Age
	ul.Address = u.Address
	ul.Role = ur.Role
	return
}

func BookReportToBookReportList(br BookReport, b Book, u User) (brl BookReportList) {

	actualReturnDate := br.ActualReturnDate

	//	actualReturnDate.IsZero)

	//	actualReturnDate.String()

	brl.BookName = b.BookName
	brl.UserName = u.FirstName
	brl.BookReportID = br.BookReportID
	brl.IssueDate = br.IssueDate
	brl.ReturnDate = br.ReturnDate
	if actualReturnDate != nil {
		brl.BookStatus = "Book Returned"
		brl.ActualReturnDate = *actualReturnDate
	} else {
		brl.BookStatus = "Book Issued"
	}

	return
}
func ReturnBookReportfunc(u User, b Book, br BookReport) (rbr BookReportList) {

	actualReturnDate := br.ActualReturnDate

	rbr.BookReportID = br.BookReportID
	rbr.BookName = b.BookName
	rbr.UserName = u.FirstName
	rbr.IssueDate = br.IssueDate
	rbr.ReturnDate = br.ReturnDate
	if actualReturnDate != nil {
		rbr.ActualReturnDate = *br.ActualReturnDate
	}

	return
}

func UpdatedUserDetails(ou User, nu User) (uu UpdateUser) {
	uu.OldFirstName = ou.FirstName
	uu.OldLastName = ou.LastName
	uu.OldPassword = ou.Password
	uu.NewFirstName = nu.FirstName
	uu.NewLastName = nu.LastName
	uu.NewPassword = nu.Password
	return
}

func GenerateJWT(email, role string) (string, error) {
	var mySigningKey = []byte("secret_key")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		//fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
