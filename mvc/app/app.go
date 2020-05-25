package app

import (
	"net/http"

	"github.com/gereltod/go_microservices/mvc/controllers"
)

// StartApp string application
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	if err:=http.ListenAndServe(":8080",nil); err!=nil{
		panic(err)
	}
}