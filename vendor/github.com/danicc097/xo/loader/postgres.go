package loader

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/danicc097/xo/models"
	xo "github.com/danicc097/xo/types"
)

func init() {
	Register("postgres", Loader{
		Mask:             "$%d",
		Flags:            PostgresFlags,
		Schema:           models.PostgresSchema,
		Enums:            models.PostgresEnums,
		EnumValues:       models.PostgresEnumValues,
		Procs:            models.PostgresProcs,
		Constraints:      models.PostgresConstraints,
		ProcParams:       models.PostgresProcParams,
		Tables:           models.PostgresTables,
		TableColumns:     PostgresTableColumns,
		TableSequences:   models.PostgresTableSequences,
		TableGenerations: models.PostgresTableGenerations,
		TableForeignKeys: models.PostgresTableForeignKeys,
		TableIndexes:     models.PostgresTableIndexes,
		IndexColumns:     PostgresIndexColumns,
		ViewCreate:       models.PostgresViewCreate,
		ViewSchema:       models.PostgresViewSchema,
		ViewDrop:         models.PostgresViewDrop,
		ViewStrip:        PostgresViewStrip,
	})
}

// PostgresFlags returnss the postgres flags.
func PostgresFlags() []xo.Flag {
	return []xo.Flag{
		{
			ContextKey: oidsKey,
			Type:       "bool",
			Desc:       "enable postgres OIDs",
			Default:    "false",
		},
	}
}

// StdlibPostgresGoType parses a type into a Go type based on the databate type definition.
//
// For array types, it returns the standard go array ([]<type>).
func StdlibPostgresGoType(d xo.Type, schema, itype, _ string) (string, string, error) {
	goType, zero, err := PostgresGoType(d, schema, itype)
	if err != nil {
		return "", "", err
	}
	if d.IsArray {
		arrType, ok := pgStdArrMapping[goType]
		goType, zero = "[]byte", "nil"
		if ok {
			goType = arrType
		}
	}
	return goType, zero, nil
}

// PQPostgresGoType parses a type into a Go type based on the databate type definition.
//
// For array types, it returns the equivalent as defined in `github.com/lib/pq`.
func PQPostgresGoType(d xo.Type, schema, itype, _ string) (string, string, error) {
	goType, zero, err := PostgresGoType(d, schema, itype)
	if err != nil {
		return "", "", err
	}
	if d.IsArray {
		arrType, ok := pgxArrMapping[goType]
		goType, zero = "[]any", "[]any{}" // is of type struct { A any }; can't be nil
		if ok {
			goType, zero = arrType, "nil"
		}
	}
	return goType, zero, nil
}

// PostgresGoType parse a type into a Go type based on the database type
// definition.
func PostgresGoType(d xo.Type, schema, itype string) (string, string, error) {
	// SETOF -> []T
	if strings.HasPrefix(d.Type, "SETOF ") {
		d.Type = d.Type[len("SETOF "):]
		goType, _, err := PostgresGoType(d, schema, itype)
		if err != nil {
			return "", "", err
		}
		return "[]" + goType, "nil", nil
	}
	// If it's an array, the underlying type shouldn't also be set as an array
	typNullable := d.Nullable && !d.IsArray
	// special type handling
	typ := d.Type
	switch {
	case typ == `"char"`:
		typ = "char"
	case strings.HasPrefix(typ, "information_schema."):
		switch strings.TrimPrefix(typ, "information_schema.") {
		case "cardinal_number":
			typ = "integer"
		case "character_data", "sql_identifier", "yes_or_no":
			typ = "character varying"
		case "time_stamp":
			typ = "timestamp with time zone"
		}
	}
	var goType, zero string
	switch typ {
	case "boolean":
		goType, zero = "bool", "false"
		if typNullable {
			goType, zero = "*bool", "nil"
		}
	case "bpchar", "character varying", "character", "inet", "money", "text", "name":
		goType, zero = "string", `""`
		if typNullable {
			goType, zero = "*string", "nil"
		}
	case "smallint":
		goType, zero = "int16", "0"
		if typNullable {
			goType, zero = "*int16", "nil"
		}
	case "integer":
		goType, zero = itype, "0"
		if typNullable {
			goType, zero = "*int", "nil"
		}
	case "bigint":
		goType, zero = "int64", "0"
		if typNullable {
			goType, zero = "*int64", "nil"
		}
	case "real":
		goType, zero = "float32", "0.0"
		if typNullable {
			goType, zero = "*float32", "nil"
		}
	case "double precision", "numeric":
		goType, zero = "float64", "0.0"
		if typNullable {
			goType, zero = "*float64", "nil"
		}
		// TODO pgtypes for EACH different ttyype in this case
	case "date", "timestamp with time zone", "time with time zone", "time without time zone", "timestamp without time zone":
		goType, zero = "time.Time", "nil"
		if typNullable {
			goType, zero = "*time.Time", "nil"
		}
	case "bit":
		goType, zero = "uint8", "0"
		if typNullable {
			goType, zero = "*uint8", "nil"
		}
	case "any", "bit varying", "bytea", "interval", "xml":
		// TODO: write custom type for interval marshaling
		goType, zero = "[]byte", "nil"
	case "json":
		goType, zero = "pgtype.JSON", "nil"
	case "jsonb":
		goType, zero = "pgtype.JSONB", "nil"
	case "hstore":
		goType, zero = "hstore.Hstore", "nil"
	case "uuid":
		goType, zero = "uuid.UUID", "uuid.UUID{}"
		if typNullable {
			goType, zero = "*uuid.UUID", "nil"
		}
	default:
		goType, zero = schemaType(d.Type, typNullable, schema)
	}
	return goType, zero, nil
}

// PostgresTableColumns returns the columns for a table.
func PostgresTableColumns(ctx context.Context, db models.DB, schema string, table string) ([]*models.Column, error) {
	return models.PostgresTableColumns(ctx, db, schema, table, enableOids(ctx))
}

// PostgresIndexColumns returns the column list for an index.
//
// FIXME: rewrite this using SQL exclusively using OVER
func PostgresIndexColumns(ctx context.Context, db models.DB, schema string, table string, index string) ([]*models.IndexColumn, error) {
	// load columns
	cols, err := models.PostgresIndexColumns(ctx, db, schema, index)
	if err != nil {
		return nil, err
	}
	// load col order
	colOrd, err := models.PostgresGetColOrder(ctx, db, schema, index)
	if err != nil {
		return nil, err
	}
	// build schema name used in errors
	s := schema
	if s != "" {
		s += "."
	}
	// put cols in order using colOrder
	var ret []*models.IndexColumn
	for _, v := range strings.Split(colOrd.Ord, " ") {
		cid, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("could not convert %s%s index %s column %s to int", s, table, index, v)
		}
		// find column
		found := false
		var c *models.IndexColumn
		for _, ic := range cols {
			if cid == ic.Cid {
				found, c = true, ic
				break
			}
		}
		// sanity check
		if !found {
			return nil, fmt.Errorf("could not find %s%s index %s column id %d", s, table, index, cid)
		}
		ret = append(ret, c)
	}
	return ret, nil
}

// PostgresViewStrip strips '::type AS name' in queries.
func PostgresViewStrip(query, inspect []string) ([]string, []string, []string, error) {
	comments := make([]string, len(query))
	for i, line := range query {
		if pos := stripRE.FindStringIndex(line); pos != nil {
			query[i] = line[:pos[0]] + line[pos[1]:]
			comments[i] = line[pos[0]:pos[1]]
		}
	}
	return query, inspect, comments, nil
}

// stripRE is the regexp to match the '::type AS name' portion in a query,
// which is a quirk/requirement of generating queries for postgres.
var stripRE = regexp.MustCompile(`(?i)::[a-z][a-z0-9_\.]+\s+AS\s+[a-z][a-z0-9_\.]+`)

var pgStdArrMapping = map[string]string{
	"bool":    "[]bool",
	"[]byte":  "[][]byte",
	"float64": "[]float64",
	"float32": "[]float32",
	"int64":   "[]int64",
	"int32":   "[]int32",
	"string":  "[]string",
	// default: "[]byte"
}

var pgxArrMapping = map[string]string{
	"bool":    "[]bool",
	"[]byte":  "[][]byte",
	"float64": "[]float64",
	"float32": "[]float32",
	"int64":   "[]int64",
	"int32":   "[]int32",
	"string":  "[]string",
	// default: "pq.GenericArray"
}

// oidsKey is the oids context key.
const oidsKey xo.ContextKey = "oids"

// enableOids returns the enableOids value from the context.
func enableOids(ctx context.Context) bool {
	b, _ := ctx.Value(oidsKey).(bool)
	return b
}
