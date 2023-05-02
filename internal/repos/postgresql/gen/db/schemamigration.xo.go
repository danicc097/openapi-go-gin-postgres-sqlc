package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// SchemaMigration represents a row from 'public.schema_migrations'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|O2M|M2O|M2M" to generate joins (not executed by default).
type SchemaMigration struct {
	Version int64 `json:"version" db:"version" required:"true"` // version
	Dirty   bool  `json:"dirty" db:"dirty" required:"true"`     // dirty

}

// SchemaMigrationCreateParams represents insert params for 'public.schema_migrations'
type SchemaMigrationCreateParams struct {
	Version int64 `json:"version"` // version
	Dirty   bool  `json:"dirty"`   // dirty
}

func NewSchemaMigration(params *SchemaMigrationCreateParams) *SchemaMigration {
	return &SchemaMigration{
		Version: params.Version,
		Dirty:   params.Dirty,
	}
}

// SchemaMigrationUpdateParams represents update params for 'public.schema_migrations'
type SchemaMigrationUpdateParams struct {
	Version *int64 `json:"version"` // version
	Dirty   *bool  `json:"dirty"`   // dirty
}

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
}
type SchemaMigrationSelectConfigOption func(*SchemaMigrationSelectConfig)

// WithSchemaMigrationLimit limits row selection.
func WithSchemaMigrationLimit(limit int) SchemaMigrationSelectConfigOption {
	return func(s *SchemaMigrationSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type SchemaMigrationOrderBy = string

const ()

type SchemaMigrationJoins struct {
}

// WithSchemaMigrationJoin joins with the given tables.
func WithSchemaMigrationJoin(joins SchemaMigrationJoins) SchemaMigrationSelectConfigOption {
	return func(s *SchemaMigrationSelectConfig) {
		s.joins = joins
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

// Upsert performs an upsert for SchemaMigration.
func (sm *SchemaMigration) Upsert(ctx context.Context, db DB) error {
	// upsert
	sqlstr := `INSERT INTO public.schema_migrations (` +
		`version, dirty` +
		`) VALUES (` +
		`$1, $2` +
		`)` +
		` ON CONFLICT (version) DO ` +
		`UPDATE SET ` +
		`dirty = EXCLUDED.dirty ` +
		` RETURNING * `
	// run
	logf(sqlstr, sm.Version, sm.Dirty)
	if _, err := db.Exec(ctx, sqlstr, sm.Version, sm.Dirty); err != nil {
		return logerror(err)
	}
	// set exists
	return nil
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

// SchemaMigrationByVersion retrieves a row from 'public.schema_migrations' as a SchemaMigration.
//
// Generated from index 'schema_migrations_pkey'.
func SchemaMigrationByVersion(ctx context.Context, db DB, version int64, opts ...SchemaMigrationSelectConfigOption) (*SchemaMigration, error) {
	c := &SchemaMigrationSelectConfig{joins: SchemaMigrationJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`schema_migrations.version,
schema_migrations.dirty ` +
		`FROM public.schema_migrations ` +
		`` +
		` WHERE schema_migrations.version = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, version)
	rows, err := db.Query(ctx, sqlstr, version)
	if err != nil {
		return nil, logerror(fmt.Errorf("schema_migrations/SchemaMigrationByVersion/db.Query: %w", err))
	}
	sm, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[SchemaMigration])
	if err != nil {
		return nil, logerror(fmt.Errorf("schema_migrations/SchemaMigrationByVersion/pgx.CollectOneRow: %w", err))
	}

	return &sm, nil
}
