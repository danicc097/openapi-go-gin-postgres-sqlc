/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { RestNotification } from './restNotification'
import type { RestPaginationPage } from './restPaginationPage'

export interface PaginatedNotificationsResponse {
  items?: RestNotification[] | null
  page?: RestPaginationPage
}
