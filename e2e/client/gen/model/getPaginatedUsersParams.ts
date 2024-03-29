/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { Direction } from './direction'
import type { GetPaginatedUsersFilterObjectsItem } from './getPaginatedUsersFilterObjectsItem'
import type { GetPaginatedUsersNestedObj } from './getPaginatedUsersNestedObj'

export type GetPaginatedUsersParams = {
  limit: number
  direction: Direction
  cursor: string
  filter?: {
    bools?: boolean[]
    ints?: number[]
    objects?: GetPaginatedUsersFilterObjectsItem[]
    post?: string[]
  }
  nested?: {
    obj?: GetPaginatedUsersNestedObj
  }
}
