package services_test

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pointers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
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

	type fields struct {
		urepo repos.User
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

	userRole, advancedUserRole, managerRole, adminRole := getRoles(t, authzsvc)

	normalUser, advancedUser, _, adminUser := fakeUsers(userRole, advancedUserRole, managerRole, adminRole)

	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
		error  string
	}{
		{
			name: "user_updated",
			fields: fields{
				urepo: &repostesting.FakeUser{
					UserByIDStub: func(ctx context.Context, d db.DBTX, s string) (*db.User, error) {
						return normalUser, nil
					},
					UpdateStub: func(ctx context.Context, d db.DBTX, uup repos.UserUpdateParams) (*db.User, error) {
						u := *normalUser
						u.FirstName = uup.FirstName

						return &u, nil
					},
				},
			},
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
			fields: fields{
				urepo: &repostesting.FakeUser{
					UserByIDStub: func(ctx context.Context, d db.DBTX, s string) (*db.User, error) {
						return normalUser, nil
					},
				},
			},
			args: args{
				params: &models.UpdateUserRequest{},
				id:     normalUser.UserID.String(),
				caller: advancedUser,
			},
			error: "cannot change another user's information",
		},
		{
			name: "admin_can_update_different_user",
			fields: fields{
				urepo: &repostesting.FakeUser{
					UserByIDStub: func(ctx context.Context, d db.DBTX, s string) (*db.User, error) {
						return normalUser, nil
					},
					UpdateStub: func(ctx context.Context, d db.DBTX, uup repos.UserUpdateParams) (*db.User, error) {
						u := *normalUser
						u.FirstName = uup.FirstName
						u.LastName = uup.LastName

						return &u, nil
					},
				},
			},
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

			u := services.NewUser(logger, tc.fields.urepo, authzsvc)
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

	userRole, advancedUserRole, managerRole, adminRole := getRoles(t, authzsvc)

	normalUser, advancedUser, managerUser, adminUser := fakeUsers(userRole, advancedUserRole, managerRole, adminRole)

	type fields struct {
		urepo repos.User
	}
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
		name   string
		fields fields
		args   args
		want   want
		error  string
	}{
		{
			name: "user_updated_up_to_same_rank_and_scopes",
			fields: fields{
				urepo: &repostesting.FakeUser{
					UserByIDStub: func(ctx context.Context, d db.DBTX, s string) (*db.User, error) {
						return normalUser, nil
					},
					UpdateStub: func(ctx context.Context, d db.DBTX, uup repos.UserUpdateParams) (*db.User, error) {
						u := *normalUser
						u.Scopes = *uup.Scopes
						u.RoleRank = *uup.Rank

						return &u, nil
					},
				},
			},
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
			fields: fields{
				urepo: &repostesting.FakeUser{
					UserByIDStub: func(ctx context.Context, d db.DBTX, s string) (*db.User, error) {
						return normalUser, nil
					},
				},
			},
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
			fields: fields{
				urepo: &repostesting.FakeUser{
					UserByIDStub: func(ctx context.Context, d db.DBTX, s string) (*db.User, error) {
						return normalUser, nil
					},
				},
			},
			args: args{
				params: &models.UpdateUserAuthRequest{
					Scopes: &[]models.Scope{models.ScopeUsersRead, models.ScopeProjectSettingsWrite, models.ScopeUsersWrite},
				},
				id:     normalUser.UserID.String(),
				caller: managerUser,
			},
			error: "cannot set a scope unassigned to self",
		},
		{
			name: "cannot_update_own_auth_information",
			fields: fields{
				urepo: &repostesting.FakeUser{
					UserByIDStub: func(ctx context.Context, d db.DBTX, s string) (*db.User, error) {
						return managerUser, nil
					},
				},
			},
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
			fields: fields{
				urepo: &repostesting.FakeUser{
					UserByIDStub: func(ctx context.Context, d db.DBTX, s string) (*db.User, error) {
						return advancedUser, nil
					},
				},
			},
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
			fields: fields{
				urepo: &repostesting.FakeUser{
					UserByIDStub: func(ctx context.Context, d db.DBTX, s string) (*db.User, error) {
						return normalUser, nil
					},
				},
			},
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
			fields: fields{
				urepo: &repostesting.FakeUser{
					UserByIDStub: func(ctx context.Context, d db.DBTX, s string) (*db.User, error) {
						return normalUser, nil
					},
					UpdateStub: func(ctx context.Context, d db.DBTX, uup repos.UserUpdateParams) (*db.User, error) {
						u := *normalUser
						u.Scopes = *uup.Scopes

						return &u, nil
					},
				},
			},
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
			fields: fields{
				urepo: &repostesting.FakeUser{
					UserByIDStub: func(ctx context.Context, d db.DBTX, s string) (*db.User, error) {
						return advancedUser, nil
					},
					UpdateStub: func(ctx context.Context, d db.DBTX, uup repos.UserUpdateParams) (*db.User, error) {
						u := *advancedUser

						return &u, nil
					},
				},
			},
			args: args{
				params: &models.UpdateUserAuthRequest{
					Role: (*models.Role)(pointers.New(string(models.RoleGuest))),
				},
				id:     advancedUser.UserID.String(),
				caller: adminUser,
			},
			want: want{
				Rank:   advancedUser.RoleRank,
				Scopes: advancedUser.Scopes,
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			u := services.NewUser(logger, tc.fields.urepo, authzsvc)
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

func fakeUsers(userRole, advancedUserRole, managerRole, adminRole services.Role) (*db.User, *db.User, *db.User, *db.User) {
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

	return normalUser, advancedUser, managerUser, adminUser
}

func getRoles(t *testing.T, authzsvc *services.Authorization) (services.Role, services.Role, services.Role, services.Role) {
	t.Helper()

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

	return userRole, advancedUserRole, managerRole, adminRole
}
