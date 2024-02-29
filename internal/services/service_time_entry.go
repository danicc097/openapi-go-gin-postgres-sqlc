package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

type TimeEntry struct {
	logger *zap.SugaredLogger
	repos  *repos.Repos
}

// NewTimeEntry returns a new TimeEntry service.
func NewTimeEntry(logger *zap.SugaredLogger, repos *repos.Repos) *TimeEntry {
	return &TimeEntry{
		logger: logger,
		repos:  repos,
	}
}

// ByID gets a time entry by ID.
func (te *TimeEntry) ByID(ctx context.Context, d db.DBTX, id db.TimeEntryID) (*db.TimeEntry, error) {
	defer newOTelSpan().Build(ctx).End()

	teObj, err := te.repos.TimeEntry.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.TimeEntry.ByID: %w", err)
	}

	return teObj, nil
}

// Create creates a new time entry.
func (te *TimeEntry) Create(ctx context.Context, d db.DBTX, caller CtxUser, params *db.TimeEntryCreateParams) (*db.TimeEntry, error) {
	defer newOTelSpan().Build(ctx).End()

	if err := te.validateUserAuthorized(caller, params.UserID); err != nil {
		return nil, err
	}
	if err := te.validateUserIsInTeam(caller, params.TeamID); err != nil {
		return nil, err
	}

	if params.WorkItemID != nil {
		wi, err := te.repos.WorkItem.ByID(ctx, d, *params.WorkItemID, db.WithWorkItemJoin(db.WorkItemJoins{Assignees: true}))
		if err != nil {
			return nil, fmt.Errorf("repos.WorkItem.ByID: %w", err)
		}

		memberIDs := make(map[db.UserID]bool)
		for _, m := range *wi.AssigneesJoin {
			memberIDs[m.User.UserID] = true
		}
		if _, ok := memberIDs[caller.UserID]; !ok {
			// FIXME filter where not null for m2m in assigned members not doing what we think
			return nil, internal.NewErrorf(models.ErrorCodeUnauthorized, "cannot link activity to an unassigned work item")
		}
	}

	teObj, err := te.repos.TimeEntry.Create(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("repos.TimeEntry.Create: %w", err)
	}

	te.logger.Infof("created time entry by user %q", teObj.UserID)

	return teObj, nil
}

func (te *TimeEntry) validateUserAuthorized(caller CtxUser, userID db.UserID) error {
	if caller.UserID != userID {
		return internal.NewErrorf(models.ErrorCodeUnauthorized, "cannot add activity for a different user")
	}

	return nil
}

func (te *TimeEntry) validateUserIsInTeam(caller CtxUser, teamID *db.TeamID) error {
	if teamID != nil {
		teamIDs := make([]db.TeamID, len(caller.Teams))
		for i, t := range caller.Teams {
			teamIDs[i] = t.TeamID
		}
		if !slices.Contains(teamIDs, *teamID) {
			return internal.NewErrorf(models.ErrorCodeUnauthorized, "cannot link activity to an unassigned team")
		}
	}

	return nil
}

// Update updates an existing time entry.
func (te *TimeEntry) Update(ctx context.Context, d db.DBTX, caller CtxUser, id db.TimeEntryID, params *db.TimeEntryUpdateParams) (*db.TimeEntry, error) {
	defer newOTelSpan().Build(ctx).End()

	// TODO: xo should create interface for params, with GetUserID shared. therefore we could have a generic validate function that works for both create and update params.
	// if for some reason we want to skip a validation on either update or create we could pass a flag easily.
	if params.UserID != nil {
		if err := te.validateUserAuthorized(caller, *params.UserID); err != nil {
			return nil, err
		}
	}

	if params.TeamID != nil {
		if err := te.validateUserIsInTeam(caller, *params.TeamID); err != nil {
			return nil, err
		}
	}

	teObj, err := te.repos.TimeEntry.Update(ctx, d, id, params)
	if err != nil {
		return nil, fmt.Errorf("repos.TimeEntry.Update: %w", err)
	}

	return teObj, nil
}

// Delete deletes a time entry by ID.
func (te *TimeEntry) Delete(ctx context.Context, d db.DBTX, id db.TimeEntryID) (*db.TimeEntry, error) {
	defer newOTelSpan().Build(ctx).End()

	teObj, err := te.repos.TimeEntry.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.TimeEntry.Delete: %w", err)
	}

	return teObj, nil
}
