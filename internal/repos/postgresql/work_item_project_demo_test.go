package postgresql_test

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/stretchr/testify/assert"
)

func TestDemoTwoWorkItem_Update(t *testing.T) {
	t.Parallel()

	projectID := internal.ProjectIDByName[models.ProjectDemo]
	team, _ := postgresqltestutil.NewRandomTeam(t, testPool, projectID)

	kanbanStepID := internal.DemoKanbanStepsIDByName[models.DemoKanbanStepsReceived]
	workItemTypeID := internal.DemoTwoWorkItemTypesIDByName[models.DemoTwoWorkItemTypesType1]
	demoWorkItem, _ := postgresqltestutil.NewRandomDemoTwoWorkItem(t, testPool, kanbanStepID, workItemTypeID, team.TeamID)

	type args struct {
		id     int64
		params repos.DemoTwoWorkItemUpdateParams
	}
	type params struct {
		name    string
		args    args
		want    *db.WorkItem
		wantErr bool
	}
	d := pointers.New(pointers.New(testutil.RandomLocalDate()))

	tests := []params{
		{
			name: "updated",
			args: args{
				id: demoWorkItem.WorkItemID,
				params: repos.DemoTwoWorkItemUpdateParams{
					Base:           &db.WorkItemUpdateParams{Description: pointers.New("new description")},
					DemoTwoProject: &db.DemoTwoWorkItemUpdateParams{CustomDateForProject2: d},
				},
			},
			want: func() *db.WorkItem {
				u := *demoWorkItem
				u.Description = "new description"
				u.DemoTwoWorkItemJoin.CustomDateForProject2 = *d

				return &u
			}(),
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			u := postgresql.NewDemoTwoWorkItem()
			got, err := u.Update(context.Background(), testPool, tc.args.id, tc.args.params)
			if (err != nil) != tc.wantErr {
				t.Errorf("DemoTwoWorkItem.Update() error = %v, wantErr %v", err, tc.wantErr)

				return
			}

			got.UpdatedAt = demoWorkItem.UpdatedAt // ignore
			assert.Equal(t, tc.want, got)
		})
	}
}
