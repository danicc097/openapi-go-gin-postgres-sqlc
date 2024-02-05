import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */

export interface DbKanbanStep {
  color: string;
  description: string;
  kanbanStepID: EntityIDs.KanbanStepID;
  name: string;
  projectID: EntityIDs.ProjectID;
  stepOrder: number;
  timeTrackable: boolean;
}
