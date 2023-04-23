package postgresql_test

import (
	"context"
	"testing"

	internalmodels "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/stretchr/testify/assert"
)

// TODO test Create, etc. completely

func TestDemoProjectWorkItem_Update(t *testing.T) {
	t.Parallel()

	projectRepo := postgresql.NewProject()

	ctx := context.Background()

	project, err := projectRepo.ByName(ctx, testPool, internalmodels.ProjectDemoProject)
	if err != nil {
		t.Fatalf("projectRepo.ByName unexpected error = %v", err)
	}
	// user, _ := postgresqltestutil.NewRandomUser(t, testPool)
	team, _ := postgresqltestutil.NewRandomTeam(t, testPool, project.ProjectID)
	workItemType, _ := postgresqltestutil.NewRandomWorkItemType(t, testPool, project.ProjectID)
	kanbanStep, _ := postgresqltestutil.NewRandomKanbanStep(t, testPool, project.ProjectID)
	demoprojectworkitem, _ := postgresqltestutil.NewRandomDemoProjectWorkItem(t, testPool, project.ProjectID, kanbanStep.KanbanStepID, workItemType.WorkItemTypeID, team.TeamID)

	type args struct {
		id     int64
		params repos.DemoProjectWorkItemUpdateParams
	}
	type params struct {
		name    string
		args    args
		want    *db.DemoProjectWorkItem
		wantErr bool
	}
	tests := []params{
		{
			name: "updated",
			args: args{
				id: demoprojectworkitem.WorkItemID,
				params: repos.DemoProjectWorkItemUpdateParams{
					Base:        &db.WorkItemUpdateParams{Description: pointers.New("new description")},
					DemoProject: &db.DemoProjectWorkItemUpdateParams{Line: pointers.New("new line")},
				},
			},
			want: func() *db.DemoProjectWorkItem {
				u := *demoprojectworkitem
				u.WorkItem().Description = "new description"
				u.Line = "new line"

				return &u
			}(),
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			u := postgresql.NewDemoProjectWorkItem()
			got, err := u.Update(context.Background(), testPool, tc.args.id, tc.args.params)
			if (err != nil) != tc.wantErr {
				t.Errorf("DemoProjectWorkItem.Update() error = %v, wantErr %v", err, tc.wantErr)

				return
			}

			got.WorkItem().UpdatedAt = demoprojectworkitem.WorkItem().UpdatedAt // ignore
			assert.Equal(t, tc.want, got)
		})
	}
}
