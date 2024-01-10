delete_method=$(test -n "$has_deleted_at" && echo "SoftDelete" || echo "Delete")

# shellcheck disable=SC2028,SC2154
echo "package postgresql

import (
	\"context\"
	\"fmt\"

	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos\"
	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db\"
)

// ${pascal_name} represents the repository used for interacting with ${sentence_name} records.
type ${pascal_name} struct {
	q db.Querier
}

// New${pascal_name} instantiates the ${sentence_name} repository.
func New${pascal_name}() *${pascal_name} {
	return &${pascal_name}{
		q: NewQuerierWrapper(db.New()),
	}
}

var _ repos.${pascal_name} = (*${pascal_name})(nil)

func (t *${pascal_name}) Create(ctx context.Context, d db.DBTX, params *db.${pascal_name}CreateParams) (*db.${pascal_name}, error) {
	${camel_name}, err := db.Create${pascal_name}(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf(\"could not create ${camel_name}: %w\", ParseDBErrorDetail(err))
	}

	return ${camel_name}, nil
}

func (t *${pascal_name}) Update(ctx context.Context, d db.DBTX, id db.${pascal_name}ID, params *db.${pascal_name}UpdateParams) (*db.${pascal_name}, error) {
	${camel_name}, err := t.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf(\"could not get ${sentence_name} by id %w\", ParseDBErrorDetail(err))
	}

	${camel_name}.SetUpdateParams(params)

	${camel_name}, err = ${camel_name}.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf(\"could not update ${sentence_name}: %w\", ParseDBErrorDetail(err))
	}

	return ${camel_name}, err
}

func (t *${pascal_name}) ByID(ctx context.Context, d db.DBTX, id db.${pascal_name}ID, opts ...db.${pascal_name}SelectConfigOption) (*db.${pascal_name}, error) {
	${camel_name}, err := db.${pascal_name}By${pascal_name}ID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf(\"could not get ${sentence_name}: %w\", ParseDBErrorDetail(err))
	}

	return ${camel_name}, nil
}

func (t *${pascal_name}) Delete(ctx context.Context, d db.DBTX, id db.${pascal_name}ID) (*db.${pascal_name}, error) {
	${camel_name} := &db.${pascal_name}{
		${pascal_name}ID: id,
	}

	err := ${camel_name}.${delete_method}(ctx, d) // use SoftDelete if a deleted_at column exists.
	if err != nil {
		return nil, fmt.Errorf(\"could not delete ${sentence_name}: %w\", ParseDBErrorDetail(err))
	}

	return ${camel_name}, err
}
"
