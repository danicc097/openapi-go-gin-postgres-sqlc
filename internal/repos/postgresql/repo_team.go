package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

// Team represents the repository used for interacting with Team records.
type Team struct {
	q models.Querier
}

// NewTeam instantiates the Team repository.
func NewTeam() *Team {
	return &Team{
		q: NewQuerierWrapper(models.New()),
	}
}

var _ repos.Team = (*Team)(nil)

func (t *Team) Create(ctx context.Context, d models.DBTX, params *models.TeamCreateParams) (*models.Team, error) {
	team, err := models.CreateTeam(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not create team: %w", ParseDBErrorDetail(err))
	}

	return team, nil
}

func (t *Team) Update(ctx context.Context, d models.DBTX, id models.TeamID, params *models.TeamUpdateParams) (*models.Team, error) {
	team, err := t.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get team by id %w", ParseDBErrorDetail(err))
	}

	team.SetUpdateParams(params)

	team, err = team.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update team: %w", ParseDBErrorDetail(err))
	}

	return team, err
}

func (t *Team) ByName(ctx context.Context, d models.DBTX, name string, projectID models.ProjectID, opts ...models.TeamSelectConfigOption) (*models.Team, error) {
	team, err := models.TeamByNameProjectID(ctx, d, name, projectID, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get team: %w", ParseDBErrorDetail(err))
	}

	return team, nil
}

func (t *Team) ByID(ctx context.Context, d models.DBTX, id models.TeamID, opts ...models.TeamSelectConfigOption) (*models.Team, error) {
	team, err := models.TeamByTeamID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get team: %w", ParseDBErrorDetail(err))
	}

	return team, nil
}

func (t *Team) Delete(ctx context.Context, d models.DBTX, id models.TeamID) (*models.Team, error) {
	team := &models.Team{
		TeamID: id,
	}

	err := team.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete team: %w", ParseDBErrorDetail(err))
	}

	return team, err
}
