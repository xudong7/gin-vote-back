package middlewares

import (
	"community/back/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		token := c.GetHeader("Authorization")

		// If the token is empty, return a 401 Unauthorized response
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Missing Authorization Header",
			})
			c.Abort()
			return
		}

		// Check if the user is authenticated
		username, err := utils.ParseJWT(token)
		// If not, return a 401 Unauthorized response
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Token",
			})
			c.Abort()
			return
		}

		// If the user is authenticated, set the username in the context
		c.Set("username", username)
		c.Next() // Continue to the next middleware
	}
}
