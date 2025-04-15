package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/techcoder01/go-auth/internal/auth"
	"github.com/techcoder01/go-auth/internal/config"
	"github.com/techcoder01/go-auth/internal/handlers"
)

func SetupRoutes(router *gin.Engine, cfg *config.Config) {
	// Health check endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the API"})
	})

	api := router.Group("/api")
	{
		// Auth routes
		api.POST("/login", handlers.Login(cfg))
		api.POST("/register", handlers.Register)
		
		// Protected routes
		protected := api.Group("")
		protected.Use(auth.AuthMiddleware(cfg.JWTSecret))
		{
			protected.GET("/user", handlers.GetUser)
			protected.POST("/logout", handlers.Logout)
		}
	}
}
