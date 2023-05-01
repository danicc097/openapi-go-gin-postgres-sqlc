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

	demoWorkItem := &db.DemoWorkItem{
		WorkItemID:    workItem.WorkItemID,
		Ref:           params.DemoProject.Ref,
		Line:          params.DemoProject.Line,
		LastMessageAt: params.DemoProject.LastMessageAt,
		Reopened:      params.DemoProject.Reopened,
	}

	demoWorkItem, err = demoWorkItem.Save(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not save demoWorkItem: %w", parseErrorDetail(err))
	}

	demoWorkItem.WorkItemJoin = workItem

	return demoWorkItem, nil
}

func (u *DemoWorkItem) Update(ctx context.Context, d db.DBTX, id int64, params repos.DemoWorkItemUpdateParams) (*db.DemoWorkItem, error) {
	demoWorkItem, err := u.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get demoWorkItem by id: %w", parseErrorDetail(err))
	}
	workItem, err := demoWorkItem.FKWorkItem_WorkItemID(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not get associated workItem: %w", parseErrorDetail(err))
	}

	if params.Base != nil {
		updateEntityWithParams(workItem, params.Base)
	}

	if params.DemoProject != nil {
		updateEntityWithParams(demoWorkItem, params.DemoProject)
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
	workItem, err := db.WorkItemByWorkItemID(ctx, d, id, db.WithWorkItemJoin(db.WorkItemJoins{DemoWorkItem: true}))
	if err != nil {
		return nil, fmt.Errorf("could not get workItem: %w", parseErrorDetail(err))
	}

	err = workItem.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete workItem: %w", parseErrorDetail(err))
	}

	return workItem.DemoWorkItemJoin, err
}
