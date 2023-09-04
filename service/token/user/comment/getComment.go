package comment

import (
	"PSHOP/model/database/mysql"
	"PSHOP/model/user"
	H "PSHOP/utils/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetComment(c *gin.Context) {
	var good []user.GoodInfo
	mysql.Db.Preload("Comment", func(g *gorm.DB) *gorm.DB {
		return g.Preload("Replay")
	}).Find(&good)
	H.OK(c, good)
}
