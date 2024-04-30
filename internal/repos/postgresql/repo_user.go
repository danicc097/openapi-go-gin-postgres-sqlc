package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	models1 "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/slices"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// User represents the repository used for interacting with User records.
type User struct {
	q      models1.Querier
	logger *zap.SugaredLogger
}

// NewUser instantiates the User repository.
func NewUser() *User {
	return &User{
		q: NewQuerierWrapper(models1.New()),
	}
}

var _ repos.User = (*User)(nil)

func (u *User) Create(ctx context.Context, d models1.DBTX, params *models1.UserCreateParams) (*models1.User, error) {
	params.Scopes = slices.Unique(params.Scopes)
	user, err := models1.CreateUser(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not create user: %w", ParseDBErrorDetail(err))
	}

	return user, nil
}

func (u *User) Paginated(ctx context.Context, d models1.DBTX, params repos.GetPaginatedUsersParams) ([]models1.User, error) {
	var err error
	filters := make(map[string][]interface{})
	if ii := params.Items; ii != nil {
		filters, err = GenerateDefaultFilters(models1.TableEntityUser, *ii)
		if err != nil {
			return nil, internal.WrapErrorf(err, models.ErrorCodeInvalidArgument, "invalid default filters")
		}
	}

	// handle custom keys as desired. They should be set in spec directly and
	// not via rest/models.go
	if r := params.RoleRank; r != nil {
		filters["role_rank = $i"] = []interface{}{r}
	}

	opts := []models1.UserSelectConfigOption{
		models1.WithUserFilters(filters),
		models1.WithUserJoin(models1.UserJoins{MemberTeams: true, MemberProjects: true}),
	}
	if params.Limit > 0 { // for users, allow 0 or less to fetch all
		opts = append(opts, models1.WithUserLimit(params.Limit))
	}

	if err := setDefaultCursor(d, models1.TableEntityUser, &params.Cursor); err != nil {
		return nil, fmt.Errorf("could not set default cursors: %w", ParseDBErrorDetail(err))
	}

	users, err := models1.UserPaginated(ctx, d, params.Cursor, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get paginated users: %w", ParseDBErrorDetail(err))
	}

	return users, nil
}

func (u *User) Update(ctx context.Context, d models1.DBTX, id models1.UserID, params *models1.UserUpdateParams) (*models1.User, error) {
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

func (u *User) Delete(ctx context.Context, d models1.DBTX, id models1.UserID) (*models1.User, error) {
	user := &models1.User{
		UserID: id,
	}

	if err := user.SoftDelete(ctx, d); err != nil {
		return nil, fmt.Errorf("could not mark user as deleted: %w", ParseDBErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByExternalID(ctx context.Context, d models1.DBTX, extID string, opts ...models1.UserSelectConfigOption) (*models1.User, error) {
	user, err := models1.UserByExternalID(ctx, d, extID, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get user by external id: %w", ParseDBErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByEmail(ctx context.Context, d models1.DBTX, email string, opts ...models1.UserSelectConfigOption) (*models1.User, error) {
	user, err := models1.UserByEmail(ctx, d, email, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get user by email: %w", ParseDBErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByTeam(ctx context.Context, d models1.DBTX, teamID models1.TeamID) ([]models1.User, error) {
	team, err := models1.TeamByTeamID(ctx, d, teamID, models1.WithTeamJoin(models1.TeamJoins{Members: true}))
	if err != nil {
		return []models1.User{}, fmt.Errorf("could not get users by team: %w", ParseDBErrorDetail(err))
	}

	return *team.MembersJoin, nil
}

func (u *User) ByProject(ctx context.Context, d models1.DBTX, projectID models1.ProjectID) ([]models1.User, error) {
	teams, err := models1.TeamsByProjectID(ctx, d, projectID)
	if err != nil {
		return []models1.User{}, fmt.Errorf("could not get teams in project: %w", ParseDBErrorDetail(err))
	}

	var users []models1.User
	for _, t := range teams {
		uu, err := u.ByTeam(ctx, d, t.TeamID)
		if err != nil {
			return []models1.User{}, fmt.Errorf("u.ByTeam: %w", ParseDBErrorDetail(err))
		}
		users = append(users, uu...)
	}

	return users, nil
}

func (u *User) ByUsername(ctx context.Context, d models1.DBTX, username string, opts ...models1.UserSelectConfigOption) (*models1.User, error) {
	user, err := models1.UserByUsername(ctx, d, username, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get user by username: %w", ParseDBErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByID(ctx context.Context, d models1.DBTX, id models1.UserID, opts ...models1.UserSelectConfigOption) (*models1.User, error) {
	user, err := models1.UserByUserID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get user: %w", ParseDBErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByAPIKey(ctx context.Context, d models1.DBTX, apiKey string) (*models1.User, error) {
	uak, err := models1.UserAPIKeyByAPIKey(ctx, d, apiKey, models1.WithUserAPIKeyJoin(models1.UserAPIKeyJoins{User: true}))
	if err != nil {
		return nil, fmt.Errorf("could not get api key: %w", ParseDBErrorDetail(err))
	}

	if uak.UserJoin == nil {
		return nil, fmt.Errorf("could not join user by api key")
	}

	return uak.UserJoin, nil
}

func (u *User) DeleteAPIKey(ctx context.Context, d models1.DBTX, apiKey string) (*models1.UserAPIKey, error) {
	uak, err := models1.UserAPIKeyByAPIKey(ctx, d, apiKey)
	if err != nil {
		return nil, fmt.Errorf("could not get api key: %w", ParseDBErrorDetail(err))
	}

	err = uak.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete api key: %w", ParseDBErrorDetail(err))
	}

	return uak, nil
}

func (u *User) CreateAPIKey(ctx context.Context, d models1.DBTX, user *models1.User) (*models1.UserAPIKey, error) {
	uak := &models1.UserAPIKey{
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
