/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { ModelsDemoWorkItem } from './modelsDemoWorkItem'
import type { ModelsWorkItemM2MAssigneeWIA } from './modelsWorkItemM2MAssigneeWIA'
import type { DemoWorkItemResponseMetadata } from './demoWorkItemResponseMetadata'
import type { DemoWorkItemResponseProjectName } from './demoWorkItemResponseProjectName'
import type { ModelsTimeEntry } from './modelsTimeEntry'
import type { ModelsWorkItemComment } from './modelsWorkItemComment'
import type { ModelsWorkItemTag } from './modelsWorkItemTag'
import type { ModelsWorkItemType } from './modelsWorkItemType'

export interface DemoWorkItemResponse {
  closedAt?: Date | null
  createdAt: Date
  deletedAt?: Date | null
  demoWorkItem: ModelsDemoWorkItem
  description: string
  kanbanStepID: number
  members?: ModelsWorkItemM2MAssigneeWIA[] | null
  metadata: DemoWorkItemResponseMetadata
  projectName: DemoWorkItemResponseProjectName
  targetDate: Date
  teamID: number | null
  timeEntries?: ModelsTimeEntry[] | null
  title: string
  updatedAt: Date
  workItemComments?: ModelsWorkItemComment[] | null
  workItemID: number
  workItemTags?: ModelsWorkItemTag[] | null
  workItemType?: ModelsWorkItemType
  workItemTypeID: number
}
