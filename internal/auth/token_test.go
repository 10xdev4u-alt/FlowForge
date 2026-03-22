package auth

import (
	"testing"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestJWT(t *testing.T) {
	userID := uuid.New()
	secret := "test-secret"

	token, err := GenerateToken(userID, secret)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	claims, err := VerifyToken(token, secret)
	assert.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)

	_, err = VerifyToken(token, "wrong-secret")
	assert.Error(t, err)
}
