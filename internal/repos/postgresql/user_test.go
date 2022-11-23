package postgresql_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pointers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/stretchr/testify/assert"
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

	userRepo := postgresql.NewUser()

	// role, err := ff.authzsvc.RoleByName(string(params.Role))
	// if err != nil {
	// 	return nil, fmt.Errorf("authzsvc.RoleByName: %w", err)
	// }
	type want struct {
		FullName *string
		repos.UserCreateParams
	}

	type args struct {
		params repos.UserCreateParams
	}

	t.Run("correct user", func(t *testing.T) {
		t.Parallel()

		ucp := repos.UserCreateParams{
			Username:   testutil.RandomNameIdentifier(1, "-") + testutil.RandomName(),
			Email:      testutil.RandomEmail(),
			FirstName:  pointers.New(testutil.RandomFirstName()),
			LastName:   pointers.New(testutil.RandomLastName()),
			ExternalID: testutil.RandomString(10),
			Scopes:     []string{"scope1", "scope2"},
			RoleRank:   int16(1),
		}

		want := want{
			FullName:         pointers.New(*ucp.FirstName + " " + *ucp.LastName),
			UserCreateParams: ucp,
		}

		args := args{
			params: ucp,
		}

		got, err := userRepo.Create(context.Background(), testpool, args.params)
		if err != nil {
			t.Fatalf("unexpected error = %v", err)
		}

		assert.Equal(t, want.FullName, got.FullName)
		assert.Equal(t, want.ExternalID, got.ExternalID)
		assert.Equal(t, want.Email, got.Email)
		assert.Equal(t, want.Username, got.Username)
		assert.Equal(t, want.RoleRank, got.RoleRank)
		assert.Equal(t, want.Scopes, got.Scopes)
		assert.Equal(t, want.FirstName, got.FirstName)
		assert.Equal(t, want.LastName, got.LastName)
	})

	t.Run("role rank less than zero", func(t *testing.T) {
		t.Parallel()

		ucp := repos.UserCreateParams{
			Username:   testutil.RandomNameIdentifier(1, "-") + testutil.RandomName(),
			Email:      testutil.RandomEmail(),
			FirstName:  pointers.New(testutil.RandomFirstName()),
			LastName:   pointers.New(testutil.RandomLastName()),
			ExternalID: testutil.RandomString(10),
			Scopes:     []string{"scope1", "scope2"},
			RoleRank:   int16(-1),
		}

		args := args{
			params: ucp,
		}

		errContains := "violates check constraint"

		_, err := userRepo.Create(context.Background(), testpool, args.params)
		if err == nil {
			t.Fatalf("expected error = '%v' but got nothing", errContains)
		}
		assert.Contains(t, err.Error(), errContains)
	})
}
