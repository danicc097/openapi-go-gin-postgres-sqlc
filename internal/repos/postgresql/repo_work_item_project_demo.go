package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
)

// DemoWorkItem represents the repository used for interacting with DemoWorkItem records.
type DemoWorkItem struct {
	q models.Querier
}

// NewDemoWorkItem instantiates the DemoWorkItem repository.
func NewDemoWorkItem() *DemoWorkItem {
	return &DemoWorkItem{
		q: NewQuerierWrapper(models.New()),
	}
}

var _ repos.DemoWorkItem = (*DemoWorkItem)(nil)

func (u *DemoWorkItem) ByID(ctx context.Context, d models.DBTX, id models.WorkItemID, opts ...models.WorkItemSelectConfigOption) (*models.WorkItem, error) {
	extraOpts := []models.WorkItemSelectConfigOption{models.WithWorkItemJoin(models.WorkItemJoins{DemoWorkItem: true})}

	return models.WorkItemByWorkItemID(ctx, d, id, (append(extraOpts, opts...))...)
}

func (u *DemoWorkItem) Paginated(ctx context.Context, d models.DBTX, cursor models.WorkItemID, opts ...models.CacheDemoWorkItemSelectConfigOption) ([]models.CacheDemoWorkItem, error) {
	extraOpts := []models.CacheDemoWorkItemSelectConfigOption{models.WithCacheDemoWorkItemJoin(models.CacheDemoWorkItemJoins{})}

	// TODO: params models.GetPaginatedCacheDemoWorkItemParams instead of cursor
	c := models.PaginationCursor{Column: "workItemID", Value: pointers.New[interface{}](cursor), Direction: models.DirectionDesc}

	return models.CacheDemoWorkItemPaginated(ctx, d, c, (append(extraOpts, opts...))...)
}

func (u *DemoWorkItem) Create(ctx context.Context, d models.DBTX, params repos.DemoWorkItemCreateParams) (*models.WorkItem, error) {
	workItem, err := models.CreateWorkItem(ctx, d, &params.Base)
	if err != nil {
		return nil, internal.WrapErrorWithLocf(ParseDBErrorDetail(err), "", []string{"base"}, "could not create workItem")
	}

	params.DemoProject.WorkItemID = workItem.WorkItemID
	demoWorkItem, err := models.CreateDemoWorkItem(ctx, d, &params.DemoProject)
	if err != nil {
		return nil, internal.WrapErrorWithLocf(ParseDBErrorDetail(err), "", []string{"demoWorkItem"}, "could not create demoWorkItem")
	}

	workItem.DemoWorkItemJoin = demoWorkItem

	return workItem, nil
}

func (u *DemoWorkItem) Update(ctx context.Context, d models.DBTX, id models.WorkItemID, params repos.DemoWorkItemUpdateParams) (*models.WorkItem, error) {
	workItem, err := u.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItem by id: %w", ParseDBErrorDetail(err))
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
		return nil, fmt.Errorf("could not update workItem: %w", ParseDBErrorDetail(err))
	}
	demoWorkItem, err = demoWorkItem.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update demoWorkItem: %w", ParseDBErrorDetail(err))
	}

	workItem.DemoWorkItemJoin = demoWorkItem

	return workItem, err
}
