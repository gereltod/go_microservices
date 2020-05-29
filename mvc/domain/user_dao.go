package domain

import (
	"fmt"
	"net/http"

	"github.com/gereltod/go_microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		1: {ID: 1, FirstName: "Gerelt-Od", LastName: "Surenjav", Email: "gereltod.s@gmail.com"},
		2: {ID: 2, FirstName: "Test", LastName: "Surenjav", Email: "gereltod.s@gmail.com"},
	}

	//UserDao domain
	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

// GetUser from db
func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	// if user==nil {
	// 	return User{}, errors.New(fmt.Sprintf("user %v was not found", userId))
	// }

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
