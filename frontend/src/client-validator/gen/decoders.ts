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
  User,
  ValidationError,
  PgtypeJSONB,
  Task,
  TaskType,
  Team,
  TimeEntry,
  UserAPIKey,
  UuidUUID,
  WorkItem,
  WorkItemComment,
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
export const UserDecoder: Decoder<User> = {
  definitionName: 'User',
  schemaRef: '#/definitions/User',

  decode(json: unknown): User {
    const schema = ajv.getSchema(UserDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${UserDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, UserDecoder.definitionName)
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
export const PgtypeJSONBDecoder: Decoder<PgtypeJSONB> = {
  definitionName: 'PgtypeJSONB',
  schemaRef: '#/definitions/PgtypeJSONB',

  decode(json: unknown): PgtypeJSONB {
    const schema = ajv.getSchema(PgtypeJSONBDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${PgtypeJSONBDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, PgtypeJSONBDecoder.definitionName)
  },
}
export const TaskDecoder: Decoder<Task> = {
  definitionName: 'Task',
  schemaRef: '#/definitions/Task',

  decode(json: unknown): Task {
    const schema = ajv.getSchema(TaskDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${TaskDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, TaskDecoder.definitionName)
  },
}
export const TaskTypeDecoder: Decoder<TaskType> = {
  definitionName: 'TaskType',
  schemaRef: '#/definitions/TaskType',

  decode(json: unknown): TaskType {
    const schema = ajv.getSchema(TaskTypeDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${TaskTypeDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, TaskTypeDecoder.definitionName)
  },
}
export const TeamDecoder: Decoder<Team> = {
  definitionName: 'Team',
  schemaRef: '#/definitions/Team',

  decode(json: unknown): Team {
    const schema = ajv.getSchema(TeamDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${TeamDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, TeamDecoder.definitionName)
  },
}
export const TimeEntryDecoder: Decoder<TimeEntry> = {
  definitionName: 'TimeEntry',
  schemaRef: '#/definitions/TimeEntry',

  decode(json: unknown): TimeEntry {
    const schema = ajv.getSchema(TimeEntryDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${TimeEntryDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, TimeEntryDecoder.definitionName)
  },
}
export const UserAPIKeyDecoder: Decoder<UserAPIKey> = {
  definitionName: 'UserAPIKey',
  schemaRef: '#/definitions/UserAPIKey',

  decode(json: unknown): UserAPIKey {
    const schema = ajv.getSchema(UserAPIKeyDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${UserAPIKeyDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, UserAPIKeyDecoder.definitionName)
  },
}
export const UuidUUIDDecoder: Decoder<UuidUUID> = {
  definitionName: 'UuidUUID',
  schemaRef: '#/definitions/UuidUUID',

  decode(json: unknown): UuidUUID {
    const schema = ajv.getSchema(UuidUUIDDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${UuidUUIDDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, UuidUUIDDecoder.definitionName)
  },
}
export const WorkItemDecoder: Decoder<WorkItem> = {
  definitionName: 'WorkItem',
  schemaRef: '#/definitions/WorkItem',

  decode(json: unknown): WorkItem {
    const schema = ajv.getSchema(WorkItemDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${WorkItemDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, WorkItemDecoder.definitionName)
  },
}
export const WorkItemCommentDecoder: Decoder<WorkItemComment> = {
  definitionName: 'WorkItemComment',
  schemaRef: '#/definitions/WorkItemComment',

  decode(json: unknown): WorkItemComment {
    const schema = ajv.getSchema(WorkItemCommentDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${WorkItemCommentDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, WorkItemCommentDecoder.definitionName)
  },
}
