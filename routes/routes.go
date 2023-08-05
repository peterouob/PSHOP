package routes

import (
	"PSHOP/middleware"
	"PSHOP/service"
	"PSHOP/service/session"
	"PSHOP/service/token"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/login", session.Login)
	r.POST("/create", session.Create)
	r.POST("/logout", session.Logout)
	s := r.Group("session")
	s.Use(middleware.SessionAuth())
	{
		s.GET("/", session.Backbord)
	}
	tk := r.Group("tk")
	{
		tk.POST("/login", serviceToken.Login)
		rauth := r.Group("/rauth")
		rauth.Use(middleware.RDBAuthToken())
		{
			rauth.GET("/", service.Backbord)
		}
	}
	t := r.Group("/testSession")
	{
		//經測試後發現POST無法獲得session,GET可以
		t.GET("/set", session.Set)
		t.GET("/get", session.Get)
	}
}
