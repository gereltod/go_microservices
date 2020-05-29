package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	//Initialization:


	//Execution:
	user, err := UserDao.GetUser(0)

	//Validation:
	assert.Nil(t, user, "we were not expection a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 not found", err.Message)

}

func TestGetUserNoError(t *testing.T) {
	user, err := UserDao.GetUser(1)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.ID)
	assert.EqualValues(t, "Gerelt-Od", user.FirstName)
	assert.EqualValues(t, "Surenjav", user.LastName)
	assert.EqualValues(t, "gereltod.s@gmail.com", user.Email)

}
