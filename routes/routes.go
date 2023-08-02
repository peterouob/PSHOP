package routes

import (
	"PSHOP/service/session"
	service "PSHOP/service/token"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/login", session.Login)
	r.POST("/create", session.Create)
	tk := r.Group("tk")
	{
		tk.POST("/login", service.Login)
	}
	t := r.Group("/testSession")
	{
		//經測試後發現POST無法獲得session,GET可以
		t.GET("/set", session.Set)
		t.GET("/get", session.Get)
	}
}
