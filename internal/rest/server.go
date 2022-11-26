package rest

import (
	"context"
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

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	rv8 "github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zitadel/oidc/pkg/client/rp"
	httphelper "github.com/zitadel/oidc/pkg/http"
	"github.com/zitadel/oidc/pkg/oidc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	internaldomain "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	v1 "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pb/python-ml-app-protos/tfidf/v1"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/redis"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/static"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/tracing"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Config struct {
	// Port to listen to. Use ":0" for a random port.
	Address string
	Pool    *pgxpool.Pool
	Redis   *rv8.Client
	Logger  *zap.Logger
	// SpecPath is the OpenAPI spec filepath.
	SpecPath               string
	MovieSvcClient         v1.MovieGenreClient
	ScopePolicyPath        string
	RolePolicyPath         string
	MyProviderCallbackPath string
}

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

type serverOption func(*server)

// WithMiddlewares adds the given middlewares before registering the main routers.
func WithMiddlewares(mws ...gin.HandlerFunc) serverOption {
	return func(s *server) {
		s.middlewares = mws
	}
}

var key = []byte("test1234test1234")

// NewServer returns a new http server.
func NewServer(conf Config, opts ...serverOption) (*server, error) {
	err := conf.validate()
	if err != nil {
		return nil, fmt.Errorf("server config validation: %w", err)
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

	// -- oidc
	clientID := os.Getenv("OIDC_CLIENT_ID")
	clientSecret := os.Getenv("OIDC_CLIENT_SECRET")
	keyPath := os.Getenv("OIDC_KEY_PATH") // not used
	issuer := os.Getenv("OIDC_ISSUER")
	scopes := strings.Split(os.Getenv("OIDC_SCOPES"), " ")

	redirectURI := internaldomain.BuildAPIURL(conf.MyProviderCallbackPath)
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
	conf.Logger.Sugar().Infof("issuer %s", issuer)
	conf.Logger.Sugar().Infof("redirectURI %s", redirectURI)
	if err != nil {
		return nil, fmt.Errorf("error creating provider %s", err)
	}

	// generate some state (representing the state of the user in your application,
	// e.g. the page where he was before sending him to login
	state := func() string {
		return uuid.New().String()
	}

	// register the AuthURLHandler at your preferred path
	// the AuthURLHandler creates the auth request and redirects the user to the auth server
	// including state handling with secure cookie and the possibility to use PKCE
	vg.Any("/auth/myprovider/login", gin.WrapH(rp.AuthURLHandler(state, provider)))

	// for demonstration purposes the returned userinfo response is written as JSON object onto response
	marshalUserinfo := func(w http.ResponseWriter, r *http.Request, tokens *oidc.Tokens, state string, rp rp.RelyingParty, info oidc.UserInfo) {
		data, err := json.Marshal(info)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}

	// you could also just take the access_token and id_token without calling the userinfo endpoint:
	//
	// marshalToken := func(w http.ResponseWriter, r *http.Request, tokens *oidc.Tokens, state string, rp rp.RelyingParty) {
	//	data, err := json.Marshal(tokens)
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//	w.Write(data)
	//}

	// register the CodeExchangeHandler at the conf.MyProviderCallbackPath
	// the CodeExchangeHandler handles the auth response, creates the token request and calls the callback function
	// with the returned tokens from the token endpoint
	// in this example the callback function itself is wrapped by the UserinfoCallback which
	// will call the Userinfo endpoint, check the sub and pass the info into the callback function
	// TODO in reality we would redirect back to our app frontend instead
	// this callback must also be extended to create an access token (or any other login method) for our own
	// app (and create user if not found) once the external auth request is authorized
	vg.Any(conf.MyProviderCallbackPath, gin.WrapH(rp.CodeExchangeHandler(rp.UserinfoCallback(marshalUserinfo), provider)))

	// if you would use the callback without calling the userinfo endpoint, simply switch the callback handler for:
	//
	// http.Handle(conf.MyProviderCallbackPath, rp.CodeExchangeHandler(marshalToken, provider))

	// TODO /auth/logout

	// -- openapi
	oaOptions := OAValidatorOptions{
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

	openapi, err := ReadOpenAPI(conf.SpecPath)
	if err != nil {
		return nil, err
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

	rlMw := newRateLimitMiddleware(conf.Logger, 25, 10)
	switch os.Getenv("APP_ENV") {
	case "prod":
		vg.Use(rlMw.Limit())
	}

	authzsvc, err := services.NewAuthorization(conf.Logger, conf.ScopePolicyPath, conf.RolePolicyPath)
	if err != nil {
		return nil, fmt.Errorf("NewAuthorization: %w", err)
	}
	retryCount := 2
	retryInterval := 1 * time.Second
	usvc := services.NewUser(
		conf.Logger,
		repos.NewUserWithTracing(
			repos.NewUserWithRetry(
				repos.NewUserWithTimeout(
					postgresql.NewUser(), repos.UserWithTimeoutConfig{CreateTimeout: 10 * time.Second}),
				retryCount, retryInterval),
			postgresql.OtelName, nil),
		authzsvc,
	)

	authnsvc := services.NewAuthentication(conf.Logger, usvc, conf.Pool)

	vg.Use(oasMw.RequestValidatorWithOptions(&oaOptions))

	authmw := newAuthMiddleware(conf.Logger, conf.Pool, authnsvc, authzsvc, usvc)

	handlers := NewHandlers(conf.Logger, conf.Pool, conf.MovieSvcClient, usvc, authmw)

	RegisterHandlersWithOptions(vg, handlers, GinServerOptions{BaseURL: ""})

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
func Run(env, address, specPath, rolePolicyPath, scopePolicyPath string) (<-chan error, error) {
	var err error

	if err = envvar.Load(env); err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "envvar.Load")
	}

	conf := envvar.New()

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

	pool, err := postgresql.New(conf, logger)
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "postgresql.New")
	}

	rdb, err := redis.New(conf)
	if err != nil {
		return nil, internaldomain.WrapErrorf(err, internaldomain.ErrorCodeUnknown, "redis.New")
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
		Address:                address,
		Pool:                   pool,
		Redis:                  rdb,
		Logger:                 logger,
		SpecPath:               specPath,
		ScopePolicyPath:        scopePolicyPath,
		RolePolicyPath:         rolePolicyPath,
		MovieSvcClient:         v1.NewMovieGenreClient(movieSvcConn),
		MyProviderCallbackPath: "/auth/myprovider/callback",
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
