/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { ModelsProjectConfigField } from './modelsProjectConfigField'
import type { ModelsProjectConfigVisualization } from './modelsProjectConfigVisualization'

export interface ModelsProjectConfig {
  fields?: ModelsProjectConfigField[] | null
  header?: string[] | null
  visualization?: ModelsProjectConfigVisualization
}
