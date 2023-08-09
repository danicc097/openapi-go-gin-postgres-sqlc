package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
)

type Team struct {
	logger *zap.SugaredLogger
	trepo  repos.Team
}

// NewTeam returns a new Team service.
func NewTeam(logger *zap.SugaredLogger, trepo repos.Team) *Team {
	return &Team{
		logger: logger,
		trepo:  trepo,
	}
}

// ByID gets a team by ID.
func (t *Team) ByID(ctx context.Context, d db.DBTX, id int) (*db.Team, error) {
	defer newOTELSpan(ctx, "").End()

	team, err := t.trepo.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("trepo.ByID: %w", err)
	}

	return team, nil
}

// ByName gets a team by name.
func (t *Team) ByName(ctx context.Context, d db.DBTX, name string, projectID int) (*db.Team, error) {
	defer newOTELSpan(ctx, "").End()

	team, err := t.trepo.ByName(ctx, d, name, projectID)
	if err != nil {
		return nil, fmt.Errorf("trepo.ByName: %w", err)
	}

	return team, nil
}

// Create creates a new team.
func (t *Team) Create(ctx context.Context, d db.DBTX, params *db.TeamCreateParams) (*db.Team, error) {
	defer newOTELSpan(ctx, "").End()

	team, err := t.trepo.Create(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("trepo.Create: %w", err)
	}

	return team, nil
}

// Update updates an existing team.
func (t *Team) Update(ctx context.Context, d db.DBTX, id int, params *db.TeamUpdateParams) (*db.Team, error) {
	defer newOTELSpan(ctx, "").End()

	team, err := t.trepo.Update(ctx, d, id, params)
	if err != nil {
		return nil, fmt.Errorf("trepo.Update: %w", err)
	}

	return team, nil
}

// Delete deletes an existing team.
func (t *Team) Delete(ctx context.Context, d db.DBTX, id int) (*db.Team, error) {
	defer newOTELSpan(ctx, "").End()

	team, err := t.trepo.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("trepo.Delete: %w", err)
	}

	return team, nil
}
