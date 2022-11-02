package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

type SchemaMigrationOrderBy = string

// SchemaMigration represents a row from 'public.schema_migrations'.
type SchemaMigration struct {
	Version int64 `json:"version"` // version
	Dirty   bool  `json:"dirty"`   // dirty
	// xo fields
	_exists, _deleted bool
}

// TODO only create if exists
// GetMostRecentSchemaMigration returns n most recent rows from 'schema_migrations',
// ordered by "created_at" in descending order.
func GetMostRecentSchemaMigration(ctx context.Context, db DB, n int) ([]*SchemaMigration, error) {
	// list
	const sqlstr = `SELECT ` +
		`version, dirty ` +
		`FROM public.schema_migrations ` +
		`ORDER BY created_at DESC LIMIT $1`
	// run
	logf(sqlstr, n)

	rows, err := db.Query(ctx, sqlstr, n)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()

	// load results
	var res []*SchemaMigration
	for rows.Next() {
		sm := SchemaMigration{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&sm.Version, &sm.Dirty); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &sm)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
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
func (sm *SchemaMigration) Insert(ctx context.Context, db DB) error {
	switch {
	case sm._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case sm._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.schema_migrations (` +
		`version, dirty` +
		`) VALUES (` +
		`$1, $2` +
		`)`
	// run
	logf(sqlstr, sm.Version, sm.Dirty)
	if _, err := db.Exec(ctx, sqlstr, sm.Version, sm.Dirty); err != nil {
		return logerror(err)
	}
	// set exists
	sm._exists = true
	return nil
}

// Update updates a SchemaMigration in the database.
func (sm *SchemaMigration) Update(ctx context.Context, db DB) error {
	switch {
	case !sm._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case sm._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.schema_migrations SET ` +
		`dirty = $1 ` +
		`WHERE version = $2`
	// run
	logf(sqlstr, sm.Dirty, sm.Version)
	if _, err := db.Exec(ctx, sqlstr, sm.Dirty, sm.Version); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the SchemaMigration to the database.
func (sm *SchemaMigration) Save(ctx context.Context, db DB) error {
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
	const sqlstr = `INSERT INTO public.schema_migrations (` +
		`version, dirty` +
		`) VALUES (` +
		`$1, $2` +
		`)` +
		` ON CONFLICT (version) DO ` +
		`UPDATE SET ` +
		`dirty = EXCLUDED.dirty `
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
	const sqlstr = `DELETE FROM public.schema_migrations ` +
		`WHERE version = $1`
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
func SchemaMigrationByVersion(ctx context.Context, db DB, version int64) (*SchemaMigration, error) {
	// query
	const sqlstr = `SELECT ` +
		`version, dirty ` +
		`FROM public.schema_migrations ` +
		`WHERE version = $1`
	// run
	logf(sqlstr, version)
	sm := SchemaMigration{
		_exists: true,
	}
	if err := db.QueryRow(ctx, sqlstr, version).Scan(&sm.Version, &sm.Dirty); err != nil {
		return nil, logerror(err)
	}
	return &sm, nil
}
