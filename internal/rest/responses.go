package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/format"
	"github.com/gin-gonic/gin"
)

// ErrorResponse represents a response containing an error message.
type ErrorResponse struct {
	Error           string                     `json:"error"`
	Message         string                     `json:"message"`
	ValidationError models.HTTPValidationError `json:"validationError,omitempty"`
}

func renderErrorResponse(c *gin.Context, msg string, err error) {
	resp := ErrorResponse{Error: msg}
	status := http.StatusInternalServerError

	var ierr *internal.Error
	if !errors.As(err, &ierr) {
		resp.Error = "internal error"
		resp.Message = msg
	} else {
		resp.Message = ierr.Cause().Error()
		switch ierr.Code() {
		case internal.ErrorCodeNotFound:
			status = http.StatusNotFound
		case internal.ErrorCodeInvalidArgument:
			status = http.StatusBadRequest
		case internal.ErrorCodeRequestValidation:
			fmt.Printf("internal.ErrorCodeRequestValidation err: %v\n", err)
			u1 := errors.Unwrap(err)
			fmt.Printf("errors.Unwrap(err) u1: %v\n", u1)
			u2 := errors.Unwrap(u1)
			fmt.Printf("errors.Unwrap(err) u2: %v\n", u2)
			u3 := errors.Unwrap(u2)
			fmt.Printf("errors.Unwrap(err) u3: %v\n", u3)
			status = http.StatusBadRequest
			resp.Message = "OpenAPI request validation failed"
			resp.ValidationError = extractValidationError(err, c, "request")
		case internal.ErrorCodeResponseValidation:
			fmt.Printf("internal.ErrorCodeResponseValidation err: %v", err)
			status = http.StatusInternalServerError
			resp.Message = "OpenAPI response validation failed"
			resp.ValidationError = extractValidationError(err, c, "response")
		case internal.ErrorCodeAlreadyExists:
			status = http.StatusConflict
		case internal.ErrorCodeUnauthorized:
			status = http.StatusForbidden
		case internal.ErrorCodeUnauthenticated:
			status = http.StatusUnauthorized
		case internal.ErrorCodeUnknown:
			fallthrough
		default:
			status = http.StatusInternalServerError
		}
	}

	if err != nil {
		span := newOTELSpan(c.Request.Context(), "renderErrorResponse")
		defer span.End()

		span.RecordError(err)
	}

	renderResponse(c, resp, status)
}

func extractValidationError(err error, c *gin.Context, typ string) models.HTTPValidationError {
	unwrappedErr := errors.Unwrap(err) // first wrap can always be discarded
	if err := errors.Unwrap(unwrappedErr); err != nil {
		// don't want to unwrap up to the custom schema error
		if !strings.HasPrefix(strings.TrimSpace(err.Error()), ValidationErrorSeparator) {
			unwrappedErr = err
		}
	}

	s := strings.Split(unwrappedErr.Error(), ValidationErrorSeparator)
	oe := s[0] // kin-openapi original concatenated error string
	validationErrors := s[1:]

	oe = strings.TrimSuffix(strings.TrimSpace(oe), ":")
	origErrs := strings.Split(oe, " | ")
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
	//
	// FIXME actually we first need to split by " | " and then grab schemaError (if any splitting $$$$), not the other way around: OpenAPI request validation failed: validation errors encountered: parameter "id" in query has an error: value abc: an invalid integer: invalid syntax | parameter "id2" in query has an error: $$$${"detail":"\nSchema:\n  {\n    \"maximum\": 100,\n    \"minimum\": 10,\n    \"type\": \"integer\"\n  }\n\nValue:\n  1\n","loc":null,"msg":"number must be at least 10","type":"unknown"}
	format.PrintJSON(origErrs)
	fmt.Printf("validationErrors: %+v\n\n", validationErrors)

	var vErrs []models.ValidationError

	for _, vErrStrings := range validationErrors {
		for _, vErrString := range strings.Split(vErrStrings, " | ") { // kin-openapi sep
			var vErr models.ValidationError

			if err := json.Unmarshal([]byte(vErrString), &vErr); err != nil {
				// instead err could be a string, which will only shown in callout via origErr, or be badly formatted
				continue
			}

			// in any case we don't want validation errors with empty loc
			if len(vErr.Loc) == 0 {
				continue
			}

			if typ == "request" {
				vErr.Type = models.HttpErrorTypeRequestValidation
			} else {
				vErr.Type = models.HttpErrorTypeResponseValidation
			}

			vErrs = append(vErrs, vErr)
		}
	}

	httpValidationError := models.HTTPValidationError{
		Detail:   &vErrs,
		Messages: origErrs,
	}

	return httpValidationError
}

func renderResponse(c *gin.Context, res interface{}, status int) {
	c.Header("Content-Type", "application/json")

	content, err := json.Marshal(res)
	if err != nil {
		// TODO Do something with the error
		fmt.Printf("error in renderResponse Marshal: %s", err)
		// c.Status(http.StatusInternalServerError)

		return
	}

	c.Status(status)

	if _, err = c.Writer.Write(content); err != nil { //nolint: staticcheck
		// TODO Do something with the error
		fmt.Printf("error in renderResponse Write: %s", err)
	}
}
