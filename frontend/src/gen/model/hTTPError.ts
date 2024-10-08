import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.30.2 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { ErrorCode } from './errorCode';
import type { HTTPValidationError } from './hTTPValidationError';

/**
 * represents an error message response.
 */
export interface HTTPError {
  detail: string;
  error: string;
  /** location in body path, if any */
  loc?: string[];
  status: number;
  title: string;
  type: ErrorCode;
  validationError?: HTTPValidationError;
}
