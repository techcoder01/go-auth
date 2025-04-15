package handlers

import (
	"testing"               // Used for test functions
	"github.com/stretchr/testify/assert" // Used for assertions in tests
	"github.com/techcoder01/go-auth/internal/models" // Required for the models
	"github.com/DATA-DOG/go-sqlmock" // For mocking the database queries
	"golang.org/x/crypto/bcrypt" // 👈 Add this line
	"github.com/techcoder01/go-auth/internal/database"  // 👈 ye import zaroori hai
)


func TestAuthenticateUser(t *testing.T) {
	// Mock database setup
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	database.DB = db // Assign mock DB

	// User credentials
	email := "test@example.com"
	password := "password"

	// Generate valid bcrypt hash of password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// Setup expected DB response
	rows := sqlmock.NewRows([]string{"id", "email", "password"}).
		AddRow("1", email, string(hashedPassword))

	mock.ExpectQuery("SELECT id, email, password FROM users WHERE email = ?").
		WithArgs(email).
		WillReturnRows(rows)

	// Call the function
	_, err = models.AuthenticateUser(email, password)
	assert.NoError(t, err, "Authentication should succeed")
}

