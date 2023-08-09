package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
)

type WorkItemType struct {
	logger  *zap.SugaredLogger
	witRepo repos.WorkItemType
}

// NewWorkItemType returns a new WorkItemType service.
func NewWorkItemType(logger *zap.SugaredLogger, witRepo repos.WorkItemType) *WorkItemType {
	return &WorkItemType{
		logger:  logger,
		witRepo: witRepo,
	}
}

// ByID gets a work item type by ID.
func (wit *WorkItemType) ByID(ctx context.Context, d db.DBTX, id int) (*db.WorkItemType, error) {
	defer newOTELSpan(ctx, "").End()

	witObj, err := wit.witRepo.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("witRepo.ByID: %w", err)
	}

	return witObj, nil
}
