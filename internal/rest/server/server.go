package server

import (
	"context"
	"errors"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	rv8 "github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"

	internaldomain "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/handlers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/redis"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest/middleware"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/static"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/vault"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Config struct {
	Address string
	DB      *pgxpool.Pool
	Redis   *rv8.Client
	Metrics http.Handler
	Logger  *zap.Logger
}

// New returns a new http server.
func New(conf Config) (*http.Server, error) {
	router := gin.Default()

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

	fsys, _ := fs.Sub(static.SwaggerUI, "swagger-ui")
	authMw := middleware.NewAuth(&middleware.AuthConf{Logger: conf.Logger, Pool: conf.DB})
	vg := router.Group(os.Getenv("API_VERSION"))

	fakeSvc := services.Fake{Logger: conf.Logger, Pool: conf.DB}
	petSvc := services.Pet{Logger: conf.Logger, Pool: conf.DB}
	storeSvc := services.Store{Logger: conf.Logger, Pool: conf.DB}
	handlers.
		NewDefault().
		Register(vg, []gin.HandlerFunc{authMw.EnsureAuthenticated(), authMw.EnsureAuthorized()})
	handlers.
		NewFake(fakeSvc).
		Register(vg, []gin.HandlerFunc{})
	handlers.
		NewPet(petSvc).
		Register(vg, []gin.HandlerFunc{})
	handlers.
		NewStore(storeSvc).
		Register(vg, []gin.HandlerFunc{})
	handlers.
		NewUser(&handlers.UserConf{Logger: conf.Logger, Pool: conf.DB}).
		Register(vg, []gin.HandlerFunc{})
	// TODO /admin with authMw.EnsureAuthorized() in group

	vg.StaticFS("/docs", http.FS(fsys))

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
func Run(env, address string) (<-chan error, error) {
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
	case "dev":
		logger, err = zap.NewDevelopment()
	default:
		logger, err = zap.NewProduction()
	}

	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "internal.zapNew")
	}

	srv, err := New(Config{
		Address: address,
		DB:      pool,
		Redis:   rdb,
		Logger:  logger,
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
		err := srv.ListenAndServeTLS("certificates/localhost.pem", "certificates/localhost-key.pem")
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errC <- err
		}
	}()

	return errC, nil
}
