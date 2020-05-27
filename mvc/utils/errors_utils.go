package utils

//ApplicationError for http error handle
type ApplicationError struct {
	Message string `json:"message`
	StatusCode int `json:"status"`
	Code string `jsong:"code"`
}