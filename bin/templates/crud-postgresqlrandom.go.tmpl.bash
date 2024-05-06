#!/usr/bin/env bash

create_params="$(test -n "$with_project" && echo "projectID models.ProjectID")"
create_args="$(test -n "$with_project" && echo "projectID")"

cat <<EOF
package postgresqlrandom

import (
  "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

// NOTE: FKs should always be passed explicitly.
func ${pascal_name}CreateParams($create_params) *models.${pascal_name}CreateParams {
	return &models.${pascal_name}CreateParams{
		// TODO: fill in with testutil randomizer helpers or add parameters accordingly
$(test -n "$with_project" && echo "		ProjectID: projectID,")
	}
}
EOF
