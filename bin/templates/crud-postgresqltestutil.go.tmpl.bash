echo "package postgresqltestutil

import (
	\"context\"
	\"testing\"

	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql\"
	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db\"
	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil\"
	\"github.com/stretchr/testify/require\"
)

func NewRandom${pascal_name}(t *testing.T, d db.DBTX, projectID db.ProjectID) (*db.${pascal_name}, error) {
	t.Helper()

	${camel_name}Repo := postgresql.New${pascal_name}()

	ucp := Random${pascal_name}CreateParams(t, projectID)

	${camel_name}, err := ${camel_name}Repo.Create(context.Background(), d, ucp)
	require.NoError(t, err, \"failed to create random entity\") // IMPORTANT: must fail. If testing actual failures use random create params instead

	return ${camel_name}, nil
}

func Random${pascal_name}CreateParams(t *testing.T, projectID db.ProjectID) *db.${pascal_name}CreateParams {
	t.Helper()

	return &db.${pascal_name}CreateParams{
		// TODO: fill in with testutil randomizer helpers
	}
}
"
