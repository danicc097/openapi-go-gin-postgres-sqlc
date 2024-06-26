// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: project.sql

package models

import (
	"context"
)

const IsTeamInProject = `-- name: IsTeamInProject :one
select
  exists (
    select
      1
    from
      teams
      join projects using (project_id)
    where
      teams.team_id = $1
      and projects.project_id = $2)
`

type IsTeamInProjectParams struct {
	TeamID    int32 `db:"team_id" json:"team_id"`
	ProjectID int32 `db:"project_id" json:"project_id"`
}

func (q *Queries) IsTeamInProject(ctx context.Context, db DBTX, arg IsTeamInProjectParams) (bool, error) {
	row := db.QueryRow(ctx, IsTeamInProject, arg.TeamID, arg.ProjectID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
