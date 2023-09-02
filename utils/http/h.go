package H

import (
	"PSHOP/model/database/redis"
	"PSHOP/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func FectApi(authD *utils.AccessDetails) (uint64, error) {
	userid, err := redisdao.Rdb.Get(context.Background(), authD.AccessUid).Result()
	if err != nil {
		return 0, err
	}
	userId, _ := strconv.ParseUint(userid, 10, 64)
	return userId, nil
}

func Cors(c *gin.Context) {
	method := c.Request.Method
	c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
	fmt.Println(c.GetHeader("Origin"))
	c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token")
	c.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
	c.Header("Access-Control-Allow-Credentials", "true")
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.Next()
}
