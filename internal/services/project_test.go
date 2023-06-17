package services_test

import (
	"context"
	"sort"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/repostesting"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/google/go-cmp/cmp"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func Test_MergeConfigFields(t *testing.T) {
	t.Parallel()

	proj := &db.Project{
		Name: models.ProjectDemo,
		BoardConfig: models.ProjectConfig{
			Header: []string{"demoProject.ref", "workItemType"},
			Fields: []models.ProjectConfigField{
				{
					IsEditable:    true,
					ShowCollapsed: true,
					IsVisible:     true,
					Path:          "demoWorkItem",
					Name:          "Demo project",
				},
				{
					IsEditable:    true,
					ShowCollapsed: true,
					IsVisible:     true,
					Path:          "demoWorkItem.ref",
					Name:          "Reference",
				},
			},
		},
	}

	fakeProjectRepo := &repostesting.FakeProject{}
	fakeProjectRepo.ByIDStub = func(ctx context.Context, d db.DBTX, i int) (*db.Project, error) {
		return proj, nil
	}
	fakeProjectRepo.ByNameStub = func(ctx context.Context, d db.DBTX, p models.Project) (*db.Project, error) {
		return proj, nil
	}
	fakeTeamRepo := &repostesting.FakeTeam{}
	p := services.NewProject(zaptest.NewLogger(t).Sugar(), fakeProjectRepo, fakeTeamRepo)

	type args struct {
		update map[string]any
	}
	tests := []struct {
		name  string
		args  args
		want  *models.ProjectConfig
		error string
	}{
		// TODO: expand test cases with different stubs, test bad config in db/update request (no fields key, wrong type of array elements...)
		{
			name: "example",
			args: args{
				update: map[string]any{"fields": []any{ // []any to test proper conversion later on
					map[string]any{"path": "workItemTypeID", "name": "Updated", "isEditable": false},
					map[string]any{"path": "inexistent", "name": "inexistent"}, // will be ignored
				}},
			},
			want: &models.ProjectConfig{
				Header: []string{"demoProject.ref", "workItemType"},
				Fields: []models.ProjectConfigField{
					{IsEditable: false, ShowCollapsed: true, IsVisible: true, Path: "workItemTypeID", Name: "Updated"}, // updated

					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "description", Name: "description"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "teamID", Name: "teamID"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "updatedAt", Name: "updatedAt"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "kanbanStepID", Name: "kanbanStepID"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "targetDate", Name: "targetDate"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "workItemID", Name: "workItemID"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "createdAt", Name: "createdAt"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "title", Name: "title"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "demoWorkItem", Name: "Demo project"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "demoWorkItem.workItemID", Name: "workItemID"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "demoWorkItem.ref", Name: "Reference"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "demoWorkItem.reopened", Name: "reopened"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "demoWorkItem.lastMessageAt", Name: "lastMessageAt"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "demoWorkItem.line", Name: "line"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "closed", Name: "closed"},
				},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := p.MergeConfigFields(context.Background(), &pgxpool.Pool{}, models.ProjectDemo, tc.args.update)
			if (err != nil) && tc.error == "" {
				t.Fatalf("unexpected error = %v", err)
			}
			if tc.error != "" {
				if err == nil {
					t.Fatalf("expected error = '%v' but got nothing", tc.error)
				}
				assert.Equal(t, tc.error, err.Error())

				return
			}

			// opts := cmp.Options{
			// 	cmp.FilterPath(func(p cmp.Path) bool {
			// 		return p.Last().Type().Kind() == reflect.Slice &&
			// 			p.Last().Type().Elem().Kind() == reflect.Struct
			// 	},
			// 		cmp.Comparer(customSliceComparer),
			// 	),
			// }

			sort.SliceStable(tc.want.Fields, func(i, j int) bool {
				return tc.want.Fields[i].Path < tc.want.Fields[j].Path
			})
			sort.SliceStable(got.Fields, func(i, j int) bool {
				return got.Fields[i].Path < got.Fields[j].Path
			})

			if diff := cmp.Diff(tc.want.Fields, got.Fields); diff != "" {
				t.Errorf("Fields mismatch (-want +got):\n%s", diff)
			}

			sort.SliceStable(tc.want.Header, func(i, j int) bool {
				return tc.want.Header[i] < tc.want.Header[j]
			})
			sort.SliceStable(got.Header, func(i, j int) bool {
				return got.Header[i] < got.Header[j]
			})

			if diff := cmp.Diff(tc.want.Header, got.Header); diff != "" {
				t.Errorf("Header mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
