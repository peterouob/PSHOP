package session

import (
	"PSHOP/dao/sessions"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Set(c *gin.Context) {
	session, _ := sessions.GetSession(c, c.Request)
	session.Values["user"] = "hey"
	session.Sync()
	fmt.Println("setting session successfully", session.Values["user"])
	c.JSON(200, session.Values["user"])
}

func Get(c *gin.Context) {
	session, _ := sessions.GetSession(c, c.Request)
	jsonstr, _ := json.Marshal(session.Values["user"])
	c.JSON(200, string(jsonstr))
}
