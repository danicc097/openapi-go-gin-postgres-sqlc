import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.25.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { ModelsNotification } from './modelsNotification';
import type { ModelsUserID } from './modelsUserID';

export interface NotificationResponse {
  notification: ModelsNotification;
  notificationID: EntityIDs.NotificationID;
  read: boolean;
  userID: EntityIDs.UserID;
  userNotificationID: EntityIDs.UserNotificationID;
}
