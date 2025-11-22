package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oloomoses/go-auth/internals/model"
	"github.com/oloomoses/go-auth/internals/repository"
)

type UserHandler struct {
	repo repository.UserRepo
}

func NewUserHandler() UserHandler {
	return UserHandler{}
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var input model.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if input.Username == "" || input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username Or Password cannot be empty"})
		return
	}

	input.Username = strings.ToLower(input.Username)

	user, token, err := h.repo.CreateUser(input)

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
		"token":   token,
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var input model.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	username := strings.ToLower(input.Username)
	password := input.Password
	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password cannot be blank"})
		return
	}

	isLoggedIn := h.repo.VerifyPassword(username, password)

	if !isLoggedIn {
		c.JSON(http.StatusExpectationFailed, gin.H{"error": "Login failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Loggin Success!"})
}
