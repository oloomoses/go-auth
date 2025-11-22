package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/oloomoses/go-auth/internals/auth"
)

func RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("jwt")

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "login required"})
			return
		}

		claims, err := auth.ValidateToken(token)

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Bad token"})
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
