package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/slices"
	"github.com/gin-gonic/gin"
)

// ErrorResponse represents a response containing an error message.
type ErrorResponse struct {
	Title           string                      `json:"title"`
	Detail          string                      `json:"detail"`
	Status          int                         `json:"status"`
	Error           string                      `json:"error"`
	Type            string                      `json:"type"`
	ValidationError *models.HTTPValidationError `json:"validationError,omitempty"`
}

// renderErrorResponse writes an error response from title and error.
// Inspired by https://www.rfc-editor.org/rfc/rfc7807.
func renderErrorResponse(c *gin.Context, title string, err error) {
	resp := ErrorResponse{
		Title: title, Error: err.Error(),
		Type:   internal.ErrorCodeUnknown.String(),
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
		resp.Type = ierr.Code().String()
		resp.Detail = ierr.Cause().Error()
		switch ierr.Code() {
		case internal.ErrorCodeNotFound:
			resp.Status = http.StatusNotFound
		case internal.ErrorCodeInvalidArgument:
			resp.Status = http.StatusBadRequest
		case internal.ErrorCodeInvalidRole, internal.ErrorCodeInvalidScope, internal.ErrorCodeInvalidUUID:
			resp.Status = http.StatusBadRequest
		case internal.ErrorCodeRequestValidation:
			resp.Status = http.StatusBadRequest
			resp.Detail = "OpenAPI request validation failed"
			resp.ValidationError = extractValidationError(err, "request")
		case internal.ErrorCodeResponseValidation:
			resp.Status = http.StatusInternalServerError
			resp.Detail = "OpenAPI response validation failed"
			resp.ValidationError = extractValidationError(err, "response")
		case internal.ErrorCodeAlreadyExists:
			resp.Status = http.StatusConflict
		case internal.ErrorCodeUnauthorized:
			resp.Status = http.StatusForbidden
		case internal.ErrorCodeUnauthenticated:
			resp.Status = http.StatusUnauthorized
		case internal.ErrorCodePrivate:
			resp = ErrorResponse{Title: "internal error", Detail: "internal error"}

			fallthrough
		case internal.ErrorCodeUnknown:
			fallthrough
		default:
			resp.Status = http.StatusInternalServerError
		}
	}

	if err != nil {
		span := newOTELSpan(c.Request.Context(), "renderErrorResponse")
		defer span.End()

		span.RecordError(err)
	}

	renderResponse(c, resp, resp.Status)
}

func extractValidationError(err error, typ string) *models.HTTPValidationError {
	var origErrs []string
	var vErrs []models.ValidationError

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

		switch typ {
		case "request":
			vErr.Type = models.HttpErrorTypeRequestValidation
		case "response":
			vErr.Type = models.HttpErrorTypeResponseValidation
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
