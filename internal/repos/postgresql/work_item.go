package postgresql

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// WorkItem represents the repository used for interacting with WorkItem records.
type WorkItem struct {
	q *db.Queries
}

// NewWorkItem instantiates the WorkItem repository.
func NewWorkItem() *WorkItem {
	return &WorkItem{
		q: db.New(),
	}
}

var _ repos.WorkItem = (*WorkItem)(nil)

func (u *WorkItem) ByID(ctx context.Context, d db.DBTX, id int64, opts ...db.WorkItemSelectConfigOption) (*db.WorkItem, error) {
	return db.WorkItemByWorkItemID(ctx, d, id, opts...)
}
