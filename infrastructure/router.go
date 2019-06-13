package infrastructure

import (
	"net/http"
	"fmt"
	"../interfaces/controllers"
)

func Run()  {

	userController := controllers.NewUserController()
	authController := controllers.NewAuthController()

	http.HandleFunc("/users", userController.Users)
	http.HandleFunc("/users/", userController.User)
	http.HandleFunc("/auth", authController.Auth)
	
	err := http.ListenAndServe(":9000", nil)
	
	if err != nil {
		fmt.Println(err)
	}
}

// func routingGen() {

// }
