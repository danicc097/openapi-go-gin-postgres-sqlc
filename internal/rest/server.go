package rest

import (
	"context"
	"database/sql"
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
	"github.com/gin-gonic/gin"
	rv8 "github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zitadel/oidc/pkg/client/rp"
	httphelper "github.com/zitadel/oidc/pkg/http"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	internal "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
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

type Config struct {
	// Port to listen to. Use ":0" for a random port.
	Address string
	Pool    *pgxpool.Pool
	SQLPool *sql.DB
	Redis   *rv8.Client
	Logger  *zap.Logger
	// SpecPath is the OpenAPI spec filepath.
	SpecPath               string
	MovieSvcClient         v1.MovieGenreClient
	ScopePolicyPath        string
	RolePolicyPath         string
	MyProviderCallbackPath string
}

// TODO BuildServerConfig with implicit validation instead
func (c *Config) validate() error {
	if c.Address == "" {
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

type server struct {
	httpsrv     *http.Server
	middlewares []gin.HandlerFunc
}

type ServerOption func(*server)

// WithMiddlewares adds the given middlewares before registering the main routers.
func WithMiddlewares(mws ...gin.HandlerFunc) ServerOption {
	return func(s *server) {
		s.middlewares = mws
	}
}

var key = []byte("test1234test1234")

// NewServer returns a new http server.
func NewServer(conf Config, opts ...ServerOption) (*server, error) {
	if err := conf.validate(); err != nil {
		return nil, fmt.Errorf("server config validation: %w", err)
	}

	srv := &server{}
	for _, o := range opts {
		o(srv)
	}

	router := gin.Default()
	// Add a ginzap middleware, which:
	// - Logs all requests, like a combined access and error log.
	// - Logs to stdout.
	// - RFC3339 with UTC time format.
	router.Use(ginzap.GinzapWithConfig(conf.Logger, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		// SkipPaths:  []string{"/no_log"},
	}))
	router.Use(ginzap.RecoveryWithZap(conf.Logger, true))
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
	// pprof.Register(router, "dev/pprof")
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	cfg := internal.Config()

	for _, mw := range srv.middlewares {
		router.Use(mw)
	}
	fsys, _ := fs.Sub(static.SwaggerUI, "swagger-ui")
	vg := router.Group(cfg.APIVersion)
	vg.StaticFS("/docs", http.FS(fsys)) // can't validate if not in spec

	// oidc
	clientID := cfg.OIDC.ClientID
	clientSecret := cfg.OIDC.ClientSecret
	keyPath := "" // not used
	issuer := cfg.OIDC.Issuer
	scopes := strings.Split(cfg.OIDC.Scopes, " ")

	redirectURI := internal.BuildAPIURL(conf.MyProviderCallbackPath)
	cookieHandler := httphelper.NewCookieHandler(key, key, httphelper.WithUnsecure())

	options := []rp.Option{
		rp.WithCookieHandler(cookieHandler),
		rp.WithVerifierOpts(rp.WithIssuedAtOffset(5 * time.Second)),
	}
	if clientSecret == "" {
		options = append(options, rp.WithPKCE(cookieHandler))
	}
	if keyPath != "" {
		options = append(options, rp.WithJWTProfile(rp.SignerFromKeyPath(keyPath)))
	}
	provider, err := rp.NewRelyingPartyOIDC(issuer, clientID, clientSecret, redirectURI, scopes, options...)
	if err != nil {
		return nil, fmt.Errorf("error creating provider %s", err)
	}
	//

	// -- openapi
	oafilterOpts := openapi3filter.Options{
		ExcludeRequestBody:    false,
		ExcludeResponseBody:   false,
		IncludeResponseStatus: false,
		MultiError:            true,
		AuthenticationFunc:    verifyAuthentication,
	}
	oafilterOpts.WithCustomSchemaErrorFunc(func(err *openapi3.SchemaError) string {
		return fmt.Sprintf("%s: %s", err.SchemaField, err.Reason)
	})
	oaOptions := OAValidatorOptions{
		ValidateResponse: true,
		Options:          oafilterOpts,
		// MultiErrorHandler: func(me openapi3.MultiError) error {
		// 	return fmt.Errorf("multiple errors:  %s", me.Error())
		// },
	}

	openapi, err := ReadOpenAPI(conf.SpecPath)
	if err != nil {
		return nil, err
	}

	oasMw := newOpenapiMiddleware(conf.Logger, openapi)

	rlMw := newRateLimitMiddleware(conf.Logger, 25, 10)
	switch cfg.AppEnv {
	case "prod":
		vg.Use(rlMw.Limit())
	}
	vg.Use(oasMw.RequestValidatorWithOptions(&oaOptions))

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
	authnsvc := services.NewAuthentication(conf.Logger, usvc, conf.Pool)
	authmw := newAuthMiddleware(conf.Logger, conf.Pool, authnsvc, authzsvc, usvc)

	handlers := NewHandlers(conf.Logger, conf.Pool, conf.MovieSvcClient, usvc, authzsvc, authnsvc, authmw, provider)

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
// TODO should take in AppConfig.
// RunTestServer also takes AppConfig
// NewServer takes its own config as is now
func Run(env, address, specPath, rolePolicyPath, scopePolicyPath string) (<-chan error, error) {
	var err error

	if err = envvar.Load(env); err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "envvar.Load")
	}

	conf := envvar.New()

	cfg := internal.Config()

	var logger *zap.Logger
	// XXX there's work being done in https://github.com/uptrace/opentelemetry-go-extra/tree/main/otelzap
	switch cfg.AppEnv {
	case "prod":
		logger, err = zap.NewProduction()
	default:
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "zap.New")
	}

	pool, sqlpool, err := postgresql.New(logger)
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "postgresql.New")
	}

	rdb, err := redis.New(conf)
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "redis.New")
	}

	tp := tracing.InitTracer()

	ctx := context.Background()

	_, span := tp.Tracer("server-start-tracer").Start(ctx, "server-start")
	defer span.End()

	registerValidators()

	movieSvcConn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "movieSvcConn")
	}

	srv, err := NewServer(Config{
		Address:                address,
		Pool:                   pool,
		SQLPool:                sqlpool,
		Redis:                  rdb,
		Logger:                 logger,
		SpecPath:               specPath,
		ScopePolicyPath:        scopePolicyPath,
		RolePolicyPath:         rolePolicyPath,
		MovieSvcClient:         v1.NewMovieGenreClient(movieSvcConn),
		MyProviderCallbackPath: "/auth/myprovider/callback",
	})
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "NewServer")
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
		logger.Info("Listening and serving", zap.String("address", address))

		// "ListenAndServe always returns a non-nil error. After Shutdown or Close, the returned error is
		// ErrServerClosed."
		var err error

		switch cfg.AppEnv {
		case "dev", "ci":
			// err = srv.httpsrv.ListenAndServe()
			err = srv.httpsrv.ListenAndServeTLS("certificates/localhost.pem", "certificates/localhost-key.pem")
		case "prod":
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
