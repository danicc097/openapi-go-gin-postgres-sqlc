import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.25.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { DbUserID } from './dbUserID';

export interface CreateWorkItemCommentRequest {
  message: string;
  userID: EntityIDs.UserID;
  workItemID: EntityIDs.WorkItemID;
}
