/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { PaginatedUsersResponse } from '../model/paginatedUsersResponse'
import type { UpdateUserAuthRequest } from '../model/updateUserAuthRequest'
import type { UpdateUserRequest } from '../model/updateUserRequest'
import type { User } from '../model/user'
import { customInstance } from '../../api/mutator'

// eslint-disable-next-line
type SecondParameter<T extends (...args: any) => any> = T extends (config: any, args: infer P) => any ? P : never

/**
 * @summary returns all users
 */
export const getUsers = (options?: SecondParameter<typeof customInstance>) => {
  return customInstance<PaginatedUsersResponse>({ url: `/user/`, method: 'GET' }, options)
}
/**
 * @summary returns the logged in user
 */
export const getCurrentUser = (options?: SecondParameter<typeof customInstance>) => {
  return customInstance<User>({ url: `/user/me`, method: 'GET' }, options)
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
  return customInstance<User>(
    { url: `/user/${id}`, method: 'PATCH', headers: { 'Content-Type': 'application/json' }, data: updateUserRequest },
    options,
  )
}
export type GetUsersResult = NonNullable<Awaited<ReturnType<typeof getUsers>>>
export type GetCurrentUserResult = NonNullable<Awaited<ReturnType<typeof getCurrentUser>>>
export type UpdateUserAuthorizationResult = NonNullable<Awaited<ReturnType<typeof updateUserAuthorization>>>
export type DeleteUserResult = NonNullable<Awaited<ReturnType<typeof deleteUser>>>
export type UpdateUserResult = NonNullable<Awaited<ReturnType<typeof updateUser>>>
