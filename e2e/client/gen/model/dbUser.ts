/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { Scopes } from './scopes'
import type { UuidUUID } from './uuidUUID'

export interface DbUser {
  createdAt: Date
  deletedAt?: Date | null
  email: string
  firstName?: string | null
  fullName?: string | null
  hasGlobalNotifications: boolean
  hasPersonalNotifications: boolean
  lastName?: string | null
  scopes: Scopes
  userID: UuidUUID
  username: string
}