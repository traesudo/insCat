package main

import (
	"github.com/gin-gonic/gin"
	"insCat/api_server/user"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 64 << 20
	r.Use()
	r.POST("/api/login", user.LoginUser)
	r.Run()
}
