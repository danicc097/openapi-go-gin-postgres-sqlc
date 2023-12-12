package services

import (
	"context"
	"errors"
	"strconv"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
)

/**
 *
 * TODO: generic logic here, eg
AssignTags
RemoveTags
AssignUsers
RemoveAssignedUsers

it will accept projectName (tags) and teamID (members) if necessary, e.g. to ensure asigntag called with
tag belonging to project and a member that belongs to the team
(notice how seeming redundancy between repo/service starts to lose strength)

*/

type WorkItem struct {
	logger *zap.SugaredLogger
	repos  *repos.Repos
}

// NewWorkItem returns a new WorkItem service with common logic for all project workitems.
func NewWorkItem(logger *zap.SugaredLogger, repos *repos.Repos) *WorkItem {
	return &WorkItem{
		logger: logger,
		repos:  repos,
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

func (w *WorkItem) AssignTags(ctx context.Context, d db.DBTX, projectName models.Project, workItemID db.WorkItemID, tagIDs []db.WorkItemTagID) error {
	for idx, tagID := range tagIDs {
		tag, err := w.repos.WorkItemTag.ByID(ctx, d, tagID)
		if err != nil {
			return internal.WrapErrorWithLocf(err, models.ErrorCodeNotFound, []string{strconv.Itoa(idx)}, "tag with id %d not found", tagID)
		}

		if internal.ProjectIDByName[projectName] != tag.ProjectID {
			return internal.WrapErrorWithLocf(nil, models.ErrorCodeUnauthorized, []string{strconv.Itoa(idx)}, "tag %q does not belong to project %q", tag.Name, tag.ProjectJoin.Name)
		}

		err = w.repos.WorkItem.AssignTag(ctx, d, &db.WorkItemWorkItemTagCreateParams{
			WorkItemTagID: tagID,
			WorkItemID:    workItemID,
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

func (w *WorkItem) RemoveTags(ctx context.Context, d db.DBTX, workItemID db.WorkItemID, tagIDs []db.WorkItemTagID) error {
	for idx, tagID := range tagIDs {
		lookup := &db.WorkItemWorkItemTag{
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
