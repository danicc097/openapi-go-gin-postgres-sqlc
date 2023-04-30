package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// KanbanStep represents the repository used for interacting with KanbanStep records.
type KanbanStep struct {
	q *db.Queries
}

// NewKanbanStep instantiates the KanbanStep repository.
func NewKanbanStep() *KanbanStep {
	return &KanbanStep{
		q: db.New(),
	}
}

var _ repos.KanbanStep = (*KanbanStep)(nil)

func (k *KanbanStep) ByProject(ctx context.Context, d db.DBTX, projectID int) ([]db.KanbanStep, error) {
	kss, err := db.KanbanStepsByProjectID(ctx, d, projectID)
	if err != nil {
		return nil, fmt.Errorf("could not get kanban steps: %w", parseErrorDetail(err))
	}

	return kss, nil
}

func (k *KanbanStep) ByID(ctx context.Context, d db.DBTX, id int) (*db.KanbanStep, error) {
	ks, err := db.KanbanStepByKanbanStepID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get kanban step: %w", parseErrorDetail(err))
	}

	return ks, nil
}
