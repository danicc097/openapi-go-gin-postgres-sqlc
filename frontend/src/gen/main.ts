/**
 * Generated by orval v6.10.3 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import axios from 'axios'
import type { AxiosRequestConfig, AxiosResponse, AxiosError } from 'axios'
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
import type {
  HTTPValidationError,
  UserResponse,
  UpdateUserAuthRequest,
  UpdateUserRequest,
  InitializeProjectRequest,
  ProjectBoardResponse,
} from './model'

export const myProviderCallback = (options?: AxiosRequestConfig): Promise<AxiosResponse<void>> => {
  return axios.get(`/auth/myprovider/callback`, options)
}

export const getMyProviderCallbackQueryKey = () => [`/auth/myprovider/callback`]

export type MyProviderCallbackInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof myProviderCallback>>>
export type MyProviderCallbackInfiniteQueryError = AxiosError<unknown>

export const useMyProviderCallbackInfinite = <
  TData = Awaited<ReturnType<typeof myProviderCallback>>,
  TError = AxiosError<unknown>,
>(options?: {
  query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof myProviderCallback>>, TError, TData>
  axios?: AxiosRequestConfig
}): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getMyProviderCallbackQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof myProviderCallback>>> = ({ signal }) =>
    myProviderCallback({ signal, ...axiosOptions })

  const query = useInfiniteQuery<Awaited<ReturnType<typeof myProviderCallback>>, TError, TData>(queryKey, queryFn, {
    staleTime: 3600000,
    ...queryOptions,
  }) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

export type MyProviderCallbackQueryResult = NonNullable<Awaited<ReturnType<typeof myProviderCallback>>>
export type MyProviderCallbackQueryError = AxiosError<unknown>

export const useMyProviderCallback = <
  TData = Awaited<ReturnType<typeof myProviderCallback>>,
  TError = AxiosError<unknown>,
>(options?: {
  query?: UseQueryOptions<Awaited<ReturnType<typeof myProviderCallback>>, TError, TData>
  axios?: AxiosRequestConfig
}): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getMyProviderCallbackQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof myProviderCallback>>> = ({ signal }) =>
    myProviderCallback({ signal, ...axiosOptions })

  const query = useQuery<Awaited<ReturnType<typeof myProviderCallback>>, TError, TData>(queryKey, queryFn, {
    staleTime: 3600000,
    ...queryOptions,
  }) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

export const myProviderLogin = (options?: AxiosRequestConfig): Promise<AxiosResponse<unknown>> => {
  return axios.get(`/auth/myprovider/login`, options)
}

export const getMyProviderLoginQueryKey = () => [`/auth/myprovider/login`]

export type MyProviderLoginInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof myProviderLogin>>>
export type MyProviderLoginInfiniteQueryError = AxiosError<void>

export const useMyProviderLoginInfinite = <
  TData = Awaited<ReturnType<typeof myProviderLogin>>,
  TError = AxiosError<void>,
>(options?: {
  query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof myProviderLogin>>, TError, TData>
  axios?: AxiosRequestConfig
}): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getMyProviderLoginQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof myProviderLogin>>> = ({ signal }) =>
    myProviderLogin({ signal, ...axiosOptions })

  const query = useInfiniteQuery<Awaited<ReturnType<typeof myProviderLogin>>, TError, TData>(queryKey, queryFn, {
    staleTime: 3600000,
    ...queryOptions,
  }) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

export type MyProviderLoginQueryResult = NonNullable<Awaited<ReturnType<typeof myProviderLogin>>>
export type MyProviderLoginQueryError = AxiosError<void>

export const useMyProviderLogin = <
  TData = Awaited<ReturnType<typeof myProviderLogin>>,
  TError = AxiosError<void>,
>(options?: {
  query?: UseQueryOptions<Awaited<ReturnType<typeof myProviderLogin>>, TError, TData>
  axios?: AxiosRequestConfig
}): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getMyProviderLoginQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof myProviderLogin>>> = ({ signal }) =>
    myProviderLogin({ signal, ...axiosOptions })

  const query = useQuery<Awaited<ReturnType<typeof myProviderLogin>>, TError, TData>(queryKey, queryFn, {
    staleTime: 3600000,
    ...queryOptions,
  }) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

export const events = (options?: AxiosRequestConfig): Promise<AxiosResponse<string>> => {
  return axios.get(`/events`, options)
}

export const getEventsQueryKey = () => [`/events`]

export type EventsInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof events>>>
export type EventsInfiniteQueryError = AxiosError<unknown>

export const useEventsInfinite = <TData = Awaited<ReturnType<typeof events>>, TError = AxiosError<unknown>>(options?: {
  query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof events>>, TError, TData>
  axios?: AxiosRequestConfig
}): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getEventsQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof events>>> = ({ signal }) => events({ signal, ...axiosOptions })

  const query = useInfiniteQuery<Awaited<ReturnType<typeof events>>, TError, TData>(queryKey, queryFn, {
    staleTime: 3600000,
    ...queryOptions,
  }) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

export type EventsQueryResult = NonNullable<Awaited<ReturnType<typeof events>>>
export type EventsQueryError = AxiosError<unknown>

export const useEvents = <TData = Awaited<ReturnType<typeof events>>, TError = AxiosError<unknown>>(options?: {
  query?: UseQueryOptions<Awaited<ReturnType<typeof events>>, TError, TData>
  axios?: AxiosRequestConfig
}): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getEventsQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof events>>> = ({ signal }) => events({ signal, ...axiosOptions })

  const query = useQuery<Awaited<ReturnType<typeof events>>, TError, TData>(queryKey, queryFn, {
    staleTime: 3600000,
    ...queryOptions,
  }) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

/**
 * @summary Ping pongs
 */
export const ping = (options?: AxiosRequestConfig): Promise<AxiosResponse<string>> => {
  return axios.get(`/ping`, options)
}

export const getPingQueryKey = () => [`/ping`]

export type PingInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof ping>>>
export type PingInfiniteQueryError = AxiosError<HTTPValidationError>

export const usePingInfinite = <
  TData = Awaited<ReturnType<typeof ping>>,
  TError = AxiosError<HTTPValidationError>,
>(options?: {
  query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof ping>>, TError, TData>
  axios?: AxiosRequestConfig
}): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getPingQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof ping>>> = ({ signal }) => ping({ signal, ...axiosOptions })

  const query = useInfiniteQuery<Awaited<ReturnType<typeof ping>>, TError, TData>(queryKey, queryFn, {
    staleTime: 3600000,
    ...queryOptions,
  }) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

export type PingQueryResult = NonNullable<Awaited<ReturnType<typeof ping>>>
export type PingQueryError = AxiosError<HTTPValidationError>

export const usePing = <TData = Awaited<ReturnType<typeof ping>>, TError = AxiosError<HTTPValidationError>>(options?: {
  query?: UseQueryOptions<Awaited<ReturnType<typeof ping>>, TError, TData>
  axios?: AxiosRequestConfig
}): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getPingQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof ping>>> = ({ signal }) => ping({ signal, ...axiosOptions })

  const query = useQuery<Awaited<ReturnType<typeof ping>>, TError, TData>(queryKey, queryFn, {
    staleTime: 3600000,
    ...queryOptions,
  }) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

/**
 * @summary Returns this very OpenAPI spec.
 */
export const openapiYamlGet = (options?: AxiosRequestConfig): Promise<AxiosResponse<Blob>> => {
  return axios.get(`/openapi.yaml`, {
    responseType: 'blob',
    ...options,
  })
}

export const getOpenapiYamlGetQueryKey = () => [`/openapi.yaml`]

export type OpenapiYamlGetInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof openapiYamlGet>>>
export type OpenapiYamlGetInfiniteQueryError = AxiosError<unknown>

export const useOpenapiYamlGetInfinite = <
  TData = Awaited<ReturnType<typeof openapiYamlGet>>,
  TError = AxiosError<unknown>,
>(options?: {
  query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof openapiYamlGet>>, TError, TData>
  axios?: AxiosRequestConfig
}): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getOpenapiYamlGetQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof openapiYamlGet>>> = ({ signal }) =>
    openapiYamlGet({ signal, ...axiosOptions })

  const query = useInfiniteQuery<Awaited<ReturnType<typeof openapiYamlGet>>, TError, TData>(queryKey, queryFn, {
    staleTime: 3600000,
    ...queryOptions,
  }) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

export type OpenapiYamlGetQueryResult = NonNullable<Awaited<ReturnType<typeof openapiYamlGet>>>
export type OpenapiYamlGetQueryError = AxiosError<unknown>

export const useOpenapiYamlGet = <
  TData = Awaited<ReturnType<typeof openapiYamlGet>>,
  TError = AxiosError<unknown>,
>(options?: {
  query?: UseQueryOptions<Awaited<ReturnType<typeof openapiYamlGet>>, TError, TData>
  axios?: AxiosRequestConfig
}): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getOpenapiYamlGetQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof openapiYamlGet>>> = ({ signal }) =>
    openapiYamlGet({ signal, ...axiosOptions })

  const query = useQuery<Awaited<ReturnType<typeof openapiYamlGet>>, TError, TData>(queryKey, queryFn, {
    staleTime: 3600000,
    ...queryOptions,
  }) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

/**
 * @summary Ping pongs
 */
export const adminPing = (options?: AxiosRequestConfig): Promise<AxiosResponse<string>> => {
  return axios.get(`/admin/ping`, options)
}

export const getAdminPingQueryKey = () => [`/admin/ping`]

export type AdminPingInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof adminPing>>>
export type AdminPingInfiniteQueryError = AxiosError<HTTPValidationError>

export const useAdminPingInfinite = <
  TData = Awaited<ReturnType<typeof adminPing>>,
  TError = AxiosError<HTTPValidationError>,
>(options?: {
  query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof adminPing>>, TError, TData>
  axios?: AxiosRequestConfig
}): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getAdminPingQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof adminPing>>> = ({ signal }) =>
    adminPing({ signal, ...axiosOptions })

  const query = useInfiniteQuery<Awaited<ReturnType<typeof adminPing>>, TError, TData>(queryKey, queryFn, {
    staleTime: 3600000,
    ...queryOptions,
  }) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

export type AdminPingQueryResult = NonNullable<Awaited<ReturnType<typeof adminPing>>>
export type AdminPingQueryError = AxiosError<HTTPValidationError>

export const useAdminPing = <
  TData = Awaited<ReturnType<typeof adminPing>>,
  TError = AxiosError<HTTPValidationError>,
>(options?: {
  query?: UseQueryOptions<Awaited<ReturnType<typeof adminPing>>, TError, TData>
  axios?: AxiosRequestConfig
}): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getAdminPingQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof adminPing>>> = ({ signal }) =>
    adminPing({ signal, ...axiosOptions })

  const query = useQuery<Awaited<ReturnType<typeof adminPing>>, TError, TData>(queryKey, queryFn, {
    staleTime: 3600000,
    ...queryOptions,
  }) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

/**
 * @summary returns the logged in user
 */
export const getCurrentUser = (options?: AxiosRequestConfig): Promise<AxiosResponse<UserResponse>> => {
  return axios.get(`/user/me`, options)
}

export const getGetCurrentUserQueryKey = () => [`/user/me`]

export type GetCurrentUserInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getCurrentUser>>>
export type GetCurrentUserInfiniteQueryError = AxiosError<unknown>

export const useGetCurrentUserInfinite = <
  TData = Awaited<ReturnType<typeof getCurrentUser>>,
  TError = AxiosError<unknown>,
>(options?: {
  query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getCurrentUser>>, TError, TData>
  axios?: AxiosRequestConfig
}): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetCurrentUserQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getCurrentUser>>> = ({ signal }) =>
    getCurrentUser({ signal, ...axiosOptions })

  const query = useInfiniteQuery<Awaited<ReturnType<typeof getCurrentUser>>, TError, TData>(queryKey, queryFn, {
    staleTime: 3600000,
    ...queryOptions,
  }) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

export type GetCurrentUserQueryResult = NonNullable<Awaited<ReturnType<typeof getCurrentUser>>>
export type GetCurrentUserQueryError = AxiosError<unknown>

export const useGetCurrentUser = <
  TData = Awaited<ReturnType<typeof getCurrentUser>>,
  TError = AxiosError<unknown>,
>(options?: {
  query?: UseQueryOptions<Awaited<ReturnType<typeof getCurrentUser>>, TError, TData>
  axios?: AxiosRequestConfig
}): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetCurrentUserQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getCurrentUser>>> = ({ signal }) =>
    getCurrentUser({ signal, ...axiosOptions })

  const query = useQuery<Awaited<ReturnType<typeof getCurrentUser>>, TError, TData>(queryKey, queryFn, {
    staleTime: 3600000,
    ...queryOptions,
  }) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

/**
 * @summary updates user role and scopes by id
 */
export const updateUserAuthorization = (
  id: string,
  updateUserAuthRequest: UpdateUserAuthRequest,
  options?: AxiosRequestConfig,
): Promise<AxiosResponse<void>> => {
  return axios.patch(`/user/${id}/authorization`, updateUserAuthRequest, options)
}

export type UpdateUserAuthorizationMutationResult = NonNullable<Awaited<ReturnType<typeof updateUserAuthorization>>>
export type UpdateUserAuthorizationMutationBody = UpdateUserAuthRequest
export type UpdateUserAuthorizationMutationError = AxiosError<unknown>

export const useUpdateUserAuthorization = <TError = AxiosError<unknown>, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof updateUserAuthorization>>,
    TError,
    { id: string; data: UpdateUserAuthRequest },
    TContext
  >
  axios?: AxiosRequestConfig
}) => {
  const { mutation: mutationOptions, axios: axiosOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof updateUserAuthorization>>,
    { id: string; data: UpdateUserAuthRequest }
  > = (props) => {
    const { id, data } = props ?? {}

    return updateUserAuthorization(id, data, axiosOptions)
  }

  return useMutation<
    Awaited<ReturnType<typeof updateUserAuthorization>>,
    TError,
    { id: string; data: UpdateUserAuthRequest },
    TContext
  >(mutationFn, mutationOptions)
}

/**
 * @summary deletes the user by id
 */
export const deleteUser = (id: string, options?: AxiosRequestConfig): Promise<AxiosResponse<unknown>> => {
  return axios.delete(`/user/${id}`, options)
}

export type DeleteUserMutationResult = NonNullable<Awaited<ReturnType<typeof deleteUser>>>

export type DeleteUserMutationError = AxiosError<void>

export const useDeleteUser = <TError = AxiosError<void>, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<Awaited<ReturnType<typeof deleteUser>>, TError, { id: string }, TContext>
  axios?: AxiosRequestConfig
}) => {
  const { mutation: mutationOptions, axios: axiosOptions } = options ?? {}

  const mutationFn: MutationFunction<Awaited<ReturnType<typeof deleteUser>>, { id: string }> = (props) => {
    const { id } = props ?? {}

    return deleteUser(id, axiosOptions)
  }

  return useMutation<Awaited<ReturnType<typeof deleteUser>>, TError, { id: string }, TContext>(
    mutationFn,
    mutationOptions,
  )
}

/**
 * @summary updates the user by id
 */
export const updateUser = (
  id: string,
  updateUserRequest: UpdateUserRequest,
  options?: AxiosRequestConfig,
): Promise<AxiosResponse<UserResponse>> => {
  return axios.patch(`/user/${id}`, updateUserRequest, options)
}

export type UpdateUserMutationResult = NonNullable<Awaited<ReturnType<typeof updateUser>>>
export type UpdateUserMutationBody = UpdateUserRequest
export type UpdateUserMutationError = AxiosError<unknown>

export const useUpdateUser = <TError = AxiosError<unknown>, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof updateUser>>,
    TError,
    { id: string; data: UpdateUserRequest },
    TContext
  >
  axios?: AxiosRequestConfig
}) => {
  const { mutation: mutationOptions, axios: axiosOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof updateUser>>,
    { id: string; data: UpdateUserRequest }
  > = (props) => {
    const { id, data } = props ?? {}

    return updateUser(id, data, axiosOptions)
  }

  return useMutation<Awaited<ReturnType<typeof updateUser>>, TError, { id: string; data: UpdateUserRequest }, TContext>(
    mutationFn,
    mutationOptions,
  )
}

/**
 * @summary creates initial data (teams, work item types, tags...) for a new project
 */
export const initializeProject = (
  id: number,
  initializeProjectRequest: InitializeProjectRequest,
  options?: AxiosRequestConfig,
): Promise<AxiosResponse<void>> => {
  return axios.post(`/project/${id}/initialize`, initializeProjectRequest, options)
}

export type InitializeProjectMutationResult = NonNullable<Awaited<ReturnType<typeof initializeProject>>>
export type InitializeProjectMutationBody = InitializeProjectRequest
export type InitializeProjectMutationError = AxiosError<unknown>

export const useInitializeProject = <TError = AxiosError<unknown>, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof initializeProject>>,
    TError,
    { id: number; data: InitializeProjectRequest },
    TContext
  >
  axios?: AxiosRequestConfig
}) => {
  const { mutation: mutationOptions, axios: axiosOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof initializeProject>>,
    { id: number; data: InitializeProjectRequest }
  > = (props) => {
    const { id, data } = props ?? {}

    return initializeProject(id, data, axiosOptions)
  }

  return useMutation<
    Awaited<ReturnType<typeof initializeProject>>,
    TError,
    { id: number; data: InitializeProjectRequest },
    TContext
  >(mutationFn, mutationOptions)
}

/**
 * @summary returns board data for a project
 */
export const getProjectBoard = (
  id: number,
  options?: AxiosRequestConfig,
): Promise<AxiosResponse<ProjectBoardResponse>> => {
  return axios.get(`/project/${id}/board`, options)
}

export const getGetProjectBoardQueryKey = (id: number) => [`/project/${id}/board`]

export type GetProjectBoardInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectBoard>>>
export type GetProjectBoardInfiniteQueryError = AxiosError<unknown>

export const useGetProjectBoardInfinite = <
  TData = Awaited<ReturnType<typeof getProjectBoard>>,
  TError = AxiosError<unknown>,
>(
  id: number,
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectBoardQueryKey(id)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectBoard>>> = ({ signal }) =>
    getProjectBoard(id, { signal, ...axiosOptions })

  const query = useInfiniteQuery<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData>(queryKey, queryFn, {
    enabled: !!id,
    staleTime: 3600000,
    ...queryOptions,
  }) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}

export type GetProjectBoardQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectBoard>>>
export type GetProjectBoardQueryError = AxiosError<unknown>

export const useGetProjectBoard = <TData = Awaited<ReturnType<typeof getProjectBoard>>, TError = AxiosError<unknown>>(
  id: number,
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectBoardQueryKey(id)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectBoard>>> = ({ signal }) =>
    getProjectBoard(id, { signal, ...axiosOptions })

  const query = useQuery<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData>(queryKey, queryFn, {
    enabled: !!id,
    staleTime: 3600000,
    ...queryOptions,
  }) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryKey

  return query
}
