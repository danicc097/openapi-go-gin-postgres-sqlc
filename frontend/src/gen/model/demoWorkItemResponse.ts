import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.30.2 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { ModelsDemoWorkItem } from './modelsDemoWorkItem';
import type { ModelsWorkItemM2MAssigneeWIA } from './modelsWorkItemM2MAssigneeWIA';
import type { DemoWorkItemResponseMetadata } from './demoWorkItemResponseMetadata';
import type { DemoWorkItemResponseProjectName } from './demoWorkItemResponseProjectName';
import type { ModelsTimeEntry } from './modelsTimeEntry';
import type { ModelsWorkItemComment } from './modelsWorkItemComment';
import type { ModelsWorkItemTag } from './modelsWorkItemTag';
import type { ModelsWorkItemType } from './modelsWorkItemType';

export interface DemoWorkItemResponse {
  /** @nullable */
  closedAt?: Date | null;
  createdAt: Date;
  /** @nullable */
  deletedAt?: Date | null;
  demoWorkItem: ModelsDemoWorkItem;
  description: string;
  kanbanStepID: EntityIDs.KanbanStepID;
  /** @nullable */
  members?: ModelsWorkItemM2MAssigneeWIA[] | null;
  metadata: DemoWorkItemResponseMetadata;
  projectName: DemoWorkItemResponseProjectName;
  targetDate: Date;
  /** @nullable */
  teamID: EntityIDs.TeamID | null;
  /** @nullable */
  timeEntries?: ModelsTimeEntry[] | null;
  title: string;
  updatedAt: Date;
  /** @nullable */
  workItemComments?: ModelsWorkItemComment[] | null;
  workItemID: EntityIDs.WorkItemID;
  /** @nullable */
  workItemTags?: ModelsWorkItemTag[] | null;
  workItemType?: ModelsWorkItemType;
  workItemTypeID: EntityIDs.WorkItemTypeID;
}
