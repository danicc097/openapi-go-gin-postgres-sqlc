package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/gin-gonic/gin"
)

const (
	GinContextKey = "oapi-codegen/gin-context"
	UserDataKey   = "oapi-codegen/user-data"
)

// ErrorHandler is called when there is an error in validation
type ErrorHandler func(c *gin.Context, message string, statusCode int)

// MultiErrorHandler is called when oapi returns a MultiError type
type MultiErrorHandler func(openapi3.MultiError) error

// OAValidatorOptions to customize request validation. These are passed through to
// openapi3filter.
type OAValidatorOptions struct {
	ErrorHandler      ErrorHandler
	Options           openapi3filter.Options
	ParamDecoder      openapi3filter.ContentParameterDecoder
	UserData          interface{}
	MultiErrorHandler MultiErrorHandler
}

// TODO use renderErrorResponse instead
// Create a validator from a openapi object, with validation options
func OapiRequestValidatorWithOptions(openapi *openapi3.T, options *OAValidatorOptions) gin.HandlerFunc {
	router, err := gorillamux.NewRouter(openapi)
	if err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		defer newOTELSpan(c.Request.Context(), "OapiRequestValidatorWithOptions").End()

		err := ValidateRequestFromContext(c, router, options)
		if err != nil {
			if options != nil && options.ErrorHandler != nil {
				options.ErrorHandler(c, err.Error(), http.StatusBadRequest)
				// in case the handler didn't internally call Abort, stop the chain
				c.Abort()
			} else {
				// TODO renderErrorResponse instead. Should parse errors to be more rfc7807 friendly
				c.Abort()

				renderErrorResponse(c, "openapi validation bad request", err)
			}
		}
		c.Next()
	}
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
			return internal.NewErrorf(internal.ErrorCodeValidationError, e.Reason)
		default:
			// This should never happen today, but if our upstream code changes,
			// we don't want to crash the server, so handle the unexpected error.
			return internal.NewErrorf(internal.ErrorCodeValidationError, "unknown error validating route: %s", err.Error())
		}
	}

	validationInput := &openapi3filter.RequestValidationInput{
		Request:    req,
		PathParams: pathParams,
		Route:      route,
	}

	// Pass the gin context into the request validator, so that any callbacks
	// which it invokes make it available.
	requestContext := context.WithValue(context.Background(), GinContextKey, c)

	if options != nil {
		validationInput.Options = &options.Options
		validationInput.ParamDecoder = options.ParamDecoder
		requestContext = context.WithValue(requestContext, UserDataKey, options.UserData)
	}

	err = openapi3filter.ValidateRequest(requestContext, validationInput)
	if err != nil {
		fmt.Printf("err: %T ::: %v\n", err, err)
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
			return internal.NewErrorf(internal.ErrorCodeValidationError, "error in openapi3filter.RequestError: %s", errorLines[0])
		case *openapi3filter.SecurityRequirementsError:
			return internal.NewErrorf(internal.ErrorCodeValidationError, "error in openapi3filter.SecurityRequirementsError: %s", e.Error())
		default:
			// This should never happen today, but if our upstream code changes,
			// we don't want to crash the server, so handle the unexpected error.
			return internal.NewErrorf(internal.ErrorCodeValidationError, "unknown error validating request: %s", err)
		}
	}
	return nil
}

// Helper function to get the gin context from within requests. It returns
// nil if not found or wrong type.
func GetGinContext(c context.Context) *gin.Context {
	iface := c.Value(GinContextKey)
	if iface == nil {
		return nil
	}
	ginCtx, ok := iface.(*gin.Context)
	if !ok {
		return nil
	}
	return ginCtx
}

func GetUserData(c context.Context) interface{} {
	return c.Value(UserDataKey)
}

// attempt to get the MultiErrorHandler from the options. If it is not set,
// return a default handler
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
	return internal.NewErrorf(internal.ErrorCodeValidationError, "multiple errors encountered: %s", me)
}
