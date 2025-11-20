package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oloomoses/go-auth/internals/handler"
)

func main() {
	r := gin.Default()

	userAuth := handler.NewUserHandler()

	r.POST("/signup", userAuth.SignUp)
	r.POST("/signin", userAuth.Login)

	r.Run(":8080")
}
