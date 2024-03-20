/* eslint-disable */
// @ts-nocheck
/* eslint-disable */
/* eslint-disable @typescript-eslint/ban-ts-comment */

import Ajv from 'ajv'
import addFormats from 'ajv-formats'
import { Decoder } from './helpers'
import { validateJson } from '../validate'
import {
  Activity,
  CreateActivityRequest,
  CreateDemoTwoWorkItemRequest,
  CreateDemoWorkItemRequest,
  CreateProjectBoardRequest,
  CreateTeamRequest,
  CreateWorkItemCommentRequest,
  CreateWorkItemTagRequest,
  CreateWorkItemTypeRequest,
  DbDemoTwoWorkItem,
  DbDemoTwoWorkItemCreateParams,
  DbDemoWorkItem,
  DbDemoWorkItemCreateParams,
  DbKanbanStep,
  DbNotification,
  DbProject,
  DbTeam,
  DbTeamCreateParams,
  DbTimeEntry,
  DbUser,
  DbUserAPIKey,
  DbUserID,
  DbUserWIAUWorkItem,
  DbWorkItem,
  DbWorkItemComment,
  DbWorkItemCreateParams,
  DbWorkItemTag,
  DbWorkItemTagCreateParams,
  DbWorkItemType,
  Notification,
  PaginatedNotificationsResponse,
  PaginatedUsersResponse,
  PaginationPage,
  ProjectBoard,
  ServicesMember,
  SharedWorkItemJoins,
  Team,
  UpdateActivityRequest,
  UpdateTeamRequest,
  UpdateWorkItemCommentRequest,
  UpdateWorkItemTagRequest,
  UpdateWorkItemTypeRequest,
  User,
  WorkItemComment,
  WorkItemTag,
  WorkItemType,
  Direction,
  DbActivity,
  ProjectConfig,
  ProjectConfigField,
  HTTPValidationError,
  ErrorCode,
  HTTPError,
  Topics,
  Topic,
  Scope,
  Scopes,
  Role,
  WorkItemRole,
  UpdateUserRequest,
  UpdateUserAuthRequest,
  ValidationError,
  UuidUUID,
  WorkItem,
  CreateWorkItemRequest,
  Project,
  DbActivityCreateParams,
  DbWorkItemRole,
  NotificationType,
  DemoTwoWorkItemTypes,
  DemoWorkItemTypes,
  DbWorkItemID,
  DbProjectID,
  DbWorkItemTypeID,
  DbNotificationID,
  DbUserNotification,
  DemoKanbanSteps,
  DemoTwoKanbanSteps,
  DbUserWIAWorkItem,
  DbWorkItemM2MAssigneeWIA,
  CreateTimeEntryRequest,
  TimeEntry,
  UpdateTimeEntryRequest,
  DemoTwoWorkItem,
  DemoWorkItem,
  WorkItemBase,
  PaginationFilterModes,
  DbCacheDemoWorkItemJoins,
  DbUserJoins,
  GetCacheDemoWorkItemQueryParameters,
  GetPaginatedUsersQueryParameters,
  GetCurrentUserQueryParameters,
} from './models'
import jsonSchema from './schema.json'

const ajv = new Ajv({ strict: false, allErrors: true })
addFormats(ajv, { formats: ['int64', 'int32', 'binary', 'date-time', 'date'] })
ajv.compile(jsonSchema)

// Decoders
export const ActivityDecoder: Decoder<Activity> = {
  definitionName: 'Activity',
  schemaRef: '#/definitions/Activity',

  decode(json: unknown): Activity {
    const schema = ajv.getSchema(ActivityDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ActivityDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ActivityDecoder.definitionName)
  },
}
export const CreateActivityRequestDecoder: Decoder<CreateActivityRequest> = {
  definitionName: 'CreateActivityRequest',
  schemaRef: '#/definitions/CreateActivityRequest',

  decode(json: unknown): CreateActivityRequest {
    const schema = ajv.getSchema(CreateActivityRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${CreateActivityRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, CreateActivityRequestDecoder.definitionName)
  },
}
export const CreateDemoTwoWorkItemRequestDecoder: Decoder<CreateDemoTwoWorkItemRequest> = {
  definitionName: 'CreateDemoTwoWorkItemRequest',
  schemaRef: '#/definitions/CreateDemoTwoWorkItemRequest',

  decode(json: unknown): CreateDemoTwoWorkItemRequest {
    const schema = ajv.getSchema(CreateDemoTwoWorkItemRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${CreateDemoTwoWorkItemRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, CreateDemoTwoWorkItemRequestDecoder.definitionName)
  },
}
export const CreateDemoWorkItemRequestDecoder: Decoder<CreateDemoWorkItemRequest> = {
  definitionName: 'CreateDemoWorkItemRequest',
  schemaRef: '#/definitions/CreateDemoWorkItemRequest',

  decode(json: unknown): CreateDemoWorkItemRequest {
    const schema = ajv.getSchema(CreateDemoWorkItemRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${CreateDemoWorkItemRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, CreateDemoWorkItemRequestDecoder.definitionName)
  },
}
export const CreateProjectBoardRequestDecoder: Decoder<CreateProjectBoardRequest> = {
  definitionName: 'CreateProjectBoardRequest',
  schemaRef: '#/definitions/CreateProjectBoardRequest',

  decode(json: unknown): CreateProjectBoardRequest {
    const schema = ajv.getSchema(CreateProjectBoardRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${CreateProjectBoardRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, CreateProjectBoardRequestDecoder.definitionName)
  },
}
export const CreateTeamRequestDecoder: Decoder<CreateTeamRequest> = {
  definitionName: 'CreateTeamRequest',
  schemaRef: '#/definitions/CreateTeamRequest',

  decode(json: unknown): CreateTeamRequest {
    const schema = ajv.getSchema(CreateTeamRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${CreateTeamRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, CreateTeamRequestDecoder.definitionName)
  },
}
export const CreateWorkItemCommentRequestDecoder: Decoder<CreateWorkItemCommentRequest> = {
  definitionName: 'CreateWorkItemCommentRequest',
  schemaRef: '#/definitions/CreateWorkItemCommentRequest',

  decode(json: unknown): CreateWorkItemCommentRequest {
    const schema = ajv.getSchema(CreateWorkItemCommentRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${CreateWorkItemCommentRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, CreateWorkItemCommentRequestDecoder.definitionName)
  },
}
export const CreateWorkItemTagRequestDecoder: Decoder<CreateWorkItemTagRequest> = {
  definitionName: 'CreateWorkItemTagRequest',
  schemaRef: '#/definitions/CreateWorkItemTagRequest',

  decode(json: unknown): CreateWorkItemTagRequest {
    const schema = ajv.getSchema(CreateWorkItemTagRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${CreateWorkItemTagRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, CreateWorkItemTagRequestDecoder.definitionName)
  },
}
export const CreateWorkItemTypeRequestDecoder: Decoder<CreateWorkItemTypeRequest> = {
  definitionName: 'CreateWorkItemTypeRequest',
  schemaRef: '#/definitions/CreateWorkItemTypeRequest',

  decode(json: unknown): CreateWorkItemTypeRequest {
    const schema = ajv.getSchema(CreateWorkItemTypeRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${CreateWorkItemTypeRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, CreateWorkItemTypeRequestDecoder.definitionName)
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
export const DbUserWIAUWorkItemDecoder: Decoder<DbUserWIAUWorkItem> = {
  definitionName: 'DbUserWIAUWorkItem',
  schemaRef: '#/definitions/DbUserWIAUWorkItem',

  decode(json: unknown): DbUserWIAUWorkItem {
    const schema = ajv.getSchema(DbUserWIAUWorkItemDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbUserWIAUWorkItemDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbUserWIAUWorkItemDecoder.definitionName)
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
export const NotificationDecoder: Decoder<Notification> = {
  definitionName: 'Notification',
  schemaRef: '#/definitions/Notification',

  decode(json: unknown): Notification {
    const schema = ajv.getSchema(NotificationDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${NotificationDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, NotificationDecoder.definitionName)
  },
}
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
export const PaginatedUsersResponseDecoder: Decoder<PaginatedUsersResponse> = {
  definitionName: 'PaginatedUsersResponse',
  schemaRef: '#/definitions/PaginatedUsersResponse',

  decode(json: unknown): PaginatedUsersResponse {
    const schema = ajv.getSchema(PaginatedUsersResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${PaginatedUsersResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, PaginatedUsersResponseDecoder.definitionName)
  },
}
export const PaginationPageDecoder: Decoder<PaginationPage> = {
  definitionName: 'PaginationPage',
  schemaRef: '#/definitions/PaginationPage',

  decode(json: unknown): PaginationPage {
    const schema = ajv.getSchema(PaginationPageDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${PaginationPageDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, PaginationPageDecoder.definitionName)
  },
}
export const ProjectBoardDecoder: Decoder<ProjectBoard> = {
  definitionName: 'ProjectBoard',
  schemaRef: '#/definitions/ProjectBoard',

  decode(json: unknown): ProjectBoard {
    const schema = ajv.getSchema(ProjectBoardDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ProjectBoardDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ProjectBoardDecoder.definitionName)
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
export const SharedWorkItemJoinsDecoder: Decoder<SharedWorkItemJoins> = {
  definitionName: 'SharedWorkItemJoins',
  schemaRef: '#/definitions/SharedWorkItemJoins',

  decode(json: unknown): SharedWorkItemJoins {
    const schema = ajv.getSchema(SharedWorkItemJoinsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${SharedWorkItemJoinsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, SharedWorkItemJoinsDecoder.definitionName)
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
export const UpdateActivityRequestDecoder: Decoder<UpdateActivityRequest> = {
  definitionName: 'UpdateActivityRequest',
  schemaRef: '#/definitions/UpdateActivityRequest',

  decode(json: unknown): UpdateActivityRequest {
    const schema = ajv.getSchema(UpdateActivityRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${UpdateActivityRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, UpdateActivityRequestDecoder.definitionName)
  },
}
export const UpdateTeamRequestDecoder: Decoder<UpdateTeamRequest> = {
  definitionName: 'UpdateTeamRequest',
  schemaRef: '#/definitions/UpdateTeamRequest',

  decode(json: unknown): UpdateTeamRequest {
    const schema = ajv.getSchema(UpdateTeamRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${UpdateTeamRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, UpdateTeamRequestDecoder.definitionName)
  },
}
export const UpdateWorkItemCommentRequestDecoder: Decoder<UpdateWorkItemCommentRequest> = {
  definitionName: 'UpdateWorkItemCommentRequest',
  schemaRef: '#/definitions/UpdateWorkItemCommentRequest',

  decode(json: unknown): UpdateWorkItemCommentRequest {
    const schema = ajv.getSchema(UpdateWorkItemCommentRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${UpdateWorkItemCommentRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, UpdateWorkItemCommentRequestDecoder.definitionName)
  },
}
export const UpdateWorkItemTagRequestDecoder: Decoder<UpdateWorkItemTagRequest> = {
  definitionName: 'UpdateWorkItemTagRequest',
  schemaRef: '#/definitions/UpdateWorkItemTagRequest',

  decode(json: unknown): UpdateWorkItemTagRequest {
    const schema = ajv.getSchema(UpdateWorkItemTagRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${UpdateWorkItemTagRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, UpdateWorkItemTagRequestDecoder.definitionName)
  },
}
export const UpdateWorkItemTypeRequestDecoder: Decoder<UpdateWorkItemTypeRequest> = {
  definitionName: 'UpdateWorkItemTypeRequest',
  schemaRef: '#/definitions/UpdateWorkItemTypeRequest',

  decode(json: unknown): UpdateWorkItemTypeRequest {
    const schema = ajv.getSchema(UpdateWorkItemTypeRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${UpdateWorkItemTypeRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, UpdateWorkItemTypeRequestDecoder.definitionName)
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
export const WorkItemTagDecoder: Decoder<WorkItemTag> = {
  definitionName: 'WorkItemTag',
  schemaRef: '#/definitions/WorkItemTag',

  decode(json: unknown): WorkItemTag {
    const schema = ajv.getSchema(WorkItemTagDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${WorkItemTagDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, WorkItemTagDecoder.definitionName)
  },
}
export const WorkItemTypeDecoder: Decoder<WorkItemType> = {
  definitionName: 'WorkItemType',
  schemaRef: '#/definitions/WorkItemType',

  decode(json: unknown): WorkItemType {
    const schema = ajv.getSchema(WorkItemTypeDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${WorkItemTypeDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, WorkItemTypeDecoder.definitionName)
  },
}
export const DirectionDecoder: Decoder<Direction> = {
  definitionName: 'Direction',
  schemaRef: '#/definitions/Direction',

  decode(json: unknown): Direction {
    const schema = ajv.getSchema(DirectionDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DirectionDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DirectionDecoder.definitionName)
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
export const TopicDecoder: Decoder<Topic> = {
  definitionName: 'Topic',
  schemaRef: '#/definitions/Topic',

  decode(json: unknown): Topic {
    const schema = ajv.getSchema(TopicDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${TopicDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, TopicDecoder.definitionName)
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
export const CreateWorkItemRequestDecoder: Decoder<CreateWorkItemRequest> = {
  definitionName: 'CreateWorkItemRequest',
  schemaRef: '#/definitions/CreateWorkItemRequest',

  decode(json: unknown): CreateWorkItemRequest {
    const schema = ajv.getSchema(CreateWorkItemRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${CreateWorkItemRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, CreateWorkItemRequestDecoder.definitionName)
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
export const DbUserNotificationDecoder: Decoder<DbUserNotification> = {
  definitionName: 'DbUserNotification',
  schemaRef: '#/definitions/DbUserNotification',

  decode(json: unknown): DbUserNotification {
    const schema = ajv.getSchema(DbUserNotificationDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbUserNotificationDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbUserNotificationDecoder.definitionName)
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
export const DbUserWIAWorkItemDecoder: Decoder<DbUserWIAWorkItem> = {
  definitionName: 'DbUserWIAWorkItem',
  schemaRef: '#/definitions/DbUserWIAWorkItem',

  decode(json: unknown): DbUserWIAWorkItem {
    const schema = ajv.getSchema(DbUserWIAWorkItemDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbUserWIAWorkItemDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbUserWIAWorkItemDecoder.definitionName)
  },
}
export const DbWorkItemM2MAssigneeWIADecoder: Decoder<DbWorkItemM2MAssigneeWIA> = {
  definitionName: 'DbWorkItemM2MAssigneeWIA',
  schemaRef: '#/definitions/DbWorkItemM2MAssigneeWIA',

  decode(json: unknown): DbWorkItemM2MAssigneeWIA {
    const schema = ajv.getSchema(DbWorkItemM2MAssigneeWIADecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbWorkItemM2MAssigneeWIADecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbWorkItemM2MAssigneeWIADecoder.definitionName)
  },
}
export const CreateTimeEntryRequestDecoder: Decoder<CreateTimeEntryRequest> = {
  definitionName: 'CreateTimeEntryRequest',
  schemaRef: '#/definitions/CreateTimeEntryRequest',

  decode(json: unknown): CreateTimeEntryRequest {
    const schema = ajv.getSchema(CreateTimeEntryRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${CreateTimeEntryRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, CreateTimeEntryRequestDecoder.definitionName)
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
export const UpdateTimeEntryRequestDecoder: Decoder<UpdateTimeEntryRequest> = {
  definitionName: 'UpdateTimeEntryRequest',
  schemaRef: '#/definitions/UpdateTimeEntryRequest',

  decode(json: unknown): UpdateTimeEntryRequest {
    const schema = ajv.getSchema(UpdateTimeEntryRequestDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${UpdateTimeEntryRequestDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, UpdateTimeEntryRequestDecoder.definitionName)
  },
}
export const DemoTwoWorkItemDecoder: Decoder<DemoTwoWorkItem> = {
  definitionName: 'DemoTwoWorkItem',
  schemaRef: '#/definitions/DemoTwoWorkItem',

  decode(json: unknown): DemoTwoWorkItem {
    const schema = ajv.getSchema(DemoTwoWorkItemDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DemoTwoWorkItemDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DemoTwoWorkItemDecoder.definitionName)
  },
}
export const DemoWorkItemDecoder: Decoder<DemoWorkItem> = {
  definitionName: 'DemoWorkItem',
  schemaRef: '#/definitions/DemoWorkItem',

  decode(json: unknown): DemoWorkItem {
    const schema = ajv.getSchema(DemoWorkItemDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DemoWorkItemDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DemoWorkItemDecoder.definitionName)
  },
}
export const WorkItemBaseDecoder: Decoder<WorkItemBase> = {
  definitionName: 'WorkItemBase',
  schemaRef: '#/definitions/WorkItemBase',

  decode(json: unknown): WorkItemBase {
    const schema = ajv.getSchema(WorkItemBaseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${WorkItemBaseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, WorkItemBaseDecoder.definitionName)
  },
}
export const PaginationFilterModesDecoder: Decoder<PaginationFilterModes> = {
  definitionName: 'PaginationFilterModes',
  schemaRef: '#/definitions/PaginationFilterModes',

  decode(json: unknown): PaginationFilterModes {
    const schema = ajv.getSchema(PaginationFilterModesDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${PaginationFilterModesDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, PaginationFilterModesDecoder.definitionName)
  },
}
export const DbCacheDemoWorkItemJoinsDecoder: Decoder<DbCacheDemoWorkItemJoins> = {
  definitionName: 'DbCacheDemoWorkItemJoins',
  schemaRef: '#/definitions/DbCacheDemoWorkItemJoins',

  decode(json: unknown): DbCacheDemoWorkItemJoins {
    const schema = ajv.getSchema(DbCacheDemoWorkItemJoinsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbCacheDemoWorkItemJoinsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbCacheDemoWorkItemJoinsDecoder.definitionName)
  },
}
export const DbUserJoinsDecoder: Decoder<DbUserJoins> = {
  definitionName: 'DbUserJoins',
  schemaRef: '#/definitions/DbUserJoins',

  decode(json: unknown): DbUserJoins {
    const schema = ajv.getSchema(DbUserJoinsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DbUserJoinsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DbUserJoinsDecoder.definitionName)
  },
}
export const GetCacheDemoWorkItemQueryParametersDecoder: Decoder<GetCacheDemoWorkItemQueryParameters> = {
  definitionName: 'GetCacheDemoWorkItemQueryParameters',
  schemaRef: '#/definitions/GetCacheDemoWorkItemQueryParameters',

  decode(json: unknown): GetCacheDemoWorkItemQueryParameters {
    const schema = ajv.getSchema(GetCacheDemoWorkItemQueryParametersDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${GetCacheDemoWorkItemQueryParametersDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, GetCacheDemoWorkItemQueryParametersDecoder.definitionName)
  },
}
export const GetPaginatedUsersQueryParametersDecoder: Decoder<GetPaginatedUsersQueryParameters> = {
  definitionName: 'GetPaginatedUsersQueryParameters',
  schemaRef: '#/definitions/GetPaginatedUsersQueryParameters',

  decode(json: unknown): GetPaginatedUsersQueryParameters {
    const schema = ajv.getSchema(GetPaginatedUsersQueryParametersDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${GetPaginatedUsersQueryParametersDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, GetPaginatedUsersQueryParametersDecoder.definitionName)
  },
}
export const GetCurrentUserQueryParametersDecoder: Decoder<GetCurrentUserQueryParameters> = {
  definitionName: 'GetCurrentUserQueryParameters',
  schemaRef: '#/definitions/GetCurrentUserQueryParameters',

  decode(json: unknown): GetCurrentUserQueryParameters {
    const schema = ajv.getSchema(GetCurrentUserQueryParametersDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${GetCurrentUserQueryParametersDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, GetCurrentUserQueryParametersDecoder.definitionName)
  },
}
