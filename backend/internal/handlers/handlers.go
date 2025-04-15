package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/techcoder01/go-auth/internal/config"
	"github.com/techcoder01/go-auth/internal/models"
	
)

// Login handler
func Login(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		// Authenticate the user
		user, err := models.AuthenticateUser(input.Email, input.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		// Generate JWT token
		token, err := models.GenerateToken(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
	}
}

// Register handler
func Register(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// Create user
	user, err := models.CreateUser(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created", "user": user})
}

// GetUser handler
func GetUser(c *gin.Context) {
	userID := c.Param("id")
	user, err := models.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Logout handler
func Logout(c *gin.Context) {
	// Invalidate the JWT or session (not implemented here)
	c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
