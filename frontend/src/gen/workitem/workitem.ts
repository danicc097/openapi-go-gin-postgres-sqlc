/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import { useQuery, useInfiniteQuery, useMutation } from '@tanstack/react-query'
import type {
  UseQueryOptions,
  UseInfiniteQueryOptions,
  UseMutationOptions,
  QueryFunction,
  MutationFunction,
  UseQueryResult,
  UseInfiniteQueryResult,
  QueryKey,
} from '@tanstack/react-query'
import type { DbWorkItem, CreateWorkitemBody, DbWorkItemComment, WorkItemCommentCreateRequest } from '.././model'
import { customInstance } from '../../api/mutator'

type AwaitedInput<T> = PromiseLike<T> | T

type Awaited<O> = O extends AwaitedInput<infer T> ? T : never

// eslint-disable-next-line
type SecondParameter<T extends (...args: any) => any> = T extends (config: any, args: infer P) => any ? P : never

/**
 * @summary create workitem
 */
export const createWorkitem = (
  createWorkitemBody: CreateWorkitemBody,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<DbWorkItem>(
    { url: `/workitem/`, method: 'post', headers: { 'Content-Type': 'application/json' }, data: createWorkitemBody },
    options,
  )
}

export const getCreateWorkitemMutationOptions = <TError = unknown, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof createWorkitem>>,
    TError,
    { data: CreateWorkitemBody },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}): UseMutationOptions<Awaited<ReturnType<typeof createWorkitem>>, TError, { data: CreateWorkitemBody }, TContext> => {
  const { mutation: mutationOptions, request: requestOptions } = options ?? {}

  const mutationFn: MutationFunction<Awaited<ReturnType<typeof createWorkitem>>, { data: CreateWorkitemBody }> = (
    props,
  ) => {
    const { data } = props ?? {}

    return createWorkitem(data, requestOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type CreateWorkitemMutationResult = NonNullable<Awaited<ReturnType<typeof createWorkitem>>>
export type CreateWorkitemMutationBody = CreateWorkitemBody
export type CreateWorkitemMutationError = unknown

/**
 * @summary create workitem
 */
export const useCreateWorkitem = <TError = unknown, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof createWorkitem>>,
    TError,
    { data: CreateWorkitemBody },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}) => {
  const mutationOptions = getCreateWorkitemMutationOptions(options)

  return useMutation(mutationOptions)
}
/**
 * @summary get workitem
 */
export const getWorkitem = (id: number, options?: SecondParameter<typeof customInstance>, signal?: AbortSignal) => {
  return customInstance<DbWorkItem>({ url: `/workitem/${id}/`, method: 'get', signal }, options)
}

export const getGetWorkitemQueryKey = (id: number) => [`/workitem/${id}/`] as const

export const getGetWorkitemInfiniteQueryOptions = <TData = Awaited<ReturnType<typeof getWorkitem>>, TError = unknown>(
  id: number,
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getWorkitem>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseInfiniteQueryOptions<Awaited<ReturnType<typeof getWorkitem>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetWorkitemQueryKey(id)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getWorkitem>>> = ({ signal }) =>
    getWorkitem(id, requestOptions, signal)

  return { queryKey, queryFn, enabled: !!id, staleTime: 3600000, ...queryOptions }
}

export type GetWorkitemInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getWorkitem>>>
export type GetWorkitemInfiniteQueryError = unknown

/**
 * @summary get workitem
 */
export const useGetWorkitemInfinite = <TData = Awaited<ReturnType<typeof getWorkitem>>, TError = unknown>(
  id: number,
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getWorkitem>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetWorkitemInfiniteQueryOptions(id, options)

  const query = useInfiniteQuery(queryOptions) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

export const getGetWorkitemQueryOptions = <TData = Awaited<ReturnType<typeof getWorkitem>>, TError = unknown>(
  id: number,
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getWorkitem>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseQueryOptions<Awaited<ReturnType<typeof getWorkitem>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetWorkitemQueryKey(id)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getWorkitem>>> = ({ signal }) =>
    getWorkitem(id, requestOptions, signal)

  return { queryKey, queryFn, enabled: !!id, staleTime: 3600000, ...queryOptions }
}

export type GetWorkitemQueryResult = NonNullable<Awaited<ReturnType<typeof getWorkitem>>>
export type GetWorkitemQueryError = unknown

/**
 * @summary get workitem
 */
export const useGetWorkitem = <TData = Awaited<ReturnType<typeof getWorkitem>>, TError = unknown>(
  id: number,
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getWorkitem>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetWorkitemQueryOptions(id, options)

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

/**
 * @summary update workitem
 */
export const updateWorkitem = (id: number, options?: SecondParameter<typeof customInstance>) => {
  return customInstance<DbWorkItem>({ url: `/workitem/${id}/`, method: 'patch' }, options)
}

export const getUpdateWorkitemMutationOptions = <TError = unknown, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<Awaited<ReturnType<typeof updateWorkitem>>, TError, { id: number }, TContext>
  request?: SecondParameter<typeof customInstance>
}): UseMutationOptions<Awaited<ReturnType<typeof updateWorkitem>>, TError, { id: number }, TContext> => {
  const { mutation: mutationOptions, request: requestOptions } = options ?? {}

  const mutationFn: MutationFunction<Awaited<ReturnType<typeof updateWorkitem>>, { id: number }> = (props) => {
    const { id } = props ?? {}

    return updateWorkitem(id, requestOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type UpdateWorkitemMutationResult = NonNullable<Awaited<ReturnType<typeof updateWorkitem>>>

export type UpdateWorkitemMutationError = unknown

/**
 * @summary update workitem
 */
export const useUpdateWorkitem = <TError = unknown, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<Awaited<ReturnType<typeof updateWorkitem>>, TError, { id: number }, TContext>
  request?: SecondParameter<typeof customInstance>
}) => {
  const mutationOptions = getUpdateWorkitemMutationOptions(options)

  return useMutation(mutationOptions)
}
/**
 * @summary delete workitem
 */
export const deleteWorkitem = (id: number, options?: SecondParameter<typeof customInstance>) => {
  return customInstance<void>({ url: `/workitem/${id}/`, method: 'delete' }, options)
}

export const getDeleteWorkitemMutationOptions = <TError = unknown, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<Awaited<ReturnType<typeof deleteWorkitem>>, TError, { id: number }, TContext>
  request?: SecondParameter<typeof customInstance>
}): UseMutationOptions<Awaited<ReturnType<typeof deleteWorkitem>>, TError, { id: number }, TContext> => {
  const { mutation: mutationOptions, request: requestOptions } = options ?? {}

  const mutationFn: MutationFunction<Awaited<ReturnType<typeof deleteWorkitem>>, { id: number }> = (props) => {
    const { id } = props ?? {}

    return deleteWorkitem(id, requestOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type DeleteWorkitemMutationResult = NonNullable<Awaited<ReturnType<typeof deleteWorkitem>>>

export type DeleteWorkitemMutationError = unknown

/**
 * @summary delete workitem
 */
export const useDeleteWorkitem = <TError = unknown, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<Awaited<ReturnType<typeof deleteWorkitem>>, TError, { id: number }, TContext>
  request?: SecondParameter<typeof customInstance>
}) => {
  const mutationOptions = getDeleteWorkitemMutationOptions(options)

  return useMutation(mutationOptions)
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

export const getCreateWorkitemCommentMutationOptions = <TError = unknown, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof createWorkitemComment>>,
    TError,
    { id: number; data: WorkItemCommentCreateRequest },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}): UseMutationOptions<
  Awaited<ReturnType<typeof createWorkitemComment>>,
  TError,
  { id: number; data: WorkItemCommentCreateRequest },
  TContext
> => {
  const { mutation: mutationOptions, request: requestOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof createWorkitemComment>>,
    { id: number; data: WorkItemCommentCreateRequest }
  > = (props) => {
    const { id, data } = props ?? {}

    return createWorkitemComment(id, data, requestOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type CreateWorkitemCommentMutationResult = NonNullable<Awaited<ReturnType<typeof createWorkitemComment>>>
export type CreateWorkitemCommentMutationBody = WorkItemCommentCreateRequest
export type CreateWorkitemCommentMutationError = unknown

/**
 * @summary create workitem comment
 */
export const useCreateWorkitemComment = <TError = unknown, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof createWorkitemComment>>,
    TError,
    { id: number; data: WorkItemCommentCreateRequest },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}) => {
  const mutationOptions = getCreateWorkitemCommentMutationOptions(options)

  return useMutation(mutationOptions)
}
