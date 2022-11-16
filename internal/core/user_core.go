package core

import (
	"DJ-Blog/internal/http/request"
	"DJ-Blog/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type IUserController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	Logout(c *gin.Context)
}

type IGithubUserController interface {
	GithubLogin(c *gin.Context)
	GithubLoginCallback(c *gin.Context)
}

func UserRegister(req *request.UserRegisterReq) (interface{}, bool) {
	errs := make(map[string][]string)
	errs = request.ValidateUserRegisterReq(req)

	if len(errs) > 0 {
		return errs, false
	}

	if !service.CreateUser(req.Username, req.Password) {
		logrus.Error("服务器创建用户失败")
		errs["server"] = []string{"服务器创建用户失败"}
		return errs, false
	}

	return errs, true
}

func UserLogin(req *request.UserLoginReq) (interface{}, bool) {

	errs := make(map[string][]string)
	errs = request.ValidateUserLoginReq(req)
	if len(errs) > 0 {
		return errs, false
	}

	user, ok := service.GetUserByUsername(req.Username)
	if !ok {
		logrus.Error("服务器获取用户失败")
		errs["server"] = []string{"服务器获取用户失败"}
		return errs, false
	}

	return user, true
}
