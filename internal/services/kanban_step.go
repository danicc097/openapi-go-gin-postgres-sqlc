package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
)

type KanbanStep struct {
	logger *zap.Logger
	ksrepo repos.KanbanStep
}

// NewKanbanStep returns a new KanbanStep service.
func NewKanbanStep(logger *zap.Logger, ksrepo repos.KanbanStep, notificationrepo repos.Notification, authzsvc *Authorization) *KanbanStep {
	return &KanbanStep{
		logger: logger,
		ksrepo: ksrepo,
	}
}

// ByID gets a KanbanStep by ID.
func (ks *KanbanStep) ByID(ctx context.Context, d db.DBTX, id int) (*db.KanbanStep, error) {
	defer newOTELSpan(ctx, "KanbanStep.ByID").End()

	kanbanStep, err := ks.ksrepo.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("ksrepo.ByID: %w", err)
	}

	return kanbanStep, nil
}

// ByProject gets all KanbanSteps for a project.
func (ks *KanbanStep) ByProject(ctx context.Context, d db.DBTX, projectID int) ([]db.KanbanStep, error) {
	defer newOTELSpan(ctx, "KanbanStep.ByProject").End()

	kanbanSteps, err := ks.ksrepo.ByProject(ctx, d, projectID)
	if err != nil {
		return nil, fmt.Errorf("ksrepo.ByProject: %w", err)
	}

	return kanbanSteps, nil
}

// Create creates a new KanbanStep.
func (ks *KanbanStep) Create(ctx context.Context, d db.DBTX, params db.KanbanStepCreateParams) (*db.KanbanStep, error) {
	defer newOTELSpan(ctx, "KanbanStep.Create").End()

	kanbanStep, err := ks.ksrepo.Create(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("ksrepo.Create: %w", err)
	}

	return kanbanStep, nil
}

// Update updates an existing KanbanStep.
func (ks *KanbanStep) Update(ctx context.Context, d db.DBTX, id int, params db.KanbanStepUpdateParams) (*db.KanbanStep, error) {
	defer newOTELSpan(ctx, "KanbanStep.Update").End()

	kanbanStep, err := ks.ksrepo.Update(ctx, d, id, params)
	if err != nil {
		return nil, fmt.Errorf("ksrepo.Update: %w", err)
	}

	return kanbanStep, nil
}

// Delete deletes a KanbanStep.
func (ks *KanbanStep) Delete(ctx context.Context, d db.DBTX, id int) (*db.KanbanStep, error) {
	defer newOTELSpan(ctx, "KanbanStep.Delete").End()

	kanbanStep, err := ks.ksrepo.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("ksrepo.Delete: %w", err)
	}

	return kanbanStep, nil
}
