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
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	internaldomain "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	v1 "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pb/python-ml-app-protos/tfidf/v1"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/redis"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/static"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/tracing"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/vault"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

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
	SpecPath       string
	MovieSvcClient v1.MovieGenreClient
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
	if c.MovieSvcClient == nil {
		return fmt.Errorf("no movie service client provided")
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

	// TODO need to instantiate the repo with the conn/transaction already.
	// hence we need to create a new service for every new request
	// with a pool conn (already do this in py impl), else we
	// would need to have helpers to change db dBTX in the repo itself, and the repo
	// should not be concerned with this.
	// TLDR refactor: all services need to be instantiated in the handler itself,
	// beginning a transaction every time and sharing it for most use cases.
	// handler structs receive everything necessary to construct all services.
	// then all services share the same conn, _ := conf.Pool.Acquire() initialized in the handler
	// (dont want to handle transactions or any committing in services, its up to caller - i.e. handlers, cli...)
	// unless we have a reason not to (e.g. conn is not concurrency safe)
	// finally we must call conn.Close and Rollback. we commit along the way

	// use a direct pool connection if a query cannot be run in a transaction.
	// IMPORTANT: read https://groups.google.com/g/golang-nuts/c/y8uLMofW2-E and then this comment tree
	// https://medium.com/@florian_445/thanks-for-your-answer-6d03846860fa, then the article
	// itself.
	// takeaways:
	// - we start any transactions in each handler method. Each service method calls the necessary
	// repos OR services. Services are built in each handler, else imagine
	//   the need for usvc := users.New(nsvc) and nsvc := notifications.New(usvc) at the same time
	//   if we create the service inside NewXX() the problem is gone as long as services
	//   remain in the same package, which they should.
	// - repos must not be concerned with transaction details
	// - also note services dont necessarily need an equivalently named repository or viceversa.

	switch os.Getenv("APP_ENV") {
	case "prod":
		rlMw := newRateLimitMiddleware(conf.Logger, 15)
		vg.Use(rlMw.Limit())
	}

	usvc := services.NewUser(postgresql.NewUser(), conf.Logger)
	authzsvc := services.NewAuthorization(conf.Logger)
	authnsvc := services.NewAuthentication(conf.Logger, usvc)

	vg.Use(oasMw.RequestValidatorWithOptions(&options))

	authmw := newAuthMiddleware(conf.Logger, conf.Pool, authnsvc, authzsvc, usvc)

	NewAdmin(conf.Logger, conf.Pool).
		Register(vg, []gin.HandlerFunc{authmw.EnsureAuthorized(db.RoleAdmin)})

	NewDefault().
		Register(vg, []gin.HandlerFunc{authmw.EnsureAuthenticated()})

	NewUser(conf.Logger, conf.Pool, conf.MovieSvcClient, usvc, authmw).
		Register(vg, []gin.HandlerFunc{authmw.EnsureAuthenticated()})

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

	movieSvcConn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "movieSvcConn")
	}

	srv, err := NewServer(Config{
		Address:        address,
		Pool:           pool,
		Redis:          rdb,
		Logger:         logger,
		SpecPath:       specPath,
		MovieSvcClient: v1.NewMovieGenreClient(movieSvcConn),
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
			movieSvcConn.Close()
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
