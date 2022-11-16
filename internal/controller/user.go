package controller

import "github.com/gin-gonic/gin"

type UserController struct {
}

func NewUserController() *UserController {
	return new(UserController)
}

func (u *UserController) Login(c *gin.Context) {

}

func (u *UserController) Register(c *gin.Context) {

}

func (u *UserController) DoLogin(c *gin.Context) {

}

func (u *UserController) DoRegister(c *gin.Context) {

}

func (u *UserController) Logout(c *gin.Context) {

}
