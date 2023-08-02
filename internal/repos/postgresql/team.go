package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// Team represents the repository used for interacting with Team records.
type Team struct {
	q *db.Queries
}

// NewTeam instantiates the Team repository.
func NewTeam() *Team {
	return &Team{
		q: db.New(),
	}
}

var _ repos.Team = (*Team)(nil)

func (t *Team) Create(ctx context.Context, d db.DBTX, params *db.TeamCreateParams) (*db.Team, error) {
	team, err := db.CreateTeam(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not create team: %w", parseErrorDetail(err))
	}

	return team, nil
}

func (t *Team) Update(ctx context.Context, d db.DBTX, id int, params *db.TeamUpdateParams) (*db.Team, error) {
	team, err := t.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get team by id %w", parseErrorDetail(err))
	}

	team.SetUpdateParams(params)

	team, err = team.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update team: %w", parseErrorDetail(err))
	}

	return team, err
}

func (t *Team) ByName(ctx context.Context, d db.DBTX, name string, projectID int, opts ...db.TeamSelectConfigOption) (*db.Team, error) {
	team, err := db.TeamByNameProjectID(ctx, d, name, projectID, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get team: %w", parseErrorDetail(err))
	}

	return team, nil
}

func (t *Team) ByID(ctx context.Context, d db.DBTX, id int, opts ...db.TeamSelectConfigOption) (*db.Team, error) {
	team, err := db.TeamByTeamID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get team: %w", parseErrorDetail(err))
	}

	return team, nil
}

func (t *Team) Delete(ctx context.Context, d db.DBTX, id int) (*db.Team, error) {
	team, err := t.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get team by id %w", parseErrorDetail(err))
	}

	err = team.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete team: %w", parseErrorDetail(err))
	}

	return team, err
}
