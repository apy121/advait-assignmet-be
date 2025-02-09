package handlers

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/service"
	"awesomeProject/internal/storage"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

var userStore = storage.NewUserStorage()

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)
	// Generate a unique ID (you could use UUID for better uniqueness)
	user.ID = generateUniqueID()
	if err := userStore.CreateUser(user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
}

func SignIn(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	authUser, err := userStore.AuthenticateUser(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	token, err := service.CreateToken(authUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	storage.TokenStore.StoreToken(models.Token{
		Value:  token,
		UserId: authUser.ID,
	})
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Protected checks if the request is authorized
func Protected(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"message": "Protected endpoint response", "user": user})
}

// RevokeToken removes the token from the store
func RevokeToken(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token := tokenString[7:] // Strip 'Bearer ' from token string
	if err := storage.TokenStore.RevokeToken(token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not revoke token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Token revoked successfully"})
}

// RefreshToken refreshes an existing token
func RefreshToken(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token := tokenString[7:] // Strip 'Bearer ' from token string
	claims, err := service.ParseToken(token)

	errw := storage.TokenStore.RevokeToken(token)
	if errw != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred in revoking token"})
		return
	}

	user := models.User{}
	if userMap, ok := claims["user"].(map[string]interface{}); ok {
		user = models.User{
			ID:       userMap["id"].(string),
			Email:    userMap["email"].(string),
			Password: userMap["password"].(string),
		}
	}
	newToken, err := service.CreateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	storage.TokenStore.StoreToken(models.Token{
		Value:  newToken,
		UserId: user.ID,
	})

	c.JSON(http.StatusOK, gin.H{"token": newToken})
}

// Helper function to generate a unique ID (simplified for example purposes)
func generateUniqueID() string {
	// In a real application, use something like github.com/google/uuid
	return "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
}
