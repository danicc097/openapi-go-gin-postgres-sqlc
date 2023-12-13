/**
 * Generated by orval v6.19.1 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { DbDemoWorkItem } from './dbDemoWorkItem'
import type { DbUserWIAUWorkItem } from './dbUserWIAUWorkItem'
import type { DemoWorkItemsMetadata } from './demoWorkItemsMetadata'
import type { DbTimeEntry } from './dbTimeEntry'
import type { DbWorkItemComment } from './dbWorkItemComment'
import type { DbWorkItemTag } from './dbWorkItemTag'
import type { DbWorkItemType } from './dbWorkItemType'

export interface DemoWorkItems {
  closedAt?: Date | null
  createdAt: Date
  deletedAt?: Date | null
  demoWorkItem: DbDemoWorkItem
  description: string
  kanbanStepID: number
  members?: DbUserWIAUWorkItem[] | null
  metadata: DemoWorkItemsMetadata
  targetDate: Date
  teamID: number
  timeEntries?: DbTimeEntry[] | null
  title: string
  updatedAt: Date
  workItemComments?: DbWorkItemComment[] | null
  workItemID: number
  workItemTags?: DbWorkItemTag[] | null
  workItemType?: DbWorkItemType
  workItemTypeID: number
}
