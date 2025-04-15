package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
	ServerPort  string
	JWTSecret   string
	JWTTTL      time.Duration
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cfg := &Config{
		Environment: getEnv("ENVIRONMENT", "development"),
		ServerPort:  getEnv("PORT", "8080"),
		JWTSecret:   getEnv("JWT_SECRET", "default-secret-key"),
	}

	ttl, err := time.ParseDuration(getEnv("JWT_TTL", "24h"))
	if err != nil {
		return nil, err
	}
	cfg.JWTTTL = ttl

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
