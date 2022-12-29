package services

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func Test_MergeConfigFields(t *testing.T) {
	p := NewProject(zaptest.NewLogger(t), postgresql.NewProject(), postgresql.NewTeam())

	type args struct {
		obj1     models.ProjectConfig
		obj2     any
		pathKeys []string
	}
	tests := []struct {
		name  string
		args  args
		want  models.ProjectConfig
		error string
	}{
		// TODO: Add test cases.
		// - empty array of initialization keys
		// - array coming from GetStructKeys(boardConfig) where boardConfig is initialized with the nested fields we specify

		// {
		// 	name: "no initialization",
		// 	args: args{
		// 		obj1: models.ProjectConfig{
		// 			Fields: []models.ProjectConfigField{{}},
		// 		},
		// 		obj2: map[string]any{"fields": []map[string]any{
		// 			{"path:": "test", "name": "test"},
		// 			{"path:": "test.nested", "name": "test nested"},
		// 		}},
		// 		pathKeys: []string{""},
		// 	},
		// },
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
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

			assert.Equal(t, tc.want, got)
		})
	}
}
