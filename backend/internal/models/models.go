// models/user.go - fixed version
package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/techcoder01/go-auth/internal/database"
	"golang.org/x/crypto/bcrypt"
)

// Define custom errors
var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// User struct represents the user model
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"` // Don't include password in JSON responses
}

// HashPassword hashes the given password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPasswordHash compares a hashed password with the given plain text password
func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// AuthenticateUser authenticates the user by checking the credentials
func AuthenticateUser(email, password string) (*User, error) {
	// Create a query to get the user from the database
	var user User
	query := "SELECT id, email, password FROM users WHERE email = $1"
	err := database.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// Check the password
	if !CheckPasswordHash(password, user.Password) {
		return nil, ErrInvalidCredentials
	}

	// Don't return the password
	user.Password = ""
	return &user, nil
}

// GenerateToken generates a JWT token for the user
func GenerateToken(userID string) (string, error) {
	// Define the secret key (in a real application, keep this secret)
	secretKey := "your_secret_key"

	// Create JWT claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// CreateUser creates a new user in the database
func CreateUser(email, password string) (*User, error) {
	// Check if user already exists
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", email).Scan(&count)
	if err != nil {
		log.Printf("Error checking for existing user: %v", err)
		return nil, fmt.Errorf("error checking for existing user: %w", err)
	}
	
	if count > 0 {
		return nil, fmt.Errorf("user with email %s already exists", email)
	}

	// Hash the password
	hash, err := HashPassword(password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	// Create a new user in the database
	var id int64
	query := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id"
	err = database.DB.QueryRow(query, email, hash).Scan(&id)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return nil, err
	}

	// Convert int64 id to string properly
	userID := strconv.FormatInt(id, 10)

	// Return the created user (without password)
	return &User{
		ID:    userID,
		Email: email,
	}, nil
}

// GetUser retrieves a user by their ID
func GetUser(userID string) (*User, error) {
	var user User
	query := "SELECT id, email FROM users WHERE id = $1"
	err := database.DB.QueryRow(query, userID).Scan(&user.ID, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}