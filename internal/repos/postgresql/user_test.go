package postgresql_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
					RoleRank: pointers.New(10),
					Scopes:   &models.Scopes{"test", "test", "test"},
				},
			},
			want: func() *db.User {
				u := *user
				u.RoleRank = 10
				u.Scopes = models.Scopes{"test"}

				return &u
			}(),
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			u := postgresql.NewUser()
			got, err := u.Update(context.Background(), testPool, tc.args.id, &tc.args.params)
			if (err != nil) != tc.wantErr {
				t.Errorf("User.Update() error = %v, wantErr %v", err, tc.wantErr)

				return
			}

			got.UpdatedAt = user.UpdatedAt // ignore

			// NOTE: this should not fail when running notification tests (from this package) in transaction
			// // since we run tests in parallel, notification fan out effects changes on all users
			got.HasGlobalNotifications = user.HasGlobalNotifications     // ignore
			got.HasPersonalNotifications = user.HasPersonalNotifications // ignore

			assert.Equal(t, tc.want, got)
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
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			u := postgresql.NewUser()
			_, err := u.Delete(context.Background(), testPool, tc.args.id)
			if err != nil {
				t.Errorf("User.Delete() unexpected error = %v", err)

				return
			}

			_, err = u.ByID(context.Background(), testPool, tc.args.id)
			if err == nil {
				t.Error("wanted error but got nothing", err)

				return
			}
			assert.ErrorContains(t, err, tc.errorContains)
		})
	}
}

func TestUser_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	userRepo := postgresql.NewUser()

	user, _ := postgresqltestutil.NewRandomUser(t, testPool)

	testCases := []filterTestCase{
		{
			name: "external_id",
			args: filterTestCaseArgs{
				filter: user.ExternalID,
				fn:     reflect.ValueOf(userRepo.ByExternalID),
			},
		},
		{
			name: "email",
			args: filterTestCaseArgs{
				filter: user.Email,
				fn:     reflect.ValueOf(userRepo.ByEmail),
			},
		},
		{
			name: "username",
			args: filterTestCaseArgs{
				filter: user.Username,
				fn:     reflect.ValueOf(userRepo.ByUsername),
			},
		},
		{
			name: "user_id",
			args: filterTestCaseArgs{
				filter: user.UserID,
				fn:     reflect.ValueOf(userRepo.ByID),
			},
		},
		/**
		 * TODO: ByTeam and ByProject tests which return []Entity
		 * will need custom comparerFunc ()
		 */
	}

	for _, tc := range testCases {
		runGenericUniqueFilterTests(t, tc, func(t *testing.T, foundEntity any) {
			gotIDField := reflect.ValueOf(foundEntity).Elem().FieldByName("UserID").Interface()
			wantIDField := reflect.ValueOf(user).Elem().FieldByName("UserID").Interface()
			assert.Equal(t, gotIDField, wantIDField)
		})
	}
}

func TestUser_UserAPIKeys(t *testing.T) {
	t.Parallel()

	userRepo := postgresql.NewUser()

	t.Run("correct_api_key_creation", func(t *testing.T) {
		t.Parallel()

		user, _ := postgresqltestutil.NewRandomUser(t, testPool)

		uak, err := userRepo.CreateAPIKey(context.Background(), testPool, user)
		require.NoError(t, err)
		assert.NotEmpty(t, uak.APIKey)
		assert.Equal(t, uak.UserID, user.UserID)
		assert.Equal(t, uak.UserAPIKeyID, *user.APIKeyID)
	})

	t.Run("no_api_key_created_when_user_does_not_exist", func(t *testing.T) {
		t.Parallel()

		errContains := "could not save api key"

		_, err := userRepo.CreateAPIKey(context.Background(), testPool, &db.User{UserID: uuid.New()})
		if err == nil {
			t.Fatalf("expected error = '%v' but got nothing", errContains)
		}
		assert.ErrorContains(t, err, errContains)
	})

	t.Run("can_get_user_by_api_key", func(t *testing.T) {
		t.Parallel()

		newUser, _ := postgresqltestutil.NewRandomUser(t, testPool)

		uak, err := userRepo.CreateAPIKey(context.Background(), testPool, newUser)
		require.NoError(t, err)

		user, err := userRepo.ByAPIKey(context.Background(), testPool, uak.APIKey)
		require.NoError(t, err)

		assert.Equal(t, user.UserID, newUser.UserID)
		assert.Equal(t, *user.APIKeyID, uak.UserAPIKeyID)
	})

	t.Run("cannot_get_user_by_api_key_if_key_does_not_exist", func(t *testing.T) {
		t.Parallel()

		errContains := errNoRows

		_, err := userRepo.ByAPIKey(context.Background(), testPool, "missing")
		if err == nil {
			t.Fatalf("expected error = '%v' but got nothing", errContains)
		}
		assert.ErrorContains(t, err, errContains)
	})

	t.Run("can_delete_an_api_key", func(t *testing.T) {
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

	t.Run("correct_user", func(t *testing.T) {
		t.Parallel()

		ucp := postgresqltestutil.RandomUserCreateParams(t)

		want := want{
			FullName:         pointers.New(*ucp.FirstName + " " + *ucp.LastName),
			UserCreateParams: *ucp,
		}

		args := args{
			params: *ucp,
		}

		got, err := userRepo.Create(context.Background(), testPool, &args.params)
		require.NoError(t, err)

		assert.Equal(t, want.FullName, got.FullName)
		assert.Equal(t, want.ExternalID, got.ExternalID)
		assert.Equal(t, want.Email, got.Email)
		assert.Equal(t, want.Username, got.Username)
		assert.Equal(t, want.RoleRank, got.RoleRank)
		assert.Equal(t, want.Scopes, got.Scopes)
		assert.Equal(t, want.FirstName, got.FirstName)
		assert.Equal(t, want.LastName, got.LastName)
	})

	t.Run("role_rank_less_than_zero", func(t *testing.T) {
		t.Parallel()

		ucp := postgresqltestutil.RandomUserCreateParams(t)
		ucp.RoleRank = -1

		args := args{
			params: *ucp,
		}

		errContains := errViolatesCheckConstraint

		_, err := userRepo.Create(context.Background(), testPool, &args.params)
		if err == nil {
			t.Fatalf("expected error = '%v' but got nothing", errContains)
		}
		assert.ErrorContains(t, err, errContains)
	})
}
