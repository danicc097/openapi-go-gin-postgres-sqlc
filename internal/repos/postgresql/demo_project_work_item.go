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
	wi := &db.WorkItem{
		Title:          params.Base.Title,
		Description:    params.Base.Description,
		WorkItemTypeID: params.Base.WorkItemTypeID,
		Metadata:       params.Base.Metadata,
		TeamID:         params.Base.TeamID,
		KanbanStepID:   params.Base.KanbanStepID,
		Closed:         params.Base.Closed,
		TargetDate:     params.Base.TargetDate,
	}

	wi, err := wi.Save(ctx, d)
	if err != nil {
		return nil, err
	}

	dpwi := &db.DemoProjectWorkItem{
		WorkItemID:    wi.WorkItemID,
		Ref:           params.DemoProject.Ref,
		Line:          params.DemoProject.Line,
		LastMessageAt: params.DemoProject.LastMessageAt,
		Reopened:      params.DemoProject.Reopened,
	}

	dpwi, err = dpwi.Save(ctx, d)
	if err != nil {
		return nil, err
	}

	return dpwi, nil
}

func (u *DemoProjectWorkItem) Update(ctx context.Context, d db.DBTX, params repos.DemoProjectWorkItemUpdateParams) (*db.DemoProjectWorkItem, error) {
	demoProjectWorkItem, err := u.ByID(ctx, d, *params.DemoProject.WorkItemID)
	if err != nil {
		return nil, fmt.Errorf("could not get demoProjectWorkItem by id: %w", parseErrorDetail(err))
	}
	workItem, err := demoProjectWorkItem.FKWorkItem_WorkItemID(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not get associated workItem: %w", parseErrorDetail(err))
	}

	if params.Base != nil {
		workItem.Closed = params.Base.Closed

		if params.Base.Description != nil {
			workItem.Description = *params.Base.Description
		}
		if params.Base.TargetDate != nil {
			workItem.TargetDate = *params.Base.TargetDate
		}
		if params.Base.KanbanStepID != nil {
			workItem.KanbanStepID = *params.Base.KanbanStepID
		}
		if params.Base.Metadata != nil {
			workItem.Metadata = *params.Base.Metadata
		}
		if params.Base.Title != nil {
			workItem.Title = *params.Base.Title
		}
		if params.Base.WorkItemTypeID != nil {
			workItem.WorkItemTypeID = *params.Base.WorkItemTypeID
		}
	}

	if params.DemoProject != nil {
		if params.DemoProject.LastMessageAt != nil {
			demoProjectWorkItem.LastMessageAt = *params.DemoProject.LastMessageAt
		}
		if params.DemoProject.Line != nil {
			demoProjectWorkItem.Line = *params.DemoProject.Line
		}
		if params.DemoProject.Ref != nil {
			demoProjectWorkItem.Ref = *params.DemoProject.Ref
		}
		if params.DemoProject.Reopened != nil {
			demoProjectWorkItem.Reopened = *params.DemoProject.Reopened
		}
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

func (u *DemoProjectWorkItem) Delete(ctx context.Context, d db.DBTX, workItemID int64) (*db.DemoProjectWorkItem, error) {
	workItem, err := db.WorkItemByWorkItemID(ctx, d, workItemID, db.WithWorkItemJoin(db.WorkItemJoins{DemoProjectWorkItem: true}))
	if err != nil {
		return nil, fmt.Errorf("could not get workItem: %w", parseErrorDetail(err))
	}

	err = workItem.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete workItem: %w", parseErrorDetail(err))
	}

	return workItem.DemoProjectWorkItem, err
}
