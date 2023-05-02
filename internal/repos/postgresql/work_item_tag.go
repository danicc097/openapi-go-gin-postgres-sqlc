package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// WorkItemTag represents the repository used for interacting with WorkItemTag records.
type WorkItemTag struct {
	q *db.Queries
}

// NewWorkItemTag instantiates the WorkItemTag repository.
func NewWorkItemTag() *WorkItemTag {
	return &WorkItemTag{
		q: db.New(),
	}
}

var _ repos.WorkItemTag = (*WorkItemTag)(nil)

func (wit *WorkItemTag) Create(ctx context.Context, d db.DBTX, params db.WorkItemTagCreateParams) (*db.WorkItemTag, error) {
	activity := &db.WorkItemTag{
		Name:        params.Name,
		Description: params.Description,
		ProjectID:   params.ProjectID,
		Color:       params.Color,
	}

	if _, err := activity.Insert(ctx, d); err != nil {
		return nil, err
	}

	return activity, nil
}

func (wit *WorkItemTag) Update(ctx context.Context, d db.DBTX, id int, params db.WorkItemTagUpdateParams) (*db.WorkItemTag, error) {
	activity, err := wit.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get activity by id %w", parseErrorDetail(err))
	}

	updateEntityWithParams(activity, &params)

	activity, err = activity.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update activity: %w", parseErrorDetail(err))
	}

	return activity, err
}

func (wit *WorkItemTag) ByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.WorkItemTag, error) {
	activity, err := db.WorkItemTagByNameProjectID(ctx, d, name, projectID)
	if err != nil {
		return nil, fmt.Errorf("could not get activity: %w", parseErrorDetail(err))
	}

	return activity, nil
}

func (wit *WorkItemTag) ByID(ctx context.Context, d db.DBTX, id int) (*db.WorkItemTag, error) {
	activity, err := db.WorkItemTagByWorkItemTagID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get activity: %w", parseErrorDetail(err))
	}

	return activity, nil
}

func (wit *WorkItemTag) Delete(ctx context.Context, d db.DBTX, id int) (*db.WorkItemTag, error) {
	activity, err := wit.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get activity by id %w", parseErrorDetail(err))
	}

	err = activity.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete activity: %w", parseErrorDetail(err))
	}

	return activity, err
}
