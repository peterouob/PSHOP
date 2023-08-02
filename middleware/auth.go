package middleware

import (
	H "PSHOP/http"
	"github.com/gin-gonic/gin"
)

func SessionAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("session_cookie")
		if err != nil {
			H.Forbidden(c, err.Error())
		} else {
			H.OK(c, "hello")
		}
	}
}

func RDBAuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("accessToken")
		if err != nil {
			H.Forbidden(c, "not authorized")
		} else {
			H.OK(c, "hello")
		}
	}
}
