package routes

import (
	"PSHOP/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/login", service.Login)
	r.POST("/create", service.Create)
}
