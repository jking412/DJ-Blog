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
	"net/http"
)

type UserController struct {
}

func NewUserController() *UserController {
	return new(UserController)
}

// Login godoc
// @Summary     用户登录
// @Description 用户登录
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       UserLoginRequest body     request.UserLoginReq               true "用户登录请求"
// @Success     200            {object} response.Status{data=service.User} "登录成功"
// @Failure     401            {object} response.Status                    "请求参数格式错误"
// @Failure     402            {object} response.Status                    "请求参数不符合要求"
// @Router      /user/login [post]
func (u *UserController) Login(c *gin.Context) {
	req := &request.UserLoginReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Info("注册请求JSON格式错误", err)
		response.EndWithUnprocessableData(c, nil)
		return
	}

	value, ok := core.UserLogin(req)
	if !ok {
		logrus.Info("登录失败", value)
		response.EndWithUnsatisfiedRequest(c, value)
		return
	}

	user := value.(*service.User)

	s := sessions.Default(c)
	session.SetUserId(user.Id, s, sessions.Options{})

	response.EndWithOK(c, user)
}

// Register godoc
// @Summary     用户注册
// @Description 用户注册
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       userRegisterRequest body     request.UserRegisterReq            true "用户注册请求"
// @Success     200               {object} response.Status{data=service.User} "注册成功"
// @Failure     401               {object} response.Status                    "请求参数格式错误"
// @Failure     402               {object} response.Status                    "请求参数不符合要求"
// @Router      /user/register [post]
func (u *UserController) Register(c *gin.Context) {
	req := &request.UserRegisterReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Info("注册请求JSON格式错误", err)
		response.EndWithUnprocessableData(c, nil)
		return
	}

	value, ok := core.UserRegister(req)
	if !ok {
		logrus.Info("注册失败", value)
		response.EndWithUnsatisfiedRequest(c, value)
		return
	}

	user := value.(*service.User)

	s := sessions.Default(c)
	session.SetUserId(user.Id, s, sessions.Options{})

	response.EndWithOK(c, user)
}

// Logout godoc
// @Summary     退出登录
// @Description 退出登录
// @Tags        User
// @Success     200 {object} response.Status "退出成功"
// @Router      /user/logout [get]
func (u *UserController) Logout(c *gin.Context) {
	session.ClearUserId(sessions.Default(c))
	response.EndWithOK(c, nil)
}

func (u *UserController) UploadAvatarImg(c *gin.Context) {

	avatarImg, err := c.FormFile("avatarImg")
	if err != nil {
		logrus.Info("上传文件失败", err)
		response.EndWithUnprocessableData(c, nil)
		return
	}

	user := session.GetUser(c)
	ok := core.UploadAvatarImg(user.Username, avatarImg)
	if !ok {
		logrus.Warn("上传头像失败")
		response.EndWithInternalServerError(c, nil)
		return
	}

	response.EndWithOK(c, nil)
}

func (u *UserController) GetAvatarImg(c *gin.Context) {

	user := session.GetUser(c)

	avatarImg, ok := core.GetAvatarImg(user.Username)
	if !ok {
		logrus.Warn("获取头像失败")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "获取头像失败",
		})
		return
	}

	// TODO: 重构response包，封装响应图片的方法，感觉自己封装的response不如原来的简洁
	// TODO: 数据混乱，要重新处理，包括整个处理的逻辑
	c.Data(200, "image/jpeg", avatarImg)
}

//func (u *UserController) GithubLogin(c *gin.Context) {
//	c.Redirect(http.StatusFound, fmt.Sprintf(conf.GithubCallbackUrl, conf.GithubClientId))
//}

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
