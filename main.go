package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

var (
	addr string
)

func main() {
	r := gin.New()
	addr = os.Getenv("ADDR")
	if addr == "" {
		addr = "8081"
	}
	r.Run(":" + addr)
}

//utils.Config.GetString("mysql.dsn")
