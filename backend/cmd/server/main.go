package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/techcoder01/go-auth/internal/config"
	"github.com/techcoder01/go-auth/internal/database"
	"github.com/techcoder01/go-auth/internal/routes"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	database.InitDB()
	defer database.CloseDB()

	// Check environment
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = cfg.Environment
	}

	// Set Gin mode based on environment
	if environment == "production" {
		gin.SetMode(gin.ReleaseMode)
		log.Printf("Running in production mode")
	} else {
		log.Printf("Running in development mode")
	}

	router := gin.Default()
	routes.SetupRoutes(router, cfg)

	// Get port from environment or config
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.ServerPort
	}
	
	log.Printf("Environment: %s", environment)
	log.Printf("Server starting on port %s", port)
	
	// In production, bind to all interfaces
	if environment == "production" {
		log.Printf("Binding to all interfaces (0.0.0.0)")
		log.Fatal(router.Run("0.0.0.0:" + port))
	} else {
		// In development, use default binding
		log.Fatal(router.Run(":" + port))
	}
}