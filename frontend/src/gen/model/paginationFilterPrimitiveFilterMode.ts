import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.30.2 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */

export type PaginationFilterPrimitiveFilterMode = typeof PaginationFilterPrimitiveFilterMode[keyof typeof PaginationFilterPrimitiveFilterMode];


// eslint-disable-next-line @typescript-eslint/no-redeclare
export const PaginationFilterPrimitiveFilterMode = {
  contains: 'contains',
  empty: 'empty',
  endsWith: 'endsWith',
  equals: 'equals',
  fuzzy: 'fuzzy',
  greaterThan: 'greaterThan',
  greaterThanOrEqualTo: 'greaterThanOrEqualTo',
  lessThan: 'lessThan',
  lessThanOrEqualTo: 'lessThanOrEqualTo',
  notEmpty: 'notEmpty',
  notEquals: 'notEquals',
  startsWith: 'startsWith',
} as const;
