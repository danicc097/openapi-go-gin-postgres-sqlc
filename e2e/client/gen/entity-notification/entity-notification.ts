/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { CreateEntityNotificationRequest } from '../model/createEntityNotificationRequest'
import type { EntityNotification } from '../model/entityNotification'
import type { UpdateEntityNotificationRequest } from '../model/updateEntityNotificationRequest'
import { customInstance } from '../../api/mutator'

// eslint-disable-next-line
type SecondParameter<T extends (...args: any) => any> = T extends (config: any, args: infer P) => any ? P : never

/**
 * @summary create entity notification.
 */
export const createEntityNotification = (
  createEntityNotificationRequest: CreateEntityNotificationRequest,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<EntityNotification>(
    {
      url: `/entity-notification/`,
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      data: createEntityNotificationRequest,
    },
    options,
  )
}
/**
 * @summary get entity notification.
 */
export const getEntityNotification = (id: number, options?: SecondParameter<typeof customInstance>) => {
  return customInstance<EntityNotification>({ url: `/entity-notification/${id}/`, method: 'GET' }, options)
}
/**
 * @summary update entity notification.
 */
export const updateEntityNotification = (
  id: number,
  updateEntityNotificationRequest: UpdateEntityNotificationRequest,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<EntityNotification>(
    {
      url: `/entity-notification/${id}/`,
      method: 'PATCH',
      headers: { 'Content-Type': 'application/json' },
      data: updateEntityNotificationRequest,
    },
    options,
  )
}
/**
 * @summary delete .
 */
export const deleteEntityNotification = (id: number, options?: SecondParameter<typeof customInstance>) => {
  return customInstance<void>({ url: `/entity-notification/${id}/`, method: 'DELETE' }, options)
}
export type CreateEntityNotificationResult = NonNullable<Awaited<ReturnType<typeof createEntityNotification>>>
export type GetEntityNotificationResult = NonNullable<Awaited<ReturnType<typeof getEntityNotification>>>
export type UpdateEntityNotificationResult = NonNullable<Awaited<ReturnType<typeof updateEntityNotification>>>
export type DeleteEntityNotificationResult = NonNullable<Awaited<ReturnType<typeof deleteEntityNotification>>>
