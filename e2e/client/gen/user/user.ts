/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { User, UpdateUserAuthRequest, UpdateUserRequest } from '.././model'
import { customInstance } from '../../api/mutator'

// eslint-disable-next-line
type SecondParameter<T extends (...args: any) => any> = T extends (config: any, args: infer P) => any ? P : never

/**
 * @summary returns the logged in user
 */
export const getCurrentUser = (options?: SecondParameter<typeof customInstance>) => {
  return customInstance<User>({ url: `/user/me`, method: 'get' }, options)
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
      method: 'patch',
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
  return customInstance<void>({ url: `/user/${id}`, method: 'delete' }, options)
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
    { url: `/user/${id}`, method: 'patch', headers: { 'Content-Type': 'application/json' }, data: updateUserRequest },
    options,
  )
}
export type GetCurrentUserResult = NonNullable<Awaited<ReturnType<typeof getCurrentUser>>>
export type UpdateUserAuthorizationResult = NonNullable<Awaited<ReturnType<typeof updateUserAuthorization>>>
export type DeleteUserResult = NonNullable<Awaited<ReturnType<typeof deleteUser>>>
export type UpdateUserResult = NonNullable<Awaited<ReturnType<typeof updateUser>>>