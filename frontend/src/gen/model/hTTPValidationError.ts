import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { ValidationError } from './validationError';

export interface HTTPValidationError {
  /** Additional details for validation errors */
  detail?: ValidationError[];
  /** Descriptive error messages to show in a callout */
  messages: string[];
}
