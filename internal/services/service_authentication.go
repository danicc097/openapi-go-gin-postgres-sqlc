package services

import (
	"context"
	"fmt"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
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
	repos  *repos.Repos
	usvc   *User
}

// NewAuthentication returns a new authentication service.
// TODO should we use tx instead of providing pool only.
func NewAuthentication(logger *zap.SugaredLogger, repos *repos.Repos, pool *pgxpool.Pool) *Authentication {
	usvc := NewUser(logger, repos)

	return &Authentication{
		logger: logger,
		repos:  repos,
		usvc:   usvc,
		pool:   pool,
	}
}

// GetUserFromAccessToken returns a user from a token.
func (a *Authentication) GetUserFromAccessToken(ctx context.Context, token string) (*models.User, error) {
	claims, err := a.ParseToken(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	user, err := a.usvc.ByExternalID(ctx, a.pool, claims.Subject)
	if err != nil {
		return nil, internal.WrapErrorf(err, models.ErrorCodeNotFound, "user from token not found")
	}

	return user, nil
}

// GetUserFromAPIKey returns a user from an api key.
func (a *Authentication) GetUserFromAPIKey(ctx context.Context, apiKey string) (*models.User, error) {
	u, err := a.usvc.ByAPIKey(ctx, a.pool, apiKey)
	if err != nil {
		return nil, internal.WrapErrorf(err, models.ErrorCodeNotFound, "user from api key not found")
	}

	if u.UserAPIKeyJoin.ExpiresOn.Before(time.Now()) {
		return nil, internal.WrapErrorf(err, models.ErrorCodeUnauthorized, "api key expired")
	}

	u.UserAPIKeyJoin = nil

	return u, nil
}

// GetOrRegisterUserFromUserInfo returns a user from user info.
func (a *Authentication) GetOrRegisterUserFromUserInfo(ctx context.Context, userinfo oidc.UserInfo) (*models.User, error) {
	u, _ := a.usvc.ByExternalID(ctx, a.pool, userinfo.Subject)
	role := models.RoleUser

	guestRole := a.usvc.authzsvc.RoleByName(models.RoleGuest)

	cfg := internal.Config

	superAdmin, err := a.usvc.ByEmail(ctx, a.pool, cfg.SuperAdmin.DefaultEmail)
	if err != nil {
		return nil, internal.WrapErrorf(err, models.ErrorCodePrivate, "could not get admin user %s", cfg.SuperAdmin.DefaultEmail)
	}

	// superAdmin is registered without id since an account needs to exist beforehand (created via initial-data, for any env)
	if userinfo.Email == cfg.SuperAdmin.DefaultEmail && superAdmin.ExternalID == "" {
		// external ID is not editable via services.
		superAdmin, err = a.repos.User.Update(ctx, a.pool, superAdmin.UserID, &models.UserUpdateParams{
			ExternalID: pointers.New(userinfo.Subject),
		})
		if err != nil {
			return nil, internal.WrapErrorf(err, models.ErrorCodeUnknown, "could not update super admin external ID after first login %s", cfg.SuperAdmin.DefaultEmail)
		}
		// continue as normal to update superAdmin if necessary.
		// default superAdmin account can be changed on startup via SUPERADMIN_EMAIL and info will always be synced with auth server
		u = superAdmin
	}

	// create user on first login
	if u == nil {
		// userInfo.AppendClaims(AuthClaim, map[string]interface{}{
		// 	"is_admin": user.IsAdmin(),
		// })
		if auth, ok := userinfo.Claims["auth"].(map[string]any); ok {
			if isAdmin, _ := auth["is_admin"].(bool); isAdmin {
				role = models.RoleSuperAdmin
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
			return nil, internal.WrapErrorf(err, models.ErrorCodeUnknown, "could not register user from provider")
		}
	}

	// update guest when verified
	if u.RoleRank == guestRole.Rank && userinfo.EmailVerified {
		u, err = a.usvc.UpdateUserAuthorization(ctx, a.pool, u.UserID, CtxUser{User: superAdmin}, &models.UpdateUserAuthRequest{
			Role:   &guestRole.Name,
			Scopes: nil,
		})
		if err != nil {
			return nil, internal.WrapErrorf(err, models.ErrorCodeUnknown, "could not update user auth after email verification")
		}
	}

	// update out of sync non editable fields
	if u.Email != userinfo.Email || u.Username != userinfo.PreferredUsername {
		u.Email = userinfo.Email
		u.Username = userinfo.PreferredUsername
		u, err = u.Update(ctx, a.pool)
		if err != nil {
			return nil, internal.WrapErrorf(err, models.ErrorCodeUnknown, "could not update out of sync userinfo")
		}
	}

	return u, nil
}

// CreateAccessTokenForUser creates a new token for a user.
func (a *Authentication) CreateAccessTokenForUser(ctx context.Context, user *models.User) (string, error) {
	cfg := internal.Config
	claims := AppClaims{
		Email:    user.Email,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), // mandatory
			Issuer:    cfg.OIDC.Issuer,                                        // mandatory
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   user.ExternalID,
			// ID:        "1", // to explicitly revoke tokens. No longer stateless
			Audience: []string{"myapp"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(cfg.SigningKey)) // cannot handle strings for some reason https://github.com/dgrijalva/jwt-go/issues/65
	if err != nil {
		return "", fmt.Errorf("could not sign token: %w", err)
	}

	return ss, nil
}

// CreateAPIKeyForUser creates a new API key for a user.
func (a *Authentication) CreateAPIKeyForUser(ctx context.Context, user *models.User) (*models.UserAPIKey, error) {
	uak, err := a.usvc.CreateAPIKey(ctx, a.pool, user)
	if err != nil {
		return nil, fmt.Errorf("usvc.CreateAPIKey: %w", err)
	}

	return uak, nil
}

// ParseToken returns a token string claims.
func (a *Authentication) ParseToken(ctx context.Context, token string) (*AppClaims, error) {
	cfg := internal.Config
	jwtToken, err := jwt.ParseWithClaims(token, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.SigningKey), nil // can't handle string in signing keys here either
	})
	if err != nil || jwtToken == nil {
		return nil, fmt.Errorf("could not parse token: %w", err)
	}

	if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
	}

	claims, ok := jwtToken.Claims.(*AppClaims)
	if ok && jwtToken.Valid {
		a.logger.Debugf("Token valid for user: %v", claims.Email)
	} else {
		return nil, fmt.Errorf("could not parse token string: %w", err)
	}

	return claims, nil
}
