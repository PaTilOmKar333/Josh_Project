package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"project/models"
	"project/service"
	"strconv"

	"github.com/gorilla/mux"
)

// func LoginHandler(userService service.UserServiceInterface) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var authdetails models.Authentication
// 		err := json.NewDecoder(r.Body).Decode(&authdetails)
// 		if err != nil {
// 			err = errors.New("Error in reading body")
// 			return
// 		}
// 		_, err = login(authdetails)

// 	}
// }

func AllUsersHandler(userService service.UserServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var listUsers models.UserListResponse

		users, err := userService.ListUsers()
		if err != nil {
			listUsers.Message = err.Error()
			listUsers.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(listUsers)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}
		listUsers.UserList = users
		listUsers.StatusCode = http.StatusOK
		listUsers.Message = "These are all users."

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(listUsers)
		w.Write(res)

	}
}

func CreateUsersHandler(userService service.UserServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var cretaeUserResponse models.CretaeUserResponse

		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			log.Fatalf("unable to decode the request body. %v", err)
		}

		createduser, err := userService.CreateUser(user)

		if err != nil {
			fmt.Sprintln("error....")
			//	cretaeUserResponse.Message = "User is already exists with same EmailID. Please use another EmailID to create user."
			cretaeUserResponse.Message = err.Error()
			cretaeUserResponse.StatusCode = http.StatusInternalServerError
			//	cretaeUserResponse.ErrorMsg = err.Error()
			res, _ := json.Marshal(cretaeUserResponse)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}
		cretaeUserResponse.CreateUser.UserID = createduser.User_ID
		cretaeUserResponse.CreateUser.FirstName = createduser.FirstName
		cretaeUserResponse.CreateUser.Email = createduser.Email
		cretaeUserResponse.CreateUser.Password = createduser.Password
		cretaeUserResponse.Message = "User Created successfully."
		cretaeUserResponse.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)
		//	fmt.Println("handler layer", users)

		//json.NewEncoder(w).Encode(users)
		res, _ := json.Marshal(cretaeUserResponse)
		w.Write(res)
	}
}

func GetUsersByEmailOrIDHandler(userService service.UserServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var newvariable interface{}
		var getuser models.GetUserResponse
		var err error
		param := r.URL.Query().Get("filter_by")
		variable := r.URL.Query().Get(param)

		if param == "user_id" {
			newvariable, err = strconv.Atoi(variable)

			if err != nil {
				log.Println(err)
				getuser.ErrorMsg = "unable to convert the string into int"
				return
			}

		} else if param == "email_id" {
			newvariable = variable
		} else {
			_ = errors.New("wrong parameters")
		}
		user, err := userService.GetUser(newvariable)

		if err != nil {
			getuser.ErrorMsg = err.Error()
			getuser.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(getuser)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}
		getuser.GotUser = user
		getuser.StatusCode = http.StatusOK
		getuser.ErrorMsg = "user fetched successfully"

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(getuser)
		w.Write(res)

	}
}

func UpdateUserHandler(userService service.UserServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var getuser models.GetUserResponse
		var updateUserResponse models.UpdateUserResponse
		//var cretaeUserResponse models.CretaeUserResponse
		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			log.Println(err)
			return
		}
		params := mux.Vars(r)
		email := params["email_id"]

		updateduser, err := userService.UpdateUser(email, user)

		if err != nil {
			//fmt.Sprintln("error....")
			updateUserResponse.Message = err.Error()
			updateUserResponse.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(updateUserResponse)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}
		updateUserResponse.UpdatedUser = updateduser
		updateUserResponse.Message = "User updated successfully."
		updateUserResponse.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)
		//	fmt.Println("handler layer", users)

		//json.NewEncoder(w).Encode(users)
		res, _ := json.Marshal(updateUserResponse)
		w.Write(res)

	}
}

// func GetUsersByEmailHandler(userService service.UserServiceInterface) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var getuser models.GetUserResponse
// 		params := mux.Vars(r)

// 		// convert the id type from string to int
// 		email := params["email_id"]

// 		user, err := userService.GetUserByEmail(email)

// 		if err != nil {
// 			fmt.Sprintln("error....")
// 			getuser.ErrorMsg = "user not render"
// 			getuser.StatusCode = http.StatusInternalServerError
// 			res, _ := json.Marshal(getuser)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write(res)

// 			return
// 		}
// 		getuser.GotUser = user
// 		getuser.StatusCode = http.StatusOK

// 		w.WriteHeader(http.StatusOK)
// 		//	fmt.Println("handler layer", users)

// 		//json.NewEncoder(w).Encode(users)
// 		res, _ := json.Marshal(getuser)
// 		w.Write(res)

// 	}
// }

// func GetUsersByEmailOrIDHandler(userService service.UserServiceInterface) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var getuser models.GetUserResponse
// 		params := mux.Vars(r)

// 		// convert the id type from string to int
// 		email := params["email_id"]

// 		user, err := userService.GetUserByEmail(email)

// 		if err != nil {
// 			fmt.Sprintln("error....")
// 			getuser.ErrorMsg = "user not render"
// 			getuser.StatusCode = http.StatusInternalServerError
// 			res, _ := json.Marshal(getuser)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write(res)

// 			return
// 		}
// 		getuser.GotUser = user
// 		getuser.StatusCode = http.StatusOK

// 		w.WriteHeader(http.StatusOK)
// 		//	fmt.Println("handler layer", users)

// 		//json.NewEncoder(w).Encode(users)
// 		res, _ := json.Marshal(getuser)
// 		w.Write(res)

// 	}
// }

func DeleteUserHandler(userService service.UserServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cretaeUserResponse models.CretaeUserResponse
		params := mux.Vars(r)

		// convert the id type from string to int
		uid, err := strconv.Atoi(params["user_id"])

		if err != nil {
			//log.Fatalf("Unable to convert the string into int.  %v", err)
			log.Println(err)
			return
		}
		_, err = userService.DeleteUser(uid)

		if err != nil {
			cretaeUserResponse.Message = err.Error()
			cretaeUserResponse.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(cretaeUserResponse)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}

		cretaeUserResponse.Message = "User Deleted successfully."
		cretaeUserResponse.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(cretaeUserResponse)
		w.Write(res)
	}
}
