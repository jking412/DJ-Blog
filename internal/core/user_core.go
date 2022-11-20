package core

import (
	"DJ-Blog/internal/http/request"
	"DJ-Blog/internal/model"
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

	user := &model.UserModel{
		Username:  req.Username,
		Password:  req.Password,
		AvatarUrl: req.AvatarUrl,
	}

	var serviceUser *service.User

	if serviceUser = service.CreateUser(user); serviceUser.Id == 0 {
		logrus.Error("服务器创建用户失败")
		errs["server"] = []string{"服务器创建用户失败"}
		return errs, false
	}

	serviceUser.Password = ""

	return serviceUser, true
}

func UserLogin(req *request.UserLoginReq) (interface{}, bool) {

	errs := make(map[string][]string)
	errs = request.ValidateUserLoginReq(req)
	if len(errs) > 0 {
		return errs, false
	}

	serviceUser, ok := service.GetUserByUsername(req.Username)
	if !ok {
		logrus.Error("服务器获取用户失败")
		errs["server"] = []string{"服务器获取用户失败"}
		return errs, false
	}

	serviceUser.Password = ""

	return serviceUser, true
}
