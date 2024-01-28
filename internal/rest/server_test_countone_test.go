//go:build !skip_countone

package rest_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
)

// TODO: exclude CountOne_ in -run flag by default.
// need to pass default flag when doing watch-for later with AND.
// -run uses go regex so see: https://pkg.go.dev/regexp/syntax
func TestTracing(t *testing.T) {
	// for better architecture see
	// https://github.com/open-telemetry/opentelemetry-go/discussions/4532
	// (still not suitable for unit tests), see this instead -> https://github.com/open-telemetry/opentelemetry-go/pull/4539
	// as of now must run with count=1

	t.Parallel()

	srv, err := runTestServer(t, testPool)
	srv.setupCleanup(t)
	require.NoError(t, err, "Couldn't run test server: %s\n")

	otel.SetTracerProvider(srv.tp) // IMPORTANT: most likely leaks into other tests.

	svc := services.New(testutil.NewLogger(t), services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	ufixture := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
		WithAPIKey: true,
		Scopes:     []models.Scope{models.ScopeWorkItemCommentDelete},
	})

	workItemCommentf := ff.CreateWorkItemComment(context.Background(), servicetestutil.CreateWorkItemCommentParams{Project: models.ProjectDemo, UserID: ufixture.User.UserID})

	id := workItemCommentf.WorkItemComment.WorkItemCommentID
	res, err := srv.client.DeleteWorkItemCommentWithResponse(context.Background(), workItemCommentf.WorkItem.WorkItemID, id, ReqWithAPIKey(ufixture.APIKey.APIKey))
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, res.StatusCode(), string(res.Body))

	spans := srv.spanRecorder.Ended()
	for _, ros := range spans {
		t.Logf("%+v", ros.Name())
	}
	require.NotEmpty(t, spans)
	require.Equal(t, "/v2/work-item/:workItemID/comment/:workItemCommentID", spans[len(spans)-1].Name())
}
