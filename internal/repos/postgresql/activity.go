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

func (u *Activity) Create(ctx context.Context, d db.DBTX, params repos.ActivityCreateParams) (*db.Activity, error) {
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

func (u *Activity) Update(ctx context.Context, d db.DBTX, id int, params repos.ActivityUpdateParams) (*db.Activity, error) {
	workItemTag, err := u.ActivityByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemTag by id %w", parseErrorDetail(err))
	}

	if params.Description != nil {
		workItemTag.Description = *params.Description
	}
	if params.Name != nil {
		workItemTag.Name = *params.Name
	}
	if params.IsProductive != nil {
		workItemTag.IsProductive = *params.IsProductive
	}

	_, err = workItemTag.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update workItemTag: %w", parseErrorDetail(err))
	}

	return workItemTag, err
}

func (u *Activity) ActivityByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.Activity, error) {
	workItemTag, err := db.ActivityByNameProjectID(ctx, d, name, projectID)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemTag: %w", parseErrorDetail(err))
	}

	return workItemTag, nil
}

func (u *Activity) ActivityByID(ctx context.Context, d db.DBTX, id int) (*db.Activity, error) {
	workItemTag, err := db.ActivityByActivityID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemTag: %w", parseErrorDetail(err))
	}

	return workItemTag, nil
}

func (u *Activity) Delete(ctx context.Context, d db.DBTX, id int) (*db.Activity, error) {
	workItemTag, err := u.ActivityByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get workItemTag by id %w", parseErrorDetail(err))
	}

	err = workItemTag.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete workItemTag: %w", parseErrorDetail(err))
	}

	return workItemTag, err
}
