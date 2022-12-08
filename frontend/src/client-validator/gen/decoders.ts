/* eslint-disable */
// @ts-nocheck
/* eslint-disable */
/* eslint-disable @typescript-eslint/ban-ts-comment */

import Ajv from 'ajv'
import addFormats from 'ajv-formats'
import { Decoder } from './helpers'
import { validateJson } from '../validate'
import {
  ProjectBoardCreateRequest,
  ProjectBoardResponse,
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
  PgtypeJSONB,
  UuidUUID,
  TeamPublic,
  TimeEntryPublic,
  WorkItemCommentPublic,
  ModelsRole,
  RestUserResponse,
  ModelsScope,
  UserAPIKeyPublic,
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
  RestProjectBoardCreateRequest,
  RestProjectBoardResponse,
} from './models'
import jsonSchema from './schema.json'

const ajv = new Ajv({ strict: false, allErrors: true })
addFormats(ajv, { formats: ['int64', 'int32', 'binary', 'date-time'] })
ajv.compile(jsonSchema)

// Decoders
export const ProjectBoardCreateRequestDecoder: Decoder<ProjectBoardCreateRequest> = {
  definitionName: 'ProjectBoardCreateRequest',
  schemaRef: '#/definitions/ProjectBoardCreateRequest',

  decode(json: unknown): ProjectBoardCreateRequest {
    const schema = ajv.getSchema(ProjectBoardCreateRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ProjectBoardCreateRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ProjectBoardCreateRequestDecoder.definitionName)
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
export const RestProjectBoardCreateRequestDecoder: Decoder<RestProjectBoardCreateRequest> = {
  definitionName: 'RestProjectBoardCreateRequest',
  schemaRef: '#/definitions/RestProjectBoardCreateRequest',

  decode(json: unknown): RestProjectBoardCreateRequest {
    const schema = ajv.getSchema(RestProjectBoardCreateRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${RestProjectBoardCreateRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, RestProjectBoardCreateRequestDecoder.definitionName)
  },
}
export const RestProjectBoardResponseDecoder: Decoder<RestProjectBoardResponse> = {
  definitionName: 'RestProjectBoardResponse',
  schemaRef: '#/definitions/RestProjectBoardResponse',

  decode(json: unknown): RestProjectBoardResponse {
    const schema = ajv.getSchema(RestProjectBoardResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${RestProjectBoardResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, RestProjectBoardResponseDecoder.definitionName)
  },
}
