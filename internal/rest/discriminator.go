package rest

import (
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

type PolymorphicBody interface {
	Discriminator() (string, error)
}

func projectByDiscriminator(b PolymorphicBody) (models.ProjectName, error) {
	d, err := b.Discriminator()
	if err != nil {
		return "", fmt.Errorf("could not get project discriminator: %w", err)
	}

	return models.ProjectName(d), nil
}
