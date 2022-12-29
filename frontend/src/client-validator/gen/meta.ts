/* eslint-disable */
// @ts-nocheck
/* eslint-disable */
/* eslint-disable @typescript-eslint/ban-ts-comment */

import {
  ProjectConfig,
  DemoProjectWorkItemsResponse,
  InitializeProjectRequest,
  ProjectBoardResponse,
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
  ModelsRole,
  UuidUUID,
  DbWorkItemPublic,
  PgtypeJSONB,
  DbDemoProjectWorkItemPublic,
  DbTimeEntryPublic,
  DbUserPublic,
  DbWorkItemCommentPublic,
  RestProjectConfigField,
  Project,
  demoProjectKanbanSteps,
} from './models'

export const schemaDefinitions = {
  ProjectConfig: info<ProjectConfig>('ProjectConfig', '#/definitions/ProjectConfig'),
  DemoProjectWorkItemsResponse: info<DemoProjectWorkItemsResponse>(
    'DemoProjectWorkItemsResponse',
    '#/definitions/DemoProjectWorkItemsResponse',
  ),
  InitializeProjectRequest: info<InitializeProjectRequest>(
    'InitializeProjectRequest',
    '#/definitions/InitializeProjectRequest',
  ),
  ProjectBoardResponse: info<ProjectBoardResponse>('ProjectBoardResponse', '#/definitions/ProjectBoardResponse'),
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
  DbTeamPublic: info<DbTeamPublic>('DbTeamPublic', '#/definitions/DbTeamPublic'),
  DbUserAPIKeyPublic: info<DbUserAPIKeyPublic>('DbUserAPIKeyPublic', '#/definitions/DbUserAPIKeyPublic'),
  DbActivityPublic: info<DbActivityPublic>('DbActivityPublic', '#/definitions/DbActivityPublic'),
  DbKanbanStepPublic: info<DbKanbanStepPublic>('DbKanbanStepPublic', '#/definitions/DbKanbanStepPublic'),
  DbProjectPublic: info<DbProjectPublic>('DbProjectPublic', '#/definitions/DbProjectPublic'),
  DbWorkItemTagPublic: info<DbWorkItemTagPublic>('DbWorkItemTagPublic', '#/definitions/DbWorkItemTagPublic'),
  DbWorkItemTypePublic: info<DbWorkItemTypePublic>('DbWorkItemTypePublic', '#/definitions/DbWorkItemTypePublic'),
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
  DbWorkItemPublic: info<DbWorkItemPublic>('DbWorkItemPublic', '#/definitions/DbWorkItemPublic'),
  PgtypeJSONB: info<PgtypeJSONB>('PgtypeJSONB', '#/definitions/PgtypeJSONB'),
  DbDemoProjectWorkItemPublic: info<DbDemoProjectWorkItemPublic>(
    'DbDemoProjectWorkItemPublic',
    '#/definitions/DbDemoProjectWorkItemPublic',
  ),
  DbTimeEntryPublic: info<DbTimeEntryPublic>('DbTimeEntryPublic', '#/definitions/DbTimeEntryPublic'),
  DbUserPublic: info<DbUserPublic>('DbUserPublic', '#/definitions/DbUserPublic'),
  DbWorkItemCommentPublic: info<DbWorkItemCommentPublic>(
    'DbWorkItemCommentPublic',
    '#/definitions/DbWorkItemCommentPublic',
  ),
  RestProjectConfigField: info<RestProjectConfigField>(
    'RestProjectConfigField',
    '#/definitions/RestProjectConfigField',
  ),
  Project: info<Project>('Project', '#/definitions/Project'),
  demoProjectKanbanSteps: info<demoProjectKanbanSteps>(
    'demoProjectKanbanSteps',
    '#/definitions/demoProjectKanbanSteps',
  ),
}

export interface SchemaInfo<T> {
  definitionName: string
  schemaRef: string
}

function info<T>(definitionName: string, schemaRef: string): SchemaInfo<T> {
  return { definitionName, schemaRef }
}
