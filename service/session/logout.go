package session

import (
	H "PSHOP/http"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	session.Remove()
	H.RemoveCookie(c, "session_cookie")
	H.OK(c, "logout successful")
}
