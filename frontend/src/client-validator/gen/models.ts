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
export type DbProjectPublic = {
  createdAt: string
  description: string
  initialized: boolean
  name: string
  projectID: number
  updatedAt: string
} & DbProjectPublic1
export type DbProjectPublic1 = {
  createdAt: string
  description: string
  initialized: boolean
  name: string
  projectID: number
  updatedAt: string
} | null

export interface ProjectBoardCreateRequest {}
export interface ProjectBoardResponse {}
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
export interface PgtypeJSONB {}
export interface TeamPublic {
  createdAt: string
  description: string
  metadata: PgtypeJSONB
  name: string
  projectID: number
  teamID: number
  updatedAt: string
}
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
export interface DbActivityPublic {
  activityID: number
  description: string
  isProductive: boolean
  name: string
  projectID: number
}
export interface DbKanbanStepPublic {
  color: string
  description: string
  kanbanStepID: number
  name: string
  projectID: number
  stepOrder: number | null
  timeTrackable: boolean
}
export interface DbWorkItemTagPublic {
  color: string
  description: string
  name: string
  projectID: number
  workItemTagID: number
}
export interface DbWorkItemTypePublic {
  color: string
  description: string
  name: string
  projectID: number
  workItemTypeID: number
}
export interface ReposActivityCreateParams {
  description?: string
  isProductive?: boolean
  name?: string
  projectID?: number
}
export interface ReposKanbanStepCreateParams {
  color?: string
  description?: string
  name?: string
  projectID?: number
  stepOrder?: number
  timeTrackable?: boolean
}
export interface ReposTeamCreateParams {
  description?: string
  name?: string
  projectID?: number
}
export interface ReposWorkItemTagCreateParams {
  color?: string
  description?: string
  name?: string
  projectID?: number
}
export interface ReposWorkItemTypeCreateParams {
  color?: string
  description?: string
  name?: string
  projectID?: number
}
export interface RestProjectBoardCreateRequest {
  activities?: ReposActivityCreateParams[] | null
  kanbanSteps?: ReposKanbanStepCreateParams[] | null
  projectID?: number
  teams?: ReposTeamCreateParams[] | null
  workItemTags?: ReposWorkItemTagCreateParams[] | null
  workItemTypes?: ReposWorkItemTypeCreateParams[] | null
}
export interface RestProjectBoardResponse {
  activities?: DbActivityPublic[] | null
  kanbanSteps?: DbKanbanStepPublic[] | null
  project?: DbProjectPublic1
  teams?: DbTeamPublic[] | null
  workItemTags?: DbWorkItemTagPublic[] | null
  workItemTypes?: DbWorkItemTypePublic[] | null
}
