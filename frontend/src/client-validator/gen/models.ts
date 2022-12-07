/* eslint-disable */
/* tslint:disable */
/**
 * This file was automatically generated by json-schema-to-typescript.
 * DO NOT MODIFY IT BY HAND. Instead, modify the source JSONSchema file,
 * and run json-schema-to-typescript to regenerate this file.
 */

export type Location = string[]
export type Message = string
export type ErrorType = string
export type Detail = ValidationError[]
/**
 * string identifiers for SSE event listeners.
 */
export type Topics =
  | 'UserNotifications'
  | 'ManagerNotifications'
  | 'AdminNotifications'
  | 'WorkItemMoved'
  | 'WorkItemClosed'
export type Scope =
  | 'test-scope'
  | 'users:read'
  | 'users:write'
  | 'scopes:write'
  | 'team-settings:write'
  | 'project-settings:write'
  | 'work-item:review'
export type Scopes = Scope[]
export type Role = 'guest' | 'user' | 'advancedUser' | 'manager' | 'admin' | 'superAdmin'
/**
 * User notification type.
 */
export type NotificationType = 'personal' | 'global'
/**
 * Role in work item for a member.
 */
export type WorkItemRole = 'preparer' | 'reviewer'
export type UuidUUID = string
export type TaskTypePublic = {
  color?: string
  description?: string
  name?: string
  taskTypeID?: number
  teamID?: number
} & TaskTypePublic1
export type TaskTypePublic1 = {
  color?: string
  description?: string
  name?: string
  taskTypeID?: number
  teamID?: number
} | null
export type ModelsRole = string
export type DbUserAPIKeyPublic = {
  apiKey: string
  expiresOn: string
  userID: UuidUUID
} & DbUserAPIKeyPublic1
export type DbUserAPIKeyPublic1 = {
  apiKey: string
  expiresOn: string
  userID: UuidUUID
} | null
export type ModelsScope = string
export type UserAPIKeyPublic = {
  apiKey: string
  expiresOn: string
  userID: UuidUUID
} & UserAPIKeyPublic1
export type UserAPIKeyPublic1 = {
  apiKey: string
  expiresOn: string
  userID: UuidUUID
} | null

export interface HTTPValidationError {
  detail?: Detail
}
export interface ValidationError {
  loc: Location
  msg: Message
  type: ErrorType
}
/**
 * represents User data to update
 */
export interface UpdateUserRequest {
  /**
   * originally from auth server but updatable
   */
  first_name?: string
  /**
   * originally from auth server but updatable
   */
  last_name?: string
}
/**
 * represents User authorization data to update
 */
export interface UpdateUserAuthRequest {
  role?: Role
  scopes?: Scopes
}
export interface UserPublic {
  apiKeyID?: number | null
  createdAt?: string
  deletedAt?: string | null
  email?: string
  firstName?: string | null
  fullName?: string | null
  lastName?: string | null
  teams?: TeamPublic[] | null
  timeEntries?: TimeEntryPublic[] | null
  userID?: UuidUUID
  username?: string
  workItems?: WorkItemPublic[] | null
}
export interface TeamPublic {
  createdAt: string
  description: string
  metadata: PgtypeJSONB
  name: string
  projectID: number
  teamID: number
  updatedAt: string
}
export interface PgtypeJSONB {}
export interface TimeEntryPublic {
  activityID?: number
  comment?: string
  durationMinutes?: number | null
  start?: string
  teamID?: number | null
  timeEntryID?: number
  userID?: UuidUUID
  workItemID?: number | null
}
export interface WorkItemPublic {
  closed?: boolean
  createdAt?: string
  deletedAt?: string | null
  kanbanStepID?: number
  metadata?: PgtypeJSONB
  tasks?: TaskPublic[] | null
  teamID?: number
  timeEntries?: TimeEntryPublic[] | null
  title?: string
  updatedAt?: string
  users?: UserPublic[] | null
  workItemComments?: WorkItemCommentPublic[] | null
  workItemID?: number
  workItemTypeID?: number
}
export interface TaskPublic {
  createdAt?: string
  deletedAt?: string | null
  finished?: boolean | null
  metadata?: PgtypeJSONB
  taskID?: number
  taskType?: TaskTypePublic
  taskTypeID?: number
  title?: string
  updatedAt?: string
  workItemID?: number
}
export interface WorkItemCommentPublic {
  createdAt?: string
  message?: string
  updatedAt?: string
  userID?: UuidUUID
  workItemCommentID?: number
  workItemID?: number
}
export interface RestUserResponse {
  apiKey?: DbUserAPIKeyPublic
  createdAt: string
  deletedAt: string | null
  email: string
  firstName: string | null
  fullName: string | null
  hasGlobalNotifications: boolean
  hasPersonalNotifications: boolean
  lastName: string | null
  role: Role
  scopes: Scopes
  teams?: DbTeamPublic[] | null
  userID: UuidUUID
  username: string
}
export interface DbTeamPublic {
  createdAt: string
  description: string
  name: string
  projectID: number
  teamID: number
  updatedAt: string
}
