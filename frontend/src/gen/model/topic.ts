import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */

/**
 * string identifiers for SSE event listeners.
 */
export type Topic = typeof Topic[keyof typeof Topic];


// eslint-disable-next-line @typescript-eslint/no-redeclare
export const Topic = {
  WorkItemUpdated: 'WorkItemUpdated',
  TeamCreated: 'TeamCreated',
  GlobalAlerts: 'GlobalAlerts',
} as const;
