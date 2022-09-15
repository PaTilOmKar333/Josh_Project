package models

import "github.com/golang-jwt/jwt"

// User struct declaration
type User struct {
	User_ID   int    `json:"u_id ,omitempty" db:"user_id"`
	FirstName string `json:"u_firstname,omitempty" db:"first_name"`
	LastName  string `json:"u_lastname,omitempty" db:"last_name"`
	Age       int    `json:"age,omitempty" db:"age"`
	Email     string `json:"email,omitempty" db:"email"`
	Password  string `json:"password,omitempty" db:"password"`
	Address   string `json:"address,omitempty" db:"address"`
	Role_ID   int    `json:"role_id,omitempty" db:"role_id"`
}

type UserAuth struct {
	User_ID   int    `json:"u_id"`
	FirstName string `json:"u_firstname"`
	LastName  string `json:"u_lastname"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	Address   string `json:"address"`
	Role      string `json:"role"`
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
	UserID    int    `json:"user_id,omitempty"`
	FirstName string `json:"u_firstname,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
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
	UserID   int    `json:"u_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Claims struct {
	UserID int    `json:"u_id"`
	Role   string `json:"role"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

type TokenResponse struct {
	Token      string `json:"token,omitempty"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type CreateUserResponse struct {
	CreateUser CreateUser `json:"created_user_details,omitempty"`
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

type AuthenticationResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"error_msg,omitempty"`
}
