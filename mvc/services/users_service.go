package services

import (
	"github.com/gereltod/go_microservices/mvc/domain"
	"github.com/gereltod/go_microservices/mvc/utils"
)

//GetUser from domain
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}