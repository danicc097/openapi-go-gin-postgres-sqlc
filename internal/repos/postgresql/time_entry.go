package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// TimeEntry represents the repository used for interacting with TimeEntry records.
type TimeEntry struct {
	q *db.Queries
}

// NewTimeEntry instantiates the TimeEntry repository.
func NewTimeEntry() *TimeEntry {
	return &TimeEntry{
		q: db.New(),
	}
}

var _ repos.TimeEntry = (*TimeEntry)(nil)

func (wit *TimeEntry) Create(ctx context.Context, d db.DBTX, params db.TimeEntryCreateParams) (*db.TimeEntry, error) {
	timeEntry := &db.TimeEntry{
		ActivityID:      params.ActivityID,
		TeamID:          params.TeamID,
		WorkItemID:      params.WorkItemID,
		UserID:          params.UserID,
		Comment:         params.Comment,
		Start:           params.Start,
		DurationMinutes: params.DurationMinutes,
	}

	if _, err := timeEntry.Save(ctx, d); err != nil {
		return nil, err
	}

	return timeEntry, nil
}

func (wit *TimeEntry) Update(ctx context.Context, d db.DBTX, id int64, params db.TimeEntryUpdateParams) (*db.TimeEntry, error) {
	timeEntry, err := wit.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get timeEntry by id %w", parseErrorDetail(err))
	}

	if params.TeamID != nil {
		timeEntry.TeamID = *params.TeamID
	}
	if params.WorkItemID != nil {
		timeEntry.WorkItemID = *params.WorkItemID
	}
	if params.DurationMinutes != nil {
		timeEntry.DurationMinutes = *params.DurationMinutes
	}
	if params.Comment != nil {
		timeEntry.Comment = *params.Comment
	}
	if params.Start != nil {
		timeEntry.Start = *params.Start
	}
	if params.UserID != nil {
		timeEntry.UserID = *params.UserID
	}
	if params.ActivityID != nil {
		timeEntry.ActivityID = *params.ActivityID
	}
	timeEntry, err = timeEntry.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update timeEntry: %w", parseErrorDetail(err))
	}

	return timeEntry, err
}

func (wit *TimeEntry) ByID(ctx context.Context, d db.DBTX, id int64) (*db.TimeEntry, error) {
	timeEntry, err := db.TimeEntryByTimeEntryID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get timeEntry: %w", parseErrorDetail(err))
	}

	return timeEntry, nil
}

func (wit *TimeEntry) Delete(ctx context.Context, d db.DBTX, id int64) (*db.TimeEntry, error) {
	timeEntry, err := wit.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get timeEntry by id %w", parseErrorDetail(err))
	}

	err = timeEntry.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete timeEntry: %w", parseErrorDetail(err))
	}

	return timeEntry, err
}
