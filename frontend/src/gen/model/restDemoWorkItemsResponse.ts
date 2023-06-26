/**
 * Generated by orval v6.15.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { DbDemoWorkItem } from './dbDemoWorkItem'
import type { DbUser } from './dbUser'
import type { RestDemoWorkItemsResponseMetadata } from './restDemoWorkItemsResponseMetadata'
import type { DbTimeEntry } from './dbTimeEntry'
import type { DbWorkItemComment } from './dbWorkItemComment'
import type { DbWorkItemTag } from './dbWorkItemTag'
import type { DbWorkItemType } from './dbWorkItemType'

export interface RestDemoWorkItemsResponse {
  closed: Date | null
  createdAt: Date
  deletedAt: Date | null
  demoWorkItem: DbDemoWorkItem
  description: string
  kanbanStepID: number
  members?: DbUser[] | null
  metadata: RestDemoWorkItemsResponseMetadata
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
