package good

import (
	"PSHOP/model/database/mysql"
	"PSHOP/model/user"
	H "PSHOP/utils/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetGood(c *gin.Context) {
	var goodinfo user.GoodInfo
	goodId := c.Param("goodId")
	result := mysql.Db.Where("good_id = ?", goodId).Preload("Comment", func(g *gorm.DB) *gorm.DB {
		return g.Preload("Replay")
	}).First(&goodinfo)
	if result.RowsAffected == 0 {
		H.FailStatus(c, http.StatusNotFound, "商品不存在")
	}
	H.OK(c, goodinfo)
}
