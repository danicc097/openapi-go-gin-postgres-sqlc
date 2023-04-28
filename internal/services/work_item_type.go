package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
)

type WorkItemTag struct {
	logger  *zap.Logger
	witRepo repos.WorkItemTag
}

// NewWorkItemTag returns a new WorkItemTag service.
func NewWorkItemTag(logger *zap.Logger, witRepo repos.WorkItemTag, notificationrepo repos.Notification, authzsvc *Authorization) *WorkItemTag {
	return &WorkItemTag{
		logger:  logger,
		witRepo: witRepo,
	}
}

// ByID gets a work item type by ID.
func (wit *WorkItemTag) ByID(ctx context.Context, d db.DBTX, id int) (*db.WorkItemTag, error) {
	defer newOTELSpan(ctx, "WorkItemTag.ByID").End()

	witObj, err := wit.witRepo.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("witRepo.ByID: %w", err)
	}

	return witObj, nil
}

// Create creates a new work item type.
func (wit *WorkItemTag) Create(ctx context.Context, d db.DBTX, params db.WorkItemTagCreateParams) (*db.WorkItemTag, error) {
	defer newOTELSpan(ctx, "WorkItemTag.Create").End()

	witObj, err := wit.witRepo.Create(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("witRepo.Create: %w", err)
	}

	return witObj, nil
}

// Update updates an existing work item type.
func (wit *WorkItemTag) Update(ctx context.Context, d db.DBTX, id int, params db.WorkItemTagUpdateParams) (*db.WorkItemTag, error) {
	defer newOTELSpan(ctx, "WorkItemTag.Update").End()

	witObj, err := wit.witRepo.Update(ctx, d, id, params)
	if err != nil {
		return nil, fmt.Errorf("witRepo.Update: %w", err)
	}

	return witObj, nil
}

// Delete deletes a work item type by ID.
func (wit *WorkItemTag) Delete(ctx context.Context, d db.DBTX, id int) (*db.WorkItemTag, error) {
	defer newOTELSpan(ctx, "WorkItemTag.Delete").End()

	witObj, err := wit.witRepo.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("witRepo.Delete: %w", err)
	}

	return witObj, nil
}
