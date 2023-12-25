/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { Activity } from '../model/activity'
import type { CreateActivityRequest } from '../model/createActivityRequest'
import type { UpdateActivityRequest } from '../model/updateActivityRequest'
import { customInstance } from '../../api/mutator'

// eslint-disable-next-line
type SecondParameter<T extends (...args: any) => any> = T extends (config: any, args: infer P) => any ? P : never

/**
 * @summary create activity.
 */
export const createActivity = (
  projectName: 'demo' | 'demo_two',
  createActivityRequest: CreateActivityRequest,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<Activity>(
    {
      url: `/project/${projectName}/activity/`,
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      data: createActivityRequest,
    },
    options,
  )
}
/**
 * @summary get activity.
 */
export const getActivity = (
  projectName: 'demo' | 'demo_two',
  id: number,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<Activity>({ url: `/project/${projectName}/activity/${id}/`, method: 'GET' }, options)
}
/**
 * @summary update activity.
 */
export const updateActivity = (
  projectName: 'demo' | 'demo_two',
  id: number,
  updateActivityRequest: UpdateActivityRequest,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<Activity>(
    {
      url: `/project/${projectName}/activity/${id}/`,
      method: 'PATCH',
      headers: { 'Content-Type': 'application/json' },
      data: updateActivityRequest,
    },
    options,
  )
}
/**
 * @summary delete activity.
 */
export const deleteActivity = (
  projectName: 'demo' | 'demo_two',
  id: number,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<void>({ url: `/project/${projectName}/activity/${id}/`, method: 'DELETE' }, options)
}
export type CreateActivityResult = NonNullable<Awaited<ReturnType<typeof createActivity>>>
export type GetActivityResult = NonNullable<Awaited<ReturnType<typeof getActivity>>>
export type UpdateActivityResult = NonNullable<Awaited<ReturnType<typeof updateActivity>>>
export type DeleteActivityResult = NonNullable<Awaited<ReturnType<typeof deleteActivity>>>
