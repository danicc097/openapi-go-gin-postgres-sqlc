package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

// TimeEntry represents the repository used for interacting with TimeEntry records.
type TimeEntry struct {
	q models.Querier
}

// NewTimeEntry instantiates the TimeEntry repository.
func NewTimeEntry() *TimeEntry {
	return &TimeEntry{
		q: NewQuerierWrapper(models.New()),
	}
}

var _ repos.TimeEntry = (*TimeEntry)(nil)

func (wit *TimeEntry) Create(ctx context.Context, d models.DBTX, params *models.TimeEntryCreateParams) (*models.TimeEntry, error) {
	timeEntry, err := models.CreateTimeEntry(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not create time entry: %w", ParseDBErrorDetail(err))
	}

	return timeEntry, nil
}

func (wit *TimeEntry) Update(ctx context.Context, d models.DBTX, id models.TimeEntryID, params *models.TimeEntryUpdateParams) (*models.TimeEntry, error) {
	timeEntry, err := wit.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get timeEntry by id %w", ParseDBErrorDetail(err))
	}

	timeEntry.SetUpdateParams(params)

	timeEntry, err = timeEntry.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update timeEntry: %w", ParseDBErrorDetail(err))
	}

	return timeEntry, err
}

func (wit *TimeEntry) ByID(ctx context.Context, d models.DBTX, id models.TimeEntryID, opts ...models.TimeEntrySelectConfigOption) (*models.TimeEntry, error) {
	timeEntry, err := models.TimeEntryByTimeEntryID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get timeEntry: %w", ParseDBErrorDetail(err))
	}

	return timeEntry, nil
}

func (wit *TimeEntry) Delete(ctx context.Context, d models.DBTX, id models.TimeEntryID) (*models.TimeEntry, error) {
	timeEntry := &models.TimeEntry{
		TimeEntryID: id,
	}

	err := timeEntry.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete timeEntry: %w", ParseDBErrorDetail(err))
	}

	return timeEntry, err
}
