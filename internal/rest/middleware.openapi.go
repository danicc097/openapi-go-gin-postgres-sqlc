package rest

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ErrorHandler is called when there is an error in validation.
type ErrorHandler func(c *gin.Context, message string, statusCode int)

// MultiErrorHandler is called when oapi returns a MultiError type.
type MultiErrorHandler func(openapi3.MultiError) error

// OAValidatorOptions customizes request validation.
type OAValidatorOptions struct {
	ValidateResponse  bool
	ErrorHandler      ErrorHandler
	Options           openapi3filter.Options
	ParamDecoder      openapi3filter.ContentParameterDecoder
	UserData          interface{}
	MultiErrorHandler MultiErrorHandler
}

// openapiMiddleware handles authentication and authorization middleware.
type openapiMiddleware struct {
	Logger *zap.SugaredLogger
	router routers.Router //
}

// TODO kin-openapi already has middleware, possibly added after this was created
// - see openapi3filter.NewValidator and tests/examples
// we just need to add our own onError func and wrap it all in gin.WrapH
// there's also onLog to use zap.
func newOpenapiMiddleware(
	logger *zap.SugaredLogger, spec *openapi3.T,
) *openapiMiddleware {
	// kinopenapi's own mux based on gorilla for validation only
	router, err := gorillamux.NewRouter(spec)
	if err != nil {
		panic(err)
	}

	return &openapiMiddleware{
		Logger: logger,
		router: router,
	}
}

// RequestValidatorWithOptions creates a validator middlewares from an openapi object.
// TODO validate responses for dev and ci at least (with openapi3filter.Strict(true)).
// need custom response writer for gin: https://github.com/gin-gonic/gin/issues/1363#issuecomment-577722498
// for reference middelware see https://github.com/aereal/go-openapi3-validation-middleware/blob/main/middleware.go
func (m *openapiMiddleware) RequestValidatorWithOptions(options *OAValidatorOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		if getSkipRequestValidationFromCtx(c) {
			c.Next()

			return
		}

		defer newOTelSpan(c.Request.Context(), "RequestValidator").End()

		rbw := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = rbw

		err := ValidateRequestFromContext(c, m.router, options)
		if err != nil {
			if options != nil && options.ErrorHandler != nil {
				options.ErrorHandler(c, err.Error(), http.StatusBadRequest)
			} else {
				// error response customized via WithCustomSchemaErrorFunc
				renderErrorResponse(c, "Invalid request", internal.WrapErrorf(err, models.ErrorCodeRequestValidation, "OpenAPI request validation failed"))
			}

			rbw.ResponseWriter.Write(rbw.body.Bytes())
			c.Abort()

			return
		}

		c.Next() // handle actual endpoint

		if !options.ValidateResponse && !getValidateResponseFromCtx(c) {
			rbw.ResponseWriter.Write(rbw.body.Bytes())

			return
		}

		rvi, err := buildRequestValidationInput(m.router, c.Request, &options.Options)
		if err != nil {
			// error response customized via WithCustomSchemaErrorFunc
			renderErrorResponse(c, fmt.Sprintf("openapi request validation input: %v", err), err)
			rbw.ResponseWriter.Write(rbw.body.Bytes())

			return
		}

		input := &openapi3filter.ResponseValidationInput{
			RequestValidationInput: rvi,
			Status:                 rbw.Status(),
			Header:                 rbw.Header(),
			Options:                &options.Options,
		}

		bodyBytes := rbw.body.Bytes()
		input.SetBodyBytes(bodyBytes)

		if err := openapi3filter.ValidateResponse(c.Request.Context(), input); err != nil {
			rbw.body.Reset()
			renderErrorResponse(c, "Invalid response", internal.WrapErrorf(err, models.ErrorCodeResponseValidation, "OpenAPI response validation failed"))
			rbw.ResponseWriter.Write(rbw.body.Bytes())

			return
		}

		rbw.ResponseWriter.Write(rbw.body.Bytes())
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	return r.body.Write(b)
}

func buildRequestValidationInput(router routers.Router, r *http.Request, options *openapi3filter.Options) (*openapi3filter.RequestValidationInput, error) {
	route, pathParams, err := router.FindRoute(r)
	if err != nil {
		return nil, fmt.Errorf("route not found: %w", err)
	}
	input := &openapi3filter.RequestValidationInput{
		Request:    r,
		PathParams: pathParams,
		Route:      route,
		Options:    options,
	}
	return input, nil
}

// ValidateRequestFromContext is called from the middleware above and actually does the work
// of validating a request.
func ValidateRequestFromContext(c *gin.Context, router routers.Router, options *OAValidatorOptions) error {
	req := c.Request
	route, pathParams, err := router.FindRoute(req)
	// We failed to find a matching route for the request.
	if err != nil {
		switch e := err.(type) {
		case *routers.RouteError:
			// We've got a bad request, the path requested doesn't match
			// either server, or path, or something.
			return internal.NewErrorf(models.ErrorCodeRequestValidation, e.Reason)
		default:
			// This should never happen today, but if our upstream code changes,
			// we don't want to crash the server, so handle the unexpected error.
			return internal.NewErrorf(models.ErrorCodeRequestValidation, "unknown error validating route: %s", err.Error())
		}
	}

	validationInput := &openapi3filter.RequestValidationInput{
		Request:    req,
		PathParams: pathParams,
		Route:      route,
	}

	// Pass the gin context into the request validator, so that any callbacks
	// which it invokes make it available.
	requestContext := context.WithValue(context.Background(), ginContextCtxKey, c)

	if options != nil {
		validationInput.Options = &options.Options
		validationInput.ParamDecoder = options.ParamDecoder
		requestContext = context.WithValue(requestContext, userDataCtxKey, options.UserData)
	}

	err = openapi3filter.ValidateRequest(requestContext, validationInput)
	if err != nil {
		// fmt.Printf("err: %T ::: %v\n", err, err)
		me := openapi3.MultiError{}
		if errors.As(err, &me) {
			errFunc := getMultiErrorHandlerFromOptions(options)
			return errFunc(me)
		}

		switch e := err.(type) {
		case *openapi3filter.RequestError:
			// We've got a bad request
			// Split up the verbose error by lines and return the first one
			// openapi errors seem to be multi-line with a decent message on the first
			errorLines := strings.Split(e.Error(), "\n")
			return internal.NewErrorf(models.ErrorCodeRequestValidation, "error in openapi3filter.RequestError: %s", errorLines[0])
		case *openapi3filter.SecurityRequirementsError:
			return internal.NewErrorf(models.ErrorCodeRequestValidation, "error in openapi3filter.SecurityRequirementsError: %s", e.Error())
		default:
			// This should never happen today, but if our upstream code changes,
			// we don't want to crash the server, so handle the unexpected error.
			return internal.NewErrorf(models.ErrorCodeRequestValidation, "unknown error validating request: %s", err)
		}
	}
	return nil
}

// attempt to get the MultiErrorHandler from the options. If it is not set,
// return a default handler.
func getMultiErrorHandlerFromOptions(options *OAValidatorOptions) MultiErrorHandler {
	if options == nil {
		return defaultMultiErrorHandler
	}

	if options.MultiErrorHandler == nil {
		return defaultMultiErrorHandler
	}

	return options.MultiErrorHandler
}

// defaultMultiErrorHandler returns a StatusBadRequest (400) and a list
// of all of the errors. This method is called if there are no other
// methods defined on the options.
func defaultMultiErrorHandler(me openapi3.MultiError) error {
	return internal.WrapErrorf(me, models.ErrorCodeRequestValidation, "validation errors encountered")
}
