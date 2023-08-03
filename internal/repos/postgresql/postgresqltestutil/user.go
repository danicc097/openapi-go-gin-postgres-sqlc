package postgresqltestutil

import (
	"context"
	"testing"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/reposwrappers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func NewRandomUser(t *testing.T, pool *pgxpool.Pool) (*db.User, error) {
	t.Helper()

	userRepo := reposwrappers.NewUserWithRetry(postgresql.NewUser(), 10, 65*time.Millisecond)

	ucp := RandomUserCreateParams(t)

	user, err := userRepo.Create(context.Background(), pool, ucp)
	require.NoError(t, err, "failed to create random entity")

	return user, nil
}

func RandomUserCreateParams(t *testing.T) *db.UserCreateParams {
	t.Helper()

	return &db.UserCreateParams{
		Username:   testutil.RandomNameIdentifier(1, "-") + testutil.RandomName(),
		Email:      testutil.RandomEmail(),
		FirstName:  pointers.New(testutil.RandomFirstName()),
		LastName:   pointers.New(testutil.RandomLastName()),
		ExternalID: testutil.RandomString(10),
		Scopes:     models.Scopes{"scope1", "scope2"},
		RoleRank:   testutil.RandomInt(2, 4),
	}
}
