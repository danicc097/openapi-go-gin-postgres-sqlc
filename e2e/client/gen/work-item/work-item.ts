/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { CreateWorkItemRequest } from '../model/createWorkItemRequest'
import type { WorkItem } from '../model/workItem'
import { customInstance } from '../../api/mutator'

// eslint-disable-next-line
type SecondParameter<T extends (...args: any) => any> = T extends (config: any, args: infer P) => any ? P : never

/**
 * @summary create workitem
 */
export const createWorkitem = (
  createWorkItemRequest: CreateWorkItemRequest,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<WorkItem>(
    {
      url: `/work-item/`,
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      data: createWorkItemRequest,
    },
    options,
  )
}
/**
 * @summary get workitem
 */
export const getWorkItem = (workItemID: number, options?: SecondParameter<typeof customInstance>) => {
  return customInstance<WorkItem>({ url: `/work-item/${workItemID}/`, method: 'GET' }, options)
}
/**
 * @summary update workitem
 */
export const updateWorkitem = (workItemID: number, options?: SecondParameter<typeof customInstance>) => {
  return customInstance<WorkItem>({ url: `/work-item/${workItemID}/`, method: 'PATCH' }, options)
}
/**
 * @summary delete workitem
 */
export const deleteWorkitem = (workItemID: number, options?: SecondParameter<typeof customInstance>) => {
  return customInstance<void>({ url: `/work-item/${workItemID}/`, method: 'DELETE' }, options)
}
export type CreateWorkitemResult = NonNullable<Awaited<ReturnType<typeof createWorkitem>>>
export type GetWorkItemResult = NonNullable<Awaited<ReturnType<typeof getWorkItem>>>
export type UpdateWorkitemResult = NonNullable<Awaited<ReturnType<typeof updateWorkitem>>>
export type DeleteWorkitemResult = NonNullable<Awaited<ReturnType<typeof deleteWorkitem>>>
