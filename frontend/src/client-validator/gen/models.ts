/* eslint-disable */
/* eslint-disable */
/**
 * This file was automatically generated by json-schema-to-typescript.
 * DO NOT MODIFY IT BY HAND. Instead, modify the source JSONSchema file,
 * and run json-schema-to-typescript to regenerate this file.
 */

export type Project = 'demo' | 'demo_two'
export type UuidUUID = string
export type Scope =
  | 'users:read'
  | 'users:write'
  | 'scopes:write'
  | 'team-settings:write'
  | 'project-settings:write'
  | 'work-item-tag:create'
  | 'work-item-tag:edit'
  | 'work-item-tag:delete'
  | 'work-item:review'
export type Scopes = Scope[]
export type Role = 'guest' | 'user' | 'advancedUser' | 'manager' | 'admin' | 'superAdmin'
/**
 * location in body path, if any
 */
export type Location = string[]
/**
 * should always be shown to the user
 */
export type Message = string
/**
 * Additional details for validation errors
 */
export type Detail = ValidationError[]
/**
 * Descriptive error messages to show in a callout
 */
export type Messages = string[]
/**
 * Represents standardized HTTP error types.
 * Notes:
 * - 'Private' marks an error to be hidden in response.
 *
 */
export type ErrorCode =
  | 'Unknown'
  | 'Private'
  | 'NotFound'
  | 'InvalidArgument'
  | 'AlreadyExists'
  | 'Unauthorized'
  | 'Unauthenticated'
  | 'RequestValidation'
  | 'ResponseValidation'
  | 'InvalidRole'
  | 'InvalidScope'
  | 'InvalidUUID'
/**
 * string identifiers for SSE event listeners.
 */
export type Topics = 'GlobalAlerts'
/**
 * represents a database 'work_item_role'
 */
export type WorkItemRole = 'preparer' | 'reviewer'
export type DbWorkItemRole = string
/**
 * represents a database 'notification_type'
 */
export type NotificationType = 'personal' | 'global'
export type DemoProjectKanbanSteps = 'Disabled' | 'Received' | 'Under review' | 'Work in progress'
export type DemoProject2KanbanSteps = 'Received'
export type Demo2WorkItemTypes = 'Type 1' | 'Type 2' | 'Another type'
export type DemoKanbanSteps = 'Disabled' | 'Received' | 'Under review' | 'Work in progress'
export type DemoTwoKanbanSteps = 'Received'
export type DemoTwoWorkItemTypes = 'Type 1' | 'Type 2' | 'Another type'
export type DemoWorkItemTypes = 'Type 1'

export interface DbActivity {
  activityID: number
  description: string
  isProductive: boolean
  name: string
  projectID: number
}
export interface DbKanbanStep {
  color: string
  description: string
  kanbanStepID: number
  name: string
  projectID: number
  stepOrder: number
  timeTrackable: boolean
}
export interface DbProject {
  boardConfig: ProjectConfig
  createdAt: string
  description: string
  name: Project
  projectID: number
  updatedAt: string
}
export interface ProjectConfig {
  fields: ProjectConfigField[]
  header: string[]
}
export interface ProjectConfigField {
  isEditable: boolean
  isVisible: boolean
  name: string
  path: string
  showCollapsed: boolean
}
export interface DbTeam {
  createdAt: string
  description: string
  name: string
  projectID: number
  teamID: number
  updatedAt: string
}
export interface DbWorkItemTag {
  color: string
  description: string
  name: string
  projectID: number
  workItemTagID: number
}
export interface DbWorkItemType {
  color: string
  description: string
  name: string
  projectID: number
  workItemTypeID: number
}
export interface DbDemoWorkItem {
  lastMessageAt: string
  line: string
  ref: string
  reopened: boolean
  workItemID: number
}
export interface DbUserAPIKey {
  apiKey: string
  expiresOn: string
  userID: UuidUUID
}
export interface DbUser {
  createdAt: string
  deletedAt: string | null
  email: string
  firstName: string | null
  fullName: string | null
  hasGlobalNotifications: boolean
  hasPersonalNotifications: boolean
  lastName: string | null
  scopes: Scopes
  userID: UuidUUID
  username: string
}
export interface DbTimeEntry {
  activityID: number
  comment: string
  durationMinutes: number | null
  start: string
  teamID: number | null
  timeEntryID: number
  userID: UuidUUID
  workItemID: number | null
}
export interface DbWorkItemComment {
  createdAt: string
  message: string
  updatedAt: string
  userID: UuidUUID
  workItemCommentID: number
  workItemID: number
}
export interface RestDemoWorkItemsResponse {
  closedAt: string | null
  createdAt: string
  deletedAt: string | null
  demoWorkItem: DbDemoWorkItem
  description: string
  kanbanStepID: number
  members?: DbUser[] | null
  metadata: {} | null
  targetDate: string
  teamID: number
  timeEntries?: DbTimeEntry[] | null
  title: string
  updatedAt: string
  workItemComments?: DbWorkItemComment[] | null
  workItemID: number
  workItemTags?: DbWorkItemTag[] | null
  workItemType?: DbWorkItemType
  workItemTypeID: number
}
export interface InitializeProjectRequest {
  activities?: DbActivityCreateParams[] | null
  projectID?: number
  teams?: DbTeamCreateParams[] | null
  workItemTags?: DbWorkItemTagCreateParams[] | null
}
export interface DbActivityCreateParams {
  description: string
  isProductive: boolean
  name: string
  projectID: number
}
export interface DbTeamCreateParams {
  description: string
  name: string
  projectID: number
}
export interface DbWorkItemTagCreateParams {
  color: string
  description: string
  name: string
  projectID: number
}
export interface RestProjectBoardResponse {
  activities: DbActivity[] | null
  boardConfig: ProjectConfig
  createdAt: string
  description: string
  kanbanSteps: DbKanbanStep[] | null
  name: Project
  projectID: number
  teams: DbTeam[] | null
  updatedAt: string
  workItemTags: DbWorkItemTag[] | null
  workItemTypes: DbWorkItemType[] | null
}
export interface User {
  apiKey?: DbUserAPIKey
  createdAt: string
  deletedAt: string | null
  email: string
  firstName: string | null
  fullName: string | null
  hasGlobalNotifications: boolean
  hasPersonalNotifications: boolean
  lastName: string | null
  projects?: DbProject[] | null
  role: Role
  scopes: Scopes
  teams?: DbTeam[] | null
  userID: UuidUUID
  username: string
}
export interface HTTPValidationError {
  detail?: Detail
  messages: Messages
}
export interface ValidationError {
  loc: Location
  msg: Message
  detail: ErrorDetails
  ctx?: ContextualInformation
}
/**
 * verbose details of the error
 */
export interface ErrorDetails {
  schema: {}
  value: string
}
export interface ContextualInformation {}
/**
 * represents an error message response.
 */
export interface HTTPError {
  title: string
  detail: string
  status: number
  error: string
  type: ErrorCode
  validationError?: HTTPValidationError
}
/**
 * represents User data to update
 */
export interface UpdateUserRequest {
  /**
   * originally from auth server but updatable
   */
  firstName?: string
  /**
   * originally from auth server but updatable
   */
  lastName?: string
}
/**
 * represents User authorization data to update
 */
export interface UpdateUserAuthRequest {
  role?: Role
  scopes?: Scopes
}
export interface PgtypeJSONB {}
export interface DbWorkItem {
  closedAt: string | null
  createdAt: string
  deletedAt: string | null
  description: string
  kanbanStepID: number
  metadata: {} | null
  targetDate: string
  teamID: number
  title: string
  updatedAt: string
  workItemID: number
  workItemTypeID: number
}
export interface RestWorkItemTagCreateRequest {
  color: string
  description: string
  name: string
  projectID: number
}
export interface RestDemoWorkItemCreateRequest {
  base: DbWorkItemCreateParams
  demoProject: DbDemoWorkItemCreateParams
  members: ServicesMember[] | null
  tagIDs: number[] | null
}
export interface DbWorkItemCreateParams {
  closedAt: string | null
  description: string
  kanbanStepID: number
  metadata: {} | null
  targetDate: string
  teamID: number
  title: string
  workItemTypeID: number
}
export interface DbDemoWorkItemCreateParams {
  lastMessageAt: string
  line: string
  ref: string
  reopened: boolean
  workItemID: number
}
export interface ServicesMember {
  role: WorkItemRole
  userID: UuidUUID
}
export interface RestWorkItemCommentCreateRequest {
  message: string
  userID: UuidUUID
  workItemID: number
}
