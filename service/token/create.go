package serviceToken

import (
	"PSHOP/model/database/mysql"
	user2 "PSHOP/model/user"
	"PSHOP/utils"
	H "PSHOP/utils/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

func CreateUser(c *gin.Context) {
	var u user2.UserModel
	if err := c.ShouldBind(&u); err != nil {
		H.Fail(c, "bind failed")
	}
	d := mysql.Db.Where("user_name = ? and password = ?", u.UserName, u.Password).Find(&u)
	if d.RowsAffected >= 1 {
		H.Fail(c, "have same user")
	}
	identity := uuid.NewString()
	u.UserIdentity = identity
	u.CreateAt = time.Now()
	err := mysql.Db.Create(&u).Error
	if utils.MysqlErrcode(err) == utils.ErrDuplicateEntryCode {
		H.Fail(c, "發生主鍵衝突")
	}
	H.OK(c, "success create")
}
