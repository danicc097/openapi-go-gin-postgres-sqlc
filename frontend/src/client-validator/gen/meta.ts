/* eslint-disable */
// @ts-nocheck
/* eslint-disable */
/* eslint-disable @typescript-eslint/ban-ts-comment */

import {
  HTTPValidationError,
  Scope,
  Role,
  WorkItemRole,
  UpdateUserRequest,
  UpdateUserAuthRequest,
  User,
  ValidationError,
  PgtypeJSONB,
  Task,
  TaskType,
  Team,
  TimeEntry,
  UserAPIKey,
  UuidUUID,
  WorkItem,
  WorkItemComment,
} from './models'

export const schemaDefinitions = {
  HTTPValidationError: info<HTTPValidationError>('HTTPValidationError', '#/definitions/HTTPValidationError'),
  Scope: info<Scope>('Scope', '#/definitions/Scope'),
  Role: info<Role>('Role', '#/definitions/Role'),
  WorkItemRole: info<WorkItemRole>('WorkItemRole', '#/definitions/WorkItemRole'),
  UpdateUserRequest: info<UpdateUserRequest>('UpdateUserRequest', '#/definitions/UpdateUserRequest'),
  UpdateUserAuthRequest: info<UpdateUserAuthRequest>('UpdateUserAuthRequest', '#/definitions/UpdateUserAuthRequest'),
  User: info<User>('User', '#/definitions/User'),
  ValidationError: info<ValidationError>('ValidationError', '#/definitions/ValidationError'),
  PgtypeJSONB: info<PgtypeJSONB>('PgtypeJSONB', '#/definitions/PgtypeJSONB'),
  Task: info<Task>('Task', '#/definitions/Task'),
  TaskType: info<TaskType>('TaskType', '#/definitions/TaskType'),
  Team: info<Team>('Team', '#/definitions/Team'),
  TimeEntry: info<TimeEntry>('TimeEntry', '#/definitions/TimeEntry'),
  UserAPIKey: info<UserAPIKey>('UserAPIKey', '#/definitions/UserAPIKey'),
  UuidUUID: info<UuidUUID>('UuidUUID', '#/definitions/UuidUUID'),
  WorkItem: info<WorkItem>('WorkItem', '#/definitions/WorkItem'),
  WorkItemComment: info<WorkItemComment>('WorkItemComment', '#/definitions/WorkItemComment'),
}

export interface SchemaInfo<T> {
  definitionName: string
  schemaRef: string
}

function info<T>(definitionName: string, schemaRef: string): SchemaInfo<T> {
  return { definitionName, schemaRef }
}
