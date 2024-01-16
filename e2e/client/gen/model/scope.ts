/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */

/**
 * is generated from scopes.json keys.
 */
export type Scope = typeof Scope[keyof typeof Scope]

// eslint-disable-next-line @typescript-eslint/no-redeclare
export const Scope = {
  'project-member': 'project-member',
  'users:read': 'users:read',
  'users:write': 'users:write',
  'users:delete': 'users:delete',
  'scopes:write': 'scopes:write',
  'team-settings:write': 'team-settings:write',
  'project-settings:write': 'project-settings:write',
  'activity:create': 'activity:create',
  'activity:edit': 'activity:edit',
  'activity:delete': 'activity:delete',
  'work-item-tag:create': 'work-item-tag:create',
  'work-item-tag:edit': 'work-item-tag:edit',
  'work-item-tag:delete': 'work-item-tag:delete',
  'work-item:review': 'work-item:review',
  'work-item-comment:create': 'work-item-comment:create',
  'work-item-comment:edit': 'work-item-comment:edit',
  'work-item-comment:delete': 'work-item-comment:delete',
} as const
