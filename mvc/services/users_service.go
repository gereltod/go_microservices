package services

import (
	"github.com/gereltod/go_microservices/mvc/domain"
	"github.com/gereltod/go_microservices/mvc/utils"
)

type usersService struct{}

var (
	//UsersService services
	UsersService usersService
)

//GetUser from domain
func (s *usersService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userID)
}
