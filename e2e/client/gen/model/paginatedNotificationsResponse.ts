/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { NotificationResponse } from './notificationResponse'
import type { PaginationPage } from './paginationPage'

export interface PaginatedNotificationsResponse {
  items: NotificationResponse[] | null
  page: PaginationPage
}
