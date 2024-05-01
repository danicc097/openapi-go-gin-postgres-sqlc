package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/slices"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// renderErrorResponse writes an error response from title and error.
// title represents an error title which will be shown to end users.
// Inspired by https://www.rfc-editor.org/rfc/rfc7807.
func renderErrorResponse(c *gin.Context, title string, err error) {
	if err == nil {
		err = errors.New("unknown error")
	}

	resp := models.HTTPError{
		Title:  title,
		Error:  err.Error(),
		Type:   models.ErrorCodeUnknown,
		Status: http.StatusInternalServerError,
	}

	/**
	 *
	 *
	 *

	o  "type" (string) - A URI reference [RFC3986] that identifies the
		 problem type.  This specification encourages that, when
		 dereferenced, it provide human-readable documentation for the
		 problem type (e.g., using HTML [W3C.REC-html5-20141028]).  When
		 this member is not present, its value is assumed to be
		 "about:blank".
		 TODO: simple html page with fragments, generated from mapping
		 ErrCode to go tmpl
		 : /problems#NotFound

	o  "title" (string) - A short, human-readable summary of the problem
		 type.  It SHOULD NOT change from occurrence to occurrence of the
		 problem, except for purposes of localization (e.g., using
		 proactive content negotiation; see [RFC7231], Section 3.4).

	o  "status" (number) - The HTTP status code ([RFC7231], Section 6)
		 generated by the origin server for this occurrence of the problem.

	o  "detail" (string) - A human-readable explanation specific to this
		 occurrence of the problem.
	*/

	var ierr *internal.Error
	if !errors.As(err, &ierr) {
		resp.Title = "internal error"
		resp.Detail = title
	} else {
		resp.Loc = pointers.New(ierr.Loc())
		resp.Type = ierr.Code()
		resp.Detail = ierr.Cause().Error()
		switch ierr.Code() {
		case models.ErrorCodeNotFound:
			resp.Status = http.StatusNotFound
		case models.ErrorCodeInvalidArgument:
			resp.Status = http.StatusBadRequest
		case models.ErrorCodeInvalidRole, models.ErrorCodeInvalidScope, models.ErrorCodeInvalidUUID:
			resp.Status = http.StatusBadRequest
		case models.ErrorCodeRequestValidation:
			resp.Status = http.StatusBadRequest
			resp.Detail = "OpenAPI request validation failed"
			resp.Error = "" // will use validationError
			resp.ValidationError = extractValidationError(err)
		case models.ErrorCodeResponseValidation:
			resp.Status = http.StatusInternalServerError
			resp.Detail = "OpenAPI response validation failed"
			resp.Error = "" // will use validationError
			resp.ValidationError = extractValidationError(err)
		case models.ErrorCodeAlreadyExists:
			resp.Status = http.StatusConflict
		case models.ErrorCodeUnauthorized:
			resp.Status = http.StatusForbidden
		case models.ErrorCodeUnauthenticated:
			resp.Status = http.StatusUnauthorized
		case models.ErrorCodePrivate:
			if os.Getenv("TESTING") != "" {
				resp = models.HTTPError{Title: "internal error", Detail: "internal error"}
			}

			fallthrough
		case models.ErrorCodeUnknown:
			fallthrough
		default:
			resp.Status = http.StatusInternalServerError
		}
	}

	if err != nil {
		span := newOTelSpan().Build(c.Request.Context())
		defer span.End()

		opts := []trace.EventOption{}

		var ierr *internal.Error
		if errors.As(err, &ierr) {
			opts = append(opts, trace.WithAttributes(
				attribute.Key("type").String(string(ierr.Code())),
				attribute.Key("loc").StringSlice(ierr.Loc())),
			)
		}

		span.RecordError(err, opts...)
	}

	CtxWithRequestError(c)

	renderResponse(c, resp, resp.Status)
}

func extractValidationError(err error) *models.HTTPValidationError {
	var origErrs []string
	var vErrs []models.ValidationError // nolint: prealloc

	unwrappedErr := err

	for maxCalls, uErr := 10, errors.Unwrap(unwrappedErr); uErr != nil && maxCalls > 0; {
		e := strings.TrimSpace(uErr.Error())
		if strings.HasPrefix(e, "response body doesn't match schema") {
			unwrappedErr = errors.Unwrap(uErr)

			origErrs = append(origErrs, "response body error") // this is obviously not catched on client-side validation

			break
		}
		if strings.HasPrefix(e, "validation errors encountered:") {
			unwrappedErr = errors.Unwrap(uErr)

			break
		}
		maxCalls--
	}

	// NOTE: custom schema error count may not match error strings, e.g. missing parameters, etc.
	// what we can do is have HTTPError have a messages []string field that contains splitted errors, e.g.
	// - parameter "id" in query has an error: value is required but missing
	// - parameter "id2" in query has an error: value is required but missing
	// which is good enough to always show in a callout. Note we already have client-side validation
	// that should take care of most errors. But if we miss something, it is sufficient
	// to get descriptive errors (HTTPValidationError.messages) and additionally validationErrors with locs
	// to mark as invalid in form with error text of ValidationError.Msg (this Msg is not in HTTPValidationError.messages)
	// then, to get exact location (if available), we parse validationErrors independently of the above, and if they have a loc
	// we mark the form field. empty locs -> don't save to validationErrors in the first place (will correspond to e.g. parameters and will be
	// shown in callout only)

	validationErrors := unwrappedErr.Error()

	for _, validationError := range strings.Split(validationErrors, " | ") { // kin-openapi sep
		origErr, vErrString, _ := strings.Cut(validationError, ValidationErrorSeparator)
		origErr = strings.TrimSuffix(strings.TrimSpace(origErr), ":")

		var vErr models.ValidationError

		if err := json.Unmarshal([]byte(vErrString), &vErr); err != nil {
			// instead err could be a string (which will only be shown in callout via origErr) or badly formatted
			origErrs = append(origErrs, origErr)

			continue
		}

		if len(vErr.Loc) == 0 {
			// in any case we don't want validation errors with empty loc.
			// but do keep the err message
			origErrs = append(origErrs, origErr+": "+vErr.Msg)

			continue
		} else {
			origErrs = append(origErrs, origErr)
		}

		vErrs = append(vErrs, vErr)
	}

	httpValidationError := models.HTTPValidationError{
		Detail:   &vErrs,
		Messages: slices.RemoveEmptyString(origErrs),
	}

	return &httpValidationError
}

func renderResponse(c *gin.Context, res any, status int) {
	c.Header("Content-Type", "application/json")

	content, err := json.Marshal(res)
	if err != nil {
		// TODO Do something with the error
		fmt.Printf("error in renderResponse Marshal: %s\n", err)
		// c.Status(http.StatusInternalServerError)

		return
	}

	c.Status(status)

	if _, err = c.Writer.Write(content); err != nil { //nolint: staticcheck
		// TODO Do something with the error
		fmt.Printf("error in renderResponse Write: %s\n", err)
	}
}

// parseBody attempts to bind the given struct to the request body
// and returns whether the caller should exit early or not.
func parseBody(c *gin.Context, body any) bool {
	if err := c.BindJSON(body); err != nil {
		// openapi validator should have caught a validation error beforehand
		// for routes that have validator mw enabled.
		renderErrorResponse(c, "Invalid data", internal.WrapErrorf(err, models.ErrorCodeInvalidArgument, "invalid data"))

		return true
	}

	return false
}

func rawMessage(data any) json.RawMessage {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return json.RawMessage{}
	}

	return json.RawMessage(jsonData)
}
