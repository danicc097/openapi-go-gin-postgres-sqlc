package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// DemoProjectWorkItem represents the repository used for interacting with DemoProjectWorkItem records.
type DemoProjectWorkItem struct {
	q *db.Queries
}

// NewDemoProjectWorkItem instantiates the DemoProjectWorkItem repository.
func NewDemoProjectWorkItem() *DemoProjectWorkItem {
	return &DemoProjectWorkItem{
		q: db.New(),
	}
}

var _ repos.DemoProjectWorkItem = (*DemoProjectWorkItem)(nil)

func (u *DemoProjectWorkItem) ByID(ctx context.Context, d db.DBTX, id int64, opts ...db.DemoProjectWorkItemSelectConfigOption) (*db.DemoProjectWorkItem, error) {
	return db.DemoProjectWorkItemByWorkItemID(ctx, d, id, opts...)
}

func (u *DemoProjectWorkItem) Create(ctx context.Context, d db.DBTX, params repos.DemoProjectWorkItemCreateParams) (*db.DemoProjectWorkItem, error) {
	workItem := &db.WorkItem{
		Title:          params.Base.Title,
		Description:    params.Base.Description,
		WorkItemTypeID: params.Base.WorkItemTypeID,
		Metadata:       params.Base.Metadata,
		TeamID:         params.Base.TeamID,
		KanbanStepID:   params.Base.KanbanStepID,
		Closed:         params.Base.Closed,
		TargetDate:     params.Base.TargetDate,
	}

	workItem, err := workItem.Save(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not save workItem: %w", parseErrorDetail(err))
	}

	demoProjectWorkItem := &db.DemoProjectWorkItem{
		WorkItemID:    workItem.WorkItemID,
		Ref:           params.DemoProject.Ref,
		Line:          params.DemoProject.Line,
		LastMessageAt: params.DemoProject.LastMessageAt,
		Reopened:      params.DemoProject.Reopened,
	}

	demoProjectWorkItem, err = demoProjectWorkItem.Save(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not save demoProjectWorkItem: %w", parseErrorDetail(err))
	}

	demoProjectWorkItem.WorkItem = workItem

	return demoProjectWorkItem, nil
}

func (u *DemoProjectWorkItem) Update(ctx context.Context, d db.DBTX, id int64, params repos.DemoProjectWorkItemUpdateParams) (*db.DemoProjectWorkItem, error) {
	demoProjectWorkItem, err := u.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get demoProjectWorkItem by id: %w", parseErrorDetail(err))
	}
	workItem, err := demoProjectWorkItem.FKWorkItem_WorkItemID(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not get associated workItem: %w", parseErrorDetail(err))
	}

	if params.Base != nil {
		updateEntityWithParams(workItem, params.Base)
	}

	if params.DemoProject != nil {
		updateEntityWithParams(demoProjectWorkItem, params.DemoProject)
	}

	workItem, err = workItem.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update workItem: %w", parseErrorDetail(err))
	}
	demoProjectWorkItem, err = demoProjectWorkItem.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update demoProjectWorkItem: %w", parseErrorDetail(err))
	}

	demoProjectWorkItem.WorkItem = workItem

	return demoProjectWorkItem, err
}

func (u *DemoProjectWorkItem) Delete(ctx context.Context, d db.DBTX, id int64) (*db.DemoProjectWorkItem, error) {
	workItem, err := db.WorkItemByWorkItemID(ctx, d, id, db.WithWorkItemJoin(db.WorkItemJoins{DemoProjectWorkItem: true}))
	if err != nil {
		return nil, fmt.Errorf("could not get workItem: %w", parseErrorDetail(err))
	}

	err = workItem.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete workItem: %w", parseErrorDetail(err))
	}

	return workItem.DemoProjectWorkItem, err
}
