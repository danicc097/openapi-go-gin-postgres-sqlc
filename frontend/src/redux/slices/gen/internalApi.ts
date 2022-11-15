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
        query: (queryArg) => ({ url: `/user/${queryArg.id}`, method: 'PUT', body: queryArg.updateUserRequest }),
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
export type AdminPingRes = /** status 200 OK */ string
export type AdminPingArgs = void
export type GetCurrentUserRes = /** status 200 ok */ User
export type GetCurrentUserArgs = void
export type DeleteUserRes = unknown
export type DeleteUserArgs = /** user_id that needs to be deleted */ string
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
  activity_id?: number
  comment?: string
  duration_minutes?: number | null
  start?: string
  task_id?: number | null
  team_id?: number | null
  time_entry_id?: number
  user_id?: UuidUuid
}
export type Team = {
  created_at?: string
  description?: string
  metadata?: PgtypeJsonb
  name?: string
  project_id?: number
  team_id?: number
  time_entries?: TimeEntry[] | null
  updated_at?: string
  users?: User[] | null
}
export type UserApiKey = {
  api_key?: string
  expires_on?: string
  user_api_key_id?: number
} | null
export type TaskType = {
  color?: string
  description?: string
  name?: string
  task_type_id?: number
  team_id?: number
} | null
export type Task = {
  created_at?: string
  deleted_at?: string | null
  metadata?: PgtypeJsonb
  target_date?: string
  target_date_timezone?: string
  task_id?: number
  task_type?: TaskType
  task_type_id?: number
  time_entries?: TimeEntry[] | null
  title?: string
  updated_at?: string
  work_item_id?: number
}
export type WorkItemComment = {
  created_at?: string
  message?: string
  updated_at?: string
  user_id?: UuidUuid
  work_item_comment_id?: number
  work_item_id?: number
}
export type WorkItem = {
  closed?: boolean
  created_at?: string
  deleted_at?: string | null
  kanban_step_id?: number
  metadata?: PgtypeJsonb
  tasks?: Task[] | null
  team_id?: number
  title?: string
  updated_at?: string
  users?: User[] | null
  work_item_comments?: WorkItemComment[] | null
  work_item_id?: number
}
export type User = {
  api_key_id?: number | null
  created_at?: string
  deleted_at?: string | null
  email?: string
  external_id?: string | null
  first_name?: string | null
  full_name?: string | null
  last_name?: string | null
  role_rank?: number
  scopes?: string[] | null
  teams?: Team[] | null
  time_entries?: TimeEntry[] | null
  updated_at?: string
  user_api_key?: UserApiKey
  user_id?: UuidUuid
  username?: string
  work_items?: WorkItem[] | null
}
export type Role = 'guest' | 'user' | 'advancedUser' | 'manager' | 'admin' | 'superAdmin'
export type AUser = {
  role?: Role
  first_name?: string
  last_name?: string
}
export const {
  usePingQuery,
  useOpenapiYamlGetQuery,
  useAdminPingQuery,
  useGetCurrentUserQuery,
  useDeleteUserMutation,
  useUpdateUserMutation,
} = injectedRtkApi
