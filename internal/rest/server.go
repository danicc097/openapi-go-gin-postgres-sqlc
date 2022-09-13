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
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/vault"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Config struct {
	Address string
	Pool    *pgxpool.Pool
	Redis   *rv8.Client
	Metrics http.Handler
	Logger  *zap.Logger
	// SpecPath is the OpenAPI spec filepath.
	SpecPath string
}

// NewServer returns a new http server.
func NewServer(conf Config) (*http.Server, error) {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("alphanumspace", Alphanumspace)
	}

	router.Use(gin.Recovery())
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
	err = openapi.Validate(sl.Context)
	if err != nil {
		return nil, err
	}

	options := OAValidatorOptions{
		Options: openapi3filter.Options{
			ExcludeRequestBody:    false,
			ExcludeResponseBody:   false,
			IncludeResponseStatus: true,
			MultiError:            true,
		},
		// MultiErrorHandler: func(me openapi3.MultiError) error {
		// 	return fmt.Errorf("multiple errors:  %s", me.Error())
		// },
	}

	vg.Use(OapiRequestValidatorWithOptions(openapi, &options))

	authnSvc := services.Authentication{Logger: conf.Logger, Pool: conf.Pool}
	authzSvc := services.Authorization{Logger: conf.Logger}
	fakeSvc := services.Fake{Logger: conf.Logger, Pool: conf.Pool}
	petSvc := services.Pet{Logger: conf.Logger, Pool: conf.Pool}
	storeSvc := services.Store{Logger: conf.Logger, Pool: conf.Pool}
	userSvc := services.NewUser(postgresql.NewUser(conf.Pool), conf.Logger, conf.Pool)

	authMw := NewAuthMw(conf.Logger, authnSvc, authzSvc, userSvc)

	NewAdmin(userSvc).
		Register(vg, []gin.HandlerFunc{authMw.EnsureAuthorized(db.RoleAdmin)})

	NewDefault().
		Register(vg, []gin.HandlerFunc{authMw.EnsureAuthenticated(), authMw.EnsureVerified()})

	NewFake(fakeSvc).
		Register(vg, []gin.HandlerFunc{})

	NewPet(petSvc).
		Register(vg, []gin.HandlerFunc{})

	NewStore(storeSvc).
		Register(vg, []gin.HandlerFunc{})

	NewUser(conf.Logger, userSvc, authnSvc, authzSvc).
		Register(vg, []gin.HandlerFunc{})
	// TODO /admin with authMw.EnsureAuthorized() in group

	conf.Logger.Info("Server started")

	return &http.Server{
		Handler:           router,
		Addr:              conf.Address,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       1 * time.Second,
	}, nil
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
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "internal.NewPostgreSQL")
	}

	rdb, err := redis.New(conf)
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "internal.NewRedis")
	}

	var logger *zap.Logger

	switch os.Getenv("APP_ENV") {
	case "prod":
		logger, err = zap.NewProduction()
	default:
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "internal.zapNew")
	}

	srv, err := NewServer(Config{
		Address:  address,
		Pool:     pool,
		Redis:    rdb,
		Logger:   logger,
		SpecPath: specPath,
	})
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "New")
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

		defer func() {
			_ = logger.Sync()

			pool.Close()
			// rmq.Close()
			rdb.Close()
			stop()
			cancel()
			close(errC)
		}()

		srv.SetKeepAlivesEnabled(false)

		if err := srv.Shutdown(ctxTimeout); err != nil { //nolint: contextcheck
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
			err = srv.ListenAndServeTLS("certificates/localhost.pem", "certificates/localhost-key.pem")
		case "prod":
			err = srv.ListenAndServe()
		default:
			err = fmt.Errorf("unknown APP_ENV: %s", env)
		}
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errC <- err
		}
	}()

	return errC, nil
}
