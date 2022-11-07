/* eslint-disable */
// @ts-nocheck
/* eslint-disable */
/* eslint-disable @typescript-eslint/ban-ts-comment */

import Ajv from 'ajv'
import addFormats from 'ajv-formats'
import { Decoder } from './helpers'
import { validateJson } from '../validate'
import {
  HTTPValidationError,
  UpdateUserRequest,
  Scope,
  Role,
  TaskRole,
  Organization,
  GetCurrentUserRes,
  ValidationError,
} from './models'
import jsonSchema from './schema.json'

const ajv = new Ajv({ strict: false, allErrors: true })
addFormats(ajv, { formats: ['int64', 'int32', 'binary', 'date-time'] })
ajv.compile(jsonSchema)

// Decoders
export const HTTPValidationErrorDecoder: Decoder<HTTPValidationError> = {
  definitionName: 'HTTPValidationError',
  schemaRef: '#/definitions/HTTPValidationError',

  decode(json: unknown): HTTPValidationError {
    const schema = ajv.getSchema(HTTPValidationErrorDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${HTTPValidationErrorDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, HTTPValidationErrorDecoder.definitionName)
  },
}
export const UpdateUserRequestDecoder: Decoder<UpdateUserRequest> = {
  definitionName: 'UpdateUserRequest',
  schemaRef: '#/definitions/UpdateUserRequest',

  decode(json: unknown): UpdateUserRequest {
    const schema = ajv.getSchema(UpdateUserRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${UpdateUserRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, UpdateUserRequestDecoder.definitionName)
  },
}
export const ScopeDecoder: Decoder<Scope> = {
  definitionName: 'Scope',
  schemaRef: '#/definitions/Scope',

  decode(json: unknown): Scope {
    const schema = ajv.getSchema(ScopeDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ScopeDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ScopeDecoder.definitionName)
  },
}
export const RoleDecoder: Decoder<Role> = {
  definitionName: 'Role',
  schemaRef: '#/definitions/Role',

  decode(json: unknown): Role {
    const schema = ajv.getSchema(RoleDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${RoleDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, RoleDecoder.definitionName)
  },
}
export const TaskRoleDecoder: Decoder<TaskRole> = {
  definitionName: 'TaskRole',
  schemaRef: '#/definitions/TaskRole',

  decode(json: unknown): TaskRole {
    const schema = ajv.getSchema(TaskRoleDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${TaskRoleDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, TaskRoleDecoder.definitionName)
  },
}
export const OrganizationDecoder: Decoder<Organization> = {
  definitionName: 'Organization',
  schemaRef: '#/definitions/Organization',

  decode(json: unknown): Organization {
    const schema = ajv.getSchema(OrganizationDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${OrganizationDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, OrganizationDecoder.definitionName)
  },
}
export const GetCurrentUserResDecoder: Decoder<GetCurrentUserRes> = {
  definitionName: 'GetCurrentUserRes',
  schemaRef: '#/definitions/GetCurrentUserRes',

  decode(json: unknown): GetCurrentUserRes {
    const schema = ajv.getSchema(GetCurrentUserResDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${GetCurrentUserResDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, GetCurrentUserResDecoder.definitionName)
  },
}
export const ValidationErrorDecoder: Decoder<ValidationError> = {
  definitionName: 'ValidationError',
  schemaRef: '#/definitions/ValidationError',

  decode(json: unknown): ValidationError {
    const schema = ajv.getSchema(ValidationErrorDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ValidationErrorDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ValidationErrorDecoder.definitionName)
  },
}
