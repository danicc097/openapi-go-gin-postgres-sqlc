package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

type TimeEntry struct {
	logger *zap.SugaredLogger
	teRepo repos.TimeEntry
}

// NewTimeEntry returns a new TimeEntry service.
func NewTimeEntry(logger *zap.SugaredLogger, teRepo repos.TimeEntry) *TimeEntry {
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
func (a *TimeEntry) Create(ctx context.Context, d db.DBTX, caller *db.User, params *db.TimeEntryCreateParams) (*db.TimeEntry, error) {
	defer newOTELSpan(ctx, "TimeEntry.Create").End()

	if caller.UserID != params.UserID {
		return nil, internal.NewErrorf(internal.ErrorCodeUnauthorized, "cannot add activity for a different user")
	}

	if params.TeamID != nil {
		teamIDs := make([]int, len(*caller.TeamsJoin))
		for i, t := range *caller.TeamsJoin {
			teamIDs[i] = t.TeamID
		}
		if !slices.Contains(teamIDs, *params.TeamID) {
			return nil, internal.NewErrorf(internal.ErrorCodeUnauthorized, "cannot link activity to an unassigned team")
		}
	}

	// TODO if activity linked to workitem, dont allow if caller is not a member

	teObj, err := a.teRepo.Create(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("teRepo.Create: %w", err)
	}

	a.logger.Infof("created time entry by user %q", teObj.UserID)

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
