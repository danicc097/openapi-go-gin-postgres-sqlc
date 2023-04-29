package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
)

type Activity struct {
	logger *zap.Logger
	aRepo  repos.Activity
}

// NewActivity returns a new Activity service.
func NewActivity(logger *zap.Logger, aRepo repos.Activity) *Activity {
	return &Activity{
		logger: logger,
		aRepo:  aRepo,
	}
}

// ByID gets an activity by ID.
func (a *Activity) ByID(ctx context.Context, d db.DBTX, id int) (*db.Activity, error) {
	defer newOTELSpan(ctx, "Activity.ByID").End()

	activity, err := a.aRepo.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("aRepo.ByID: %w", err)
	}

	return activity, nil
}

// ByName gets an activity by name.
func (a *Activity) ByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.Activity, error) {
	defer newOTELSpan(ctx, "Activity.ByName").End()

	activity, err := a.aRepo.ByName(ctx, d, name, projectID)
	if err != nil {
		return nil, fmt.Errorf("aRepo.ByName: %w", err)
	}

	return activity, nil
}

// ByProjectID gets activities by project ID.
func (a *Activity) ByProjectID(ctx context.Context, d db.DBTX, projectID int) ([]db.Activity, error) {
	defer newOTELSpan(ctx, "Activity.ByProjectID").End()

	activity, err := a.aRepo.ByProjectID(ctx, d, projectID)
	if err != nil {
		return nil, fmt.Errorf("aRepo.ByProjectID: %w", err)
	}

	return activity, nil
}

// Create creates a new activity.
func (a *Activity) Create(ctx context.Context, d db.DBTX, params db.ActivityCreateParams) (*db.Activity, error) {
	defer newOTELSpan(ctx, "Activity.Create").End()

	activity, err := a.aRepo.Create(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("aRepo.Create: %w", err)
	}

	return activity, nil
}

// Update updates an existing activity.
func (a *Activity) Update(ctx context.Context, d db.DBTX, id int, params db.ActivityUpdateParams) (*db.Activity, error) {
	defer newOTELSpan(ctx, "Activity.Update").End()

	activity, err := a.aRepo.Update(ctx, d, id, params)
	if err != nil {
		return nil, fmt.Errorf("aRepo.Update: %w", err)
	}

	return activity, nil
}

// Delete deletes an activity by ID.
func (a *Activity) Delete(ctx context.Context, d db.DBTX, id int) (*db.Activity, error) {
	defer newOTELSpan(ctx, "Activity.Delete").End()

	activity, err := a.aRepo.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("aRepo.Delete: %w", err)
	}

	return activity, nil
}
