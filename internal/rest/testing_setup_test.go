package rest_test

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pb/python-ml-app-protos/tfidf/v1/v1testing"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	redis "github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"

	_ "embed"
)

var (
	testPool    *pgxpool.Pool
	testSQLPool *sql.DB // for jet, use .Sql() to use pgx directly
)

//go:embed testdata/test_spec.yaml
var testSchema []byte

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	testutil.Setup()

	// call flag.Parse() here if TestMain uses flags
	var err error

	internal.Config.RolePolicyPath = "../../roles.json"
	internal.Config.ScopePolicyPath = "../../scopes.json"

	testPool, testSQLPool, err = testutil.NewDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create testPool: %s\n", err)
		os.Exit(1)
	}
	defer testPool.Close()

	return m.Run()
}

type testServer struct {
	server *http.Server
	client *ClientWithResponses
}

func (s *testServer) setupCleanup(t *testing.T) {
	t.Cleanup(func() {
		s.server.Close()
	})
}

// runTestServer returns a test server and client.
// We will require different middlewares depending on the test case, so a shared global instance
// is not possible.
func runTestServer(t *testing.T, testPool *pgxpool.Pool, middlewares ...gin.HandlerFunc) (*testServer, error) {
	t.Helper()

	ctx := context.Background()

	// race. also already done in testutils setup.
	// if err := envvar.Load(fmt.Sprintf("../../.env.%s", os.Getenv("APP_ENV"))); err != nil {
	// 	return nil, internal.WrapErrorf(err, models.ErrorCodeUnknown, "envvar.Load")
	// }

	// provider, err := vault.New()
	// if err != nil {
	// 	return nil, internal.WrapErrorf(err, models.ErrorCodeUnknown, "internal.NewVaultProvider")
	// }

	// conf := envvar.New(provider)

	s := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	// sanity check
	rdb.Set(ctx, "foo", "bar", 10*time.Second)
	assert.Equal(t, "bar", rdb.Get(ctx, "foo").Val())

	logger := zaptest.NewLogger(t).Sugar()

	_, err := openapi3.NewLoader().LoadFromFile("../../openapi.yaml")
	if err != nil {
		panic(fmt.Sprintf("openapi3.NewLoader: %v", err))
	}

	srv, err := rest.NewServer(rest.Config{
		// not necessary when using ServeHTTP. Won't actually listen.
		// Address:         ":0", // random next available for each test server
		Pool:           testPool,
		Redis:          rdb,
		Logger:         logger,
		SpecPath:       "../../openapi.yaml",
		MovieSvcClient: &v1testing.FakeMovieGenreClient{},
	}, rest.WithMiddlewares(middlewares...))
	if err != nil {
		return nil, internal.WrapErrorf(err, models.ErrorCodeUnknown, "NewServer")
	}

	client, err := NewTestClient(MustConstructInternalPath(""), srv.Httpsrv.Handler)
	if err != nil {
		return nil, internal.WrapErrorf(err, models.ErrorCodeUnknown, "NewTestClient")
	}

	return &testServer{server: srv.Httpsrv, client: client}, nil
}
