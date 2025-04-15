package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Set environment variables for testing
	os.Setenv("ENVIRONMENT", "development")
	os.Setenv("PORT", "8080")
	os.Setenv("JWT_SECRET", "test-secret")
	os.Setenv("JWT_TTL", "1h") // Set a TTL value for testing

	// Load the config
	cfg, err := LoadConfig()

	// Assert no error occurred
	assert.NoError(t, err)

	// Assert that the config is not nil
	assert.NotNil(t, cfg)

	// Assert the expected values
	assert.Equal(t, "development", cfg.Environment)
	assert.Equal(t, "8080", cfg.ServerPort)
	assert.Equal(t, "test-secret", cfg.JWTSecret)
	assert.Equal(t, time.Hour, cfg.JWTTTL) // Assert that the TTL is 1 hour
}
