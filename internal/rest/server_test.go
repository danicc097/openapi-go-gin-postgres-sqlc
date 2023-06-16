package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestValidationErrorsResponse(t *testing.T) {
	t.Parallel()

	openapi, err := ReadOpenAPI("testdata/test_spec.yaml")
	if err != nil {
		t.Fatalf("could not read test spec: %s", err)
	}

	logger, _ := zap.NewDevelopment()
	oasMw := newOpenapiMiddleware(logger.Sugar(), openapi)
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

		jsonErr := "{\"title\":\"invalid response\",\"detail\":\"OpenAPI response validation failed\",\"status\":500,\"error\":\"OpenAPI response validation failed: response body doesn't match schema: $$$${\\\"detail\\\":{\\\"schema\\\":{\\\"type\\\":\\\"integer\\\"},\\\"value\\\":\\\"\\\\\\\"a_wrong_id\\\\\\\"\\\"},\\\"loc\\\":[\\\"id\\\"],\\\"msg\\\":\\\"value must be an integer\\\",\\\"type\\\":\\\"unknown\\\"}\",\"validationError\":{\"detail\":[{\"detail\":{\"schema\":{\"type\":\"integer\"},\"value\":\"\\\"a_wrong_id\\\"\"},\"loc\":[\"id\"],\"msg\":\"value must be an integer\",\"type\":\"response_validation\"}],\"messages\":[\"response body error\"]}}"

		assert.Equal(t, jsonErr, resp.Body.String())
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

		jsonErr := "{\"title\":\"invalid request\",\"detail\":\"OpenAPI request validation failed\",\"status\":400,\"error\":\"OpenAPI request validation failed: validation errors encountered: request body has an error: doesn't match schema: $$$${\\\"detail\\\":{\\\"schema\\\":{\\\"type\\\":\\\"integer\\\"},\\\"value\\\":\\\"\\\\\\\"a_wrong_id\\\\\\\"\\\"},\\\"loc\\\":[\\\"id\\\"],\\\"msg\\\":\\\"value must be an integer\\\",\\\"type\\\":\\\"unknown\\\"} | $$$${\\\"detail\\\":{\\\"schema\\\":{\\\"type\\\":\\\"string\\\"},\\\"value\\\":\\\"1234\\\"},\\\"loc\\\":[\\\"name\\\"],\\\"msg\\\":\\\"value must be a string\\\",\\\"type\\\":\\\"unknown\\\"} | $$$${\\\"detail\\\":{\\\"schema\\\":{\\\"properties\\\":{\\\"color\\\":{\\\"type\\\":\\\"string\\\"},\\\"nestedProperty\\\":{\\\"type\\\":\\\"string\\\"}},\\\"required\\\":[\\\"nestedProperty\\\"],\\\"type\\\":\\\"object\\\"},\\\"value\\\":\\\"{\\\\\\\"color\\\\\\\":\\\\\\\"color\\\\\\\"}\\\"},\\\"loc\\\":[\\\"nested\\\",\\\"nestedProperty\\\"],\\\"msg\\\":\\\"property \\\\\\\"nestedProperty\\\\\\\" is missing\\\",\\\"type\\\":\\\"unknown\\\"}\",\"validationError\":{\"detail\":[{\"detail\":{\"schema\":{\"type\":\"integer\"},\"value\":\"\\\"a_wrong_id\\\"\"},\"loc\":[\"id\"],\"msg\":\"value must be an integer\",\"type\":\"request_validation\"},{\"detail\":{\"schema\":{\"type\":\"string\"},\"value\":\"1234\"},\"loc\":[\"name\"],\"msg\":\"value must be a string\",\"type\":\"request_validation\"},{\"detail\":{\"schema\":{\"properties\":{\"color\":{\"type\":\"string\"},\"nestedProperty\":{\"type\":\"string\"}},\"required\":[\"nestedProperty\"],\"type\":\"object\"},\"value\":\"{\\\"color\\\":\\\"color\\\"}\"},\"loc\":[\"nested\",\"nestedProperty\"],\"msg\":\"property \\\"nestedProperty\\\" is missing\",\"type\":\"request_validation\"}],\"messages\":[\"request body has an error: doesn't match schema\"]}}"

		assert.Equal(t, jsonErr, resp.Body.String())
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})
}
