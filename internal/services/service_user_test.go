package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/repostesting"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/reposwrappers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

type testUsers struct {
	guest, user, advancedUser, manager, admin *servicetestutil.CreateUserResult
}

func TestUser_UpdateUser(t *testing.T) {
	t.Parallel()

	logger := zaptest.NewLogger(t).Sugar()

	authzsvc := newTestAuthzService(t)

	type args struct {
		params *models.UpdateUserRequest
		id     db.UserID
		caller *db.User
	}
	type want struct {
		FirstName *string
		LastName  *string
	}

	testUsers := createTestUsers(t)

	tests := []struct {
		name  string
		args  args
		want  want
		error string
	}{
		{
			name: "user_updated",
			args: args{
				params: &models.UpdateUserRequest{
					FirstName: pointers.New("changed"),
				},
				id:     testUsers.user.User.UserID,
				caller: testUsers.user.User,
			},
			want: want{
				FirstName: pointers.New("changed"),
				LastName:  testUsers.user.User.LastName,
			},
		},
		{
			name: "cannot_update_different_user",
			args: args{
				params: &models.UpdateUserRequest{},
				id:     testUsers.user.User.UserID,
				caller: testUsers.advancedUser.User,
			},
			error: "cannot change another user's information",
		},
		{
			name: "admin_can_update_different_user",
			args: args{
				params: &models.UpdateUserRequest{
					FirstName: pointers.New("changed"),
					LastName:  pointers.New("changed"),
				},
				id:     testUsers.user.User.UserID,
				caller: testUsers.admin.User,
			},
			want: want{
				FirstName: pointers.New("changed"),
				LastName:  pointers.New("changed"),
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			urepo := reposwrappers.NewUserWithRetry(postgresql.NewUser(), 10, 65*time.Millisecond)

			notificationrepo := repostesting.NewFakeNotification()

			ctx := context.Background()
			tx, _ := testPool.BeginTx(ctx, pgx.TxOptions{})
			defer tx.Rollback(ctx)

			u := services.NewUser(logger, urepo, notificationrepo, authzsvc)
			got, err := u.Update(ctx, tx, tc.args.id, tc.args.caller, tc.args.params)
			if (err != nil) && tc.error == "" {
				t.Fatalf("unexpected error = %v", err)
			}
			if tc.error != "" {
				require.Error(t, err)

				assert.Equal(t, tc.error, err.Error())

				return
			}

			assert.Equal(t, *tc.want.FirstName, *got.FirstName)
			assert.Equal(t, *tc.want.LastName, *got.LastName)
		})
	}
}

func TestUser_UpdateUserAuthorization(t *testing.T) {
	t.Parallel()

	logger := zaptest.NewLogger(t).Sugar()

	authzsvc := newTestAuthzService(t)

	// TODO create users on demand with parameterized tests. same as repo ucp but using FakeUserRepo instead
	// e.g. cannot_set_scope_unassigned_to_self  and can_set_scopes_asigned_to_self
	// should have test struct field{callerScopes: []...} , therefore when we look at the test case
	// we see all relevant parameters and input.

	testUsers := createTestUsers(t)

	type args struct {
		params *models.UpdateUserAuthRequest
		id     db.UserID
		caller *db.User
	}
	type want struct {
		Scopes models.Scopes
		Rank   int
	}

	tests := []struct {
		name  string
		args  args
		want  want
		error string
	}{
		{
			name: "user_updated_up_to_same_rank_and_scopes_allowed",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Scopes: pointers.New(authzsvc.DefaultScopes(models.RoleManager)),
					Role:   pointers.New(models.RoleManager),
				},
				id:     testUsers.user.User.UserID,
				caller: testUsers.manager.User,
			},
			want: want{
				Scopes: authzsvc.DefaultScopes(models.RoleManager),
				Rank:   authzsvc.RoleByName(models.RoleManager).Rank,
			},
		},
		{
			name: "cannot_update_to_role_higher_than_self",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Role: pointers.New(models.RoleAdmin),
				},
				id:     testUsers.user.User.UserID,
				caller: testUsers.manager.User,
			},
			error: "cannot set a user rank higher than self",
		},
		{
			name: "cannot_set_scope_unassigned_to_self",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Scopes: &models.Scopes{models.ScopeUsersRead, models.ScopeProjectSettingsWrite, models.ScopeUsersWrite},
				},
				id:     testUsers.user.User.UserID,
				caller: testUsers.admin.User,
			},
			error: "cannot set a scope unassigned to self",
		},
		{
			name: "can_set_scopes_assigned_to_self_without_role_update",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Scopes: pointers.New(authzsvc.DefaultScopes(models.RoleAdmin)),
				},
				id:     testUsers.user.User.UserID,
				caller: testUsers.admin.User,
			},
			want: want{
				Scopes: testUsers.admin.User.Scopes,
				Rank:   testUsers.user.User.RoleRank, // unchanged
			},
		},
		{
			name: "cannot_update_own_auth_information",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Scopes: &models.Scopes{},
				},
				id:     testUsers.manager.User.UserID,
				caller: testUsers.manager.User,
			},
			error: "cannot update your own authorization information",
		},
		{
			name: "cannot_demote_role_if_not_admin",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Role: pointers.New(models.RoleGuest),
				},
				id:     testUsers.advancedUser.User.UserID,
				caller: testUsers.manager.User,
			},
			error: "cannot demote a user role",
		},
		{
			name: "cannot_unassign_scopes_if_not_admin",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Scopes: &models.Scopes{},
				},
				id:     testUsers.user.User.UserID,
				caller: testUsers.manager.User,
			},
			error: "cannot unassign a user's scope",
		},
		{
			name: "can_unassign_scopes_if_admin",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Scopes: &models.Scopes{},
				},
				id:     testUsers.user.User.UserID,
				caller: testUsers.admin.User,
			},
			want: want{
				Scopes: models.Scopes{},
				Rank:   testUsers.user.User.RoleRank, // unchanged
			},
		},
		{
			name: "can_demote_role_if_admin",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Role: pointers.New(models.RoleGuest),
				},
				id:     testUsers.advancedUser.User.UserID,
				caller: testUsers.admin.User,
			},
			want: want{
				Rank:   authzsvc.RoleByName(models.RoleGuest).Rank,
				Scopes: authzsvc.DefaultScopes(models.RoleGuest), // scopes are reset on role change
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			urepo := reposwrappers.NewUserWithRetry(postgresql.NewUser(), 10, 65*time.Millisecond)

			notificationrepo := repostesting.NewFakeNotification()

			ctx := context.Background()
			tx, _ := testPool.BeginTx(ctx, pgx.TxOptions{})
			defer tx.Rollback(ctx)

			u := services.NewUser(logger, urepo, notificationrepo, authzsvc)
			got, err := u.UpdateUserAuthorization(ctx, tx, tc.args.id, tc.args.caller, tc.args.params)
			if (err != nil) && tc.error == "" {
				t.Fatalf("unexpected error = %v", err)
			}
			if tc.error != "" {
				require.Error(t, err)

				assert.Equal(t, tc.error, err.Error())

				return
			}

			assert.Equal(t, tc.want.Scopes, got.Scopes)
			assert.Equal(t, tc.want.Rank, got.RoleRank)
		})
	}
}

// dont use repos here, we want the actual services logic.
func createTestUsers(t *testing.T) testUsers {
	t.Helper()

	ff := newTestFixtureFactory(t)

	guest, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
		Role:       models.RoleAdmin,
		WithAPIKey: true,
	})
	require.NoError(t, err)

	user, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
		Role:       models.RoleUser,
		WithAPIKey: true,
	})
	require.NoError(t, err)

	advancedUser, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
		Role:       models.RoleAdvancedUser,
		WithAPIKey: true,
	})
	require.NoError(t, err)

	manager, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
		Role:       models.RoleManager,
		WithAPIKey: true,
	})
	require.NoError(t, err)

	admin, err := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
		Role:       models.RoleAdmin,
		WithAPIKey: true,
	})
	require.NoError(t, err)

	return testUsers{
		guest:        guest,
		user:         user,
		advancedUser: advancedUser,
		manager:      manager,
		admin:        admin,
	}
}
