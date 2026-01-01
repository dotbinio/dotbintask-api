package middleware

import (
	"net/http"
	"strings"

	"github.com/dotbinio/taskwarrior-api/internal/auth"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware creates a Gin middleware for token authentication
func AuthMiddleware(validator *auth.TokenValidator) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "missing authorization header",
				"code":  "MISSING_AUTH_HEADER",
			})
			c.Abort()
			return
		}

		// Check if it's a Bearer token
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization header format (expected 'Bearer <token>')",
				"code":  "INVALID_AUTH_FORMAT",
			})
			c.Abort()
			return
		}

		token := parts[1]

		// Validate token
		if err := validator.Validate(token); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authentication token",
				"code":  "INVALID_TOKEN",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
