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

func TestUser_Update(t *testing.T) {
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
							FirstName: pointers.String("changed"),
							LastName:  pointers.String("last"),
						}, nil
					},
				},
			},
			args: args{
				params: &models.UpdateUserRequest{
					FirstName: pointers.String("changed"),
				},
				id:     uuid1.String(),
				caller: &db.User{UserID: uuid1},
			},
			want: want{
				FirstName: pointers.String("changed"),
				LastName:  pointers.String("last"),
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
