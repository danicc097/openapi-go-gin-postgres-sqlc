/* eslint-disable */
// @ts-nocheck
/* eslint-disable */
/* eslint-disable @typescript-eslint/ban-ts-comment */

import Ajv from 'ajv'
import addFormats from 'ajv-formats'
import { Decoder } from './helpers'
import { validateJson } from '../validate'
import {
  InitializeProjectRequest,
  ProjectBoardResponse,
  UserResponse,
  HTTPValidationError,
  Topics,
  Scope,
  Scopes,
  Role,
  NotificationType,
  WorkItemRole,
  UpdateUserRequest,
  UpdateUserAuthRequest,
  ValidationError,
  DbTeamPublic,
  DbUserAPIKeyPublic,
  DbActivityPublic,
  DbKanbanStepPublic,
  DbProjectPublic,
  DbWorkItemTagPublic,
  DbWorkItemTypePublic,
  ReposActivityCreateParams,
  ReposKanbanStepCreateParams,
  ReposTeamCreateParams,
  ReposWorkItemTagCreateParams,
  ReposWorkItemTypeCreateParams,
  ModelsRole,
  UuidUUID,
} from './models'
import jsonSchema from './schema.json'

const ajv = new Ajv({ strict: false, allErrors: true })
addFormats(ajv, { formats: ['int64', 'int32', 'binary', 'date-time'] })
ajv.compile(jsonSchema)

// Decoders
export const InitializeProjectRequestDecoder: Decoder<InitializeProjectRequest> = {
  definitionName: 'InitializeProjectRequest',
  schemaRef: '#/definitions/InitializeProjectRequest',

  decode(json: unknown): InitializeProjectRequest {
    const schema = ajv.getSchema(InitializeProjectRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${InitializeProjectRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, InitializeProjectRequestDecoder.definitionName)
  },
}
export const ProjectBoardResponseDecoder: Decoder<ProjectBoardResponse> = {
  definitionName: 'ProjectBoardResponse',
  schemaRef: '#/definitions/ProjectBoardResponse',

  decode(json: unknown): ProjectBoardResponse {
    const schema = ajv.getSchema(ProjectBoardResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ProjectBoardResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ProjectBoardResponseDecoder.definitionName)
  },
}
export const UserResponseDecoder: Decoder<UserResponse> = {
  definitionName: 'UserResponse',
  schemaRef: '#/definitions/UserResponse',

  decode(json: unknown): UserResponse {
    const schema = ajv.getSchema(UserResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${UserResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, UserResponseDecoder.definitionName)
  },
}
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
export const TopicsDecoder: Decoder<Topics> = {
  definitionName: 'Topics',
  schemaRef: '#/definitions/Topics',

  decode(json: unknown): Topics {
    const schema = ajv.getSchema(TopicsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${TopicsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, TopicsDecoder.definitionName)
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
export const DbActivityPublicDecoder: Decoder<DbActivityPublic> = {
  definitionName: 'DbActivityPublic',
  schemaRef: '#/definitions/DbActivityPublic',

  decode(json: unknown): DbActivityPublic {
    const schema = ajv.getSchema(DbActivityPublicDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbActivityPublicDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbActivityPublicDecoder.definitionName)
  },
}
export const DbKanbanStepPublicDecoder: Decoder<DbKanbanStepPublic> = {
  definitionName: 'DbKanbanStepPublic',
  schemaRef: '#/definitions/DbKanbanStepPublic',

  decode(json: unknown): DbKanbanStepPublic {
    const schema = ajv.getSchema(DbKanbanStepPublicDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbKanbanStepPublicDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbKanbanStepPublicDecoder.definitionName)
  },
}
export const DbProjectPublicDecoder: Decoder<DbProjectPublic> = {
  definitionName: 'DbProjectPublic',
  schemaRef: '#/definitions/DbProjectPublic',

  decode(json: unknown): DbProjectPublic {
    const schema = ajv.getSchema(DbProjectPublicDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbProjectPublicDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbProjectPublicDecoder.definitionName)
  },
}
export const DbWorkItemTagPublicDecoder: Decoder<DbWorkItemTagPublic> = {
  definitionName: 'DbWorkItemTagPublic',
  schemaRef: '#/definitions/DbWorkItemTagPublic',

  decode(json: unknown): DbWorkItemTagPublic {
    const schema = ajv.getSchema(DbWorkItemTagPublicDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbWorkItemTagPublicDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbWorkItemTagPublicDecoder.definitionName)
  },
}
export const DbWorkItemTypePublicDecoder: Decoder<DbWorkItemTypePublic> = {
  definitionName: 'DbWorkItemTypePublic',
  schemaRef: '#/definitions/DbWorkItemTypePublic',

  decode(json: unknown): DbWorkItemTypePublic {
    const schema = ajv.getSchema(DbWorkItemTypePublicDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbWorkItemTypePublicDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbWorkItemTypePublicDecoder.definitionName)
  },
}
export const ReposActivityCreateParamsDecoder: Decoder<ReposActivityCreateParams> = {
  definitionName: 'ReposActivityCreateParams',
  schemaRef: '#/definitions/ReposActivityCreateParams',

  decode(json: unknown): ReposActivityCreateParams {
    const schema = ajv.getSchema(ReposActivityCreateParamsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ReposActivityCreateParamsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ReposActivityCreateParamsDecoder.definitionName)
  },
}
export const ReposKanbanStepCreateParamsDecoder: Decoder<ReposKanbanStepCreateParams> = {
  definitionName: 'ReposKanbanStepCreateParams',
  schemaRef: '#/definitions/ReposKanbanStepCreateParams',

  decode(json: unknown): ReposKanbanStepCreateParams {
    const schema = ajv.getSchema(ReposKanbanStepCreateParamsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ReposKanbanStepCreateParamsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ReposKanbanStepCreateParamsDecoder.definitionName)
  },
}
export const ReposTeamCreateParamsDecoder: Decoder<ReposTeamCreateParams> = {
  definitionName: 'ReposTeamCreateParams',
  schemaRef: '#/definitions/ReposTeamCreateParams',

  decode(json: unknown): ReposTeamCreateParams {
    const schema = ajv.getSchema(ReposTeamCreateParamsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ReposTeamCreateParamsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ReposTeamCreateParamsDecoder.definitionName)
  },
}
export const ReposWorkItemTagCreateParamsDecoder: Decoder<ReposWorkItemTagCreateParams> = {
  definitionName: 'ReposWorkItemTagCreateParams',
  schemaRef: '#/definitions/ReposWorkItemTagCreateParams',

  decode(json: unknown): ReposWorkItemTagCreateParams {
    const schema = ajv.getSchema(ReposWorkItemTagCreateParamsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ReposWorkItemTagCreateParamsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ReposWorkItemTagCreateParamsDecoder.definitionName)
  },
}
export const ReposWorkItemTypeCreateParamsDecoder: Decoder<ReposWorkItemTypeCreateParams> = {
  definitionName: 'ReposWorkItemTypeCreateParams',
  schemaRef: '#/definitions/ReposWorkItemTypeCreateParams',

  decode(json: unknown): ReposWorkItemTypeCreateParams {
    const schema = ajv.getSchema(ReposWorkItemTypeCreateParamsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ReposWorkItemTypeCreateParamsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ReposWorkItemTypeCreateParamsDecoder.definitionName)
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
