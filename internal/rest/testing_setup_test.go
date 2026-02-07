package rest_test

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/assert"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pb/python-ml-app-protos/tfidf/v1/v1testing"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest/resttesting"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
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
	gin.SetMode(gin.TestMode)

	testutil.Setup()
	// call flag.Parse() here if TestMain uses flags
	var err error

	internal.Config.RolePolicyPath = "../../roles.json"
	internal.Config.ScopePolicyPath = "../../scopes.json"

	testPool, testSQLPool, err = testutil.NewDB(testutil.WithMigrations())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create testPool: %s\n", err)
		os.Exit(1)
	}
	defer testPool.Close()

	return m.Run()
}

type testServer struct {
	server       *http.Server
	client       *resttesting.ClientWithResponses
	tp           *sdktrace.TracerProvider
	spanRecorder *tracetest.SpanRecorder
	event        *rest.EventServer
}

func (s *testServer) setupCleanup(t *testing.T) {
	t.Cleanup(func() {
		s.server.Close()
		s.tp.Shutdown(t.Context())
	})
}

// runTestServer returns a test server and client.
// We will require different middlewares depending on the test case, so a shared global instance
// is not possible.
func runTestServer(t *testing.T, ctx context.Context, testPool *pgxpool.Pool, middlewares ...gin.HandlerFunc) (*testServer, error) {
	t.Helper()

	spanRecorder := tracetest.NewSpanRecorder()
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(tracetest.NewInMemoryExporter()),
		sdktrace.WithSpanProcessor(spanRecorder),
	)

	s := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	// sanity check
	rdb.Set(ctx, "foo", "bar", 10*time.Second)
	assert.Equal(t, "bar", rdb.Get(ctx, "foo").Val())

	logger := testutil.NewLogger(t)

	_, err := openapi3.NewLoader().LoadFromFile("../../openapi.yaml")
	if err != nil {
		panic(fmt.Sprintf("openapi3.NewLoader: %v", err))
	}

	srv, err := rest.NewServer(ctx, rest.Config{
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

	client, err := resttesting.NewTestClient(MustConstructInternalPath(""), srv.Httpsrv.Handler)
	if err != nil {
		return nil, internal.WrapErrorf(err, models.ErrorCodeUnknown, "NewTestClient")
	}

	return &testServer{
		server:       srv.Httpsrv,
		event:        srv.Event,
		client:       client,
		tp:           tp,
		spanRecorder: spanRecorder,
	}, nil
}
