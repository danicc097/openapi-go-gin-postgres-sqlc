package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

type User struct {
	logger   *zap.SugaredLogger
	repos    *repos.Repos
	authzsvc *Authorization
	// sharedDBOpts represents shared db select options for all work item entities
	// for returned values
	getSharedDBOpts func() []models.UserSelectConfigOption
}

// NOTE: the most important distinction about repositories is that they represent collections of entities. They do not represent database storage or caching or any number of technical concerns. Repositories represent collections. How you hold those collections is simply an implementation detail.
// Repo should not be aware of models Role and Scope, its conversion or its default values. That's all
// for upper layers convenience. e.g roles: entity uses rank internally. Repo should not care about mappings to user-friendly names.
type UserRegisterParams struct {
	Username   string         `json:"username"`
	Email      string         `json:"email"`
	FirstName  *string        `json:"firstName"`
	LastName   *string        `json:"lastName"`
	ExternalID string         `json:"externalID"`
	Scopes     []models.Scope `json:"scopes"`
	Role       models.Role    `json:"role"`
}

// NewUser returns a new User service.
func NewUser(logger *zap.SugaredLogger, repos *repos.Repos) *User {
	authzsvc := NewAuthorization(logger)

	return &User{
		logger:   logger,
		repos:    repos,
		authzsvc: authzsvc,
		getSharedDBOpts: func() []models.UserSelectConfigOption {
			return []models.UserSelectConfigOption{models.WithUserJoin(models.UserJoins{MemberProjects: true, MemberTeams: true})}
		},
	}
}

// Register registers a user.
// IMPORTANT: for internal use only.
func (u *User) Register(ctx context.Context, d models.DBTX, params UserRegisterParams) (*models.User, error) {
	defer newOTelSpan().Build(ctx).End()

	if params.Role == "" {
		params.Role = models.RoleUser
	}
	role := u.authzsvc.RoleByName(params.Role)

	// append default scopes for role upon registration regardless of provided params
	params.Scopes = append(params.Scopes, u.authzsvc.DefaultScopes(params.Role)...)

	repoParams := models.UserCreateParams{
		FirstName:                params.FirstName,
		LastName:                 params.LastName,
		Username:                 params.Username,
		Email:                    params.Email,
		ExternalID:               params.ExternalID,
		RoleRank:                 role.Rank,
		Scopes:                   params.Scopes,
		APIKeyID:                 nil,
		HasGlobalNotifications:   false,
		HasPersonalNotifications: false,
		Age:                      nil,
	}

	user, err := u.repos.User.Create(ctx, d, &repoParams)
	if err != nil {
		return nil, fmt.Errorf("repos.User.Create: %w", err)
	}

	u.logger.Infof("user %q registered", user.UserID)

	// TODO: publish internal event --> consumer send mail, send teams message, etc.
	// so we dont block here, make it easier to test, much cleaner (passing lots of unrelated services to constructor) and decoupled
	// see watermill lib for event-driven.
	// we want persistence for these, as well as retries (notifications).

	return user, nil
}

func (u *User) ByID(ctx context.Context, d models.DBTX, id models.UserID, dbOpts ...models.UserSelectConfigOption) (*models.User, error) {
	opts := append(u.getSharedDBOpts(), dbOpts...)
	user, err := u.repos.User.ByID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("repos.User.ByID: %w", err)
	}

	return user, nil
}

// Update updates a user. In this case request params are defined in the spec
// and converted to db params for demo purposes.
// The same can be achieved excluding fields from db update params via SQL column comments.
func (u *User) Update(ctx context.Context, d models.DBTX, id models.UserID, caller CtxUser, params *models.UpdateUserRequest) (*models.User, error) {
	defer newOTelSpan().Build(ctx).End()

	if params == nil {
		return nil, errors.New("params cannot be nil")
	}

	user, err := u.repos.User.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.User.ByID: %w", err)
	}

	adminRole := u.authzsvc.RoleByName(models.RoleAdmin)

	if user.UserID != caller.UserID &&
		caller.RoleRank < adminRole.Rank {
		return nil, internal.NewErrorf(models.ErrorCodeUnauthorized, "cannot change another user's information")
	}

	up := models.UserUpdateParams{}
	if params.FirstName != nil {
		up.FirstName = pointers.New(params.FirstName)
	}
	if params.LastName != nil {
		up.LastName = pointers.New(params.LastName)
	}

	user, err = u.repos.User.Update(ctx, d, id, &up)
	if err != nil {
		return nil, fmt.Errorf("repos.User.Update: %w", err)
	}

	u.logger.Infof("user %q updated", user.UserID)

	return user, nil
}

func (u *User) UpdateUserAuthorization(ctx context.Context, d models.DBTX, id models.UserID, caller CtxUser, params *models.UpdateUserAuthRequest) (*models.User, error) {
	defer newOTelSpan().Build(ctx).End()

	if params == nil {
		return nil, errors.New("params cannot be nil")
	}

	user, err := u.repos.User.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.User.ByID: %w", err)
	}

	adminRole := u.authzsvc.RoleByName(models.RoleAdmin)

	if caller.RoleRank < adminRole.Rank {
		if user.UserID == caller.UserID { // exit early, though it's not possible to update to something not assigned to self already anyway
			return nil, internal.NewErrorf(models.ErrorCodeUnauthorized, "cannot update your own authorization information")
		}
	}

	if params.Scopes != nil {
		for _, s := range *params.Scopes {
			if !slices.Contains(caller.Scopes, s) && caller.RoleRank < adminRole.Rank {
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

	user, err = u.repos.User.Update(ctx, d, id, &models.UserUpdateParams{
		Scopes:   params.Scopes,
		RoleRank: rank,
	})
	if err != nil {
		return nil, fmt.Errorf("repos.User.Update: %w", err)
	}

	u.logger.Infof("user %q authorization updated", user.UserID)

	return user, nil
}

func (u *User) CreateAPIKey(ctx context.Context, d models.DBTX, user *models.User) (*models.UserAPIKey, error) {
	defer newOTelSpan().Build(ctx).End()

	uak, err := u.repos.User.CreateAPIKey(ctx, d, user)
	if err != nil {
		return nil, fmt.Errorf("repos.User.CreateAPIKey: %w", err)
	}

	u.logger.Infof("user %q api key created", user.UserID)

	return uak, nil
}

// ByExternalID gets a user by ExternalID.
func (u *User) ByExternalID(ctx context.Context, d models.DBTX, id string, dbOpts ...models.UserSelectConfigOption) (*models.User, error) {
	defer newOTelSpan().Build(ctx).End()

	opts := append(u.getSharedDBOpts(), dbOpts...)
	user, err := u.repos.User.ByExternalID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("repos.User.ByExternalID: %w", err)
	}

	return user, nil
}

func (u *User) Paginated(ctx context.Context, d models.DBTX, params models.GetPaginatedUsersParams) ([]models.User, error) {
	defer newOTelSpan().Build(ctx).End()

	var roleRank *int
	if r := params.SearchQuery.Role; r != nil {
		role := u.authzsvc.RoleByName(*r)
		roleRank = &role.Rank
	}
	var c interface{}
	if params.Cursor != nil {
		c = *params.Cursor
	}

	users, err := u.repos.User.Paginated(ctx, d, repos.GetPaginatedUsersParams{
		Limit:     params.Limit,
		Direction: params.Direction,
		Cursor: models.PaginationCursor{
			Column:    params.Column,
			Direction: params.Direction,
			Value:     &c,
		},
		Items:    params.SearchQuery.Items,
		RoleRank: roleRank,
	})
	if err != nil {
		return nil, fmt.Errorf("repos.User.Paginated: %w", err)
	}

	return users, nil
}

// ByEmail gets a user by email.
func (u *User) ByEmail(ctx context.Context, d models.DBTX, email string, dbOpts ...models.UserSelectConfigOption) (*models.User, error) {
	defer newOTelSpan().Build(ctx).End()

	opts := append(u.getSharedDBOpts(), dbOpts...)
	user, err := u.repos.User.ByEmail(ctx, d, email, opts...)
	if err != nil {
		return nil, fmt.Errorf("repos.User.ByEmail: %w", err)
	}

	return user, nil
}

// ByUsername gets a user by username.
func (u *User) ByUsername(ctx context.Context, d models.DBTX, username string, dbOpts ...models.UserSelectConfigOption) (*models.User, error) {
	defer newOTelSpan().Build(ctx).End()

	opts := append(u.getSharedDBOpts(), dbOpts...)
	user, err := u.repos.User.ByUsername(ctx, d, username, opts...)
	if err != nil {
		return nil, fmt.Errorf("repos.User.ByUsername: %w", err)
	}

	return user, nil
}

// ByAPIKey gets a user by apiKey.
func (u *User) ByAPIKey(ctx context.Context, d models.DBTX, apiKey string) (*models.User, error) {
	defer newOTelSpan().Build(ctx).End()

	user, err := u.repos.User.ByAPIKey(ctx, d, apiKey)
	if err != nil {
		return nil, fmt.Errorf("repos.User.ByAPIKey: %w", err)
	}

	opts := append(u.getSharedDBOpts(), models.WithUserJoin(models.UserJoins{UserAPIKey: true}))
	user, err = u.repos.User.ByID(ctx, d, user.UserID, opts...)
	if err != nil {
		return nil, fmt.Errorf("repos.User.ByID: %w", err)
	}

	return user, nil
}

// Delete marks a user as deleted.
func (u *User) Delete(ctx context.Context, d models.DBTX, id models.UserID) (*models.User, error) {
	defer newOTelSpan().Build(ctx).End()

	user, err := u.repos.User.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.User.Delete: %w", err)
	}

	u.logger.Infof("user %q deleted", user.UserID)

	return user, nil
}

// TODO.
func (u *User) LatestPersonalNotifications(ctx context.Context, d models.DBTX, userID models.UserID) ([]models.UserNotification, error) {
	// this will also set user.has_new_personal_notifications to false in the same tx
	return []models.UserNotification{}, nil

	// defer newOTelSpan().Build(ctx).End()

	// uid, err := uuid.Parse(userID)
	// if err != nil {
	// 	return nil, internal.NewErrorf(models.ErrorCodeInvalidUUID, "could not parse UUID")
	// }

	// user, err := u.notificationrepo.LatestNotifications(ctx, d, db.GetUserNotificationsParams{UserID: uid})
	// if err != nil {
	// 	return nil, fmt.Errorf("repos.User.ByAPIKey: %w", err)
	// }

	// return user, nil
}

// TODO.
func (u *User) LatestGlobalNotifications(ctx context.Context, d models.DBTX, userID models.UserID) ([]models.GetUserNotificationsRow, error) {
	// this will also set user.has_new_global_notifications to false in the same tx
	return []models.GetUserNotificationsRow{}, nil
}

func (u *User) AssignTeam(ctx context.Context, d models.DBTX, userID models.UserID, teamID models.TeamID) (*models.User, error) {
	defer newOTelSpan().Build(ctx).End()

	_, err := models.CreateUserTeam(ctx, d, &models.UserTeamCreateParams{
		TeamID: teamID,
		Member: userID,
	})
	var ierr *internal.Error
	if err != nil {
		err := postgresql.ParseDBErrorDetail(err)
		if !(errors.As(err, &ierr) && ierr.Code() == models.ErrorCodeAlreadyExists) {
			return nil, internal.WrapErrorf(err, models.ErrorCodeUnknown, "could not assign user to team")
		}
	}

	user, err := u.repos.User.ByID(ctx, d, userID, u.getSharedDBOpts()...)
	if err != nil {
		return nil, fmt.Errorf("repos.User.ByID: %w", err)
	}

	return user, nil
}
