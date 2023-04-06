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
  DemoProjectWorkItemsResponse,
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
  ReposActivityCreateParams,
  ReposKanbanStepCreateParams,
  ReposTeamCreateParams,
  ReposWorkItemTagCreateParams,
  ReposWorkItemTypeCreateParams,
  ModelsRole,
  UuidUUID,
  PgtypeJSONB,
  Project,
  demoProjectKanbanSteps,
  ModelsProjectConfigField,
  DbProject2WorkItem,
  DbWorkItem,
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
  DemoProjectWorkItemsResponse: info<DemoProjectWorkItemsResponse>(
    'DemoProjectWorkItemsResponse',
    '#/definitions/DemoProjectWorkItemsResponse',
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
  ReposActivityCreateParams: info<ReposActivityCreateParams>(
    'ReposActivityCreateParams',
    '#/definitions/ReposActivityCreateParams',
  ),
  ReposKanbanStepCreateParams: info<ReposKanbanStepCreateParams>(
    'ReposKanbanStepCreateParams',
    '#/definitions/ReposKanbanStepCreateParams',
  ),
  ReposTeamCreateParams: info<ReposTeamCreateParams>('ReposTeamCreateParams', '#/definitions/ReposTeamCreateParams'),
  ReposWorkItemTagCreateParams: info<ReposWorkItemTagCreateParams>(
    'ReposWorkItemTagCreateParams',
    '#/definitions/ReposWorkItemTagCreateParams',
  ),
  ReposWorkItemTypeCreateParams: info<ReposWorkItemTypeCreateParams>(
    'ReposWorkItemTypeCreateParams',
    '#/definitions/ReposWorkItemTypeCreateParams',
  ),
  ModelsRole: info<ModelsRole>('ModelsRole', '#/definitions/ModelsRole'),
  UuidUUID: info<UuidUUID>('UuidUUID', '#/definitions/UuidUUID'),
  PgtypeJSONB: info<PgtypeJSONB>('PgtypeJSONB', '#/definitions/PgtypeJSONB'),
  Project: info<Project>('Project', '#/definitions/Project'),
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
}

export interface SchemaInfo<T> {
  definitionName: string
  schemaRef: string
}

function info<T>(definitionName: string, schemaRef: string): SchemaInfo<T> {
  return { definitionName, schemaRef }
}
