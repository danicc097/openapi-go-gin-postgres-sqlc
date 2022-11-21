package services_test

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pointers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/repostesting"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

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

	guestRole, userRole, advancedUserRole, managerRole, adminRole := getRoles(t, authzsvc)

	_, normalUser, advancedUser, _, adminUser := fakeUsers(guestRole, userRole, advancedUserRole, managerRole, adminRole)

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
				id:     normalUser.UserID.String(),
				caller: normalUser,
			},
			want: want{
				FirstName: pointers.New("changed"),
				LastName:  advancedUser.LastName,
			},
		},
		{
			name: "cannot_update_different_user",
			args: args{
				params: &models.UpdateUserRequest{},
				id:     normalUser.UserID.String(),
				caller: advancedUser,
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
				id:     normalUser.UserID.String(),
				caller: adminUser,
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

			urepo := repostesting.NewFakeUser([]*db.User{normalUser, advancedUser, adminUser})

			u := services.NewUser(logger, urepo, authzsvc)
			got, err := u.Update(context.Background(), &pgxpool.Pool{}, tc.args.id, tc.args.caller, tc.args.params)
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

	// TODO create users on demand with parameterized tests.
	// e.g. cannot_set_scope_unassigned_to_self  and can_set_scopes_asigned_to_self
	// should have test struct field{callerScopes: []...} , therefore
	// with all relevant parameters set.
	guestRole, userRole, advancedUserRole, managerRole, adminRole := getRoles(t, authzsvc)

	guestUser, normalUser, advancedUser, managerUser, adminUser := fakeUsers(guestRole, userRole, advancedUserRole, managerRole, adminRole)

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
				id:     normalUser.UserID.String(),
				caller: managerUser,
			},
			want: want{
				Scopes: []string{string(models.ScopeUsersRead), string(models.ScopeTestScope)},
				Rank:   managerRole.Rank,
			},
		},
		{
			name: "cannot_update_to_role_higher_than_self",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Role: (*models.Role)(pointers.New(string(models.RoleAdmin))),
				},
				id:     normalUser.UserID.String(),
				caller: managerUser,
			},
			error: "cannot set a user rank higher than self",
		},
		{
			name: "cannot_set_scope_unassigned_to_self",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Scopes: &[]models.Scope{models.ScopeUsersRead, models.ScopeProjectSettingsWrite, models.ScopeUsersWrite},
				},
				id:     normalUser.UserID.String(),
				caller: adminUser,
			},
			error: "cannot set a scope unassigned to self",
		},
		{
			name: "can_set_scope_assigned_to_self",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Scopes: &[]models.Scope{models.ScopeUsersRead, models.ScopeProjectSettingsWrite},
				},
				id:     normalUser.UserID.String(),
				caller: adminUser,
			},
			want: want{
				Scopes: []string{string(models.ScopeUsersRead), string(models.ScopeProjectSettingsWrite)},
				Rank:   normalUser.RoleRank,
			},
		},
		{
			name: "cannot_update_own_auth_information",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Scopes: &[]models.Scope{},
				},
				id:     managerUser.UserID.String(),
				caller: managerUser,
			},
			error: "cannot update your own authorization information",
		},
		{
			name: "cannot_demote_role_if_not_admin",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Role: (*models.Role)(pointers.New(string(models.RoleGuest))),
				},
				id:     advancedUser.UserID.String(),
				caller: managerUser,
			},
			error: "cannot demote a user role",
		},
		{
			name: "cannot_unassign_scopes_if_not_admin",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Scopes: &[]models.Scope{},
				},
				id:     normalUser.UserID.String(),
				caller: managerUser,
			},
			error: "cannot unassign a user's scope",
		},
		{
			name: "can_unassign_scopes_if_admin",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Scopes: &[]models.Scope{},
				},
				id:     normalUser.UserID.String(),
				caller: adminUser,
			},
			want: want{
				Scopes: []string{},
				Rank:   normalUser.RoleRank,
			},
		},
		{
			name: "can_demote_role_if_admin",
			args: args{
				params: &models.UpdateUserAuthRequest{
					Role: (*models.Role)(pointers.New(string(models.RoleGuest))),
				},
				id:     advancedUser.UserID.String(),
				caller: adminUser,
			},
			want: want{
				Rank:   guestUser.RoleRank,
				Scopes: advancedUser.Scopes,
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			urepo := repostesting.NewFakeUser([]*db.User{normalUser, advancedUser, managerUser, adminUser})

			u := services.NewUser(logger, urepo, authzsvc)
			got, err := u.UpdateUserAuthorization(context.Background(), &pgxpool.Pool{}, tc.args.id, tc.args.caller, tc.args.params)
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

func fakeUsers(guestRole, userRole, advancedUserRole, managerRole, adminRole services.Role) (*db.User, *db.User, *db.User, *db.User, *db.User) {
	guestUser := &db.User{
		UserID:   uuid.New(),
		RoleRank: guestRole.Rank,
		Scopes:   []string{string(models.ScopeTestScope)},
	}
	normalUser := &db.User{
		UserID:   uuid.New(),
		RoleRank: userRole.Rank,
		Scopes:   []string{string(models.ScopeTestScope)},
	}
	advancedUser := &db.User{
		UserID:   uuid.New(),
		RoleRank: advancedUserRole.Rank,
		Scopes:   []string{string(models.ScopeTestScope)},
	}
	managerUser := &db.User{
		UserID:   uuid.New(),
		RoleRank: managerRole.Rank,
		Scopes:   []string{string(models.ScopeUsersRead), string(models.ScopeTestScope)},
	}
	adminUser := &db.User{
		UserID:   uuid.New(),
		RoleRank: adminRole.Rank,
		Scopes:   []string{string(models.ScopeUsersRead), string(models.ScopeProjectSettingsWrite)},
	}

	return guestUser, normalUser, advancedUser, managerUser, adminUser
}

func getRoles(t *testing.T, authzsvc *services.Authorization) (services.Role, services.Role, services.Role, services.Role, services.Role) {
	t.Helper()

	guestRole, err := authzsvc.RoleByName(string(models.RoleGuest))
	if err != nil {
		t.Fatalf("RoleByName: %v", err)
	}
	userRole, err := authzsvc.RoleByName(string(models.RoleUser))
	if err != nil {
		t.Fatalf("RoleByName: %v", err)
	}
	advancedUserRole, err := authzsvc.RoleByName(string(models.RoleAdvancedUser))
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

	return guestRole, userRole, advancedUserRole, managerRole, adminRole
}
