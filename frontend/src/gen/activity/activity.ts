/**
 * Generated by orval v6.19.1 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import { useInfiniteQuery, useMutation, useQuery } from '@tanstack/react-query'
import type {
  MutationFunction,
  QueryFunction,
  QueryKey,
  UseInfiniteQueryOptions,
  UseInfiniteQueryResult,
  UseMutationOptions,
  UseQueryOptions,
  UseQueryResult,
} from '@tanstack/react-query'
import type { Activity, CreateActivityRequest, HTTPError, UpdateActivityRequest } from '.././model'
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
      method: 'post',
      headers: { 'Content-Type': 'application/json' },
      data: createActivityRequest,
    },
    options,
  )
}

export const getCreateActivityMutationOptions = <TError = void | HTTPError, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof createActivity>>,
    TError,
    { projectName: 'demo' | 'demo_two'; data: CreateActivityRequest },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}): UseMutationOptions<
  Awaited<ReturnType<typeof createActivity>>,
  TError,
  { projectName: 'demo' | 'demo_two'; data: CreateActivityRequest },
  TContext
> => {
  const { mutation: mutationOptions, request: requestOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof createActivity>>,
    { projectName: 'demo' | 'demo_two'; data: CreateActivityRequest }
  > = (props) => {
    const { projectName, data } = props ?? {}

    return createActivity(projectName, data, requestOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type CreateActivityMutationResult = NonNullable<Awaited<ReturnType<typeof createActivity>>>
export type CreateActivityMutationBody = CreateActivityRequest
export type CreateActivityMutationError = void | HTTPError

/**
 * @summary create activity.
 */
export const useCreateActivity = <TError = void | HTTPError, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof createActivity>>,
    TError,
    { projectName: 'demo' | 'demo_two'; data: CreateActivityRequest },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}) => {
  const mutationOptions = getCreateActivityMutationOptions(options)

  return useMutation(mutationOptions)
}
/**
 * @summary get activity.
 */
export const getActivity = (
  projectName: 'demo' | 'demo_two',
  id: number,
  options?: SecondParameter<typeof customInstance>,
  signal?: AbortSignal,
) => {
  return customInstance<Activity>({ url: `/project/${projectName}/activity/${id}/`, method: 'get', signal }, options)
}

export const getGetActivityQueryKey = (projectName: 'demo' | 'demo_two', id: number) => {
  return [`/project/${projectName}/activity/${id}/`] as const
}

export const getGetActivityInfiniteQueryOptions = <
  TData = Awaited<ReturnType<typeof getActivity>>,
  TError = void | HTTPError,
>(
  projectName: 'demo' | 'demo_two',
  id: number,
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getActivity>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
) => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetActivityQueryKey(projectName, id)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getActivity>>> = ({ signal }) =>
    getActivity(projectName, id, requestOptions, signal)

  return {
    queryKey,
    queryFn,
    enabled: !!(projectName && id),
    staleTime: 3600000,
    ...queryOptions,
  } as UseInfiniteQueryOptions<Awaited<ReturnType<typeof getActivity>>, TError, TData> & { queryKey: QueryKey }
}

export type GetActivityInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getActivity>>>
export type GetActivityInfiniteQueryError = void | HTTPError

/**
 * @summary get activity.
 */
export const useGetActivityInfinite = <TData = Awaited<ReturnType<typeof getActivity>>, TError = void | HTTPError>(
  projectName: 'demo' | 'demo_two',
  id: number,
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getActivity>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetActivityInfiniteQueryOptions(projectName, id, options)

  const query = useInfiniteQuery(queryOptions) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

export const getGetActivityQueryOptions = <TData = Awaited<ReturnType<typeof getActivity>>, TError = void | HTTPError>(
  projectName: 'demo' | 'demo_two',
  id: number,
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getActivity>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
) => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetActivityQueryKey(projectName, id)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getActivity>>> = ({ signal }) =>
    getActivity(projectName, id, requestOptions, signal)

  return { queryKey, queryFn, enabled: !!(projectName && id), staleTime: 3600000, ...queryOptions } as UseQueryOptions<
    Awaited<ReturnType<typeof getActivity>>,
    TError,
    TData
  > & { queryKey: QueryKey }
}

export type GetActivityQueryResult = NonNullable<Awaited<ReturnType<typeof getActivity>>>
export type GetActivityQueryError = void | HTTPError

/**
 * @summary get activity.
 */
export const useGetActivity = <TData = Awaited<ReturnType<typeof getActivity>>, TError = void | HTTPError>(
  projectName: 'demo' | 'demo_two',
  id: number,
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getActivity>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetActivityQueryOptions(projectName, id, options)

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
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
      method: 'patch',
      headers: { 'Content-Type': 'application/json' },
      data: updateActivityRequest,
    },
    options,
  )
}

export const getUpdateActivityMutationOptions = <TError = void | HTTPError, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof updateActivity>>,
    TError,
    { projectName: 'demo' | 'demo_two'; id: number; data: UpdateActivityRequest },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}): UseMutationOptions<
  Awaited<ReturnType<typeof updateActivity>>,
  TError,
  { projectName: 'demo' | 'demo_two'; id: number; data: UpdateActivityRequest },
  TContext
> => {
  const { mutation: mutationOptions, request: requestOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof updateActivity>>,
    { projectName: 'demo' | 'demo_two'; id: number; data: UpdateActivityRequest }
  > = (props) => {
    const { projectName, id, data } = props ?? {}

    return updateActivity(projectName, id, data, requestOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type UpdateActivityMutationResult = NonNullable<Awaited<ReturnType<typeof updateActivity>>>
export type UpdateActivityMutationBody = UpdateActivityRequest
export type UpdateActivityMutationError = void | HTTPError

/**
 * @summary update activity.
 */
export const useUpdateActivity = <TError = void | HTTPError, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof updateActivity>>,
    TError,
    { projectName: 'demo' | 'demo_two'; id: number; data: UpdateActivityRequest },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}) => {
  const mutationOptions = getUpdateActivityMutationOptions(options)

  return useMutation(mutationOptions)
}
/**
 * @summary delete activity.
 */
export const deleteActivity = (
  projectName: 'demo' | 'demo_two',
  id: number,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<void>({ url: `/project/${projectName}/activity/${id}/`, method: 'delete' }, options)
}

export const getDeleteActivityMutationOptions = <TError = HTTPError, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof deleteActivity>>,
    TError,
    { projectName: 'demo' | 'demo_two'; id: number },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}): UseMutationOptions<
  Awaited<ReturnType<typeof deleteActivity>>,
  TError,
  { projectName: 'demo' | 'demo_two'; id: number },
  TContext
> => {
  const { mutation: mutationOptions, request: requestOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof deleteActivity>>,
    { projectName: 'demo' | 'demo_two'; id: number }
  > = (props) => {
    const { projectName, id } = props ?? {}

    return deleteActivity(projectName, id, requestOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type DeleteActivityMutationResult = NonNullable<Awaited<ReturnType<typeof deleteActivity>>>

export type DeleteActivityMutationError = HTTPError

/**
 * @summary delete activity.
 */
export const useDeleteActivity = <TError = HTTPError, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof deleteActivity>>,
    TError,
    { projectName: 'demo' | 'demo_two'; id: number },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}) => {
  const mutationOptions = getDeleteActivityMutationOptions(options)

  return useMutation(mutationOptions)
}
