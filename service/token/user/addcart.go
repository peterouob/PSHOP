package u

import (
	"PSHOP/model/database/mysql"
	"PSHOP/model/user"
	"PSHOP/service/token/user/cart"
	H "PSHOP/utils/http"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	var cartModel user.Cart
	identity, err := c.Cookie("userIdentity")
	if err != nil {
		H.Fail(c, "get cookie error")
	}
	c.ShouldBind(&cartModel)
	cartModel.UserIdentity = identity
	if err := cart.Add(cartModel); err != nil {
		H.Fail(c, err.Error())
	}
	if err := mysql.Db.Create(&cartModel).Error; err != nil {
		H.Fail(c, err.Error())
	}
	H.OK(c, "success insert cart")
}
