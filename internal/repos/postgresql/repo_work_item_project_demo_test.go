package postgresql_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
)

func TestDemoWorkItem_Update(t *testing.T) {
	t.Parallel()

	demoWorkItem := newRandomDemoWorkItem(t, testPool)

	type args struct {
		id     models.WorkItemID
		params repos.DemoWorkItemUpdateParams
	}
	type params struct {
		name    string
		args    args
		want    *models.WorkItem
		wantErr bool
	}
	tests := []params{
		{
			name: "updated",
			args: args{
				id: demoWorkItem.WorkItemID,
				params: repos.DemoWorkItemUpdateParams{
					Base:        &models.WorkItemUpdateParams{Description: pointers.New("new description")},
					DemoProject: &models.DemoWorkItemUpdateParams{Line: pointers.New("new line")},
				},
			},
			want: func() *models.WorkItem {
				u := *demoWorkItem // copy
				u.Description = "new description"
				u.DemoWorkItemJoin.Line = "new line"

				return &u
			}(),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			u := postgresql.NewDemoWorkItem()
			got, err := u.Update(t.Context(), testPool, tc.args.id, tc.args.params)
			if (err != nil) != tc.wantErr {
				t.Errorf("DemoTwoWorkItem.Update() error = %v, wantErr %v", err, tc.wantErr)

				return
			}

			got.UpdatedAt = demoWorkItem.UpdatedAt // ignore
			assert.Equal(t, tc.want, got)
		})
	}
}
