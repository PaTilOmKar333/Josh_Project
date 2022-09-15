package handler

import (
	"encoding/json"
	"net/http"
	"project/models"
	"project/service"
	"strconv"

	"github.com/gorilla/mux"
)

func LoginHandler(userService service.UserServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var authdetails models.Authentication
		var tokenResponse models.TokenResponse

		err := json.NewDecoder(r.Body).Decode(&authdetails)
		if err != nil {

			tokenResponse.Message = "Error in reading body."
			tokenResponse.StatusCode = http.StatusBadRequest
			res, _ := json.Marshal(tokenResponse)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)
			return
		}
		token, err := userService.Login(authdetails)
		if err != nil {
			tokenResponse.Message = err.Error()
			tokenResponse.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(tokenResponse)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}
		tokenResponse.Token = token
		tokenResponse.StatusCode = http.StatusOK
		tokenResponse.Message = "Token Generated Successfully."

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(tokenResponse)
		w.Write(res)
	}
}

func AllUsersHandler(userService service.UserServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ListUsersResponse models.UserListResponse

		users, err := userService.ListUsers()
		if err != nil {
			ListUsersResponse.Message = err.Error()
			ListUsersResponse.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(ListUsersResponse)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}
		ListUsersResponse.UserList = users
		ListUsersResponse.StatusCode = http.StatusOK
		ListUsersResponse.Message = "These are all users."

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(ListUsersResponse)
		w.Write(res)

	}
}

func CreateUsersHandler(userService service.UserServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var createUserResponse models.CreateUserResponse

		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			createUserResponse.Message = "unable to decode the request body."
			createUserResponse.StatusCode = http.StatusBadRequest
			w.WriteHeader(http.StatusBadRequest)
			res, _ := json.Marshal(createUserResponse)
			w.Write(res)
			return
		}

		createduser, err := userService.CreateUser(user)
		if err != nil {
			createUserResponse.Message = err.Error()
			createUserResponse.StatusCode = http.StatusBadRequest
			res, _ := json.Marshal(createUserResponse)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)

			return
		}
		createUserResponse.CreateUser.UserID = createduser.User_ID
		createUserResponse.CreateUser.FirstName = createduser.FirstName
		createUserResponse.CreateUser.Email = createduser.Email
		createUserResponse.CreateUser.Password = createduser.Password
		createUserResponse.Message = "User Created successfully."
		createUserResponse.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(createUserResponse)
		w.Write(res)

		//	}
		// else {
		// 	createUserResponse.Message = "Invalid email. please type valid email address."
		// 	createUserResponse.StatusCode = http.StatusBadRequest
		// 	res, _ := json.Marshal(createUserResponse)
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	w.Write(res)
		// }

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
				getuser.ErrorMsg = "unable to decode the request body."
				getuser.StatusCode = http.StatusBadRequest
				w.WriteHeader(http.StatusBadRequest)
				res, _ := json.Marshal(getuser)
				w.Write(res)
				return
			}

		} else if param == "email_id" {
			newvariable = variable
		} else {
			getuser.ErrorMsg = "wrong parameters."
			getuser.StatusCode = http.StatusBadRequest
			w.WriteHeader(http.StatusBadRequest)
			res, _ := json.Marshal(getuser)
			w.Write(res)
			return
		}
		user, err := userService.GetUser(r.Context(), newvariable)

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
		var updateUserResponse models.UpdateUserResponse
		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			updateUserResponse.Message = "unable to decode the request body."
			updateUserResponse.StatusCode = http.StatusBadRequest
			w.WriteHeader(http.StatusBadRequest)
			res, _ := json.Marshal(updateUserResponse)
			w.Write(res)
			return
		}
		params := mux.Vars(r)
		email := params["email_id"]

		updateduser, err := userService.UpdateUser(r.Context(), email, user)

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

		res, _ := json.Marshal(updateUserResponse)
		w.Write(res)

	}
}

func DeleteUserHandler(userService service.UserServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var DeleteUserResponse models.CreateUserResponse
		params := mux.Vars(r)

		// convert the id type from string to int
		uid, err := strconv.Atoi(params["user_id"])

		if err != nil {
			DeleteUserResponse.Message = "unable to convert userid in int."
			DeleteUserResponse.StatusCode = http.StatusBadRequest
			w.WriteHeader(http.StatusBadRequest)
			res, _ := json.Marshal(DeleteUserResponse)
			w.Write(res)
			return
		}
		_, err = userService.DeleteUser(uid)

		if err != nil {
			DeleteUserResponse.Message = err.Error()
			DeleteUserResponse.StatusCode = http.StatusInternalServerError
			res, _ := json.Marshal(DeleteUserResponse)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(res)

			return
		}

		DeleteUserResponse.Message = "User Deleted successfully."
		DeleteUserResponse.StatusCode = http.StatusOK

		w.WriteHeader(http.StatusOK)

		res, _ := json.Marshal(DeleteUserResponse)
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
