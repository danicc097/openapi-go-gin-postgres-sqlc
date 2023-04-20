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
  NotificationType,
  WorkItemRole,
  UpdateUserRequest,
  UpdateUserAuthRequest,
  ValidationError,
  UuidUUID,
  PgtypeJSONB,
  demoProjectKanbanSteps,
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
} from './models'

export const schemaDefinitions = {
  DbActivity: info<DbActivity>('DbActivity', '#/definitions/DbActivity'),
  DbKanbanStep: info<DbKanbanStep>('DbKanbanStep', '#/definitions/DbKanbanStep'),
  DbProject: info<DbProject>('DbProject', '#/definitions/DbProject'),
  DbTeam: info<DbTeam>('DbTeam', '#/definitions/DbTeam'),
  DbWorkItemTag: info<DbWorkItemTag>('DbWorkItemTag', '#/definitions/DbWorkItemTag'),
  DbWorkItemType: info<DbWorkItemType>('DbWorkItemType', '#/definitions/DbWorkItemType'),
  DbDemoProjectWorkItem: info<DbDemoProjectWorkItem>('DbDemoProjectWorkItem', '#/definitions/DbDemoProjectWorkItem'),
  DbUserAPIKey: info<DbUserAPIKey>('DbUserAPIKey', '#/definitions/DbUserAPIKey'),
  DbUser: info<DbUser>('DbUser', '#/definitions/DbUser'),
  DbTimeEntry: info<DbTimeEntry>('DbTimeEntry', '#/definitions/DbTimeEntry'),
  DbWorkItemComment: info<DbWorkItemComment>('DbWorkItemComment', '#/definitions/DbWorkItemComment'),
  ProjectConfig: info<ProjectConfig>('ProjectConfig', '#/definitions/ProjectConfig'),
  RestDemoProjectWorkItemsResponse: info<RestDemoProjectWorkItemsResponse>(
    'RestDemoProjectWorkItemsResponse',
    '#/definitions/RestDemoProjectWorkItemsResponse',
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
  NotificationType: info<NotificationType>('NotificationType', '#/definitions/NotificationType'),
  WorkItemRole: info<WorkItemRole>('WorkItemRole', '#/definitions/WorkItemRole'),
  UpdateUserRequest: info<UpdateUserRequest>('UpdateUserRequest', '#/definitions/UpdateUserRequest'),
  UpdateUserAuthRequest: info<UpdateUserAuthRequest>('UpdateUserAuthRequest', '#/definitions/UpdateUserAuthRequest'),
  ValidationError: info<ValidationError>('ValidationError', '#/definitions/ValidationError'),
  UuidUUID: info<UuidUUID>('UuidUUID', '#/definitions/UuidUUID'),
  PgtypeJSONB: info<PgtypeJSONB>('PgtypeJSONB', '#/definitions/PgtypeJSONB'),
  demoProjectKanbanSteps: info<demoProjectKanbanSteps>(
    'demoProjectKanbanSteps',
    '#/definitions/demoProjectKanbanSteps',
  ),
  ModelsProjectConfigField: info<ModelsProjectConfigField>(
    'ModelsProjectConfigField',
    '#/definitions/ModelsProjectConfigField',
  ),
  DbProject2WorkItem: info<DbProject2WorkItem>('DbProject2WorkItem', '#/definitions/DbProject2WorkItem'),
  DbWorkItem: info<DbWorkItem>('DbWorkItem', '#/definitions/DbWorkItem'),
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
  ModelsProject: info<ModelsProject>('ModelsProject', '#/definitions/ModelsProject'),
  ModelsRole: info<ModelsRole>('ModelsRole', '#/definitions/ModelsRole'),
  DbWorkItemRole: info<DbWorkItemRole>('DbWorkItemRole', '#/definitions/DbWorkItemRole'),
  DbWorkItem_Member: info<DbWorkItem_Member>('DbWorkItem_Member', '#/definitions/DbWorkItem_Member'),
}

export interface SchemaInfo<T> {
  definitionName: string
  schemaRef: string
}

function info<T>(definitionName: string, schemaRef: string): SchemaInfo<T> {
  return { definitionName, schemaRef }
}
