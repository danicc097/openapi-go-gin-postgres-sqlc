package services

import (
	"context"
	"fmt"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type AppClaims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Authentication struct {
	pool   *pgxpool.Pool
	logger *zap.SugaredLogger
	usvc   *User
}

// NewAuthentication returns a new authentication service.
// TODO should we use tx instead of providing pool only
func NewAuthentication(logger *zap.SugaredLogger, usvc *User, pool *pgxpool.Pool) *Authentication {
	return &Authentication{
		logger: logger,
		usvc:   usvc,
		pool:   pool,
	}
}

// GetUserFromAccessToken returns a user from a token.
func (a *Authentication) GetUserFromAccessToken(ctx context.Context, token string) (*db.User, error) {
	claims, err := a.ParseToken(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	user, err := a.usvc.ByEmail(ctx, a.pool, claims.Email)
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeNotFound, "user from token not found: %s", err)
	}

	return user, nil
}

// GetUserFromAPIKey returns a user from an api key.
func (a *Authentication) GetUserFromAPIKey(ctx context.Context, apiKey string) (*db.User, error) {
	return a.usvc.ByAPIKey(ctx, a.pool, apiKey)
}

// GetOrRegisterUserFromProvider returns a user from a JWT.
func (a *Authentication) GetOrRegisterUserFromProvider(ctx context.Context, token map[string]any) (*db.User, error) {
	u, err := a.usvc.ByExternalID(ctx, a.pool, token["sub"].(string))
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeNotFound, "could not get user from external id: %s", err)
	}
	if u == nil {
		role := models.RoleUser
		if isAdmin, _ := token["is_admin"].(bool); isAdmin { // TODO is_admin not returned currently
			role = models.RoleAdmin
		}

		// {"email":"admin@admin.com","email_verified":true,"family_name":"Admin","given_name":"Mr","locale":"de","name":"Mr Admin","preferred_username":"admin","sub":"id1"}
		u, err = a.usvc.Register(ctx, a.pool, UserRegisterParams{
			Username:   token["preferred_username"].(string),
			Email:      token["email"].(string),
			ExternalID: token["sub"].(string),
			FirstName:  pointers.New(token["given_name"].(string)),
			LastName:   pointers.New(token["family_name"].(string)),
			Role:       role,
		})
		if err != nil {
			return nil, internal.WrapErrorf(err, internal.ErrorCodeNotFound, "could not get register user from provider: %s", err)
		}
	}

	return u, nil
}

// CreateAccessTokenForUser creates a new token for a user.
func (a *Authentication) CreateAccessTokenForUser(ctx context.Context, user *db.User) (string, error) {
	cfg := internal.Config()
	claims := AppClaims{
		Email:    user.Email,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // mandatory
			Issuer:    cfg.OIDC.Issuer,                                    // mandatory
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   user.ExternalID,
			// ID:        "1", // to explicitly revoke tokens. No longer stateless
			Audience: []string{"myapp"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(cfg.SigningKey)
	if err != nil {
		return "", fmt.Errorf("could not sign token: %w", err)
	}

	return ss, nil
}

// CreateAPIKeyForUser creates a new API key for a user.
func (a *Authentication) CreateAPIKeyForUser(ctx context.Context, user *db.User) (*db.UserAPIKey, error) {
	uak, err := a.usvc.CreateAPIKey(ctx, a.pool, user)
	if err != nil {
		return nil, fmt.Errorf("usvc.CreateAPIKey: %w", err)
	}

	return uak, nil
}

// ParseToken returns a token string claims.
func (a *Authentication) ParseToken(ctx context.Context, token string) (*AppClaims, error) {
	cfg := internal.Config()
	jwtToken, err := jwt.ParseWithClaims(token, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return cfg.SigningKey, nil
	})

	if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
	}

	claims, ok := jwtToken.Claims.(*AppClaims)
	if ok && jwtToken.Valid {
		fmt.Printf("%v %v", claims.Email, claims.Username)
	} else {
		return nil, fmt.Errorf("could not parse token string: %w", err)
	}

	return claims, nil
}
