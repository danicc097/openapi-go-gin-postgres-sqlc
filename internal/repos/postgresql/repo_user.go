package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/slices"
	"github.com/google/uuid"
)

// User represents the repository used for interacting with User records.
type User struct {
	q *db.Queries
}

// NewUser instantiates the User repository.
func NewUser() *User {
	return &User{
		q: db.New(),
	}
}

var _ repos.User = (*User)(nil)

func (u *User) Create(ctx context.Context, d db.DBTX, params *db.UserCreateParams) (*db.User, error) {
	params.Scopes = slices.Unique(params.Scopes)
	user, err := db.CreateUser(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not create user: %w", parseErrorDetail(err))
	}

	return user, nil
}

func (u *User) Update(ctx context.Context, d db.DBTX, id uuid.UUID, params *db.UserUpdateParams) (*db.User, error) {
	user, err := u.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get user by id: %w", parseErrorDetail(err))
	}

	if params.Scopes != nil {
		*params.Scopes = slices.Unique(*params.Scopes)
	}

	user.SetUpdateParams(params)

	user, err = user.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update user: %w", parseErrorDetail(err))
	}

	return user, err
}

func (u *User) Delete(ctx context.Context, d db.DBTX, id uuid.UUID) (*db.User, error) {
	user, err := u.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get user by id %w", parseErrorDetail(err))
	}

	if err := user.SoftDelete(ctx, d); err != nil {
		return nil, fmt.Errorf("could not mark user as deleted: %w", parseErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByExternalID(ctx context.Context, d db.DBTX, extID string, opts ...db.UserSelectConfigOption) (*db.User, error) {
	user, err := db.UserByExternalID(ctx, d, extID, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get user by external id: %w", parseErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByEmail(ctx context.Context, d db.DBTX, email string, opts ...db.UserSelectConfigOption) (*db.User, error) {
	user, err := db.UserByEmail(ctx, d, email, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get user by email: %w", parseErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByTeam(ctx context.Context, d db.DBTX, teamID int) ([]db.User, error) {
	team, err := db.TeamByTeamID(ctx, d, teamID, db.WithTeamJoin(db.TeamJoins{Members: true}))
	if err != nil {
		return []db.User{}, fmt.Errorf("could not get users by team: %w", parseErrorDetail(err))
	}

	return *team.TeamMembersJoin, nil
}

func (u *User) ByProject(ctx context.Context, d db.DBTX, projectID int) ([]db.User, error) {
	teams, err := db.TeamsByProjectID(ctx, d, projectID)
	if err != nil {
		return []db.User{}, fmt.Errorf("could not get teams in project: %w", parseErrorDetail(err))
	}

	var users []db.User
	for _, t := range teams {
		uu, err := u.ByTeam(ctx, d, t.TeamID)
		if err != nil {
			return []db.User{}, fmt.Errorf("u.ByTeam: %w", parseErrorDetail(err))
		}
		users = append(users, uu...)
	}

	return users, nil
}

func (u *User) ByUsername(ctx context.Context, d db.DBTX, username string, opts ...db.UserSelectConfigOption) (*db.User, error) {
	user, err := db.UserByUsername(ctx, d, username, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get user by username: %w", parseErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByID(ctx context.Context, d db.DBTX, id uuid.UUID, opts ...db.UserSelectConfigOption) (*db.User, error) {
	user, err := db.UserByUserID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get user: %w", parseErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByAPIKey(ctx context.Context, d db.DBTX, apiKey string) (*db.User, error) {
	uak, err := db.UserAPIKeyByAPIKey(ctx, d, apiKey, db.WithUserAPIKeyJoin(db.UserAPIKeyJoins{User: true}))
	if err != nil {
		return nil, fmt.Errorf("could not get api key: %w", parseErrorDetail(err))
	}

	if uak.UserJoin == nil {
		return nil, fmt.Errorf("could not join user by api key")
	}

	return uak.UserJoin, nil
}

func (u *User) DeleteAPIKey(ctx context.Context, d db.DBTX, apiKey string) (*db.UserAPIKey, error) {
	uak, err := db.UserAPIKeyByAPIKey(ctx, d, apiKey)
	if err != nil {
		return nil, fmt.Errorf("could not get api key: %w", parseErrorDetail(err))
	}

	err = uak.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete api key: %w", parseErrorDetail(err))
	}

	return uak, nil
}

func (u *User) CreateAPIKey(ctx context.Context, d db.DBTX, user *db.User) (*db.UserAPIKey, error) {
	uak := &db.UserAPIKey{
		APIKey:    uuid.NewString(),
		ExpiresOn: time.Now().AddDate(1, 0, 0),
		UserID:    user.UserID,
	}
	if _, err := uak.Insert(ctx, d); err != nil {
		return nil, fmt.Errorf("could not save api key: %w", parseErrorDetail(err))
	}

	user.APIKeyID = pointers.New(uak.UserAPIKeyID)
	if _, err := user.Update(ctx, d); err != nil {
		return nil, fmt.Errorf("could not update user: %w", parseErrorDetail(err))
	}

	return uak, nil
}