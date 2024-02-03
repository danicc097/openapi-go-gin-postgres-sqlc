import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { DbUserID } from './dbUserID'

export interface RestWorkItemComment {
  createdAt: Date
  message: string
  updatedAt: Date
  userID: EntityIDs.UserID
  workItemCommentID: EntityIDs.WorkItemCommentID
  workItemID: EntityIDs.WorkItemID
}
