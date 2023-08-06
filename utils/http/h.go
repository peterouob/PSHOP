package H

import (
	"PSHOP/model/dao/redis"
	"PSHOP/utils"
	"context"
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
