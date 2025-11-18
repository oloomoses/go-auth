package main

import (
	"fmt"

	"github.com/oloomoses/go-auth/internals/handler"
	"github.com/oloomoses/go-auth/internals/model"
)

func main() {
	user := model.User{
		Username: "Moses",
		Password: "1234",
	}

	createdUser, _ := handler.CreateUser(user)

	fmt.Println(createdUser)

}
