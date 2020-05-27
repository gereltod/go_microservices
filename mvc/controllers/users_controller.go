package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gereltod/go_microservices/mvc/services"
	"github.com/gereltod/go_microservices/mvc/utils"
)

// GetUser info
func GetUser(resp http.ResponseWriter, req *http.Request) {

	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	log.Printf("about to user_id %v", userID)

	user, apiErr := services.GetUser(userID)

	if apiErr != nil {
		jsonValue, _:=json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
