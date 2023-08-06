package service

import (
	"PSHOP/utils/http"
	"github.com/gin-gonic/gin"
)

func Backbord(c *gin.Context) {
	H.OK(c, "hey")
}
