/**
 * Generated by orval v6.10.3 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { DbActivity } from './dbActivity'
import type { DbKanbanStep } from './dbKanbanStep'
import type { ModelsProject } from './modelsProject'
import type { DbTeam } from './dbTeam'
import type { DbWorkItemTag } from './dbWorkItemTag'
import type { DbWorkItemType } from './dbWorkItemType'

export interface DbProject {
  activities?: DbActivity[] | null
  createdAt: Date
  description: string
  initialized: boolean
  kanbanSteps?: DbKanbanStep[] | null
  name: ModelsProject
  projectID: number
  teams?: DbTeam[] | null
  updatedAt: Date
  workItemTags?: DbWorkItemTag[] | null
  workItemTypes?: DbWorkItemType[] | null
}
