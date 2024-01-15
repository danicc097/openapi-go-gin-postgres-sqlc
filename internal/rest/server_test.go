package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestValidationErrorsResponse(t *testing.T) {
	t.Parallel()

	openapi, err := ReadOpenAPI("testdata/test_spec.yaml")
	if err != nil {
		t.Fatalf("could not read test spec: %s", err)
	}

	logger := testutil.NewLogger(t)
	oasMw := NewOpenapiMiddleware(logger, openapi)
	oaOptions := createOpenAPIValidatorOptions()

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
