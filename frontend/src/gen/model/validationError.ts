/**
 * Generated by orval v6.15.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { ValidationErrorDetail } from './validationErrorDetail'
import type { ValidationErrorCtx } from './validationErrorCtx'

export interface ValidationError {
  /** location in body path, if any */
  loc: string[]
  /** should always be shown to the user */
  msg: string
  /** verbose details of the error */
  detail: ValidationErrorDetail
  ctx?: ValidationErrorCtx
}
