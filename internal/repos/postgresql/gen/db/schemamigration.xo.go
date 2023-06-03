package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// SchemaMigration represents a row from 'public.schema_migrations'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":private to exclude a field from JSON.
//   - "type":<pkg.type> to override the type annotation.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type SchemaMigration struct {
	Version int64 `json:"version" db:"version" required:"true"` // version
	Dirty   bool  `json:"dirty" db:"dirty" required:"true"`     // dirty

}

// SchemaMigrationCreateParams represents insert params for 'public.schema_migrations'.
type SchemaMigrationCreateParams struct {
	Version int64 `json:"version" required:"true"` // version
	Dirty   bool  `json:"dirty" required:"true"`   // dirty
}

// CreateSchemaMigration creates a new SchemaMigration in the database with the given params.
func CreateSchemaMigration(ctx context.Context, db DB, params *SchemaMigrationCreateParams) (*SchemaMigration, error) {
	sm := &SchemaMigration{
		Version: params.Version,
		Dirty:   params.Dirty,
	}

	return sm.Insert(ctx, db)
}

// SchemaMigrationUpdateParams represents update params for 'public.schema_migrations'.
type SchemaMigrationUpdateParams struct {
	Version *int64 `json:"version" required:"true"` // version
	Dirty   *bool  `json:"dirty" required:"true"`   // dirty
}

// SetUpdateParams updates public.schema_migrations struct fields with the specified params.
func (sm *SchemaMigration) SetUpdateParams(params *SchemaMigrationUpdateParams) {
	if params.Version != nil {
		sm.Version = *params.Version
	}
	if params.Dirty != nil {
		sm.Dirty = *params.Dirty
	}
}

type SchemaMigrationSelectConfig struct {
	limit   string
	orderBy string
	joins   SchemaMigrationJoins
	filters map[string][]any
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

// WithSchemaMigrationFilters adds the given filters, which may be parameterized with $i.
// Filters are joined with AND.
// NOTE: SQL injection prone.
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

// Insert inserts the SchemaMigration to the database.
func (sm *SchemaMigration) Insert(ctx context.Context, db DB) (*SchemaMigration, error) {
	// insert (manual)
	sqlstr := `INSERT INTO public.schema_migrations (` +
		`version, dirty` +
		`) VALUES (` +
		`$1, $2` +
		`)` +
		` RETURNING * `
	// run
	logf(sqlstr, sm.Version, sm.Dirty)
	rows, err := db.Query(ctx, sqlstr, sm.Version, sm.Dirty)
	if err != nil {
		return nil, logerror(fmt.Errorf("SchemaMigration/Insert/db.Query: %w", err))
	}
	newsm, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[SchemaMigration])
	if err != nil {
		return nil, logerror(fmt.Errorf("SchemaMigration/Insert/pgx.CollectOneRow: %w", err))
	}
	*sm = newsm

	return sm, nil
}

// Update updates a SchemaMigration in the database.
func (sm *SchemaMigration) Update(ctx context.Context, db DB) (*SchemaMigration, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.schema_migrations SET ` +
		`dirty = $1 ` +
		`WHERE version = $2 ` +
		`RETURNING * `
	// run
	logf(sqlstr, sm.Dirty, sm.Version)

	rows, err := db.Query(ctx, sqlstr, sm.Dirty, sm.Version)
	if err != nil {
		return nil, logerror(fmt.Errorf("SchemaMigration/Update/db.Query: %w", err))
	}
	newsm, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[SchemaMigration])
	if err != nil {
		return nil, logerror(fmt.Errorf("SchemaMigration/Update/pgx.CollectOneRow: %w", err))
	}
	*sm = newsm

	return sm, nil
}

// Upsert upserts a SchemaMigration in the database.
// Requires appropiate PK(s) to be set beforehand.
func (sm *SchemaMigration) Upsert(ctx context.Context, db DB, params *SchemaMigrationCreateParams) (*SchemaMigration, error) {
	var err error

	sm.Version = params.Version
	sm.Dirty = params.Dirty

	sm, err = sm.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", err)
			}
			sm, err = sm.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", err)
			}
		}
	}

	return sm, err
}

// Delete deletes the SchemaMigration from the database.
func (sm *SchemaMigration) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.schema_migrations ` +
		`WHERE version = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, sm.Version); err != nil {
		return logerror(err)
	}
	return nil
}

// SchemaMigrationPaginatedByVersionAsc returns a cursor-paginated list of SchemaMigration in Asc order.
func SchemaMigrationPaginatedByVersionAsc(ctx context.Context, db DB, version int64, opts ...SchemaMigrationSelectConfigOption) ([]SchemaMigration, error) {
	c := &SchemaMigrationSelectConfig{joins: SchemaMigrationJoins{}, filters: make(map[string][]any)}

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

	sqlstr := fmt.Sprintf(`SELECT `+
		`schema_migrations.version,
schema_migrations.dirty %s `+
		`FROM public.schema_migrations %s `+
		` WHERE schema_migrations.version > $1`+
		` %s   %s 
  ORDER BY 
		version Asc`, selects, joins, filters, groupbys)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{version}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("SchemaMigration/Paginated/Asc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[SchemaMigration])
	if err != nil {
		return nil, logerror(fmt.Errorf("SchemaMigration/Paginated/Asc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// SchemaMigrationPaginatedByVersionDesc returns a cursor-paginated list of SchemaMigration in Desc order.
func SchemaMigrationPaginatedByVersionDesc(ctx context.Context, db DB, version int64, opts ...SchemaMigrationSelectConfigOption) ([]SchemaMigration, error) {
	c := &SchemaMigrationSelectConfig{joins: SchemaMigrationJoins{}, filters: make(map[string][]any)}

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

	sqlstr := fmt.Sprintf(`SELECT `+
		`schema_migrations.version,
schema_migrations.dirty %s `+
		`FROM public.schema_migrations %s `+
		` WHERE schema_migrations.version < $1`+
		` %s   %s 
  ORDER BY 
		version Desc`, selects, joins, filters, groupbys)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{version}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("SchemaMigration/Paginated/Desc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[SchemaMigration])
	if err != nil {
		return nil, logerror(fmt.Errorf("SchemaMigration/Paginated/Desc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// SchemaMigrationByVersion retrieves a row from 'public.schema_migrations' as a SchemaMigration.
//
// Generated from index 'schema_migrations_pkey'.
func SchemaMigrationByVersion(ctx context.Context, db DB, version int64, opts ...SchemaMigrationSelectConfigOption) (*SchemaMigration, error) {
	c := &SchemaMigrationSelectConfig{joins: SchemaMigrationJoins{}, filters: make(map[string][]any)}

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

	sqlstr := fmt.Sprintf(`SELECT `+
		`schema_migrations.version,
schema_migrations.dirty %s `+
		`FROM public.schema_migrations %s `+
		` WHERE schema_migrations.version = $1`+
		` %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, version)
	rows, err := db.Query(ctx, sqlstr, append([]any{version}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("schema_migrations/SchemaMigrationByVersion/db.Query: %w", err))
	}
	sm, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[SchemaMigration])
	if err != nil {
		return nil, logerror(fmt.Errorf("schema_migrations/SchemaMigrationByVersion/pgx.CollectOneRow: %w", err))
	}

	return &sm, nil
}
