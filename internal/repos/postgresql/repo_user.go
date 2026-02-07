package postgresql

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/slices"
)

// User represents the repository used for interacting with User records.
type User struct {
	q      models.Querier
	logger *zap.SugaredLogger
}

// NewUser instantiates the User repository.
func NewUser() *User {
	return &User{
		q: NewQuerierWrapper(models.New()),
	}
}

var _ repos.User = (*User)(nil)

func (u *User) Create(ctx context.Context, d models.DBTX, params *models.UserCreateParams) (*models.User, error) {
	params.Scopes = slices.Unique(params.Scopes)
	user, err := models.CreateUser(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not create user: %w", ParseDBErrorDetail(err))
	}

	return user, nil
}

func (u *User) Paginated(ctx context.Context, d models.DBTX, params repos.GetPaginatedUsersParams) ([]models.User, error) {
	var err error
	filters := make(map[string][]interface{})
	if ii := params.Items; ii != nil {
		filters, err = GenerateDefaultFilters(models.TableEntityUser, *ii)
		if err != nil {
			return nil, internal.WrapErrorf(err, models.ErrorCodeInvalidArgument, "invalid default filters")
		}
	}

	// handle custom keys as desired. They should be set in spec directly and
	// not via rest/models.spec.go
	if r := params.RoleRank; r != nil {
		filters["role_rank = $i"] = []interface{}{r}
	}

	opts := []models.UserSelectConfigOption{
		models.WithUserFilters(filters),
		models.WithUserJoin(models.UserJoins{MemberTeams: true, MemberProjects: true}),
	}
	if params.Limit > 0 { // for users, allow 0 or less to fetch all
		opts = append(opts, models.WithUserLimit(params.Limit))
	}

	if err := setDefaultCursor(d, models.TableEntityUser, &params.Cursor); err != nil {
		return nil, fmt.Errorf("could not set default cursors: %w", ParseDBErrorDetail(err))
	}

	users, err := models.UserPaginated(ctx, d, params.Cursor, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get paginated users: %w", ParseDBErrorDetail(err))
	}

	return users, nil
}

func (u *User) Update(ctx context.Context, d models.DBTX, id models.UserID, params *models.UserUpdateParams) (*models.User, error) {
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

func (u *User) Delete(ctx context.Context, d models.DBTX, id models.UserID) (*models.User, error) {
	user := &models.User{
		UserID: id,
	}

	if err := user.SoftDelete(ctx, d); err != nil {
		return nil, fmt.Errorf("could not mark user as deleted: %w", ParseDBErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByExternalID(ctx context.Context, d models.DBTX, extID string, opts ...models.UserSelectConfigOption) (*models.User, error) {
	user, err := models.UserByExternalID(ctx, d, extID, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get user by external id: %w", ParseDBErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByEmail(ctx context.Context, d models.DBTX, email string, opts ...models.UserSelectConfigOption) (*models.User, error) {
	user, err := models.UserByEmail(ctx, d, email, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get user by email: %w", ParseDBErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByTeam(ctx context.Context, d models.DBTX, teamID models.TeamID) ([]models.User, error) {
	team, err := models.TeamByTeamID(ctx, d, teamID, models.WithTeamJoin(models.TeamJoins{Members: true}))
	if err != nil {
		return []models.User{}, fmt.Errorf("could not get users by team: %w", ParseDBErrorDetail(err))
	}

	return *team.MembersJoin, nil
}

func (u *User) ByProject(ctx context.Context, d models.DBTX, projectID models.ProjectID) ([]models.User, error) {
	teams, err := models.TeamsByProjectID(ctx, d, projectID)
	if err != nil {
		return []models.User{}, fmt.Errorf("could not get teams in project: %w", ParseDBErrorDetail(err))
	}

	var users []models.User
	for _, t := range teams {
		uu, err := u.ByTeam(ctx, d, t.TeamID)
		if err != nil {
			return []models.User{}, fmt.Errorf("u.ByTeam: %w", ParseDBErrorDetail(err))
		}
		users = append(users, uu...)
	}

	return users, nil
}

func (u *User) ByUsername(ctx context.Context, d models.DBTX, username string, opts ...models.UserSelectConfigOption) (*models.User, error) {
	user, err := models.UserByUsername(ctx, d, username, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get user by username: %w", ParseDBErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByID(ctx context.Context, d models.DBTX, id models.UserID, opts ...models.UserSelectConfigOption) (*models.User, error) {
	user, err := models.UserByUserID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get user: %w", ParseDBErrorDetail(err))
	}

	return user, nil
}

func (u *User) ByAPIKey(ctx context.Context, d models.DBTX, apiKey string) (*models.User, error) {
	uak, err := models.UserAPIKeyByAPIKey(ctx, d, apiKey, models.WithUserAPIKeyJoin(models.UserAPIKeyJoins{User: true}))
	if err != nil {
		return nil, fmt.Errorf("could not get api key: %w", ParseDBErrorDetail(err))
	}

	if uak.UserJoin == nil {
		return nil, errors.New("could not join user by api key")
	}

	return uak.UserJoin, nil
}

func (u *User) DeleteAPIKey(ctx context.Context, d models.DBTX, apiKey string) (*models.UserAPIKey, error) {
	uak, err := models.UserAPIKeyByAPIKey(ctx, d, apiKey)
	if err != nil {
		return nil, fmt.Errorf("could not get api key: %w", ParseDBErrorDetail(err))
	}

	err = uak.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete api key: %w", ParseDBErrorDetail(err))
	}

	return uak, nil
}

func (u *User) CreateAPIKey(ctx context.Context, d models.DBTX, user *models.User) (*models.UserAPIKey, error) {
	uak := &models.UserAPIKey{
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
