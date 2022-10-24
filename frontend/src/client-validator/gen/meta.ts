/* eslint-disable */
// @ts-nocheck
/* eslint-disable */
/* eslint-disable @typescript-eslint/ban-ts-comment */

import { HTTPValidationError, UpdateUserRequest, Scope, Role, GetCurrentUserResponse, ValidationError } from './models'

export const schemaDefinitions = {
  HTTPValidationError: info<HTTPValidationError>('HTTPValidationError', '#/definitions/HTTPValidationError'),
  UpdateUserRequest: info<UpdateUserRequest>('UpdateUserRequest', '#/definitions/UpdateUserRequest'),
  Scope: info<Scope>('Scope', '#/definitions/Scope'),
  Role: info<Role>('Role', '#/definitions/Role'),
  GetCurrentUserResponse: info<GetCurrentUserResponse>(
    'GetCurrentUserResponse',
    '#/definitions/GetCurrentUserResponse',
  ),
  ValidationError: info<ValidationError>('ValidationError', '#/definitions/ValidationError'),
}

export interface SchemaInfo<T> {
  definitionName: string
  schemaRef: string
}

function info<T>(definitionName: string, schemaRef: string): SchemaInfo<T> {
  return { definitionName, schemaRef }
}
