import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.30.2 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { NotificationType } from './notificationType';
import type { ModelsUserID } from './modelsUserID';

export interface ModelsNotification {
  body: string;
  createdAt: Date;
  labels: string[];
  /** @nullable */
  link?: string | null;
  notificationID: EntityIDs.NotificationID;
  notificationType: NotificationType;
  receiver?: ModelsUserID;
  sender: ModelsUserID;
  title: string;
}
