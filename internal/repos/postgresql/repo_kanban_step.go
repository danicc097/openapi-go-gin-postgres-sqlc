package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

// KanbanStep represents the repository used for interacting with KanbanStep records.
type KanbanStep struct {
	q models.Querier
}

// NewKanbanStep instantiates the KanbanStep repository.
func NewKanbanStep() *KanbanStep {
	return &KanbanStep{
		q: NewQuerierWrapper(models.New()),
	}
}

var _ repos.KanbanStep = (*KanbanStep)(nil)

func (k *KanbanStep) ByProject(ctx context.Context, d models.DBTX, projectID models.ProjectID, opts ...models.KanbanStepSelectConfigOption) ([]models.KanbanStep, error) {
	kss, err := models.KanbanStepsByProjectID(ctx, d, projectID, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get kanban steps: %w", ParseDBErrorDetail(err))
	}

	return kss, nil
}

func (k *KanbanStep) ByID(ctx context.Context, d models.DBTX, id models.KanbanStepID, opts ...models.KanbanStepSelectConfigOption) (*models.KanbanStep, error) {
	ks, err := models.KanbanStepByKanbanStepID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get kanban step: %w", ParseDBErrorDetail(err))
	}

	return ks, nil
}
