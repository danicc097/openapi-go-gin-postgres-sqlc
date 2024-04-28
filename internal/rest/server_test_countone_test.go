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
// nolint: paralleltest
func TestTracing(t *testing.T) {
	// for better architecture see
	// https://github.com/open-telemetry/opentelemetry-go/discussions/4532
	// (still not suitable for unit tests), see this instead -> https://github.com/open-telemetry/opentelemetry-go/pull/4539
	// as of now must run with count=1
	t.Parallel()

	srv, err := runTestServer(t, context.Background(), testPool)
	srv.setupCleanup(t)
	require.NoErrorf(t, err, "Couldn't run test server\n")

	otel.SetTracerProvider(srv.tp) // IMPORTANT: leaks into other tests.

	svc := services.New(testutil.NewLogger(t), services.CreateTestRepos(t), testPool)
	ff := servicetestutil.NewFixtureFactory(t, testPool, svc)

	ufixture := ff.CreateUser(context.Background(), servicetestutil.CreateUserParams{
		WithAPIKey: true,
		Scopes:     []models.Scope{models.ScopeWorkItemCommentDelete},
	})

	requiredProject := models.ProjectDemo
	teamf := ff.CreateTeam(context.Background(), servicetestutil.CreateTeamParams{Project: requiredProject})
	workItemf := ff.CreateWorkItem(context.Background(), requiredProject, *services.NewCtxUser(ufixture.User), teamf.TeamID)
	workItemCommentf := ff.CreateWorkItemComment(context.Background(), ufixture.UserID, workItemf.WorkItemID)

	id := workItemCommentf.WorkItemCommentID
	res, err := srv.client.DeleteWorkItemCommentWithResponse(context.Background(), workItemf.WorkItemID, id, ReqWithAPIKey(ufixture.APIKey.APIKey))
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, res.StatusCode(), string(res.Body))

	spans := srv.spanRecorder.Ended()
	require.NotEmpty(t, spans)
	// otelgin's tracer sometimes doesn't contain first call
	// require.True(t, slices.ContainsMatch(spans, func(item trace.ReadOnlySpan) bool {
	// 	return strings.Contains(item.Name(), "/v2/work-item/:workItemID/comment/:workItemCommentID")
	// }))
}
