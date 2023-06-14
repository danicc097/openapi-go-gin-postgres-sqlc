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

func (wit *TimeEntry) Create(ctx context.Context, d db.DBTX, params *db.TimeEntryCreateParams) (*db.TimeEntry, error) {
	timeEntry, err := db.CreateTimeEntry(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not create time entry: %w", parseErrorDetail(err))
	}

	return timeEntry, nil
}

func (wit *TimeEntry) Update(ctx context.Context, d db.DBTX, id int64, params *db.TimeEntryUpdateParams) (*db.TimeEntry, error) {
	timeEntry, err := wit.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get timeEntry by id %w", parseErrorDetail(err))
	}

	timeEntry, err = timeEntry.Update(ctx, d, params)
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
