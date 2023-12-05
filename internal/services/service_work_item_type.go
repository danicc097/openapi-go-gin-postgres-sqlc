package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
)

type WorkItemType struct {
	logger *zap.SugaredLogger
	repos  *repos.Repos
}

// NewWorkItemType returns a new WorkItemType service.
func NewWorkItemType(logger *zap.SugaredLogger, repos *repos.Repos) *WorkItemType {
	return &WorkItemType{
		logger: logger,
		repos:  repos,
	}
}

// ByID gets a work item type by ID.
func (wit *WorkItemType) ByID(ctx context.Context, d db.DBTX, id db.WorkItemTypeID) (*db.WorkItemType, error) {
	defer newOTelSpan().Build(ctx).End()

	witObj, err := wit.repos.WorkItemType.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.WorkItemType.ByID: %w", err)
	}

	return witObj, nil
}
