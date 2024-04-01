package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/slices"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// User represents the repository used for interacting with User records.
type User struct {
	q      db.Querier
	logger *zap.SugaredLogger
}

// NewUser instantiates the User repository.
func NewUser() *User {
	return &User{
		q: NewQuerierWrapper(db.New()),
	}
}

var _ repos.User = (*User)(nil)

func (u *User) Create(ctx context.Context, d db.DBTX, params *db.UserCreateParams) (*db.User, error) {
	params.Scopes = slices.Unique(params.Scopes)
	user, err := db.CreateUser(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not create user: %w", ParseDBErrorDetail(err))
	}

	return user, nil
}

func (u *User) Paginated(ctx context.Context, d db.DBTX, params models.GetPaginatedUsersParams) ([]db.User, error) {
	createdAt, err := time.Parse(time.RFC3339, params.Cursor)
	if err != nil {
		return nil, internal.NewErrorf(models.ErrorCodeInvalidArgument, "invalid createdAt cursor for paginated user: %s", params.Cursor)
	}
	var filters map[string][]interface{}
	if params.SearchQuery.Items != nil {
		filters, err = GenerateDefaultFilters(db.TableEntityUser, *params.SearchQuery.Items)
		if err != nil {
			return nil, internal.NewErrorf(models.ErrorCodeInvalidArgument, "invalid default filters")
		}

		// TODO: sort mapping could also be generated, just like db.EntityFilters
		// we can have db.EntitySorting indexed by entity, field and direction,
		// ignoring nulls first/last.

		// handle custom keys as desired. They should be set in spec directly and
		// not via rest/models.go
	}

	opts := []db.UserSelectConfigOption{
		db.WithUserFilters(filters),
		db.WithUserJoin(db.UserJoins{MemberTeams: true, MemberProjects: true}),
	}
	if params.Limit > 0 { // for users, allow 0 or less to fetch all
		opts = append(opts, db.WithUserLimit(params.Limit))
	}

	users, err := db.UserPaginatedByCreatedAt(ctx, d, createdAt, params.Direction, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get paginated users: %w", ParseDBErrorDetail(err))
	}

	return users, nil
}

func (u *User) Update(ctx context.Context, d db.DBTX, id db.UserID, params *db.UserUpdateParams) (*db.User, error) {
	user, err := u.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get user by id: %w", ParseDBErrorDetail(err))
	}

	if params.Scopes != nil {
		*params.Scopes = slices.Unique(*params.Scopes)
	}

	user.SetUpdateParams(params)

	user, err = user.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update user: %w", ParseDBErrorDetail(err))
	}

	return user, err
}

func (u *User) Delete(ctx context.Context, d db.DBTX, id db.UserID) (*db.User, error) {
	user := &db.User{
		UserID: id,
	}

	if err := user.SoftDelete(ctx, d); err != nil {
		return nil, fmt.Errorf("could not mark user as deleted: %w", ParseDBErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByExternalID(ctx context.Context, d db.DBTX, extID string, opts ...db.UserSelectConfigOption) (*db.User, error) {
	user, err := db.UserByExternalID(ctx, d, extID, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get user by external id: %w", ParseDBErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByEmail(ctx context.Context, d db.DBTX, email string, opts ...db.UserSelectConfigOption) (*db.User, error) {
	user, err := db.UserByEmail(ctx, d, email, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get user by email: %w", ParseDBErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByTeam(ctx context.Context, d db.DBTX, teamID db.TeamID) ([]db.User, error) {
	team, err := db.TeamByTeamID(ctx, d, teamID, db.WithTeamJoin(db.TeamJoins{Members: true}))
	if err != nil {
		return []db.User{}, fmt.Errorf("could not get users by team: %w", ParseDBErrorDetail(err))
	}

	return *team.MembersJoin, nil
}

func (u *User) ByProject(ctx context.Context, d db.DBTX, projectID db.ProjectID) ([]db.User, error) {
	teams, err := db.TeamsByProjectID(ctx, d, projectID)
	if err != nil {
		return []db.User{}, fmt.Errorf("could not get teams in project: %w", ParseDBErrorDetail(err))
	}

	var users []db.User
	for _, t := range teams {
		uu, err := u.ByTeam(ctx, d, t.TeamID)
		if err != nil {
			return []db.User{}, fmt.Errorf("u.ByTeam: %w", ParseDBErrorDetail(err))
		}
		users = append(users, uu...)
	}

	return users, nil
}

func (u *User) ByUsername(ctx context.Context, d db.DBTX, username string, opts ...db.UserSelectConfigOption) (*db.User, error) {
	user, err := db.UserByUsername(ctx, d, username, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get user by username: %w", ParseDBErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByID(ctx context.Context, d db.DBTX, id db.UserID, opts ...db.UserSelectConfigOption) (*db.User, error) {
	user, err := db.UserByUserID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get user: %w", ParseDBErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByAPIKey(ctx context.Context, d db.DBTX, apiKey string) (*db.User, error) {
	uak, err := db.UserAPIKeyByAPIKey(ctx, d, apiKey, db.WithUserAPIKeyJoin(db.UserAPIKeyJoins{User: true}))
	if err != nil {
		return nil, fmt.Errorf("could not get api key: %w", ParseDBErrorDetail(err))
	}

	if uak.UserJoin == nil {
		return nil, fmt.Errorf("could not join user by api key")
	}

	return uak.UserJoin, nil
}

func (u *User) DeleteAPIKey(ctx context.Context, d db.DBTX, apiKey string) (*db.UserAPIKey, error) {
	uak, err := db.UserAPIKeyByAPIKey(ctx, d, apiKey)
	if err != nil {
		return nil, fmt.Errorf("could not get api key: %w", ParseDBErrorDetail(err))
	}

	err = uak.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete api key: %w", ParseDBErrorDetail(err))
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
		return nil, fmt.Errorf("could not save api key: %w", ParseDBErrorDetail(err))
	}

	user.APIKeyID = pointers.New(uak.UserAPIKeyID)
	if _, err := user.Update(ctx, d); err != nil {
		return nil, fmt.Errorf("could not update user: %w", ParseDBErrorDetail(err))
	}

	return uak, nil
}
