#!/usr/bin/env bash

create_params="$(test -n "$with_project" && echo "projectID db.ProjectID")"
create_args="$(test -n "$with_project" && echo "projectID")"

cat <<EOF
package postgresqlrandom

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// NOTE: FKs should always be passed explicitly.
func ${pascal_name}CreateParams($create_params) *db.${pascal_name}CreateParams {
	return &db.${pascal_name}CreateParams{
		// TODO: fill in with testutil randomizer helpers or add parameters accordingly
$(test -n "$with_project" && echo "		ProjectID: projectID,")
	}
}
EOF
