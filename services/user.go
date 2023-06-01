package services

import (
	"kvm-dashboard/conf"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (svc *Service) Login(username, password string, c *gin.Context) bool {
	if username == conf.C.User.Username && password == conf.C.User.Password {
		setSession(c, username)
		return true
	} else {
		return false
	}
}

func setSession(c *gin.Context, username string) {
	session := sessions.Default(c)

	session.Options(sessions.Options{
		MaxAge: 60,
		Path:   "/", // set cookie path
	})
	session.Set("username", username)
	session.Save()
}
