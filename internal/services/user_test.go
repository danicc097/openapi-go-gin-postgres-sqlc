package services_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicestesting"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func TestUser_Create(t *testing.T) {
	type fields struct {
		urepo  services.UserRepo
		logger *zap.Logger
		pool   *pgxpool.Pool
	}
	type args struct {
		ctx    context.Context
		params models.CreateUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.CreateUserResponse
		wantErr bool
	}{
		{
			name: "user created",
			fields: fields{
				urepo: &servicestesting.FakeUserRepo{CreateStub: func(ctx context.Context, params models.CreateUserRequest) (models.CreateUserResponse, error) {
					return models.CreateUserResponse{AccessToken: "abcd", UserId: 1}, nil
				}},
				logger: zaptest.NewLogger(t),
				pool:   &pgxpool.Pool{},
			},
			args: args{
				params: models.CreateUserRequest{
					Username: "username",
					Email:    "email@mail.com",
					Password: "password",
				},
			},
			want: models.CreateUserResponse{AccessToken: "abcd", UserId: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := services.NewUser(tt.fields.urepo, tt.fields.logger, tt.fields.pool)
			got, err := u.Create(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
