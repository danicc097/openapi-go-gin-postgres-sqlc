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
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/crud"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/redis"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/static"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/tracing"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/vault"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"

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
	SpecPath     string
	MovieSvcConn *grpc.ClientConn
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
	if c.MovieSvcConn == nil {
		return fmt.Errorf("no movie service connection provided")
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

	authnSvc := services.Authentication{Logger: conf.Logger, Pool: conf.Pool}
	authzSvc := services.Authorization{Logger: conf.Logger}
	movieClient := v1.NewMovieGenreClient(conf.MovieSvcConn)

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
	userSvc := services.NewUser(postgresql.NewUser(conf.Pool), conf.Logger, conf.Pool, movieClient)

	switch os.Getenv("APP_ENV") {
	case "prod":
		rlMw := newRateLimitMiddleware(conf.Logger, 15)
		vg.Use(rlMw.Limit())
	}

	// TODO REMOVE
	/*
		curl -X 'POST'   'https://localhost:8090/v2/upsert-user'   -H 'accept: application/json'   -H 'Authorization: Bearer fsefse'  -d '{"username":"user","email":"email","role":"admin"}'
	*/
	// https://github.com/xo/xo/blob/master/_examples/booktest/sql/postgres_schema.sql
	// https://github.com/xo/xo/blob/master/_examples/booktest/postgres.go
	// we can call functions directly: presumably should also work for update on mat views, vacuum etc.
	// it can also generate custom queries like sqlc:
	// https://github.com/xo/xo/blob/master/_examples/booktest/sql/postgres_query.sql
	// is AuthorBookResultsByTags
	vg.POST("/upsert-user", func(c *gin.Context) {
		ctx := c.Request.Context()

		// span attribute not inheritable:
		// see https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/14026
		s := newOTELSpan(ctx, "User.CreateUser", trace.WithAttributes(userIDAttribute(c)))
		s.AddEvent("create-user") // filterable with event="create-user"
		defer s.End()

		// tx, err := conf.Pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
		// if err != nil {
		// 	renderErrorResponse(c, "err::", err)
		// 	return
		// }
		// defer tx.Rollback(ctx)

		type UpsertUserRequest struct {
			Username string `json:"username,omitempty" binding:"required"`
			Email    string `json:"email,omitempty" binding:"required"`
			Role     string `json:"role,omitempty" binding:"required"`
		}
		var body UpsertUserRequest
		if err := c.BindJSON(&body); err != nil {
			renderErrorResponse(c, "err::", err)
			return
		}

		// TODO extract to helper
		var role crud.Role
		err = role.UnmarshalText([]byte(body.Role))
		if err != nil {
			renderErrorResponse(c, "err::", err)
			return
		}
		conf.Logger.Sugar().Infof("body is :%#v", body)

		user, err := userSvc.UserByEmail(c, body.Email)
		if err != nil {
			fmt.Printf("failed userSvc.UserByEmail: %s\n", err)
		}
		conf.Logger.Sugar().Infof("user by email: %v", user)
		if user == nil {
			err = userSvc.Register(c, &crud.User{
				Username:  body.Username,
				Email:     body.Email,
				Role:      role,
				FirstName: sql.NullString{String: "firstname", Valid: true},
			})
			if err != nil {
				fmt.Printf("failed userSvc.UserByEmail: %s\n", err)
				renderErrorResponse(c, "user could not be created", err)
				return
			}
			renderResponse(c, "user created", http.StatusOK)
			return
		}
		user.Username = body.Username
		user.Email = body.Email
		user.Role = role
		err = userSvc.Upsert(c, user)
		if err != nil {
			renderErrorResponse(c, "err: ", err)
			return
		}
		// tx.Commit(ctx)
	})

	vg.Use(oasMw.RequestValidatorWithOptions(&options))

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

	movieSvcConn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "movieSvcConn")
	}

	srv, err := NewServer(Config{
		Address:      address,
		Pool:         pool,
		Redis:        rdb,
		Logger:       logger,
		SpecPath:     specPath,
		MovieSvcConn: movieSvcConn,
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
