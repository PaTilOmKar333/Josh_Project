//struct and methods

package repo

import (
	"errors"
	"fmt"
	"log"
	"project/app"
	"project/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

type UserRepoInterface interface {
	//user methods
	ListUser() (usersList []models.UserList, err error)
	CreateUser(user models.User) (createduser models.User, err error)
	//GetUser(uid int) (user models.User, err error)
	//	GetUser(param interface{}) (user models.User, err error)
	UpdateUser(email string, user models.User) (updateUser models.UpdateUser, err error)
	GetUserByEmail(email string) (user models.User, err error)
	DeleteUser(uid int) (id int, err error)
}

type userRepo struct {
	db *sqlx.DB
}

func InitUserRepo() UserRepoInterface {
	//var err error
	var ur userRepo
	ur.db = app.GetDB()
	// return object
	return &ur
}

// func (ur *userRepo) login(authdetails models.Authentication) (err error) {
// 	var authuser models.User
// 	ur.db.Where("email = ?", authdetails.Email).First(&authuser)
// 	if authuser.Email == "" {
// 		err = SetError(err, "Username or Password is incorrect")
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(err)
// 		return
// 	}

// 	check := CheckPasswordHash(authdetails.Password, authuser.Password)

// 	if !check {
// 		var err Error
// 		err = SetError(err, "Username or Password is incorrect")
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(err)
// 		return
// 	}

// 	validToken, err := GenerateJWT(authuser.Email, authuser.Role)
// 	if err != nil {
// 		var err Error
// 		err = SetError(err, "Failed to generate token")
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(err)
// 		return
// 	}

// 	var token Token
// 	token.Email = authuser.Email
// 	token.Role = authuser.Role
// 	token.TokenString = validToken
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(token)
// }

func (ur *userRepo) ListUser() (usersList []models.UserList, err error) {
	var users []models.User
	sqlStatement := `SELECT * FROM users`

	//rows, err := ur.db.Query(sqlStatement)
	err = ur.db.Select(&users, sqlStatement)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in fetching list of users. we are working on this")
		return
	}

	for _, user := range users {
		var userRole models.UserRole
		sqlStatement1 := `select * from roles where role_id=$1`
		err = ur.db.Get(&userRole, sqlStatement1, user.Role_ID)
		if err != nil {
			log.Println(err)
			err = errors.New("sorry for inconvenience, there is error in fetching list of users. we are working on this")
			return
		}
		userList := models.UserToUserList(user, userRole)
		usersList = append(usersList, userList)
	}

	// defer rows.Close()

	// for rows.Next() {
	// 	//err = rows.Scan(&user.User_ID, &user.FirstName, &user.LastName, &user.Age, &user.Email, &user.Password, &user.Address, &user.Role_ID)
	// 	err = rows.Scan(&user.User_ID, &user.FirstName, &user.LastName)
	// 	if err != nil {
	// 		log.Println(err)
	// 		err = errors.New("sorry for inconvenience, there is error in fetching list of users. we are working on this")
	// 		return
	// 	}
	// 	users = append(users, user)
	// }
	//	fmt.Println("repo layer: ", users)

	return
}

func (ur *userRepo) CreateUser(user models.User) (createduser models.User, err error) {

	sqlStatement := `INSERT INTO users(first_name, last_name, age, email, password, address, role_id) VALUES ($1, $2, $3,$4, $5, $6, $7) RETURNING user_id`
	//var id int
	//err = ur.db.QueryRow(sqlStatement, user.FirstName, user.LastName, user.Age, user.Email, user.Password, user.Address, user.Role_ID)
	err = ur.db.Get(&user, sqlStatement, user.FirstName, user.LastName, user.Age, user.Email, user.Password, user.Address, user.Role_ID)
	if err != nil {
		// log.Println(err)
		// err = errors.New("user is already exists with same emailid. please use another email to create user")
		// return
		errorstring := err.Error()

		if strings.Contains(errorstring, "constraint_email_unique") {
			fmt.Println("errorstring:", errorstring)
			err = errors.New("user is already exists with same emailid. please use another email to create user")
			return
		} else {
			log.Println(err)
			err = errors.New("sorry for inconvenience, there is error in creating user. we are working on this")
			return
		}

	}
	user_id := user.User_ID
	sqlStatement1 := `select * FROM users WHERE user_id=$1`
	err = ur.db.Get(&createduser, sqlStatement1, user_id)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in creating user. we are working on this")
		return
	}
	return
}

// func (ur *userRepo) GetUser(uid int) (user models.User, err error) {

// 	sqlStatement := `select user_id, first_name, last_name, age, email, address FROM users where user_id=$1 `
// 	//var id int
// 	//err = ur.db.QueryRow(sqlStatement, user.User_ID, user.FirstName, user.LastName, user.Age, user.Email, user.Password, user.Address, user.Role_ID).Scan(&id)
// 	err = ur.db.Get(&user, sqlStatement, uid)
// 	if err != nil {
// 		log.Println(err)
// 		err = errors.New("sorry for inconvenience, there is error in fetching user. we are working on this")
// 		return
// 	}
// 	return
// }

// func (ur *userRepo) GetUser(param interface{}) (user models.User, err error) {

// 	sqlStatement := `select user_id, first_name, last_name, age, email, address FROM users where user_id=$1 `
// 	//var id int
// 	//err = ur.db.QueryRow(sqlStatement, user.User_ID, user.FirstName, user.LastName, user.Age, user.Email, user.Password, user.Address, user.Role_ID).Scan(&id)
// 	err = ur.db.Get(&user, sqlStatement, uid)
// 	if err != nil {
// 		log.Println(err)
// 		err = errors.New("sorry for inconvenience, there is error in fetching user. we are working on this")
// 		return
// 	}
// 	return
// }

func (ur *userRepo) UpdateUser(email string, user models.User) (updateUser models.UpdateUser, err error) {
	var newuser, olduser models.User

	sqlStatement1 := `select * FROM users WHERE email=$1`
	err = ur.db.Get(&olduser, sqlStatement1, email)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in updating user. we are working on this")
		return
	}

	sqlStatement2 := `UPDATE users SET first_name=$2, last_name=$3, password=$4 WHERE email=$1 RETURNING user_id`

	err = ur.db.Get(&user, sqlStatement2, email, user.FirstName, user.LastName, user.Password)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in updating user. we are working on this")
		return
	}

	id := user.User_ID
	sqlStatement3 := `select * FROM users where user_id=$1 `
	err = ur.db.Get(&newuser, sqlStatement3, id)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in updating user. we are working on this")
		return
	}

	updateUser = models.UpdatedUserDetails(olduser, newuser)
	//	fmt.Printf("user updated successfully,%v", id)
	return
}

func (ur *userRepo) GetUserByEmail(email string) (user models.User, err error) {

	sqlStatement := `select user_id, first_name, last_name, age, email, address FROM users where email=$1 `
	//var id int
	//err = ur.db.QueryRow(sqlStatement, user.User_ID, user.FirstName, user.LastName, user.Age, user.Email, user.Password, user.Address, user.Role_ID).Scan(&id)
	err = ur.db.Get(&user, sqlStatement, email)
	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in fetching user. we are working on this")
		return
	}
	//	fmt.Printf("inserted single record %v", user.User_ID)
	return
}

func (ur *userRepo) DeleteUser(uid int) (id int, err error) {

	sqlStatement := `DELETE FROM users WHERE user_id=$1 `
	_, err = ur.db.Exec(sqlStatement, uid)

	if err != nil {
		log.Println(err)
		err = errors.New("sorry for inconvenience, there is error in deleting user. we are working on this")
		return
	}
	//fmt.Printf("UserDeleted Successfully %v", uid)
	id = uid
	return
}
