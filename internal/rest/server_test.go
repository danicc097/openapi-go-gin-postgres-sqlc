package rest_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services/servicetestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
)

func TestValidationErrorsResponse(t *testing.T) {
	t.Parallel()

	openapi, err := rest.ReadOpenAPI("testdata/test_spec.yaml")
	if err != nil {
		t.Fatalf("could not read test spec: %s", err)
	}

	logger := testutil.NewLogger(t)
	oasMw := rest.NewOpenapiMiddleware(logger, openapi)
	oaOptions := rest.CreateOpenAPIValidatorOptions()

	t.Run("response_validation", func(t *testing.T) {
		t.Parallel()

		type ResponseValidationTestSchema struct {
			Name any `json:"name"`
			ID   any `json:"id"`
		}

		resp := httptest.NewRecorder()
		_, engine := gin.CreateTestContext(resp)
		engine.Use(oasMw.RequestValidatorWithOptions(&oaOptions))

		engine.GET("/validation_errors", func(c *gin.Context) {
			c.JSON(http.StatusOK, ResponseValidationTestSchema{
				Name: "ok",
				ID:   "a_wrong_id",
			})
		})

		req, _ := http.NewRequest(http.MethodGet, "/validation_errors", nil)
		engine.ServeHTTP(resp, req)

		var gotResp *models.HTTPError
		json.Unmarshal(resp.Body.Bytes(), &gotResp)

		wantResp := models.HTTPError{
			Detail: "OpenAPI response validation failed",
			Status: 500,
			Title:  "Invalid response",
			Type:   "ResponseValidation",
			ValidationError: &models.HTTPValidationError{
				Detail: &[]models.ValidationError{{
					Detail: struct {
						Schema map[string]any `json:"schema"`
						Value  string         `json:"value"`
					}{Schema: map[string]any{"type": string("integer")}, Value: `"a_wrong_id"`},
					Loc: []string{"id"},
					Msg: "value must be an integer",
				}},
				Messages: []string{"response body error"},
			},
		}

		if diff := cmp.Diff(wantResp, *gotResp); diff != "" {
			t.Errorf("HTTPError mismatch (-want +got):\n%s", diff)
		}
		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})

	t.Run("request_validation", func(t *testing.T) {
		t.Parallel()

		type Nested struct {
			Color          any `json:"color"`
			NestedProperty any `json:"nestedProperty,omitempty"`
		}

		type RequestValidationTestSchema struct {
			Name   any    `json:"name"`
			ID     any    `json:"id"`
			Nested Nested `json:"nested"`
		}

		resp := httptest.NewRecorder()
		_, engine := gin.CreateTestContext(resp)
		engine.Use(oasMw.RequestValidatorWithOptions(&oaOptions))

		engine.POST("/validation_errors", func(c *gin.Context) {
			c.String(http.StatusOK, "good validation")
		})

		var buf bytes.Buffer

		if err := json.NewEncoder(&buf).Encode(RequestValidationTestSchema{
			Name: 1234,
			ID:   "a_wrong_id",
			Nested: Nested{
				Color: "color",
			},
		}); err != nil {
			t.Errorf("unexpected error %v", err)
		}

		req, _ := http.NewRequest(http.MethodPost, "/validation_errors", &buf)
		req.Header.Add("Content-Type", "application/json")
		engine.ServeHTTP(resp, req)

		var gotResp *models.HTTPError
		json.Unmarshal(resp.Body.Bytes(), &gotResp)

		wantResp := models.HTTPError{
			Detail: "OpenAPI request validation failed",
			Status: 400,
			Title:  "Invalid request",
			Type:   "RequestValidation",
			ValidationError: &models.HTTPValidationError{
				Detail: &[]models.ValidationError{
					{
						Detail: struct {
							Schema map[string]any `json:"schema"`
							Value  string         `json:"value"`
						}{Schema: map[string]any{"type": string("integer")}, Value: `"a_wrong_id"`},
						Loc: []string{"id"},
						Msg: "value must be an integer",
					},
					{
						Detail: struct {
							Schema map[string]any `json:"schema"`
							Value  string         `json:"value"`
						}{Schema: map[string]any{"type": string("string")}, Value: "1234"},
						Loc: []string{"name"},
						Msg: "value must be a string",
					},
					{
						Detail: struct {
							Schema map[string]any `json:"schema"`
							Value  string         `json:"value"`
						}{
							Schema: map[string]any{
								"properties": map[string]any{
									"color":          map[string]any{"type": string("string")},
									"nestedProperty": map[string]any{"type": string("string")},
								},
								"required": []any{string("nestedProperty")},
								"type":     string("object"),
							},
							Value: `{"color":"color"}`,
						},
						Loc: []string{"nested", "nestedProperty"},
						Msg: `property "nestedProperty" is missing`,
					},
				},
				Messages: []string{"request body has an error: doesn't match schema"},
			},
		}

		if diff := cmp.Diff(wantResp, *gotResp); diff != "" {
			t.Errorf("HTTPError mismatch (-want +got):\n%s", diff)
		}
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})
}

func TestTracing(t *testing.T) {
	t.Skip("fails with -count= higher than 1 since we are using a global tracer provider.")

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
	// FIXME: fails with -count= higher than 1 since we are using a global tracer provider -> spans out of order and mixed.
	require.Equal(t, "/v2/work-item/:workItemID/comment/:workItemCommentID", spans[len(spans)-1].Name())
}
