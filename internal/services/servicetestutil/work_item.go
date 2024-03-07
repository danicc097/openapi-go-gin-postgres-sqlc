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
	Caller  services.CtxUser
	TeamID  db.TeamID
}

type CreateWorkItemFixture struct {
	*db.WorkItem
}

// CreateWorkItem creates a new random work item comment with the given configuration.
func (ff *FixtureFactory) CreateWorkItem(ctx context.Context, params CreateWorkItemParams) *CreateWorkItemFixture {
	var workItem *db.WorkItem
	var err error

	switch params.Project {
	case models.ProjectDemo:
		p := postgresqlrandom.DemoWorkItemCreateParams(
			postgresqlrandom.KanbanStepID(params.Project),
			postgresqlrandom.WorkItemTypeID(params.Project),
			params.TeamID,
		)
		workItem, err = ff.svc.DemoWorkItem.Create(ctx, ff.d, params.Caller, services.DemoWorkItemCreateParams{
			DemoWorkItemCreateParams: p,
		})
		require.NoError(ff.t, err)
	case models.ProjectDemoTwo:
		p := postgresqlrandom.DemoTwoWorkItemCreateParams(
			postgresqlrandom.KanbanStepID(params.Project),
			postgresqlrandom.WorkItemTypeID(params.Project),
			params.TeamID,
		)
		workItem, err = ff.svc.DemoTwoWorkItem.Create(ctx, ff.d, params.Caller, services.DemoTwoWorkItemCreateParams{
			DemoTwoWorkItemCreateParams: p,
		})
		require.NoError(ff.t, err)
	}

	return &CreateWorkItemFixture{
		WorkItem: workItem,
	}
}
