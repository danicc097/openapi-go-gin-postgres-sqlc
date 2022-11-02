package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
)

type KanbanStepSelectConfig struct {
	limit    string
	orderBy  string
	joinWith []KanbanStepJoinBy
}

type KanbanStepSelectConfigOption func(*KanbanStepSelectConfig)

// KanbanStepWithLimit limits row selection.
func KanbanStepWithLimit(limit int) KanbanStepSelectConfigOption {
	return func(s *KanbanStepSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

// KanbanStepWithOrderBy orders results by the given columns.
func KanbanStepWithOrderBy(rows ...KanbanStepOrderBy) KanbanStepSelectConfigOption {
	return func(s *KanbanStepSelectConfig) {
		s.orderBy = strings.Join(rows, ", ")
	}
}

type (
	KanbanStepJoinBy  = string
	KanbanStepOrderBy = string
)

// KanbanStep represents a row from 'public.kanban_steps'.
type KanbanStep struct {
	KanbanStepID  int    `json:"kanban_step_id"` // kanban_step_id
	TeamID        int    `json:"team_id"`        // team_id
	StepOrder     int16  `json:"step_order"`     // step_order
	Name          string `json:"name"`           // name
	Description   string `json:"description"`    // description
	TimeTrackable bool   `json:"time_trackable"` // time_trackable
	Disabled      bool   `json:"disabled"`       // disabled
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the KanbanStep exists in the database.
func (ks *KanbanStep) Exists() bool {
	return ks._exists
}

// Deleted returns true when the KanbanStep has been marked for deletion from
// the database.
func (ks *KanbanStep) Deleted() bool {
	return ks._deleted
}

// Insert inserts the KanbanStep to the database.
func (ks *KanbanStep) Insert(ctx context.Context, db DB) error {
	switch {
	case ks._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case ks._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	sqlstr := `INSERT INTO public.kanban_steps (` +
		`kanban_step_id, team_id, step_order, name, description, time_trackable, disabled` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) `
	// run
	logf(sqlstr, ks.KanbanStepID, ks.TeamID, ks.StepOrder, ks.Name, ks.Description, ks.TimeTrackable, ks.Disabled)
	if _, err := db.Exec(ctx, sqlstr, ks.KanbanStepID, ks.TeamID, ks.StepOrder, ks.Name, ks.Description, ks.TimeTrackable, ks.Disabled); err != nil {
		return logerror(err)
	}
	// set exists
	ks._exists = true
	return nil
}

// Update updates a KanbanStep in the database.
func (ks *KanbanStep) Update(ctx context.Context, db DB) error {
	switch {
	case !ks._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case ks._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.kanban_steps SET ` +
		`team_id = $1, step_order = $2, name = $3, description = $4, time_trackable = $5, disabled = $6 ` +
		`WHERE kanban_step_id = $7 `
	// run
	logf(sqlstr, ks.TeamID, ks.StepOrder, ks.Name, ks.Description, ks.TimeTrackable, ks.Disabled, ks.KanbanStepID)
	if _, err := db.Exec(ctx, sqlstr, ks.TeamID, ks.StepOrder, ks.Name, ks.Description, ks.TimeTrackable, ks.Disabled, ks.KanbanStepID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the KanbanStep to the database.
func (ks *KanbanStep) Save(ctx context.Context, db DB) error {
	if ks.Exists() {
		return ks.Update(ctx, db)
	}
	return ks.Insert(ctx, db)
}

// Upsert performs an upsert for KanbanStep.
func (ks *KanbanStep) Upsert(ctx context.Context, db DB) error {
	switch {
	case ks._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.kanban_steps (` +
		`kanban_step_id, team_id, step_order, name, description, time_trackable, disabled` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`)` +
		` ON CONFLICT (kanban_step_id) DO ` +
		`UPDATE SET ` +
		`team_id = EXCLUDED.team_id, step_order = EXCLUDED.step_order, name = EXCLUDED.name, description = EXCLUDED.description, time_trackable = EXCLUDED.time_trackable, disabled = EXCLUDED.disabled  `
	// run
	logf(sqlstr, ks.KanbanStepID, ks.TeamID, ks.StepOrder, ks.Name, ks.Description, ks.TimeTrackable, ks.Disabled)
	if _, err := db.Exec(ctx, sqlstr, ks.KanbanStepID, ks.TeamID, ks.StepOrder, ks.Name, ks.Description, ks.TimeTrackable, ks.Disabled); err != nil {
		return logerror(err)
	}
	// set exists
	ks._exists = true
	return nil
}

// Delete deletes the KanbanStep from the database.
func (ks *KanbanStep) Delete(ctx context.Context, db DB) error {
	switch {
	case !ks._exists: // doesn't exist
		return nil
	case ks._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.kanban_steps ` +
		`WHERE kanban_step_id = $1 `
	// run
	logf(sqlstr, ks.KanbanStepID)
	if _, err := db.Exec(ctx, sqlstr, ks.KanbanStepID); err != nil {
		return logerror(err)
	}
	// set deleted
	ks._deleted = true
	return nil
}

// KanbanStepByKanbanStepID retrieves a row from 'public.kanban_steps' as a KanbanStep.
//
// Generated from index 'kanban_steps_pkey'.
func KanbanStepByKanbanStepID(ctx context.Context, db DB, kanbanStepID int, opts ...KanbanStepSelectConfigOption) (*KanbanStep, error) {
	c := &KanbanStepSelectConfig{}
	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`kanban_step_id, team_id, step_order, name, description, time_trackable, disabled ` +
		`FROM public.kanban_steps ` +
		`WHERE kanban_step_id = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, kanbanStepID)
	ks := KanbanStep{
		_exists: true,
	}
	if err := db.QueryRow(ctx, sqlstr, kanbanStepID).Scan(&ks.KanbanStepID, &ks.TeamID, &ks.StepOrder, &ks.Name, &ks.Description, &ks.TimeTrackable, &ks.Disabled); err != nil {
		return nil, logerror(err)
	}
	return &ks, nil
}

// Team returns the Team associated with the KanbanStep's (TeamID).
//
// Generated from foreign key 'kanban_steps_team_id_fkey'.
func (ks *KanbanStep) Team(ctx context.Context, db DB) (*Team, error) {
	return TeamByTeamID(ctx, db, ks.TeamID)
}
