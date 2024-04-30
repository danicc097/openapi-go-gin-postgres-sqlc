#!/usr/bin/env bash

create_args="$(test -n "$with_project" && echo "params.ProjectID")"

# shellcheck disable=SC2028,SC2154
cat <<EOF
package servicetestutil

import (
	"context"

$(test -n "$has_deleted_at" && echo "	\"time\"")
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqlrandom"
  "github.com/stretchr/testify/require"
)

type Create${pascal_name}Params struct {
$(test -n "$with_project" && echo "	ProjectID  db.ProjectID")
$(test -n "$has_deleted_at" && echo "	// DeletedAt allows returning a soft deleted ${sentence_name} when a deleted_at column exists.
	// Note that the service Delete call should make use of the SoftDelete method.
	DeletedAt  *time.Time")
}

type Create${pascal_name}Fixture struct {
	${pascal_name}   *db.${pascal_name}
}

// Create${pascal_name} creates a new random ${sentence_name} with the given configuration.
func (ff *FixtureFactory) Create${pascal_name}(ctx context.Context, params Create${pascal_name}Params) *Create${pascal_name}Fixture {
	randomRepoCreateParams := postgresqlrandom.${pascal_name}CreateParams($create_args)
	// don't use repos for test fixtures, use service logic
	${camel_name}, err := ff.svc.${pascal_name}.Create(ctx, ff.d, randomRepoCreateParams)
	require.NoError(ff.t, err)

$(test -n "$has_deleted_at" && echo "
	if params.DeletedAt != nil {
		${camel_name}, err = ff.svc.${pascal_name}.Delete(ctx, ff.d, ${camel_name}.${pascal_name}ID)
	  require.NoError(ff.t, err)
	}
")

	return &Create${pascal_name}Fixture{
		${pascal_name}:   ${camel_name},
	}
}

EOF
