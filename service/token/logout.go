package serviceToken

import (
	redisdao "PSHOP/model/database/redis"
	"PSHOP/utils"
	H "PSHOP/utils/http"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	au, err := utils.ExtractTokenMetadata(c.Request)
	if err != nil {
		H.Fail(c, "take token have error :"+err.Error())
	}
	deleted, deletErr := redisdao.DeleteTokenAuth(au.AccessUid)
	if deletErr != nil || deleted == 0 {
		H.Fail(c, "unauthorized")
	}
	H.RemoveCookie(c, "accessToken")
	H.OK(c, "success logout")
}
