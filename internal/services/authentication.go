package services

import (
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Authentication struct {
	Logger *zap.Logger
	Pool   *pgxpool.Pool
}

type AuthenticationService interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password string, hashedPassword string) error
}

// TODO for jwt, refresh token, redis...
// https://developer.vonage.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr
// we would need a redis repo for auth

// HashPassword returns the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

// VerifyPassword checks if the provided password is correct or not
func VerifyPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
