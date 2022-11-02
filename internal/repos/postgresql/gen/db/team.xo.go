package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"
)

type TeamOrderBy = string

// Team represents a row from 'public.teams'.
type Team struct {
	TeamID      int       `json:"team_id"`     // team_id
	ProjectID   int       `json:"project_id"`  // project_id
	Name        string    `json:"name"`        // name
	Description string    `json:"description"` // description
	Metadata    []byte    `json:"metadata"`    // metadata
	CreatedAt   time.Time `json:"created_at"`  // created_at
	UpdatedAt   time.Time `json:"updated_at"`  // updated_at
	// xo fields
	_exists, _deleted bool
}

// TODO only create if exists
// GetMostRecentTeam returns n most recent rows from 'teams',
// ordered by "created_at" in descending order.
func GetMostRecentTeam(ctx context.Context, db DB, n int) ([]*Team, error) {
	// list
	const sqlstr = `SELECT ` +
		`team_id, project_id, name, description, metadata, created_at, updated_at ` +
		`FROM public.teams ` +
		`ORDER BY created_at DESC LIMIT $1`
	// run
	logf(sqlstr, n)

	rows, err := db.Query(ctx, sqlstr, n)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()

	// load results
	var res []*Team
	for rows.Next() {
		t := Team{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&t.TeamID, &t.ProjectID, &t.Name, &t.Description, &t.Metadata, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// Exists returns true when the Team exists in the database.
func (t *Team) Exists() bool {
	return t._exists
}

// Deleted returns true when the Team has been marked for deletion from
// the database.
func (t *Team) Deleted() bool {
	return t._deleted
}

// Insert inserts the Team to the database.
func (t *Team) Insert(ctx context.Context, db DB) error {
	switch {
	case t._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case t._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO public.teams (` +
		`project_id, name, description, metadata` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING team_id`
	// run
	logf(sqlstr, t.ProjectID, t.Name, t.Description, t.Metadata)
	if err := db.QueryRow(ctx, sqlstr, t.ProjectID, t.Name, t.Description, t.Metadata).Scan(&t.TeamID); err != nil {
		return logerror(err)
	}
	// set exists
	t._exists = true
	return nil
}

// Update updates a Team in the database.
func (t *Team) Update(ctx context.Context, db DB) error {
	switch {
	case !t._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case t._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.teams SET ` +
		`project_id = $1, name = $2, description = $3, metadata = $4 ` +
		`WHERE team_id = $5`
	// run
	logf(sqlstr, t.ProjectID, t.Name, t.Description, t.Metadata, t.CreatedAt, t.UpdatedAt, t.TeamID)
	if _, err := db.Exec(ctx, sqlstr, t.ProjectID, t.Name, t.Description, t.Metadata, t.CreatedAt, t.UpdatedAt, t.TeamID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Team to the database.
func (t *Team) Save(ctx context.Context, db DB) error {
	if t.Exists() {
		return t.Update(ctx, db)
	}
	return t.Insert(ctx, db)
}

// Upsert performs an upsert for Team.
func (t *Team) Upsert(ctx context.Context, db DB) error {
	switch {
	case t._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.teams (` +
		`team_id, project_id, name, description, metadata` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`)` +
		` ON CONFLICT (team_id) DO ` +
		`UPDATE SET ` +
		`project_id = EXCLUDED.project_id, name = EXCLUDED.name, description = EXCLUDED.description, metadata = EXCLUDED.metadata `
	// run
	logf(sqlstr, t.TeamID, t.ProjectID, t.Name, t.Description, t.Metadata)
	if _, err := db.Exec(ctx, sqlstr, t.TeamID, t.ProjectID, t.Name, t.Description, t.Metadata); err != nil {
		return logerror(err)
	}
	// set exists
	t._exists = true
	return nil
}

// Delete deletes the Team from the database.
func (t *Team) Delete(ctx context.Context, db DB) error {
	switch {
	case !t._exists: // doesn't exist
		return nil
	case t._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.teams ` +
		`WHERE team_id = $1`
	// run
	logf(sqlstr, t.TeamID)
	if _, err := db.Exec(ctx, sqlstr, t.TeamID); err != nil {
		return logerror(err)
	}
	// set deleted
	t._deleted = true
	return nil
}

// TeamByName retrieves a row from 'public.teams' as a Team.
//
// Generated from index 'teams_name_key'.
func TeamByName(ctx context.Context, db DB, name string) (*Team, error) {
	// query
	const sqlstr = `SELECT ` +
		`team_id, project_id, name, description, metadata, created_at, updated_at ` +
		`FROM public.teams ` +
		`WHERE name = $1`
	// run
	logf(sqlstr, name)
	t := Team{
		_exists: true,
	}
	if err := db.QueryRow(ctx, sqlstr, name).Scan(&t.TeamID, &t.ProjectID, &t.Name, &t.Description, &t.Metadata, &t.CreatedAt, &t.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &t, nil
}

// TeamByTeamID retrieves a row from 'public.teams' as a Team.
//
// Generated from index 'teams_pkey'.
func TeamByTeamID(ctx context.Context, db DB, teamID int) (*Team, error) {
	// query
	const sqlstr = `SELECT ` +
		`team_id, project_id, name, description, metadata, created_at, updated_at ` +
		`FROM public.teams ` +
		`WHERE team_id = $1`
	// run
	logf(sqlstr, teamID)
	t := Team{
		_exists: true,
	}
	if err := db.QueryRow(ctx, sqlstr, teamID).Scan(&t.TeamID, &t.ProjectID, &t.Name, &t.Description, &t.Metadata, &t.CreatedAt, &t.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &t, nil
}

// Project returns the Project associated with the Team's (ProjectID).
//
// Generated from foreign key 'teams_project_id_fkey'.
func (t *Team) Project(ctx context.Context, db DB) (*Project, error) {
	return ProjectByProjectID(ctx, db, t.ProjectID)
}
