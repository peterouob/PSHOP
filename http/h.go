package H

import (
	"PSHOP/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func OK(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  msg,
	})
}

func Fail(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": -1,
		"msg":  msg,
	})
	return
}

func Forbidden(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusForbidden, gin.H{
		"code": -1,
		"msg":  msg,
	})
	c.Abort()
}

func SetCookie(c *gin.Context, name, value string) {
	c.SetCookie(name, value, 365*3600, "/", utils.Config.GetString("server.host"), false, true)
}

func RemoveCookie(c *gin.Context, key string) {
	c.SetCookie(key, "", -1, "", utils.Config.GetString("server.host"), false, true)
}
