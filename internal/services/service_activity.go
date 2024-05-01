package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"go.uber.org/zap"
)

type Activity struct {
	logger *zap.SugaredLogger
	repos  *repos.Repos
}

// NewActivity returns a new Activity service.
func NewActivity(logger *zap.SugaredLogger, repos *repos.Repos) *Activity {
	return &Activity{
		logger: logger,
		repos:  repos,
	}
}

// ByID gets an activity by ID.
func (a *Activity) ByID(ctx context.Context, d models.DBTX, id models.ActivityID) (*models.Activity, error) {
	defer newOTelSpan().Build(ctx).End()

	activity, err := a.repos.Activity.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.Activity.ByID: %w", err)
	}

	return activity, nil
}

// ByName gets an activity by name.
func (a *Activity) ByName(ctx context.Context, d models.DBTX, name string, projectID models.ProjectID) (*models.Activity, error) {
	defer newOTelSpan().Build(ctx).End()

	activity, err := a.repos.Activity.ByName(ctx, d, name, projectID)
	if err != nil {
		return nil, fmt.Errorf("repos.Activity.ByName: %w", err)
	}

	return activity, nil
}

// ByProjectID gets activities by project ID.
func (a *Activity) ByProjectID(ctx context.Context, d models.DBTX, projectID models.ProjectID) ([]models.Activity, error) {
	defer newOTelSpan().Build(ctx).End()

	activity, err := a.repos.Activity.ByProjectID(ctx, d, projectID)
	if err != nil {
		return nil, fmt.Errorf("repos.Activity.ByProjectID: %w", err)
	}

	return activity, nil
}

// Create creates a new activity.
func (a *Activity) Create(ctx context.Context, d models.DBTX, projectName models.ProjectName, params *models.ActivityCreateParams) (*models.Activity, error) {
	defer newOTelSpan().Build(ctx).End()

	params.ProjectID = internal.ProjectIDByName[projectName]

	activity, err := a.repos.Activity.Create(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("repos.Activity.Create: %w", err)
	}

	return activity, nil
}

// Update updates an existing activity.
func (a *Activity) Update(ctx context.Context, d models.DBTX, id models.ActivityID, params *models.ActivityUpdateParams) (*models.Activity, error) {
	defer newOTelSpan().Build(ctx).End()

	activity, err := a.repos.Activity.Update(ctx, d, id, params)
	if err != nil {
		return nil, fmt.Errorf("repos.Activity.Update: %w", err)
	}

	return activity, nil
}

// Delete deletes an activity by ID.
func (a *Activity) Delete(ctx context.Context, d models.DBTX, id models.ActivityID) (*models.Activity, error) {
	defer newOTelSpan().Build(ctx).End()

	activity, err := a.repos.Activity.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.Activity.Delete: %w", err)
	}

	return activity, nil
}

func (a *Activity) Restore(ctx context.Context, d models.DBTX, id models.ActivityID) error {
	defer newOTelSpan().Build(ctx).End()

	if err := a.repos.Activity.Restore(ctx, d, id); err != nil {
		return fmt.Errorf("repos.Activity.Restore: %w", err)
	}

	return nil
}
