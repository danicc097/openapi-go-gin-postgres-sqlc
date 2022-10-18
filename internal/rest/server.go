package rest

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	rv8 "github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"

	internaldomain "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/redis"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/static"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/tracing"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/vault"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Config struct {
	// Port to listen to. Use ":0" for a random port.
	Address string
	Pool    *pgxpool.Pool
	Redis   *rv8.Client
	Logger  *zap.Logger
	// SpecPath is the OpenAPI spec filepath.
	SpecPath string
}

func (c *Config) validate() error {
	if c.Address == "" {
		return fmt.Errorf("no server address provided")
	}
	if c.SpecPath == "" {
		return fmt.Errorf("no openapi spec path provided")
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

	return nil
}

type server struct {
	httpsrv     *http.Server
	middlewares []gin.HandlerFunc
}

type serverOption func(*server)

// WithMiddlewares adds the given middlewares before registering the main routers.
func WithMiddlewares(mws []gin.HandlerFunc) serverOption {
	return func(s *server) {
		s.middlewares = mws
	}
}

// NewServer returns a new http server.
func NewServer(conf Config, opts ...serverOption) (*server, error) {
	err := conf.validate()
	if err != nil {
		return nil, err
	}

	srv := &server{}
	for _, o := range opts {
		o(srv)
	}

	router := gin.Default()
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
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

	for _, mw := range srv.middlewares {
		router.Use(mw)
	}

	fsys, _ := fs.Sub(static.SwaggerUI, "swagger-ui")
	vg := router.Group(os.Getenv("API_VERSION"))
	vg.StaticFS("/docs", http.FS(fsys)) // can't validate if not in spec

	schemaBlob, err := os.ReadFile(conf.SpecPath)
	if err != nil {
		return nil, err
	}
	sl := openapi3.NewLoader()

	openapi, err := sl.LoadFromData(schemaBlob)
	if err != nil {
		return nil, err
	}

	if err = openapi.Validate(sl.Context); err != nil {
		return nil, err
	}

	options := OAValidatorOptions{
		Options: openapi3filter.Options{
			ExcludeRequestBody:    false,
			ExcludeResponseBody:   false,
			IncludeResponseStatus: true,
			MultiError:            true,
			AuthenticationFunc:    verifyAuthentication,
		},
		// MultiErrorHandler: func(me openapi3.MultiError) error {
		// 	return fmt.Errorf("multiple errors:  %s", me.Error())
		// },
	}

	oasMw := newOpenapiMiddleware(conf.Logger, openapi)
	vg.Use(oasMw.RequestValidatorWithOptions(&options))

	authnSvc := services.Authentication{Logger: conf.Logger, Pool: conf.Pool}
	authzSvc := services.Authorization{Logger: conf.Logger}
	userSvc := services.NewUser(postgresql.NewUser(conf.Pool), conf.Logger, conf.Pool)

	authMw := newAuthMiddleware(conf.Logger, authnSvc, authzSvc, userSvc)

	NewAdmin(userSvc).
		Register(vg, []gin.HandlerFunc{authMw.EnsureAuthorized(db.RoleAdmin)})

	NewDefault().
		Register(vg, []gin.HandlerFunc{authMw.EnsureAuthenticated(), authMw.EnsureVerified()})

	NewUser(conf.Logger, userSvc, authnSvc, authzSvc).
		Register(vg, []gin.HandlerFunc{})

	conf.Logger.Info("Server started")
	srv.httpsrv = &http.Server{
		Handler:           router,
		Addr:              conf.Address,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       1 * time.Second,
	}
	return srv, nil
}

// Run configures a server and underlying services with the given configuration.
func Run(env, address, specPath string) (<-chan error, error) {
	if err := envvar.Load(env); err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "envvar.Load")
	}

	provider, err := vault.New()
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "internal.NewVaultProvider")
	}

	conf := envvar.New(provider)

	pool, err := postgresql.New(conf)
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "postgresql.New")
	}

	rdb, err := redis.New(conf)
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "redis.New")
	}

	var logger *zap.Logger

	// XXX there's work being done in https://github.com/uptrace/opentelemetry-go-extra/tree/main/otelzap
	switch os.Getenv("APP_ENV") {
	case "prod":
		logger, err = zap.NewProduction()
	default:
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "zap.New")
	}

	tp := tracing.InitTracer()

	ctx := context.Background()

	_, span := tp.Tracer("server-start-tracer").Start(ctx, "server-start")
	defer span.End()

	registerValidators()

	srv, err := NewServer(Config{
		Address:  address,
		Pool:     pool,
		Redis:    rdb,
		Logger:   logger,
		SpecPath: specPath,
	})
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "NewServer")
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
			pool.Close()
			// rmq.Close()
			rdb.Close()
			stop()
			cancel()
			close(errC)
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

		switch env := os.Getenv("APP_ENV"); env {
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
