package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authorization(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userData, _ := c.MustGet("userData").(jwt.MapClaims)
		userRole := userData["role"].(string)

		for _, role := range roles {
			if role == userRole {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "you are not authorized to access this data",
		})
	}
}
