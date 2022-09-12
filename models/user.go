package models

import "github.com/golang-jwt/jwt"

// User struct declaration
type User struct {
	User_ID   int    `json:"u_id" db:"user_id"`
	FirstName string `json:"u_firstname" db:"first_name"`
	LastName  string `json:"u_lastname" db:"last_name"`
	Age       int    `json:"age" db:"age"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password,omitempty" db:"password"`
	Address   string `json:"address" db:"address"`
	Role_ID   int    `json:"role_id,omitempty" db:"role_id"`
}

type UpdateUser struct {
	OldFirstName string `json:"old_firstname,omitempty"`
	OldLastName  string `json:"old_lastname,omitempty"`
	OldPassword  string `json:"old_password,omitempty"`
	NewFirstName string `json:"new_firstname,omitempty"`
	NewLastName  string `json:"new_lastname,omitempty"`
	NewPassword  string `json:"new_password,omitempty"`
}

type CreateUser struct {
	UserID    int    `json:"user_id"`
	FirstName string `json:"u_firstname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserRole struct {
	RoleID int    `json:"role_id" db:"role_id"`
	Role   string `json:"role" db:"role_name"`
}

type UserList struct {
	User_ID   int    `json:"u_id"`
	FirstName string `json:"u_firstname"`
	LastName  string `json:"u_lastname"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Role      string `json:"role"`
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	UserID      int    `json:"u_id"`
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
	jwt.StandardClaims
}

type CretaeUserResponse struct {
	CreateUser CreateUser `json:"created_user_details"`
	Message    string     `json:"message"`
	StatusCode int        `json:"status_code"`
}

type UserListResponse struct {
	UserList   []UserList `json:"user_list,omitempty"`
	Message    string     `json:"error_msg,omitempty"`
	StatusCode int        `json:"status_code"`
}

type UpdateUserResponse struct {
	UpdatedUser UpdateUser `json:"updated_user_details,omitempty"`
	StatusCode  int        `json:"status_code"`
	Message     string     `json:"error_msg,omitempty"`
}

type GetUserResponse struct {
	GotUser    User   `json:"user_list,omitempty"`
	StatusCode int    `json:"status_code"`
	ErrorMsg   string `json:"error_msg,omitempty"`
}
