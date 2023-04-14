package postgresql_test

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUser_Update(t *testing.T) {
	t.Parallel()

	user, _ := postgresqltestutil.NewRandomUser(t, testPool)

	type args struct {
		id     uuid.UUID
		params db.UserUpdateParams
	}
	type params struct {
		name    string
		args    args
		want    *db.User
		wantErr bool
	}
	tests := []params{
		{
			name: "updated",
			args: args{
				id: user.UserID,
				params: db.UserUpdateParams{
					RoleRank: pointers.New[int16](10),
					Scopes:   &[]string{"test", "test", "test"},
				},
			},
			want: func() *db.User {
				u := *user
				u.RoleRank = 10
				u.Scopes = []string{"test"}

				return &u
			}(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			u := postgresql.NewUser()
			got, err := u.Update(context.Background(), testPool, tt.args.id, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.Update() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			got.UpdatedAt = user.UpdatedAt // ignore
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUser_SoftDelete(t *testing.T) {
	t.Parallel()

	user, _ := postgresqltestutil.NewRandomUser(t, testPool)

	type args struct {
		id uuid.UUID
	}
	type params struct {
		name          string
		args          args
		errorContains string
	}
	tests := []params{
		{
			name: "deleted",
			args: args{
				id: user.UserID,
			},
			errorContains: errNoRows,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			u := postgresql.NewUser()
			_, err := u.Delete(context.Background(), testPool, tt.args.id)
			if err != nil {
				t.Errorf("User.Delete() unexpected error = %v", err)

				return
			}

			_, err = u.ByID(context.Background(), testPool, tt.args.id)
			if err == nil {
				t.Error("wanted error but got nothing", err)

				return
			}
			assert.ErrorContains(t, err, tt.errorContains)
		})
	}
}

type testCase struct {
	name string
	args args
}

// no type parameter to allow direct assertion
type args struct {
	filter any
	fn     any
}

func TestUser_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	userRepo := postgresql.NewUser()

	user, _ := postgresqltestutil.NewRandomUser(t, testPool)

	testCases := []testCase{
		{
			name: "external_id",
			args: args{
				filter: user.ExternalID,
				fn:     (userRepo.ByExternalID),
			},
		},
		{
			name: "email",
			args: args{
				filter: user.Email,
				fn:     (userRepo.ByEmail),
			},
		},
		{
			name: "username",
			args: args{
				filter: user.Username,
				fn:     (userRepo.ByUsername),
			},
		},
		{
			name: "user_id",
			args: args{
				filter: user.UserID,
				fn:     (userRepo.ByID),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		runGenericFilterTests(t, tc, user)
	}
}

// runGenericFilterTests runs queries to find a unique item in db and expects no matches with an inexistent filter.
func runGenericFilterTests(t *testing.T, tc testCase, user *db.User) {
	t.Run(tc.name, func(t *testing.T) {
		t.Parallel()

		var foundUser *db.User
		var err error

		switch fn := tc.args.fn.(type) {
		case func(context.Context, db.DBTX, string) (*db.User, error):
			foundUser, err = fn(context.Background(), testPool, tc.args.filter.(string))
		case func(context.Context, db.DBTX, int) (*db.User, error):
			foundUser, err = fn(context.Background(), testPool, tc.args.filter.(int))
		case func(context.Context, db.DBTX, int64) (*db.User, error):
			foundUser, err = fn(context.Background(), testPool, tc.args.filter.(int64))
		case func(context.Context, db.DBTX, uuid.UUID) (*db.User, error):
			foundUser, err = fn(context.Background(), testPool, tc.args.filter.(uuid.UUID))
		}
		if err != nil {
			t.Fatalf("unexpected error = %v", err)
		}
		assert.Equal(t, foundUser.UserID, user.UserID)
	})

	t.Run(tc.name+"__no_rows_if_does_not_exist", func(t *testing.T) {
		t.Parallel()

		errContains := errNoRows

		var err error

		switch fn := tc.args.fn.(type) {
		case func(context.Context, db.DBTX, string) (*db.User, error):
			filter := "does not exist"
			_, err = fn(context.Background(), testPool, filter)
		case func(context.Context, db.DBTX, int) (*db.User, error):
			filter := int(732745)
			_, err = fn(context.Background(), testPool, filter)
		case func(context.Context, db.DBTX, int64) (*db.User, error):
			filter := int64(732745)
			_, err = fn(context.Background(), testPool, filter)
		case func(context.Context, db.DBTX, uuid.UUID) (*db.User, error):
			filter := uuid.New()
			_, err = fn(context.Background(), testPool, filter)
		}
		if err == nil {
			t.Fatalf("expected error = '%v' but got nothing", errContains)
		}
		assert.Contains(t, err.Error(), errContains)
	})
}

func TestUser_UserAPIKeys(t *testing.T) {
	t.Parallel()

	userRepo := postgresql.NewUser()

	t.Run("correct api key creation", func(t *testing.T) {
		t.Parallel()

		user, _ := postgresqltestutil.NewRandomUser(t, testPool)

		uak, err := userRepo.CreateAPIKey(context.Background(), testPool, user)
		if err != nil {
			t.Fatalf("unexpected error = %v", err)
		}
		assert.NotEmpty(t, uak.APIKey)
		assert.Equal(t, uak.UserID, user.UserID)
		assert.Equal(t, uak.UserAPIKeyID, *user.APIKeyID)
	})

	t.Run("no api key created when user does not exist", func(t *testing.T) {
		t.Parallel()

		errContains := "could not save api key"

		_, err := userRepo.CreateAPIKey(context.Background(), testPool, &db.User{UserID: uuid.New()})
		if err == nil {
			t.Fatalf("expected error = '%v' but got nothing", errContains)
		}
		assert.Contains(t, err.Error(), errContains)
	})

	t.Run("can get user by api key", func(t *testing.T) {
		t.Parallel()

		newUser, _ := postgresqltestutil.NewRandomUser(t, testPool)

		uak, err := userRepo.CreateAPIKey(context.Background(), testPool, newUser)
		if err != nil {
			t.Fatalf("unexpected error = %v", err)
		}

		user, err := userRepo.ByAPIKey(context.Background(), testPool, uak.APIKey)
		if err != nil {
			t.Fatalf("unexpected error = %v", err)
		}

		assert.Equal(t, user.UserID, newUser.UserID)
		assert.Equal(t, *user.APIKeyID, uak.UserAPIKeyID)
	})

	t.Run("cannot get user by api key if key does not exist", func(t *testing.T) {
		t.Parallel()

		errContains := errNoRows

		_, err := userRepo.ByAPIKey(context.Background(), testPool, "missing")
		if err == nil {
			t.Fatalf("expected error = '%v' but got nothing", errContains)
		}
		assert.Contains(t, err.Error(), errContains)
	})

	t.Run("can delete an api key", func(t *testing.T) {
		// TODO
		t.Parallel()
	})
}

func TestUser_Create(t *testing.T) {
	t.Parallel()

	userRepo := postgresql.NewUser()

	type want struct {
		FullName *string
		db.UserCreateParams
	}

	type args struct {
		params db.UserCreateParams
	}

	t.Run("correct user", func(t *testing.T) {
		t.Parallel()

		ucp := postgresqltestutil.RandomUserCreateParams(t)

		want := want{
			FullName:         pointers.New(*ucp.FirstName + " " + *ucp.LastName),
			UserCreateParams: ucp,
		}

		args := args{
			params: ucp,
		}

		got, err := userRepo.Create(context.Background(), testPool, args.params)
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

	// TODO replacing all names with snake case: sed -i -e 's/t\.Run\(.*\)/echo $(echo \1)/e' internal/repos/postgresql/user_test.go
	t.Run("role rank less than zero", func(t *testing.T) {
		t.Parallel()

		ucp := postgresqltestutil.RandomUserCreateParams(t)
		ucp.RoleRank = int16(-1)

		args := args{
			params: ucp,
		}

		errContains := errViolatesCheckConstraint

		_, err := userRepo.Create(context.Background(), testPool, args.params)
		if err == nil {
			t.Fatalf("expected error = '%v' but got nothing", errContains)
		}
		assert.Contains(t, err.Error(), errContains)
	})
}
