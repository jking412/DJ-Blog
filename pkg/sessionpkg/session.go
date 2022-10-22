package sessionpkg

import (
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

func GetUserId(c *gin.Context) uint64 {
	session := sessions.Default(c)
	userId := session.Get("userId")
	if userId == nil {
		return 0
	}
	return userId.(uint64)
}

func GetUsername(c *gin.Context) string {
	session := sessions.Default(c)
	username := session.Get("username")
	if username == nil {
		return ""
	}
	return username.(string)
}

func SetSession(key string, value interface{}, session sessions.Session, options sessions.Options) {
	session.Set(key, value)
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
	SetSession("userId", GetUserId(c), session, sessions.Options{})
	SetSession("username", GetUsername(c), session, sessions.Options{})
}
