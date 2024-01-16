create_params="$(test -n "$with_project" && echo ", projectID db.ProjectID")"
create_args="$(test -n "$with_project" && echo ", projectID")"

# shellcheck disable=SC2028,SC2154
echo "package postgresqltestutil

import (
	\"context\"
	\"testing\"

	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql\"
	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db\"
	\"github.com/stretchr/testify/require\"
  \"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/reposwrappers\"
	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil\"

)

// NOTE: FKs should always be passed explicitly.
func Random${pascal_name}CreateParams(t *testing.T $create_params) *db.${pascal_name}CreateParams {
	t.Helper()

	return &db.${pascal_name}CreateParams{
		// TODO: fill in with testutil randomizer helpers or add parameters accordingly
$(test -n "$with_project" && echo "		ProjectID: projectID,")
	}
}
"
