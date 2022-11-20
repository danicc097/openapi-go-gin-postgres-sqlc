package services_test

// import (
// 	"context"
// 	"reflect"
// 	"testing"

// 	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
// 	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
// 	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
// 	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/repostesting"
// 	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
// 	"github.com/jackc/pgx/v4/pgxpool"
// 	"go.uber.org/zap/zaptest"
// )

// func TestUser_Update(t *testing.T) {
// 	logger := zaptest.NewLogger(t)

// 	authzsvc, err := NewAuthorization(zaptest.NewLogger(t), "testdata/scopes.json", "testdata/roles.json")
// 	if err != nil {
// 		t.Fatalf("NewAuthorization: %v", err)
// 	}

// 	type fields struct {
// 		urepo repos.User
// 		pool  *pgxpool.Pool
// 	}
// 	type args struct {
// 		ctx context.Context
// 		// TODO model will come from xoxo (ideally)/sqlc
// 		params models.UpdateUserRequest
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		// TODO model will come from xoxo (ideally)/sqlc
// 		want    models.User
// 		wantErr bool
// 	}{
// 		{
// 			name: "user_created",
// 			fields: fields{
// 				urepo: &repostesting.FakeUser{UpdateStub: func(ctx context.Context, d db.DBTX, params repos.UserUpdateParams) (*db.User, error) {
// 					return db.User{AccessToken: "abcd", UserId: 1}, nil
// 				}},
// 				pool: &pgxpool.Pool{},
// 			},
// 			args: args{
// 				params: models.UpdateUserRequest{
// 					Username: "username",
// 					Email:    "email@mail.com",
// 				},
// 				ctx: context.Background(),
// 			},
// 			want: models.CreateUserResponse{AccessToken: "abcd", UserId: 1},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			u := services.NewUser(logger, tt.fields.urepo, authzsvc)
// 			got, err := u.Update(tt.args.ctx, tt.args.params)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("User.Create() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("User.Create() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
