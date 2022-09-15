package models

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

func UserToUserAuth(u User, ur UserRole) (ul UserAuth) {
	ul.User_ID = u.User_ID
	ul.FirstName = u.FirstName
	ul.LastName = u.LastName
	ul.Email = u.Email
	ul.Age = u.Age
	ul.Address = u.Address
	ul.Password = u.Password
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

	//	issuedate := br.IssueDate

	// fmt.Println(issuedate)
	// if issuedate.Valid {
	// 	brl.IssueDate = issuedate.Time
	// 	fmt.Println(brl.IssueDate)
	// }
	// returndate := br.ReturnDate
	// if returndate.Valid {
	// 	brl.ReturnDate = returndate.Time
	// }
	if actualReturnDate != nil {
		brl.BookStatus = "Book Returned"
		brl.ActualReturnDate = *actualReturnDate
	} else {
		brl.BookStatus = "Book Issued"
	}

	return
}
func ReturnBookReportfunc(u User, b Book, br BookReport) (rbr BookReportList) {
	//actualReturnDate := br.ActualReturnDate

	actualReturnDate := br.ActualReturnDate

	rbr.BookReportID = br.BookReportID
	rbr.BookName = b.BookName
	rbr.UserName = u.FirstName

	// issuedate := br.IssueDate
	// if issuedate.Valid {
	// 	rbr.IssueDate = issuedate.Time
	// 	fmt.Println(rbr.IssueDate)
	// }
	// returndate := br.ReturnDate
	// if returndate.Valid {
	// 	rbr.ReturnDate = returndate.Time
	// }
	// if actualReturnDate.Valid {
	// 	rbr.BookStatus = "Book Returned"
	// 	rbr.ActualReturnDate = *&actualReturnDate.Time
	// } else {
	// 	rbr.BookStatus = "Book Issued"
	// }

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
