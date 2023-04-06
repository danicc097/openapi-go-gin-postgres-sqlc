package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// SchemaMigration represents a row from 'public.schema_migrations'.
// Include "property:private" in a SQL column comment to exclude a field from JSON.
type SchemaMigration struct {
	Version int64 `json:"version" db:"version"` // version
	Dirty   bool  `json:"dirty" db:"dirty"`     // dirty

	// xo fields
	_exists, _deleted bool
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

// Exists returns true when the SchemaMigration exists in the database.
func (sm *SchemaMigration) Exists() bool {
	return sm._exists
}

// Deleted returns true when the SchemaMigration has been marked for deletion from
// the database.
func (sm *SchemaMigration) Deleted() bool {
	return sm._deleted
}

// Insert inserts the SchemaMigration to the database.

func (sm *SchemaMigration) Insert(ctx context.Context, db DB) (*SchemaMigration, error) {
	switch {
	case sm._exists: // already exists
		return nil, logerror(&ErrInsertFailed{ErrAlreadyExists})
	case sm._deleted: // deleted
		return nil, logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	sqlstr := `INSERT INTO public.schema_migrations (` +
		`version, dirty` +
		`) VALUES (` +
		`$1, $2` +
		`) `
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
	newsm._exists = true
	sm = &newsm

	return sm, nil
}

// Update updates a SchemaMigration in the database.
func (sm *SchemaMigration) Update(ctx context.Context, db DB) (*SchemaMigration, error) {
	switch {
	case !sm._exists: // doesn't exist
		return nil, logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case sm._deleted: // deleted
		return nil, logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
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
	newsm._exists = true
	sm = &newsm

	return sm, nil
}

// Save saves the SchemaMigration to the database.
func (sm *SchemaMigration) Save(ctx context.Context, db DB) (*SchemaMigration, error) {
	if sm.Exists() {
		return sm.Update(ctx, db)
	}
	return sm.Insert(ctx, db)
}

// Upsert performs an upsert for SchemaMigration.
func (sm *SchemaMigration) Upsert(ctx context.Context, db DB) error {
	switch {
	case sm._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.schema_migrations (` +
		`version, dirty` +
		`) VALUES (` +
		`$1, $2` +
		`)` +
		` ON CONFLICT (version) DO ` +
		`UPDATE SET ` +
		`dirty = EXCLUDED.dirty  `
	// run
	logf(sqlstr, sm.Version, sm.Dirty)
	if _, err := db.Exec(ctx, sqlstr, sm.Version, sm.Dirty); err != nil {
		return logerror(err)
	}
	// set exists
	sm._exists = true
	return nil
}

// Delete deletes the SchemaMigration from the database.
func (sm *SchemaMigration) Delete(ctx context.Context, db DB) error {
	switch {
	case !sm._exists: // doesn't exist
		return nil
	case sm._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.schema_migrations ` +
		`WHERE version = $1 `
	// run
	logf(sqlstr, sm.Version)
	if _, err := db.Exec(ctx, sqlstr, sm.Version); err != nil {
		return logerror(err)
	}
	// set deleted
	sm._deleted = true
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
	logf(sqlstr, version)
	rows, err := db.Query(ctx, sqlstr, version)
	if err != nil {
		return nil, logerror(fmt.Errorf("schema_migrations/SchemaMigrationByVersion/db.Query: %w", err))
	}
	sm, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[SchemaMigration])
	if err != nil {
		return nil, logerror(fmt.Errorf("schema_migrations/SchemaMigrationByVersion/pgx.CollectOneRow: %w", err))
	}
	sm._exists = true
	return &sm, nil
}
