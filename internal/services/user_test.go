package services_test

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/repostesting"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap/zaptest"
)

func TestUser_Update(t *testing.T) {
	_ = repostesting.NewFakeUser()

	_ = zaptest.NewLogger(t)

	_, err := services.NewAuthorization(zaptest.NewLogger(t), "testdata/scopes.json", "testdata/roles.json")
	if err != nil {
		t.Fatalf("NewAuthorization: %v", err)
	}

	type fields struct {
		urepo repos.User
		pool  *pgxpool.Pool
	}
	type args struct {
		ctx    context.Context
		params models.UpdateUserRequest
	}
	// tests := []struct {
	// 	name    string
	// 	fields  fields
	// 	args    args
	// 	want    models.User
	// 	wantErr bool
	// }{
	// 	{
	// 		name: "user_updated",
	// 		fields: fields{
	// 			urepo: fakeUserRepo,
	// 			pool:  &pgxpool.Pool{},
	// 		},
	// 		args: args{
	// 			params: models.UpdateUserRequest{
	// 				FirstName: pointers.String("first"),
	// 				LastName:  pointers.String("last"),
	// 			},
	// 			ctx: context.Background(),
	// 		},
	// 		want: models.User{AccessToken: "abcd", UserId: 1},
	// 	},
	// }

	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		u := services.NewUser(logger, tt.fields.urepo, authzsvc)
	// 		got, err := u.up(tt.args.ctx, tt.args.params)
	// 		if (err != nil) != tt.wantErr {
	// 			t.Errorf("User.Create() error = %v, wantErr %v", err, tt.wantErr)
	// 			return
	// 		}
	// 		if !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("User.Create() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}
