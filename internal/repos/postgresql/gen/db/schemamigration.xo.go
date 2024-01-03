package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// SchemaMigration represents a row from 'public.schema_migrations'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//     -- private to exclude a field from JSON.
//     -- not-required to make a schema field not required.
//     -- hidden to exclude field from OpenAPI generation.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type SchemaMigration struct {
	Version SchemaMigrationID `json:"version" db:"version" required:"true" nullable:"false"` // version
	Dirty   bool              `json:"dirty" db:"dirty" required:"true" nullable:"false"`     // dirty

}

// SchemaMigrationCreateParams represents insert params for 'public.schema_migrations'.
type SchemaMigrationCreateParams struct {
	Dirty   bool              `json:"dirty" required:"true" nullable:"false"`   // dirty
	Version SchemaMigrationID `json:"version" required:"true" nullable:"false"` // version
}

type SchemaMigrationID int

// CreateSchemaMigration creates a new SchemaMigration in the database with the given params.
func CreateSchemaMigration(ctx context.Context, db DB, params *SchemaMigrationCreateParams) (*SchemaMigration, error) {
	sm := &SchemaMigration{
		Dirty:   params.Dirty,
		Version: params.Version,
	}

	return sm.Insert(ctx, db)
}

// SchemaMigrationUpdateParams represents update params for 'public.schema_migrations'.
type SchemaMigrationUpdateParams struct {
	Dirty   *bool              `json:"dirty" nullable:"false"`   // dirty
	Version *SchemaMigrationID `json:"version" nullable:"false"` // version
}

// SetUpdateParams updates public.schema_migrations struct fields with the specified params.
func (sm *SchemaMigration) SetUpdateParams(params *SchemaMigrationUpdateParams) {
	if params.Dirty != nil {
		sm.Dirty = *params.Dirty
	}
	if params.Version != nil {
		sm.Version = *params.Version
	}
}

type SchemaMigrationSelectConfig struct {
	limit   string
	orderBy string
	joins   SchemaMigrationJoins
	filters map[string][]any
	having  map[string][]any
}
type SchemaMigrationSelectConfigOption func(*SchemaMigrationSelectConfig)

// WithSchemaMigrationLimit limits row selection.
func WithSchemaMigrationLimit(limit int) SchemaMigrationSelectConfigOption {
	return func(s *SchemaMigrationSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type SchemaMigrationOrderBy string

const ()

type SchemaMigrationJoins struct {
}

// WithSchemaMigrationJoin joins with the given tables.
func WithSchemaMigrationJoin(joins SchemaMigrationJoins) SchemaMigrationSelectConfigOption {
	return func(s *SchemaMigrationSelectConfig) {
		s.joins = SchemaMigrationJoins{}
	}
}

// WithSchemaMigrationFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithSchemaMigrationFilters(filters map[string][]any) SchemaMigrationSelectConfigOption {
	return func(s *SchemaMigrationSelectConfig) {
		s.filters = filters
	}
}

// WithSchemaMigrationHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	// filter a given aggregate of assigned users to return results where at least one of them has id of userId
//	filters := map[string][]any{
//	"$i = ANY(ARRAY_AGG(assigned_users_join.user_id))": {userId},
//	}
func WithSchemaMigrationHavingClause(conditions map[string][]any) SchemaMigrationSelectConfigOption {
	return func(s *SchemaMigrationSelectConfig) {
		s.having = conditions
	}
}

// Insert inserts the SchemaMigration to the database.
func (sm *SchemaMigration) Insert(ctx context.Context, db DB) (*SchemaMigration, error) {
	// insert (manual)
	sqlstr := `INSERT INTO public.schema_migrations (
	dirty, version
	) VALUES (
	$1, $2
	)
	 RETURNING * `
	// run
	logf(sqlstr, sm.Dirty, sm.Version)
	rows, err := db.Query(ctx, sqlstr, sm.Dirty, sm.Version)
	if err != nil {
		return nil, logerror(fmt.Errorf("SchemaMigration/Insert/db.Query: %w", &XoError{Entity: "Schema migration", Err: err}))
	}
	newsm, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[SchemaMigration])
	if err != nil {
		return nil, logerror(fmt.Errorf("SchemaMigration/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Schema migration", Err: err}))
	}
	*sm = newsm

	return sm, nil
}

// Update updates a SchemaMigration in the database.
func (sm *SchemaMigration) Update(ctx context.Context, db DB) (*SchemaMigration, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.schema_migrations SET 
	dirty = $1 
	WHERE version = $2 
	RETURNING * `
	// run
	logf(sqlstr, sm.Dirty, sm.Version)

	rows, err := db.Query(ctx, sqlstr, sm.Dirty, sm.Version)
	if err != nil {
		return nil, logerror(fmt.Errorf("SchemaMigration/Update/db.Query: %w", &XoError{Entity: "Schema migration", Err: err}))
	}
	newsm, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[SchemaMigration])
	if err != nil {
		return nil, logerror(fmt.Errorf("SchemaMigration/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Schema migration", Err: err}))
	}
	*sm = newsm

	return sm, nil
}

// Upsert upserts a SchemaMigration in the database.
// Requires appropriate PK(s) to be set beforehand.
func (sm *SchemaMigration) Upsert(ctx context.Context, db DB, params *SchemaMigrationCreateParams) (*SchemaMigration, error) {
	var err error

	sm.Dirty = params.Dirty
	sm.Version = params.Version

	sm, err = sm.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Schema migration", Err: err})
			}
			sm, err = sm.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Schema migration", Err: err})
			}
		}
	}

	return sm, err
}

// Delete deletes the SchemaMigration from the database.
func (sm *SchemaMigration) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.schema_migrations 
	WHERE version = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, sm.Version); err != nil {
		return logerror(err)
	}
	return nil
}

// SchemaMigrationPaginatedByVersion returns a cursor-paginated list of SchemaMigration.
func SchemaMigrationPaginatedByVersion(ctx context.Context, db DB, version SchemaMigrationID, direction models.Direction, opts ...SchemaMigrationSelectConfigOption) ([]SchemaMigration, error) {
	c := &SchemaMigrationSelectConfig{joins: SchemaMigrationJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	operator := "<"
	if direction == models.DirectionAsc {
		operator = ">"
	}

	sqlstr := fmt.Sprintf(`SELECT 
	schema_migrations.dirty,
	schema_migrations.version %s 
	 FROM public.schema_migrations %s 
	 WHERE schema_migrations.version %s $1
	 %s   %s 
  %s 
  ORDER BY 
		version %s `, selects, joins, operator, filters, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* SchemaMigrationPaginatedByVersion */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{version}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("SchemaMigration/Paginated/db.Query: %w", &XoError{Entity: "Schema migration", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[SchemaMigration])
	if err != nil {
		return nil, logerror(fmt.Errorf("SchemaMigration/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Schema migration", Err: err}))
	}
	return res, nil
}

// SchemaMigrationByVersion retrieves a row from 'public.schema_migrations' as a SchemaMigration.
//
// Generated from index 'schema_migrations_pkey'.
func SchemaMigrationByVersion(ctx context.Context, db DB, version SchemaMigrationID, opts ...SchemaMigrationSelectConfigOption) (*SchemaMigration, error) {
	c := &SchemaMigrationSelectConfig{joins: SchemaMigrationJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	schema_migrations.dirty,
	schema_migrations.version %s 
	 FROM public.schema_migrations %s 
	 WHERE schema_migrations.version = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* SchemaMigrationByVersion */\n" + sqlstr

	// run
	// logf(sqlstr, version)
	rows, err := db.Query(ctx, sqlstr, append([]any{version}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("schema_migrations/SchemaMigrationByVersion/db.Query: %w", &XoError{Entity: "Schema migration", Err: err}))
	}
	sm, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[SchemaMigration])
	if err != nil {
		return nil, logerror(fmt.Errorf("schema_migrations/SchemaMigrationByVersion/pgx.CollectOneRow: %w", &XoError{Entity: "Schema migration", Err: err}))
	}

	return &sm, nil
}
