package postgresql

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// DemoProjectWorkItem represents the repository used for interacting with DemoProjectWorkItem records.
type DemoProjectWorkItem struct {
	q *db.Queries
}

// NewWorkItem instantiates the DemoProjectWorkItem repository.
func NewWorkItem() *DemoProjectWorkItem {
	return &DemoProjectWorkItem{
		q: db.New(),
	}
}

var _ repos.DemoProjectWorkItem = (*DemoProjectWorkItem)(nil)

func (u *DemoProjectWorkItem) WorkItemByID(ctx context.Context, d db.DBTX, id int64, opts ...db.DemoProjectWorkItemSelectConfigOption) (*db.DemoProjectWorkItem, error) {
	return db.DemoProjectWorkItemByWorkItemID(ctx, d, id, opts...)
}
