/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { CacheDemoWorkItemResponse } from './cacheDemoWorkItemResponse'
import type { PaginationPage } from './paginationPage'

export interface PaginatedDemoWorkItemsResponse {
  items: CacheDemoWorkItemResponse[] | null
  page: PaginationPage
}
