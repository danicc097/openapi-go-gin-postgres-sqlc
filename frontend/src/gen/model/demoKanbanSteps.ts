import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */

/**
 * is generated from kanban_steps table.
 */
export type DemoKanbanSteps = typeof DemoKanbanSteps[keyof typeof DemoKanbanSteps]

// eslint-disable-next-line @typescript-eslint/no-redeclare
export const DemoKanbanSteps = {
  Disabled: 'Disabled',
  Received: 'Received',
  Under_review: 'Under review',
  Work_in_progress: 'Work in progress',
} as const
