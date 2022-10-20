package controller

import (
	"DJ-Blog/model"
	"DJ-Blog/pkg/helper"
	"DJ-Blog/pkg/oauth2"
	"DJ-Blog/pkg/sessionpkg"
	"DJ-Blog/pkg/viperlib"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

type UserController struct {
}

type UserRegisterReq struct {
	Username string
	Password string
}

type UserLoginReq struct {
	Username string
	Password string
}

func (uc *UserController) ShowRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register", gin.H{})
}

func (uc *UserController) Register(c *gin.Context) {
	username, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")
	userRegisterReq := &UserRegisterReq{
		Username: username,
		Password: password,
	}
	user := &model.User{
		Username: userRegisterReq.Username,
		Password: userRegisterReq.Password,
	}
	if flag := user.IsExistUsername(); flag {
		logrus.Warn("用户名已存在")
		c.HTML(http.StatusBadRequest, "error", gin.H{})
		return
	}
	if err := user.Create(); err != nil {
		logrus.Warn("注册失败")
		c.HTML(http.StatusInternalServerError, "error", gin.H{})
		return
	}
	var err error
	if user, err = user.GetUserByUsername(); err != nil {
		logrus.Warn("获取用户信息失败")
		c.HTML(http.StatusInternalServerError, "error", gin.H{})
		return
	}
	sessionpkg.SetSession("userId", user.Id, sessions.Default(c), sessions.Options{})
	c.Redirect(http.StatusFound, "/user/login")
}

func (uc *UserController) ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login", gin.H{})
}

func (uc *UserController) Login(c *gin.Context) {
	username, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")
	userLoginReq := &UserRegisterReq{
		Username: username,
		Password: password,
	}
	user := &model.User{
		Username: userLoginReq.Username,
		Password: userLoginReq.Password,
	}
	if flag := user.CheckPassword(); !flag {
		logrus.Warn("密码错误")
		c.HTML(http.StatusBadRequest, "error", gin.H{})
		return
	}
	var err error
	if user, err = user.GetUserByUsername(); err != nil {
		logrus.Warn("获取用户信息失败")
		c.HTML(http.StatusInternalServerError, "error", gin.H{})
		return
	}
	sessionpkg.SetSession("userId", user.Id, sessions.Default(c), sessions.Options{})
	c.Redirect(http.StatusFound, "/")
}

func (uc *UserController) GithubLogin(c *gin.Context) {
	c.Redirect(http.StatusFound, "https://github.com/login/oauth/authorize?client_id=5404b18fe048a5ef58a0")
}

func (uc *UserController) GithubLoginCallback(c *gin.Context) {
	code := c.Query("code")

	resp, _ := http.Post("https://github.com/login/oauth/access_token",
		"application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf("client_id=%s&client_secret=%s&code=%s",
			viperlib.GetString("github.clientId"),
			viperlib.GetString("github.clientSecret"),
			code)))
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var token oauth2.Token
	form := helper.ParseForm(string(body))
	for k, v := range form {
		if k == "access_token" {
			token.AccessToken = v
		}
	}
	req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	resp, _ = http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	var githubUserInfo oauth2.GithubUserInfo
	if err := json.Unmarshal(body, &githubUserInfo); err != nil {
		logrus.Warn("github用户信息解析失败")
		c.HTML(http.StatusInternalServerError, "error", gin.H{})
		return
	}
	user := &model.User{
		Username:  githubUserInfo.Username,
		AvatarUrl: githubUserInfo.AvatarUrl,
	}

	if flag := user.IsExistUsername(); !flag {
		if err := user.Create(); err != nil {
			logrus.Warn("注册失败")
			c.HTML(http.StatusInternalServerError, "error", gin.H{})
			return
		}
	}

	var err error
	if user, err = user.GetUserByUsername(); err != nil {
		logrus.Warn("获取用户信息失败")
		c.HTML(http.StatusInternalServerError, "error", gin.H{})
		return
	}
	sessionpkg.SetSession("userId", user.Id, sessions.Default(c), sessions.Options{})
	c.Redirect(http.StatusFound, "/")
}

func (uc *UserController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userId")
	session.Save()
	c.Redirect(http.StatusFound, "/user/login")
}
