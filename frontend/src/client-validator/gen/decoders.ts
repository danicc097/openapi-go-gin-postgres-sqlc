/* eslint-disable */
// @ts-nocheck
/* eslint-disable */
/* eslint-disable @typescript-eslint/ban-ts-comment */

import Ajv from 'ajv'
import addFormats from 'ajv-formats'
import { Decoder } from './helpers'
import { validateJson } from '../validate'
import {
  PaginatedNotificationsResponse,
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
  DemoWorkItemsResponse,
  DemoTwoWorkItemsResponse,
  InitializeProjectRequest,
  ProjectBoardResponse,
  User,
  HTTPValidationError,
  ErrorCode,
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
  WorkItemCreateRequest,
  DbWorkItem,
  WorkItemTagCreateRequest,
  DemoTwoWorkItemCreateRequest,
  DemoWorkItemCreateRequest,
  WorkItemCommentCreateRequest,
  Project,
  DbActivityCreateParams,
  DbTeamCreateParams,
  DbWorkItemTagCreateParams,
  DbWorkItemRole,
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
  ServicesMember,
  DbDemoTwoWorkItem,
  DbDemoTwoWorkItemCreateParams,
  DbWorkItemID,
  DbProjectID,
  DbUserID,
  DbWorkItemTypeID,
  DbNotificationID,
  RestNotification,
  DbNotification,
} from './models'
import jsonSchema from './schema.json'

const ajv = new Ajv({ strict: false, allErrors: true })
addFormats(ajv, { formats: ['int64', 'int32', 'binary', 'date-time', 'date'] })
ajv.compile(jsonSchema)

// Decoders
export const PaginatedNotificationsResponseDecoder: Decoder<PaginatedNotificationsResponse> = {
  definitionName: 'PaginatedNotificationsResponse',
  schemaRef: '#/definitions/PaginatedNotificationsResponse',

  decode(json: unknown): PaginatedNotificationsResponse {
    const schema = ajv.getSchema(PaginatedNotificationsResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${PaginatedNotificationsResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, PaginatedNotificationsResponseDecoder.definitionName)
  },
}
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
export const DemoWorkItemsResponseDecoder: Decoder<DemoWorkItemsResponse> = {
  definitionName: 'DemoWorkItemsResponse',
  schemaRef: '#/definitions/DemoWorkItemsResponse',

  decode(json: unknown): DemoWorkItemsResponse {
    const schema = ajv.getSchema(DemoWorkItemsResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DemoWorkItemsResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DemoWorkItemsResponseDecoder.definitionName)
  },
}
export const DemoTwoWorkItemsResponseDecoder: Decoder<DemoTwoWorkItemsResponse> = {
  definitionName: 'DemoTwoWorkItemsResponse',
  schemaRef: '#/definitions/DemoTwoWorkItemsResponse',

  decode(json: unknown): DemoTwoWorkItemsResponse {
    const schema = ajv.getSchema(DemoTwoWorkItemsResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DemoTwoWorkItemsResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DemoTwoWorkItemsResponseDecoder.definitionName)
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
export const ErrorCodeDecoder: Decoder<ErrorCode> = {
  definitionName: 'ErrorCode',
  schemaRef: '#/definitions/ErrorCode',

  decode(json: unknown): ErrorCode {
    const schema = ajv.getSchema(ErrorCodeDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ErrorCodeDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ErrorCodeDecoder.definitionName)
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
export const WorkItemCreateRequestDecoder: Decoder<WorkItemCreateRequest> = {
  definitionName: 'WorkItemCreateRequest',
  schemaRef: '#/definitions/WorkItemCreateRequest',

  decode(json: unknown): WorkItemCreateRequest {
    const schema = ajv.getSchema(WorkItemCreateRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${WorkItemCreateRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, WorkItemCreateRequestDecoder.definitionName)
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
export const WorkItemTagCreateRequestDecoder: Decoder<WorkItemTagCreateRequest> = {
  definitionName: 'WorkItemTagCreateRequest',
  schemaRef: '#/definitions/WorkItemTagCreateRequest',

  decode(json: unknown): WorkItemTagCreateRequest {
    const schema = ajv.getSchema(WorkItemTagCreateRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${WorkItemTagCreateRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, WorkItemTagCreateRequestDecoder.definitionName)
  },
}
export const DemoTwoWorkItemCreateRequestDecoder: Decoder<DemoTwoWorkItemCreateRequest> = {
  definitionName: 'DemoTwoWorkItemCreateRequest',
  schemaRef: '#/definitions/DemoTwoWorkItemCreateRequest',

  decode(json: unknown): DemoTwoWorkItemCreateRequest {
    const schema = ajv.getSchema(DemoTwoWorkItemCreateRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DemoTwoWorkItemCreateRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DemoTwoWorkItemCreateRequestDecoder.definitionName)
  },
}
export const DemoWorkItemCreateRequestDecoder: Decoder<DemoWorkItemCreateRequest> = {
  definitionName: 'DemoWorkItemCreateRequest',
  schemaRef: '#/definitions/DemoWorkItemCreateRequest',

  decode(json: unknown): DemoWorkItemCreateRequest {
    const schema = ajv.getSchema(DemoWorkItemCreateRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DemoWorkItemCreateRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DemoWorkItemCreateRequestDecoder.definitionName)
  },
}
export const WorkItemCommentCreateRequestDecoder: Decoder<WorkItemCommentCreateRequest> = {
  definitionName: 'WorkItemCommentCreateRequest',
  schemaRef: '#/definitions/WorkItemCommentCreateRequest',

  decode(json: unknown): WorkItemCommentCreateRequest {
    const schema = ajv.getSchema(WorkItemCommentCreateRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${WorkItemCommentCreateRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, WorkItemCommentCreateRequestDecoder.definitionName)
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
export const DbDemoTwoWorkItemDecoder: Decoder<DbDemoTwoWorkItem> = {
  definitionName: 'DbDemoTwoWorkItem',
  schemaRef: '#/definitions/DbDemoTwoWorkItem',

  decode(json: unknown): DbDemoTwoWorkItem {
    const schema = ajv.getSchema(DbDemoTwoWorkItemDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbDemoTwoWorkItemDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbDemoTwoWorkItemDecoder.definitionName)
  },
}
export const DbDemoTwoWorkItemCreateParamsDecoder: Decoder<DbDemoTwoWorkItemCreateParams> = {
  definitionName: 'DbDemoTwoWorkItemCreateParams',
  schemaRef: '#/definitions/DbDemoTwoWorkItemCreateParams',

  decode(json: unknown): DbDemoTwoWorkItemCreateParams {
    const schema = ajv.getSchema(DbDemoTwoWorkItemCreateParamsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbDemoTwoWorkItemCreateParamsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbDemoTwoWorkItemCreateParamsDecoder.definitionName)
  },
}
export const DbWorkItemIDDecoder: Decoder<DbWorkItemID> = {
  definitionName: 'DbWorkItemID',
  schemaRef: '#/definitions/DbWorkItemID',

  decode(json: unknown): DbWorkItemID {
    const schema = ajv.getSchema(DbWorkItemIDDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbWorkItemIDDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbWorkItemIDDecoder.definitionName)
  },
}
export const DbProjectIDDecoder: Decoder<DbProjectID> = {
  definitionName: 'DbProjectID',
  schemaRef: '#/definitions/DbProjectID',

  decode(json: unknown): DbProjectID {
    const schema = ajv.getSchema(DbProjectIDDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbProjectIDDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbProjectIDDecoder.definitionName)
  },
}
export const DbUserIDDecoder: Decoder<DbUserID> = {
  definitionName: 'DbUserID',
  schemaRef: '#/definitions/DbUserID',

  decode(json: unknown): DbUserID {
    const schema = ajv.getSchema(DbUserIDDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbUserIDDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbUserIDDecoder.definitionName)
  },
}
export const DbWorkItemTypeIDDecoder: Decoder<DbWorkItemTypeID> = {
  definitionName: 'DbWorkItemTypeID',
  schemaRef: '#/definitions/DbWorkItemTypeID',

  decode(json: unknown): DbWorkItemTypeID {
    const schema = ajv.getSchema(DbWorkItemTypeIDDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbWorkItemTypeIDDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbWorkItemTypeIDDecoder.definitionName)
  },
}
export const DbNotificationIDDecoder: Decoder<DbNotificationID> = {
  definitionName: 'DbNotificationID',
  schemaRef: '#/definitions/DbNotificationID',

  decode(json: unknown): DbNotificationID {
    const schema = ajv.getSchema(DbNotificationIDDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbNotificationIDDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbNotificationIDDecoder.definitionName)
  },
}
export const RestNotificationDecoder: Decoder<RestNotification> = {
  definitionName: 'RestNotification',
  schemaRef: '#/definitions/RestNotification',

  decode(json: unknown): RestNotification {
    const schema = ajv.getSchema(RestNotificationDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${RestNotificationDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, RestNotificationDecoder.definitionName)
  },
}
export const DbNotificationDecoder: Decoder<DbNotification> = {
  definitionName: 'DbNotification',
  schemaRef: '#/definitions/DbNotification',

  decode(json: unknown): DbNotification {
    const schema = ajv.getSchema(DbNotificationDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbNotificationDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbNotificationDecoder.definitionName)
  },
}
