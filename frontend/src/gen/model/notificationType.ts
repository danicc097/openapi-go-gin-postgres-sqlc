/**
 * Generated by orval v6.19.1 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */

/**
 * is generated from database enum 'notification_type'.
 */
export type NotificationType = typeof NotificationType[keyof typeof NotificationType]

// eslint-disable-next-line @typescript-eslint/no-redeclare
export const NotificationType = {
  personal: 'personal',
  global: 'global',
} as const
