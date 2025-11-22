package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oloomoses/go-auth/internals/handler"
	"github.com/oloomoses/go-auth/internals/middleware"
)

func main() {
	r := gin.Default()

	userAuth := handler.NewUserHandler()

	protected := r.Group("api/v1")

	r.POST("/signup", userAuth.SignUp)
	r.POST("/signin", userAuth.Login)
	protected.Use(middleware.RequireLogin())
	protected.GET("/mewith-header", userAuth.MeWithHeader)
	protected.GET("/me", userAuth.Me)
	protected.GET("/logout", userAuth.Logout)

	r.Run(":8080")
}
