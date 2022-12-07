import { emptyInternalApi as api } from '../emptyApi'
export const addTagTypes = ['admin', 'user', 'project'] as const
const injectedRtkApi = api
  .enhanceEndpoints({
    addTagTypes,
  })
  .injectEndpoints({
    endpoints: (build) => ({
      myProviderCallback: build.query<MyProviderCallbackRes, MyProviderCallbackArgs>({
        query: () => ({ url: `/auth/myprovider/callback` }),
      }),
      myProviderLogin: build.query<MyProviderLoginRes, MyProviderLoginArgs>({
        query: () => ({ url: `/auth/myprovider/login` }),
      }),
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
      initializeProject: build.mutation<InitializeProjectRes, InitializeProjectArgs>({
        query: (queryArg) => ({
          url: `/project/${queryArg.id}/initialize`,
          method: 'POST',
          body: queryArg.projectBoardCreateRequest,
        }),
        invalidatesTags: ['project'],
      }),
      getProjectBoard: build.query<GetProjectBoardRes, GetProjectBoardArgs>({
        query: (queryArg) => ({ url: `/project/${queryArg}/board` }),
        providesTags: ['project'],
      }),
    }),
    overrideExisting: false,
  })
export { injectedRtkApi as internalApi }
export type MyProviderCallbackRes = unknown
export type MyProviderCallbackArgs = void
export type MyProviderLoginRes = unknown
export type MyProviderLoginArgs = void
export type EventsRes = unknown
export type EventsArgs = void
export type PingRes = /** status 200 OK */ string
export type PingArgs = void
export type OpenapiYamlGetRes = unknown
export type OpenapiYamlGetArgs = void
export type AdminPingRes = /** status 200 OK */ string
export type AdminPingArgs = void
export type GetCurrentUserRes = /** status 200 ok */ RestUserResponse
export type GetCurrentUserArgs = void
export type UpdateUserAuthorizationRes = unknown
export type UpdateUserAuthorizationArgs = {
  /** UUID identifier of entity that needs to be updated */
  id: string
  /** Updated user object */
  updateUserAuthRequest: UpdateUserAuthRequest
}
export type DeleteUserRes = unknown
export type DeleteUserArgs = /** UUID identifier of entity that needs to be updated */ string
export type UpdateUserRes = /** status 200 ok */ RestUserResponse
export type UpdateUserArgs = {
  /** UUID identifier of entity that needs to be updated */
  id: string
  /** Updated user object */
  updateUserRequest: UpdateUserRequest
}
export type InitializeProjectRes = unknown
export type InitializeProjectArgs = {
  /** integer identifier that needs to be updated */
  id: number
  /** Updated user object */
  projectBoardCreateRequest: ProjectBoardCreateRequest
}
export type GetProjectBoardRes = /** status 200 Project successfully initialized. */ ProjectBoardResponse
export type GetProjectBoardArgs = /** integer identifier that needs to be updated */ number
export type ValidationError = {
  loc: string[]
  msg: string
  type: string
}
export type HttpValidationError = {
  detail?: ValidationError[]
}
export type UuidUuid = string
export type DbUserApiKeyPublic = {
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
export type DbTeamPublic = {
  createdAt: string
  description: string
  name: string
  projectID: number
  teamID: number
  updatedAt: string
}
export type RestUserResponse = {
  apiKey?: DbUserApiKeyPublic
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
  userID: UuidUuid
  username: string
}
export type UpdateUserAuthRequest = {
  role?: Role
  scopes?: Scopes
}
export type UpdateUserRequest = {
  first_name?: string
  last_name?: string
}
export type ProjectBoardCreateRequest = any
export type ProjectBoardResponse = any
export const {
  useMyProviderCallbackQuery,
  useMyProviderLoginQuery,
  useEventsQuery,
  usePingQuery,
  useOpenapiYamlGetQuery,
  useAdminPingQuery,
  useGetCurrentUserQuery,
  useUpdateUserAuthorizationMutation,
  useDeleteUserMutation,
  useUpdateUserMutation,
  useInitializeProjectMutation,
  useGetProjectBoardQuery,
} = injectedRtkApi
