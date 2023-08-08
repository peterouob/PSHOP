package session

import (
	"PSHOP/model/dao/mysql"
	"PSHOP/model/user"
	"PSHOP/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func Create(c *gin.Context) {
	var user user.UserModel
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
	user.UserIdentity = uuid.NewString()
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
