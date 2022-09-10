package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/gin-gonic/gin"
)

const otelName = "github.com/MarioCarrion/todo-api/internal/rest"

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
	if !errors.As(err, &ierr) {
		resp.Error = "internal error"
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

		case internal.ErrorCodeAlreadyExists:
			status = http.StatusConflict
		case internal.ErrorCodeUnknown:
			fallthrough
		default:
			status = http.StatusInternalServerError
		}
	}

	fmt.Printf("Error: %v\n", err)

	renderResponse(c, resp, status)
}

func renderResponse(c *gin.Context, res interface{}, status int) {
	c.Header("Content-Type", "application/json")

	content, err := json.Marshal(res)
	if err != nil {
		// XXX Do something with the error ;)
		fmt.Printf("error in renderResponse Marshal: %s", err)
		// c.Status(http.StatusInternalServerError)

		return
	}

	c.Status(status)

	if _, err = c.Writer.Write(content); err != nil { //nolint: staticcheck
		// XXX Do something with the error ;)
		fmt.Printf("error in renderResponse Write: %s", err)
	}
}
