#!/bin/bash

# shellcheck disable=SC2028,SC2154
cat <<EOF
package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
)

type ${pascal_name} struct {
	logger *zap.SugaredLogger
	repos  *repos.Repos
}

// New${pascal_name} returns a new ${sentence_name} service.
func New${pascal_name}(logger *zap.SugaredLogger, repos *repos.Repos) *${pascal_name} {
	return &${pascal_name}{
		logger: logger,
		repos:  repos,
	}
}

// ByID gets a ${sentence_name} by ID.
func (t *${pascal_name}) ByID(ctx context.Context, d db.DBTX, id db.${pascal_name}ID) (*db.${pascal_name}, error) {
	defer newOTelSpan().Build(ctx).End()

	${camel_name}, err := t.repos.${pascal_name}.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.${pascal_name}.ByID: %w", err)
	}

	return ${camel_name}, nil
}

// Create creates a new ${sentence_name}.
func (t *${pascal_name}) Create(ctx context.Context, d db.DBTX, params *db.${pascal_name}CreateParams) (*db.${pascal_name}, error) {
	defer newOTelSpan().Build(ctx).End()

	${camel_name}, err := t.repos.${pascal_name}.Create(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("repos.${pascal_name}.Create: %w", err)
	}

	return ${camel_name}, nil
}

// Update updates an existing ${sentence_name}.
func (t *${pascal_name}) Update(ctx context.Context, d db.DBTX, id db.${pascal_name}ID, params *db.${pascal_name}UpdateParams) (*db.${pascal_name}, error) {
	defer newOTelSpan().Build(ctx).End()

	${camel_name}, err := t.repos.${pascal_name}.Update(ctx, d, id, params)
	if err != nil {
		return nil, fmt.Errorf("repos.${pascal_name}.Update: %w", err)
	}

	return ${camel_name}, nil
}

// Delete deletes an existing ${sentence_name}.
func (t *${pascal_name}) Delete(ctx context.Context, d db.DBTX, id db.${pascal_name}ID) (*db.${pascal_name}, error) {
	defer newOTelSpan().Build(ctx).End()

	${camel_name}, err := t.repos.${pascal_name}.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.${pascal_name}.Delete: %w", err)
	}

	return ${camel_name}, nil
}
EOF
