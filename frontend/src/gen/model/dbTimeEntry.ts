/**
 * Generated by orval v6.15.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { UuidUUID } from './uuidUUID'

export interface DbTimeEntry {
  activityID: number
  comment: string
  durationMinutes: number | null
  start: Date
  teamID: number | null
  timeEntryID: number
  userID: UuidUUID
  workItemID: number | null
}
