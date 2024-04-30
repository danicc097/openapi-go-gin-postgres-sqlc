package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	models1 "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

type TimeEntry struct {
	logger   *zap.SugaredLogger
	repos    *repos.Repos
	authzsvc *Authorization
}

// NewTimeEntry returns a new TimeEntry service.
func NewTimeEntry(logger *zap.SugaredLogger, repos *repos.Repos) *TimeEntry {
	authzsvc := NewAuthorization(logger)

	return &TimeEntry{
		logger:   logger,
		repos:    repos,
		authzsvc: authzsvc,
	}
}

// ByID gets a time entry by ID.
func (te *TimeEntry) ByID(ctx context.Context, d models1.DBTX, id models1.TimeEntryID) (*models1.TimeEntry, error) {
	defer newOTelSpan().Build(ctx).End()

	teObj, err := te.repos.TimeEntry.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.TimeEntry.ByID: %w", err)
	}

	return teObj, nil
}

// Create creates a new time entry.
func (te *TimeEntry) Create(ctx context.Context, d models1.DBTX, caller CtxUser, params *models1.TimeEntryCreateParams) (*models1.TimeEntry, error) {
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

func (te *TimeEntry) validateCreateParams(d models1.DBTX, caller CtxUser, params *models1.TimeEntryCreateParams) error {
	if err := te.validateBaseParams(validateModeCreate, d, caller, params); err != nil {
		return err
	}

	// extra create param validation, if any

	return nil
}

func (te *TimeEntry) validateUpdateParams(d models1.DBTX, caller CtxUser, params *models1.TimeEntryUpdateParams) error {
	if err := te.validateBaseParams(validateModeUpdate, d, caller, params); err != nil {
		return err
	}

	// extra update param validation, if any

	return nil
}

/**
 * TODO:
 * xo : new property annotations: no-update, no-create, no-update-json, no-create-json to remove json fields or struct fields from either create or update or both,
 * instead of current `private` annotations.
 */
// example: extra fields required in services for update (same applies for create)
type TimeEntryUpdateParams struct {
	models1.TimeEntryUpdateParams
	// no need for getters for new field, validate in dedicated validateUpdateParams only.
	// if it conflicts with base validation, just skip based on validatemode
	NewField string `json:"newField"`
}

// if we need service update/create params later, embed db params in services.*{Create|Update}Params
// and add new fields+accessors as required.
func (te *TimeEntry) validateBaseParams(mode validateMode, d models1.DBTX, caller CtxUser, params models1.TimeEntryParams) error {
	if params.GetTeamID() != nil && params.GetWorkItemID() != nil {
		// checked in db, but can verify here too
		return internal.NewErrorf(models.ErrorCodeInvalidArgument, "cannot link activity to both team and work item")
	}

	if params.GetUserID() != nil {
		if caller.UserID != *params.GetUserID() {
			return internal.NewErrorf(models.ErrorCodeUnauthorized, "cannot add activity for a different user")
		}
	}

	if params.GetTeamID() != nil {
		teamIDs := make([]models1.TeamID, len(caller.Teams))
		for i, t := range caller.Teams {
			teamIDs[i] = t.TeamID
		}
		if !slices.Contains(teamIDs, *params.GetTeamID()) {
			return internal.NewErrorf(models.ErrorCodeUnauthorized, "cannot link activity to an unassigned team")
		}
	}

	if params.GetWorkItemID() != nil {
		wi, err := te.repos.WorkItem.ByID(context.Background(), d, *params.GetWorkItemID(), models1.WithWorkItemJoin(models1.WorkItemJoins{Assignees: true}))
		if err != nil {
			return fmt.Errorf("repos.WorkItem.ByID: %w", err)
		}

		memberIDs := make(map[models1.UserID]bool)
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
func (te *TimeEntry) Update(ctx context.Context, d models1.DBTX, caller CtxUser, id models1.TimeEntryID, params *models1.TimeEntryUpdateParams) (*models1.TimeEntry, error) {
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
func (te *TimeEntry) Delete(ctx context.Context, d models1.DBTX, id models1.TimeEntryID) (*models1.TimeEntry, error) {
	defer newOTelSpan().Build(ctx).End()

	teObj, err := te.repos.TimeEntry.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.TimeEntry.Delete: %w", err)
	}

	return teObj, nil
}
