package main

import (
	"awesomeProject/internal/handlers"
	"awesomeProject/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	// Middleware for JWT main
	authMiddleware := middleware.JWTAuthMiddleware()

	r.POST("/signup", handlers.SignUp)
	r.POST("/signin", handlers.SignIn)
	r.GET("/protected", authMiddleware, handlers.Protected)
	r.POST("/revoke", authMiddleware, handlers.RevokeToken)
	r.POST("/refresh", authMiddleware, handlers.RefreshToken)

	r.Run(":8080")
}
