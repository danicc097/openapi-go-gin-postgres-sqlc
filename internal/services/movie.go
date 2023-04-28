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

const synopsis = `
Asian horror cinema often depicts stomach-churning scenes of gore and zombie outbreaks quite vividly and The Sadness ticks all the right boxes.
Chaos and anarchy descend on the city of Taipei as residents turn into mass killers. In the wake of such a deadly viral pandemic, Jim and Kat are a young couple who seek to find each other. Violence, killing and massacre only seem to rise while the government and authorities remain complacent.
Among the most gruesome horror movies of 2022, The Sadness lives up to its name and is not for the faint-hearted. In fact, a trigger warning is also issued at the beginning for those who may not be able to endure watching all the slashing and blood.
`

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

	if _, err := movie.Insert(ctx, m.d); err != nil {
		return fmt.Errorf("movierepo.Create: %w", err)
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
