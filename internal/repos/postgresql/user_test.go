package postgresql_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/go-playground/assert/v2"
)

/**
 *
 * TODO integration tests only
 *
 * https://stackoverflow.com/questions/4452928/should-i-bother-unit-testing-my-repository-layer
 *
 *
 */

func TestUser_Update(t *testing.T) {
	t.Parallel()

	type args struct {
		params repos.UserUpdateParams
	}
	tests := []struct {
		name    string
		args    args
		want    *db.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			u := postgresql.NewUser()
			got, err := u.Update(context.Background(), testpool, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO at least one index query test should test joins, orderby...
func TestUser_UserByEmail(t *testing.T) {
	t.Parallel()

	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    *db.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			u := postgresql.NewUser()
			got, err := u.UserByEmail(context.Background(), testpool, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.UserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.UserByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_UserByID(t *testing.T) {
	t.Parallel()

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *db.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			u := postgresql.NewUser()
			got, err := u.UserByID(context.Background(), testpool, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.UserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.UserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_UserByAPIKey(t *testing.T) {
	t.Parallel()

	type args struct {
		apiKey string
	}
	tests := []struct {
		name    string
		args    args
		want    *db.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			u := postgresql.NewUser()
			got, err := u.UserByAPIKey(context.Background(), testpool, tt.args.apiKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.UserByAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.UserByAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_CreateAPIKey(t *testing.T) {
	t.Parallel()

	type args struct {
		user *db.User
	}
	tests := []struct {
		name    string
		args    args
		want    *db.UserAPIKey
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			u := postgresql.NewUser()
			got, err := u.CreateAPIKey(context.Background(), testpool, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.CreateAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.CreateAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Create(t *testing.T) {
	t.Parallel()

	// role, err := ff.authzsvc.RoleByName(string(params.Role))
	// if err != nil {
	// 	return nil, fmt.Errorf("authzsvc.RoleByName: %w", err)
	// }

	type args struct {
		params repos.UserCreateParams
	}
	tests := []struct {
		name  string
		args  args
		want  *db.User
		error string
	}{
		// {
		// 	name: "test",
		// 	args: args{
		// 		params: repos.UserCreateParams{
		// 			Username:   testutil.RandomNameIdentifier(1, "-") + testutil.RandomName(),
		// 			Email:      testutil.RandomEmail(),
		// 			FirstName:  pointers.New(testutil.RandomFirstName()),
		// 			LastName:   pointers.New(testutil.RandomLastName()),
		// 			ExternalID: testutil.RandomString(10),
		// 			Scopes:     []string{"scope1", "scope2"},
		// 			RoleRank:   int16(1),
		// 		},
		// 	},
		// 	want: &db.User{},
		// },
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			u := postgresql.NewUser()
			got, err := u.Create(context.Background(), testpool, tc.args.params)
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

			assert.Equal(t, tc.want.FullName, got.FullName)
			assert.Equal(t, tc.want.Teams, got.Teams)
		})
	}
}
