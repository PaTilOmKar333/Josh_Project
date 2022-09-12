//struct and methods

package service

import (
	"project/models"
	"project/repo"
)

type UserServiceInterface interface {
	//user methods
	ListUsers() (users []models.UserList, err error)
	CreateUser(user models.User) (createduser models.User, err error)
	//GetUser(param interface{}) (user models.User, err error)
	UpdateUser(email string, user models.User) (updateUser models.UpdateUser, err error)
	GetUserByEmail(email string) (user models.User, err error)
	DeleteUser(uid int) (id int, err error)
	// login(authdetails models.Authentication) (err error)
}

type userService struct {
	repo repo.UserRepoInterface
}

func InitUserService(r repo.UserRepoInterface) UserServiceInterface {

	//initialies
	//repo.InitUserRepo()
	return &userService{
		repo: r,
	}
}

func (us *userService) ListUsers() (users []models.UserList, err error) {
	users, err = us.repo.ListUser()
	//fmt.Println("service layer: ", users)
	if err != nil {
		return
	}
	return
}

// func (us *userService) login(authdetails models.Authentication) (err error) {
// 	_, err = us.repo.login(authdetails)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

func (us *userService) CreateUser(user models.User) (createduser models.User, err error) {
	createduser, err = us.repo.CreateUser(user)
	if err != nil {
		return
	}
	return
}

// func (us *userService) GetUser(uid int) (user models.User, err error) {
// 	user, err = us.repo.GetUser(uid)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

// func (us *userService) GetUser(param interface{}) (user models.User, err error) {
// 	user, err = us.repo.GetUser(param)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

func (us *userService) UpdateUser(email string, user models.User) (updateUser models.UpdateUser, err error) {
	updateUser, err = us.repo.UpdateUser(email, user)
	if err != nil {
		return
	}
	return
}

func (us *userService) GetUserByEmail(email string) (user models.User, err error) {
	user, err = us.repo.GetUserByEmail(email)
	if err != nil {
		return
	}
	return
}

func (us *userService) DeleteUser(uid int) (id int, err error) {
	id, err = us.repo.DeleteUser(uid)
	if err != nil {
		return
	}
	return
}
