import { emptyInternalApi as api } from '../emptyApi'
export const addTagTypes = ['admin', 'user'] as const
const injectedRtkApi = api
  .enhanceEndpoints({
    addTagTypes,
  })
  .injectEndpoints({
    endpoints: (build) => ({
      ping: build.query<PingRes, PingArgs>({
        query: () => ({ url: `/ping` }),
      }),
      openapiYamlGet: build.query<OpenapiYamlGetRes, OpenapiYamlGetArgs>({
        query: () => ({ url: `/openapi.yaml` }),
      }),
      adminPing: build.query<AdminPingRes, AdminPingArgs>({
        query: () => ({ url: `/admin/ping` }),
        providesTags: ['admin'],
      }),
      getCurrentUser: build.query<GetCurrentUserRes, GetCurrentUserArgs>({
        query: () => ({ url: `/user/me` }),
        providesTags: ['user'],
      }),
      deleteUser: build.mutation<DeleteUserRes, DeleteUserArgs>({
        query: (queryArg) => ({ url: `/user/${queryArg}`, method: 'DELETE' }),
        invalidatesTags: ['user'],
      }),
      updateUser: build.mutation<UpdateUserRes, UpdateUserArgs>({
        query: (queryArg) => ({ url: `/user/${queryArg.id}`, method: 'PATCH', body: queryArg.updateUserRequest }),
        invalidatesTags: ['user'],
      }),
    }),
    overrideExisting: false,
  })
export { injectedRtkApi as internalApi }
export type PingRes = /** status 200 OK */ string
export type PingArgs = void
export type OpenapiYamlGetRes = unknown
export type OpenapiYamlGetArgs = void
export type AdminPingRes = unknown
export type AdminPingArgs = void
export type GetCurrentUserRes = /** status 200 ok */ User
export type GetCurrentUserArgs = void
export type DeleteUserRes = unknown
export type DeleteUserArgs = /** user_id that needs to be updated */ string
export type UpdateUserRes = unknown
export type UpdateUserArgs = {
  /** user_id that needs to be updated */
  id: string
  /** Updated user object */
  updateUserRequest: AUser
}
export type ValidationError = {
  loc: string[]
  msg: string
  type: string
}
export type HttpValidationError = {
  detail?: ValidationError[]
}
export type PgtypeJsonb = object
export type UuidUuid = string
export type TimeEntry = {
  activityID?: number
  comment?: string
  durationMinutes?: number | null
  start?: string
  teamID?: number | null
  timeEntryID?: number
  userID?: UuidUuid
  workItemID?: number | null
}
export type Team = {
  createdAt?: string
  description?: string
  metadata?: PgtypeJsonb
  name?: string
  projectID?: number
  teamID?: number
  time_entries?: TimeEntry[] | null
  updatedAt?: string
  users?: User[] | null
}
export type TaskType = {
  color?: string
  description?: string
  name?: string
  taskTypeID?: number
  teamID?: number
} | null
export type Task = {
  createdAt?: string
  deletedAt?: string | null
  finished?: boolean | null
  metadata?: PgtypeJsonb
  task_type?: TaskType
  taskID?: number
  taskTypeID?: number
  title?: string
  updatedAt?: string
  workItemID?: number
}
export type WorkItemComment = {
  createdAt?: string
  message?: string
  updatedAt?: string
  userID?: UuidUuid
  workItemCommentID?: number
  workItemID?: number
}
export type WorkItem = {
  closed?: boolean
  createdAt?: string
  deletedAt?: string | null
  kanbanStepID?: number
  metadata?: PgtypeJsonb
  tasks?: Task[] | null
  teamID?: number
  time_entries?: TimeEntry[] | null
  title?: string
  updatedAt?: string
  users?: User[] | null
  work_item_comments?: WorkItemComment[] | null
  workItemID?: number
  workItemTypeID?: number
}
export type User = {
  apiKeyID?: number | null
  createdAt?: string
  deletedAt?: string | null
  email?: string
  externalID?: string
  firstName?: string | null
  fullName?: string | null
  lastName?: string | null
  roleRank?: number
  scopes?: string[] | null
  teams?: Team[] | null
  time_entries?: TimeEntry[] | null
  updatedAt?: string
  userID?: UuidUuid
  username?: string
  work_items?: WorkItem[] | null
}
export type Role = 'guest' | 'user' | 'advancedUser' | 'manager' | 'admin' | 'superAdmin'
export type Scope =
  | 'test-scope'
  | 'users:read'
  | 'users:write'
  | 'scopes:write'
  | 'team-settings:write'
  | 'project-settings:write'
  | 'work-item:review'
export type AUser = {
  first_name?: string
  last_name?: string
  role?: Role
  scopes?: Scope[]
}
export const {
  usePingQuery,
  useOpenapiYamlGetQuery,
  useAdminPingQuery,
  useGetCurrentUserQuery,
  useDeleteUserMutation,
  useUpdateUserMutation,
} = injectedRtkApi
