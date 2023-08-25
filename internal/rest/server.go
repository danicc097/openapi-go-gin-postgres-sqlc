package rest

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	rv8 "github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zitadel/oidc/v2/pkg/client/rp"
	httphelper "github.com/zitadel/oidc/v2/pkg/http"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	internal "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	v1 "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pb/python-ml-app-protos/tfidf/v1"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/redis"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/reposwrappers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/static"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/tracing"
	ginzap "github.com/gin-contrib/zap"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const ValidationErrorSeparator = "$$$$"

type Config struct {
	// Port to listen to. Use ":0" for a random port.
	Address string
	Pool    *pgxpool.Pool
	SQLPool *sql.DB
	Redis   *rv8.Client
	Logger  *zap.SugaredLogger
	// SpecPath is the OpenAPI spec filepath.
	SpecPath               string
	MovieSvcClient         v1.MovieGenreClient
	ScopePolicyPath        string
	RolePolicyPath         string
	MyProviderCallbackPath string
}

// TODO BuildServerConfig with implicit validation instead.
func (c *Config) validate() error {
	if c.Address == "" && os.Getenv("IS_TESTING") == "" {
		return fmt.Errorf("no server address provided")
	}
	if c.SpecPath == "" {
		return fmt.Errorf("no openapi spec path provided")
	}
	if c.ScopePolicyPath == "" {
		return fmt.Errorf("no scope policy path provided")
	}
	if c.RolePolicyPath == "" {
		return fmt.Errorf("no role policy path provided")
	}
	if c.Pool == nil {
		return fmt.Errorf("no Postgres pool provided")
	}
	if c.Redis == nil {
		return fmt.Errorf("no Redis client provided")
	}
	if c.Logger == nil {
		return fmt.Errorf("no logger provided")
	}
	if c.MovieSvcClient == nil {
		return fmt.Errorf("no movie service client provided")
	}

	return nil
}

type Server struct {
	httpsrv     *http.Server
	middlewares []gin.HandlerFunc
}

type ServerOption func(*Server)

// WithMiddlewares adds the given middlewares before registering the main routers.
func WithMiddlewares(mws ...gin.HandlerFunc) ServerOption {
	return func(s *Server) {
		s.middlewares = mws
	}
}

var key = []byte("test1234test1234")

// NewServer returns a new http server.
func NewServer(conf Config, opts ...ServerOption) (*Server, error) {
	if err := conf.validate(); err != nil {
		return nil, fmt.Errorf("server config validation: %w", err)
	}

	srv := &Server{}
	for _, o := range opts {
		o(srv)
	}

	router := gin.Default()
	// Add a ginzap middleware, which:
	// - Logs all requests, like a combined access and error log.
	// - Logs to stdout.
	// - RFC3339 with UTC time format.
	router.Use(ginzap.GinzapWithConfig(conf.Logger.Desugar(), &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		// SkipPaths:  []string{"/no_log"},
	}))
	router.Use(ginzap.RecoveryWithZap(conf.Logger.Desugar(), true))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))

	// no need to set provider and propagator again, will use server global's
	router.Use(otelgin.Middleware(""))
	pprof.Register(router, "dev/pprof")
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	cfg := internal.Config

	for _, mw := range srv.middlewares {
		router.Use(mw)
	}

	fsys, _ := fs.Sub(static.SwaggerUI, "swagger-ui")
	vg := router.Group(cfg.APIVersion)
	// register before spec validation as routes are not in spec
	vg.StaticFS("/docs", http.FS(fsys))
	vg.GET("/docs-redoc", redocHandler)

	// oidc
	keyPath := "" // not used
	scopes := strings.Split(cfg.OIDC.Scopes, " ")

	redirectURI := internal.BuildAPIURL(conf.MyProviderCallbackPath)
	cookieHandler := httphelper.NewCookieHandler(key, key, httphelper.WithUnsecure())

	options := []rp.Option{
		rp.WithCookieHandler(cookieHandler),
		rp.WithVerifierOpts(rp.WithIssuedAtOffset(5 * time.Second)),
	}
	if cfg.OIDC.ClientSecret == "" {
		options = append(options, rp.WithPKCE(cookieHandler))
	}
	if keyPath != "" {
		options = append(options, rp.WithJWTProfile(rp.SignerFromKeyPath(keyPath)))
	}
	provider, err := rp.NewRelyingPartyOIDC(cfg.OIDC.Issuer, cfg.OIDC.ClientID, cfg.OIDC.ClientSecret, redirectURI, scopes, options...)
	if err != nil {
		return nil, fmt.Errorf("error creating provider: %w", err)
	}
	//

	// -- openapi
	openapi, err := ReadOpenAPI(conf.SpecPath)
	if err != nil {
		return nil, err
	}

	oasMw := newOpenapiMiddleware(conf.Logger, openapi)
	oaOptions := createOpenAPIValidatorOptions()
	vg.Use(oasMw.RequestValidatorWithOptions(&oaOptions))

	rlMw := newRateLimitMiddleware(conf.Logger, 25, 10)
	switch cfg.AppEnv {
	case "prod", "e2e":
		vg.Use(rlMw.Limit())
	}
	workitemrepo := reposwrappers.NewWorkItemWithTracing(
		reposwrappers.NewWorkItemWithTimeout(
			postgresql.NewWorkItem(),
			reposwrappers.WorkItemWithTimeoutConfig{},
		),
		postgresql.OtelName,
		nil,
	)
	demoworkitemrepo := reposwrappers.NewDemoWorkItemWithTracing(
		reposwrappers.NewDemoWorkItemWithTimeout(
			postgresql.NewDemoWorkItem(),
			reposwrappers.DemoWorkItemWithTimeoutConfig{},
		),
		postgresql.OtelName,
		nil,
	)
	demotwoworkitemrepo := reposwrappers.NewDemoTwoWorkItemWithTracing(
		reposwrappers.NewDemoTwoWorkItemWithTimeout(
			postgresql.NewDemoTwoWorkItem(),
			reposwrappers.DemoTwoWorkItemWithTimeoutConfig{},
		),
		postgresql.OtelName,
		nil,
	)
	workitemtagrepo := reposwrappers.NewWorkItemTagWithTracing(
		reposwrappers.NewWorkItemTagWithTimeout(
			postgresql.NewWorkItemTag(),
			reposwrappers.WorkItemTagWithTimeoutConfig{},
		),
		postgresql.OtelName,
		nil,
	)
	projectrepo := reposwrappers.NewProjectWithTracing(
		reposwrappers.NewProjectWithTimeout(
			postgresql.NewProject(),
			reposwrappers.ProjectWithTimeoutConfig{},
		),
		postgresql.OtelName,
		nil,
	)
	urepo := reposwrappers.NewUserWithTracing(
		reposwrappers.NewUserWithTimeout(
			postgresql.NewUser(),
			reposwrappers.UserWithTimeoutConfig{},
		),
		postgresql.OtelName,
		nil,
	)
	notifrepo := reposwrappers.NewNotificationWithTracing(
		reposwrappers.NewNotificationWithTimeout(
			postgresql.NewNotification(),
			reposwrappers.NotificationWithTimeoutConfig{},
		),
		postgresql.OtelName,
		nil,
	)

	authzsvc, err := services.NewAuthorization(conf.Logger, conf.ScopePolicyPath, conf.RolePolicyPath)
	if err != nil {
		return nil, fmt.Errorf("NewAuthorization: %w", err)
	}
	usvc := services.NewUser(conf.Logger, urepo, notifrepo, authzsvc)
	workitemsvc := services.NewWorkItem(conf.Logger, workitemtagrepo, workitemrepo, urepo, projectrepo)
	demoworkitemsvc := services.NewDemoWorkItem(conf.Logger, demoworkitemrepo, workitemrepo, urepo, workitemsvc)
	demotwoworkitemsvc := services.NewDemoTwoWorkItem(conf.Logger, demotwoworkitemrepo, workitemrepo, urepo, workitemsvc)
	workitemtagsvc := services.NewWorkItemTag(conf.Logger, workitemtagrepo)
	authnsvc := services.NewAuthentication(conf.Logger, usvc, conf.Pool)
	authmw := newAuthMiddleware(conf.Logger, conf.Pool, authnsvc, authzsvc, usvc)

	handlers := NewHandlers(
		conf.Logger,
		conf.Pool,
		conf.MovieSvcClient,
		conf.SpecPath,
		usvc,
		demoworkitemsvc,
		demotwoworkitemsvc,
		workitemtagsvc,
		authzsvc,
		authnsvc,
		authmw, // middleware needed here since it's generated code
		provider,
	)

	RegisterHandlers(vg, handlers)

	conf.Logger.Info("Server started")

	srv.httpsrv = &http.Server{
		Handler: router,
		Addr:    conf.Address,
		// ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		// WriteTimeout:      10 * time.Second,
		// IdleTimeout:       10 * time.Second,
	}
	return srv, nil
}

// Run configures a server and underlying services with the given configuration.
// NewServer takes its own config as is now.
func Run(env, specPath, rolePolicyPath, scopePolicyPath string) (<-chan error, error) {
	var err error

	if err = envvar.Load(env); err != nil {
		return nil, internal.WrapErrorf(err, models.ErrorCodeUnknown, "envvar.Load")
	}

	cfg := internal.Config

	var logger *zap.Logger
	// XXX there's work being done in https://github.com/uptrace/opentelemetry-go-extra/tree/main/otelzap
	switch cfg.AppEnv {
	case "prod":
		logger, err = zap.NewProduction()
	default:
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		return nil, internal.WrapErrorf(err, models.ErrorCodeUnknown, "zap.New")
	}

	// slogger := logger

	pool, sqlpool, err := postgresql.New(logger.Sugar())
	if err != nil {
		return nil, internal.WrapErrorf(err, models.ErrorCodeUnknown, "postgresql.New")
	}

	rdb, err := redis.New()
	if err != nil {
		return nil, internal.WrapErrorf(err, models.ErrorCodeUnknown, "redis.New")
	}

	tp := tracing.InitOTelTracer()

	ctx := context.Background()

	_, span := tp.Tracer("server-start-tracer").Start(ctx, "server-start")
	defer span.End()

	registerValidators()

	movieSvcConn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return nil, internal.WrapErrorf(err, models.ErrorCodeUnknown, "movieSvcConn")
	}

	srv, err := NewServer(Config{
		Address:                ":" + strings.TrimPrefix(cfg.APIPort, ":"),
		Pool:                   pool,
		SQLPool:                sqlpool,
		Redis:                  rdb,
		Logger:                 logger.Sugar(),
		SpecPath:               specPath,
		ScopePolicyPath:        scopePolicyPath,
		RolePolicyPath:         rolePolicyPath,
		MovieSvcClient:         v1.NewMovieGenreClient(movieSvcConn),
		MyProviderCallbackPath: "/auth/myprovider/callback",
	})
	if err != nil {
		return nil, internal.WrapErrorf(err, models.ErrorCodeUnknown, "NewServer")
	}

	errC := make(chan error, 1)

	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-ctx.Done()

		logger.Info("Shutdown signal received")

		ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		// any action on shutdown must be deferred here and not in the main block
		defer func() {
			_ = logger.Sync()

			tp.Shutdown(context.Background())
			movieSvcConn.Close()
			pool.Close()
			// rmq.Close()
			rdb.Close()
			stop()
			cancel()
			close(errC)

			// TODO close SSE channels
		}()

		srv.httpsrv.SetKeepAlivesEnabled(false)

		if err := srv.httpsrv.Shutdown(ctxTimeout); err != nil { //nolint: contextcheck
			errC <- err
		}

		logger.Info("Shutdown completed")
	}()

	go func() {
		logger.Info("Listening and serving", zap.String("address", cfg.APIPort))

		// "ListenAndServe always returns a non-nil error. After Shutdown or Close, the returned error is
		// ErrServerClosed."
		var err error

		switch cfg.AppEnv {
		case "dev", "ci":
			// err = srv.httpsrv.ListenAndServe()
			err = srv.httpsrv.ListenAndServeTLS("certificates/localhost.pem", "certificates/localhost-key.pem")
		case "prod", "e2e":
			err = srv.httpsrv.ListenAndServe()
		default:
			err = fmt.Errorf("unknown APP_ENV: %s", env)
		}

		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errC <- err
		}
	}()

	return errC, nil
}

func createOpenAPIValidatorOptions() OAValidatorOptions {
	// TODO if env != prod so that no values are shown in `details` (secrets, passwords)
	// openapi3.SchemaErrorDetailsDisabled = true

	oafilterOpts := openapi3filter.Options{
		ExcludeRequestBody:    false,
		ExcludeResponseBody:   false,
		IncludeResponseStatus: false,
		MultiError:            true,
		AuthenticationFunc:    verifyAuthentication,
	}

	oafilterOpts.WithCustomSchemaErrorFunc(CustomSchemaErrorFunc)
	oaOptions := OAValidatorOptions{
		ValidateResponse: os.Getenv("IS_TESTING") != "",
		Options:          oafilterOpts,
	}

	return oaOptions
}

func CustomSchemaErrorFunc(err *openapi3.SchemaError) string {
	value, _ := json.Marshal(err.Value)

	var schema map[string]any
	s, _ := err.Schema.MarshalJSON()
	_ = json.Unmarshal(s, &schema)

	ve := &models.ValidationError{
		Loc: err.JSONPointer(),
		Msg: err.Reason,
		Detail: struct {
			Schema map[string]interface{} `json:"schema"`
			Value  string                 `json:"value"`
		}{
			Schema: schema,
			Value:  string(value),
		},
	}

	b, _ := json.Marshal(ve)

	return ValidationErrorSeparator + string(b)
}

func redocHandler(c *gin.Context) {
	htmlString := fmt.Sprintf(`<!DOCTYPE html>
	<html>
		<head>
			<title>Redoc</title>
			<meta charset="utf-8"/>
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet">

			<style>
				body {
					margin: 0;
					padding: 0;
				}
			</style>
		</head>
		<body>
			<redoc spec-url='%s'></redoc>
			<script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js"> </script>
		</body>
	</html>`, internal.BuildAPIURL("openapi.yaml"))

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlString))
}
