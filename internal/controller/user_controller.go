package controller

import (
	"DJ-Blog/internal/core"
	"DJ-Blog/internal/http/request"
	"DJ-Blog/internal/http/response"
	"DJ-Blog/internal/service"
	"DJ-Blog/pkg/session"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserController struct {
}

func NewUserController() *UserController {
	return new(UserController)
}

func (u *UserController) Login(c *gin.Context) {

}

func (u *UserController) Register(c *gin.Context) {
	req := &request.UserRegisterReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Warn("注册请求JSON格式错误", err)
		response.EndWithUnprocessableJSON(c, nil)
	}

	errs, ok := core.UserRegister(req)
	if !ok {
		logrus.Warn("注册失败", errs)
		response.EndWithUnsatisfiedRequest(c, errs)
	}

	user, ok := service.GetUserByUsername(req.Username)
	if !ok {
		logrus.Warn("创建用户后获取用户失败")
		response.EndWithInternalServerError(c, errs)
	}

	s := sessions.Default(c)
	session.SetUserId(user.Id, s, sessions.Options{})

	response.EndWithOK(c, user)
}

func (u *UserController) Logout(c *gin.Context) {

}
