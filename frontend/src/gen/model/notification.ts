import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { DbNotification } from './dbNotification';
import type { DbUserID } from './dbUserID';

export interface Notification {
  notification: DbNotification;
  notificationID: EntityIDs.NotificationID;
  read: boolean;
  userID: EntityIDs.UserID;
  userNotificationID: EntityIDs.UserNotificationID;
}
