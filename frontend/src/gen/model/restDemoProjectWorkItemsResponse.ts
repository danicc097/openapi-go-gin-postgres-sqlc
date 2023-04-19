/**
 * Generated by orval v6.10.3 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { DbDemoProjectWorkItem } from './dbDemoProjectWorkItem'
import type { DbUser } from './dbUser'
import type { DbProject2WorkItem } from './dbProject2WorkItem'
import type { DbTimeEntry } from './dbTimeEntry'
import type { DbWorkItemComment } from './dbWorkItemComment'
import type { DbWorkItemTag } from './dbWorkItemTag'
import type { DbWorkItemType } from './dbWorkItemType'

export interface RestDemoProjectWorkItemsResponse {
  closed: Date | null
  createdAt: Date
  deletedAt: Date | null
  demoProjectWorkItem: DbDemoProjectWorkItem
  description: string
  kanbanStepID: number
  members?: DbUser[] | null
  metadata: number[] | null
  project2WorkItem?: DbProject2WorkItem
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
