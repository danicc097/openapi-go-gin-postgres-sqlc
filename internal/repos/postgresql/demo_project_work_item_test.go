package postgresql_test

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	internalmodels "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/stretchr/testify/assert"
)

// TODO test Create, etc. completely

func TestDemoWorkItem_Update(t *testing.T) {
	t.Parallel()

	projectID := internal.ProjectIDByName[internalmodels.ProjectDemo]
	team, _ := postgresqltestutil.NewRandomTeam(t, testPool, projectID)

	kanbanStepID := internal.DemoKanbanStepsIDByName[internalmodels.DemoKanbanStepsReceived]
	workItemTypeID := internal.DemoWorkItemTypesIDByName[internalmodels.DemoWorkItemTypesType1]
	demoWorkItem, _ := postgresqltestutil.NewRandomDemoWorkItem(t, testPool, projectID, kanbanStepID, workItemTypeID, team.TeamID)

	type args struct {
		id     int64
		params repos.DemoWorkItemUpdateParams
	}
	type params struct {
		name    string
		args    args
		want    *db.DemoWorkItem
		wantErr bool
	}
	tests := []params{
		{
			name: "updated",
			args: args{
				id: demoWorkItem.WorkItemID,
				params: repos.DemoWorkItemUpdateParams{
					Base:        &db.WorkItemUpdateParams{Description: pointers.New("new description")},
					DemoProject: &db.DemoWorkItemUpdateParams{Line: pointers.New("new line")},
				},
			},
			want: func() *db.DemoWorkItem {
				u := *demoWorkItem
				u.WorkItem.Description = "new description"
				u.Line = "new line"

				return &u
			}(),
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			u := postgresql.NewDemoWorkItem()
			got, err := u.Update(context.Background(), testPool, tc.args.id, tc.args.params)
			if (err != nil) != tc.wantErr {
				t.Errorf("DemoWorkItem.Update() error = %v, wantErr %v", err, tc.wantErr)

				return
			}

			got.WorkItem.UpdatedAt = demoWorkItem.WorkItem.UpdatedAt // ignore
			assert.Equal(t, tc.want, got)
		})
	}
}
