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
	log.Println("Starting application...")

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Println("Config loaded successfully")

	// Initialize database
	log.Println("Initializing database connection...")
	database.InitDB()
	defer database.CloseDB()
	log.Println("Database connection established")

	// Set up router
	router := gin.Default()
	routes.SetupRoutes(router, cfg)
	log.Println("Routes configured")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is required by Render")
	}

	log.Printf("Environment: %s", os.Getenv("ENVIRONMENT"))
	log.Printf("Attempting to start server on port %s", port)

	err = router.Run("0.0.0.0:" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}