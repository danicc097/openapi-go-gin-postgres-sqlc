package servicetestutil

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqlrandom"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/stretchr/testify/require"
)

type CreateWorkItemParams struct{}

type CreateWorkItemFixture struct {
	*models.WorkItem
}

// CreateWorkItem creates a new random work item comment with the given configuration.
func (ff *FixtureFactory) CreateWorkItem(ctx context.Context, project models.ProjectName, caller services.CtxUser, teamID models.TeamID) *CreateWorkItemFixture {
	var workItem *models.WorkItem
	var err error

	switch project {
	case models.ProjectNameDemo:
		p := postgresqlrandom.DemoWorkItemCreateParams(
			postgresqlrandom.KanbanStepID(project),
			postgresqlrandom.WorkItemTypeID(project),
			teamID,
		)
		workItem, err = ff.svc.DemoWorkItem.Create(ctx, ff.d, caller, services.DemoWorkItemCreateParams{
			DemoWorkItemCreateParams: p,
		})
		require.NoError(ff.t, err)
	case models.ProjectNameDemoTwo:
		p := postgresqlrandom.DemoTwoWorkItemCreateParams(
			postgresqlrandom.KanbanStepID(project),
			postgresqlrandom.WorkItemTypeID(project),
			teamID,
		)
		workItem, err = ff.svc.DemoTwoWorkItem.Create(ctx, ff.d, caller, services.DemoTwoWorkItemCreateParams{
			DemoTwoWorkItemCreateParams: p,
		})
		require.NoError(ff.t, err)
	}

	return &CreateWorkItemFixture{
		WorkItem: workItem,
	}
}
