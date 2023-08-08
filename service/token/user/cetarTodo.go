package user

import (
	"PSHOP/utils"
	H "PSHOP/utils/http"
	"github.com/gin-gonic/gin"
)

type Todo struct {
	Title string
}

func Create(c *gin.Context) {
	var todo *Todo
	if err := c.ShouldBind(&todo); err != nil {
		H.Fail(c, "bind todo failed")
		return
	}
	tokenAuth, err := utils.ExtractTokenMetadata(c.Request)
	if err != nil {
		H.Fail(c, err.Error())
		return
	}
	_, err = H.FectApi(tokenAuth)
	if err != nil {
		H.Fail(c, "unauthorized")
	}
	H.OK(c, todo)
}
