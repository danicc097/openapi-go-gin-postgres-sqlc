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
	repos  *repos.Repos
	// sharedDBOpts represents shared db select options for all team entities
	// for returned values
	getSharedDBOpts func() []db.TeamSelectConfigOption
}

// NewTeam returns a new Team service.
func NewTeam(logger *zap.SugaredLogger, repos *repos.Repos) *Team {
	return &Team{
		logger: logger,
		repos:  repos,
		getSharedDBOpts: func() []db.TeamSelectConfigOption {
			return []db.TeamSelectConfigOption{db.WithTeamJoin(db.TeamJoins{Project: true})}
		},
	}
}

// ByID gets a team by ID.
func (t *Team) ByID(ctx context.Context, d db.DBTX, id db.TeamID) (*db.Team, error) {
	defer newOTelSpan().Build(ctx).End()

	team, err := t.repos.Team.ByID(ctx, d, id, t.getSharedDBOpts()...)
	if err != nil {
		return nil, fmt.Errorf("repos.Team.ByID: %w", err)
	}

	return team, nil
}

// ByName gets a team by name.
func (t *Team) ByName(ctx context.Context, d db.DBTX, name string, projectID db.ProjectID) (*db.Team, error) {
	defer newOTelSpan().Build(ctx).End()

	team, err := t.repos.Team.ByName(ctx, d, name, projectID, t.getSharedDBOpts()...)
	if err != nil {
		return nil, fmt.Errorf("repos.Team.ByName: %w", err)
	}

	return team, nil
}

// Create creates a new team.
func (t *Team) Create(ctx context.Context, d db.DBTX, params *db.TeamCreateParams) (*db.Team, error) {
	defer newOTelSpan().Build(ctx).End()

	team, err := t.repos.Team.Create(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("repos.Team.Create: %w", err)
	}

	team, err = t.repos.Team.ByID(ctx, d, team.TeamID, t.getSharedDBOpts()...)
	if err != nil {
		return nil, fmt.Errorf("repos.Team.ByID: %w", err)
	}

	return team, nil
}

// Update updates an existing team.
func (t *Team) Update(ctx context.Context, d db.DBTX, id db.TeamID, params *db.TeamUpdateParams) (*db.Team, error) {
	defer newOTelSpan().Build(ctx).End()

	team, err := t.repos.Team.Update(ctx, d, id, params)
	if err != nil {
		return nil, fmt.Errorf("repos.Team.Update: %w", err)
	}

	team, err = t.repos.Team.ByID(ctx, d, team.TeamID, t.getSharedDBOpts()...)
	if err != nil {
		return nil, fmt.Errorf("repos.Team.ByID: %w", err)
	}

	return team, nil
}

// Delete deletes an existing team.
func (t *Team) Delete(ctx context.Context, d db.DBTX, id db.TeamID) (*db.Team, error) {
	defer newOTelSpan().Build(ctx).End()

	team, err := t.repos.Team.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.Team.Delete: %w", err)
	}

	return team, nil
}
