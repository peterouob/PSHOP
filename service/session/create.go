package session

import (
	"PSHOP/model/database/mysql"
	"PSHOP/model/user"
	"PSHOP/utils"
	H "PSHOP/utils/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func Create(c *gin.Context) {
	var user user.UserModel
	var err error
	if err = c.ShouldBind(&user); err != nil {
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
	if err != nil {
		H.Fail(c, "parse uuid failed")
	}
	err = mysql.Db.Create(&user).Error
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
