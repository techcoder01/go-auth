package auth

import (
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
    secret := "test-secret"
    userID := "test-user"
    ttl := time.Hour

    // Generate token
    token, err := GenerateToken(userID, secret, ttl)

    // Assert no error occurred
    assert.NoError(t, err)

    // Assert that the token is not empty
    assert.NotEmpty(t, token)
}

func TestParseToken(t *testing.T) {
    secret := "test-secret"
    userID := "test-user"
    ttl := time.Hour

    // Generate token
    token, err := GenerateToken(userID, secret, ttl)

    // Assert no error occurred
    assert.NoError(t, err)

    // Parse the token
    parsedClaims, err := ParseToken(token, secret)

    // Assert no error and the parsed token's userID matches
    assert.NoError(t, err)
    assert.Equal(t, userID, parsedClaims.UserID)
}
