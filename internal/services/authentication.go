package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zitadel/oidc/v2/pkg/oidc"
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

// GetOrRegisterUserFromUserInfo returns a user from user info.
func (a *Authentication) GetOrRegisterUserFromUserInfo(ctx context.Context, userinfo oidc.UserInfo) (*db.User, error) {
	u, err := a.usvc.ByExternalID(ctx, a.pool, userinfo.Subject)
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeNotFound, "could not get user from external id: %s", err)
	}
	role := models.RoleUser

	guestRole, _ := a.usvc.authzsvc.RoleByName(models.RoleGuest)
	superAdminRole, _ := a.usvc.authzsvc.RoleByName(models.RoleSuperAdmin)

	cfg := internal.Config()

	superAdmin, err := a.usvc.ByEmail(ctx, a.pool, cfg.SuperAdmins.DefaultEmail)
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodePrivate, "could not get admin user %s: %s", cfg.SuperAdmins.DefaultEmail, err)
	}

	// superAdmin is registered without id since an account needs to exist beforehand (created via initial-data, for any env)
	if userinfo.Email == cfg.SuperAdmins.DefaultEmail && superAdmin.ExternalID == "" {
		superAdmin.ExternalID = userinfo.Subject
		superAdmin, err = superAdmin.Update(ctx, a.pool) // TODO external ID is not editable via services but should be. if params external id is set we just ensure caller is admin.
		if err != nil {
			return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "could not update super admin external ID after first login %s: %s", cfg.SuperAdmins.DefaultEmail, err)
		}
		// continue as normal to update superAdmin with updated info.
		// superAdmin account can be changed on demand via SUPERADMIN_EMAIL and info will always be synced with auth server
		u = superAdmin
	}

	// create user on first login
	if u == nil {
		if auth, ok := userinfo.Claims["auth"].(map[string]any); ok {
			if isAdmin, _ := auth["is_admin"].(bool); isAdmin {
				role = models.RoleAdmin
			}
		}

		if !userinfo.EmailVerified {
			role = models.RoleGuest
		}

		u, err = a.usvc.Register(ctx, a.pool, UserRegisterParams{
			Username:   userinfo.PreferredUsername,
			Email:      userinfo.Email,
			ExternalID: userinfo.Subject,
			FirstName:  pointers.New(userinfo.GivenName),
			LastName:   pointers.New(userinfo.FamilyName),
			Role:       role,
		})
		if err != nil {
			return nil, internal.WrapErrorf(err, internal.ErrorCodeNotFound, "could not get register user from provider: %s", err)
		}
	}

	// check if user is superadmin
	if ee := cfg.SuperAdmins.Emails; ee != nil {
		for _, email := range strings.Split(*ee, " ") {
			if u.Email != email {
				continue
			}
			if role, _ := a.usvc.authzsvc.RoleByRank(u.RoleRank); role.Rank == superAdminRole.Rank {
				continue
			}
			u, err = a.usvc.UpdateUserAuthorization(ctx, a.pool, u.UserID.String(), superAdmin, &models.UpdateUserAuthRequest{Role: &superAdminRole.Name})
			if err != nil {
				return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "could not update superadmin role: %s", err)
			}
		}
	}

	// update guest when verified
	if u.RoleRank == guestRole.Rank && userinfo.EmailVerified {
		u, err = a.usvc.UpdateUserAuthorization(ctx, a.pool, u.UserID.String(), superAdmin, &models.UpdateUserAuthRequest{Role: &guestRole.Name})
		if err != nil {
			return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "could not update user auth after email verification: %s", err)
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
