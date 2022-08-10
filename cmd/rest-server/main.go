// Code adapted from:
// https://github.com/MarioCarrion/todo-api-microservice-example

package main

import (
	"context"
	"errors"
	"flag"
	"io/fs"
	"log"
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

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/cmd/internal"
	internaldomain "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/environment"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	gen "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/gen"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/static"
	// "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgresql"
	// "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/redis"
	// "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	// "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/service"
)

func main() {
	var env, address string

	flag.StringVar(&env, "env", "", "Environment Variables filename")
	flag.StringVar(&address, "address", ":8090", "HTTP Server Address")
	flag.Parse()

	errC, err := run(env, address)
	if err != nil {
		log.Fatalf("Couldn't run: %s", err)
	}

	if err := <-errC; err != nil {
		log.Fatalf("Error while running: %s", err)
	}
}

func run(env, address string) (<-chan error, error) {
	if err := envvar.Load(env); err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "envvar.Load")
	}

	vault, err := internal.NewVaultProvider()
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "internal.NewVaultProvider")
	}

	conf := envvar.New(vault)

	pool, err := internal.NewPostgreSQL(conf)
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "internal.NewPostgreSQL")
	}

	environment.Pool = pool

	rdb, err := internal.NewRedis(conf)
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "internal.NewRedis")
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "internal.NewProduction")
	}

	environment.Logger = logger

	srv, err := newServer(serverConfig{
		Address: address,
		DB:      pool,
		Redis:   rdb,
		Logger:  logger,
	})
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "newServer")
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
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errC <- err
		}
	}()

	return errC, nil
}

type serverConfig struct {
	Address string
	DB      *pgxpool.Pool
	Redis   *rv8.Client
	Metrics http.Handler
	Logger  *zap.Logger
}

func newServer(conf serverConfig) (*http.Server, error) {
	router := gen.NewRouter()

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

	// TODO defining static file serving in spec is not supported?
	fsys, _ := fs.Sub(static.SwaggerUI, "swagger-ui")
	router.StaticFS("/v2/docs", http.FS(fsys))

	conf.Logger.Info("Server started")
	log.Fatal(router.Run(conf.Address))

	return &http.Server{
		Handler:           router,
		Addr:              conf.Address,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       1 * time.Second,
	}, nil
}
