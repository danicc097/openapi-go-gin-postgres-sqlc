package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// KanbanStep represents the repository used for interacting with KanbanStep records.
type KanbanStep struct {
	q db.Querier
}

// NewKanbanStep instantiates the KanbanStep repository.
func NewKanbanStep() *KanbanStep {
	return &KanbanStep{
		q: NewQuerierWrapper(db.New()),
	}
}

var _ repos.KanbanStep = (*KanbanStep)(nil)

func (k *KanbanStep) ByProject(ctx context.Context, d db.DBTX, projectID db.ProjectID, opts ...db.KanbanStepSelectConfigOption) ([]db.KanbanStep, error) {
	kss, err := db.KanbanStepsByProjectID(ctx, d, projectID, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get kanban steps: %w", parseDBErrorDetail(err))
	}

	return kss, nil
}

func (k *KanbanStep) ByID(ctx context.Context, d db.DBTX, id db.KanbanStepID, opts ...db.KanbanStepSelectConfigOption) (*db.KanbanStep, error) {
	ks, err := db.KanbanStepByKanbanStepID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get kanban step: %w", parseDBErrorDetail(err))
	}

	return ks, nil
}
