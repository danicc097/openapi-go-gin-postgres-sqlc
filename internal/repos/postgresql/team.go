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

func (u *Team) Create(ctx context.Context, d db.DBTX, params repos.TeamCreateParams) (*db.Team, error) {
	description := ""
	if params.Description != nil {
		description = *params.Description
	}

	team := &db.Team{
		Name:        params.Name,
		Description: description,
		ProjectID:   params.ProjectID,
	}

	if err := team.Save(ctx, d); err != nil {
		return nil, err
	}

	return team, nil
}

func (u *Team) Update(ctx context.Context, d db.DBTX, id int, params repos.TeamUpdateParams) (*db.Team, error) {
	team, err := u.TeamByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get team by id %w", parseErrorDetail(err))
	}

	if params.Description != nil {
		team.Description = *params.Description
	}
	if params.Name != nil {
		team.Name = *params.Name
	}

	err = team.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update team: %w", parseErrorDetail(err))
	}

	return team, err
}

func (u *Team) TeamByName(ctx context.Context, d db.DBTX, name string) (*db.Team, error) {
	team, err := db.TeamByName(ctx, d, name)
	if err != nil {
		return nil, fmt.Errorf("could not get team: %w", parseErrorDetail(err))
	}

	return team, nil
}

func (u *Team) TeamByID(ctx context.Context, d db.DBTX, id int) (*db.Team, error) {
	team, err := db.TeamByTeamID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get team: %w", parseErrorDetail(err))
	}

	return team, nil
}

func (u *Team) Delete(ctx context.Context, d db.DBTX, id int) (*db.Team, error) {
	team, err := u.TeamByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get team by id %w", parseErrorDetail(err))
	}

	err = team.Delete(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not delete team: %w", parseErrorDetail(err))
	}

	return team, err
}
