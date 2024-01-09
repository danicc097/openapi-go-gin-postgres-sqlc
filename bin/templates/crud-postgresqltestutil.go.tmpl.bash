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
)

func NewRandom${pascal_name}(t *testing.T, d db.DBTX $create_params) *db.${pascal_name} {
	t.Helper()

	${camel_name}Repo := postgresql.New${pascal_name}()

	ucp := Random${pascal_name}CreateParams(t $create_args)

	${camel_name}, err := ${camel_name}Repo.Create(context.Background(), d, ucp)
	require.NoError(t, err, \"failed to create random entity\") // IMPORTANT: must fail. If testing actual failures use random create params instead

	return ${camel_name}
}

func Random${pascal_name}CreateParams(t *testing.T $create_params) *db.${pascal_name}CreateParams {
	t.Helper()

	return &db.${pascal_name}CreateParams{
		// TODO: fill in with testutil randomizer helpers or add parameters accordingly
$(test -n "$with_project" && echo "		ProjectID: projectID,")
	}
}
"
