import { emptyInternalApi as api } from '../emptyApi'
export const addTagTypes = ['admin', 'user'] as const
const injectedRtkApi = api
  .enhanceEndpoints({
    addTagTypes,
  })
  .injectEndpoints({
    endpoints: (build) => ({
      events: build.query<EventsRes, EventsArgs>({
        query: () => ({ url: `/events` }),
      }),
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
      updateUserAuthorization: build.mutation<UpdateUserAuthorizationRes, UpdateUserAuthorizationArgs>({
        query: (queryArg) => ({
          url: `/user/${queryArg.id}/authorization`,
          method: 'PATCH',
          body: queryArg.updateUserAuthRequest,
        }),
        invalidatesTags: ['user'],
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
export type EventsRes = unknown
export type EventsArgs = void
export type PingRes = /** status 200 OK */ string
export type PingArgs = void
export type OpenapiYamlGetRes = unknown
export type OpenapiYamlGetArgs = void
export type AdminPingRes = /** status 200 OK */ string
export type AdminPingArgs = void
export type GetCurrentUserRes = /** status 200 ok */ UserResponse
export type GetCurrentUserArgs = void
export type UpdateUserAuthorizationRes = /** status 200 ok */ UserResponse
export type UpdateUserAuthorizationArgs = {
  /** user_id that needs to be updated */
  id: string
  /** Updated user object */
  updateUserAuthRequest: UpdateUserAuthRequest
}
export type DeleteUserRes = unknown
export type DeleteUserArgs = /** user_id that needs to be updated */ string
export type UpdateUserRes = /** status 200 ok */ UserPublic
export type UpdateUserArgs = {
  /** user_id that needs to be updated */
  id: string
  /** Updated user object */
  updateUserRequest: UpdateUserRequest
}
export type ValidationError = {
  loc: string[]
  msg: string
  type: string
}
export type HttpValidationError = {
  detail?: ValidationError[]
}
export type UuidUuid = string
export type UserApiKeyPublic = {
  apiKey: string
  expiresOn: string
  userID: UuidUuid
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
export type PgtypeJsonb = object
export type TeamPublic = {
  createdAt: string
  description: string
  metadata: PgtypeJsonb
  name: string
  projectID: number
  teamID: number
  updatedAt: string
}
export type UserResponse = {
  apiKey?: UserApiKeyPublic
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
  teams?: TeamPublic[] | null
  userID: UuidUuid
  username: string
}
export type UpdateUserAuthRequest = {
  role?: Role
  scopes?: Scopes
}
export type TimeEntryPublic = {
  activityID?: number
  comment?: string
  durationMinutes?: number | null
  start?: string
  teamID?: number | null
  timeEntryID?: number
  userID?: UuidUuid
  workItemID?: number | null
}
export type TaskTypePublic = {
  color?: string
  description?: string
  name?: string
  taskTypeID?: number
  teamID?: number
} | null
export type TaskPublic = {
  createdAt?: string
  deletedAt?: string | null
  finished?: boolean | null
  metadata?: PgtypeJsonb
  taskID?: number
  taskType?: TaskTypePublic
  taskTypeID?: number
  title?: string
  updatedAt?: string
  workItemID?: number
}
export type WorkItemCommentPublic = {
  createdAt?: string
  message?: string
  updatedAt?: string
  userID?: UuidUuid
  workItemCommentID?: number
  workItemID?: number
}
export type WorkItemPublic = {
  closed?: boolean
  createdAt?: string
  deletedAt?: string | null
  kanbanStepID?: number
  metadata?: PgtypeJsonb
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
export type UserPublic = {
  apiKeyID?: number | null
  createdAt?: string
  deletedAt?: string | null
  email?: string
  firstName?: string | null
  fullName?: string | null
  lastName?: string | null
  teams?: TeamPublic[] | null
  timeEntries?: TimeEntryPublic[] | null
  userID?: UuidUuid
  username?: string
  workItems?: WorkItemPublic[] | null
}
export type UpdateUserRequest = {
  first_name?: string
  last_name?: string
}
export const {
  useEventsQuery,
  usePingQuery,
  useOpenapiYamlGetQuery,
  useAdminPingQuery,
  useGetCurrentUserQuery,
  useUpdateUserAuthorizationMutation,
  useDeleteUserMutation,
  useUpdateUserMutation,
} = injectedRtkApi
