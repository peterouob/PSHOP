package main

import (
	"PSHOP/model/dao/mysql"
	"PSHOP/model/dao/redis"
	"PSHOP/routes"
	"github.com/gin-gonic/gin"
	"os"
)

var (
	addr string
)

func main() {
	r := gin.Default()
	go func() {
		mysql.MysqlInit()
		redisdao.RedisInit()
	}()
	addr = os.Getenv("ADDR")
	if addr == "" {
		addr = "8081"
	}
	routes.SetupRouter(r)
	r.Run(":" + addr)
}
