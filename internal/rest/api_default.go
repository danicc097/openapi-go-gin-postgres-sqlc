package rest

import (
	"strings"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/static"
	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) OpenapiYamlGet(c *gin.Context, request OpenapiYamlGetRequestObject) (OpenapiYamlGetResponseObject, error) {
	oas, err := static.SwaggerUI.ReadFile("swagger-ui/openapi.yaml")
	if err != nil {
		panic("openapi spec not found")
	}

	return OpenapiYamlGet200ApplicationxYamlResponse{
		Body: strings.NewReader(string(oas)),
	}, nil
}

func (h *StrictHandlers) Ping(c *gin.Context, request PingRequestObject) (PingResponseObject, error) {
	return Ping200TextResponse("pong"), nil
}
