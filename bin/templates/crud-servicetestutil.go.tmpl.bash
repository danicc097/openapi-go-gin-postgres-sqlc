echo "package servicetestutil

import (
	\"context\"
	\"fmt\"
	\"time\"

	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db\"
	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil\"
)

type Create${pascal_name}Params struct {
	ProjectID  db.ProjectID
	$([[ -z "$has_deleted_at" ]] && echo "// DeletedAt allows returning a soft deleted ${sentence_name} when a deleted_at column exists.
	// Note that the service Delete call should make use of the SoftDelete method.
	DeletedAt  *time.Time")
}

type Create${pascal_name}Fixture struct {
	${pascal_name}   *db.${pascal_name}
}

// Create${pascal_name} creates a new random ${sentence_name} with the given configuration.
func (ff *FixtureFactory) Create${pascal_name}(ctx context.Context, params Create${pascal_name}Params) (*Create${pascal_name}Fixture, error) {
	randomRepoCreateParams := postgresqltestutil.Random${pascal_name}CreateParams(ff.t) // , params.ProjectID
	// don't use repos for tests
	${camel_name}, err := ff.svc.${pascal_name}.Create(ctx, ff.db, randomRepoCreateParams)
	if err != nil {
		return nil, fmt.Errorf(\"svc.${pascal_name}.Create: %w\", err)
	}

$([[ -z "$has_deleted_at" ]] && echo "
	if params.DeletedAt != nil {
		${camel_name}, err = ff.svc.${pascal_name}.Delete(ctx, ff.db, ${camel_name}.${pascal_name}ID)
		if err != nil {
			return nil, fmt.Errorf(\"svc.${pascal_name}.Delete: %w\", err)
		}
	}
")

	return &Create${pascal_name}Fixture{
		${pascal_name}:   ${camel_name},
	}, nil
}

"
