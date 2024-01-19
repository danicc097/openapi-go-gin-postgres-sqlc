package servicetestutil

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqlrandom"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/stretchr/testify/require"
)

type CreateWorkItemParams struct {
	Project models.Project
}

type CreateWorkItemFixture struct {
	WorkItem *db.WorkItem
}

// CreateWorkItem creates a new random work item comment with the given configuration.
func (ff *FixtureFactory) CreateWorkItem(ctx context.Context, params CreateWorkItemParams) *CreateWorkItemFixture {
	teamf := ff.CreateTeam(ctx, CreateTeamParams{Project: params.Project})

	var workItem *db.WorkItem
	var err error

	switch params.Project {
	case models.ProjectDemo:
		params := postgresqlrandom.DemoWorkItemCreateParams(
			postgresqlrandom.KanbanStepID(params.Project),
			postgresqlrandom.WorkItemTypeID(params.Project),
			teamf.Team.TeamID,
		)
		workItem, err = ff.svc.DemoWorkItem.Create(ctx, ff.d, services.DemoWorkItemCreateParams{
			DemoWorkItemCreateParams: params,
		})
		require.NoError(ff.t, err)
	case models.ProjectDemoTwo:
		params := postgresqlrandom.DemoTwoWorkItemCreateParams(
			postgresqlrandom.KanbanStepID(params.Project),
			postgresqlrandom.WorkItemTypeID(params.Project),
			teamf.Team.TeamID,
		)
		workItem, err = ff.svc.DemoTwoWorkItem.Create(ctx, ff.d, services.DemoTwoWorkItemCreateParams{
			DemoTwoWorkItemCreateParams: params,
		})
		require.NoError(ff.t, err)
	}

	return &CreateWorkItemFixture{
		WorkItem: workItem,
	}
}
