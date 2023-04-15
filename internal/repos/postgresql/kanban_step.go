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

func (k *KanbanStep) Create(ctx context.Context, d db.DBTX, params db.KanbanStepCreateParams) (*db.KanbanStep, error) {
	team := &db.KanbanStep{
		Name:          params.Name,
		Description:   params.Description,
		ProjectID:     params.ProjectID,
		Color:         params.Color,
		StepOrder:     params.StepOrder,
		TimeTrackable: params.TimeTrackable,
	}

	if _, err := team.Save(ctx, d); err != nil {
		return nil, err
	}

	return team, nil
}

func (k *KanbanStep) Update(ctx context.Context, d db.DBTX, id int, params db.KanbanStepUpdateParams) (*db.KanbanStep, error) {
	kanbanStep, err := k.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get kanban step by id %w", parseErrorDetail(err))
	}

	updateEntityWithParams(kanbanStep, &params)

	kanbanStep, err = kanbanStep.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update kanban step: %w", parseErrorDetail(err))
	}

	return kanbanStep, err
}

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

func (k *KanbanStep) Delete(ctx context.Context, d db.DBTX, id int) (*db.KanbanStep, error) {
	ks, err := k.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get kanban step by id %w", parseErrorDetail(err))
	}

	err = ks.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete kanban step: %w", parseErrorDetail(err))
	}

	return ks, err
}
