package rest

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	internaldomain "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pb/python-ml-app-protos/tfidf/v1/v1testing"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest/resttestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	redis "github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

var (
	testPool    *pgxpool.Pool
	testSQLPool *sql.DB
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

func runTestServer(t *testing.T, testPool *pgxpool.Pool, middlewares []gin.HandlerFunc) (*http.Server, error) {
	ctx := context.Background()

	if err := envvar.Load(fmt.Sprintf("../../.env.%s", os.Getenv("APP_ENV"))); err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "envvar.Load")
	}

	// provider, err := vault.New()
	// if err != nil {
	// 	return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "internal.NewVaultProvider")
	// }

	// conf := envvar.New(provider)

	s := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	// sanity check
	rdb.Set(ctx, "foo", "bar", 10*time.Second)
	assert.Equal(t, "bar", rdb.Get(ctx, "foo").Val())

	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "internal.zapNew")
	}

	_, err = openapi3.NewLoader().LoadFromFile("../../openapi.yaml")
	if err != nil {
		panic(err)
	}

	srv, err := NewServer(Config{
		Address:         ":0", // random next available for each test server
		Pool:            testPool,
		Redis:           rdb,
		Logger:          logger,
		SpecPath:        "../../openapi.yaml",
		MovieSvcClient:  &v1testing.FakeMovieGenreClient{},
		ScopePolicyPath: "../../scopes.json",
		RolePolicyPath:  "../../roles.json",
	}, WithMiddlewares(middlewares...))
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "New")
	}

	return srv.httpsrv, nil
}

func newTestFixtureFactory(t *testing.T) *resttestutil.FixtureFactory {
	logger := zaptest.NewLogger(t)
	authzsvc, err := services.NewAuthorization(logger, "../../scopes.json", "../../roles.json")
	if err != nil {
		t.Fatalf("services.NewAuthorization: %v", err)
	}
	retryCount := 1
	retryInterval := 1 * time.Second
	usvc := services.NewUser(
		logger,
		repos.NewUserWithTracing(
			repos.NewUserWithRetry(
				repos.NewUserWithTimeout(
					postgresql.NewUser(), repos.UserWithTimeoutConfig{CreateTimeout: 10 * time.Second}),
				retryCount, retryInterval),
			postgresql.OtelName, nil),
		postgresql.NewNotification(),
		authzsvc,
	)
	authnsvc := services.NewAuthentication(logger, usvc, testPool)

	ff := resttestutil.NewFixtureFactory(usvc, testPool, authnsvc, authzsvc)
	return ff
}
