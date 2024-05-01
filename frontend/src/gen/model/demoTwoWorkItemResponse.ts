import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.25.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { ModelsDemoTwoWorkItem } from './modelsDemoTwoWorkItem';
import type { ModelsWorkItemM2MAssigneeWIA } from './modelsWorkItemM2MAssigneeWIA';
import type { DemoTwoWorkItemResponseMetadata } from './demoTwoWorkItemResponseMetadata';
import type { DemoTwoWorkItemResponseProjectName } from './demoTwoWorkItemResponseProjectName';
import type { ModelsTimeEntry } from './modelsTimeEntry';
import type { ModelsWorkItemComment } from './modelsWorkItemComment';
import type { ModelsWorkItemTag } from './modelsWorkItemTag';
import type { ModelsWorkItemType } from './modelsWorkItemType';

export interface DemoTwoWorkItemResponse {
  closedAt?: Date | null;
  createdAt: Date;
  deletedAt?: Date | null;
  demoTwoWorkItem: ModelsDemoTwoWorkItem;
  description: string;
  kanbanStepID: EntityIDs.KanbanStepID;
  members?: ModelsWorkItemM2MAssigneeWIA[] | null;
  metadata: DemoTwoWorkItemResponseMetadata;
  projectName: DemoTwoWorkItemResponseProjectName;
  targetDate: Date;
  teamID: EntityIDs.TeamID | null;
  timeEntries?: ModelsTimeEntry[] | null;
  title: string;
  updatedAt: Date;
  workItemComments?: ModelsWorkItemComment[] | null;
  workItemID: EntityIDs.WorkItemID;
  workItemTags?: ModelsWorkItemTag[] | null;
  workItemType?: ModelsWorkItemType;
  workItemTypeID: EntityIDs.WorkItemTypeID;
}
