package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/techcoder01/go-auth/internal/auth"
	"github.com/techcoder01/go-auth/internal/config"
	"github.com/techcoder01/go-auth/internal/handlers"
	"time"
)

func SetupRoutes(router *gin.Engine, cfg *config.Config) {
	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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