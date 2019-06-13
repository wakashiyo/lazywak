package controllers

import (
	"net/http"
	"fmt"
)

type UserController struct {}

func NewUserController() *UserController {
	return &UserController{}
}

//Users : ユーザー作成 or ユーザー一覧 
func (c *UserController) Users(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		c.index()
	} else if req.Method == "POST" {
		c.create()
	} else {
		fmt.Println("error")
	}
}

//User : 特定のユーザーに対する処理
func (c *UserController) User(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		c.show()
	case "PATCH":
		c.update()
	case "DELETE":
		c.delete()
	default:
		fmt.Println("error")
	}
}

//create : create user
func (c *UserController) create() {
	fmt.Println("create")
}

//index : show users
func (c *UserController) index() {
	fmt.Println("index")
}

//show : show user by id
func (c *UserController) show() {
	fmt.Println("show")
}

//update : update user profile
func (c *UserController) update() {
	fmt.Println("update")
}

//delete : delete user
func (c *UserController) delete() {
	fmt.Println("delete")
}
