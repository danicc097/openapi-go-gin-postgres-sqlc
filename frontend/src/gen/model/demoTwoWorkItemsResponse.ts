/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { DbDemoTwoWorkItem } from './dbDemoTwoWorkItem'
import type { DbUser } from './dbUser'
import type { DemoTwoWorkItemsResponseMetadata } from './demoTwoWorkItemsResponseMetadata'
import type { DbTimeEntry } from './dbTimeEntry'
import type { DbWorkItemComment } from './dbWorkItemComment'
import type { DbWorkItemID } from './dbWorkItemID'
import type { DbWorkItemTag } from './dbWorkItemTag'
import type { DbWorkItemType } from './dbWorkItemType'

export interface DemoTwoWorkItemsResponse {
  closedAt?: Date | null
  createdAt: Date
  deletedAt?: Date | null
  demoTwoWorkItem: DbDemoTwoWorkItem
  description: string
  kanbanStepID: number
  members?: DbUser[] | null
  metadata: DemoTwoWorkItemsResponseMetadata
  targetDate: Date
  teamID: number
  timeEntries?: DbTimeEntry[] | null
  title: string
  updatedAt: Date
  workItemComments?: DbWorkItemComment[] | null
  workItemID: DbWorkItemID
  workItemTags?: DbWorkItemTag[] | null
  workItemType?: DbWorkItemType
  workItemTypeID: number
}
