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
  DbWorkItem,
  WorkItemTagCreateRequest,
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
  DemoWorkItemsResponse: info<DemoWorkItemsResponse>('DemoWorkItemsResponse', '#/definitions/DemoWorkItemsResponse'),
  DemoTwoWorkItemsResponse: info<DemoTwoWorkItemsResponse>(
    'DemoTwoWorkItemsResponse',
    '#/definitions/DemoTwoWorkItemsResponse',
  ),
  InitializeProjectRequest: info<InitializeProjectRequest>(
    'InitializeProjectRequest',
    '#/definitions/InitializeProjectRequest',
  ),
  ProjectBoardResponse: info<ProjectBoardResponse>('ProjectBoardResponse', '#/definitions/ProjectBoardResponse'),
  User: info<User>('User', '#/definitions/User'),
  HTTPValidationError: info<HTTPValidationError>('HTTPValidationError', '#/definitions/HTTPValidationError'),
  ErrorCode: info<ErrorCode>('ErrorCode', '#/definitions/ErrorCode'),
  HTTPError: info<HTTPError>('HTTPError', '#/definitions/HTTPError'),
  Topics: info<Topics>('Topics', '#/definitions/Topics'),
  Scope: info<Scope>('Scope', '#/definitions/Scope'),
  Scopes: info<Scopes>('Scopes', '#/definitions/Scopes'),
  Role: info<Role>('Role', '#/definitions/Role'),
  WorkItemRole: info<WorkItemRole>('WorkItemRole', '#/definitions/WorkItemRole'),
  UpdateUserRequest: info<UpdateUserRequest>('UpdateUserRequest', '#/definitions/UpdateUserRequest'),
  UpdateUserAuthRequest: info<UpdateUserAuthRequest>('UpdateUserAuthRequest', '#/definitions/UpdateUserAuthRequest'),
  ValidationError: info<ValidationError>('ValidationError', '#/definitions/ValidationError'),
  UuidUUID: info<UuidUUID>('UuidUUID', '#/definitions/UuidUUID'),
  DbWorkItem: info<DbWorkItem>('DbWorkItem', '#/definitions/DbWorkItem'),
  WorkItemTagCreateRequest: info<WorkItemTagCreateRequest>(
    'WorkItemTagCreateRequest',
    '#/definitions/WorkItemTagCreateRequest',
  ),
  DemoWorkItemCreateRequest: info<DemoWorkItemCreateRequest>(
    'DemoWorkItemCreateRequest',
    '#/definitions/DemoWorkItemCreateRequest',
  ),
  WorkItemCommentCreateRequest: info<WorkItemCommentCreateRequest>(
    'WorkItemCommentCreateRequest',
    '#/definitions/WorkItemCommentCreateRequest',
  ),
  Project: info<Project>('Project', '#/definitions/Project'),
  DbActivityCreateParams: info<DbActivityCreateParams>(
    'DbActivityCreateParams',
    '#/definitions/DbActivityCreateParams',
  ),
  DbTeamCreateParams: info<DbTeamCreateParams>('DbTeamCreateParams', '#/definitions/DbTeamCreateParams'),
  DbWorkItemTagCreateParams: info<DbWorkItemTagCreateParams>(
    'DbWorkItemTagCreateParams',
    '#/definitions/DbWorkItemTagCreateParams',
  ),
  DbWorkItemRole: info<DbWorkItemRole>('DbWorkItemRole', '#/definitions/DbWorkItemRole'),
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
  ServicesMember: info<ServicesMember>('ServicesMember', '#/definitions/ServicesMember'),
  DbDemoTwoWorkItem: info<DbDemoTwoWorkItem>('DbDemoTwoWorkItem', '#/definitions/DbDemoTwoWorkItem'),
}

export interface SchemaInfo<T> {
  definitionName: string
  schemaRef: string
}

function info<T>(definitionName: string, schemaRef: string): SchemaInfo<T> {
  return { definitionName, schemaRef }
}
