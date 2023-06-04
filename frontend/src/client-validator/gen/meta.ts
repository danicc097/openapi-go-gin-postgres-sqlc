/* eslint-disable */
// @ts-nocheck
/* eslint-disable */
/* eslint-disable @typescript-eslint/ban-ts-comment */

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

export const schemaDefinitions = {
  DbActivity: info<DbActivity>('DbActivity', '#/definitions/DbActivity'),
  DbKanbanStep: info<DbKanbanStep>('DbKanbanStep', '#/definitions/DbKanbanStep'),
  DbProject: info<DbProject>('DbProject', '#/definitions/DbProject'),
  DbTeam: info<DbTeam>('DbTeam', '#/definitions/DbTeam'),
  DbWorkItemTag: info<DbWorkItemTag>('DbWorkItemTag', '#/definitions/DbWorkItemTag'),
  DbWorkItemType: info<DbWorkItemType>('DbWorkItemType', '#/definitions/DbWorkItemType'),
  DbDemoWorkItem: info<DbDemoWorkItem>('DbDemoWorkItem', '#/definitions/DbDemoWorkItem'),
  DbUserAPIKey: info<DbUserAPIKey>('DbUserAPIKey', '#/definitions/DbUserAPIKey'),
  DbUser: info<DbUser>('DbUser', '#/definitions/DbUser'),
  DbTimeEntry: info<DbTimeEntry>('DbTimeEntry', '#/definitions/DbTimeEntry'),
  DbWorkItemComment: info<DbWorkItemComment>('DbWorkItemComment', '#/definitions/DbWorkItemComment'),
  ProjectConfig: info<ProjectConfig>('ProjectConfig', '#/definitions/ProjectConfig'),
  ProjectConfigField: info<ProjectConfigField>('ProjectConfigField', '#/definitions/ProjectConfigField'),
  RestDemoWorkItemsResponse: info<RestDemoWorkItemsResponse>(
    'RestDemoWorkItemsResponse',
    '#/definitions/RestDemoWorkItemsResponse',
  ),
  InitializeProjectRequest: info<InitializeProjectRequest>(
    'InitializeProjectRequest',
    '#/definitions/InitializeProjectRequest',
  ),
  RestProjectBoardResponse: info<RestProjectBoardResponse>(
    'RestProjectBoardResponse',
    '#/definitions/RestProjectBoardResponse',
  ),
  UserResponse: info<UserResponse>('UserResponse', '#/definitions/UserResponse'),
  HTTPValidationError: info<HTTPValidationError>('HTTPValidationError', '#/definitions/HTTPValidationError'),
  Topics: info<Topics>('Topics', '#/definitions/Topics'),
  Scope: info<Scope>('Scope', '#/definitions/Scope'),
  Scopes: info<Scopes>('Scopes', '#/definitions/Scopes'),
  Role: info<Role>('Role', '#/definitions/Role'),
  WorkItemRole: info<WorkItemRole>('WorkItemRole', '#/definitions/WorkItemRole'),
  UpdateUserRequest: info<UpdateUserRequest>('UpdateUserRequest', '#/definitions/UpdateUserRequest'),
  UpdateUserAuthRequest: info<UpdateUserAuthRequest>('UpdateUserAuthRequest', '#/definitions/UpdateUserAuthRequest'),
  ValidationError: info<ValidationError>('ValidationError', '#/definitions/ValidationError'),
  HttpErrorType: info<HttpErrorType>('HttpErrorType', '#/definitions/HttpErrorType'),
  UuidUUID: info<UuidUUID>('UuidUUID', '#/definitions/UuidUUID'),
  PgtypeJSONB: info<PgtypeJSONB>('PgtypeJSONB', '#/definitions/PgtypeJSONB'),
  DbWorkItem: info<DbWorkItem>('DbWorkItem', '#/definitions/DbWorkItem'),
  RestWorkItemTagCreateRequest: info<RestWorkItemTagCreateRequest>(
    'RestWorkItemTagCreateRequest',
    '#/definitions/RestWorkItemTagCreateRequest',
  ),
  RestDemoWorkItemCreateRequest: info<RestDemoWorkItemCreateRequest>(
    'RestDemoWorkItemCreateRequest',
    '#/definitions/RestDemoWorkItemCreateRequest',
  ),
  RestWorkItemCommentCreateRequest: info<RestWorkItemCommentCreateRequest>(
    'RestWorkItemCommentCreateRequest',
    '#/definitions/RestWorkItemCommentCreateRequest',
  ),
  Project: info<Project>('Project', '#/definitions/Project'),
  DbActivityCreateParams: info<DbActivityCreateParams>(
    'DbActivityCreateParams',
    '#/definitions/DbActivityCreateParams',
  ),
  DbKanbanStepCreateParams: info<DbKanbanStepCreateParams>(
    'DbKanbanStepCreateParams',
    '#/definitions/DbKanbanStepCreateParams',
  ),
  DbTeamCreateParams: info<DbTeamCreateParams>('DbTeamCreateParams', '#/definitions/DbTeamCreateParams'),
  DbWorkItemTagCreateParams: info<DbWorkItemTagCreateParams>(
    'DbWorkItemTagCreateParams',
    '#/definitions/DbWorkItemTagCreateParams',
  ),
  DbWorkItemTypeCreateParams: info<DbWorkItemTypeCreateParams>(
    'DbWorkItemTypeCreateParams',
    '#/definitions/DbWorkItemTypeCreateParams',
  ),
  DbWorkItemRole: info<DbWorkItemRole>('DbWorkItemRole', '#/definitions/DbWorkItemRole'),
  DbWorkItem_AssignedUser: info<DbWorkItem_AssignedUser>(
    'DbWorkItem_AssignedUser',
    '#/definitions/DbWorkItem_AssignedUser',
  ),
  NotificationType: info<NotificationType>('NotificationType', '#/definitions/NotificationType'),
  DemoProjectKanbanSteps: info<DemoProjectKanbanSteps>(
    'DemoProjectKanbanSteps',
    '#/definitions/DemoProjectKanbanSteps',
  ),
  DemoProject2KanbanSteps: info<DemoProject2KanbanSteps>(
    'DemoProject2KanbanSteps',
    '#/definitions/DemoProject2KanbanSteps',
  ),
  Demo2WorkItemTypes: info<Demo2WorkItemTypes>('Demo2WorkItemTypes', '#/definitions/Demo2WorkItemTypes'),
  DemoKanbanSteps: info<DemoKanbanSteps>('DemoKanbanSteps', '#/definitions/DemoKanbanSteps'),
  DemoTwoKanbanSteps: info<DemoTwoKanbanSteps>('DemoTwoKanbanSteps', '#/definitions/DemoTwoKanbanSteps'),
  DemoTwoWorkItemTypes: info<DemoTwoWorkItemTypes>('DemoTwoWorkItemTypes', '#/definitions/DemoTwoWorkItemTypes'),
  DemoWorkItemTypes: info<DemoWorkItemTypes>('DemoWorkItemTypes', '#/definitions/DemoWorkItemTypes'),
  DbDemoWorkItemCreateParams: info<DbDemoWorkItemCreateParams>(
    'DbDemoWorkItemCreateParams',
    '#/definitions/DbDemoWorkItemCreateParams',
  ),
  DbWorkItemCreateParams: info<DbWorkItemCreateParams>(
    'DbWorkItemCreateParams',
    '#/definitions/DbWorkItemCreateParams',
  ),
  ModelsWorkItemRole: info<ModelsWorkItemRole>('ModelsWorkItemRole', '#/definitions/ModelsWorkItemRole'),
  ServicesMember: info<ServicesMember>('ServicesMember', '#/definitions/ServicesMember'),
}

export interface SchemaInfo<T> {
  definitionName: string
  schemaRef: string
}

function info<T>(definitionName: string, schemaRef: string): SchemaInfo<T> {
  return { definitionName, schemaRef }
}
