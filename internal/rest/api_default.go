package rest

import (
	"net/http"
	"strings"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/static"
	"github.com/gin-gonic/gin"
)

// OpenapiYamlGet returns this very openapi spec.
func (h *dummyStrictHandlers) OpenapiYamlGet(c *gin.Context, request OpenapiYamlGetRequestObject) (OpenapiYamlGetResponseObject, error) {
	oas, err := static.SwaggerUI.ReadFile("swagger-ui/openapi.yaml")
	if err != nil {
		panic("openapi spec not found")
	}

	return OpenapiYamlGet200ApplicationxYamlResponse{
		Body: strings.NewReader(string(oas)),
	}, nil
}

// Ping ping pongs.
func (h *dummyStrictHandlers) Ping(c *gin.Context, request PingRequestObject) (PingResponseObject, error) {
	return Ping200TextResponse("pong"), nil
}

func (h *StrictHandlers) OpenapiYamlGet(c *gin.Context, request OpenapiYamlGetRequestObject) (OpenapiYamlGetResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}

func (h *StrictHandlers) Ping(c *gin.Context, request PingRequestObject) (PingResponseObject, error) {
	c.JSON(http.StatusNotImplemented, "not implemented")

	return nil, nil
}
