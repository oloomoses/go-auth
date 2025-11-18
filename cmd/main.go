package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oloomoses/go-auth/internals/handler"
)

func main() {
	r := gin.Default()

	r.POST("/signup", handler.SignUp)

	r.Run(":8080")
}
