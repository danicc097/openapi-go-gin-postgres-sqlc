/* eslint-disable */
// @ts-nocheck
/* eslint-disable */
/* eslint-disable @typescript-eslint/ban-ts-comment */

import {
  HTTPValidationError,
  ServerSentEvents,
  Scope,
  Scopes,
  Role,
  NotificationType,
  WorkItemRole,
  UpdateUserRequest,
  UpdateUserAuthRequest,
  UserPublic,
  ValidationError,
  PgtypeJSONB,
  UuidUUID,
  TaskPublic,
  TaskTypePublic,
  TeamPublic,
  TimeEntryPublic,
  WorkItemCommentPublic,
  WorkItemPublic,
  ModelsRole,
  RestUserResponse,
  ModelsScope,
  UserAPIKeyPublic,
  DbTeamPublic,
  DbUserAPIKeyPublic,
} from './models'

export const schemaDefinitions = {
  HTTPValidationError: info<HTTPValidationError>('HTTPValidationError', '#/definitions/HTTPValidationError'),
  ServerSentEvents: info<ServerSentEvents>('ServerSentEvents', '#/definitions/ServerSentEvents'),
  Scope: info<Scope>('Scope', '#/definitions/Scope'),
  Scopes: info<Scopes>('Scopes', '#/definitions/Scopes'),
  Role: info<Role>('Role', '#/definitions/Role'),
  NotificationType: info<NotificationType>('NotificationType', '#/definitions/NotificationType'),
  WorkItemRole: info<WorkItemRole>('WorkItemRole', '#/definitions/WorkItemRole'),
  UpdateUserRequest: info<UpdateUserRequest>('UpdateUserRequest', '#/definitions/UpdateUserRequest'),
  UpdateUserAuthRequest: info<UpdateUserAuthRequest>('UpdateUserAuthRequest', '#/definitions/UpdateUserAuthRequest'),
  UserPublic: info<UserPublic>('UserPublic', '#/definitions/UserPublic'),
  ValidationError: info<ValidationError>('ValidationError', '#/definitions/ValidationError'),
  PgtypeJSONB: info<PgtypeJSONB>('PgtypeJSONB', '#/definitions/PgtypeJSONB'),
  UuidUUID: info<UuidUUID>('UuidUUID', '#/definitions/UuidUUID'),
  TaskPublic: info<TaskPublic>('TaskPublic', '#/definitions/TaskPublic'),
  TaskTypePublic: info<TaskTypePublic>('TaskTypePublic', '#/definitions/TaskTypePublic'),
  TeamPublic: info<TeamPublic>('TeamPublic', '#/definitions/TeamPublic'),
  TimeEntryPublic: info<TimeEntryPublic>('TimeEntryPublic', '#/definitions/TimeEntryPublic'),
  WorkItemCommentPublic: info<WorkItemCommentPublic>('WorkItemCommentPublic', '#/definitions/WorkItemCommentPublic'),
  WorkItemPublic: info<WorkItemPublic>('WorkItemPublic', '#/definitions/WorkItemPublic'),
  ModelsRole: info<ModelsRole>('ModelsRole', '#/definitions/ModelsRole'),
  RestUserResponse: info<RestUserResponse>('RestUserResponse', '#/definitions/RestUserResponse'),
  ModelsScope: info<ModelsScope>('ModelsScope', '#/definitions/ModelsScope'),
  UserAPIKeyPublic: info<UserAPIKeyPublic>('UserAPIKeyPublic', '#/definitions/UserAPIKeyPublic'),
  DbTeamPublic: info<DbTeamPublic>('DbTeamPublic', '#/definitions/DbTeamPublic'),
  DbUserAPIKeyPublic: info<DbUserAPIKeyPublic>('DbUserAPIKeyPublic', '#/definitions/DbUserAPIKeyPublic'),
}

export interface SchemaInfo<T> {
  definitionName: string
  schemaRef: string
}

function info<T>(definitionName: string, schemaRef: string): SchemaInfo<T> {
  return { definitionName, schemaRef }
}
