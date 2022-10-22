package services

import (
	"context"

	v1 "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pb/python-ml-app-protos/tfidf/v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type Authentication struct {
	pool     *pgxpool.Pool
	logger   *zap.Logger
	movieSvc *moviePrediction
	usvc     *User
}

func NewAuthentication(pool *pgxpool.Pool, logger *zap.Logger, movieSvcClient v1.MovieGenreClient) *Authentication {
	return &Authentication{
		pool:     pool,
		logger:   logger,
		movieSvc: NewMoviePrediction(movieSvcClient),
		usvc:     NewUser(pool, logger, movieSvcClient),
	}
}

func (a *Authentication) GetUserFromToken(ctx context.Context) {
}

func (a *Authentication) GetUserFromApiKey(ctx context.Context) {
}
