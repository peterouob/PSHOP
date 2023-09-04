package u

import (
	"PSHOP/model/database/mysql"
	redisdao "PSHOP/model/database/redis"
	"PSHOP/model/user"
	H "PSHOP/utils/http"
	"context"
	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) {
	var cartModel []user.Cart
	identity, err := c.Cookie("userIdentity")
	if err != nil {
		H.Fail(c, "get cookie error")
	}
	str, err := redisdao.Rdb.Get(context.Background(), redisdao.KeyCart(identity)).Result()
	if err != nil && str == "" {
		H.OK(c, "似乎還沒有任何商品")
	}
	if str != "" {
		if err := mysql.Db.Model(&user.Cart{}).Where("user_identity = ?", identity).Find(&cartModel).Error; err != nil {
			H.Fail(c, "get mysql error"+err.Error())
		} else {
			H.OK(c, cartModel)
		}
	}
}
