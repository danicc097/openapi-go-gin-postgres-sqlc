/**
 * Generated by orval v6.15.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { DbUserAPIKey } from './dbUserAPIKey'
import type { DbProject } from './dbProject'
import type { Role } from './role'
import type { Scopes } from './scopes'
import type { DbTeam } from './dbTeam'
import type { UuidUUID } from './uuidUUID'

export interface UserResponse {
  apiKey?: DbUserAPIKey
  createdAt: Date
  deletedAt: Date | null
  email: string
  firstName: string | null
  fullName: string | null
  hasGlobalNotifications: boolean
  hasPersonalNotifications: boolean
  lastName: string | null
  projects?: DbProject[] | null
  role: Role
  scopes: Scopes
  teams?: DbTeam[] | null
  userID: UuidUUID
  username: string
}
