/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { ModelsUserID } from './modelsUserID'

export interface WorkItemCommentResponse {
  createdAt: Date
  message: string
  updatedAt: Date
  userID: ModelsUserID
  workItemCommentID: number
  workItemID: number
}
