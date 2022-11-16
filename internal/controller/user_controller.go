package controller

import (
	"DJ-Blog/internal/conf"
	"DJ-Blog/internal/core"
	"DJ-Blog/internal/http/request"
	"DJ-Blog/internal/http/response"
	"DJ-Blog/internal/service"
	"DJ-Blog/pkg/session"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserController struct {
}

func NewUserController() *UserController {
	return new(UserController)
}

func (u *UserController) Login(c *gin.Context) {
	req := &request.UserLoginReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Info("注册请求JSON格式错误", err)
		response.EndWithUnprocessableJSON(c, nil)
	}

	value, ok := core.UserLogin(req)
	if !ok {
		logrus.Info("登录失败", value)
		response.EndWithUnsatisfiedRequest(c, value)
	}

	s := sessions.Default(c)
	session.SetUserId(value.(*service.User).Id, s, sessions.Options{})

	response.EndWithOK(c, value)
}

func (u *UserController) Register(c *gin.Context) {
	req := &request.UserRegisterReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Info("注册请求JSON格式错误", err)
		response.EndWithUnprocessableJSON(c, nil)
	}

	errs, ok := core.UserRegister(req)
	if !ok {
		logrus.Info("注册失败", errs)
		response.EndWithUnsatisfiedRequest(c, errs)
	}

	user, ok := service.GetUserByUsername(req.Username)
	if !ok {
		logrus.Error("创建用户后获取用户失败")
		response.EndWithInternalServerError(c, errs)
	}

	s := sessions.Default(c)
	session.SetUserId(user.Id, s, sessions.Options{})

	response.EndWithOK(c, user)
}

func (u *UserController) Logout(c *gin.Context) {
	session.ClearUserId(sessions.Default(c))
	response.EndWithOK(c, nil)
}

func (u *UserController) GithubLogin(c *gin.Context) {
	c.Redirect(http.StatusFound, fmt.Sprintf(conf.GithubCallbackUrl, conf.GithubClientId))
}

//func (u *UserController) GithubLoginCallback(c *gin.Context) {
//	code := c.Query("code")
//
//	resp, _ := http.Post("https://github.com/login/oauth/access_token",
//		"application/x-www-form-urlencoded",
//		strings.NewReader(fmt.Sprintf("client_id=%s&client_secret=%s&code=%s",
//			conf.GithubClientId,
//			conf.GithubSecret,
//			code)))
//
//	defer resp.Body.Close()
//	body, _ := io.ReadAll(resp.Body)
//
//	var token oauth2.Token
//	form := helper.ParseForm(string(body))
//	for k, v := range form {
//		if k == "access_token" {
//			token.AccessToken = v
//		}
//	}
//
//	req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
//	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
//
//	resp, _ = http.DefaultClient.Do(req)
//	defer resp.Body.Close()
//	body, _ = io.ReadAll(resp.Body)
//
//	var githubUserInfo oauth2.GithubUserInfo
//	if err := json.Unmarshal(body, &githubUserInfo); err != nil {
//		logrus.Warn("github用户信息解析失败")
//		response.EndWithInternalServerError(c, nil)
//		return
//	}
//
//	response.EndWithOK(c, githubUserInfo)
//}
