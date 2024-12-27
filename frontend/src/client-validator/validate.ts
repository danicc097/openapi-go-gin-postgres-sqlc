/* eslint-disable */
import type { ErrorObject } from 'ajv'

export interface Validator {
  (json: unknown): boolean
  errors?: ErrorObject[] | null
}

export interface ValidationErrors {
  /* Location of the error */
  path: string
  /* Generic error message */
  message?: string
  errors?: ValidationError[]
}

interface ValidationError {
  invalidParams: {
    name: string
    reason: string
  }
}

export function validateJson(json: any, validator: Validator, definitionName: string): any {
  let validationErrors: ValidationErrors = {
    path: definitionName,
    message: 'Unexpected data received',
  }
  const jsonObject = typeof json === 'string' ? JSON.parse(json) : json

  if (validator(jsonObject)) {
    return jsonObject
  }

  if (validator.errors) {
    validationErrors.message = 'Validation errors found.'
    validationErrors.errors = parseErrors(validator.errors)
  }

  throw {
    validationErrors,
    error: new Error(),
  }
}

function parseErrors(errors: ErrorObject[]): ValidationError[] {
  let out: ValidationError[] = []
  errors.forEach(
    (error, i) =>
      (out[i] = {
        invalidParams: {
          name: error.instancePath.split('/').slice(1).join('.'),
          reason: error.message || 'Unknown error',
        },
      }),
  )
  return out
}
