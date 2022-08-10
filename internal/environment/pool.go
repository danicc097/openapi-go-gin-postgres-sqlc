package environment

import "github.com/jackc/pgx/v4/pgxpool"

// the alternative is to generate services by reparsing and editing the generated
// files again...
// we already have api_{tag}.go files, a {tag}Service struct
// that has logger, conn pool, etc. and implements all methods isn't too hard,
// but would need rethinking how non-implemented handlers are created,
// in this case we will have to
// parse existing handlers/api_*.go, not just render a brand new template,
// and remove unexisting handlers, dont touch the existing ones then append
// the new unimplemented methods.
// We will also need to regenerate, or completely rethink routes.go to initialize
// everything based on each service and handler...
// Perhaps some globals here and there wont hurt as much.
//nolint: gochecknoglobals
var Pool *pgxpool.Pool
