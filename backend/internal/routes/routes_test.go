package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/techcoder01/go-auth/internal/config"
)

func TestRoutes(t *testing.T) {
	// Setup
	r := gin.Default()
	cfg := &config.Config{
		JWTSecret: "test-secret",
	}

	// Register routes
	SetupRoutes(r, cfg)

	// Test cases
	tests := []struct {
		name       string
		method     string
		path       string
		wantStatus int
	}{
		{
			name:       "health check",
			method:     http.MethodGet,
			path:       "/",
			wantStatus: http.StatusOK,
		},
		{
			name:       "login route exists",
			method:     http.MethodPost,
			path:       "/api/login",
			wantStatus: http.StatusBadRequest, // Will fail validation
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}