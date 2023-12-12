package services

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
)

type WorkItem struct {
	logger *zap.SugaredLogger
	repos  *repos.Repos
	// sharedDBOpts represents shared db select options for all work item entities
	// for returned values
	sharedDBOpts []db.WorkItemSelectConfigOption
}

// NewWorkItem returns a new WorkItem service with common logic for all project worki tems.
func NewWorkItem(logger *zap.SugaredLogger, repos *repos.Repos) *WorkItem {
	return &WorkItem{
		logger:       logger,
		repos:        repos,
		sharedDBOpts: []db.WorkItemSelectConfigOption{db.WithWorkItemJoin(db.WorkItemJoins{AssignedUsers: true, WorkItemTags: true})},
	}
}

func (w *WorkItem) AssignUsers(ctx context.Context, d db.DBTX, workItem *db.WorkItem, members []Member) error {
	for idx, member := range members {
		user, err := w.repos.User.ByID(ctx, d, member.UserID, db.WithUserJoin(db.UserJoins{TeamsMember: true}))
		if err != nil {
			return internal.WrapErrorWithLocf(err, models.ErrorCodeNotFound, []string{strconv.Itoa(idx)}, "user with id %s not found", member.UserID)
		}

		var userInTeam bool
		for _, team := range *user.MemberTeamsJoin {
			if team.TeamID == workItem.TeamID {
				userInTeam = true
			}
		}
		if !userInTeam {
			return internal.WrapErrorWithLocf(nil, models.ErrorCodeUnauthorized, []string{strconv.Itoa(idx)}, "user %q does not belong to team %q", user.Email, workItem.TeamID)
		}

		err = w.repos.WorkItem.AssignUser(ctx, d, &db.WorkItemAssignedUserCreateParams{
			AssignedUser: member.UserID,
			WorkItemID:   workItem.WorkItemID,
			Role:         member.Role,
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

func (w *WorkItem) RemoveAssignedUsers(ctx context.Context, d db.DBTX, workItem *db.WorkItem, members []db.UserID) error {
	for idx, member := range members {
		lookup := &db.WorkItemAssignedUser{
			AssignedUser: member,
			WorkItemID:   workItem.WorkItemID,
		}

		err := lookup.Delete(ctx, d)
		if err != nil {
			return internal.WrapErrorWithLocf(err, "", []string{strconv.Itoa(idx)}, "could not remove member %s", member)
		}
	}

	return nil
}

func (w *WorkItem) AssignTags(ctx context.Context, d db.DBTX, workItemID db.WorkItemID, tagIDs []db.WorkItemTagID) error {
	wi, err := w.repos.WorkItem.ByID(ctx, d, workItemID, db.WithWorkItemJoin(db.WorkItemJoins{Team: true}))
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

		err = w.repos.WorkItem.AssignTag(ctx, d, &db.WorkItemWorkItemTagCreateParams{
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

func (w *WorkItem) RemoveTags(ctx context.Context, d db.DBTX, workItem *db.WorkItem, tagIDs []db.WorkItemTagID) error {
	for idx, tagID := range tagIDs {
		lookup := &db.WorkItemWorkItemTag{
			WorkItemTagID: tagID,
			WorkItemID:    workItem.WorkItemID,
		}

		err := lookup.Delete(ctx, d)
		if err != nil {
			return internal.WrapErrorWithLocf(err, "", []string{strconv.Itoa(idx)}, "could not remove tag %d", tagID)
		}
	}

	return nil
}
