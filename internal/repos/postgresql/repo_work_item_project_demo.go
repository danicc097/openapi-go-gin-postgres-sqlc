package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// DemoWorkItem represents the repository used for interacting with DemoWorkItem records.
type DemoWorkItem struct {
	q *db.Queries
}

// NewDemoWorkItem instantiates the DemoWorkItem repository.
func NewDemoWorkItem() *DemoWorkItem {
	return &DemoWorkItem{
		q: db.New(),
	}
}

var _ repos.DemoWorkItem = (*DemoWorkItem)(nil)

func (u *DemoWorkItem) ByID(ctx context.Context, d db.DBTX, id int, opts ...db.WorkItemSelectConfigOption) (*db.WorkItem, error) {
	extraOpts := []db.WorkItemSelectConfigOption{db.WithWorkItemJoin(db.WorkItemJoins{DemoWorkItem: true})}
	return db.WorkItemByWorkItemID(ctx, d, id, (append(extraOpts, opts...))...)
}

func (u *DemoWorkItem) Create(ctx context.Context, d db.DBTX, params repos.DemoWorkItemCreateParams) (*db.WorkItem, error) {
	workItem, err := db.CreateWorkItem(ctx, d, &params.Base)
	if err != nil {
		return nil, internal.WrapErrorWithLocf(parseDBErrorDetail(err), "", []string{"base"}, "could not create workItem")
	}

	params.DemoProject.WorkItemID = workItem.WorkItemID
	demoWorkItem, err := db.CreateDemoWorkItem(ctx, d, &params.DemoProject)
	if err != nil {
		return nil, internal.WrapErrorWithLocf(parseDBErrorDetail(err), "", []string{"demoWorkItem"}, "could not create demoWorkItem")
	}

	workItem.DemoWorkItemJoin = demoWorkItem

	return workItem, nil
}

func (u *DemoWorkItem) Update(ctx context.Context, d db.DBTX, id int, params repos.DemoWorkItemUpdateParams) (*db.WorkItem, error) {
	workItem, err := u.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItem by id: %w", parseDBErrorDetail(err))
	}
	demoWorkItem := workItem.DemoWorkItemJoin

	if params.Base != nil {
		workItem.SetUpdateParams(params.Base)
	}

	if params.DemoProject != nil {
		demoWorkItem.SetUpdateParams(params.DemoProject)
	}

	workItem, err = workItem.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update workItem: %w", parseDBErrorDetail(err))
	}
	demoWorkItem, err = demoWorkItem.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update demoWorkItem: %w", parseDBErrorDetail(err))
	}

	workItem.DemoWorkItemJoin = demoWorkItem

	return workItem, err
}
