package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// Activity represents the repository used for interacting with Activity records.
type Activity struct {
	q db.Querier
}

// NewActivity instantiates the Activity repository.
func NewActivity() *Activity {
	return &Activity{
		q: NewQuerierWrapper(db.New()),
	}
}

var _ repos.Activity = (*Activity)(nil)

func (a *Activity) Create(ctx context.Context, d db.DBTX, params *db.ActivityCreateParams) (*db.Activity, error) {
	activity, err := db.CreateActivity(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not create activity: %w", ParseDBErrorDetail(err))
	}

	return activity, nil
}

func (a *Activity) Update(ctx context.Context, d db.DBTX, id db.ActivityID, params *db.ActivityUpdateParams) (*db.Activity, error) {
	activity, err := a.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get activity by id %w", ParseDBErrorDetail(err))
	}

	activity.SetUpdateParams(params)

	activity, err = activity.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update activity: %w", ParseDBErrorDetail(err))
	}

	return activity, err
}

func (a *Activity) ByName(ctx context.Context, d db.DBTX, name string, projectID db.ProjectID, opts ...db.ActivitySelectConfigOption) (*db.Activity, error) {
	activity, err := db.ActivityByNameProjectID(ctx, d, name, projectID, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get activity: %w", ParseDBErrorDetail(err))
	}

	return activity, nil
}

func (a *Activity) ByProjectID(ctx context.Context, d db.DBTX, projectID db.ProjectID, opts ...db.ActivitySelectConfigOption) ([]db.Activity, error) {
	activities, err := db.ActivitiesByProjectID(ctx, d, projectID, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get activity: %w", ParseDBErrorDetail(err))
	}

	return activities, nil
}

func (a *Activity) ByID(ctx context.Context, d db.DBTX, id db.ActivityID, opts ...db.ActivitySelectConfigOption) (*db.Activity, error) {
	activity, err := db.ActivityByActivityID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get activity: %w", ParseDBErrorDetail(err))
	}

	return activity, nil
}

func (a *Activity) Delete(ctx context.Context, d db.DBTX, id db.ActivityID) (*db.Activity, error) {
	activity := &db.Activity{
		ActivityID: id,
	}

	err := activity.SoftDelete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete activity: %w", ParseDBErrorDetail(err))
	}

	return activity, err
}
