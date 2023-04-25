package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
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
			resp.ValidationError = extractValidationError(err.Error(), c, "request")
		case internal.ErrorCodeResponseValidation:
			status = http.StatusInternalServerError
			resp.Message = "OpenAPI response validation failed"
			resp.ValidationError = extractValidationError(err.Error(), c, "response")
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

func extractValidationError(errString string, c *gin.Context, typ string) models.HTTPValidationError {
	validationErrors := strings.Split(errString, ValidationErrorSeparator)[1:]
	vErrs := make([]models.ValidationError, len(validationErrors))

	for i, vErrString := range validationErrors {
		vErrString := strings.Split(vErrString, "|")[0]
		var vErr models.ValidationError

		err := json.Unmarshal([]byte(vErrString), &vErr)
		if typ == "request" {
			vErr.Type = models.HttpErrorTypeRequestValidation
		} else {
			vErr.Type = models.HttpErrorTypeResponseValidation
		}
		vErrs[i] = vErr
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("invalid ValidationError: %s", vErrString))
			c.Abort()
		}
	}

	httpValidationError := models.HTTPValidationError{
		Detail: &vErrs,
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
