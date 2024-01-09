package postgresql_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/reposwrappers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUser_Update(t *testing.T) {
	t.Parallel()

	user := postgresqltestutil.NewRandomUser(t, testPool)

	type args struct {
		id     db.UserID
		params db.UserUpdateParams
	}
	type params struct {
		name        string
		args        args
		want        *db.User
		errContains string
	}
	tests := []params{
		{
			name: "updated",
			args: args{
				id: user.UserID,
				params: db.UserUpdateParams{
					RoleRank: pointers.New(10),
					Scopes:   &models.Scopes{"test", "test", "test"},
				},
			},
			want: func() *db.User {
				u := *user
				u.RoleRank = 10
				u.Scopes = models.Scopes{"test"}

				return &u
			}(),
		},
	}
	for _, tc := range tests {
		tc := tc
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
				assert.ErrorContains(t, err, tc.errContains)

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

	user := postgresqltestutil.NewRandomUser(t, testPool)

	type args struct {
		id db.UserID
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
		tc := tc
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
				assert.ErrorContains(t, err, tc.errContains)

				return
			}
			assert.ErrorContains(t, err, tc.errContains)
		})
	}
}

func TestUser_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	userRepo := reposwrappers.NewUserWithRetry(postgresql.NewUser(), 10, 65*time.Millisecond)

	teamRepo := postgresql.NewTeam()
	projectRepo := postgresql.NewProject()

	project, err := projectRepo.ByName(ctx, testPool, models.ProjectDemo)
	require.NoError(t, err)

	team := postgresqltestutil.NewRandomTeam(t, testPool, project.ProjectID)
	user := postgresqltestutil.NewRandomUser(t, testPool)

	_, err = db.CreateUserTeam(ctx, testPool, &db.UserTeamCreateParams{Member: user.UserID, TeamID: team.TeamID})
	require.NoError(t, err)

	team, err = teamRepo.ByID(ctx, testPool, team.TeamID, db.WithTeamJoin(db.TeamJoins{MembersTeam: true, Project: true}))
	require.NoError(t, err)

	uniqueCallback := func(t *testing.T, res *db.User) {
		assert.Equal(t, res.UserID, user.UserID)
	}

	uniqueTestCases := []filterTestCase[*db.User]{
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
		tc := tc
		runGenericFilterTests(t, tc)
	}

	nonUniqueTestCases := []filterTestCase[[]db.User]{
		{
			// FIXME: https://github.com/danicc097/openapi-go-gin-postgres-sqlc/actions/runs/7467525226/job/20321213188?pr=206
			// due to user with project-member scope getting added since we're using projectdemo.
			// TODO: should have helper NewRandomProject just for repo tests so we can test
			// side effects on projects, which are hardcoded (might need same kind of helpers
			// for kanban steps, work item types later)
			// OR: run in tx
			name:       "team_id",
			filter:     team.TeamID,
			repoMethod: reflect.ValueOf(userRepo.ByTeam),
			callback: func(t *testing.T, res []db.User) {
				assert.Len(t, res, 1)
				assert.Equal(t, res[0].UserID, user.UserID)
				assert.Equal(t, (*team.TeamMembersJoin)[0].UserID, user.UserID)
				assert.Equal(t, team.ProjectID, project.ProjectID)
			},
		},
		{
			name:       "project_id",
			filter:     project.ProjectID,
			repoMethod: reflect.ValueOf(userRepo.ByProject),
			callback: func(t *testing.T, res []db.User) {
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
		tc := tc
		runGenericFilterTests(t, tc)
	}
}

func TestUser_UserAPIKeys(t *testing.T) {
	t.Parallel()

	userRepo := reposwrappers.NewUserWithRetry(postgresql.NewUser(), 10, 65*time.Millisecond)

	t.Run("correct_api_key_creation", func(t *testing.T) {
		t.Parallel()

		user := postgresqltestutil.NewRandomUser(t, testPool)

		uak, err := userRepo.CreateAPIKey(context.Background(), testPool, user)
		require.NoError(t, err)
		assert.NotEmpty(t, uak.APIKey)
		assert.Equal(t, uak.UserID, user.UserID)
		assert.Equal(t, uak.UserAPIKeyID, *user.APIKeyID)
	})

	t.Run("no_api_key_created_when_user_does_not_exist", func(t *testing.T) {
		t.Parallel()

		errContains := "could not save api key"

		_, err := userRepo.CreateAPIKey(context.Background(), testPool, &db.User{UserID: db.NewUserID(uuid.New())})

		assert.ErrorContains(t, err, errContains)
	})

	t.Run("can_get_user_by_api_key", func(t *testing.T) {
		t.Parallel()

		newUser := postgresqltestutil.NewRandomUser(t, testPool)

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
		assert.ErrorContains(t, err, errContains)
	})

	t.Run("can_delete_an_api_key", func(t *testing.T) {
		t.Parallel()

		newUser := postgresqltestutil.NewRandomUser(t, testPool)

		uak, err := userRepo.CreateAPIKey(context.Background(), testPool, newUser)
		require.NoError(t, err)

		deletedUak, err := userRepo.DeleteAPIKey(context.Background(), testPool, uak.APIKey)
		require.NoError(t, err)
		assert.Equal(t, deletedUak.APIKey, uak.APIKey)

		_, err = userRepo.ByAPIKey(context.Background(), testPool, uak.APIKey)
		assert.ErrorContains(t, err, errNoRows)
	})
}

func TestUser_Create(t *testing.T) {
	t.Parallel()

	userRepo := reposwrappers.NewUserWithRetry(postgresql.NewUser(), 10, 65*time.Millisecond)

	type want struct {
		FullName *string
		db.UserCreateParams
	}

	type args struct {
		params db.UserCreateParams
	}

	t.Run("correct_user", func(t *testing.T) {
		t.Parallel()

		ucp := postgresqltestutil.RandomUserCreateParams(t)

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

		ucp := postgresqltestutil.RandomUserCreateParams(t)
		ucp.RoleRank = -1

		args := args{
			params: *ucp,
		}

		_, err := userRepo.Create(context.Background(), testPool, &args.params)
		require.Error(t, err)

		assert.ErrorContains(t, err, errViolatesCheckConstraint)
	})
}
