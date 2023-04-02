package services_test

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pointers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/repostesting"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

type testUsers struct {
	guest, user, advanced, manager, admin *db.User
}

type testRoles struct {
	guest, user, advanced, manager, admin services.Role
}

func TestUser_UpdateUser(t *testing.T) {
	t.Parallel()

	logger := zaptest.NewLogger(t)

	authzsvc, err := services.NewAuthorization(zaptest.NewLogger(t), "testdata/scopes.json", "testdata/roles.json")
	if err != nil {
		t.Fatalf("NewAuthorization: %v", err)
	}

	type args struct {
		params *models.UpdateUserRequest
		id     string
		caller *db.User
	}
	type want struct {
		FirstName *string
		LastName  *string
	}

	testUsers := createTestUsers(t, testPool, authzsvc)

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
				id:     testUsers.user.UserID.String(),
				caller: testUsers.user,
			},
			want: want{
				FirstName: pointers.New("changed"),
				LastName:  testUsers.advanced.LastName,
			},
		},
		{
			name: "cannot_update_different_user",
			args: args{
				params: &models.UpdateUserRequest{},
				id:     testUsers.user.UserID.String(),
				caller: testUsers.advanced,
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
				id:     testUsers.user.UserID.String(),
				caller: testUsers.admin,
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

			urepo := postgresql.NewUser()

			notificationrepo := repostesting.NewFakeNotification()

			u := services.NewUser(logger, urepo, notificationrepo, authzsvc)
			got, err := u.Update(context.Background(), testPool, tc.args.id, tc.args.caller, tc.args.params)
			if (err != nil) && tc.error == "" {
				t.Fatalf("unexpected error = %v", err)
			}
			if tc.error != "" {
				if err == nil {
					t.Fatalf("expected error = '%v' but got nothing", tc.error)
				}
				assert.Equal(t, tc.error, err.Error())

				return
			}

			assert.Equal(t, tc.want.FirstName, got.FirstName)
			assert.Equal(t, tc.want.LastName, got.LastName)
		})
	}
}

func TestUser_UpdateUserAuthorization(t *testing.T) {
	t.Parallel()

	logger := zaptest.NewLogger(t)

	authzsvc, err := services.NewAuthorization(zaptest.NewLogger(t), "testdata/scopes.json", "testdata/roles.json")
	if err != nil {
		t.Fatalf("NewAuthorization: %v", err)
	}

	// TODO create users on demand with parameterized tests. same as repo ucp but using FakeUserRepo instead
	// e.g. cannot_set_scope_unassigned_to_self  and can_set_scopes_asigned_to_self
	// should have test struct field{callerScopes: []...} , therefore when we look at the test case
	// we see all relevant parameters and input.

	roles := getRoles(t, authzsvc)

	testUsers := createTestUsers(t, testPool, authzsvc)

	type args struct {
		params *models.UpdateUserAuthRequest
		id     string
		caller *db.User
	}
	type want struct {
		Scopes []string
		Rank   int16
	}

	tests := []struct {
		name  string
		args  args
		want  want
		error string
	}{
		{
			name: "user_updated_up_to_same_rank_and_scopes",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Scopes: &[]models.Scope{models.ScopeUsersRead, models.ScopeTestScope},
					Role:   (*models.Role)(pointers.New(string(models.RoleManager))),
				},
				id:     testUsers.user.UserID.String(),
				caller: testUsers.manager,
			},
			want: want{
				Scopes: []string{string(models.ScopeUsersRead), string(models.ScopeTestScope)},
				Rank:   roles.manager.Rank,
			},
		},
		{
			name: "cannot_update_to_role_higher_than_self",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Role: (*models.Role)(pointers.New(string(models.RoleAdmin))),
				},
				id:     testUsers.user.UserID.String(),
				caller: testUsers.manager,
			},
			error: "cannot set a user rank higher than self",
		},
		{
			name: "cannot_set_scope_unassigned_to_self",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Scopes: &[]models.Scope{models.ScopeUsersRead, models.ScopeProjectSettingsWrite, models.ScopeUsersWrite},
				},
				id:     testUsers.user.UserID.String(),
				caller: testUsers.admin,
			},
			error: "cannot set a scope unassigned to self",
		},
		{
			name: "can_set_scope_assigned_to_self",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Scopes: &[]models.Scope{models.ScopeUsersRead, models.ScopeProjectSettingsWrite},
				},
				id:     testUsers.user.UserID.String(),
				caller: testUsers.admin,
			},
			want: want{
				Scopes: []string{string(models.ScopeUsersRead), string(models.ScopeProjectSettingsWrite)},
				Rank:   testUsers.user.RoleRank,
			},
		},
		{
			name: "cannot_update_own_auth_information",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Scopes: &[]models.Scope{},
				},
				id:     testUsers.manager.UserID.String(),
				caller: testUsers.manager,
			},
			error: "cannot update your own authorization information",
		},
		{
			name: "cannot_demote_role_if_not_admin",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Role: (*models.Role)(pointers.New(string(models.RoleGuest))),
				},
				id:     testUsers.advanced.UserID.String(),
				caller: testUsers.manager,
			},
			error: "cannot demote a user role",
		},
		{
			name: "cannot_unassign_scopes_if_not_admin",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Scopes: &[]models.Scope{},
				},
				id:     testUsers.user.UserID.String(),
				caller: testUsers.manager,
			},
			error: "cannot unassign a user's scope",
		},
		{
			name: "can_unassign_scopes_if_admin",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Scopes: &[]models.Scope{},
				},
				id:     testUsers.user.UserID.String(),
				caller: testUsers.admin,
			},
			want: want{
				Scopes: []string{},
				Rank:   testUsers.user.RoleRank,
			},
		},
		{
			name: "can_demote_role_if_admin",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Role: (*models.Role)(pointers.New(string(models.RoleGuest))),
				},
				id:     testUsers.advanced.UserID.String(),
				caller: testUsers.admin,
			},
			want: want{
				Rank:   roles.guest.Rank,
				Scopes: testUsers.advanced.Scopes,
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			urepo := postgresql.NewUser()

			notificationrepo := repostesting.NewFakeNotification()

			u := services.NewUser(logger, urepo, notificationrepo, authzsvc)
			got, err := u.UpdateUserAuthorization(context.Background(), testPool, tc.args.id, tc.args.caller, tc.args.params)
			if (err != nil) && tc.error == "" {
				t.Fatalf("unexpected error = %v", err)
			}
			if tc.error != "" {
				if err == nil {
					t.Fatalf("expected error = '%v' but got nothing", tc.error)
				}
				assert.Equal(t, tc.error, err.Error())

				return
			}
			assert.Equal(t, tc.want.Scopes, got.Scopes)
			assert.Equal(t, tc.want.Rank, got.RoleRank)
		})
	}
}

func createTestUsers(t *testing.T, pool *pgxpool.Pool, authzsvc *services.Authorization) testUsers {
	t.Helper()

	roles := getRoles(t, authzsvc)

	urepo := postgresql.NewUser()

	ucp := postgresqltestutil.RandomUserCreateParams(t)
	ucp.RoleRank = roles.guest.Rank
	ucp.Scopes = []string{string(models.ScopeTestScope)}
	guest, _ := urepo.Create(context.Background(), pool, ucp)

	ucp = postgresqltestutil.RandomUserCreateParams(t)
	ucp.RoleRank = roles.user.Rank
	ucp.Scopes = []string{string(models.ScopeTestScope)}
	user, _ := urepo.Create(context.Background(), pool, ucp)

	ucp = postgresqltestutil.RandomUserCreateParams(t)
	ucp.RoleRank = roles.advanced.Rank
	ucp.Scopes = []string{string(models.ScopeTestScope)}
	advanced, _ := urepo.Create(context.Background(), pool, ucp)

	ucp = postgresqltestutil.RandomUserCreateParams(t)
	ucp.RoleRank = roles.manager.Rank
	ucp.Scopes = []string{string(models.ScopeUsersRead), string(models.ScopeTestScope)}
	manager, _ := urepo.Create(context.Background(), pool, ucp)

	ucp = postgresqltestutil.RandomUserCreateParams(t)
	ucp.RoleRank = roles.admin.Rank
	ucp.Scopes = []string{string(models.ScopeUsersRead), string(models.ScopeProjectSettingsWrite)}
	admin, _ := urepo.Create(context.Background(), pool, ucp)

	return testUsers{
		guest:    guest,
		user:     user,
		advanced: advanced,
		manager:  manager,
		admin:    admin,
	}
}

func getRoles(t *testing.T, authzsvc *services.Authorization) testRoles {
	t.Helper()

	guestRole, err := authzsvc.RoleByName(string(models.RoleGuest))
	if err != nil {
		t.Fatalf("RoleByName: %v", err)
	}
	userRole, err := authzsvc.RoleByName(string(models.RoleUser))
	if err != nil {
		t.Fatalf("RoleByName: %v", err)
	}
	advancedRole, err := authzsvc.RoleByName(string(models.RoleAdvancedUser))
	if err != nil {
		t.Fatalf("RoleByName: %v", err)
	}
	managerRole, err := authzsvc.RoleByName(string(models.RoleManager))
	if err != nil {
		t.Fatalf("RoleByName: %v", err)
	}
	adminRole, err := authzsvc.RoleByName(string(models.RoleAdmin))
	if err != nil {
		t.Fatalf("RoleByName: %v", err)
	}

	return testRoles{
		guest:    guestRole,
		user:     userRole,
		advanced: advancedRole,
		manager:  managerRole,
		admin:    adminRole,
	}
}
