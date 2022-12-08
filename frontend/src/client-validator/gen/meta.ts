/* eslint-disable */
// @ts-nocheck
/* eslint-disable */
/* eslint-disable @typescript-eslint/ban-ts-comment */

import {
  ProjectBoardCreateRequest,
  ProjectBoardResponse,
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
  PgtypeJSONB,
  UuidUUID,
  TeamPublic,
  TimeEntryPublic,
  WorkItemCommentPublic,
  ModelsRole,
  RestUserResponse,
  ModelsScope,
  UserAPIKeyPublic,
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
  RestProjectBoardCreateRequest,
  RestProjectBoardResponse,
} from './models'

export const schemaDefinitions = {
  ProjectBoardCreateRequest: info<ProjectBoardCreateRequest>(
    'ProjectBoardCreateRequest',
    '#/definitions/ProjectBoardCreateRequest',
  ),
  ProjectBoardResponse: info<ProjectBoardResponse>('ProjectBoardResponse', '#/definitions/ProjectBoardResponse'),
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
  PgtypeJSONB: info<PgtypeJSONB>('PgtypeJSONB', '#/definitions/PgtypeJSONB'),
  UuidUUID: info<UuidUUID>('UuidUUID', '#/definitions/UuidUUID'),
  TeamPublic: info<TeamPublic>('TeamPublic', '#/definitions/TeamPublic'),
  TimeEntryPublic: info<TimeEntryPublic>('TimeEntryPublic', '#/definitions/TimeEntryPublic'),
  WorkItemCommentPublic: info<WorkItemCommentPublic>('WorkItemCommentPublic', '#/definitions/WorkItemCommentPublic'),
  ModelsRole: info<ModelsRole>('ModelsRole', '#/definitions/ModelsRole'),
  RestUserResponse: info<RestUserResponse>('RestUserResponse', '#/definitions/RestUserResponse'),
  ModelsScope: info<ModelsScope>('ModelsScope', '#/definitions/ModelsScope'),
  UserAPIKeyPublic: info<UserAPIKeyPublic>('UserAPIKeyPublic', '#/definitions/UserAPIKeyPublic'),
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
  RestProjectBoardCreateRequest: info<RestProjectBoardCreateRequest>(
    'RestProjectBoardCreateRequest',
    '#/definitions/RestProjectBoardCreateRequest',
  ),
  RestProjectBoardResponse: info<RestProjectBoardResponse>(
    'RestProjectBoardResponse',
    '#/definitions/RestProjectBoardResponse',
  ),
}

export interface SchemaInfo<T> {
  definitionName: string
  schemaRef: string
}

function info<T>(definitionName: string, schemaRef: string): SchemaInfo<T> {
  return { definitionName, schemaRef }
}
