package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// Activity represents the repository used for interacting with Activity records.
type Activity struct {
	q *db.Queries
}

// NewActivity instantiates the Activity repository.
func NewActivity() *Activity {
	return &Activity{
		q: db.New(),
	}
}

var _ repos.Activity = (*Activity)(nil)

func (a *Activity) Create(ctx context.Context, d db.DBTX, params db.ActivityCreateParams) (*db.Activity, error) {
	workItemTag := &db.Activity{
		Name:         params.Name,
		Description:  params.Description,
		ProjectID:    params.ProjectID,
		IsProductive: params.IsProductive,
	}

	if _, err := workItemTag.Save(ctx, d); err != nil {
		return nil, err
	}

	return workItemTag, nil
}

func (a *Activity) Update(ctx context.Context, d db.DBTX, id int, params db.ActivityUpdateParams) (*db.Activity, error) {
	workItemTag, err := a.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemTag by id %w", parseErrorDetail(err))
	}

	updateEntityWithParams(workItemTag, &params)

	workItemTag, err = workItemTag.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update workItemTag: %w", parseErrorDetail(err))
	}

	return workItemTag, err
}

func (a *Activity) ByProjectID(ctx context.Context, d db.DBTX, name string, projectID int) (*db.Activity, error) {
	workItemTag, err := db.ActivityByNameProjectID(ctx, d, name, projectID)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemTag: %w", parseErrorDetail(err))
	}

	return workItemTag, nil
}

func (a *Activity) ByID(ctx context.Context, d db.DBTX, id int) (*db.Activity, error) {
	workItemTag, err := db.ActivityByActivityID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemTag: %w", parseErrorDetail(err))
	}

	return workItemTag, nil
}

func (a *Activity) Delete(ctx context.Context, d db.DBTX, id int) (*db.Activity, error) {
	workItemTag, err := a.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemTag by id %w", parseErrorDetail(err))
	}

	err = workItemTag.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete workItemTag: %w", parseErrorDetail(err))
	}

	return workItemTag, err
}
