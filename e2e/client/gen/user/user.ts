/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { GetPaginatedUsersParams } from '../model/getPaginatedUsersParams'
import type { PaginatedUsersResponse } from '../model/paginatedUsersResponse'
import type { RestUser } from '../model/restUser'
import type { UpdateUserAuthRequest } from '../model/updateUserAuthRequest'
import type { UpdateUserRequest } from '../model/updateUserRequest'
import { customInstance } from '../../api/mutator'

// eslint-disable-next-line
type SecondParameter<T extends (...args: any) => any> = T extends (config: any, args: infer P) => any ? P : never

/**
 * @summary Get paginated users
 */
export const getPaginatedUsers = (
  params: GetPaginatedUsersParams,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<PaginatedUsersResponse>({ url: `/user/page`, method: 'GET', params }, options)
}
/**
 * @summary returns the logged in user
 */
export const getCurrentUser = (options?: SecondParameter<typeof customInstance>) => {
  return customInstance<RestUser>({ url: `/user/me`, method: 'GET' }, options)
}
/**
 * @summary updates user role and scopes by id
 */
export const updateUserAuthorization = (
  id: string,
  updateUserAuthRequest: UpdateUserAuthRequest,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<void>(
    {
      url: `/user/${id}/authorization`,
      method: 'PATCH',
      headers: { 'Content-Type': 'application/json' },
      data: updateUserAuthRequest,
    },
    options,
  )
}
/**
 * @summary deletes the user by id
 */
export const deleteUser = (id: string, options?: SecondParameter<typeof customInstance>) => {
  return customInstance<void>({ url: `/user/${id}`, method: 'DELETE' }, options)
}
/**
 * @summary updates the user by id
 */
export const updateUser = (
  id: string,
  updateUserRequest: UpdateUserRequest,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<RestUser>(
    { url: `/user/${id}`, method: 'PATCH', headers: { 'Content-Type': 'application/json' }, data: updateUserRequest },
    options,
  )
}
export type GetPaginatedUsersResult = NonNullable<Awaited<ReturnType<typeof getPaginatedUsers>>>
export type GetCurrentUserResult = NonNullable<Awaited<ReturnType<typeof getCurrentUser>>>
export type UpdateUserAuthorizationResult = NonNullable<Awaited<ReturnType<typeof updateUserAuthorization>>>
export type DeleteUserResult = NonNullable<Awaited<ReturnType<typeof deleteUser>>>
export type UpdateUserResult = NonNullable<Awaited<ReturnType<typeof updateUser>>>
