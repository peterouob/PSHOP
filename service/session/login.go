package session

import (
	"PSHOP/model/database/mysql"
	sessions2 "PSHOP/model/database/sessions"
	"PSHOP/model/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func init() {
	sessions2.Open(sessions2.NewRDSOptions("127.0.0.1", 6379, ""))
}

var session *sessions2.Session

func Login(c *gin.Context) {
	var u user.UserModel
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "set value failed",
		})
		return
	}
	var uid = uuid.NewString()
	r := mysql.Db.Where("user_name = ? and password = ?", u.UserName, u.Password).First(&u)
	session, _ = sessions2.GetSession(c, c.Request)
	session.Values["user"] = uid
	session.Sync()

	if r.RowsAffected == 1 {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"msg":     "hello " + u.UserName,
			"session": session,
			"prename": <-user.C,
		})
	}
}
