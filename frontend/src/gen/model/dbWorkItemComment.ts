/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { UuidUUID } from './uuidUUID'

export interface DbWorkItemComment {
  createdAt: Date
  message: string
  updatedAt: Date
  userID: UuidUUID
  workItemCommentID: number
  workItemID: number
}
