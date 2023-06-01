package middleware

import (
	"kvm-dashboard/model"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CheckValidTimestamp() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		// check if login
		if session.Get("username") == nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, model.FailedWithMsg("Please login"))
			return
		}

		// expand time
		session.Options(sessions.Options{
			MaxAge: 60,
			Path:   "/",
		})
		session.Save()

		c.Next()
	}
}
