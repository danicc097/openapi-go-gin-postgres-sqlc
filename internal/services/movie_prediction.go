package services

import (
	"context"
	"fmt"
	"log"
	"time"

	tfidf "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pb/python-ml-app-protos/tfidf/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// TODO pass gin context.Request.Context somewhere around here...
// interceptor working but not associated to trace
func DummyMoviePrediction(ctx context.Context) error {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)

	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()

	c := tfidf.NewMovieGenreClient(conn)
	if err := callPredict(c, ctx); err != nil {
		return err
	}

	return nil
}

func callPredict(c tfidf.MovieGenreClient, ctx context.Context) error {
	synopsis := `
		Asian horror cinema often depicts stomach-churning scenes of gore and zombie outbreaks quite vividly and The Sadness ticks all the right boxes.

		Chaos and anarchy descend on the city of Taipei as residents turn into mass killers. In the wake of such a deadly viral pandemic, Jim and Kat are a young couple who seek to find each other. Violence, killing and massacre only seem to rise while the government and authorities remain complacent.

		Among the most gruesome horror movies of 2022, The Sadness lives up to its name and is not for the faint-hearted. In fact, a trigger warning is also issued at the beginning for those who may not be able to endure watching all the slashing and blood.
		`

	newCtx := metadata.AppendToOutgoingContext(ctx,
		"timestamp", time.Now().Format(time.StampNano),
		"client-id", "web-api-client-us-east-1",
		"user-id", "some-test-user-id")

	response, err := c.Predict(newCtx, &tfidf.PredictRequest{Synopsis: synopsis})
	if err != nil {
		return fmt.Errorf("calling Predict: %w", err)
	}
	log.Printf("Movie predictions: %v", response.Predictions)

	return nil
}
