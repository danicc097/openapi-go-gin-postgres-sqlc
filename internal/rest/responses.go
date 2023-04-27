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
			status = http.StatusBadRequest
			resp.Message = "OpenAPI request validation failed"
			resp.ValidationError = extractValidationError(err, "request")
		case internal.ErrorCodeResponseValidation:
			status = http.StatusInternalServerError
			resp.Message = "OpenAPI response validation failed"
			resp.ValidationError = extractValidationError(err, "response")
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

func extractValidationError(err error, typ string) models.HTTPValidationError {
	var origErrs []string
	var vErrs []models.ValidationError

	unwrappedErr := err
	for uErr := errors.Unwrap(unwrappedErr); uErr != nil; {
		e := strings.TrimSpace(uErr.Error())
		if strings.HasPrefix(e, "response body doesn't match schema:") {
			unwrappedErr = errors.Unwrap(uErr)

			origErrs = append(origErrs, "response body error") // this is obviously not catched on client-side validation

			break
		}
		if strings.HasPrefix(e, "validation errors encountered:") {
			unwrappedErr = errors.Unwrap(uErr)

			break
		}
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
