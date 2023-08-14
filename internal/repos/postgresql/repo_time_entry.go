package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// TimeEntry represents the repository used for interacting with TimeEntry records.
type TimeEntry struct {
	q db.Querier
}

// NewTimeEntry instantiates the TimeEntry repository.
func NewTimeEntry() *TimeEntry {
	return &TimeEntry{
		q: NewQuerierWrapper(db.New()),
	}
}

var _ repos.TimeEntry = (*TimeEntry)(nil)

func (wit *TimeEntry) Create(ctx context.Context, d db.DBTX, params *db.TimeEntryCreateParams) (*db.TimeEntry, error) {
	timeEntry, err := db.CreateTimeEntry(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not create time entry: %w", parseDBErrorDetail(err))
	}

	return timeEntry, nil
}

func (wit *TimeEntry) Update(ctx context.Context, d db.DBTX, id int, params *db.TimeEntryUpdateParams) (*db.TimeEntry, error) {
	timeEntry, err := wit.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get timeEntry by id %w", parseDBErrorDetail(err))
	}

	timeEntry.SetUpdateParams(params)

	timeEntry, err = timeEntry.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update timeEntry: %w", parseDBErrorDetail(err))
	}

	return timeEntry, err
}

func (wit *TimeEntry) ByID(ctx context.Context, d db.DBTX, id int, opts ...db.TimeEntrySelectConfigOption) (*db.TimeEntry, error) {
	timeEntry, err := db.TimeEntryByTimeEntryID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get timeEntry: %w", parseDBErrorDetail(err))
	}

	return timeEntry, nil
}

func (wit *TimeEntry) Delete(ctx context.Context, d db.DBTX, id int) (*db.TimeEntry, error) {
	timeEntry := &db.TimeEntry{
		TimeEntryID: id,
	}

	err := timeEntry.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete timeEntry: %w", parseDBErrorDetail(err))
	}

	return timeEntry, err
}
