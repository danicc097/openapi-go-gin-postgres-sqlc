package services

import (
	"context"
	"sort"
	"testing"

	internalmodels "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/repostesting"
	"github.com/google/go-cmp/cmp"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func Test_MergeConfigFields(t *testing.T) {
	fakeProjectRepo := &repostesting.FakeProject{}
	fakeProjectRepo.ProjectByIDStub = func(ctx context.Context, d db.DBTX, i int) (*db.Project, error) {
		return &db.Project{
			BoardConfig: pgtype.JSONB{Bytes: []byte(`
		{
			"header": ["demoProject.ref", "workItemType"],
			"fields": [
				{
					"isEditable": true,
					"showCollapsed": true,
					"isVisible": true,
					"path": "workItemType",
					"name": "Type"
				},
				{
					"isEditable": true,
					"showCollapsed": true,
					"isVisible": true,
					"path": "demoProjectWorkItem",
					"name": "Demo project"
				},
				{
					"isEditable": true,
					"showCollapsed": true,
					"isVisible": true,
					"path": "demoProjectWorkItem.metadata",
					"name": "Metadata"
				}
			]
		}
		`)},
			Name: string(internalmodels.ProjectDemoProject),
		}, nil
	}
	fakeTeamRepo := &repostesting.FakeTeam{}
	p := NewProject(zaptest.NewLogger(t), fakeProjectRepo, fakeTeamRepo)

	type args struct {
		obj2 map[string]any
	}
	tests := []struct {
		name  string
		args  args
		want  *models.ProjectConfig
		error string
	}{
		// TODO: expand test cases with different stubs
		{
			name: "example",
			args: args{
				obj2: map[string]any{"fields": []map[string]any{
					{"path:": "test", "name": "test"},
					{"path:": "test.nested", "name": "test nested"},
				}},
			},
			want: &models.ProjectConfig{
				Header: []string{"demoProject.ref", "workItemType"},
				Fields: []models.ProjectConfigField{
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "workItemTypeID", Name: "workItemTypeID"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "description", Name: "description"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "metadata", Name: "metadata"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "teamID", Name: "teamID"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "updatedAt", Name: "updatedAt"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "kanbanStepID", Name: "kanbanStepID"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "targetDate", Name: "targetDate"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "workItemID", Name: "workItemID"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "createdAt", Name: "createdAt"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "title", Name: "title"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "demoProjectWorkItem", Name: "demoProjectWorkItem"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "demoProjectWorkItem.workItemID", Name: "workItemID"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "demoProjectWorkItem.ref", Name: "ref"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "demoProjectWorkItem.reopened", Name: "reopened"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "demoProjectWorkItem.lastMessageAt", Name: "lastMessageAt"},
					{IsEditable: true, ShowCollapsed: true, IsVisible: true, Path: "demoProjectWorkItem.line", Name: "line"},
				},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := p.MergeConfigFields(context.Background(), &pgxpool.Pool{}, 1, tc.args.obj2)
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
