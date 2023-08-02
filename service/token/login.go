package service

import (
	"PSHOP/dao/mysql"
	H "PSHOP/http"
	"PSHOP/model"
	"PSHOP/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user model.UserModel
	if err := c.ShouldBind(&user); err != nil {
		H.Fail(c, "bind failed")
	}
	r := mysql.Db.Where("username = ? and password = ?", user.UserName, user.Password)
	if r.RowsAffected == 1 {
		fmt.Println("get User")
	}
	token, err := utils.CreateToken(user.UserIdentity)
	if err != nil {
		H.Fail(c, "failed to create token")
	} else {
		H.OK(c, "token created"+" "+token)
	}
}
