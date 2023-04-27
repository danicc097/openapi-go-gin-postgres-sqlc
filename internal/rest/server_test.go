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
	oasMw := newOpenapiMiddleware(logger, openapi)
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

		jsonErr := "{\"error\":\"invalid response\",\"message\":\"OpenAPI response validation failed\",\"validationError\":{\"detail\":[{\"detail\":\"\\nSchema:\\n  {\\n    \\\"type\\\": \\\"integer\\\"\\n  }\\n\\nValue:\\n  \\\"a_wrong_id\\\"\\n\",\"loc\":[\"id\"],\"msg\":\"value must be an integer\",\"type\":\"response_validation\"}],\"messages\":[\"response body doesn't match schema\"]}}"

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

		jsonErr := "{\"error\":\"invalid request\",\"message\":\"OpenAPI request validation failed\",\"validationError\":{\"detail\":[{\"detail\":\"\\nSchema:\\n  {\\n    \\\"type\\\": \\\"integer\\\"\\n  }\\n\\nValue:\\n  \\\"a_wrong_id\\\"\\n\",\"loc\":[\"id\"],\"msg\":\"value must be an integer\",\"type\":\"request_validation\"},{\"detail\":\"\\nSchema:\\n  {\\n    \\\"type\\\": \\\"string\\\"\\n  }\\n\\nValue:\\n  1234\\n\",\"loc\":[\"name\"],\"msg\":\"value must be a string\",\"type\":\"request_validation\"},{\"detail\":\"\\nSchema:\\n  {\\n    \\\"properties\\\": {\\n      \\\"color\\\": {\\n        \\\"type\\\": \\\"string\\\"\\n      },\\n      \\\"nestedProperty\\\": {\\n        \\\"type\\\": \\\"string\\\"\\n      }\\n    },\\n    \\\"required\\\": [\\n      \\\"nestedProperty\\\"\\n    ],\\n    \\\"type\\\": \\\"object\\\"\\n  }\\n\\nValue:\\n  {\\n    \\\"color\\\": \\\"color\\\"\\n  }\\n\",\"loc\":[\"nested\",\"nestedProperty\"],\"msg\":\"property \\\"nestedProperty\\\" is missing\",\"type\":\"request_validation\"}],\"messages\":[\"request body has an error: doesn't match schema\"]}}"

		assert.Equal(t, jsonErr, resp.Body.String())
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})
}
