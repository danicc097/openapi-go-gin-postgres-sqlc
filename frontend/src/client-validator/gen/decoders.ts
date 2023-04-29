/* eslint-disable */
// @ts-nocheck
/* eslint-disable */
/* eslint-disable @typescript-eslint/ban-ts-comment */

import Ajv from 'ajv'
import addFormats from 'ajv-formats'
import { Decoder } from './helpers'
import { validateJson } from '../validate'
import {
  DbActivity,
  DbKanbanStep,
  DbProject,
  DbTeam,
  DbWorkItemTag,
  DbWorkItemType,
  DbDemoProjectWorkItem,
  DbUserAPIKey,
  DbUser,
  DbTimeEntry,
  DbWorkItemComment,
  ProjectConfig,
  RestDemoProjectWorkItemsResponse,
  InitializeProjectRequest,
  RestProjectBoardResponse,
  UserResponse,
  HTTPValidationError,
  Topics,
  Scope,
  Scopes,
  Role,
  WorkItemRole,
  UpdateUserRequest,
  UpdateUserAuthRequest,
  ValidationError,
  HttpErrorType,
  UuidUUID,
  PgtypeJSONB,
  ModelsProjectConfigField,
  DbProject2WorkItem,
  DbWorkItem,
  Project,
  DbActivityCreateParams,
  DbKanbanStepCreateParams,
  DbTeamCreateParams,
  DbWorkItemTagCreateParams,
  DbWorkItemTypeCreateParams,
  ModelsProject,
  ModelsRole,
  DbWorkItemRole,
  DbWorkItem_Member,
  NotificationType,
  DemoProjectKanbanSteps,
  DemoProject2KanbanSteps,
} from './models'
import jsonSchema from './schema.json'

const ajv = new Ajv({ strict: false, allErrors: true })
addFormats(ajv, { formats: ['int64', 'int32', 'binary', 'date-time'] })
ajv.compile(jsonSchema)

// Decoders
export const DbActivityDecoder: Decoder<DbActivity> = {
  definitionName: 'DbActivity',
  schemaRef: '#/definitions/DbActivity',

  decode(json: unknown): DbActivity {
    const schema = ajv.getSchema(DbActivityDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbActivityDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbActivityDecoder.definitionName)
  },
}
export const DbKanbanStepDecoder: Decoder<DbKanbanStep> = {
  definitionName: 'DbKanbanStep',
  schemaRef: '#/definitions/DbKanbanStep',

  decode(json: unknown): DbKanbanStep {
    const schema = ajv.getSchema(DbKanbanStepDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbKanbanStepDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbKanbanStepDecoder.definitionName)
  },
}
export const DbProjectDecoder: Decoder<DbProject> = {
  definitionName: 'DbProject',
  schemaRef: '#/definitions/DbProject',

  decode(json: unknown): DbProject {
    const schema = ajv.getSchema(DbProjectDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbProjectDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbProjectDecoder.definitionName)
  },
}
export const DbTeamDecoder: Decoder<DbTeam> = {
  definitionName: 'DbTeam',
  schemaRef: '#/definitions/DbTeam',

  decode(json: unknown): DbTeam {
    const schema = ajv.getSchema(DbTeamDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbTeamDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbTeamDecoder.definitionName)
  },
}
export const DbWorkItemTagDecoder: Decoder<DbWorkItemTag> = {
  definitionName: 'DbWorkItemTag',
  schemaRef: '#/definitions/DbWorkItemTag',

  decode(json: unknown): DbWorkItemTag {
    const schema = ajv.getSchema(DbWorkItemTagDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbWorkItemTagDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbWorkItemTagDecoder.definitionName)
  },
}
export const DbWorkItemTypeDecoder: Decoder<DbWorkItemType> = {
  definitionName: 'DbWorkItemType',
  schemaRef: '#/definitions/DbWorkItemType',

  decode(json: unknown): DbWorkItemType {
    const schema = ajv.getSchema(DbWorkItemTypeDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbWorkItemTypeDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbWorkItemTypeDecoder.definitionName)
  },
}
export const DbDemoProjectWorkItemDecoder: Decoder<DbDemoProjectWorkItem> = {
  definitionName: 'DbDemoProjectWorkItem',
  schemaRef: '#/definitions/DbDemoProjectWorkItem',

  decode(json: unknown): DbDemoProjectWorkItem {
    const schema = ajv.getSchema(DbDemoProjectWorkItemDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbDemoProjectWorkItemDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbDemoProjectWorkItemDecoder.definitionName)
  },
}
export const DbUserAPIKeyDecoder: Decoder<DbUserAPIKey> = {
  definitionName: 'DbUserAPIKey',
  schemaRef: '#/definitions/DbUserAPIKey',

  decode(json: unknown): DbUserAPIKey {
    const schema = ajv.getSchema(DbUserAPIKeyDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbUserAPIKeyDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbUserAPIKeyDecoder.definitionName)
  },
}
export const DbUserDecoder: Decoder<DbUser> = {
  definitionName: 'DbUser',
  schemaRef: '#/definitions/DbUser',

  decode(json: unknown): DbUser {
    const schema = ajv.getSchema(DbUserDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbUserDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbUserDecoder.definitionName)
  },
}
export const DbTimeEntryDecoder: Decoder<DbTimeEntry> = {
  definitionName: 'DbTimeEntry',
  schemaRef: '#/definitions/DbTimeEntry',

  decode(json: unknown): DbTimeEntry {
    const schema = ajv.getSchema(DbTimeEntryDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbTimeEntryDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbTimeEntryDecoder.definitionName)
  },
}
export const DbWorkItemCommentDecoder: Decoder<DbWorkItemComment> = {
  definitionName: 'DbWorkItemComment',
  schemaRef: '#/definitions/DbWorkItemComment',

  decode(json: unknown): DbWorkItemComment {
    const schema = ajv.getSchema(DbWorkItemCommentDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbWorkItemCommentDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbWorkItemCommentDecoder.definitionName)
  },
}
export const ProjectConfigDecoder: Decoder<ProjectConfig> = {
  definitionName: 'ProjectConfig',
  schemaRef: '#/definitions/ProjectConfig',

  decode(json: unknown): ProjectConfig {
    const schema = ajv.getSchema(ProjectConfigDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ProjectConfigDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ProjectConfigDecoder.definitionName)
  },
}
export const RestDemoProjectWorkItemsResponseDecoder: Decoder<RestDemoProjectWorkItemsResponse> = {
  definitionName: 'RestDemoProjectWorkItemsResponse',
  schemaRef: '#/definitions/RestDemoProjectWorkItemsResponse',

  decode(json: unknown): RestDemoProjectWorkItemsResponse {
    const schema = ajv.getSchema(RestDemoProjectWorkItemsResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${RestDemoProjectWorkItemsResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, RestDemoProjectWorkItemsResponseDecoder.definitionName)
  },
}
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
export const HttpErrorTypeDecoder: Decoder<HttpErrorType> = {
  definitionName: 'HttpErrorType',
  schemaRef: '#/definitions/HttpErrorType',

  decode(json: unknown): HttpErrorType {
    const schema = ajv.getSchema(HttpErrorTypeDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${HttpErrorTypeDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, HttpErrorTypeDecoder.definitionName)
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
export const ModelsProjectConfigFieldDecoder: Decoder<ModelsProjectConfigField> = {
  definitionName: 'ModelsProjectConfigField',
  schemaRef: '#/definitions/ModelsProjectConfigField',

  decode(json: unknown): ModelsProjectConfigField {
    const schema = ajv.getSchema(ModelsProjectConfigFieldDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsProjectConfigFieldDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsProjectConfigFieldDecoder.definitionName)
  },
}
export const DbProject2WorkItemDecoder: Decoder<DbProject2WorkItem> = {
  definitionName: 'DbProject2WorkItem',
  schemaRef: '#/definitions/DbProject2WorkItem',

  decode(json: unknown): DbProject2WorkItem {
    const schema = ajv.getSchema(DbProject2WorkItemDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbProject2WorkItemDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbProject2WorkItemDecoder.definitionName)
  },
}
export const DbWorkItemDecoder: Decoder<DbWorkItem> = {
  definitionName: 'DbWorkItem',
  schemaRef: '#/definitions/DbWorkItem',

  decode(json: unknown): DbWorkItem {
    const schema = ajv.getSchema(DbWorkItemDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbWorkItemDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbWorkItemDecoder.definitionName)
  },
}
export const ProjectDecoder: Decoder<Project> = {
  definitionName: 'Project',
  schemaRef: '#/definitions/Project',

  decode(json: unknown): Project {
    const schema = ajv.getSchema(ProjectDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ProjectDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ProjectDecoder.definitionName)
  },
}
export const DbActivityCreateParamsDecoder: Decoder<DbActivityCreateParams> = {
  definitionName: 'DbActivityCreateParams',
  schemaRef: '#/definitions/DbActivityCreateParams',

  decode(json: unknown): DbActivityCreateParams {
    const schema = ajv.getSchema(DbActivityCreateParamsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbActivityCreateParamsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbActivityCreateParamsDecoder.definitionName)
  },
}
export const DbKanbanStepCreateParamsDecoder: Decoder<DbKanbanStepCreateParams> = {
  definitionName: 'DbKanbanStepCreateParams',
  schemaRef: '#/definitions/DbKanbanStepCreateParams',

  decode(json: unknown): DbKanbanStepCreateParams {
    const schema = ajv.getSchema(DbKanbanStepCreateParamsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbKanbanStepCreateParamsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbKanbanStepCreateParamsDecoder.definitionName)
  },
}
export const DbTeamCreateParamsDecoder: Decoder<DbTeamCreateParams> = {
  definitionName: 'DbTeamCreateParams',
  schemaRef: '#/definitions/DbTeamCreateParams',

  decode(json: unknown): DbTeamCreateParams {
    const schema = ajv.getSchema(DbTeamCreateParamsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbTeamCreateParamsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbTeamCreateParamsDecoder.definitionName)
  },
}
export const DbWorkItemTagCreateParamsDecoder: Decoder<DbWorkItemTagCreateParams> = {
  definitionName: 'DbWorkItemTagCreateParams',
  schemaRef: '#/definitions/DbWorkItemTagCreateParams',

  decode(json: unknown): DbWorkItemTagCreateParams {
    const schema = ajv.getSchema(DbWorkItemTagCreateParamsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbWorkItemTagCreateParamsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbWorkItemTagCreateParamsDecoder.definitionName)
  },
}
export const DbWorkItemTypeCreateParamsDecoder: Decoder<DbWorkItemTypeCreateParams> = {
  definitionName: 'DbWorkItemTypeCreateParams',
  schemaRef: '#/definitions/DbWorkItemTypeCreateParams',

  decode(json: unknown): DbWorkItemTypeCreateParams {
    const schema = ajv.getSchema(DbWorkItemTypeCreateParamsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbWorkItemTypeCreateParamsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbWorkItemTypeCreateParamsDecoder.definitionName)
  },
}
export const ModelsProjectDecoder: Decoder<ModelsProject> = {
  definitionName: 'ModelsProject',
  schemaRef: '#/definitions/ModelsProject',

  decode(json: unknown): ModelsProject {
    const schema = ajv.getSchema(ModelsProjectDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsProjectDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsProjectDecoder.definitionName)
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
export const DbWorkItemRoleDecoder: Decoder<DbWorkItemRole> = {
  definitionName: 'DbWorkItemRole',
  schemaRef: '#/definitions/DbWorkItemRole',

  decode(json: unknown): DbWorkItemRole {
    const schema = ajv.getSchema(DbWorkItemRoleDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbWorkItemRoleDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbWorkItemRoleDecoder.definitionName)
  },
}
export const DbWorkItem_MemberDecoder: Decoder<DbWorkItem_Member> = {
  definitionName: 'DbWorkItem_Member',
  schemaRef: '#/definitions/DbWorkItem_Member',

  decode(json: unknown): DbWorkItem_Member {
    const schema = ajv.getSchema(DbWorkItem_MemberDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbWorkItem_MemberDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbWorkItem_MemberDecoder.definitionName)
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
export const DemoProjectKanbanStepsDecoder: Decoder<DemoProjectKanbanSteps> = {
  definitionName: 'DemoProjectKanbanSteps',
  schemaRef: '#/definitions/DemoProjectKanbanSteps',

  decode(json: unknown): DemoProjectKanbanSteps {
    const schema = ajv.getSchema(DemoProjectKanbanStepsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DemoProjectKanbanStepsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DemoProjectKanbanStepsDecoder.definitionName)
  },
}
export const DemoProject2KanbanStepsDecoder: Decoder<DemoProject2KanbanSteps> = {
  definitionName: 'DemoProject2KanbanSteps',
  schemaRef: '#/definitions/DemoProject2KanbanSteps',

  decode(json: unknown): DemoProject2KanbanSteps {
    const schema = ajv.getSchema(DemoProject2KanbanStepsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DemoProject2KanbanStepsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DemoProject2KanbanStepsDecoder.definitionName)
  },
}
