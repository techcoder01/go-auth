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

	// Set Gin mode
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	routes.SetupRoutes(router, cfg)

	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.ServerPort
	}
	log.Printf("Server running on port %s", port)
	log.Fatal(router.Run(":" + port))
}
