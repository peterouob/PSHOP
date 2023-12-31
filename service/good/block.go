package good

import (
	"PSHOP/model/database/mysql"
	"PSHOP/model/user"
	H "PSHOP/utils/http"
	"github.com/gin-gonic/gin"
)

func BlockAll(c *gin.Context) {
	var goodsList []user.Block
	mysql.Db.Preload("Class").Find(&goodsList)
	H.OK(c, goodsList)
}

func Block(c *gin.Context) {
	name := c.Param("class")
	var goodList []user.Block
	mysql.Db.Where("class_name =?", name).Preload("Class", "goods_class = ?", name).Find(&goodList)
	H.OK(c, goodList)
}
