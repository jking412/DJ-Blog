package sessionpkg

import (
	"DJ-Blog/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Store sessions.Store

func InitSession() {
	Store = cookie.NewStore([]byte("DJ-Blog"))
	logrus.Info("Session init success")
}

func GetUser(c *gin.Context) *model.User {
	session := sessions.Default(c)
	userId := session.Get("userId")
	username := session.Get("username")
	userAvatarUrl := session.Get("userAvatarUrl")
	if userId == nil || username == nil || userAvatarUrl == nil {
		return nil
	}
	user := &model.User{
		BaseModel: model.BaseModel{
			Id: userId.(uint64),
		},
		Username:  username.(string),
		AvatarUrl: userAvatarUrl.(string),
	}
	if user == nil {
		return nil
	}
	return user
}

func SetSession(user *model.User, session sessions.Session, options sessions.Options) {
	session.Set("userId", user.Id)
	session.Set("username", user.Username)
	session.Set("userAvatarUrl", user.AvatarUrl)
	if options == (sessions.Options{}) {
		options = sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
	}
	session.Options(options)
	session.Save()
}

func FlushSession(c *gin.Context) {
	session := sessions.Default(c)
	SetSession(GetUser(c), session, sessions.Options{})
}
