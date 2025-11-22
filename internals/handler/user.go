package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oloomoses/go-auth/internals/auth"
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

	setCookie(c, token)

	c.JSON(http.StatusCreated, gin.H{
		"Message": "User created",
		"user":    user,
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

	token, isLoggedIn := h.repo.VerifyPassword(username, password)

	if !isLoggedIn {
		c.JSON(http.StatusExpectationFailed, gin.H{"error": "Login failed"})
		return
	}

	setCookie(c, token)

	c.JSON(http.StatusOK, gin.H{"Message": "Loggin Success!"})
}

func (h *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "", false, true)
	c.JSON(200, gin.H{"Message": "Bye"})
}

func (h *UserHandler) MeWithHeader(c *gin.Context) {
	token, err := c.Cookie("jwt")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no cookie"})
		return
	}

	claim, err := auth.ValidateToken(token)

	if err != nil {
		c.JSON(401, gin.H{"error": "Bad Token"})
		return
	}

	c.JSON(200, gin.H{"Message": "Hello " + strings.ToTitle(claim.Username)})
}

func (h *UserHandler) Me(c *gin.Context) {
	username := c.GetString("username")

	c.JSON(200, gin.H{"Hello ": username})
}

func setCookie(c *gin.Context, token string) {
	c.SetCookie("jwt", token, 24*3600, "/", "", false, true)
}
