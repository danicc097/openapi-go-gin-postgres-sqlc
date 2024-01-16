package servicetestutil

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/stretchr/testify/require"
)

type CreateWorkItemCommentParams struct {
	UserID  db.UserID
	Project models.Project
	TagIDs  []db.WorkItemTagID
	Members []services.Member
}

type CreateWorkItemCommentFixture struct {
	WorkItemComment *db.WorkItemComment
	WorkItem        *db.WorkItem
}

// CreateWorkItemComment creates a new random work item comment with the given configuration.
func (ff *FixtureFactory) CreateWorkItemComment(ctx context.Context, params CreateWorkItemCommentParams) (*CreateWorkItemCommentFixture, error) {
	teamf, err := ff.CreateTeam(ctx, CreateTeamParams{Project: params.Project})
	require.NoError(ff.t, err)
	kanbanStepID := internal.DemoKanbanStepsIDByName[testutil.RandomFrom(models.AllDemoKanbanStepsValues())]
	workItemTypeID := internal.DemoWorkItemTypesIDByName[testutil.RandomFrom(models.AllDemoWorkItemTypesValues())]

	// TODO: will be ff.Create{Projectname}WorkItem!
	var workItem *db.WorkItem
	switch params.Project {
	case models.ProjectDemo:
		params := postgresqltestutil.RandomDemoWorkItemCreateParams(ff.t, kanbanStepID, workItemTypeID, teamf.Team.TeamID)
		workItem, err = ff.svc.DemoWorkItem.Create(ctx, ff.d, services.DemoWorkItemCreateParams{
			DemoWorkItemCreateParams: params,
		})
	case models.ProjectDemoTwo:
		workItem = postgresqltestutil.NewRandomDemoTwoWorkItem(ff.t, ff.d)
	}

	randomRepoCreateParams := postgresqltestutil.RandomWorkItemCommentCreateParams(ff.t, params.UserID, workItem.WorkItemID)
	// don't use repos for test fixtures, useservice logic
	workItemComment, err := ff.svc.WorkItemComment.Create(ctx, ff.d, randomRepoCreateParams)
	if err != nil {
		return nil, fmt.Errorf("svc.WorkItemComment.Create: %w", err)
	}

	return &CreateWorkItemCommentFixture{
		WorkItemComment: workItemComment,
	}, nil
}
