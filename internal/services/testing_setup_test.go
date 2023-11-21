package services_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/reposwrappers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	testPool    *pgxpool.Pool
	testSQLPool *sql.DB // for jet, use .Sql() to use pgx directly
)

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	testutil.Setup()

	// call flag.Parse() here if TestMain uses flags
	var err error

	testPool, testSQLPool, err = testutil.NewDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create testPool: %s\n", err)
		os.Exit(1)
	}
	defer testPool.Close()

	return m.Run()
}

func newTestFixtureFactory(t *testing.T) *servicetestutil.FixtureFactory {
	logger := zaptest.NewLogger(t).Sugar()
	authzsvc, err := services.NewAuthorization(logger, "../../scopes.json", "../../roles.json")
	if err != nil {
		t.Fatalf("services.NewAuthorization: %v", err)
	}
	usvc := services.NewUser(
		logger,
		reposwrappers.NewUserWithTracing(
			reposwrappers.NewUserWithTimeout(
				reposwrappers.NewUserWithRetry(postgresql.NewUser(), 10, 65*time.Millisecond), reposwrappers.UserWithTimeoutConfig{}),
			postgresql.OtelName, nil),
		reposwrappers.NewNotificationWithTracing(
			reposwrappers.NewNotificationWithTimeout(
				postgresql.NewNotification(), reposwrappers.NotificationWithTimeoutConfig{}),
			postgresql.OtelName, nil),
		authzsvc,
	)
	authnsvc := services.NewAuthentication(logger, usvc, testPool)

	ff := servicetestutil.NewFixtureFactory(usvc, testPool, authnsvc, authzsvc)
	return ff
}

func newTestAuthService(t *testing.T) *services.Authorization {
	logger := zaptest.NewLogger(t).Sugar()

	authzsvc, err := services.NewAuthorization(logger, "../../scopes.json", "../../roles.json")
	require.NoError(t, err, "newTestAuthService")

	return authzsvc
}
