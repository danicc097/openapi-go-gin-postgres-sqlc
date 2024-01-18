import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.23.0 🍺
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
import type { CreateWorkItemTypeRequest } from '../model/createWorkItemTypeRequest'
import type { HTTPError } from '../model/hTTPError'
import type { UpdateWorkItemTypeRequest } from '../model/updateWorkItemTypeRequest'
import type { WorkItemType } from '../model/workItemType'
import { customInstance } from '../../api/mutator'

// eslint-disable-next-line
type SecondParameter<T extends (...args: any) => any> = T extends (config: any, args: infer P) => any ? P : never

/**
 * @summary create workitemtype.
 */
export const createWorkItemType = (
  projectName: 'demo' | 'demo_two',
  createWorkItemTypeRequest: CreateWorkItemTypeRequest,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<WorkItemType>(
    {
      url: `/project/${projectName}/work-item-type/`,
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      data: createWorkItemTypeRequest,
    },
    options,
  )
}

export const getCreateWorkItemTypeMutationOptions = <TError = void | HTTPError, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof createWorkItemType>>,
    TError,
    { projectName: 'demo' | 'demo_two'; data: CreateWorkItemTypeRequest },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}): UseMutationOptions<
  Awaited<ReturnType<typeof createWorkItemType>>,
  TError,
  { projectName: 'demo' | 'demo_two'; data: CreateWorkItemTypeRequest },
  TContext
> => {
  const { mutation: mutationOptions, request: requestOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof createWorkItemType>>,
    { projectName: 'demo' | 'demo_two'; data: CreateWorkItemTypeRequest }
  > = (props) => {
    const { projectName, data } = props ?? {}

    return createWorkItemType(projectName, data, requestOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type CreateWorkItemTypeMutationResult = NonNullable<Awaited<ReturnType<typeof createWorkItemType>>>
export type CreateWorkItemTypeMutationBody = CreateWorkItemTypeRequest
export type CreateWorkItemTypeMutationError = void | HTTPError

/**
 * @summary create workitemtype.
 */
export const useCreateWorkItemType = <TError = void | HTTPError, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof createWorkItemType>>,
    TError,
    { projectName: 'demo' | 'demo_two'; data: CreateWorkItemTypeRequest },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}) => {
  const mutationOptions = getCreateWorkItemTypeMutationOptions(options)

  return useMutation(mutationOptions)
}
/**
 * @summary get workitemtype.
 */
export const getWorkItemType = (
  workItemTypeID: EntityIDs.WorkItemTypeID,
  options?: SecondParameter<typeof customInstance>,
  signal?: AbortSignal,
) => {
  return customInstance<WorkItemType>({ url: `/work-item-type/${workItemTypeID}`, method: 'GET', signal }, options)
}

export const getGetWorkItemTypeQueryKey = (workItemTypeID: EntityIDs.WorkItemTypeID) => {
  return [`/work-item-type/${workItemTypeID}`] as const
}

export const getGetWorkItemTypeInfiniteQueryOptions = <
  TData = Awaited<ReturnType<typeof getWorkItemType>>,
  TError = void | HTTPError,
>(
  workItemTypeID: EntityIDs.WorkItemTypeID,
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getWorkItemType>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
) => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetWorkItemTypeQueryKey(workItemTypeID)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getWorkItemType>>> = ({ signal }) =>
    getWorkItemType(workItemTypeID, requestOptions, signal)

  return {
    queryKey,
    queryFn,
    enabled: !!workItemTypeID,
    cacheTime: 300000,
    refetchOnWindowFocus: false,
    refetchOnMount: false,
    retryOnMount: false,
    staleTime: Infinity,
    keepPreviousData: true,
    retry: function (failureCount, error) {
      return failureCount < 3
    },
    ...queryOptions,
  } as UseInfiniteQueryOptions<Awaited<ReturnType<typeof getWorkItemType>>, TError, TData> & { queryKey: QueryKey }
}

export type GetWorkItemTypeInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getWorkItemType>>>
export type GetWorkItemTypeInfiniteQueryError = void | HTTPError

/**
 * @summary get workitemtype.
 */
export const useGetWorkItemTypeInfinite = <
  TData = Awaited<ReturnType<typeof getWorkItemType>>,
  TError = void | HTTPError,
>(
  workItemTypeID: EntityIDs.WorkItemTypeID,
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getWorkItemType>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetWorkItemTypeInfiniteQueryOptions(workItemTypeID, options)

  const query = useInfiniteQuery(queryOptions) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

export const getGetWorkItemTypeQueryOptions = <
  TData = Awaited<ReturnType<typeof getWorkItemType>>,
  TError = void | HTTPError,
>(
  workItemTypeID: EntityIDs.WorkItemTypeID,
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getWorkItemType>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
) => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetWorkItemTypeQueryKey(workItemTypeID)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getWorkItemType>>> = ({ signal }) =>
    getWorkItemType(workItemTypeID, requestOptions, signal)

  return {
    queryKey,
    queryFn,
    enabled: !!workItemTypeID,
    cacheTime: 300000,
    refetchOnWindowFocus: false,
    refetchOnMount: false,
    retryOnMount: false,
    staleTime: Infinity,
    keepPreviousData: true,
    retry: function (failureCount, error) {
      return failureCount < 3
    },
    ...queryOptions,
  } as UseQueryOptions<Awaited<ReturnType<typeof getWorkItemType>>, TError, TData> & { queryKey: QueryKey }
}

export type GetWorkItemTypeQueryResult = NonNullable<Awaited<ReturnType<typeof getWorkItemType>>>
export type GetWorkItemTypeQueryError = void | HTTPError

/**
 * @summary get workitemtype.
 */
export const useGetWorkItemType = <TData = Awaited<ReturnType<typeof getWorkItemType>>, TError = void | HTTPError>(
  workItemTypeID: EntityIDs.WorkItemTypeID,
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getWorkItemType>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetWorkItemTypeQueryOptions(workItemTypeID, options)

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

/**
 * @summary update workitemtype.
 */
export const updateWorkItemType = (
  workItemTypeID: EntityIDs.WorkItemTypeID,
  updateWorkItemTypeRequest: UpdateWorkItemTypeRequest,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<WorkItemType>(
    {
      url: `/work-item-type/${workItemTypeID}`,
      method: 'PATCH',
      headers: { 'Content-Type': 'application/json' },
      data: updateWorkItemTypeRequest,
    },
    options,
  )
}

export const getUpdateWorkItemTypeMutationOptions = <TError = void | HTTPError, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof updateWorkItemType>>,
    TError,
    { workItemTypeID: EntityIDs.WorkItemTypeID; data: UpdateWorkItemTypeRequest },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}): UseMutationOptions<
  Awaited<ReturnType<typeof updateWorkItemType>>,
  TError,
  { workItemTypeID: EntityIDs.WorkItemTypeID; data: UpdateWorkItemTypeRequest },
  TContext
> => {
  const { mutation: mutationOptions, request: requestOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof updateWorkItemType>>,
    { workItemTypeID: EntityIDs.WorkItemTypeID; data: UpdateWorkItemTypeRequest }
  > = (props) => {
    const { workItemTypeID, data } = props ?? {}

    return updateWorkItemType(workItemTypeID, data, requestOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type UpdateWorkItemTypeMutationResult = NonNullable<Awaited<ReturnType<typeof updateWorkItemType>>>
export type UpdateWorkItemTypeMutationBody = UpdateWorkItemTypeRequest
export type UpdateWorkItemTypeMutationError = void | HTTPError

/**
 * @summary update workitemtype.
 */
export const useUpdateWorkItemType = <TError = void | HTTPError, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof updateWorkItemType>>,
    TError,
    { workItemTypeID: EntityIDs.WorkItemTypeID; data: UpdateWorkItemTypeRequest },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}) => {
  const mutationOptions = getUpdateWorkItemTypeMutationOptions(options)

  return useMutation(mutationOptions)
}
/**
 * @summary delete workitemtype.
 */
export const deleteWorkItemType = (
  workItemTypeID: EntityIDs.WorkItemTypeID,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<void>({ url: `/work-item-type/${workItemTypeID}`, method: 'DELETE' }, options)
}

export const getDeleteWorkItemTypeMutationOptions = <TError = HTTPError, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof deleteWorkItemType>>,
    TError,
    { workItemTypeID: EntityIDs.WorkItemTypeID },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}): UseMutationOptions<
  Awaited<ReturnType<typeof deleteWorkItemType>>,
  TError,
  { workItemTypeID: EntityIDs.WorkItemTypeID },
  TContext
> => {
  const { mutation: mutationOptions, request: requestOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof deleteWorkItemType>>,
    { workItemTypeID: EntityIDs.WorkItemTypeID }
  > = (props) => {
    const { workItemTypeID } = props ?? {}

    return deleteWorkItemType(workItemTypeID, requestOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type DeleteWorkItemTypeMutationResult = NonNullable<Awaited<ReturnType<typeof deleteWorkItemType>>>

export type DeleteWorkItemTypeMutationError = HTTPError

/**
 * @summary delete workitemtype.
 */
export const useDeleteWorkItemType = <TError = HTTPError, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof deleteWorkItemType>>,
    TError,
    { workItemTypeID: EntityIDs.WorkItemTypeID },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}) => {
  const mutationOptions = getDeleteWorkItemTypeMutationOptions(options)

  return useMutation(mutationOptions)
}
