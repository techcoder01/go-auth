package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// Mock database connection
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Simulate database operation for user creation
	mock.ExpectExec("INSERT INTO users").
		WithArgs("test@example.com", "password123").
		WillReturnResult(sqlmock.NewResult(1, 1)) // Simulating that 1 row is affected

	// Test CreateUser function with the mock DB
	err = CreateUser(db, "test@example.com", "password123")
	assert.NoError(t, err)

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
