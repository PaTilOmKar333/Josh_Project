//struct and methods

package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"project/models"
	"project/repo"
	"regexp"
)

type UserServiceInterface interface {
	//user methods
	Login(authdetails models.Authentication) (validToken string, err error)
	ListUsers() (users []models.UserList, err error)
	CreateUser(user models.User) (createduser models.User, err error)
	GetUser(ctx context.Context, variable interface{}) (user models.User, err error)
	UpdateUser(ctx context.Context, email string, user models.User) (updateUser models.UpdateUser, err error)
	//GetUserByEmail(email string) (user models.User, err error)
	DeleteUser(uid int) (id int, err error)
	// login(authdetails models.Authentication) (err error)
}

type userService struct {
	repo     repo.UserRepoInterface
	gentoken AuthTokenInterface
}

func InitUserService(r repo.UserRepoInterface, at AuthTokenInterface) UserServiceInterface {

	return &userService{
		repo:     r,
		gentoken: at,
	}
}

func (us *userService) ListUsers() (users []models.UserList, err error) {
	//val, ok := ctx.Value("ClaimsToVerify").(*models.Claims)
	users, err = us.repo.ListUser()
	if err != nil {
		return
	}
	return
}

func (us *userService) Login(authdetails models.Authentication) (validToken string, err error) {
	user, err := us.repo.GetUserByEmail(authdetails.Email)

	if user.Password == authdetails.Password {
		validToken, err = us.gentoken.GenerateToken(user.User_ID, user.Email, user.Role)
		return
	} else {
		log.Println(err)
		err = errors.New("login failed. please check check credantials")
		return
	}
}

func (us *userService) CreateUser(user models.User) (createduser models.User, err error) {

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	isvalid := emailRegex.MatchString(user.Email)
	fmt.Println(user.Email, isvalid)
	if isvalid {
		createduser, err = us.repo.CreateUser(user)
		if err != nil {
			return
		}
	} else {
		err = errors.New("invalid email address")
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

// func (us *userService) GetUser(ctx context.Context, variable interface{}) (user models.User, err error) {
func (us *userService) GetUser(ctx context.Context, variable interface{}) (user models.User, err error) {
	val, _ := ctx.Value("ClaimsToVerify").(*models.Claims)

	if val.Email == variable || val.UserID == variable || val.Role == "admin" || val.Role == "superadmin" {
		user, err = us.repo.GetUser(variable)
		if err != nil {
			return
		}
	} else {
		err = errors.New("you are unauthorized person")
		return
	}

	return
}

func (us *userService) UpdateUser(ctx context.Context, email string, user models.User) (updateUser models.UpdateUser, err error) {
	val, _ := ctx.Value("ClaimsToVerify").(*models.Claims)
	if val.Email == email || val.Role == "admin" || val.Role == "superadmin" {
		updateUser, err = us.repo.UpdateUser(email, user)
		if err != nil {
			return
		}
	} else {
		err = errors.New("you are unauthorized person")
		return
	}
	return
}

// func (us *userService) GetUserByEmail(email string) (user models.User, err error) {
// 	user, err = us.repo.GetUserByEmail(email)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

func (us *userService) DeleteUser(uid int) (id int, err error) {
	id, err = us.repo.DeleteUser(uid)
	if err != nil {
		return
	}
	return
}
