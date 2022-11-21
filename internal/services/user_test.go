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

	// TODO same as auth tests below
	uuid1 := uuid.New()
	uuid2 := uuid.New()

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
						return &db.User{UserID: uuid1}, nil
					},
					UpdateStub: func(ctx context.Context, d db.DBTX, uup repos.UserUpdateParams) (*db.User, error) {
						return &db.User{
							UserID:    uuid1,
							FirstName: pointers.New("changed"),
							LastName:  pointers.New("last"),
						}, nil
					},
				},
			},
			args: args{
				params: &models.UpdateUserRequest{
					FirstName: pointers.New("changed"),
				},
				id:     uuid1.String(),
				caller: &db.User{UserID: uuid1},
			},
			want: want{
				FirstName: pointers.New("changed"),
				LastName:  pointers.New("last"),
			},
		},
		{
			name: "cannot_update_different_user",
			fields: fields{
				urepo: &repostesting.FakeUser{
					UserByIDStub: func(ctx context.Context, d db.DBTX, s string) (*db.User, error) {
						return &db.User{UserID: uuid1}, nil
					},
				},
			},
			args: args{
				params: &models.UpdateUserRequest{},
				id:     uuid1.String(),
				caller: &db.User{UserID: uuid2},
			},
			error: "cannot change another user's information",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			u := services.NewUser(logger, tc.fields.urepo, authzsvc)
			got, err := u.Update(context.Background(), &pgxpool.Pool{}, tc.args.id, tc.args.caller, tc.args.params)
			if (err != nil) && tc.error == "" {
				t.Fatalf("User.Create() error = %v, error %v", err, tc.error)
			}
			if tc.error != "" {
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

	userRank, err := authzsvc.RoleByName(string(models.RoleUser))
	if err != nil {
		t.Fatalf("RoleByName: %v", err)
	}
	managerRank, err := authzsvc.RoleByName(string(models.RoleManager))
	if err != nil {
		t.Fatalf("RoleByName: %v", err)
	}
	adminRank, err := authzsvc.RoleByName(string(models.RoleAdmin))
	if err != nil {
		t.Fatalf("RoleByName: %v", err)
	}

	normalUser := &db.User{
		UserID:   uuid.New(),
		RoleRank: userRank.Rank,
		Scopes:   []string{string(models.ScopeTestScope)},
	}
	advancedUser := &db.User{
		UserID:   uuid.New(),
		RoleRank: userRank.Rank,
		Scopes:   []string{string(models.ScopeTestScope)},
	}
	managerUser := &db.User{
		UserID:   uuid.New(),
		RoleRank: managerRank.Rank,
		Scopes:   []string{string(models.ScopeUsersRead), string(models.ScopeTestScope)},
	}
	adminUser := &db.User{
		UserID:   uuid.New(),
		RoleRank: adminRank.Rank,
		Scopes:   []string{string(models.ScopeUsersRead), string(models.ScopeProjectSettingsWrite)},
	}

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
				Rank:   managerRank.Rank,
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
				Rank:   normalUser.RoleRank,
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
