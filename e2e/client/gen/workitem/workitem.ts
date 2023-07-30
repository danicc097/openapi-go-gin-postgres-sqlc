/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import type { DbWorkItem, WorkItemCreateRequest, DbWorkItemComment, WorkItemCommentCreateRequest } from '.././model'
import { customInstance } from '../../api/mutator'

// eslint-disable-next-line
type SecondParameter<T extends (...args: any) => any> = T extends (config: any, args: infer P) => any ? P : never

/**
 * @summary create workitem
 */
export const createWorkitem = (
  workItemCreateRequest: WorkItemCreateRequest,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<DbWorkItem>(
    { url: `/workitem/`, method: 'post', headers: { 'Content-Type': 'application/json' }, data: workItemCreateRequest },
    options,
  )
}
/**
 * @summary get workitem
 */
export const getWorkitem = (id: number, options?: SecondParameter<typeof customInstance>) => {
  return customInstance<DbWorkItem>({ url: `/workitem/${id}/`, method: 'get' }, options)
}
/**
 * @summary update workitem
 */
export const updateWorkitem = (id: number, options?: SecondParameter<typeof customInstance>) => {
  return customInstance<DbWorkItem>({ url: `/workitem/${id}/`, method: 'patch' }, options)
}
/**
 * @summary delete workitem
 */
export const deleteWorkitem = (id: number, options?: SecondParameter<typeof customInstance>) => {
  return customInstance<void>({ url: `/workitem/${id}/`, method: 'delete' }, options)
}
/**
 * @summary create workitem comment
 */
export const createWorkitemComment = (
  id: number,
  workItemCommentCreateRequest: WorkItemCommentCreateRequest,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<DbWorkItemComment>(
    {
      url: `/workitem/${id}/comments/`,
      method: 'post',
      headers: { 'Content-Type': 'application/json' },
      data: workItemCommentCreateRequest,
    },
    options,
  )
}
export type CreateWorkitemResult = NonNullable<Awaited<ReturnType<typeof createWorkitem>>>
export type GetWorkitemResult = NonNullable<Awaited<ReturnType<typeof getWorkitem>>>
export type UpdateWorkitemResult = NonNullable<Awaited<ReturnType<typeof updateWorkitem>>>
export type DeleteWorkitemResult = NonNullable<Awaited<ReturnType<typeof deleteWorkitem>>>
export type CreateWorkitemCommentResult = NonNullable<Awaited<ReturnType<typeof createWorkitemComment>>>