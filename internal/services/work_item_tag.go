package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
)

type WorkItemType struct {
	logger  *zap.Logger
	witRepo repos.WorkItemType
}

// NewWorkItemType returns a new WorkItemType service.
func NewWorkItemType(logger *zap.Logger, witRepo repos.WorkItemType) *WorkItemType {
	return &WorkItemType{
		logger:  logger,
		witRepo: witRepo,
	}
}

// ByID gets a work item tag by ID.
func (wit *WorkItemType) ByID(ctx context.Context, d db.DBTX, id int) (*db.WorkItemType, error) {
	defer newOTELSpan(ctx, "WorkItemType.ByID").End()

	witObj, err := wit.witRepo.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("witRepo.ByID: %w", err)
	}

	return witObj, nil
}

// ByName gets a work item tag by name.
func (wit *WorkItemType) ByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.WorkItemType, error) {
	defer newOTELSpan(ctx, "WorkItemType.ByName").End()

	witObj, err := wit.witRepo.ByName(ctx, d, name, projectID)
	if err != nil {
		return nil, fmt.Errorf("witRepo.ByName: %w", err)
	}

	return witObj, nil
}

// Create creates a new work item tag.
func (wit *WorkItemType) Create(ctx context.Context, d db.DBTX, params db.WorkItemTypeCreateParams) (*db.WorkItemType, error) {
	defer newOTELSpan(ctx, "WorkItemType.Create").End()

	witObj, err := wit.witRepo.Create(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("witRepo.Create: %w", err)
	}

	return witObj, nil
}

// Update updates an existing work item tag.
func (wit *WorkItemType) Update(ctx context.Context, d db.DBTX, id int, params db.WorkItemTypeUpdateParams) (*db.WorkItemType, error) {
	defer newOTELSpan(ctx, "WorkItemType.Update").End()

	witObj, err := wit.witRepo.Update(ctx, d, id, params)
	if err != nil {
		return nil, fmt.Errorf("witRepo.Update: %w", err)
	}

	return witObj, nil
}

// Delete deletes a work item tag by ID.
func (wit *WorkItemType) Delete(ctx context.Context, d db.DBTX, id int) (*db.WorkItemType, error) {
	defer newOTELSpan(ctx, "WorkItemType.Delete").End()

	witObj, err := wit.witRepo.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("witRepo.Delete: %w", err)
	}

	return witObj, nil
}
