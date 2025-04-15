package models

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/techcoder01/go-auth/internal/database"
	"github.com/dgrijalva/jwt-go"
	"errors"
	"time"
	"strconv"
	"database/sql" // Add this import
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
	Password string `json:"password"`
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
	query := "SELECT id, email, password FROM users WHERE email = ?"
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
	// Hash the password
	hash, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Create a new user in the database
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	res, err := database.DB.Exec(query, email, hash)
	if err != nil {
		return nil, err
	}

	// Get the last inserted ID
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Convert int64 id to string properly
	userID := strconv.FormatInt(id, 10)

	// Return the created user
	return &User{
		ID:    userID,
		Email: email,
	}, nil
}

// GetUser retrieves a user by their ID
func GetUser(userID string) (*User, error) {
	var user User
	query := "SELECT id, email FROM users WHERE id = ?"
	err := database.DB.QueryRow(query, userID).Scan(&user.ID, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
