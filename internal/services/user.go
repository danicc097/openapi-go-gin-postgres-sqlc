package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

type User struct {
	logger           *zap.SugaredLogger
	urepo            repos.User
	notificationrepo repos.Notification
	authzsvc         *Authorization
}

// NOTE: the most important distinction about repositories is that they represent collections of entities. They do not represent database storage or caching or any number of technical concerns. Repositories represent collections. How you hold those collections is simply an implementation detail.
// Repo should not be aware of models Role and Scope, its conversion or its default values. That's all
// for upper layers convenience. e.g roles: entity uses rank internally. Repo should not care about mappings to user-friendly names.
type UserRegisterParams struct {
	Username   string
	Email      string
	FirstName  *string
	LastName   *string
	ExternalID string
	Scopes     []models.Scope
	Role       models.Role
}

// NewUser returns a new User service.
func NewUser(logger *zap.SugaredLogger, urepo repos.User, notificationrepo repos.Notification, authzsvc *Authorization) *User {
	return &User{
		logger:           logger,
		urepo:            urepo,
		authzsvc:         authzsvc,
		notificationrepo: notificationrepo,
	}
}

// Register registers a user.
func (u *User) Register(ctx context.Context, d db.DBTX, params UserRegisterParams) (*db.User, error) {
	defer newOTELSpan(ctx, "User.Register").End()

	if params.Role == "" {
		params.Role = models.RoleUser
	}
	role := u.authzsvc.RoleByName(params.Role)

	fmt.Printf("role: %v\n", role)

	// append default scopes for role upon registration regardless of provided params
	params.Scopes = append(params.Scopes, u.authzsvc.DefaultScopes(params.Role)...)

	repoParams := db.UserCreateParams{
		FirstName:  params.FirstName,
		LastName:   params.LastName,
		Username:   params.Username,
		Email:      params.Email,
		ExternalID: params.ExternalID,
		RoleRank:   role.Rank,
		Scopes:     params.Scopes,
	}

	user, err := u.urepo.Create(ctx, d, &repoParams)
	if err != nil {
		return nil, fmt.Errorf("urepo.Create: %w", err)
	}

	u.logger.Infof("user %q registered", user.UserID)

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

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, internal.NewErrorf(models.ErrorCodeInvalidUUID, "could not parse UUID")
	}

	user, err := u.urepo.ByID(ctx, d, uid)
	if err != nil {
		return nil, fmt.Errorf("urepo.ByID: %w", err)
	}

	adminRole := u.authzsvc.RoleByName(models.RoleAdmin)

	if user.UserID != caller.UserID &&
		caller.RoleRank < adminRole.Rank {
		return nil, internal.NewErrorf(models.ErrorCodeUnauthorized, "cannot change another user's information")
	}

	up := db.UserUpdateParams{}
	if params.FirstName != nil {
		up.FirstName = pointers.New(params.FirstName)
	}
	if params.LastName != nil {
		up.LastName = pointers.New(params.LastName)
	}

	user, err = u.urepo.Update(ctx, d, uid, &up)
	if err != nil {
		return nil, fmt.Errorf("urepo.Update: %w", err)
	}

	u.logger.Infof("user %q updated", user.UserID)

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

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, internal.NewErrorf(models.ErrorCodeInvalidUUID, "could not parse UUID")
	}

	user, err := u.urepo.ByID(ctx, d, uid)
	if err != nil {
		return nil, fmt.Errorf("urepo.ByID: %w", err)
	}

	adminRole := u.authzsvc.RoleByName(models.RoleAdmin)

	if caller.RoleRank < adminRole.Rank {
		if user.UserID == caller.UserID { // exit early, though it's not possible to update to something not assigned to self already anyway
			return nil, internal.NewErrorf(models.ErrorCodeUnauthorized, "cannot update your own authorization information")
		}
	}

	if params.Scopes != nil {
		for _, s := range *params.Scopes {
			if !slices.Contains(caller.Scopes, s) {
				return nil, internal.NewErrorf(models.ErrorCodeUnauthorized, "cannot set a scope unassigned to self")
			}
		}

		if caller.RoleRank < adminRole.Rank {
			for _, s := range user.Scopes {
				if !slices.Contains(*params.Scopes, s) {
					return nil, internal.NewErrorf(models.ErrorCodeUnauthorized, "cannot unassign a user's scope")
				}
			}
		}
	}

	var rank *int
	if params.Role != nil {
		role := u.authzsvc.RoleByName(*params.Role)
		if role.Rank > caller.RoleRank {
			return nil, internal.NewErrorf(models.ErrorCodeUnauthorized, "cannot set a user rank higher than self")
		}
		if caller.RoleRank < adminRole.Rank {
			if role.Rank < user.RoleRank {
				return nil, internal.NewErrorf(models.ErrorCodeUnauthorized, "cannot demote a user role")
			}
		}
		rank = &role.Rank

		// always reset scopes when changing role
		params.Scopes = pointers.New(u.authzsvc.DefaultScopes(*params.Role))
	}

	user, err = u.urepo.Update(ctx, d, uid, &db.UserUpdateParams{
		Scopes:   params.Scopes,
		RoleRank: rank,
	})
	if err != nil {
		return nil, fmt.Errorf("urepo.Update: %w", err)
	}

	u.logger.Infof("user %q authorization updated", user.UserID)

	return user, nil
}

func (u *User) CreateAPIKey(ctx context.Context, d db.DBTX, user *db.User) (*db.UserAPIKey, error) {
	defer newOTELSpan(ctx, "User.CreateAPIKey").End()

	uak, err := u.urepo.CreateAPIKey(ctx, d, user)
	if err != nil {
		return nil, fmt.Errorf("urepo.CreateAPIKey: %w", err)
	}

	u.logger.Infof("user %q api key created", user.UserID)

	return uak, nil
}

// ByExternalID gets a user by ExternalID.
func (u *User) ByExternalID(ctx context.Context, d db.DBTX, id string) (*db.User, error) {
	defer newOTELSpan(ctx, "User.ByExternalID").End()

	user, err := u.urepo.ByExternalID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("urepo.ByExternalID: %w", err)
	}

	return user, nil
}

// ByEmail gets a user by email.
func (u *User) ByEmail(ctx context.Context, d db.DBTX, email string) (*db.User, error) {
	defer newOTELSpan(ctx, "User.ByEmail").End()

	user, err := u.urepo.ByEmail(ctx, d, email)
	if err != nil {
		return nil, fmt.Errorf("urepo.ByEmail: %w", err)
	}

	return user, nil
}

// ByUsername gets a user by username.
func (u *User) ByUsername(ctx context.Context, d db.DBTX, username string) (*db.User, error) {
	defer newOTELSpan(ctx, "User.ByUsername").End()

	user, err := u.urepo.ByUsername(ctx, d, username)
	if err != nil {
		return nil, fmt.Errorf("urepo.ByUsername: %w", err)
	}

	return user, nil
}

// ByAPIKey gets a user by apiKey.
func (u *User) ByAPIKey(ctx context.Context, d db.DBTX, apiKey string) (*db.User, error) {
	defer newOTELSpan(ctx, "User.ByAPIKey").End()

	user, err := u.urepo.ByAPIKey(ctx, d, apiKey)
	if err != nil {
		return nil, fmt.Errorf("urepo.ByAPIKey: %w", err)
	}

	return user, nil
}

// Delete marks a user as deleted.
func (u *User) Delete(ctx context.Context, d db.DBTX, id uuid.UUID) (*db.User, error) {
	defer newOTELSpan(ctx, "User.Delete").End()

	user, err := u.urepo.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("urepo.Delete: %w", err)
	}

	u.logger.Infof("user %q deleted", user.UserID)

	return user, nil
}

// TODO.
func (u *User) LatestPersonalNotifications(ctx context.Context, d db.DBTX, userID uuid.UUID) ([]db.UserNotification, error) {
	// this will also set user.has_new_personal_notifications to false in the same tx
	return []db.UserNotification{}, nil

	// defer newOTELSpan(ctx, "User.ByAPIKey").End()

	// uid, err := uuid.Parse(userID)
	// if err != nil {
	// 	return nil, internal.NewErrorf(models.ErrorCodeInvalidUUID, "could not parse UUID")
	// }

	// user, err := u.notificationrepo.LatestUserNotifications(ctx, d, db.GetUserNotificationsParams{UserID: uid})
	// if err != nil {
	// 	return nil, fmt.Errorf("urepo.ByAPIKey: %w", err)
	// }

	// return user, nil
}

// TODO.
func (u *User) LatestGlobalNotifications(ctx context.Context, d db.DBTX, userID uuid.UUID) ([]db.GetUserNotificationsRow, error) {
	// this will also set user.has_new_global_notifications to false in the same tx
	return []db.GetUserNotificationsRow{}, nil
}

func (u *User) AssignTeam(ctx context.Context, d db.DBTX, userID uuid.UUID, teamID int) error {
	defer newOTELSpan(ctx, "User.AssignTeam").End()

	_, err := db.CreateUserTeam(ctx, d, &db.UserTeamCreateParams{
		TeamID: teamID,
		Member: userID,
	})
	if err != nil {
		return fmt.Errorf("db.CreateUserTeam: %w", err)
	}

	u.logger.Infof("user %q assigned to team %d", userID, teamID)

	return nil
}
