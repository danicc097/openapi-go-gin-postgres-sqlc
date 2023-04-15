/* eslint-disable */
/* tslint:disable */
/**
 * This file was automatically generated by json-schema-to-typescript.
 * DO NOT MODIFY IT BY HAND. Instead, modify the source JSONSchema file,
 * and run json-schema-to-typescript to regenerate this file.
 */

export type UuidUUID = string
/**
 * Existing projects
 */
export type Project = 'demoProject' | 'demoProject2'
export type DbUserAPIKey = {
  apiKey: string
  expiresOn: string
  user?: DbUser
  userID: UuidUUID
} & DbUserAPIKey1
export type DbUserAPIKey1 = {
  apiKey: string
  expiresOn: string
  user?: DbUser
  userID: UuidUUID
} | null
export type DbDemoProjectWorkItem = {
  lastMessageAt: string
  line: string
  ref: string
  reopened: boolean
  workItem?: DbWorkItem
  workItemID: number
} & DbDemoProjectWorkItem1
export type DbDemoProjectWorkItem1 = {
  lastMessageAt: string
  line: string
  ref: string
  reopened: boolean
  workItem?: DbWorkItem
  workItemID: number
} | null
export type DbProject2WorkItem = {
  customDateForProject2: string | null
  workItem?: DbWorkItem
  workItemID: number
} & DbProject2WorkItem1
export type DbProject2WorkItem1 = {
  customDateForProject2: string | null
  workItem?: DbWorkItem
  workItemID: number
} | null
export type DbWorkItemType = {
  color: string
  description: string
  name: string
  projectID: number
  workItem?: DbWorkItem
  workItemTypeID: number
} & DbWorkItemType1
export type DbWorkItemType1 = {
  color: string
  description: string
  name: string
  projectID: number
  workItem?: DbWorkItem
  workItemTypeID: number
} | null
export type Role = 'guest' | 'user' | 'advancedUser' | 'manager' | 'admin' | 'superAdmin'
export type Scope =
  | 'test-scope'
  | 'users:read'
  | 'users:write'
  | 'scopes:write'
  | 'team-settings:write'
  | 'project-settings:write'
  | 'work-item:review'
export type Scopes = Scope[]
export type Location = string[]
export type Message = string
export type ErrorType = string
export type Detail = ValidationError[]
/**
 * string identifiers for SSE event listeners.
 */
export type Topics = 'GlobalAlerts'
/**
 * User notification type.
 */
export type NotificationType = 'personal' | 'global'
/**
 * Role in work item for a member.
 */
export type WorkItemRole = 'preparer' | 'reviewer'
/**
 * Kanban columns for project demoProject
 */
export type DemoProjectKanbanSteps = 'Disabled' | 'Received' | 'Under review' | 'Work in progress'
export type ModelsProject = string
export type ModelsRole = string

export interface DbActivity {
  activityID: number
  description: string
  isProductive: boolean
  name: string
  projectID: number
  timeEntries?: DbTimeEntry[] | null
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
export interface DbKanbanStep {
  color: string
  description: string
  kanbanStepID: number
  name: string
  projectID: number
  stepOrder: number | null
  timeTrackable: boolean
}
export interface DbProject {
  activities?: DbActivity[] | null
  createdAt: string
  description: string
  initialized: boolean
  kanbanSteps?: DbKanbanStep[] | null
  name: Project
  projectID: number
  teams?: DbTeam[] | null
  updatedAt: string
  workItemTags?: DbWorkItemTag[] | null
  workItemTypes?: DbWorkItemType1[] | null
}
export interface DbTeam {
  createdAt: string
  description: string
  name: string
  projectID: number
  teamID: number
  timeEntries?: DbTimeEntry[] | null
  updatedAt: string
  users?: DbUser[] | null
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
  teams?: DbTeam[] | null
  timeEntries?: DbTimeEntry[] | null
  userAPIKey?: DbUserAPIKey
  userID: UuidUUID
  username: string
  workItems?: DbWorkItem[] | null
}
export interface DbWorkItem {
  closed: string | null
  createdAt: string
  deletedAt: string | null
  demoProjectWorkItem?: DbDemoProjectWorkItem
  description: string
  kanbanStepID: number
  members?: DbUser[] | null
  metadata: PgtypeJSONB
  project2workItem?: DbProject2WorkItem
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
export interface PgtypeJSONB {}
export interface DbWorkItemComment {
  createdAt: string
  message: string
  updatedAt: string
  userID: UuidUUID
  workItemCommentID: number
  workItemID: number
}
export interface DbWorkItemTag {
  color: string
  description: string
  name: string
  projectID: number
  workItemTagID: number
  workItems?: DbWorkItem[] | null
}
export interface ProjectConfig {
  fields: ModelsProjectConfigField[] | null
  header: string[] | null
}
export interface ModelsProjectConfigField {
  isEditable: boolean
  isVisible: boolean
  name: string
  path: string
  showCollapsed: boolean
}
export interface RestDemoProjectWorkItemsResponse {
  closed: string | null
  createdAt: string
  deletedAt: string | null
  demoProjectWorkItem: DbDemoProjectWorkItem1
  description: string
  kanbanStepID: number
  members?: DbUser[] | null
  metadata: PgtypeJSONB
  project2workItem?: DbProject2WorkItem1
  targetDate: string
  teamID: number
  timeEntries?: DbTimeEntry[] | null
  title: string
  updatedAt: string
  workItemComments?: DbWorkItemComment[] | null
  workItemID: number
  workItemTags?: DbWorkItemTag[] | null
  workItemType?: DbWorkItemType1
  workItemTypeID: number
}
export interface InitializeProjectRequest {
  activities?: DbActivityCreateParams[] | null
  kanbanSteps?: DbKanbanStepCreateParams[] | null
  projectID?: number
  teams?: DbTeamCreateParams[] | null
  workItemTags?: DbWorkItemTagCreateParams[] | null
  workItemTypes?: DbWorkItemTypeCreateParams[] | null
}
export interface DbActivityCreateParams {
  description?: string
  isProductive?: boolean
  name?: string
  projectID?: number
}
export interface DbKanbanStepCreateParams {
  color?: string
  description?: string
  name?: string
  projectID?: number
  stepOrder?: number | null
  timeTrackable?: boolean
}
export interface DbTeamCreateParams {
  description?: string
  name?: string
  projectID?: number
}
export interface DbWorkItemTagCreateParams {
  color?: string
  description?: string
  name?: string
  projectID?: number
}
export interface DbWorkItemTypeCreateParams {
  color?: string
  description?: string
  name?: string
  projectID?: number
}
export interface RestProjectBoardResponse {
  project?: DbProject
}
export interface UserResponse {
  apiKey?: DbUserAPIKey1
  projects?: DbProject[] | null
  role: Role
  scopes: Scopes
  teams?: DbTeam[] | null
}
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
