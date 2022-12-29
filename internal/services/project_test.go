package services

import (
	"reflect"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
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
		name string
		args args
		want models.ProjectConfig
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.MergeConfigFields("1", tt.args.obj2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
