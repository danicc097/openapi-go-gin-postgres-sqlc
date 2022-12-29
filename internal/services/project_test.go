package services

import (
	"reflect"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/models"
)

func Test_mergeFields(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeFields(tt.args.obj1, tt.args.obj2, tt.args.pathKeys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
