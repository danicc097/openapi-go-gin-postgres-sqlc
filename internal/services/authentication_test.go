package services_test

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/format"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := format.RandomString(6)

	hashedPassword1, err := services.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	err = services.VerifyPassword(password, hashedPassword1)
	require.NoError(t, err)

	wrongPassword := format.RandomString(6)
	err = services.VerifyPassword(wrongPassword, hashedPassword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassword2, err := services.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}
