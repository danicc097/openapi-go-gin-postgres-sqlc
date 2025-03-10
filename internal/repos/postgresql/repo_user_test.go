package postgresql_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqlrandom"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/reposwrappers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUser_Update(t *testing.T) {
	t.Parallel()

	user := newRandomUser(t, testPool)

	type args struct {
		id     models.UserID
		params models.UserUpdateParams
	}
	type params struct {
		name        string
		args        args
		want        *models.User
		errContains string
	}
	tests := []params{
		{
			name: "updated",
			args: args{
				id: user.UserID,
				params: models.UserUpdateParams{
					RoleRank: pointers.New(10),
					Scopes:   &models.Scopes{"test", "test", "test"},
				},
			},
			want: func() *models.User {
				u := *user
				u.RoleRank = 10
				u.Scopes = models.Scopes{"test"}

				return &u
			}(),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			u := postgresql.NewUser()
			got, err := u.Update(context.Background(), testPool, tc.args.id, &tc.args.params)
			if err != nil && tc.errContains == "" {
				t.Errorf("unexpected error: %v", err)

				return
			}
			if tc.errContains != "" {
				if err == nil {
					t.Errorf("expected error but got nothing")

					return
				}
				require.ErrorContains(t, err, tc.errContains)

				return
			}

			got.UpdatedAt = user.UpdatedAt // ignore

			// NOTE: this should not fail when running notification tests (from this package) in transaction
			// // since we run tests in parallel, notification fan out effects changes on all users
			got.HasGlobalNotifications = user.HasGlobalNotifications     // ignore
			got.HasPersonalNotifications = user.HasPersonalNotifications // ignore

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestUser_SoftDelete(t *testing.T) {
	t.Parallel()

	user := newRandomUser(t, testPool)

	type args struct {
		id models.UserID
	}
	type params struct {
		name        string
		args        args
		errContains string
	}
	tests := []params{
		{
			name: "deleted",
			args: args{
				id: user.UserID,
			},
			errContains: errNoRows,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			userRepo := postgresql.NewUser()
			_, err := userRepo.Delete(context.Background(), testPool, tc.args.id)
			require.NoError(t, err)

			_, err = userRepo.ByID(context.Background(), testPool, tc.args.id)
			if err != nil && tc.errContains == "" {
				t.Errorf("unexpected error: %v", err)

				return
			}
			if tc.errContains != "" {
				if err == nil {
					t.Errorf("expected error but got nothing")

					return
				}
				require.ErrorContains(t, err, tc.errContains)

				return
			}
			require.ErrorContains(t, err, tc.errContains)
		})
	}
}

func TestUser_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	logger := testutil.NewLogger(t)
	userRepo := reposwrappers.NewUserWithRetry(postgresql.NewUser(), logger, 5, 65*time.Millisecond)

	teamRepo := postgresql.NewTeam()
	projectRepo := postgresql.NewProject()

	project, err := projectRepo.ByName(ctx, testPool, models.ProjectNameDemo)
	require.NoError(t, err)

	team := newRandomTeam(t, testPool, project.ProjectID)
	user := newRandomUser(t, testPool)

	_, err = models.CreateUserTeam(ctx, testPool, &models.UserTeamCreateParams{Member: user.UserID, TeamID: team.TeamID})
	require.NoError(t, err)

	team, err = teamRepo.ByID(ctx, testPool, team.TeamID, models.WithTeamJoin(models.TeamJoins{Members: true, Project: true}))
	require.NoError(t, err)

	uniqueCallback := func(t *testing.T, res *models.User) {
		assert.Equal(t, res.UserID, user.UserID)
	}

	uniqueTestCases := []filterTestCase[*models.User]{
		{
			name:       "external_id",
			filter:     user.ExternalID,
			repoMethod: reflect.ValueOf(userRepo.ByExternalID),
			callback:   uniqueCallback,
		},
		{
			name:       "email",
			filter:     user.Email,
			repoMethod: reflect.ValueOf(userRepo.ByEmail),
			callback:   uniqueCallback,
		},
		{
			name:       "username",
			filter:     user.Username,
			repoMethod: reflect.ValueOf(userRepo.ByUsername),
			callback:   uniqueCallback,
		},
		{
			name:       "user_id",
			filter:     user.UserID,
			repoMethod: reflect.ValueOf(userRepo.ByID),
			callback:   uniqueCallback,
		},
	}
	for _, tc := range uniqueTestCases {
		runGenericFilterTests(t, tc)
	}

	nonUniqueTestCases := []filterTestCase[[]models.User]{
		{
			name:       "team_id",
			filter:     team.TeamID,
			repoMethod: reflect.ValueOf(userRepo.ByTeam),
			callback: func(t *testing.T, res []models.User) {
				assert.Len(t, res, 1)
				assert.Equal(t, res[0].UserID, user.UserID)
				assert.Equal(t, (*team.MembersJoin)[0].UserID, user.UserID)
				assert.Equal(t, team.ProjectID, project.ProjectID)
			},
		},
		{
			name:       "project_id",
			filter:     project.ProjectID,
			repoMethod: reflect.ValueOf(userRepo.ByProject),
			callback: func(t *testing.T, res []models.User) {
				assert.GreaterOrEqual(t, len(res), 1)
				found := false
				// projects and some related entities are hardcoded. Repos could have RandomProjectCreate regardless
				// but better test the same way as services...
				for _, u := range res {
					if u.UserID == user.UserID {
						found = true
					}
				}
				assert.True(t, found)
			},
		},
	}
	for _, tc := range nonUniqueTestCases {
		runGenericFilterTests(t, tc)
	}
}

func TestUser_UserAPIKeys(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)
	userRepo := reposwrappers.NewUserWithRetry(postgresql.NewUser(), logger, 5, 65*time.Millisecond)

	t.Run("correct_api_key_creation", func(t *testing.T) {
		t.Parallel()

		user := newRandomUser(t, testPool)

		uak, err := userRepo.CreateAPIKey(context.Background(), testPool, user)
		require.NoError(t, err)
		assert.NotEmpty(t, uak.APIKey)
		assert.Equal(t, uak.UserID, user.UserID)
		assert.Equal(t, uak.UserAPIKeyID, *user.APIKeyID)
	})

	t.Run("no_api_key_created_when_user_does_not_exist", func(t *testing.T) {
		t.Parallel()

		errContains := "could not save api key"

		_, err := userRepo.CreateAPIKey(context.Background(), testPool, &models.User{UserID: models.NewUserID(uuid.New())})

		require.ErrorContains(t, err, errContains)
	})

	t.Run("can_get_user_by_api_key", func(t *testing.T) {
		t.Parallel()

		newUser := newRandomUser(t, testPool)

		uak, err := userRepo.CreateAPIKey(context.Background(), testPool, newUser)
		require.NoError(t, err)

		user, err := userRepo.ByAPIKey(context.Background(), testPool, uak.APIKey)
		require.NoError(t, err)

		assert.Equal(t, user.UserID, newUser.UserID)
		assert.Equal(t, *user.APIKeyID, uak.UserAPIKeyID)
	})

	t.Run("cannot_get_user_by_api_key_if_key_does_not_exist", func(t *testing.T) {
		t.Parallel()

		errContains := errNoRows

		_, err := userRepo.ByAPIKey(context.Background(), testPool, "missing")
		require.ErrorContains(t, err, errContains)
	})

	t.Run("can_delete_an_api_key", func(t *testing.T) {
		t.Parallel()

		newUser := newRandomUser(t, testPool)

		uak, err := userRepo.CreateAPIKey(context.Background(), testPool, newUser)
		require.NoError(t, err)

		deletedUak, err := userRepo.DeleteAPIKey(context.Background(), testPool, uak.APIKey)
		require.NoError(t, err)
		assert.Equal(t, deletedUak.APIKey, uak.APIKey)

		_, err = userRepo.ByAPIKey(context.Background(), testPool, uak.APIKey)
		require.ErrorContains(t, err, errNoRows)
	})
}

func TestUser_Create(t *testing.T) {
	t.Parallel()

	logger := testutil.NewLogger(t)
	userRepo := reposwrappers.NewUserWithRetry(postgresql.NewUser(), logger, 5, 65*time.Millisecond)

	type want struct {
		FullName *string
		models.UserCreateParams
	}

	type args struct {
		params models.UserCreateParams
	}

	t.Run("correct_user", func(t *testing.T) {
		t.Parallel()

		ucp := postgresqlrandom.UserCreateParams()

		want := want{
			FullName:         pointers.New(*ucp.FirstName + " " + *ucp.LastName), // repo responsibility
			UserCreateParams: *ucp,
		}

		args := args{
			params: *ucp,
		}

		got, err := userRepo.Create(context.Background(), testPool, &args.params)
		require.NoError(t, err)

		assert.Equal(t, want.FullName, got.FullName)
		assert.Equal(t, want.ExternalID, got.ExternalID)
		assert.Equal(t, want.Email, got.Email)
		assert.Equal(t, want.Username, got.Username)
		assert.Equal(t, want.RoleRank, got.RoleRank)
		assert.Equal(t, want.Scopes, got.Scopes)
		assert.Equal(t, want.FirstName, got.FirstName)
		assert.Equal(t, want.LastName, got.LastName)
	})

	t.Run("role_rank_less_than_zero", func(t *testing.T) {
		t.Parallel()

		ucp := postgresqlrandom.UserCreateParams()
		ucp.RoleRank = -1

		args := args{
			params: *ucp,
		}

		_, err := userRepo.Create(context.Background(), testPool, &args.params)
		require.Error(t, err)

		require.ErrorContains(t, err, errViolatesCheckConstraint)
	})
}
