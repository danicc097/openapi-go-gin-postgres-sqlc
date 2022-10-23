/* eslint-disable */
// @ts-nocheck
/* eslint-disable */
/* eslint-disable @typescript-eslint/ban-ts-comment */

import { HTTPValidationError, User, ValidationError } from './models'

export const schemaDefinitions = {
  HTTPValidationError: info<HTTPValidationError>('HTTPValidationError', '#/definitions/HTTPValidationError'),
  User: info<User>('User', '#/definitions/User'),
  ValidationError: info<ValidationError>('ValidationError', '#/definitions/ValidationError'),
}

export interface SchemaInfo<T> {
  definitionName: string
  schemaRef: string
}

function info<T>(definitionName: string, schemaRef: string): SchemaInfo<T> {
  return { definitionName, schemaRef }
}
