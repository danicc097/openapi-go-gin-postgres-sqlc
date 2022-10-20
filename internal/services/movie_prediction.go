package services

import (
	"context"
	"fmt"
	"time"

	tfidf "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pb/python-ml-app-protos/tfidf/v1"
	"google.golang.org/grpc/metadata"
)

// moviePrediction is an external ML service to showcase calling services from others.
type moviePrediction struct {
	moviec tfidf.MovieGenreClient
}

// NewMoviePrediction returns a new moviePrediction service.
func NewMoviePrediction(moviec tfidf.MovieGenreClient) *moviePrediction {
	return &moviePrediction{
		moviec: moviec,
	}
}

func (m *moviePrediction) PredictMovieGenre(ctx context.Context, synopsis string) ([]*tfidf.Prediction, error) {
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
