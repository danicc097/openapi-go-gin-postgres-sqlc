package postgresql_test

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

func TestNormalizeIndexDef(t *testing.T) {
	type testCase struct {
		name        string
		inputString string
		want        string
		wantErr     bool
	}

	tests := []testCase{
		{
			name: "normalizes stored pg index defs",
			inputString: `create index test on public.cache__demo_work_items using gin (
		description gin_trgm_ops  
		, last_message_at, ref ,"SomeColumn" 
		, reopened);`,
			want: `create index test on public.cache__demo_work_items using gin(description gin_trgm_ops,last_message_at,ref,"somecolumn",reopened)`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := models.NormalizeIndexDef(context.Background(), testPool, tt.inputString)
			if (err != nil) != tt.wantErr {
				t.Errorf("NormalizeIndexDef() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NormalizeIndexDef() = %v, want %v", got, tt.want)
			}
		})
	}
}
