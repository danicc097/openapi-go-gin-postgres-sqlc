import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.30.2 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { ModelsWorkItemM2MAssigneeWIA } from './modelsWorkItemM2MAssigneeWIA';
import type { ModelsTimeEntry } from './modelsTimeEntry';
import type { ModelsWorkItemComment } from './modelsWorkItemComment';
import type { ModelsWorkItemTag } from './modelsWorkItemTag';
import type { ModelsWorkItemType } from './modelsWorkItemType';

export interface SharedWorkItemJoins {
  /** @nullable */
  members?: ModelsWorkItemM2MAssigneeWIA[] | null;
  /** @nullable */
  timeEntries?: ModelsTimeEntry[] | null;
  /** @nullable */
  workItemComments?: ModelsWorkItemComment[] | null;
  /** @nullable */
  workItemTags?: ModelsWorkItemTag[] | null;
  workItemType?: ModelsWorkItemType;
}
