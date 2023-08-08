package serviceToken

import (
	"PSHOP/model/dao/mysql"
	"PSHOP/model/dao/redis"
	"PSHOP/model/user"
	"PSHOP/utils"
	"PSHOP/utils/http"
	"encoding/binary"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Login(c *gin.Context) {
	var u user.UserModel
	if err := c.ShouldBind(&u); err != nil {
		H.Fail(c, "bind failed")
	}
	r := mysql.Db.Where("username = ? and password = ?", u.UserName, u.Password)
	if r.RowsAffected == 1 {
		fmt.Println("get User")
	}
	u.TokenUserid = binary.BigEndian.Uint64([]byte(uuid.NewString()))
	tk, err := utils.CreateToken(u.TokenUserid)
	if err != nil {
		H.Fail(c, "create token failed"+err.Error())
	}
	saveErr := redisdao.SaveTokenAuth(u.UserIdentity, tk)
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
