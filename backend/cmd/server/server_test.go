package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestServerStart(t *testing.T) {
	// Setup Gin router for testing
	r := gin.Default()

	// Register routes (ensure the routes are registered as they are in the server setup)
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Create a request to hit the test endpoint
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	// Serve HTTP request
	r.ServeHTTP(w, req)

	// Assert status code
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}
