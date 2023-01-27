package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/gin-gonic/gin"
)

// ErrorResponse represents a response containing an error message.
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	// Validations validation.Errors `json:"validations,omitempty"`
}

func renderErrorResponse(c *gin.Context, msg string, err error) {
	resp := ErrorResponse{Error: msg}
	status := http.StatusInternalServerError

	var ierr *internal.Error
	fmt.Printf("err: %v\n", err)
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

			// TODO kin errors render response in the middleware
			// var verrors validation.Errors
			// if errors.As(ierr, &verrors) {
			// 	resp.Validations = verrors
			// }

		case internal.ErrorCodeValidationError:
			status = http.StatusBadRequest
			resp.Message = ierr.Error()
		case internal.ErrorCodeResponseValidationError:
			status = http.StatusUnprocessableEntity
			resp.Message = ierr.Error()
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
