package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
)

type KanbanStep struct {
	logger *zap.SugaredLogger
	ksrepo repos.KanbanStep
}

// NewKanbanStep returns a new KanbanStep service.
func NewKanbanStep(logger *zap.SugaredLogger, ksrepo repos.KanbanStep) *KanbanStep {
	return &KanbanStep{
		logger: logger,
		ksrepo: ksrepo,
	}
}

// ByID gets a KanbanStep by ID.
func (ks *KanbanStep) ByID(ctx context.Context, d db.DBTX, id int) (*db.KanbanStep, error) {
	defer newOTelSpan(ctx, "").End()

	kanbanStep, err := ks.ksrepo.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("ksrepo.ByID: %w", err)
	}

	return kanbanStep, nil
}

// ByProject gets all KanbanSteps for a project.
func (ks *KanbanStep) ByProject(ctx context.Context, d db.DBTX, projectID int) ([]db.KanbanStep, error) {
	defer newOTelSpan(ctx, "").End()

	kanbanSteps, err := ks.ksrepo.ByProject(ctx, d, projectID)
	if err != nil {
		return nil, fmt.Errorf("ksrepo.ByProject: %w", err)
	}

	return kanbanSteps, nil
}
