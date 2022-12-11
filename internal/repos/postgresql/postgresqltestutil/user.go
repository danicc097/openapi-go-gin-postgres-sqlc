package postgresqltestutil

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pointers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRandomUser(t *testing.T, pool *pgxpool.Pool) *db.User {
	t.Helper()

	userRepo := postgresql.NewUser()

	ucp := RandomUserCreateParams(t)

	user, err := userRepo.Create(context.Background(), pool, ucp)
	if err != nil {
		t.Fatalf("unexpected error = %v", err)
	}

	return user
}

func RandomUserCreateParams(t *testing.T) repos.UserCreateParams {
	t.Helper()

	return repos.UserCreateParams{
		Username:   testutil.RandomNameIdentifier(1, "-") + testutil.RandomName(),
		Email:      testutil.RandomEmail(),
		FirstName:  pointers.New(testutil.RandomFirstName()),
		LastName:   pointers.New(testutil.RandomLastName()),
		ExternalID: testutil.RandomString(10),
		Scopes:     []string{"scope1", "scope2"},
		RoleRank:   int16(2),
	}
}
