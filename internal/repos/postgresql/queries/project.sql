-- name: IsTeamInProject :one
select
  exists (
    select
      1
    from
      teams
      join projects using (project_id)
    where
      teams.team_id = @team_id
      and projects.project_id = @project_id);
