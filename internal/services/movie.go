package services

import (
	"context"
	"fmt"
	"time"

	tfidf "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pb/python-ml-app-protos/tfidf/v1"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

// movie is an external ML service to showcase calling services from others.
type movie struct {
	moviec    tfidf.MovieGenreClient
	movierepo tfidf.MovieGenreClient
	logger    *zap.Logger
	d         db.DBTX
}

// NewMovie returns a new movie service.
func NewMovie(d db.DBTX, logger *zap.Logger, moviec tfidf.MovieGenreClient) *movie {
	return &movie{
		d:      d,
		moviec: moviec,
		logger: logger,
	}
}

func (m *movie) Create(ctx context.Context, movie *db.Movie) error {
	// TODO repo (once figured out transactions)

	predictions, _ := m.PredictGenre(ctx, synopsis)
	m.logger.Sugar().Infof("Movie predictions: %v", predictions)

	if err := movie.Insert(ctx, m.d); err != nil {
		return errors.Wrap(err, "movierepo.Create")
	}

	return nil
}

func (m *movie) PredictGenre(ctx context.Context, synopsis string) ([]*tfidf.Prediction, error) {
	newCtx := metadata.AppendToOutgoingContext(ctx,
		"timestamp", time.Now().Format(time.StampNano),
		"client-id", "web-api-client-us-east-1",
		"user-id", "some-test-user-id")

	response, err := m.moviec.Predict(newCtx, &tfidf.PredictRequest{Synopsis: synopsis})
	if err != nil {
		return nil, fmt.Errorf("calling Predict: %w", err)
	}

	return response.Predictions, nil
}
