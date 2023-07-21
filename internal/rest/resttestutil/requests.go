package resttestutil

import (
	"context"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/client"
)

const (
	apiKeyHeaderKey = "x-api-key"
)

func ReqWithAPIKey(apiKey string) client.RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add(apiKeyHeaderKey, apiKey)

		return nil
	}
}
