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

func (a *Activity) Create(ctx context.Context, d db.DBTX, params *db.ActivityCreateParams) (*db.Activity, error) {
	activity, err := db.CreateActivity(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not create activity: %w", parseErrorDetail(err))
	}

	return activity, nil
}

func (a *Activity) Update(ctx context.Context, d db.DBTX, id int, params *db.ActivityUpdateParams) (*db.Activity, error) {
	activity, err := a.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get activity by id %w", parseErrorDetail(err))
	}

	activity, err = activity.Update(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not update activity: %w", parseErrorDetail(err))
	}

	return activity, err
}

func (a *Activity) ByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.Activity, error) {
	activity, err := db.ActivityByNameProjectID(ctx, d, name, projectID)
	if err != nil {
		return nil, fmt.Errorf("could not get activity: %w", parseErrorDetail(err))
	}

	return activity, nil
}

func (a *Activity) ByProjectID(ctx context.Context, d db.DBTX, projectID int) ([]db.Activity, error) {
	activities, err := db.ActivitiesByProjectID(ctx, d, projectID)
	if err != nil {
		return nil, fmt.Errorf("could not get activity: %w", parseErrorDetail(err))
	}

	return activities, nil
}

func (a *Activity) ByID(ctx context.Context, d db.DBTX, id int) (*db.Activity, error) {
	activity, err := db.ActivityByActivityID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get activity: %w", parseErrorDetail(err))
	}

	return activity, nil
}

func (a *Activity) Delete(ctx context.Context, d db.DBTX, id int) (*db.Activity, error) {
	activity, err := a.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get activity by id %w", parseErrorDetail(err))
	}

	err = activity.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete activity: %w", parseErrorDetail(err))
	}

	return activity, err
}
