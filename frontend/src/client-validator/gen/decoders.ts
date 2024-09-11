/* eslint-disable */
// @ts-nocheck
/* eslint-disable */
/* eslint-disable @typescript-eslint/ban-ts-comment */

import Ajv from 'ajv'
import addFormats from 'ajv-formats'
import { Decoder } from './helpers'
import { validateJson } from '../validate'
import {
  CreateActivityRequest,
  CreateDemoTwoWorkItemRequest,
  CreateDemoWorkItemRequest,
  CreateProjectBoardRequest,
  CreateTeamRequest,
  CreateWorkItemCommentRequest,
  CreateWorkItemTagRequest,
  CreateWorkItemTypeRequest,
  ModelsDemoTwoWorkItem,
  ModelsDemoTwoWorkItemCreateParams,
  ModelsDemoWorkItem,
  ModelsDemoWorkItemCreateParams,
  ModelsKanbanStep,
  ModelsNotification,
  ModelsProject,
  ModelsTeam,
  ModelsTeamCreateParams,
  ModelsTimeEntry,
  ModelsUser,
  ModelsUserAPIKey,
  ModelsUserID,
  ModelsWorkItem,
  ModelsWorkItemComment,
  ModelsWorkItemCreateParams,
  ModelsWorkItemTag,
  ModelsWorkItemTagCreateParams,
  ModelsWorkItemType,
  PaginatedNotificationsResponse,
  PaginatedUsersResponse,
  PaginationPage,
  ProjectBoard,
  ServicesMember,
  SharedWorkItemJoins,
  UpdateActivityRequest,
  UpdateTeamRequest,
  UpdateWorkItemCommentRequest,
  UpdateWorkItemTagRequest,
  UpdateWorkItemTypeRequest,
  Direction,
  ProjectConfig,
  ProjectConfigField,
  HTTPValidationError,
  ErrorCode,
  HTTPError,
  Topics,
  Topic,
  Scope,
  Role,
  WorkItemRole,
  UpdateUserRequest,
  UpdateUserAuthRequest,
  ValidationError,
  UuidUUID,
  WorkItemResponse,
  Scopes,
  CreateWorkItemRequest,
  NotificationType,
  DemoTwoWorkItemTypes,
  DemoWorkItemTypes,
  DemoKanbanSteps,
  DemoTwoKanbanSteps,
  ModelsWorkItemM2MAssigneeWIA,
  CreateTimeEntryRequest,
  UpdateTimeEntryRequest,
  DemoTwoWorkItemResponse,
  DemoWorkItemResponse,
  WorkItemBase,
  PaginationFilterPrimitive,
  PaginationFilterArray,
  PaginationFilter,
  Pagination,
  PaginationItems,
  PaginationCursor,
  GetPaginatedUsersQueryParameters,
  PaginationFilterModes,
  ModelsCacheDemoWorkItemJoins,
  ModelsUserJoins,
  PaginatedDemoWorkItemsResponse,
  GetCacheDemoWorkItemQueryParameters,
  GetCurrentUserQueryParameters,
  ProjectName,
  ActivityResponse,
  TeamResponse,
  UserResponse,
  WorkItemCommentResponse,
  WorkItemTagResponse,
  TimeEntryResponse,
  CacheDemoWorkItemResponse,
  NotificationResponse,
  WorkItemTypeResponse,
  ModelsProjectConfig,
  ModelsProjectConfigField,
} from './models'
import jsonSchema from './schema.json'

const ajv = new Ajv({ strict: false, allErrors: true })
addFormats(ajv, { formats: ['int64', 'int32', 'binary', 'date-time', 'date'] })
ajv.compile(jsonSchema)

// Decoders
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
export const ModelsDemoTwoWorkItemDecoder: Decoder<ModelsDemoTwoWorkItem> = {
  definitionName: 'ModelsDemoTwoWorkItem',
  schemaRef: '#/definitions/ModelsDemoTwoWorkItem',

  decode(json: unknown): ModelsDemoTwoWorkItem {
    const schema = ajv.getSchema(ModelsDemoTwoWorkItemDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsDemoTwoWorkItemDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsDemoTwoWorkItemDecoder.definitionName)
  },
}
export const ModelsDemoTwoWorkItemCreateParamsDecoder: Decoder<ModelsDemoTwoWorkItemCreateParams> = {
  definitionName: 'ModelsDemoTwoWorkItemCreateParams',
  schemaRef: '#/definitions/ModelsDemoTwoWorkItemCreateParams',

  decode(json: unknown): ModelsDemoTwoWorkItemCreateParams {
    const schema = ajv.getSchema(ModelsDemoTwoWorkItemCreateParamsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsDemoTwoWorkItemCreateParamsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsDemoTwoWorkItemCreateParamsDecoder.definitionName)
  },
}
export const ModelsDemoWorkItemDecoder: Decoder<ModelsDemoWorkItem> = {
  definitionName: 'ModelsDemoWorkItem',
  schemaRef: '#/definitions/ModelsDemoWorkItem',

  decode(json: unknown): ModelsDemoWorkItem {
    const schema = ajv.getSchema(ModelsDemoWorkItemDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsDemoWorkItemDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsDemoWorkItemDecoder.definitionName)
  },
}
export const ModelsDemoWorkItemCreateParamsDecoder: Decoder<ModelsDemoWorkItemCreateParams> = {
  definitionName: 'ModelsDemoWorkItemCreateParams',
  schemaRef: '#/definitions/ModelsDemoWorkItemCreateParams',

  decode(json: unknown): ModelsDemoWorkItemCreateParams {
    const schema = ajv.getSchema(ModelsDemoWorkItemCreateParamsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsDemoWorkItemCreateParamsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsDemoWorkItemCreateParamsDecoder.definitionName)
  },
}
export const ModelsKanbanStepDecoder: Decoder<ModelsKanbanStep> = {
  definitionName: 'ModelsKanbanStep',
  schemaRef: '#/definitions/ModelsKanbanStep',

  decode(json: unknown): ModelsKanbanStep {
    const schema = ajv.getSchema(ModelsKanbanStepDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsKanbanStepDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsKanbanStepDecoder.definitionName)
  },
}
export const ModelsNotificationDecoder: Decoder<ModelsNotification> = {
  definitionName: 'ModelsNotification',
  schemaRef: '#/definitions/ModelsNotification',

  decode(json: unknown): ModelsNotification {
    const schema = ajv.getSchema(ModelsNotificationDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsNotificationDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsNotificationDecoder.definitionName)
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
export const ModelsTeamDecoder: Decoder<ModelsTeam> = {
  definitionName: 'ModelsTeam',
  schemaRef: '#/definitions/ModelsTeam',

  decode(json: unknown): ModelsTeam {
    const schema = ajv.getSchema(ModelsTeamDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsTeamDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsTeamDecoder.definitionName)
  },
}
export const ModelsTeamCreateParamsDecoder: Decoder<ModelsTeamCreateParams> = {
  definitionName: 'ModelsTeamCreateParams',
  schemaRef: '#/definitions/ModelsTeamCreateParams',

  decode(json: unknown): ModelsTeamCreateParams {
    const schema = ajv.getSchema(ModelsTeamCreateParamsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsTeamCreateParamsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsTeamCreateParamsDecoder.definitionName)
  },
}
export const ModelsTimeEntryDecoder: Decoder<ModelsTimeEntry> = {
  definitionName: 'ModelsTimeEntry',
  schemaRef: '#/definitions/ModelsTimeEntry',

  decode(json: unknown): ModelsTimeEntry {
    const schema = ajv.getSchema(ModelsTimeEntryDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsTimeEntryDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsTimeEntryDecoder.definitionName)
  },
}
export const ModelsUserDecoder: Decoder<ModelsUser> = {
  definitionName: 'ModelsUser',
  schemaRef: '#/definitions/ModelsUser',

  decode(json: unknown): ModelsUser {
    const schema = ajv.getSchema(ModelsUserDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsUserDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsUserDecoder.definitionName)
  },
}
export const ModelsUserAPIKeyDecoder: Decoder<ModelsUserAPIKey> = {
  definitionName: 'ModelsUserAPIKey',
  schemaRef: '#/definitions/ModelsUserAPIKey',

  decode(json: unknown): ModelsUserAPIKey {
    const schema = ajv.getSchema(ModelsUserAPIKeyDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsUserAPIKeyDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsUserAPIKeyDecoder.definitionName)
  },
}
export const ModelsUserIDDecoder: Decoder<ModelsUserID> = {
  definitionName: 'ModelsUserID',
  schemaRef: '#/definitions/ModelsUserID',

  decode(json: unknown): ModelsUserID {
    const schema = ajv.getSchema(ModelsUserIDDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsUserIDDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsUserIDDecoder.definitionName)
  },
}
export const ModelsWorkItemDecoder: Decoder<ModelsWorkItem> = {
  definitionName: 'ModelsWorkItem',
  schemaRef: '#/definitions/ModelsWorkItem',

  decode(json: unknown): ModelsWorkItem {
    const schema = ajv.getSchema(ModelsWorkItemDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsWorkItemDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsWorkItemDecoder.definitionName)
  },
}
export const ModelsWorkItemCommentDecoder: Decoder<ModelsWorkItemComment> = {
  definitionName: 'ModelsWorkItemComment',
  schemaRef: '#/definitions/ModelsWorkItemComment',

  decode(json: unknown): ModelsWorkItemComment {
    const schema = ajv.getSchema(ModelsWorkItemCommentDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsWorkItemCommentDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsWorkItemCommentDecoder.definitionName)
  },
}
export const ModelsWorkItemCreateParamsDecoder: Decoder<ModelsWorkItemCreateParams> = {
  definitionName: 'ModelsWorkItemCreateParams',
  schemaRef: '#/definitions/ModelsWorkItemCreateParams',

  decode(json: unknown): ModelsWorkItemCreateParams {
    const schema = ajv.getSchema(ModelsWorkItemCreateParamsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsWorkItemCreateParamsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsWorkItemCreateParamsDecoder.definitionName)
  },
}
export const ModelsWorkItemTagDecoder: Decoder<ModelsWorkItemTag> = {
  definitionName: 'ModelsWorkItemTag',
  schemaRef: '#/definitions/ModelsWorkItemTag',

  decode(json: unknown): ModelsWorkItemTag {
    const schema = ajv.getSchema(ModelsWorkItemTagDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsWorkItemTagDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsWorkItemTagDecoder.definitionName)
  },
}
export const ModelsWorkItemTagCreateParamsDecoder: Decoder<ModelsWorkItemTagCreateParams> = {
  definitionName: 'ModelsWorkItemTagCreateParams',
  schemaRef: '#/definitions/ModelsWorkItemTagCreateParams',

  decode(json: unknown): ModelsWorkItemTagCreateParams {
    const schema = ajv.getSchema(ModelsWorkItemTagCreateParamsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsWorkItemTagCreateParamsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsWorkItemTagCreateParamsDecoder.definitionName)
  },
}
export const ModelsWorkItemTypeDecoder: Decoder<ModelsWorkItemType> = {
  definitionName: 'ModelsWorkItemType',
  schemaRef: '#/definitions/ModelsWorkItemType',

  decode(json: unknown): ModelsWorkItemType {
    const schema = ajv.getSchema(ModelsWorkItemTypeDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsWorkItemTypeDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsWorkItemTypeDecoder.definitionName)
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
export const WorkItemResponseDecoder: Decoder<WorkItemResponse> = {
  definitionName: 'WorkItemResponse',
  schemaRef: '#/definitions/WorkItemResponse',

  decode(json: unknown): WorkItemResponse {
    const schema = ajv.getSchema(WorkItemResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${WorkItemResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, WorkItemResponseDecoder.definitionName)
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
export const ModelsWorkItemM2MAssigneeWIADecoder: Decoder<ModelsWorkItemM2MAssigneeWIA> = {
  definitionName: 'ModelsWorkItemM2MAssigneeWIA',
  schemaRef: '#/definitions/ModelsWorkItemM2MAssigneeWIA',

  decode(json: unknown): ModelsWorkItemM2MAssigneeWIA {
    const schema = ajv.getSchema(ModelsWorkItemM2MAssigneeWIADecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsWorkItemM2MAssigneeWIADecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsWorkItemM2MAssigneeWIADecoder.definitionName)
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
export const DemoTwoWorkItemResponseDecoder: Decoder<DemoTwoWorkItemResponse> = {
  definitionName: 'DemoTwoWorkItemResponse',
  schemaRef: '#/definitions/DemoTwoWorkItemResponse',

  decode(json: unknown): DemoTwoWorkItemResponse {
    const schema = ajv.getSchema(DemoTwoWorkItemResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DemoTwoWorkItemResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DemoTwoWorkItemResponseDecoder.definitionName)
  },
}
export const DemoWorkItemResponseDecoder: Decoder<DemoWorkItemResponse> = {
  definitionName: 'DemoWorkItemResponse',
  schemaRef: '#/definitions/DemoWorkItemResponse',

  decode(json: unknown): DemoWorkItemResponse {
    const schema = ajv.getSchema(DemoWorkItemResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${DemoWorkItemResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, DemoWorkItemResponseDecoder.definitionName)
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
export const PaginationFilterPrimitiveDecoder: Decoder<PaginationFilterPrimitive> = {
  definitionName: 'PaginationFilterPrimitive',
  schemaRef: '#/definitions/PaginationFilterPrimitive',

  decode(json: unknown): PaginationFilterPrimitive {
    const schema = ajv.getSchema(PaginationFilterPrimitiveDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${PaginationFilterPrimitiveDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, PaginationFilterPrimitiveDecoder.definitionName)
  },
}
export const PaginationFilterArrayDecoder: Decoder<PaginationFilterArray> = {
  definitionName: 'PaginationFilterArray',
  schemaRef: '#/definitions/PaginationFilterArray',

  decode(json: unknown): PaginationFilterArray {
    const schema = ajv.getSchema(PaginationFilterArrayDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${PaginationFilterArrayDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, PaginationFilterArrayDecoder.definitionName)
  },
}
export const PaginationFilterDecoder: Decoder<PaginationFilter> = {
  definitionName: 'PaginationFilter',
  schemaRef: '#/definitions/PaginationFilter',

  decode(json: unknown): PaginationFilter {
    const schema = ajv.getSchema(PaginationFilterDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${PaginationFilterDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, PaginationFilterDecoder.definitionName)
  },
}
export const PaginationDecoder: Decoder<Pagination> = {
  definitionName: 'Pagination',
  schemaRef: '#/definitions/Pagination',

  decode(json: unknown): Pagination {
    const schema = ajv.getSchema(PaginationDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${PaginationDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, PaginationDecoder.definitionName)
  },
}
export const PaginationItemsDecoder: Decoder<PaginationItems> = {
  definitionName: 'PaginationItems',
  schemaRef: '#/definitions/PaginationItems',

  decode(json: unknown): PaginationItems {
    const schema = ajv.getSchema(PaginationItemsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${PaginationItemsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, PaginationItemsDecoder.definitionName)
  },
}
export const PaginationCursorDecoder: Decoder<PaginationCursor> = {
  definitionName: 'PaginationCursor',
  schemaRef: '#/definitions/PaginationCursor',

  decode(json: unknown): PaginationCursor {
    const schema = ajv.getSchema(PaginationCursorDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${PaginationCursorDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, PaginationCursorDecoder.definitionName)
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
export const ModelsCacheDemoWorkItemJoinsDecoder: Decoder<ModelsCacheDemoWorkItemJoins> = {
  definitionName: 'ModelsCacheDemoWorkItemJoins',
  schemaRef: '#/definitions/ModelsCacheDemoWorkItemJoins',

  decode(json: unknown): ModelsCacheDemoWorkItemJoins {
    const schema = ajv.getSchema(ModelsCacheDemoWorkItemJoinsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsCacheDemoWorkItemJoinsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsCacheDemoWorkItemJoinsDecoder.definitionName)
  },
}
export const ModelsUserJoinsDecoder: Decoder<ModelsUserJoins> = {
  definitionName: 'ModelsUserJoins',
  schemaRef: '#/definitions/ModelsUserJoins',

  decode(json: unknown): ModelsUserJoins {
    const schema = ajv.getSchema(ModelsUserJoinsDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsUserJoinsDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsUserJoinsDecoder.definitionName)
  },
}
export const PaginatedDemoWorkItemsResponseDecoder: Decoder<PaginatedDemoWorkItemsResponse> = {
  definitionName: 'PaginatedDemoWorkItemsResponse',
  schemaRef: '#/definitions/PaginatedDemoWorkItemsResponse',

  decode(json: unknown): PaginatedDemoWorkItemsResponse {
    const schema = ajv.getSchema(PaginatedDemoWorkItemsResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${PaginatedDemoWorkItemsResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, PaginatedDemoWorkItemsResponseDecoder.definitionName)
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
export const ProjectNameDecoder: Decoder<ProjectName> = {
  definitionName: 'ProjectName',
  schemaRef: '#/definitions/ProjectName',

  decode(json: unknown): ProjectName {
    const schema = ajv.getSchema(ProjectNameDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ProjectNameDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ProjectNameDecoder.definitionName)
  },
}
export const ActivityResponseDecoder: Decoder<ActivityResponse> = {
  definitionName: 'ActivityResponse',
  schemaRef: '#/definitions/ActivityResponse',

  decode(json: unknown): ActivityResponse {
    const schema = ajv.getSchema(ActivityResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ActivityResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ActivityResponseDecoder.definitionName)
  },
}
export const TeamResponseDecoder: Decoder<TeamResponse> = {
  definitionName: 'TeamResponse',
  schemaRef: '#/definitions/TeamResponse',

  decode(json: unknown): TeamResponse {
    const schema = ajv.getSchema(TeamResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${TeamResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, TeamResponseDecoder.definitionName)
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
export const WorkItemCommentResponseDecoder: Decoder<WorkItemCommentResponse> = {
  definitionName: 'WorkItemCommentResponse',
  schemaRef: '#/definitions/WorkItemCommentResponse',

  decode(json: unknown): WorkItemCommentResponse {
    const schema = ajv.getSchema(WorkItemCommentResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${WorkItemCommentResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, WorkItemCommentResponseDecoder.definitionName)
  },
}
export const WorkItemTagResponseDecoder: Decoder<WorkItemTagResponse> = {
  definitionName: 'WorkItemTagResponse',
  schemaRef: '#/definitions/WorkItemTagResponse',

  decode(json: unknown): WorkItemTagResponse {
    const schema = ajv.getSchema(WorkItemTagResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${WorkItemTagResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, WorkItemTagResponseDecoder.definitionName)
  },
}
export const TimeEntryResponseDecoder: Decoder<TimeEntryResponse> = {
  definitionName: 'TimeEntryResponse',
  schemaRef: '#/definitions/TimeEntryResponse',

  decode(json: unknown): TimeEntryResponse {
    const schema = ajv.getSchema(TimeEntryResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${TimeEntryResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, TimeEntryResponseDecoder.definitionName)
  },
}
export const CacheDemoWorkItemResponseDecoder: Decoder<CacheDemoWorkItemResponse> = {
  definitionName: 'CacheDemoWorkItemResponse',
  schemaRef: '#/definitions/CacheDemoWorkItemResponse',

  decode(json: unknown): CacheDemoWorkItemResponse {
    const schema = ajv.getSchema(CacheDemoWorkItemResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${CacheDemoWorkItemResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, CacheDemoWorkItemResponseDecoder.definitionName)
  },
}
export const NotificationResponseDecoder: Decoder<NotificationResponse> = {
  definitionName: 'NotificationResponse',
  schemaRef: '#/definitions/NotificationResponse',

  decode(json: unknown): NotificationResponse {
    const schema = ajv.getSchema(NotificationResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${NotificationResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, NotificationResponseDecoder.definitionName)
  },
}
export const WorkItemTypeResponseDecoder: Decoder<WorkItemTypeResponse> = {
  definitionName: 'WorkItemTypeResponse',
  schemaRef: '#/definitions/WorkItemTypeResponse',

  decode(json: unknown): WorkItemTypeResponse {
    const schema = ajv.getSchema(WorkItemTypeResponseDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${WorkItemTypeResponseDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, WorkItemTypeResponseDecoder.definitionName)
  },
}
export const ModelsProjectConfigDecoder: Decoder<ModelsProjectConfig> = {
  definitionName: 'ModelsProjectConfig',
  schemaRef: '#/definitions/ModelsProjectConfig',

  decode(json: unknown): ModelsProjectConfig {
    const schema = ajv.getSchema(ModelsProjectConfigDecoder.schemaRef)
    if (!schema) {
      throw new Error(`Schema ${ModelsProjectConfigDecoder.definitionName} not found`)
    }
    return validateJson(json, schema, ModelsProjectConfigDecoder.definitionName)
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
