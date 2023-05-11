package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

type TimeEntry struct {
	logger *zap.SugaredLogger
	teRepo repos.TimeEntry
	wiRepo repos.WorkItem
}

// NewTimeEntry returns a new TimeEntry service.
func NewTimeEntry(logger *zap.SugaredLogger, teRepo repos.TimeEntry, wiRepo repos.WorkItem) *TimeEntry {
	return &TimeEntry{
		logger: logger,
		teRepo: teRepo,
		wiRepo: wiRepo,
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
		teamIDs := make([]int, len(*caller.TeamsJoinUser))
		for i, t := range *caller.TeamsJoinUser {
			teamIDs[i] = t.TeamID
		}
		if !slices.Contains(teamIDs, *params.TeamID) {
			return nil, internal.NewErrorf(internal.ErrorCodeUnauthorized, "cannot link activity to an unassigned team")
		}
	}

	if params.WorkItemID != nil {
		wi, err := a.wiRepo.ByID(ctx, d, *params.WorkItemID, db.WithWorkItemJoin(db.WorkItemJoins{Members: true}))
		if err != nil {
			return nil, fmt.Errorf("wiRepo.ByID: %w", err)
		}

		// FIXME xo joins scanning when join table has extra fields
		fmt.Printf("wi.MembersJoin: %v\n", wi.MembersJoin)
		fmt.Printf("wi.TimeEntriesJoin: %v\n", wi.TimeEntriesJoin)

		memberIDs := make([]uuid.UUID, len(*wi.MembersJoin))
		for i, m := range *wi.MembersJoin {
			memberIDs[i] = m.User.UserID
		}
		if !slices.Contains(memberIDs, caller.UserID) {
			return nil, internal.NewErrorf(internal.ErrorCodeUnauthorized, "cannot link activity to an unassigned work item")
		}
	}

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
