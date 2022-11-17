package session

import (
	"DJ-Blog/internal/service"
	"DJ-Blog/pkg/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Store sessions.Store

func init() {
	authenticationKey := []byte(config.LoadString("session.authenticationKey"))
	encryptionKey := []byte(config.LoadString("session.encryptionKey"))
	Store = cookie.NewStore(authenticationKey, encryptionKey)
	logrus.Info("Session init success")
}

func GetUser(c *gin.Context) *service.User {
	session := sessions.Default(c)
	userId := session.Get("userId")

	if userId == nil {
		return nil
	}

	var user *service.User
	var ok bool
	if user, ok = service.GetUserById(userId.(int32)); !ok {
		logrus.Warn("session get user failed")
		return nil
	}

	return user
}

func SetUserId(id uint32, session sessions.Session, options sessions.Options) {
	session.Set("userId", id)
	if options == (sessions.Options{}) {
		session.Options(sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		})
	} else {
		session.Options(options)
	}
	session.Save()
}

func ClearUserId(session sessions.Session) {
	session.Delete("userId")
	session.Save()
}
