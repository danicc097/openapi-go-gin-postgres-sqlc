package postgresql

import (
	"context"
	"fmt"

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

func (u *DemoWorkItem) ByID(ctx context.Context, d db.DBTX, id int64, opts ...db.DemoWorkItemSelectConfigOption) (*db.DemoWorkItem, error) {
	return db.DemoWorkItemByWorkItemID(ctx, d, id, opts...)
}

func (u *DemoWorkItem) Create(ctx context.Context, d db.DBTX, params repos.DemoWorkItemCreateParams) (*db.DemoWorkItem, error) {
	workItem, err := db.CreateWorkItem(ctx, d, &params.Base)
	if err != nil {
		return nil, fmt.Errorf("could not create workItem: %w", parseErrorDetail(err))
	}

	dwicp := &db.DemoWorkItemCreateParams{
		WorkItemID:    workItem.WorkItemID,
		Ref:           params.DemoProject.Ref,
		Line:          params.DemoProject.Line,
		LastMessageAt: params.DemoProject.LastMessageAt,
		Reopened:      params.DemoProject.Reopened,
	}

	demoWorkItem, err := db.CreateDemoWorkItem(ctx, d, dwicp)
	if err != nil {
		return nil, fmt.Errorf("could not create workItem: %w", parseErrorDetail(err))
	}

	demoWorkItem.WorkItemJoin = workItem

	return demoWorkItem, nil
}

func (u *DemoWorkItem) Update(ctx context.Context, d db.DBTX, id int64, params repos.DemoWorkItemUpdateParams) (*db.DemoWorkItem, error) {
	demoWorkItem, err := u.ByID(ctx, d, id, db.WithDemoWorkItemJoin(db.DemoWorkItemJoins{WorkItem: true}))
	if err != nil {
		return nil, fmt.Errorf("could not get demoWorkItem by id: %w", parseErrorDetail(err))
	}
	workItem := demoWorkItem.WorkItemJoin

	if params.Base != nil {
		workItem.SetUpdateParams(params.Base)
	}

	if params.DemoProject != nil {
		demoWorkItem.SetUpdateParams(params.DemoProject)
	}

	workItem, err = workItem.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update workItem: %w", parseErrorDetail(err))
	}
	demoWorkItem, err = demoWorkItem.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update demoWorkItem: %w", parseErrorDetail(err))
	}

	demoWorkItem.WorkItemJoin = workItem

	return demoWorkItem, err
}

func (u *DemoWorkItem) Delete(ctx context.Context, d db.DBTX, id int64) (*db.DemoWorkItem, error) {
	workItem, err := db.WorkItemByWorkItemID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItem: %w", parseErrorDetail(err))
	}

	err = workItem.Delete(ctx, d) // cascades. PK is FK
	if err != nil {
		return nil, fmt.Errorf("could not delete workItem: %w", parseErrorDetail(err))
	}

	return workItem.DemoWorkItemJoin, err
}
