/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { ProjectConfig } from './projectConfig'
import type { ProjectName } from './projectName'

export interface DbProject {
  boardConfig: ProjectConfig
  createdAt: Date
  description: string
  name: ProjectName
  projectID: number
  updatedAt: Date
}
