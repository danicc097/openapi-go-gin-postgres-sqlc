package postgresqltestutil

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

func RandomDemoProjectWorkItemCreateParams(t *testing.T, workItemID int64) db.DemoProjectWorkItemCreateParams {
	t.Helper()

	return db.DemoProjectWorkItemCreateParams{
		WorkItemID:    workItemID,
		Ref:           "ref-" + testutil.RandomString(5),
		Line:          "line-" + testutil.RandomString(5),
		Reopened:      testutil.RandomBool(),
		LastMessageAt: testutil.RandomDate(),
	}
}
