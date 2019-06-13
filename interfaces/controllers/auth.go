package controllers

import (
	"net/http"
	"fmt"
)

type AuthController struct {}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (a *AuthController) Auth(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)
}
