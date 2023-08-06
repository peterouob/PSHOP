package serviceToken

import (
	"PSHOP/model"
	"PSHOP/model/dao"
	"PSHOP/model/dao/mysql"
	"PSHOP/utils"
	"PSHOP/utils/http"
	"encoding/binary"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	user.TokenUserid = binary.BigEndian.Uint64([]byte(uuid.NewString()))
	tk, err := utils.CreateToken(user.TokenUserid)
	if err != nil {
		H.Fail(c, "create token failed"+err.Error())
	}
	saveErr := dao.SaveTokenAuth(user.UserIdentity, tk)
	if saveErr != nil {
		H.Fail(c, "save token error"+saveErr.Error())
	}
	t := map[string]string{
		"accessToken":  tk.AccessToken,
		"refreshToken": tk.RefreshToken,
	}
	H.SetCookie(c, "accessToken", tk.AccessToken)
	//H.SetCookieForToken(c, "refreshToken", tk.RefreshToken)
	H.OK(c, t)
}
