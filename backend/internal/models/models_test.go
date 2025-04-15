package models

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt" // 👈 Add this line
	"github.com/techcoder01/go-auth/internal/database"
)

func TestHashPassword(t *testing.T) {
	password := "secret"
	hash, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEqual(t, password, hash)
}

func TestCheckPasswordHash(t *testing.T) {
	password := "secret"
	hash, _ := HashPassword(password)
	
	assert.True(t, CheckPasswordHash(password, hash))
	assert.False(t, CheckPasswordHash("wrong", hash))
}

func TestAuthenticateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	database.DB = db

	// Setup credentials
	email := "test@example.com"
	password := "password"

	// Generate hash
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// Mock DB row
	rows := sqlmock.NewRows([]string{"id", "email", "password"}).
		AddRow("1", email, string(hashedPassword))

	mock.ExpectQuery("SELECT id, email, password FROM users WHERE email = ?").
		WithArgs(email).
		WillReturnRows(rows)

	// Call function
	_, err = AuthenticateUser(email, password)
	assert.NoError(t, err)
}

