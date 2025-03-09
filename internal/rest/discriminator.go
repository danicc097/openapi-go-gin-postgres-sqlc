package rest

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/gin-gonic/gin"
)

type PolymorphicBody interface {
	Discriminator() (string, error)
	ValueByDiscriminator() (interface{}, error)
}

func projectAndBodyByDiscriminator(c *gin.Context, body PolymorphicBody) (models.ProjectName, interface{}) {
	d, err := body.Discriminator()
	if err != nil {
		renderErrorResponse(c, "could not get project discriminator: %w", err)
	}

	b, err := body.ValueByDiscriminator()
	if err != nil {
		renderErrorResponse(c, "could not convert body: %w", err)
	}

	return models.ProjectName(d), b
}
