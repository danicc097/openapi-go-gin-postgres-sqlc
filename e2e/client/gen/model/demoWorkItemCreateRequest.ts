/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { DbWorkItemCreateParams } from './dbWorkItemCreateParams'
import type { DbDemoWorkItemCreateParams } from './dbDemoWorkItemCreateParams'
import type { ServicesMember } from './servicesMember'
import type { DemoWorkItemCreateRequestProjectName } from './demoWorkItemCreateRequestProjectName'

export interface DemoWorkItemCreateRequest {
  base: DbWorkItemCreateParams
  demoProject: DbDemoWorkItemCreateParams
  members: ServicesMember[] | null
  projectName: DemoWorkItemCreateRequestProjectName
  tagIDs: number[] | null
}
