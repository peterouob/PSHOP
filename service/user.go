package service

import (
	"PSHOP/dao/mysql"
	"PSHOP/dao/sessions"
	"PSHOP/model"
	"PSHOP/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	go func() {
		sessions.Open(sessions.NewRDSOptions("127.0.0.1", 6379, ""))
	}()
	var user model.UserModel
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "set value failed",
		})
		return
	}
	r := mysql.Db.Where("user_name = ? and password = ?", user.UserName, user.Password).First(&user)
	session, _ := sessions.GetSession(c.Writer, c.Request)

	session.Values["user"] = &model.UserModel{
		UserName: user.UserName,
	}
	session.Sync()
	if r.RowsAffected == 1 {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"msg":     "hello " + user.UserName,
			"session": session,
		})
	}
}

func Create(c *gin.Context) {
	var user model.UserModel
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "set value failed",
		})
		return
	}
	d := mysql.Db.Where("user_name = ? and password = ?", user.UserName, user.Password).Find(&user)
	if d.RowsAffected >= 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "have same user",
		})
		return
	}
	user.CreateAt = time.Now()
	err := mysql.Db.Create(&user).Error
	if utils.MysqlErrcode(err) == utils.ErrDuplicateEntryCode {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "發生主鍵衝突",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}
