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
  Scope,
  Scopes,
  Role,
  NotificationType,
  WorkItemRole,
  UpdateUserRequest,
  UpdateUserAuthRequest,
  UserPublic,
  ValidationError,
  PgtypeJSONB,
  UuidUUID,
  TaskPublic,
  TaskTypePublic,
  TeamPublic,
  TimeEntryPublic,
  WorkItemCommentPublic,
  WorkItemPublic,
  ModelsRole,
  RestUserResponse,
  ModelsScope,
  UserAPIKeyPublic,
  DbTeamPublic,
  DbUserAPIKeyPublic,
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
export const ScopesDecoder: Decoder<Scopes> = {
  definitionName: 'Scopes',
  schemaRef: '#/definitions/Scopes',

  decode(json: unknown): Scopes {
    const schema = ajv.getSchema(ScopesDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ScopesDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ScopesDecoder.definitionName)
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
export const NotificationTypeDecoder: Decoder<NotificationType> = {
  definitionName: 'NotificationType',
  schemaRef: '#/definitions/NotificationType',

  decode(json: unknown): NotificationType {
    const schema = ajv.getSchema(NotificationTypeDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${NotificationTypeDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, NotificationTypeDecoder.definitionName)
  },
}
export const WorkItemRoleDecoder: Decoder<WorkItemRole> = {
  definitionName: 'WorkItemRole',
  schemaRef: '#/definitions/WorkItemRole',

  decode(json: unknown): WorkItemRole {
    const schema = ajv.getSchema(WorkItemRoleDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${WorkItemRoleDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, WorkItemRoleDecoder.definitionName)
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
export const UpdateUserAuthRequestDecoder: Decoder<UpdateUserAuthRequest> = {
  definitionName: 'UpdateUserAuthRequest',
  schemaRef: '#/definitions/UpdateUserAuthRequest',

  decode(json: unknown): UpdateUserAuthRequest {
    const schema = ajv.getSchema(UpdateUserAuthRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${UpdateUserAuthRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, UpdateUserAuthRequestDecoder.definitionName)
  },
}
export const UserPublicDecoder: Decoder<UserPublic> = {
  definitionName: 'UserPublic',
  schemaRef: '#/definitions/UserPublic',

  decode(json: unknown): UserPublic {
    const schema = ajv.getSchema(UserPublicDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${UserPublicDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, UserPublicDecoder.definitionName)
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
export const TaskPublicDecoder: Decoder<TaskPublic> = {
  definitionName: 'TaskPublic',
  schemaRef: '#/definitions/TaskPublic',

  decode(json: unknown): TaskPublic {
    const schema = ajv.getSchema(TaskPublicDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${TaskPublicDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, TaskPublicDecoder.definitionName)
  },
}
export const TaskTypePublicDecoder: Decoder<TaskTypePublic> = {
  definitionName: 'TaskTypePublic',
  schemaRef: '#/definitions/TaskTypePublic',

  decode(json: unknown): TaskTypePublic {
    const schema = ajv.getSchema(TaskTypePublicDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${TaskTypePublicDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, TaskTypePublicDecoder.definitionName)
  },
}
export const TeamPublicDecoder: Decoder<TeamPublic> = {
  definitionName: 'TeamPublic',
  schemaRef: '#/definitions/TeamPublic',

  decode(json: unknown): TeamPublic {
    const schema = ajv.getSchema(TeamPublicDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${TeamPublicDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, TeamPublicDecoder.definitionName)
  },
}
export const TimeEntryPublicDecoder: Decoder<TimeEntryPublic> = {
  definitionName: 'TimeEntryPublic',
  schemaRef: '#/definitions/TimeEntryPublic',

  decode(json: unknown): TimeEntryPublic {
    const schema = ajv.getSchema(TimeEntryPublicDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${TimeEntryPublicDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, TimeEntryPublicDecoder.definitionName)
  },
}
export const WorkItemCommentPublicDecoder: Decoder<WorkItemCommentPublic> = {
  definitionName: 'WorkItemCommentPublic',
  schemaRef: '#/definitions/WorkItemCommentPublic',

  decode(json: unknown): WorkItemCommentPublic {
    const schema = ajv.getSchema(WorkItemCommentPublicDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${WorkItemCommentPublicDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, WorkItemCommentPublicDecoder.definitionName)
  },
}
export const WorkItemPublicDecoder: Decoder<WorkItemPublic> = {
  definitionName: 'WorkItemPublic',
  schemaRef: '#/definitions/WorkItemPublic',

  decode(json: unknown): WorkItemPublic {
    const schema = ajv.getSchema(WorkItemPublicDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${WorkItemPublicDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, WorkItemPublicDecoder.definitionName)
  },
}
export const ModelsRoleDecoder: Decoder<ModelsRole> = {
  definitionName: 'ModelsRole',
  schemaRef: '#/definitions/ModelsRole',

  decode(json: unknown): ModelsRole {
    const schema = ajv.getSchema(ModelsRoleDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsRoleDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsRoleDecoder.definitionName)
  },
}
export const RestUserResponseDecoder: Decoder<RestUserResponse> = {
  definitionName: 'RestUserResponse',
  schemaRef: '#/definitions/RestUserResponse',

  decode(json: unknown): RestUserResponse {
    const schema = ajv.getSchema(RestUserResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${RestUserResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, RestUserResponseDecoder.definitionName)
  },
}
export const ModelsScopeDecoder: Decoder<ModelsScope> = {
  definitionName: 'ModelsScope',
  schemaRef: '#/definitions/ModelsScope',

  decode(json: unknown): ModelsScope {
    const schema = ajv.getSchema(ModelsScopeDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsScopeDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsScopeDecoder.definitionName)
  },
}
export const UserAPIKeyPublicDecoder: Decoder<UserAPIKeyPublic> = {
  definitionName: 'UserAPIKeyPublic',
  schemaRef: '#/definitions/UserAPIKeyPublic',

  decode(json: unknown): UserAPIKeyPublic {
    const schema = ajv.getSchema(UserAPIKeyPublicDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${UserAPIKeyPublicDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, UserAPIKeyPublicDecoder.definitionName)
  },
}
export const DbTeamPublicDecoder: Decoder<DbTeamPublic> = {
  definitionName: 'DbTeamPublic',
  schemaRef: '#/definitions/DbTeamPublic',

  decode(json: unknown): DbTeamPublic {
    const schema = ajv.getSchema(DbTeamPublicDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbTeamPublicDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbTeamPublicDecoder.definitionName)
  },
}
export const DbUserAPIKeyPublicDecoder: Decoder<DbUserAPIKeyPublic> = {
  definitionName: 'DbUserAPIKeyPublic',
  schemaRef: '#/definitions/DbUserAPIKeyPublic',

  decode(json: unknown): DbUserAPIKeyPublic {
    const schema = ajv.getSchema(DbUserAPIKeyPublicDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbUserAPIKeyPublicDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbUserAPIKeyPublicDecoder.definitionName)
  },
}
