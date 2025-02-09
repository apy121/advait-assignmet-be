package middleware

import (
	"awesomeProject/internal/service"
	"awesomeProject/internal/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// JWTAuthMiddleware checks if the request has a valid JWT token
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		token := tokenString[7:] // Strip 'Bearer ' from token string
		if len(token) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is absent, please provide correct token"})
			c.Abort()
			return
		}
		claims, err := service.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token provided"})
			c.Abort()
			return
		}

		// Check if the token has expired
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Token does not contain expiration claim"})
			c.Abort()
			return
		}
		isExist := storage.TokenStore.CheckToken(token)
		if !isExist {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token or token is revoked"})
			c.Abort()
			return
		}

		c.Set("user", claims["user"])
		c.Next()
	}
}
