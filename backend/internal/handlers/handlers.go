// handlers/auth.go - fixed version
package handlers

import (
    "log"
    "net/http"
    "strings"
    "time"
    
    "github.com/gin-gonic/gin"
    "github.com/techcoder01/go-auth/internal/auth"
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
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
            return
        }
        
        // Validate input
        if input.Email == "" || input.Password == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "email and password required"})
            return
        }
        
        // Authenticate the user
        user, err := models.AuthenticateUser(input.Email, input.Password)
        if err != nil {
            if err == models.ErrUserNotFound || err == models.ErrInvalidCredentials {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
            } else {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "authentication error"})
            }
            return
        }
        
        // Generate JWT token
        token, err := auth.GenerateToken(user.ID, cfg.JWTSecret, 24*time.Hour)
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
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
        return
    }
    
    // Validate input
    if input.Email == "" || input.Password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "email and password required"})
        return
    }
    
    if len(input.Password) < 6 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "password must be at least 6 characters"})
        return
    }
    
    // Create user
    user, err := models.CreateUser(input.Email, input.Password)
    if err != nil {
        // Check for duplicate email error
        if strings.Contains(err.Error(), "already exists") {
            c.JSON(http.StatusBadRequest, gin.H{"error": "email already in use"})
            return
        }
        
        log.Printf("Error creating user: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user", "details": err.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, gin.H{"message": "user created", "user": user})
}

// GetUser handler
func GetUser(c *gin.Context) {
    // Get userID from context (set by AuthMiddleware)
    userID, exists := c.Get("userId")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
        return
    }

    user, err := models.GetUser(userID.(string))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"user": user})
}

// Logout handler
func Logout(c *gin.Context) {
    // Simply return success - token invalidation happens client-side
    c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}