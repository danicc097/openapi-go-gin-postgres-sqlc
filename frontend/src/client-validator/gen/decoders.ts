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
  DbDemoWorkItem,
  DbUserAPIKey,
  DbUser,
  DbTimeEntry,
  DbWorkItemComment,
  ProjectConfig,
  ProjectConfigField,
  RestDemoWorkItemsResponse,
  InitializeProjectRequest,
  RestProjectBoardResponse,
  UserResponse,
  HTTPValidationError,
  HTTPError,
  Topics,
  Scope,
  Scopes,
  Role,
  WorkItemRole,
  UpdateUserRequest,
  UpdateUserAuthRequest,
  ValidationError,
  UuidUUID,
  PgtypeJSONB,
  DbWorkItem,
  RestWorkItemTagCreateRequest,
  RestDemoWorkItemCreateRequest,
  RestWorkItemCommentCreateRequest,
  Project,
  DbActivityCreateParams,
  DbKanbanStepCreateParams,
  DbTeamCreateParams,
  DbWorkItemTagCreateParams,
  DbWorkItemTypeCreateParams,
  DbWorkItemRole,
  DbWorkItem_AssignedUser,
  NotificationType,
  DemoProjectKanbanSteps,
  DemoProject2KanbanSteps,
  Demo2WorkItemTypes,
  DemoKanbanSteps,
  DemoTwoKanbanSteps,
  DemoTwoWorkItemTypes,
  DemoWorkItemTypes,
  DbDemoWorkItemCreateParams,
  DbWorkItemCreateParams,
  ModelsWorkItemRole,
  ServicesMember,
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
export const DbDemoWorkItemDecoder: Decoder<DbDemoWorkItem> = {
  definitionName: 'DbDemoWorkItem',
  schemaRef: '#/definitions/DbDemoWorkItem',

  decode(json: unknown): DbDemoWorkItem {
    const schema = ajv.getSchema(DbDemoWorkItemDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbDemoWorkItemDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbDemoWorkItemDecoder.definitionName)
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
export const ProjectConfigFieldDecoder: Decoder<ProjectConfigField> = {
  definitionName: 'ProjectConfigField',
  schemaRef: '#/definitions/ProjectConfigField',

  decode(json: unknown): ProjectConfigField {
    const schema = ajv.getSchema(ProjectConfigFieldDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ProjectConfigFieldDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ProjectConfigFieldDecoder.definitionName)
  },
}
export const RestDemoWorkItemsResponseDecoder: Decoder<RestDemoWorkItemsResponse> = {
  definitionName: 'RestDemoWorkItemsResponse',
  schemaRef: '#/definitions/RestDemoWorkItemsResponse',

  decode(json: unknown): RestDemoWorkItemsResponse {
    const schema = ajv.getSchema(RestDemoWorkItemsResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${RestDemoWorkItemsResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, RestDemoWorkItemsResponseDecoder.definitionName)
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
export const HTTPErrorDecoder: Decoder<HTTPError> = {
  definitionName: 'HTTPError',
  schemaRef: '#/definitions/HTTPError',

  decode(json: unknown): HTTPError {
    const schema = ajv.getSchema(HTTPErrorDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${HTTPErrorDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, HTTPErrorDecoder.definitionName)
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
export const RestWorkItemTagCreateRequestDecoder: Decoder<RestWorkItemTagCreateRequest> = {
  definitionName: 'RestWorkItemTagCreateRequest',
  schemaRef: '#/definitions/RestWorkItemTagCreateRequest',

  decode(json: unknown): RestWorkItemTagCreateRequest {
    const schema = ajv.getSchema(RestWorkItemTagCreateRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${RestWorkItemTagCreateRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, RestWorkItemTagCreateRequestDecoder.definitionName)
  },
}
export const RestDemoWorkItemCreateRequestDecoder: Decoder<RestDemoWorkItemCreateRequest> = {
  definitionName: 'RestDemoWorkItemCreateRequest',
  schemaRef: '#/definitions/RestDemoWorkItemCreateRequest',

  decode(json: unknown): RestDemoWorkItemCreateRequest {
    const schema = ajv.getSchema(RestDemoWorkItemCreateRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${RestDemoWorkItemCreateRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, RestDemoWorkItemCreateRequestDecoder.definitionName)
  },
}
export const RestWorkItemCommentCreateRequestDecoder: Decoder<RestWorkItemCommentCreateRequest> = {
  definitionName: 'RestWorkItemCommentCreateRequest',
  schemaRef: '#/definitions/RestWorkItemCommentCreateRequest',

  decode(json: unknown): RestWorkItemCommentCreateRequest {
    const schema = ajv.getSchema(RestWorkItemCommentCreateRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${RestWorkItemCommentCreateRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, RestWorkItemCommentCreateRequestDecoder.definitionName)
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
export const DbWorkItem_AssignedUserDecoder: Decoder<DbWorkItem_AssignedUser> = {
  definitionName: 'DbWorkItem_AssignedUser',
  schemaRef: '#/definitions/DbWorkItem_AssignedUser',

  decode(json: unknown): DbWorkItem_AssignedUser {
    const schema = ajv.getSchema(DbWorkItem_AssignedUserDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbWorkItem_AssignedUserDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbWorkItem_AssignedUserDecoder.definitionName)
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
export const Demo2WorkItemTypesDecoder: Decoder<Demo2WorkItemTypes> = {
  definitionName: 'Demo2WorkItemTypes',
  schemaRef: '#/definitions/Demo2WorkItemTypes',

  decode(json: unknown): Demo2WorkItemTypes {
    const schema = ajv.getSchema(Demo2WorkItemTypesDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${Demo2WorkItemTypesDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, Demo2WorkItemTypesDecoder.definitionName)
  },
}
export const DemoKanbanStepsDecoder: Decoder<DemoKanbanSteps> = {
  definitionName: 'DemoKanbanSteps',
  schemaRef: '#/definitions/DemoKanbanSteps',

  decode(json: unknown): DemoKanbanSteps {
    const schema = ajv.getSchema(DemoKanbanStepsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DemoKanbanStepsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DemoKanbanStepsDecoder.definitionName)
  },
}
export const DemoTwoKanbanStepsDecoder: Decoder<DemoTwoKanbanSteps> = {
  definitionName: 'DemoTwoKanbanSteps',
  schemaRef: '#/definitions/DemoTwoKanbanSteps',

  decode(json: unknown): DemoTwoKanbanSteps {
    const schema = ajv.getSchema(DemoTwoKanbanStepsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DemoTwoKanbanStepsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DemoTwoKanbanStepsDecoder.definitionName)
  },
}
export const DemoTwoWorkItemTypesDecoder: Decoder<DemoTwoWorkItemTypes> = {
  definitionName: 'DemoTwoWorkItemTypes',
  schemaRef: '#/definitions/DemoTwoWorkItemTypes',

  decode(json: unknown): DemoTwoWorkItemTypes {
    const schema = ajv.getSchema(DemoTwoWorkItemTypesDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DemoTwoWorkItemTypesDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DemoTwoWorkItemTypesDecoder.definitionName)
  },
}
export const DemoWorkItemTypesDecoder: Decoder<DemoWorkItemTypes> = {
  definitionName: 'DemoWorkItemTypes',
  schemaRef: '#/definitions/DemoWorkItemTypes',

  decode(json: unknown): DemoWorkItemTypes {
    const schema = ajv.getSchema(DemoWorkItemTypesDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DemoWorkItemTypesDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DemoWorkItemTypesDecoder.definitionName)
  },
}
export const DbDemoWorkItemCreateParamsDecoder: Decoder<DbDemoWorkItemCreateParams> = {
  definitionName: 'DbDemoWorkItemCreateParams',
  schemaRef: '#/definitions/DbDemoWorkItemCreateParams',

  decode(json: unknown): DbDemoWorkItemCreateParams {
    const schema = ajv.getSchema(DbDemoWorkItemCreateParamsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbDemoWorkItemCreateParamsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbDemoWorkItemCreateParamsDecoder.definitionName)
  },
}
export const DbWorkItemCreateParamsDecoder: Decoder<DbWorkItemCreateParams> = {
  definitionName: 'DbWorkItemCreateParams',
  schemaRef: '#/definitions/DbWorkItemCreateParams',

  decode(json: unknown): DbWorkItemCreateParams {
    const schema = ajv.getSchema(DbWorkItemCreateParamsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbWorkItemCreateParamsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbWorkItemCreateParamsDecoder.definitionName)
  },
}
export const ModelsWorkItemRoleDecoder: Decoder<ModelsWorkItemRole> = {
  definitionName: 'ModelsWorkItemRole',
  schemaRef: '#/definitions/ModelsWorkItemRole',

  decode(json: unknown): ModelsWorkItemRole {
    const schema = ajv.getSchema(ModelsWorkItemRoleDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsWorkItemRoleDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsWorkItemRoleDecoder.definitionName)
  },
}
export const ServicesMemberDecoder: Decoder<ServicesMember> = {
  definitionName: 'ServicesMember',
  schemaRef: '#/definitions/ServicesMember',

  decode(json: unknown): ServicesMember {
    const schema = ajv.getSchema(ServicesMemberDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ServicesMemberDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ServicesMemberDecoder.definitionName)
  },
}
