/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { ModelsUserAPIKey } from './modelsUserAPIKey'
import type { ModelsProject } from './modelsProject'
import type { Role } from './role'
import type { Scopes } from './scopes'
import type { ModelsTeam } from './modelsTeam'
import type { ModelsUserID } from './modelsUserID'

export interface UserResponse {
  age?: number | null
  apiKey?: ModelsUserAPIKey
  createdAt: Date
  deletedAt?: Date | null
  email: string
  firstName?: string | null
  fullName?: string | null
  hasGlobalNotifications: boolean
  hasPersonalNotifications: boolean
  lastName?: string | null
  projects?: ModelsProject[] | null
  role: Role
  scopes: Scopes
  teams?: ModelsTeam[] | null
  updatedAt: Date
  userID: ModelsUserID
  username: string
}
