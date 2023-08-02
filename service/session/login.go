package session

import (
	"PSHOP/dao/mysql"
	"PSHOP/dao/sessions"
	"PSHOP/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type LoginResponse struct {
	LoginFlag bool `json:"login_flag"`
}

func init() {
	sessions.Open(sessions.NewRDSOptions("127.0.0.1", 6379, ""))
}

func Login(c *gin.Context) {
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
	session.Values["user"] = uuid.NewString()
	session.Sync()

	if r.RowsAffected == 1 {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"msg":     "hello " + user.UserName,
			"session": session,
		})
	}
}