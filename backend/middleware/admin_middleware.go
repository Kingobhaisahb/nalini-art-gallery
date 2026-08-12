package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		role, exists := c.Get("role")

		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "User role not found",
			})
			c.Abort()
			return
		}

		userRole, ok := role.(string)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid user role",
			})
			c.Abort()
			return
		}

		if userRole != "ADMIN" {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Admin access required",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}