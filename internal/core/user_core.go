package core

import (
	"DJ-Blog/internal/http/request"
	"DJ-Blog/internal/service"
	"github.com/gin-gonic/gin"
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

func UserRegister(req *request.UserRegisterReq) (map[string][]string, bool) {
	errs := make(map[string][]string)
	errs = request.ValidateUserRegisterReq(req)
	if len(errs) > 0 {
		return errs, false
	}

	if !service.CreateUser(req.Username, req.Password) {
		errs["server"] = []string{"服务器内部创建用户失败"}
		return errs, false
	}

	return errs, true
}

//
//func (uc *IUserController) ShowRegister(c *gin.Context) {
//	c.HTML(http.StatusOK, "register", gin.H{})
//}
//
//func (uc *IUserController) Register(c *gin.Context) {
//	username, _ := c.GetPostForm("username")
//	password, _ := c.GetPostForm("password")
//	userRegisterReq := &request.UserRegisterReq{
//		Username: username,
//		Password: password,
//	}
//	errs := request.ValidateUserRegisterReq(userRegisterReq)
//	if len(errs) > 0 {
//		logrus.Warn("注册失败", errs)
//		c.HTML(http.StatusBadRequest, "error", gin.H{})
//		return
//	}
//	user := &model.User{
//		Username: userRegisterReq.Username,
//		Password: userRegisterReq.Password,
//	}
//	if flag := user.IsExistUsername(); flag {
//		logrus.Warn("用户名已存在")
//		c.HTML(http.StatusBadRequest, "error", gin.H{})
//		return
//	}
//	if err := user.Create(); err != nil {
//		logrus.Warn("注册失败")
//		c.HTML(http.StatusInternalServerError, "error", gin.H{})
//		return
//	}
//	var err error
//	if user, err = user.GetUserByUsername(); err != nil {
//		logrus.Warn("获取用户信息失败")
//		c.HTML(http.StatusInternalServerError, "error", gin.H{})
//		return
//	}
//	session.SetSession(user, sessions.Default(c), sessions.Options{})
//	c.Redirect(http.StatusFound, "/user/login")
//}
//
//func (uc *IUserController) ShowLogin(c *gin.Context) {
//	c.HTML(http.StatusOK, "login", gin.H{})
//}
//
//func (uc *IUserController) Login(c *gin.Context) {
//	username, _ := c.GetPostForm("username")
//	password, _ := c.GetPostForm("password")
//	userLoginReq := &request.UserRegisterReq{
//		Username: username,
//		Password: password,
//	}
//	errs := request.ValidateUserRegisterReq(userLoginReq)
//	if len(errs) > 0 {
//		logrus.Warn("登录失败", errs)
//		c.HTML(http.StatusBadRequest, "error", gin.H{})
//		return
//	}
//	user := &model.User{
//		Username: userLoginReq.Username,
//		Password: userLoginReq.Password,
//	}
//	if flag := user.CheckPassword(); !flag {
//		logrus.Warn("密码错误")
//		c.HTML(http.StatusBadRequest, "error", gin.H{})
//		return
//	}
//	var err error
//	if user, err = user.GetUserByUsername(); err != nil {
//		logrus.Warn("获取用户信息失败")
//		c.HTML(http.StatusInternalServerError, "error", gin.H{})
//		return
//	}
//	session.SetSession(user, sessions.Default(c), sessions.Options{})
//	c.Redirect(http.StatusFound, "/")
//}
//
//func (uc *IUserController) GithubLogin(c *gin.Context) {
//	c.Redirect(http.StatusFound, "https://github.com/login/oauth/authorize?client_id=5404b18fe048a5ef58a0")
//}
//
//func (uc *IUserController) GithubLoginCallback(c *gin.Context) {
//	code := c.Query("code")
//
//	resp, err := http.Post("https://github.com/login/oauth/access_token",
//		"application/x-www-form-urlencoded",
//		strings.NewReader(fmt.Sprintf("client_id=%s&client_secret=%s&code=%s",
//			config.GetString("github.clientId"),
//			config.GetString("github.clientSecret"),
//			code)))
//	logrus.Info(code)
//	logrus.Info(config.GetString("github.clientId"))
//	logrus.Info(config.GetString("github.clientSecret"))
//	logrus.Info(err)
//	fmt.Println(err)
//	fmt.Println(resp.Body)
//	defer resp.Body.Close()
//	body, _ := ioutil.ReadAll(resp.Body)
//	var token oauth2.Token
//	form := helper.ParseForm(string(body))
//	for k, v := range form {
//		if k == "access_token" {
//			token.AccessToken = v
//		}
//	}
//	logrus.Info(token.AccessToken)
//	req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
//	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
//	resp, _ = http.DefaultClient.Do(req)
//	defer resp.Body.Close()
//	body, _ = ioutil.ReadAll(resp.Body)
//	var githubUserInfo oauth2.GithubUserInfo
//	if err := json.Unmarshal(body, &githubUserInfo); err != nil {
//		logrus.Warn("github用户信息解析失败")
//		c.HTML(http.StatusInternalServerError, "error", gin.H{})
//		return
//	}
//	user := &model.User{
//		Username:  githubUserInfo.Username,
//		AvatarUrl: githubUserInfo.AvatarUrl,
//	}
//
//	if flag := user.IsExistUsername(); !flag {
//		if err := user.Create(); err != nil {
//			logrus.Warn("注册失败")
//			c.HTML(http.StatusInternalServerError, "error", gin.H{})
//			return
//		}
//	}
//
//	if user, err = user.GetUserByUsername(); err != nil {
//		logrus.Warn("获取用户信息失败")
//		c.HTML(http.StatusInternalServerError, "error", gin.H{})
//		return
//	}
//	session.SetSession(user, sessions.Default(c), sessions.Options{})
//	c.Redirect(http.StatusFound, "/")
//}
//
//func (uc *IUserController) Logout(c *gin.Context) {
//	session := sessions.Default(c)
//	session.Delete("user")
//	session.Save()
//	c.Redirect(http.StatusFound, "/user/login")
//}
