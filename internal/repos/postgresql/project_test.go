package postgresql_test

// import (
// 	"context"
// 	"testing"

// 	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pointers"
// 	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
// 	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
// 	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// )

// func TestProject_ProjectByIndexedQueries(t *testing.T) {
// 	t.Parallel()

// 	projectRepo := postgresql.NewProject()

// 	ucp := randomUserCreateParams(t)

// 	user, err := projectRepo.Create(context.Background(), testpool, ucp)
// 	if err != nil {
// 		t.Fatalf("unexpected error = %v", err)
// 	}

// 	type args struct {
// 		filter string
// 		fn     func(context.Context, db.DBTX, string) (*db.User, error)
// 	}

// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		{
// 			name: "external_id",
// 			args: args{
// 				filter: user.ExternalID,
// 				fn:     (userRepo.UserByExternalID),
// 			},
// 		},
// 		{
// 			name: "user_id",
// 			args: args{
// 				filter: user.UserID.String(),
// 				fn:     (userRepo.UserByID),
// 			},
// 		},
// 		{
// 			name: "email",
// 			args: args{
// 				filter: user.Email,
// 				fn:     (userRepo.UserByEmail),
// 			},
// 		},
// 		{
// 			name: "username",
// 			args: args{
// 				filter: user.Username,
// 				fn:     (userRepo.UserByUsername),
// 			},
// 		},
// 	}
// 	for _, tc := range tests {
// 		tc := tc
// 		t.Run(tc.name, func(t *testing.T) {
// 			t.Parallel()

// 			foundUser, err := tc.args.fn(context.Background(), testpool, tc.args.filter)
// 			if err != nil {
// 				t.Fatalf("unexpected error = %v", err)
// 			}
// 			assert.Equal(t, foundUser.UserID, user.UserID)
// 		})

// 		t.Run(tc.name+" - no rows when record does not exist", func(t *testing.T) {
// 			t.Parallel()

// 			errContains := errNoRows

// 			// valid as of now for any text and uuid index unless there are specific table checks
// 			filter := uuid.New().String()

// 			_, err := tc.args.fn(context.Background(), testpool, filter)
// 			if err == nil {
// 				t.Fatalf("expected error = '%v' but got nothing", errContains)
// 			}
// 			assert.Contains(t, err.Error(), errContains)
// 		})
// 	}
// }

// func TestProject_UserAPIKeys(t *testing.T) {
// 	t.Parallel()

// 	userRepo := postgresql.NewUser()

// 	t.Run("correct api key creation", func(t *testing.T) {
// 		t.Parallel()

// 		ucp := randomUserCreateParams(t)

// 		user, err := userRepo.Create(context.Background(), testpool, ucp)
// 		if err != nil {
// 			t.Fatalf("unexpected error = %v", err)
// 		}

// 		uak, err := userRepo.CreateAPIKey(context.Background(), testpool, user)
// 		if err != nil {
// 			t.Fatalf("unexpected error = %v", err)
// 		}
// 		assert.NotEmpty(t, uak.APIKey)
// 		assert.Equal(t, uak.UserID, user.UserID)
// 		assert.Equal(t, uak.UserAPIKeyID, *user.APIKeyID)
// 	})

// 	t.Run("no api key created when user does not exist", func(t *testing.T) {
// 		t.Parallel()

// 		errContains := "could not save api key"

// 		_, err := userRepo.CreateAPIKey(context.Background(), testpool, &db.User{UserID: uuid.New()})
// 		if err == nil {
// 			t.Fatalf("expected error = '%v' but got nothing", errContains)
// 		}
// 		assert.Contains(t, err.Error(), errContains)
// 	})

// 	t.Run("can get user by api key", func(t *testing.T) {
// 		t.Parallel()

// 		ucp := randomUserCreateParams(t)

// 		newUser, err := userRepo.Create(context.Background(), testpool, ucp)
// 		if err != nil {
// 			t.Fatalf("unexpected error = %v", err)
// 		}

// 		uak, err := userRepo.CreateAPIKey(context.Background(), testpool, newUser)
// 		if err != nil {
// 			t.Fatalf("unexpected error = %v", err)
// 		}

// 		user, err := userRepo.UserByAPIKey(context.Background(), testpool, uak.APIKey)
// 		if err != nil {
// 			t.Fatalf("unexpected error = %v", err)
// 		}

// 		assert.Equal(t, user.UserID, newUser.UserID)
// 		assert.Equal(t, *user.APIKeyID, uak.UserAPIKeyID)
// 	})

// 	t.Run("cannot get user by api key if key does not exist", func(t *testing.T) {
// 		t.Parallel()

// 		errContains := errNoRows

// 		_, err := userRepo.UserByAPIKey(context.Background(), testpool, "missing")
// 		if err == nil {
// 			t.Fatalf("expected error = '%v' but got nothing", errContains)
// 		}
// 		assert.Contains(t, err.Error(), errContains)
// 	})

// 	t.Run("can delete an api key", func(t *testing.T) {
// 		// TODO
// 	})
// }

// func TestProject_Create(t *testing.T) {
// 	t.Parallel()

// 	userRepo := postgresql.NewUser()

// 	type want struct {
// 		FullName *string
// 		repos.UserCreateParams
// 	}

// 	type args struct {
// 		params repos.UserCreateParams
// 	}

// 	t.Run("correct user", func(t *testing.T) {
// 		t.Parallel()

// 		ucp := randomUserCreateParams(t)

// 		want := want{
// 			FullName:         pointers.New(*ucp.FirstName + " " + *ucp.LastName),
// 			UserCreateParams: ucp,
// 		}

// 		args := args{
// 			params: ucp,
// 		}

// 		got, err := userRepo.Create(context.Background(), testpool, args.params)
// 		if err != nil {
// 			t.Fatalf("unexpected error = %v", err)
// 		}

// 		assert.Equal(t, want.FullName, got.FullName)
// 		assert.Equal(t, want.ExternalID, got.ExternalID)
// 		assert.Equal(t, want.Email, got.Email)
// 		assert.Equal(t, want.Username, got.Username)
// 		assert.Equal(t, want.RoleRank, got.RoleRank)
// 		assert.Equal(t, want.Scopes, got.Scopes)
// 		assert.Equal(t, want.FirstName, got.FirstName)
// 		assert.Equal(t, want.LastName, got.LastName)
// 	})

// 	t.Run("role rank less than zero", func(t *testing.T) {
// 		t.Parallel()

// 		ucp := randomUserCreateParams(t)
// 		ucp.RoleRank = int16(-1)

// 		args := args{
// 			params: ucp,
// 		}

// 		errContains := "violates check constraint"

// 		_, err := userRepo.Create(context.Background(), testpool, args.params)
// 		if err == nil {
// 			t.Fatalf("expected error = '%v' but got nothing", errContains)
// 		}
// 		assert.Contains(t, err.Error(), errContains)
// 	})
// }

// // func randomProjectCreateParams(t *testing.T) repos.ProjectCreateParams {
// // 	t.Helper()

// // 	return repos.ProjectCreateParams{
// // 		Name:        testutil.RandomNameIdentifier(1, "-") + testutil.RandomName(),
// // 		Description: pointers.New(testutil.RandomString(10)),
// // 		ProjectID: , // TODO we still need project repo
// // 	}
// // }
