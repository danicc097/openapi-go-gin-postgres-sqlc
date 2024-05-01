package services

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"go.uber.org/zap"
)

type WorkItemCreateParams struct {
	TagIDs  []models.WorkItemTagID `json:"tagIDs"  nullable:"false" required:"true"`
	Members []Member               `json:"members" nullable:"false" required:"true"`
}

type WorkItem struct {
	logger *zap.SugaredLogger
	repos  *repos.Repos
	// sharedDBOpts represents shared db select options for all work item entities
	// for returned values
	getSharedDBOpts func() []models.WorkItemSelectConfigOption
}

// NewWorkItem returns a new WorkItem service with common logic for all project worki tems.
func NewWorkItem(logger *zap.SugaredLogger, repos *repos.Repos) *WorkItem {
	return &WorkItem{
		logger: logger,
		repos:  repos,
		getSharedDBOpts: func() []models.WorkItemSelectConfigOption {
			// keep in sync with SharedWorkItemJoins
			return []models.WorkItemSelectConfigOption{models.WithWorkItemJoin(models.WorkItemJoins{
				Assignees:        true,
				WorkItemTags:     true,
				TimeEntries:      true,
				WorkItemComments: true,
				WorkItemType:     true,
			})}
		},
	}
}

func (w *WorkItem) AssignUsers(ctx context.Context, d models.DBTX, workItemID models.WorkItemID, members []Member) error {
	wi, err := w.repos.WorkItem.ByID(ctx, d, workItemID)
	if err != nil {
		return fmt.Errorf("repos.WorkItem.ByID: %w", err)
	}

	for idx, member := range members {
		user, err := w.repos.User.ByID(ctx, d, member.UserID, models.WithUserJoin(models.UserJoins{MemberTeams: true}))
		if err != nil {
			return internal.WrapErrorWithLocf(err, models.ErrorCodeNotFound, []string{strconv.Itoa(idx)}, "user with id %s not found", member.UserID)
		}

		var userInTeam bool
		for _, team := range *user.MemberTeamsJoin {
			if team.TeamID == wi.TeamID {
				userInTeam = true
			}
		}
		if !userInTeam {
			return internal.WrapErrorWithLocf(nil, models.ErrorCodeUnauthorized, []string{strconv.Itoa(idx)}, "user %q does not belong to team %q", user.Email, wi.TeamID)
		}

		err = w.repos.WorkItem.AssignUser(ctx, d, &models.WorkItemAssigneeCreateParams{
			Assignee:   member.UserID,
			WorkItemID: wi.WorkItemID,
			Role:       member.Role,
		})
		var ierr *internal.Error
		if err != nil {
			if errors.As(err, &ierr) && ierr.Code() == models.ErrorCodeAlreadyExists {
				continue
			}

			return internal.WrapErrorWithLocf(err, "", []string{strconv.Itoa(idx)}, "could not assign member %s", member.UserID)
		}
	}

	return nil
}

func (w *WorkItem) RemoveAssignedUsers(ctx context.Context, d models.DBTX, workItemID models.WorkItemID, members []models.UserID) error {
	for idx, member := range members {
		// nolint: exhaustruct
		lookup := &models.WorkItemAssignee{
			Assignee:   member,
			WorkItemID: workItemID,
		}

		err := lookup.Delete(ctx, d)
		if err != nil {
			return internal.WrapErrorWithLocf(err, "", []string{strconv.Itoa(idx)}, "could not remove member %s", member)
		}
	}

	return nil
}

func (w *WorkItem) AssignTags(ctx context.Context, d models.DBTX, workItemID models.WorkItemID, tagIDs []models.WorkItemTagID) error {
	// IMPORTANT: using IDs for services allows each method to grab necessary joins, etc. as needed instead of relying on a passed db entity
	// to hold them.
	wi, err := w.repos.WorkItem.ByID(ctx, d, workItemID, models.WithWorkItemJoin(models.WorkItemJoins{Team: true}))
	if err != nil {
		return fmt.Errorf("repos.WorkItem.ByID: %w", err)
	}

	for idx, tagID := range tagIDs {
		tag, err := w.repos.WorkItemTag.ByID(ctx, d, tagID)
		if err != nil {
			return internal.WrapErrorWithLocf(err, models.ErrorCodeNotFound, []string{strconv.Itoa(idx)}, "tag with id %d not found", tagID)
		}

		if wi.TeamJoin.ProjectID != tag.ProjectID {
			return internal.WrapErrorWithLocf(nil, models.ErrorCodeUnauthorized, []string{strconv.Itoa(idx)}, "tag %q does not belong to work item's project", tag.Name)
		}

		err = w.repos.WorkItem.AssignTag(ctx, d, &models.WorkItemWorkItemTagCreateParams{
			WorkItemTagID: tagID,
			WorkItemID:    wi.WorkItemID,
		})
		var ierr *internal.Error
		if err != nil {
			if errors.As(err, &ierr) && ierr.Code() == models.ErrorCodeAlreadyExists {
				continue
			}

			return internal.WrapErrorWithLocf(err, "", []string{strconv.Itoa(idx)}, "could not assign tag %s", tag.Name)
		}
	}

	return nil
}

func (w *WorkItem) RemoveTags(ctx context.Context, d models.DBTX, workItemID models.WorkItemID, tagIDs []models.WorkItemTagID) error {
	for idx, tagID := range tagIDs {
		// nolint: exhaustruct
		lookup := &models.WorkItemWorkItemTag{
			WorkItemTagID: tagID,
			WorkItemID:    workItemID,
		}

		err := lookup.Delete(ctx, d)
		if err != nil {
			return internal.WrapErrorWithLocf(err, "", []string{strconv.Itoa(idx)}, "could not remove tag %d", tagID)
		}
	}

	return nil
}

// postCreate applies changes to a workitem common to all projects after entity creation.
// NOTE: returned error should not be wrapped in calling function.
func (w *WorkItem) postCreate(ctx context.Context, d models.DBTX, workItemID models.WorkItemID, params WorkItemCreateParams) error {
	if err := w.AssignTags(ctx, d, workItemID, params.TagIDs); err != nil {
		return internal.WrapErrorWithLocf(err, "", []string{"tagIDs"}, "could not assign tags")
	}

	if err := w.AssignUsers(ctx, d, workItemID, params.Members); err != nil {
		return internal.WrapErrorWithLocf(err, "", []string{"members"}, "could not assign members")
	}

	return nil
}

// Restore restores a work item marked as deleted by ID.
func (w *WorkItem) Restore(ctx context.Context, d models.DBTX, id models.WorkItemID) (*models.WorkItem, error) {
	defer newOTelSpan().Build(ctx).End()

	wi, err := w.repos.WorkItem.Restore(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.WorkItem.Restore: %w", err)
	}

	return wi, nil
}

// Delete deletes a work item by ID.
func (w *WorkItem) Delete(ctx context.Context, d models.DBTX, id models.WorkItemID) (*models.WorkItem, error) {
	defer newOTelSpan().Build(ctx).End()

	wi, err := w.repos.WorkItem.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.WorkItem.Delete: %w", err)
	}

	return wi, nil
}

func (w *WorkItem) validateCreateParams(d models.DBTX, caller CtxUser, params *models.WorkItemCreateParams) error {
	if err := w.validateBaseParams(validateModeCreate, d, caller, params); err != nil {
		return err
	}

	return nil
}

func (w *WorkItem) validateUpdateParams(d models.DBTX, caller CtxUser, params *models.WorkItemUpdateParams) error {
	if err := w.validateBaseParams(validateModeUpdate, d, caller, params); err != nil {
		return err
	}

	return nil
}

func (w *WorkItem) validateBaseParams(mode validateMode, d models.DBTX, caller CtxUser, params models.WorkItemParams) error {
	return nil
}
