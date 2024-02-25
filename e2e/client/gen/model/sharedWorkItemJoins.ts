/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { DbUserWIAWorkItem } from './dbUserWIAWorkItem'
import type { DbTimeEntry } from './dbTimeEntry'
import type { DbWorkItemComment } from './dbWorkItemComment'
import type { DbWorkItemTag } from './dbWorkItemTag'
import type { DbWorkItemType } from './dbWorkItemType'

export interface SharedWorkItemJoins {
  members?: DbUserWIAWorkItem[] | null
  timeEntries?: DbTimeEntry[] | null
  workItemComments?: DbWorkItemComment[] | null
  workItemTags?: DbWorkItemTag[] | null
  workItemType?: DbWorkItemType
}
