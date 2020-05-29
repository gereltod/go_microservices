package services

import (
	"net/http"
	"testing"

	"github.com/gereltod/go_microservices/mvc/domain"
	"github.com/gereltod/go_microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
)

var(
	userDaoMock usersDaoMock
	
	getUserFunction func(userID int64) (*domain.User, *utils.ApplicationError)
)

func init(){
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userID int64) (*domain.User, *utils.ApplicationError){
	return getUserFunction(userID)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	user, err := UsersService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 not found", err.Message)
}
