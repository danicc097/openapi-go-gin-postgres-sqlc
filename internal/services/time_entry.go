package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
)

type TimeEntry struct {
	logger *zap.Logger
	teRepo repos.TimeEntry
}

// NewTimeEntry returns a new TimeEntry service.
func NewTimeEntry(logger *zap.Logger, teRepo repos.TimeEntry) *TimeEntry {
	return &TimeEntry{
		logger: logger,
		teRepo: teRepo,
	}
}

// ByID gets a time entry by ID.
func (a *TimeEntry) ByID(ctx context.Context, d db.DBTX, id int64) (*db.TimeEntry, error) {
	defer newOTELSpan(ctx, "TimeEntry.ByID").End()

	teObj, err := a.teRepo.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("teRepo.ByID: %w", err)
	}

	return teObj, nil
}

// Create creates a new time entry.
func (a *TimeEntry) Create(ctx context.Context, d db.DBTX, params *db.TimeEntryCreateParams) (*db.TimeEntry, error) {
	defer newOTELSpan(ctx, "TimeEntry.Create").End()

	teObj, err := a.teRepo.Create(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("teRepo.Create: %w", err)
	}

	return teObj, nil
}

// Update updates an existing time entry.
func (a *TimeEntry) Update(ctx context.Context, d db.DBTX, id int64, params *db.TimeEntryUpdateParams) (*db.TimeEntry, error) {
	defer newOTELSpan(ctx, "TimeEntry.Update").End()

	teObj, err := a.teRepo.Update(ctx, d, id, params)
	if err != nil {
		return nil, fmt.Errorf("teRepo.Update: %w", err)
	}

	return teObj, nil
}

// Delete deletes a time entry by ID.
func (a *TimeEntry) Delete(ctx context.Context, d db.DBTX, id int64) (*db.TimeEntry, error) {
	defer newOTELSpan(ctx, "TimeEntry.Delete").End()

	teObj, err := a.teRepo.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("teRepo.Delete: %w", err)
	}

	return teObj, nil
}
