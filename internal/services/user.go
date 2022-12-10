package services

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

type User struct {
	logger   *zap.Logger
	urepo    repos.User
	authzsvc *Authorization
}

type UserUpdateParams struct {
	FirstName      *string
	LastName       *string
	ID             string
	RequestingUser *db.User
}

type UserUpdateAuthorizationParams struct {
	Rank           *int16
	Scopes         *[]string
	ID             string
	RequestingUser *db.User
}

// NewUser returns a new User service.
func NewUser(logger *zap.Logger, urepo repos.User, authzsvc *Authorization) *User {
	return &User{
		logger:   logger,
		urepo:    urepo,
		authzsvc: authzsvc,
	}
}

// Register registers a user.
// TODO accepts basic parameters instead of user *db.User and everything else is default,
// returns a *db.User. must not pass a db.User here
// IMPORTANT: no endpoint for user creation. Only when coming from auth server.
// we will not support password auth.
func (u *User) Register(ctx context.Context, d db.DBTX, params repos.UserCreateParams) (*db.User, error) {
	defer newOTELSpan(ctx, "User.Register").End()

	// TODO construct db.User and fill missing fields with default roles, etc.
	// instead of passing it directly
	user, err := u.urepo.Create(ctx, d, params)
	if err != nil {
		return nil, errors.Wrap(err, "urepo.Create")
	}

	return user, nil
}

// Update updates a user.
func (u *User) Update(ctx context.Context, d db.DBTX, id string, caller *db.User, params *models.UpdateUserRequest) (*db.User, error) {
	defer newOTELSpan(ctx, "User.Update").End()

	if caller == nil {
		return nil, errors.New("caller cannot be nil")
	}
	if params == nil {
		return nil, errors.New("params cannot be nil")
	}

	user, err := u.urepo.UserByID(ctx, d, id)
	if err != nil {
		return nil, errors.Wrap(err, "urepo.UserByID")
	}

	adminRole, err := u.authzsvc.RoleByName(string(models.RoleAdmin))
	if err != nil {
		return nil, errors.Wrap(err, "authzsvc.RoleByName")
	}

	if user.UserID != caller.UserID &&
		caller.RoleRank < adminRole.Rank {
		return nil, internal.NewErrorf(internal.ErrorCodeUnauthorized, "cannot change another user's information")
	}

	user, err = u.urepo.Update(ctx, d, id, repos.UserUpdateParams{
		FirstName: params.FirstName,
		LastName:  params.LastName,
	})
	if err != nil {
		return nil, errors.Wrap(err, "urepo.Update")
	}

	return user, nil
}

func (u *User) UpdateUserAuthorization(ctx context.Context, d db.DBTX, id string, caller *db.User, params *models.UpdateUserAuthRequest) (*db.User, error) {
	defer newOTELSpan(ctx, "User.UpdateUserAuthorization").End()

	if caller == nil {
		return nil, errors.New("caller cannot be nil")
	}
	if params == nil {
		return nil, errors.New("params cannot be nil")
	}

	user, err := u.urepo.UserByID(ctx, d, id)
	if err != nil {
		return nil, errors.Wrap(err, "urepo.UserByID")
	}

	adminRole, err := u.authzsvc.RoleByName(string(models.RoleAdmin))
	if err != nil {
		return nil, errors.Wrap(err, "authzsvc.RoleByName")
	}

	if caller.RoleRank < adminRole.Rank {
		if user.UserID == caller.UserID { // exit early, though it's not possible to update to something not assigned to self already anyway
			return nil, internal.NewErrorf(internal.ErrorCodeUnauthorized, "cannot update your own authorization information")
		}
	}

	var rank *int16
	if params.Role != nil {
		role, err := u.authzsvc.RoleByName(string(*params.Role))
		if err != nil {
			return nil, errors.Wrap(err, "authzsvc.RoleByName")
		}
		if role.Rank > caller.RoleRank {
			return nil, internal.NewErrorf(internal.ErrorCodeUnauthorized, "cannot set a user rank higher than self")
		}
		if caller.RoleRank < adminRole.Rank {
			if role.Rank < user.RoleRank {
				return nil, internal.NewErrorf(internal.ErrorCodeUnauthorized, "cannot demote a user role")
			}
		}
		rank = &role.Rank
	}

	var scopes *[]string
	if params.Scopes != nil {
		ss := make([]string, 0, len(*params.Scopes))
		for _, s := range *params.Scopes {
			if !slices.Contains(caller.Scopes, string(s)) {
				return nil, internal.NewErrorf(internal.ErrorCodeUnauthorized, "cannot set a scope unassigned to self")
			}

			ss = append(ss, string(s))
		}

		if caller.RoleRank < adminRole.Rank {
			for _, s := range user.Scopes {
				if !slices.Contains(ss, s) {
					return nil, internal.NewErrorf(internal.ErrorCodeUnauthorized, "cannot unassign a user's scope")
				}
			}
		}

		scopes = &ss
	}

	user, err = u.urepo.Update(ctx, d, id, repos.UserUpdateParams{
		Scopes: scopes,
		Rank:   rank,
	})
	if err != nil {
		return nil, errors.Wrap(err, "urepo.Update")
	}

	return user, nil
}

func (u *User) CreateAPIKey(ctx context.Context, d db.DBTX, user *db.User) (*db.UserAPIKey, error) {
	defer newOTELSpan(ctx, "User.CreateAPIKey").End()

	uak, err := u.urepo.CreateAPIKey(ctx, d, user)
	if err != nil {
		return nil, errors.Wrap(err, "urepo.CreateAPIKey")
	}

	return uak, nil
}

// UserByExternalID gets a user by ExternalID.
func (u *User) UserByExternalID(ctx context.Context, d db.DBTX, id string) (*db.User, error) {
	defer newOTELSpan(ctx, "User.UserByExternalID").End()

	user, err := u.urepo.UserByExternalID(ctx, d, id)
	if err != nil {
		return nil, errors.Wrap(err, "urepo.UserByExternalID")
	}

	return user, nil
}

// UserByEmail gets a user by email.
func (u *User) UserByEmail(ctx context.Context, d db.DBTX, email string) (*db.User, error) {
	defer newOTELSpan(ctx, "User.UserByEmail").End()

	user, err := u.urepo.UserByEmail(ctx, d, email)
	if err != nil {
		return nil, errors.Wrap(err, "urepo.UserByEmail")
	}

	return user, nil
}

// UserByUsername gets a user by username.
func (u *User) UserByUsername(ctx context.Context, d db.DBTX, username string) (*db.User, error) {
	defer newOTELSpan(ctx, "User.UserByUsername").End()

	user, err := u.urepo.UserByUsername(ctx, d, username)
	if err != nil {
		return nil, errors.Wrap(err, "urepo.UserByUsername")
	}

	return user, nil
}

// UserByAPIKey gets a user by apiKey.
func (u *User) UserByAPIKey(ctx context.Context, d db.DBTX, apiKey string) (*db.User, error) {
	defer newOTELSpan(ctx, "User.UserByAPIKey").End()

	user, err := u.urepo.UserByAPIKey(ctx, d, apiKey)
	if err != nil {
		return nil, errors.Wrap(err, "urepo.UserByAPIKey")
	}

	return user, nil
}

// TODO
func (u *User) LatestPersonalNotifications(ctx context.Context, d db.DBTX, apiKey string) ([]db.GetUserNotificationsRow, error) {
	// this will also set user.has_new_personal_notifications to false in the same tx
	return []db.GetUserNotificationsRow{}, nil
}

// TODO
func (u *User) LatestGlobalNotifications(ctx context.Context, d db.DBTX, apiKey string) ([]db.GetUserNotificationsRow, error) {
	// this will also set user.has_new_global_notifications to false in the same tx
	return []db.GetUserNotificationsRow{}, nil
}
