/* eslint-disable */
// @ts-nocheck
/* eslint-disable */
/* eslint-disable @typescript-eslint/ban-ts-comment */

import {
  HTTPValidationError,
  Scope,
  Scopes,
  Role,
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
  UserResponse,
  ModelsScope,
} from './models'

export const schemaDefinitions = {
  HTTPValidationError: info<HTTPValidationError>('HTTPValidationError', '#/definitions/HTTPValidationError'),
  Scope: info<Scope>('Scope', '#/definitions/Scope'),
  Scopes: info<Scopes>('Scopes', '#/definitions/Scopes'),
  Role: info<Role>('Role', '#/definitions/Role'),
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
  UserResponse: info<UserResponse>('UserResponse', '#/definitions/UserResponse'),
  ModelsScope: info<ModelsScope>('ModelsScope', '#/definitions/ModelsScope'),
}

export interface SchemaInfo<T> {
  definitionName: string
  schemaRef: string
}

function info<T>(definitionName: string, schemaRef: string): SchemaInfo<T> {
  return { definitionName, schemaRef }
}
