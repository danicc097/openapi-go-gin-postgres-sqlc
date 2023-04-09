package models

import "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"

// ProjectBoard aggregates (for a given project id) kanban_steps, available work_item types, activities for time tracking...
// this is generic for all projects. Work items handled separately per project, since they use a dedicated table and models.
type ProjectBoard struct {
	Project *db.Project `json:"project"` // contains list of teams, work item types, kanban steps and activities if joined
}
