package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

// Activity represents the repository used for interacting with Activity records.
type Activity struct {
	q models.Querier
}

// NewActivity instantiates the Activity repository.
func NewActivity() *Activity {
	return &Activity{
		q: NewQuerierWrapper(models.New()),
	}
}

var _ repos.Activity = (*Activity)(nil)

func (a *Activity) Create(ctx context.Context, d models.DBTX, params *models.ActivityCreateParams) (*models.Activity, error) {
	activity, err := models.CreateActivity(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not create activity: %w", ParseDBErrorDetail(err))
	}

	return activity, nil
}

func (a *Activity) Update(ctx context.Context, d models.DBTX, id models.ActivityID, params *models.ActivityUpdateParams) (*models.Activity, error) {
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

func (a *Activity) ByName(ctx context.Context, d models.DBTX, name string, projectID models.ProjectID, opts ...models.ActivitySelectConfigOption) (*models.Activity, error) {
	activity, err := models.ActivityByNameProjectID(ctx, d, name, projectID, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get activity: %w", ParseDBErrorDetail(err))
	}

	return activity, nil
}

func (a *Activity) ByProjectID(ctx context.Context, d models.DBTX, projectID models.ProjectID, opts ...models.ActivitySelectConfigOption) ([]models.Activity, error) {
	activities, err := models.ActivitiesByProjectID(ctx, d, projectID, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get activity: %w", ParseDBErrorDetail(err))
	}

	return activities, nil
}

func (a *Activity) ByID(ctx context.Context, d models.DBTX, id models.ActivityID, opts ...models.ActivitySelectConfigOption) (*models.Activity, error) {
	activity, err := models.ActivityByActivityID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get activity: %w", ParseDBErrorDetail(err))
	}

	return activity, nil
}

func (a *Activity) Delete(ctx context.Context, d models.DBTX, id models.ActivityID) (*models.Activity, error) {
	activity := &models.Activity{
		ActivityID: id,
	}

	err := activity.SoftDelete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete activity: %w", ParseDBErrorDetail(err))
	}

	return activity, err
}

func (a *Activity) Restore(ctx context.Context, d models.DBTX, id models.ActivityID) error {
	activity := &models.Activity{
		ActivityID: id,
	}

	_, err := activity.Restore(ctx, d)
	if err != nil {
		return fmt.Errorf("could not restore activity: %w", ParseDBErrorDetail(err))
	}

	return err
}
