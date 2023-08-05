package services

import (
	"context"
	"errors"
	"fmt"

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
AssignMembers
RemoveMembers

it will accept projectName (tags) and teamID (members) if necessary, e.g. to ensure asigntag called with
tag belonging to project and a member that belongs to the team
(notice how seeming redundancy between repo/service starts to lose strength)

*/

type WorkItem struct {
	logger   *zap.SugaredLogger
	wiRepo   repos.WorkItem
	userRepo repos.User
}

// NewWorkItem returns a new WorkItem service with common logic for all project workitems.
func NewWorkItem(logger *zap.SugaredLogger, wiRepo repos.WorkItem, userRepo repos.User) *WorkItem {
	return &WorkItem{
		logger:   logger,
		wiRepo:   wiRepo,
		userRepo: userRepo,
	}
}

func (w *WorkItem) AssignWorkItemMembers(ctx context.Context, d db.DBTX, workItem *db.WorkItem, members []Member) error {
	for _, member := range members {
		user, err := w.userRepo.ByID(ctx, d, member.UserID, db.WithUserJoin(db.UserJoins{TeamsMember: true}))
		if err != nil {
			// TODO: WrapErrorWithLocf utility func that checks for wrapped internal.Error and appends a given `path` to loc.
			// in this case, appends the index of the members array.
			return internal.WrapErrorf(err, models.ErrorCodeNotFound, "user with id %s not found", member.UserID)
		}

		var userInTeam bool
		for _, team := range *user.MemberTeamsJoin {
			if team.TeamID == workItem.TeamID {
				userInTeam = true
			}
		}
		if !userInTeam {
			return internal.NewErrorf(models.ErrorCodeUnauthorized, "user %q does not belong to team %q", user.Email, workItem.TeamID)
		}

		// TODO: use wiRepo instead
		_, err = db.CreateWorkItemAssignedUser(ctx, d, &db.WorkItemAssignedUserCreateParams{
			AssignedUser: member.UserID,
			WorkItemID:   workItem.WorkItemID,
			Role:         member.Role,
		})
		var ierr *internal.Error
		if err != nil {
			if errors.As(err, &ierr) && ierr.Code() == models.ErrorCodeAlreadyExists {
				continue
			}

			return fmt.Errorf("could not assign member %s: %w", member.UserID, err)
		}
	}

	return nil
}
