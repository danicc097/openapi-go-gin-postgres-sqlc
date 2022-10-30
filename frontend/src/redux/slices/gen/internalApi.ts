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
export type GetCurrentUserRes = /** status 200 successful operation */ AUser
export type GetCurrentUserArgs = void
export type DeleteUserRes = unknown
export type DeleteUserArgs = /** user_id that needs to be deleted */ string
export type UpdateUserRes = unknown
export type UpdateUserArgs = {
  /** user_id that needs to be updated */
  id: string
  /** Updated user object */
  updateUserRequest: AUser2
}
export type ValidationError = {
  loc: string[]
  msg: string
  type: string
}
export type HttpValidationError = {
  detail?: ValidationError[]
}
export type Role = 'guest' | 'user' | 'advanced user' | 'manager' | 'admin' | 'superadmin'
export type Organization = string
export type AUser = {
  user_id?: number
  username?: string
  first_name?: string
  last_name?: string
  email?: string
  password?: string
  phone?: string
  role?: Role
  orgs?: Organization[]
}
export type AUser2 = {
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
