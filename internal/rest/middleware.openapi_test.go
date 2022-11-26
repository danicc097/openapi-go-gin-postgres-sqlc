package rest

// TODO take response vlaidation into account

import (
	"context"
	_ "embed"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
)

//go:embed testdata/test_spec.yaml
var testSchema []byte

func doGet(t *testing.T, handler http.Handler, rawURL string) *httptest.ResponseRecorder {
	u, err := url.Parse(rawURL)
	if err != nil {
		t.Fatalf("Invalid url: %s", rawURL)
	}

	response := testutil.NewRequest().Get(u.RequestURI()).WithHost(u.Host).WithAcceptJson().GoWithHTTPHandler(t, handler)
	return response.Recorder
}

func doPost(t *testing.T, handler http.Handler, rawURL string, jsonBody interface{}) *httptest.ResponseRecorder {
	u, err := url.Parse(rawURL)
	if err != nil {
		t.Fatalf("Invalid url: %s", rawURL)
	}

	response := testutil.NewRequest().Post(u.RequestURI()).WithHost(u.Host).WithJsonBody(jsonBody).GoWithHTTPHandler(t, handler)
	return response.Recorder
}

func TestOapiRequestValidator(t *testing.T) {
	t.Parallel()

	openapi, err := openapi3.NewLoader().LoadFromData(testSchema)
	require.NoError(t, err, "Error initializing openapi")

	g := gin.New()

	// Set up an authenticator to check authenticated function. It will allow
	// access to "someScope", but disallow others.
	options := OAValidatorOptions{
		ErrorHandler: func(c *gin.Context, message string, statusCode int) {
			c.String(statusCode, "test: "+message)
		},
		Options: openapi3filter.Options{
			AuthenticationFunc: func(c context.Context, input *openapi3filter.AuthenticationInput) error {
				// The gin context should be propagated into here.
				gCtx := getGinContextFromCtx(c)
				assert.NotNil(t, gCtx)
				// As should user data
				assert.EqualValues(t, "hi!", getUserDataFromCtx(c))

				for _, s := range input.Scopes {
					if s == "someScope" {
						return nil
					}
					if s == "unauthorized" {
						return errors.New("unauthorized")
					}
				}
				return errors.New("forbidden")
			},
		},
		UserData: "hi!",
	}

	oasMw := newOpenapiMiddleware(zaptest.NewLogger(t), openapi)
	g.Use(oasMw.RequestValidatorWithOptions(&options))

	called := false

	// Install a request handler for /resource. We want to make sure it doesn't
	// get called.
	g.GET("/resource", func(c *gin.Context) {
		called = true
	})
	// Let's send the request to the wrong server, this should fail validation
	{
		rec := doGet(t, g, "http://not.test-server.com/resource")
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.False(t, called, "Handler should not have been called")
	}

	// Let's send a good request, it should pass
	{
		rec := doGet(t, g, "http://test-server.com/resource")
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, called, "Handler should have been called")
		called = false
	}

	// Send an out-of-spec parameter
	{
		rec := doGet(t, g, "http://test-server.com/resource?id=500")
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.False(t, called, "Handler should not have been called")
		called = false
	}

	// Send a bad parameter type
	{
		rec := doGet(t, g, "http://test-server.com/resource?id=foo")
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.False(t, called, "Handler should not have been called")
		called = false
	}

	// Add a handler for the POST message
	g.POST("/resource", func(c *gin.Context) {
		called = true
		c.AbortWithStatus(http.StatusNoContent)
	})

	called = false
	// Send a good request body
	{
		body := struct {
			Name string `json:"name"`
		}{
			Name: "Marcin",
		}
		rec := doPost(t, g, "http://test-server.com/resource", body)
		assert.Equal(t, http.StatusNoContent, rec.Code)
		assert.True(t, called, "Handler should have been called")
		called = false
	}

	// Send a malformed body
	{
		body := struct {
			Name int `json:"name"`
		}{
			Name: 7,
		}
		rec := doPost(t, g, "http://test-server.com/resource", body)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.False(t, called, "Handler should not have been called")
		called = false
	}

	g.GET("/protected_resource", func(c *gin.Context) {
		called = true
		c.AbortWithStatus(http.StatusNoContent)
	})

	// Call a protected function to which we have access
	{
		rec := doGet(t, g, "http://test-server.com/protected_resource")
		assert.Equal(t, http.StatusNoContent, rec.Code)
		assert.True(t, called, "Handler should have been called")
		called = false
	}

	g.GET("/protected_resource2", func(c *gin.Context) {
		called = true
		c.AbortWithStatus(http.StatusNoContent)
	})
	// Call a protected function to which we don't have access
	{
		rec := doGet(t, g, "http://test-server.com/protected_resource2")
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.False(t, called, "Handler should not have been called")
		called = false
	}

	g.GET("/protected_resource_401", func(c *gin.Context) {
		called = true
		c.AbortWithStatus(http.StatusNoContent)
	})
	// Call a protected function without credentials
	{
		rec := doGet(t, g, "http://test-server.com/protected_resource_401")
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, "test: error in openapi3filter.SecurityRequirementsError: security requirements failed: unauthorized", rec.Body.String())
		assert.False(t, called, "Handler should not have been called")
		called = false
	}
}

func TestRequestValidatorWithOptionsMultiError(t *testing.T) {
	t.Parallel()

	openapi, err := openapi3.NewLoader().LoadFromData([]byte(testSchema))
	require.NoError(t, err, "Error initializing openapi")

	g := gin.New()

	// Set up an authenticator to check authenticated function. It will allow
	// access to "someScope", but disallow others.
	options := OAValidatorOptions{
		Options: openapi3filter.Options{
			ExcludeRequestBody:    false,
			ExcludeResponseBody:   false,
			IncludeResponseStatus: true,
			MultiError:            true,
		},
	}

	oasMw := newOpenapiMiddleware(zaptest.NewLogger(t), openapi)
	g.Use(oasMw.RequestValidatorWithOptions(&options))

	called := false

	// Install a request handler for /resource. We want to make sure it doesn't
	// get called.
	g.GET("/multiparamresource", func(c *gin.Context) {
		called = true
	})

	// Let's send a good request, it should pass
	{
		rec := doGet(t, g, "http://test-server.com/multiparamresource?id=50&id2=50")
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, called, "Handler should have been called")
		called = false
	}

	// Let's send a request with a missing parameter, it should return
	// a bad status
	{
		rec := doGet(t, g, "http://test-server.com/multiparamresource?id=50")
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		body, err := io.ReadAll(rec.Body)
		if assert.NoError(t, err) {
			assert.Contains(t, string(body), "validation errors encountered")
			assert.Contains(t, string(body), "parameter \\\"id2\\\"")
			assert.Contains(t, string(body), "value is required but missing")
		}
		assert.False(t, called, "Handler should not have been called")
		called = false
	}

	// Let's send a request with a 2 missing parameters, it should return
	// a bad status
	{
		rec := doGet(t, g, "http://test-server.com/multiparamresource")
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		body, err := io.ReadAll(rec.Body)
		if assert.NoError(t, err) {
			assert.Contains(t, string(body), "validation errors encountered")
			assert.Contains(t, string(body), "parameter \\\"id\\\"")
			assert.Contains(t, string(body), "value is required but missing")
			assert.Contains(t, string(body), "parameter \\\"id2\\\"")
			assert.Contains(t, string(body), "value is required but missing")
		}
		assert.False(t, called, "Handler should not have been called")
		called = false
	}

	// Let's send a request with a 1 missing parameter, and another outside
	// or the parameters. It should return a bad status
	{
		rec := doGet(t, g, "http://test-server.com/multiparamresource?id=500")
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		body, err := io.ReadAll(rec.Body)
		if assert.NoError(t, err) {
			assert.Contains(t, string(body), "validation errors encountered")
			assert.Contains(t, string(body), "parameter \\\"id\\\"")
			assert.Contains(t, string(body), "number must be at most 100")
			assert.Contains(t, string(body), "parameter \\\"id2\\\"")
			assert.Contains(t, string(body), "value is required but missing")
		}
		assert.False(t, called, "Handler should not have been called")
		called = false
	}

	// Let's send a request with a parameters that do not meet spec. It should
	// return a bad status
	{
		rec := doGet(t, g, "http://test-server.com/multiparamresource?id=abc&id2=1")
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		body, err := io.ReadAll(rec.Body)
		if assert.NoError(t, err) {
			assert.Contains(t, string(body), "validation errors encountered")
			assert.Contains(t, string(body), "parameter \\\"id\\\"")
			assert.Contains(t, string(body), "abc: an invalid integer: invalid syntax")
			assert.Contains(t, string(body), "parameter \\\"id2\\\"")
			assert.Contains(t, string(body), "number must be at least 10")
		}
		assert.False(t, called, "Handler should not have been called")
		called = false
	}
}

func TestRequestValidatorWithOptionsMultiErrorAndCustomHandler(t *testing.T) {
	t.Parallel()

	openapi, err := openapi3.NewLoader().LoadFromData([]byte(testSchema))
	require.NoError(t, err, "Error initializing openapi")

	g := gin.New()

	// Set up an authenticator to check authenticated function. It will allow
	// access to "someScope", but disallow others.
	options := OAValidatorOptions{
		Options: openapi3filter.Options{
			ExcludeRequestBody:    false,
			ExcludeResponseBody:   false,
			IncludeResponseStatus: true,
			MultiError:            true,
		},
		MultiErrorHandler: func(me openapi3.MultiError) error {
			return internal.NewErrorf(internal.ErrorCodeValidationError, "Bad stuff -  %s", me.Error())
		},
	}

	oasMw := newOpenapiMiddleware(zaptest.NewLogger(t), openapi)
	g.Use(oasMw.RequestValidatorWithOptions(&options))

	called := false

	// Install a request handler for /resource. We want to make sure it doesn't
	// get called.
	g.GET("/multiparamresource", func(c *gin.Context) {
		called = true
	})

	// Let's send a good request, it should pass
	{
		rec := doGet(t, g, "http://test-server.com/multiparamresource?id=50&id2=50")
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, called, "Handler should have been called")
		called = false
	}

	// Let's send a request with a missing parameter, it should return
	// a bad status
	{
		rec := doGet(t, g, "http://test-server.com/multiparamresource?id=50")
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		body, err := io.ReadAll(rec.Body)
		if assert.NoError(t, err) {
			assert.Contains(t, string(body), "Bad stuff")
			assert.Contains(t, string(body), "parameter \\\"id2\\\"")
			assert.Contains(t, string(body), "value is required but missing")
		}
		assert.False(t, called, "Handler should not have been called")
		called = false
	}

	// Let's send a request with a 2 missing parameters, it should return
	// a bad status
	{
		rec := doGet(t, g, "http://test-server.com/multiparamresource")
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		body, err := io.ReadAll(rec.Body)
		if assert.NoError(t, err) {
			assert.Contains(t, string(body), "Bad stuff")
			assert.Contains(t, string(body), "parameter \\\"id\\\"")
			assert.Contains(t, string(body), "value is required but missing")
			assert.Contains(t, string(body), "parameter \\\"id2\\\"")
			assert.Contains(t, string(body), "value is required but missing")
		}
		assert.False(t, called, "Handler should not have been called")
		called = false
	}

	// Let's send a request with a 1 missing parameter, and another outside
	// or the parameters. It should return a bad status
	{
		rec := doGet(t, g, "http://test-server.com/multiparamresource?id=500")
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		body, err := io.ReadAll(rec.Body)
		if assert.NoError(t, err) {
			assert.Contains(t, string(body), "Bad stuff")
			assert.Contains(t, string(body), "parameter \\\"id\\\"")
			assert.Contains(t, string(body), "number must be at most 100")
			assert.Contains(t, string(body), "parameter \\\"id2\\\"")
			assert.Contains(t, string(body), "value is required but missing")
		}
		assert.False(t, called, "Handler should not have been called")
		called = false
	}

	// Let's send a request with a parameters that do not meet spec. It should
	// return a bad status
	{
		rec := doGet(t, g, "http://test-server.com/multiparamresource?id=abc&id2=1")
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		body, err := io.ReadAll(rec.Body)
		if assert.NoError(t, err) {
			assert.Contains(t, string(body), "Bad stuff")
			assert.Contains(t, string(body), "parameter \\\"id\\\"")
			assert.Contains(t, string(body), "abc: an invalid integer: invalid syntax")
			assert.Contains(t, string(body), "parameter \\\"id2\\\"")
			assert.Contains(t, string(body), "number must be at least 10")
		}
		assert.False(t, called, "Handler should not have been called")
		called = false
	}
}
