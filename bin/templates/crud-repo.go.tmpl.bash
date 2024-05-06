#!/usr/bin/env bash

delete_method=$(test -n "$has_deleted_at" && echo "SoftDelete" || echo "Delete")

# shellcheck disable=SC2028,SC2154
cat <<EOF
package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
  "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"

)

// ${pascal_name} represents the repository used for interacting with ${sentence_name} records.
type ${pascal_name} struct {
	q models.Querier
}

// New${pascal_name} instantiates the ${sentence_name} repository.
func New${pascal_name}() *${pascal_name} {
	return &${pascal_name}{
		q: NewQuerierWrapper(models.New()),
	}
}

var _ repos.${pascal_name} = (*${pascal_name})(nil)

func (t *${pascal_name}) Create(ctx context.Context, d models.DBTX, params *models.${pascal_name}CreateParams) (*models.${pascal_name}, error) {
	${camel_name}, err := models.Create${pascal_name}(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not create ${camel_name}: %w", ParseDBErrorDetail(err))
	}

	return ${camel_name}, nil
}

func (t *${pascal_name}) Update(ctx context.Context, d models.DBTX, id models.${pascal_name}ID, params *models.${pascal_name}UpdateParams) (*models.${pascal_name}, error) {
	${camel_name}, err := t.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get ${sentence_name} by id %w", ParseDBErrorDetail(err))
	}

	${camel_name}.SetUpdateParams(params)

	${camel_name}, err = ${camel_name}.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update ${sentence_name}: %w", ParseDBErrorDetail(err))
	}

	return ${camel_name}, err
}

func (t *${pascal_name}) ByID(ctx context.Context, d models.DBTX, id models.${pascal_name}ID, opts ...models.${pascal_name}SelectConfigOption) (*models.${pascal_name}, error) {
	${camel_name}, err := models.${pascal_name}By${pascal_name}ID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get ${sentence_name}: %w", ParseDBErrorDetail(err))
	}

	return ${camel_name}, nil
}

func (t *${pascal_name}) Delete(ctx context.Context, d models.DBTX, id models.${pascal_name}ID) (*models.${pascal_name}, error) {
	${camel_name} := &models.${pascal_name}{
		${pascal_name}ID: id,
	}

	err := ${camel_name}.${delete_method}(ctx, d) // use SoftDelete if a deleted_at column exists.
	if err != nil {
		return nil, fmt.Errorf("could not delete ${sentence_name}: %w", ParseDBErrorDetail(err))
	}

	return ${camel_name}, err
}
EOF
