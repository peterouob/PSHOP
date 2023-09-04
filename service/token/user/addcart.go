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

	result := mysql.Db.Find(&cartModel, "good_id =?", cartModel.GoodID).RowsAffected
	if result >= 1 {
		if err := mysql.Db.Model(&user.Cart{}).Where("good_id =?", cartModel.GoodID).Update("nums", cartModel.Nums+1).Error; err != nil {
			H.Fail(c, err.Error())
		}
		H.OK(c, "success")
	} else {
		if err := mysql.Db.Create(&cartModel).Error; err != nil {
			H.Fail(c, err.Error())
		} else {
			H.OK(c, "success insert cart")
		}
	}
}
