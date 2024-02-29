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

	if err := te.validateCreateParams(d, caller, params); err != nil {
		return nil, err
	}

	teObj, err := te.repos.TimeEntry.Create(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("repos.TimeEntry.Create: %w", err)
	}

	te.logger.Infof("created time entry by user %q", teObj.UserID)

	return teObj, nil
}

func (te *TimeEntry) validateCreateParams(d db.DBTX, caller CtxUser, params *db.TimeEntryCreateParams) error {
	if err := te.validateBaseParams(d, caller, params); err != nil {
		return err
	}

	// extra create param validation, if any

	return nil
}

func (te *TimeEntry) validateUpdateParams(d db.DBTX, caller CtxUser, params *db.TimeEntryUpdateParams) error {
	if err := te.validateBaseParams(d, caller, params); err != nil {
		return err
	}

	// extra update param validation, if any

	return nil
}

// if we need service update/create params later, embed db params in services.*{Create|Update}Params
// and add new fields+accessors as required.
func (te *TimeEntry) validateBaseParams(d db.DBTX, caller CtxUser, params db.TimeEntryParams) error {
	if params.GetUserID() != nil {
		if caller.UserID != *params.GetUserID() {
			return internal.NewErrorf(models.ErrorCodeUnauthorized, "cannot add activity for a different user")
		}
	}

	if params.GetTeamID() != nil {
		teamIDs := make([]db.TeamID, len(caller.Teams))
		for i, t := range caller.Teams {
			teamIDs[i] = t.TeamID
		}
		if !slices.Contains(teamIDs, *params.GetTeamID()) {
			return internal.NewErrorf(models.ErrorCodeUnauthorized, "cannot link activity to an unassigned team")
		}
	}

	if params.GetWorkItemID() != nil {
		wi, err := te.repos.WorkItem.ByID(context.Background(), d, *params.GetWorkItemID(), db.WithWorkItemJoin(db.WorkItemJoins{Assignees: true}))
		if err != nil {
			return fmt.Errorf("repos.WorkItem.ByID: %w", err)
		}

		memberIDs := make(map[db.UserID]bool)
		for _, m := range *wi.AssigneesJoin {
			memberIDs[m.User.UserID] = true
		}
		if _, ok := memberIDs[caller.UserID]; !ok {
			return internal.NewErrorf(models.ErrorCodeUnauthorized, "cannot link activity to an unassigned work item")
		}
	}

	return nil
}

// Update updates an existing time entry.
func (te *TimeEntry) Update(ctx context.Context, d db.DBTX, caller CtxUser, id db.TimeEntryID, params *db.TimeEntryUpdateParams) (*db.TimeEntry, error) {
	defer newOTelSpan().Build(ctx).End()

	if err := te.validateUpdateParams(d, caller, params); err != nil {
		return nil, err
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
