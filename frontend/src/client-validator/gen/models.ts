/* eslint-disable */
/* eslint-disable */
/**
 * This file was automatically generated by json-schema-to-typescript.
 * DO NOT MODIFY IT BY HAND. Instead, modify the source JSONSchema file,
 * and run json-schema-to-typescript to regenerate this file.
 */

export type Direction = 'asc' | 'desc'
/**
 * is generated from database enum 'notification_type'.
 */
export type NotificationType = 'personal' | 'global'
export type DbUserID = string
/**
 * is generated from projects table.
 */
export type Project = 'demo' | 'demo_two'
/**
 * is generated from scopes.json keys.
 */
export type Scope =
  | 'project-member'
  | 'users:read'
  | 'users:write'
  | 'users:delete'
  | 'scopes:write'
  | 'team-settings:write'
  | 'project-settings:write'
  | 'activity:create'
  | 'activity:edit'
  | 'activity:delete'
  | 'work-item-tag:create'
  | 'work-item-tag:edit'
  | 'work-item-tag:delete'
  | 'work-item:review'
  | 'entity-notification:create'
  | 'entity-notification:edit'
  | 'entity-notification:delete'
export type Scopes = Scope[]
/**
 * is generated from database enum 'work_item_role'.
 */
export type WorkItemRole = 'preparer' | 'reviewer'
/**
 * is generated from roles.json keys.
 */
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
  | 'OIDC'
  | 'InvalidRole'
  | 'InvalidScope'
  | 'InvalidUUID'
/**
 * location in body path, if any
 */
export type Location1 = string[]
/**
 * string identifiers for SSE event listeners.
 */
export type Topics = 'GlobalAlerts'
export type UuidUUID = string
export type CreateWorkItemRequest = CreateDemoWorkItemRequest | CreateDemoTwoWorkItemRequest
export type DbWorkItemRole = string
/**
 * is generated from work_item_types table.
 */
export type DemoTwoWorkItemTypes = 'Type 1' | 'Type 2' | 'Another type'
/**
 * is generated from work_item_types table.
 */
export type DemoWorkItemTypes = 'Type 1'
/**
 * is generated from kanban_steps table.
 */
export type DemoKanbanSteps = 'Disabled' | 'Received' | 'Under review' | 'Work in progress'
/**
 * is generated from kanban_steps table.
 */
export type DemoTwoKanbanSteps = 'Received'

export interface CreateActivityRequest {
  description: string
  isProductive: boolean
  name: string
}
export interface UpdateActivityRequest {
  description?: string
  isProductive?: boolean
  name?: string
}
export interface Activity {
  activityID: number
  deletedAt?: string | null
  description: string
  isProductive: boolean
  name: string
  projectID: number
}
export interface CreateWorkItemTagRequest {
  color: string
  description: string
  name: string
}
export interface UpdateWorkItemTagRequest {
  color?: string
  description?: string
  name?: string
}
export interface WorkItemTag {
  color: string
  deletedAt?: string | null
  description: string
  name: string
  projectID: number
  workItemTagID: number
}
export interface CreateWorkItemTypeRequest {
  color: string
  description: string
  name: string
}
export interface UpdateWorkItemTypeRequest {
  color?: string
  description?: string
  name?: string
}
export interface WorkItemType {
  color: string
  description: string
  name: string
  projectID: number
  workItemTypeID: number
}
export interface CreateTeamRequest {
  description: string
  name: string
}
export interface UpdateTeamRequest {
  description?: string
  name?: string
}
export interface Team {
  createdAt: string
  description: string
  name: string
  projectID: number
  teamID: number
  updatedAt: string
}
export interface PaginatedNotificationsResponse {
  items: RestNotification[] | null
  page: RestPaginationPage
}
export interface RestNotification {
  notification: DbNotification
  notificationID: number
  read: boolean
  userID: DbUserID
  userNotificationID: number
}
export interface DbNotification {
  body: string
  createdAt: string
  labels: string[]
  link?: string | null
  notificationID: number
  notificationType: NotificationType
  receiver?: DbUserID
  sender: DbUserID
  title: string
}
export interface RestPaginationPage {
  nextCursor?: string
}
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
  visualization?: {}
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
  deletedAt?: string | null
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
  userID: DbUserID
}
export interface DbUser {
  createdAt: string
  deletedAt?: string | null
  email: string
  firstName?: string | null
  fullName?: string | null
  hasGlobalNotifications: boolean
  hasPersonalNotifications: boolean
  lastName?: string | null
  scopes: Scopes
  userID: DbUserID
  username: string
}
export interface DbTimeEntry {
  activityID: number
  comment: string
  durationMinutes?: number | null
  start: string
  teamID?: number | null
  timeEntryID: number
  userID: DbUserID
  workItemID?: number | null
}
export interface DbWorkItemComment {
  createdAt: string
  message: string
  updatedAt: string
  userID: DbUserID
  workItemCommentID: number
  workItemID: number
}
export interface DemoWorkItems {
  closedAt?: string | null
  createdAt: string
  deletedAt?: string | null
  demoWorkItem: DbDemoWorkItem
  description: string
  kanbanStepID: number
  members?: DbUserWIAUWorkItem[] | null
  metadata: {}
  targetDate: string
  teamID: number | null
  timeEntries?: DbTimeEntry[] | null
  title: string
  updatedAt: string
  workItemComments?: DbWorkItemComment[] | null
  workItemID: number
  workItemTags?: DbWorkItemTag[] | null
  workItemType?: DbWorkItemType
  workItemTypeID: number
}
export interface DbUserWIAUWorkItem {
  role: WorkItemRole
  user: DbUser
}
export interface DemoTwoWorkItems {
  closedAt?: string | null
  createdAt: string
  deletedAt?: string | null
  demoTwoWorkItem: DbDemoTwoWorkItem
  description: string
  kanbanStepID: number
  members?: DbUserWIAUWorkItem[] | null
  metadata: {}
  targetDate: string
  teamID: number | null
  timeEntries?: DbTimeEntry[] | null
  title: string
  updatedAt: string
  workItemComments?: DbWorkItemComment[] | null
  workItemID: number
  workItemTags?: DbWorkItemTag[] | null
  workItemType?: DbWorkItemType
  workItemTypeID: number
}
export interface DbDemoTwoWorkItem {
  customDateForProject2?: string | null
  workItemID: number
}
export interface InitializeProjectRequest {
  tags?: DbWorkItemTagCreateParams[] | null
  teams?: DbTeamCreateParams[] | null
}
export interface DbWorkItemTagCreateParams {
  color: string
  description: string
  name: string
}
export interface DbTeamCreateParams {
  description: string
  name: string
}
export interface ProjectBoard {
  projectName: Project
}
export interface User {
  apiKey?: DbUserAPIKey
  createdAt: string
  deletedAt?: string | null
  email: string
  firstName?: string | null
  fullName?: string | null
  hasGlobalNotifications: boolean
  hasPersonalNotifications: boolean
  lastName?: string | null
  projects?: DbProject[] | null
  role: Role
  scopes: Scopes
  teams?: DbTeam[] | null
  userID: DbUserID
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
  loc?: Location1
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
export interface CreateDemoWorkItemRequest {
  base: DbWorkItemCreateParams
  demoProject: DbDemoWorkItemCreateParams
  members: ServicesMember[]
  projectName: Project
  tagIDs: number[]
}
export interface DbWorkItemCreateParams {
  closedAt?: string | null
  description: string
  kanbanStepID: number
  metadata: {}
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
}
export interface ServicesMember {
  role: WorkItemRole
  userID: DbUserID
}
export interface CreateDemoTwoWorkItemRequest {
  base: DbWorkItemCreateParams
  demoTwoProject: DbDemoTwoWorkItemCreateParams
  members: ServicesMember[]
  projectName: Project
  tagIDs: number[]
}
export interface DbDemoTwoWorkItemCreateParams {
  customDateForProject2?: string | null
}
export interface DbWorkItem {
  closedAt?: string | null
  createdAt: string
  deletedAt?: string | null
  description: string
  kanbanStepID: number
  metadata: {}
  targetDate: string
  teamID: number
  title: string
  updatedAt: string
  workItemID: number
  workItemTypeID: number
}
export interface CreateWorkItemCommentRequest {
  message: string
  userID: DbUserID
  workItemID: number
}
export interface DbActivityCreateParams {
  description: string
  isProductive: boolean
  name: string
  projectID?: number
}
export interface DbWorkItemID {}
export interface DbProjectID {}
export interface DbWorkItemTypeID {}
export interface DbNotificationID {}
export interface DbUserNotification {
  notificationID: number
  read: boolean
  userID: DbUserID
  userNotificationID: number
}
export interface CreateEntityNotificationRequest {
  id: string
  message: string
  topic: Topics
}
export interface UpdateEntityNotificationRequest {
  id?: string
  message?: string
  topic?: Topics
}
export interface EntityNotification {
  createdAt: string
  deletedAt?: string | null
  entityNotificationID: number
  id: string
  message: string
  topic: Topics
}
