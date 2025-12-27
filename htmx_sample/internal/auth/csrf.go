package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GenerateCSRF(c *gin.Context) string {
	ses := sessions.Default(c)
	token := uuid.NewString()
	ses.Set("csrf_token", token)

	ses.Save()

	return token
}

func VerifyCSRF() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		formToken := c.PostForm("csrf_token")

		sessionToken := session.Get("csrf_token")

		if sessionToken == nil || formToken != sessionToken {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}
