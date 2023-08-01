package models

import "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"

// ProjectBoard aggregates (for a given project id) kanban_steps, available work_item types, activities for time tracking...
// this is generic for all projects. Work items handled separately per project, since they use a dedicated table and models.
type ProjectBoard struct {
	db.Project // contains list of teams, work item types, kanban steps and activities if joined

	Activities    []db.Activity     `json:"activities"    required:"true"`
	KanbanSteps   []db.KanbanStep   `json:"kanbanSteps"   required:"true"`
	Teams         []db.Team         `json:"teams"         required:"true"`
	WorkItemTags  []db.WorkItemTag  `json:"workItemTags"  required:"true"`
	WorkItemTypes []db.WorkItemType `json:"workItemTypes" required:"true"`
}
