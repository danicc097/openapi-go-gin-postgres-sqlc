package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
)

// WorkItemTagPublic represents fields that may be exposed from 'public.work_item_tags'
// and embedded in other response models.
// Include "property:private" in a SQL column comment to exclude a field.
// Joins may be explicitly added in the Response struct.
type WorkItemTagPublic struct {
	WorkItemTagID int    `json:"workItemTagID"` // work_item_tag_id
	Name          string `json:"name"`          // name
	Description   string `json:"description"`   // description
	Color         string `json:"color"`         // color
}

// WorkItemTag represents a row from 'public.work_item_tags'.
type WorkItemTag struct {
	WorkItemTagID int    `json:"work_item_tag_id" db:"work_item_tag_id"` // work_item_tag_id
	Name          string `json:"name" db:"name"`                         // name
	Description   string `json:"description" db:"description"`           // description
	Color         string `json:"color" db:"color"`                       // color

	// xo fields
	_exists, _deleted bool
}

func (x *WorkItemTag) ToPublic() WorkItemTagPublic {
	return WorkItemTagPublic{
		WorkItemTagID: x.WorkItemTagID, Name: x.Name, Description: x.Description, Color: x.Color,
	}
}

type WorkItemTagSelectConfig struct {
	limit   string
	orderBy string
	joins   WorkItemTagJoins
}
type WorkItemTagSelectConfigOption func(*WorkItemTagSelectConfig)

// WithWorkItemTagLimit limits row selection.
func WithWorkItemTagLimit(limit int) WorkItemTagSelectConfigOption {
	return func(s *WorkItemTagSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type WorkItemTagOrderBy = string

type WorkItemTagJoins struct{}

// WithWorkItemTagJoin orders results by the given columns.
func WithWorkItemTagJoin(joins WorkItemTagJoins) WorkItemTagSelectConfigOption {
	return func(s *WorkItemTagSelectConfig) {
		s.joins = joins
	}
}

// Exists returns true when the WorkItemTag exists in the database.
func (wit *WorkItemTag) Exists() bool {
	return wit._exists
}

// Deleted returns true when the WorkItemTag has been marked for deletion from
// the database.
func (wit *WorkItemTag) Deleted() bool {
	return wit._deleted
}

// Insert inserts the WorkItemTag to the database.
func (wit *WorkItemTag) Insert(ctx context.Context, db DB) error {
	switch {
	case wit._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case wit._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.work_item_tags (` +
		`name, description, color` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING work_item_tag_id `
	// run
	logf(sqlstr, wit.Name, wit.Description, wit.Color)
	if err := db.QueryRow(ctx, sqlstr, wit.Name, wit.Description, wit.Color).Scan(&wit.WorkItemTagID); err != nil {
		return logerror(err)
	}
	// set exists
	wit._exists = true
	return nil
}

// Update updates a WorkItemTag in the database.
func (wit *WorkItemTag) Update(ctx context.Context, db DB) error {
	switch {
	case !wit._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case wit._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.work_item_tags SET ` +
		`name = $1, description = $2, color = $3 ` +
		`WHERE work_item_tag_id = $4 ` +
		`RETURNING work_item_tag_id `
	// run
	logf(sqlstr, wit.Name, wit.Description, wit.Color, wit.WorkItemTagID)
	if err := db.QueryRow(ctx, sqlstr, wit.Name, wit.Description, wit.Color, wit.WorkItemTagID).Scan(&wit.WorkItemTagID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the WorkItemTag to the database.
func (wit *WorkItemTag) Save(ctx context.Context, db DB) error {
	if wit.Exists() {
		return wit.Update(ctx, db)
	}
	return wit.Insert(ctx, db)
}

// Upsert performs an upsert for WorkItemTag.
func (wit *WorkItemTag) Upsert(ctx context.Context, db DB) error {
	switch {
	case wit._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.work_item_tags (` +
		`work_item_tag_id, name, description, color` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)` +
		` ON CONFLICT (work_item_tag_id) DO ` +
		`UPDATE SET ` +
		`name = EXCLUDED.name, description = EXCLUDED.description, color = EXCLUDED.color  `
	// run
	logf(sqlstr, wit.WorkItemTagID, wit.Name, wit.Description, wit.Color)
	if _, err := db.Exec(ctx, sqlstr, wit.WorkItemTagID, wit.Name, wit.Description, wit.Color); err != nil {
		return logerror(err)
	}
	// set exists
	wit._exists = true
	return nil
}

// Delete deletes the WorkItemTag from the database.
func (wit *WorkItemTag) Delete(ctx context.Context, db DB) error {
	switch {
	case !wit._exists: // doesn't exist
		return nil
	case wit._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.work_item_tags ` +
		`WHERE work_item_tag_id = $1 `
	// run
	logf(sqlstr, wit.WorkItemTagID)
	if _, err := db.Exec(ctx, sqlstr, wit.WorkItemTagID); err != nil {
		return logerror(err)
	}
	// set deleted
	wit._deleted = true
	return nil
}

// WorkItemTagByName retrieves a row from 'public.work_item_tags' as a WorkItemTag.
//
// Generated from index 'work_item_tags_name_key'.
func WorkItemTagByName(ctx context.Context, db DB, name string, opts ...WorkItemTagSelectConfigOption) (*WorkItemTag, error) {
	c := &WorkItemTagSelectConfig{joins: WorkItemTagJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_tags.work_item_tag_id,
work_item_tags.name,
work_item_tags.description,
work_item_tags.color ` +
		`FROM public.work_item_tags ` +
		`` +
		` WHERE work_item_tags.name = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, name)
	wit := WorkItemTag{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, name).Scan(&wit.WorkItemTagID, &wit.Name, &wit.Description, &wit.Color); err != nil {
		return nil, logerror(err)
	}
	return &wit, nil
}

// WorkItemTagByWorkItemTagID retrieves a row from 'public.work_item_tags' as a WorkItemTag.
//
// Generated from index 'work_item_tags_pkey'.
func WorkItemTagByWorkItemTagID(ctx context.Context, db DB, workItemTagID int, opts ...WorkItemTagSelectConfigOption) (*WorkItemTag, error) {
	c := &WorkItemTagSelectConfig{joins: WorkItemTagJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`work_item_tags.work_item_tag_id,
work_item_tags.name,
work_item_tags.description,
work_item_tags.color ` +
		`FROM public.work_item_tags ` +
		`` +
		` WHERE work_item_tags.work_item_tag_id = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, workItemTagID)
	wit := WorkItemTag{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, workItemTagID).Scan(&wit.WorkItemTagID, &wit.Name, &wit.Description, &wit.Color); err != nil {
		return nil, logerror(err)
	}
	return &wit, nil
}
