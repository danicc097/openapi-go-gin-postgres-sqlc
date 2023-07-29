/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */

/**
 * Represents standardized HTTP error types.
Notes:
- 'Private' marks an error to be hidden in response.

 */
export type ErrorCode = typeof ErrorCode[keyof typeof ErrorCode]

// eslint-disable-next-line @typescript-eslint/no-redeclare
export const ErrorCode = {
  Unknown: 'Unknown',
  Private: 'Private',
  NotFound: 'NotFound',
  InvalidArgument: 'InvalidArgument',
  AlreadyExists: 'AlreadyExists',
  Unauthorized: 'Unauthorized',
  Unauthenticated: 'Unauthenticated',
  RequestValidation: 'RequestValidation',
  ResponseValidation: 'ResponseValidation',
  OIDC: 'OIDC',
  InvalidRole: 'InvalidRole',
  InvalidScope: 'InvalidScope',
  InvalidUUID: 'InvalidUUID',
} as const
