package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oloomoses/go-auth/internals/model"
	"github.com/oloomoses/go-auth/internals/repository"
)

var repo repository.UserRepo

func SignUp(c *gin.Context) {
	var input model.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if input.Username == "" || input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username Or Password cannot be empty"})
		return
	}

	user, err := repo.CreateUser(input)

	if err != nil {
		if err.Error() == "username already exits" {
			c.JSON(http.StatusConflict, gin.H{"error": "username already exits"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Message": "User created",
		"user":    user,
	})
}
