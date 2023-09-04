package routes

import (
	"PSHOP/middleware"
	"PSHOP/service"
	"PSHOP/service/good"
	"PSHOP/service/session"
	"PSHOP/service/token"
	user2 "PSHOP/service/token/user"
	"PSHOP/service/token/user/comment"
	H "PSHOP/utils/http"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.Use(H.Cors)
	r.GET("/", good.BlockAll)
	r.POST("/login", session.Login)
	r.POST("/create", session.Create)
	r.POST("/logout", session.Logout)
	r.GET("/find/:class", good.Block)
	r.GET("/:goodId", good.GetGood)
	//r.GET("/comment", user2.GetComment)
	s := r.Group("session")
	s.Use(middleware.SessionAuth())
	{
		s.GET("/", session.Backbord)
	}
	tk := r.Group("tk")
	{
		tk.POST("/login", serviceToken.Login)
		tk.POST("/create", serviceToken.CreateUser)
		rauth := r.Group("/rauth")
		rauth.Use(middleware.RDBAuthToken())
		{
			rauth.GET("/", service.Backbord)
			rauth.POST("/logout", serviceToken.Logout)
			//rauth.GET("/refresh", serviceToken.RefreshToken)
			rauth.POST("/createtodo", user2.Create)
			rauth.POST("/:goodId/addcomment", comment.AddComment)
			rauth.POST("/addCart", user2.Add)
		}
	}
}
