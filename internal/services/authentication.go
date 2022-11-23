package services

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type MyAppClaims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Authentication struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
	usvc   *User
}

func NewAuthentication(logger *zap.Logger, usvc *User, pool *pgxpool.Pool) *Authentication {
	return &Authentication{
		logger: logger,
		usvc:   usvc,
		pool:   pool,
	}
}

// GetUserFromAccessToken returns a user from a token.
func (a *Authentication) GetUserFromAccessToken(ctx context.Context, token string) (*db.User, error) {
	return &db.User{}, nil
}

// GetUserFromAPIKey returns a user from an api key.
func (a *Authentication) GetUserFromAPIKey(ctx context.Context, apiKey string) (*db.User, error) {
	return a.usvc.UserByAPIKey(ctx, a.pool, apiKey)
}

// CreateAccessTokenForUser creates a new token for a user.
func (a *Authentication) CreateAccessTokenForUser(ctx context.Context, user *db.User) (string, error) {
	issuer := os.Getenv("DOMAIN")
	claims := MyAppClaims{
		Email:    user.Email,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // mandatory
			Issuer:    issuer,                                             // mandatory
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   user.ExternalID,
			// ID:        "1", // to explicitly revoke tokens. No longer stateless
			Audience: []string{"myapp"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(os.Getenv("SIGNING_KEY"))
	if err != nil {
		return "", errors.Wrap(err, "could not sign token")
	}
	fmt.Printf("%v :error is %v\n", ss, err)

	return ss, nil
}

// CreateAccessTokenForUser creates a new API key for a user.
func (a *Authentication) CreateAPIKeyForUser(ctx context.Context, user *db.User) (*db.UserAPIKey, error) {
	uak, err := a.usvc.CreateAPIKey(ctx, a.pool, user)
	if err != nil {
		return nil, errors.Wrap(err, "usvc.CreateAPIKey")
	}

	return uak, nil
}

// ParseToken returns a token string claims.
func (a *Authentication) ParseToken(ctx context.Context, tokenString string, claim string) (*MyAppClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyAppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return os.Getenv("SIGNING_KEY"), nil
	})

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	claims, ok := token.Claims.(*MyAppClaims)
	if ok && token.Valid {
		fmt.Printf("%v %v", claims.Email, claims.Username)
	} else {
		return nil, errors.Wrap(err, "could not parse token string")
	}

	return claims, nil
}
