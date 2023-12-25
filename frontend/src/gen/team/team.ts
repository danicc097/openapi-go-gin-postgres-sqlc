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
import type { CreateTeamRequest } from '../model/createTeamRequest'
import type { HTTPError } from '../model/hTTPError'
import type { Team } from '../model/team'
import type { UpdateTeamRequest } from '../model/updateTeamRequest'
import { customInstance } from '../../api/mutator'

// eslint-disable-next-line
type SecondParameter<T extends (...args: any) => any> = T extends (config: any, args: infer P) => any ? P : never

/**
 * @summary create team.
 */
export const createTeam = (
  projectName: 'demo' | 'demo_two',
  createTeamRequest: CreateTeamRequest,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<Team>(
    {
      url: `/project/${projectName}/team/`,
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      data: createTeamRequest,
    },
    options,
  )
}

export const getCreateTeamMutationOptions = <TError = void | HTTPError, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof createTeam>>,
    TError,
    { projectName: 'demo' | 'demo_two'; data: CreateTeamRequest },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}): UseMutationOptions<
  Awaited<ReturnType<typeof createTeam>>,
  TError,
  { projectName: 'demo' | 'demo_two'; data: CreateTeamRequest },
  TContext
> => {
  const { mutation: mutationOptions, request: requestOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof createTeam>>,
    { projectName: 'demo' | 'demo_two'; data: CreateTeamRequest }
  > = (props) => {
    const { projectName, data } = props ?? {}

    return createTeam(projectName, data, requestOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type CreateTeamMutationResult = NonNullable<Awaited<ReturnType<typeof createTeam>>>
export type CreateTeamMutationBody = CreateTeamRequest
export type CreateTeamMutationError = void | HTTPError

/**
 * @summary create team.
 */
export const useCreateTeam = <TError = void | HTTPError, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof createTeam>>,
    TError,
    { projectName: 'demo' | 'demo_two'; data: CreateTeamRequest },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}) => {
  const mutationOptions = getCreateTeamMutationOptions(options)

  return useMutation(mutationOptions)
}
/**
 * @summary get team.
 */
export const getTeam = (
  projectName: 'demo' | 'demo_two',
  id: number,
  options?: SecondParameter<typeof customInstance>,
  signal?: AbortSignal,
) => {
  return customInstance<Team>({ url: `/project/${projectName}/team/${id}/`, method: 'GET', signal }, options)
}

export const getGetTeamQueryKey = (projectName: 'demo' | 'demo_two', id: number) => {
  return [`/project/${projectName}/team/${id}/`] as const
}

export const getGetTeamInfiniteQueryOptions = <TData = Awaited<ReturnType<typeof getTeam>>, TError = void | HTTPError>(
  projectName: 'demo' | 'demo_two',
  id: number,
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getTeam>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
) => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetTeamQueryKey(projectName, id)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getTeam>>> = ({ signal }) =>
    getTeam(projectName, id, requestOptions, signal)

  return {
    queryKey,
    queryFn,
    enabled: !!(projectName && id),
    staleTime: 3600000,
    ...queryOptions,
  } as UseInfiniteQueryOptions<Awaited<ReturnType<typeof getTeam>>, TError, TData> & { queryKey: QueryKey }
}

export type GetTeamInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getTeam>>>
export type GetTeamInfiniteQueryError = void | HTTPError

/**
 * @summary get team.
 */
export const useGetTeamInfinite = <TData = Awaited<ReturnType<typeof getTeam>>, TError = void | HTTPError>(
  projectName: 'demo' | 'demo_two',
  id: number,
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getTeam>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetTeamInfiniteQueryOptions(projectName, id, options)

  const query = useInfiniteQuery(queryOptions) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

export const getGetTeamQueryOptions = <TData = Awaited<ReturnType<typeof getTeam>>, TError = void | HTTPError>(
  projectName: 'demo' | 'demo_two',
  id: number,
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getTeam>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
) => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetTeamQueryKey(projectName, id)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getTeam>>> = ({ signal }) =>
    getTeam(projectName, id, requestOptions, signal)

  return { queryKey, queryFn, enabled: !!(projectName && id), staleTime: 3600000, ...queryOptions } as UseQueryOptions<
    Awaited<ReturnType<typeof getTeam>>,
    TError,
    TData
  > & { queryKey: QueryKey }
}

export type GetTeamQueryResult = NonNullable<Awaited<ReturnType<typeof getTeam>>>
export type GetTeamQueryError = void | HTTPError

/**
 * @summary get team.
 */
export const useGetTeam = <TData = Awaited<ReturnType<typeof getTeam>>, TError = void | HTTPError>(
  projectName: 'demo' | 'demo_two',
  id: number,
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getTeam>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetTeamQueryOptions(projectName, id, options)

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

/**
 * @summary update team.
 */
export const updateTeam = (
  projectName: 'demo' | 'demo_two',
  id: number,
  updateTeamRequest: UpdateTeamRequest,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<Team>(
    {
      url: `/project/${projectName}/team/${id}/`,
      method: 'PATCH',
      headers: { 'Content-Type': 'application/json' },
      data: updateTeamRequest,
    },
    options,
  )
}

export const getUpdateTeamMutationOptions = <TError = void | HTTPError, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof updateTeam>>,
    TError,
    { projectName: 'demo' | 'demo_two'; id: number; data: UpdateTeamRequest },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}): UseMutationOptions<
  Awaited<ReturnType<typeof updateTeam>>,
  TError,
  { projectName: 'demo' | 'demo_two'; id: number; data: UpdateTeamRequest },
  TContext
> => {
  const { mutation: mutationOptions, request: requestOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof updateTeam>>,
    { projectName: 'demo' | 'demo_two'; id: number; data: UpdateTeamRequest }
  > = (props) => {
    const { projectName, id, data } = props ?? {}

    return updateTeam(projectName, id, data, requestOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type UpdateTeamMutationResult = NonNullable<Awaited<ReturnType<typeof updateTeam>>>
export type UpdateTeamMutationBody = UpdateTeamRequest
export type UpdateTeamMutationError = void | HTTPError

/**
 * @summary update team.
 */
export const useUpdateTeam = <TError = void | HTTPError, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof updateTeam>>,
    TError,
    { projectName: 'demo' | 'demo_two'; id: number; data: UpdateTeamRequest },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}) => {
  const mutationOptions = getUpdateTeamMutationOptions(options)

  return useMutation(mutationOptions)
}
/**
 * @summary delete team.
 */
export const deleteTeam = (
  projectName: 'demo' | 'demo_two',
  id: number,
  options?: SecondParameter<typeof customInstance>,
) => {
  return customInstance<void>({ url: `/project/${projectName}/team/${id}/`, method: 'DELETE' }, options)
}

export const getDeleteTeamMutationOptions = <TError = HTTPError, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof deleteTeam>>,
    TError,
    { projectName: 'demo' | 'demo_two'; id: number },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}): UseMutationOptions<
  Awaited<ReturnType<typeof deleteTeam>>,
  TError,
  { projectName: 'demo' | 'demo_two'; id: number },
  TContext
> => {
  const { mutation: mutationOptions, request: requestOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof deleteTeam>>,
    { projectName: 'demo' | 'demo_two'; id: number }
  > = (props) => {
    const { projectName, id } = props ?? {}

    return deleteTeam(projectName, id, requestOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type DeleteTeamMutationResult = NonNullable<Awaited<ReturnType<typeof deleteTeam>>>

export type DeleteTeamMutationError = HTTPError

/**
 * @summary delete team.
 */
export const useDeleteTeam = <TError = HTTPError, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof deleteTeam>>,
    TError,
    { projectName: 'demo' | 'demo_two'; id: number },
    TContext
  >
  request?: SecondParameter<typeof customInstance>
}) => {
  const mutationOptions = getDeleteTeamMutationOptions(options)

  return useMutation(mutationOptions)
}
