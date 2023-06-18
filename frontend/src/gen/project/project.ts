/**
 * Generated by orval v6.15.0 🍺
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
  InitializeProjectRequest,
  DbProject,
  ProjectConfig,
  RestProjectBoardResponse,
  RestDemoWorkItemsResponse,
  GetProjectWorkitemsParams,
  DbWorkItemTag,
  RestWorkItemTagCreateRequest,
} from '.././model'

type AwaitedInput<T> = PromiseLike<T> | T

type Awaited<O> = O extends AwaitedInput<infer T> ? T : never

/**
 * @summary creates initial data (teams, work item types, tags...) for a new project
 */
export const initializeProject = (
  projectName: 'demo' | 'demo_two',
  initializeProjectRequest: InitializeProjectRequest,
  options?: AxiosRequestConfig,
): Promise<AxiosResponse<void>> => {
  return axios.post(`/project/${projectName}/initialize`, initializeProjectRequest, options)
}

export const getInitializeProjectMutationOptions = <TError = AxiosError<unknown>, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof initializeProject>>,
    TError,
    { projectName: 'demo' | 'demo_two'; data: InitializeProjectRequest },
    TContext
  >
  axios?: AxiosRequestConfig
}): UseMutationOptions<
  Awaited<ReturnType<typeof initializeProject>>,
  TError,
  { projectName: 'demo' | 'demo_two'; data: InitializeProjectRequest },
  TContext
> => {
  const { mutation: mutationOptions, axios: axiosOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof initializeProject>>,
    { projectName: 'demo' | 'demo_two'; data: InitializeProjectRequest }
  > = (props) => {
    const { projectName, data } = props ?? {}

    return initializeProject(projectName, data, axiosOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type InitializeProjectMutationResult = NonNullable<Awaited<ReturnType<typeof initializeProject>>>
export type InitializeProjectMutationBody = InitializeProjectRequest
export type InitializeProjectMutationError = AxiosError<unknown>

export const useInitializeProject = <TError = AxiosError<unknown>, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof initializeProject>>,
    TError,
    { projectName: 'demo' | 'demo_two'; data: InitializeProjectRequest },
    TContext
  >
  axios?: AxiosRequestConfig
}) => {
  const mutationOptions = getInitializeProjectMutationOptions(options)

  return useMutation(mutationOptions)
}
/**
 * @summary returns board data for a project
 */
export const getProject = (
  projectName: 'demo' | 'demo_two',
  options?: AxiosRequestConfig,
): Promise<AxiosResponse<DbProject>> => {
  return axios.get(`/project/${projectName}/`, options)
}

export const getGetProjectQueryKey = (projectName: 'demo' | 'demo_two') => [`/project/${projectName}/`] as const

export const getGetProjectInfiniteQueryOptions = <
  TData = Awaited<ReturnType<typeof getProject>>,
  TError = AxiosError<unknown>,
>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProject>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProject>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectQueryKey(projectName)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProject>>> = ({ signal }) =>
    getProject(projectName, { signal, ...axiosOptions })

  return { queryKey, queryFn, enabled: !!projectName, staleTime: 3600000, ...queryOptions }
}

export type GetProjectInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getProject>>>
export type GetProjectInfiniteQueryError = AxiosError<unknown>

export const useGetProjectInfinite = <TData = Awaited<ReturnType<typeof getProject>>, TError = AxiosError<unknown>>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProject>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetProjectInfiniteQueryOptions(projectName, options)

  const query = useInfiniteQuery(queryOptions) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

export const getGetProjectQueryOptions = <TData = Awaited<ReturnType<typeof getProject>>, TError = AxiosError<unknown>>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getProject>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseQueryOptions<Awaited<ReturnType<typeof getProject>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectQueryKey(projectName)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProject>>> = ({ signal }) =>
    getProject(projectName, { signal, ...axiosOptions })

  return { queryKey, queryFn, enabled: !!projectName, staleTime: 3600000, ...queryOptions }
}

export type GetProjectQueryResult = NonNullable<Awaited<ReturnType<typeof getProject>>>
export type GetProjectQueryError = AxiosError<unknown>

export const useGetProject = <TData = Awaited<ReturnType<typeof getProject>>, TError = AxiosError<unknown>>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getProject>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetProjectQueryOptions(projectName, options)

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

/**
 * @summary returns the project configuration
 */
export const getProjectConfig = (
  projectName: 'demo' | 'demo_two',
  options?: AxiosRequestConfig,
): Promise<AxiosResponse<ProjectConfig>> => {
  return axios.get(`/project/${projectName}/config`, options)
}

export const getGetProjectConfigQueryKey = (projectName: 'demo' | 'demo_two') =>
  [`/project/${projectName}/config`] as const

export const getGetProjectConfigInfiniteQueryOptions = <
  TData = Awaited<ReturnType<typeof getProjectConfig>>,
  TError = AxiosError<unknown>,
>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectConfigQueryKey(projectName)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectConfig>>> = ({ signal }) =>
    getProjectConfig(projectName, { signal, ...axiosOptions })

  return { queryKey, queryFn, enabled: !!projectName, staleTime: 3600000, ...queryOptions }
}

export type GetProjectConfigInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectConfig>>>
export type GetProjectConfigInfiniteQueryError = AxiosError<unknown>

export const useGetProjectConfigInfinite = <
  TData = Awaited<ReturnType<typeof getProjectConfig>>,
  TError = AxiosError<unknown>,
>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetProjectConfigInfiniteQueryOptions(projectName, options)

  const query = useInfiniteQuery(queryOptions) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

export const getGetProjectConfigQueryOptions = <
  TData = Awaited<ReturnType<typeof getProjectConfig>>,
  TError = AxiosError<unknown>,
>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseQueryOptions<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectConfigQueryKey(projectName)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectConfig>>> = ({ signal }) =>
    getProjectConfig(projectName, { signal, ...axiosOptions })

  return { queryKey, queryFn, enabled: !!projectName, staleTime: 3600000, ...queryOptions }
}

export type GetProjectConfigQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectConfig>>>
export type GetProjectConfigQueryError = AxiosError<unknown>

export const useGetProjectConfig = <TData = Awaited<ReturnType<typeof getProjectConfig>>, TError = AxiosError<unknown>>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getProjectConfig>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetProjectConfigQueryOptions(projectName, options)

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

/**
 * @summary updates the project configuration
 */
export const updateProjectConfig = (
  projectName: 'demo' | 'demo_two',
  projectConfig: ProjectConfig,
  options?: AxiosRequestConfig,
): Promise<AxiosResponse<void>> => {
  return axios.put(`/project/${projectName}/config`, projectConfig, options)
}

export const getUpdateProjectConfigMutationOptions = <TError = AxiosError<unknown>, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof updateProjectConfig>>,
    TError,
    { projectName: 'demo' | 'demo_two'; data: ProjectConfig },
    TContext
  >
  axios?: AxiosRequestConfig
}): UseMutationOptions<
  Awaited<ReturnType<typeof updateProjectConfig>>,
  TError,
  { projectName: 'demo' | 'demo_two'; data: ProjectConfig },
  TContext
> => {
  const { mutation: mutationOptions, axios: axiosOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof updateProjectConfig>>,
    { projectName: 'demo' | 'demo_two'; data: ProjectConfig }
  > = (props) => {
    const { projectName, data } = props ?? {}

    return updateProjectConfig(projectName, data, axiosOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type UpdateProjectConfigMutationResult = NonNullable<Awaited<ReturnType<typeof updateProjectConfig>>>
export type UpdateProjectConfigMutationBody = ProjectConfig
export type UpdateProjectConfigMutationError = AxiosError<unknown>

export const useUpdateProjectConfig = <TError = AxiosError<unknown>, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof updateProjectConfig>>,
    TError,
    { projectName: 'demo' | 'demo_two'; data: ProjectConfig },
    TContext
  >
  axios?: AxiosRequestConfig
}) => {
  const mutationOptions = getUpdateProjectConfigMutationOptions(options)

  return useMutation(mutationOptions)
}
/**
 * @summary returns board data for a project
 */
export const getProjectBoard = (
  projectName: 'demo' | 'demo_two',
  options?: AxiosRequestConfig,
): Promise<AxiosResponse<RestProjectBoardResponse>> => {
  return axios.get(`/project/${projectName}/board`, options)
}

export const getGetProjectBoardQueryKey = (projectName: 'demo' | 'demo_two') =>
  [`/project/${projectName}/board`] as const

export const getGetProjectBoardInfiniteQueryOptions = <
  TData = Awaited<ReturnType<typeof getProjectBoard>>,
  TError = AxiosError<unknown>,
>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectBoardQueryKey(projectName)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectBoard>>> = ({ signal }) =>
    getProjectBoard(projectName, { signal, ...axiosOptions })

  return { queryKey, queryFn, enabled: !!projectName, staleTime: 3600000, ...queryOptions }
}

export type GetProjectBoardInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectBoard>>>
export type GetProjectBoardInfiniteQueryError = AxiosError<unknown>

export const useGetProjectBoardInfinite = <
  TData = Awaited<ReturnType<typeof getProjectBoard>>,
  TError = AxiosError<unknown>,
>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetProjectBoardInfiniteQueryOptions(projectName, options)

  const query = useInfiniteQuery(queryOptions) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

export const getGetProjectBoardQueryOptions = <
  TData = Awaited<ReturnType<typeof getProjectBoard>>,
  TError = AxiosError<unknown>,
>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectBoardQueryKey(projectName)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectBoard>>> = ({ signal }) =>
    getProjectBoard(projectName, { signal, ...axiosOptions })

  return { queryKey, queryFn, enabled: !!projectName, staleTime: 3600000, ...queryOptions }
}

export type GetProjectBoardQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectBoard>>>
export type GetProjectBoardQueryError = AxiosError<unknown>

export const useGetProjectBoard = <TData = Awaited<ReturnType<typeof getProjectBoard>>, TError = AxiosError<unknown>>(
  projectName: 'demo' | 'demo_two',
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getProjectBoard>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetProjectBoardQueryOptions(projectName, options)

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

/**
 * @summary returns workitems for a project
 */
export const getProjectWorkitems = (
  projectName: 'demo' | 'demo_two',
  params?: GetProjectWorkitemsParams,
  options?: AxiosRequestConfig,
): Promise<AxiosResponse<RestDemoWorkItemsResponse>> => {
  return axios.get(`/project/${projectName}/workitems`, {
    ...options,
    params: { ...params, ...options?.params },
  })
}

export const getGetProjectWorkitemsQueryKey = (projectName: 'demo' | 'demo_two', params?: GetProjectWorkitemsParams) =>
  [`/project/${projectName}/workitems`, ...(params ? [params] : [])] as const

export const getGetProjectWorkitemsInfiniteQueryOptions = <
  TData = Awaited<ReturnType<typeof getProjectWorkitems>>,
  TError = AxiosError<unknown>,
>(
  projectName: 'demo' | 'demo_two',
  params?: GetProjectWorkitemsParams,
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectWorkitemsQueryKey(projectName, params)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectWorkitems>>> = ({ signal }) =>
    getProjectWorkitems(projectName, params, { signal, ...axiosOptions })

  return { queryKey, queryFn, enabled: !!projectName, staleTime: 3600000, ...queryOptions }
}

export type GetProjectWorkitemsInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectWorkitems>>>
export type GetProjectWorkitemsInfiniteQueryError = AxiosError<unknown>

export const useGetProjectWorkitemsInfinite = <
  TData = Awaited<ReturnType<typeof getProjectWorkitems>>,
  TError = AxiosError<unknown>,
>(
  projectName: 'demo' | 'demo_two',
  params?: GetProjectWorkitemsParams,
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetProjectWorkitemsInfiniteQueryOptions(projectName, params, options)

  const query = useInfiniteQuery(queryOptions) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

export const getGetProjectWorkitemsQueryOptions = <
  TData = Awaited<ReturnType<typeof getProjectWorkitems>>,
  TError = AxiosError<unknown>,
>(
  projectName: 'demo' | 'demo_two',
  params?: GetProjectWorkitemsParams,
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseQueryOptions<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetProjectWorkitemsQueryKey(projectName, params)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getProjectWorkitems>>> = ({ signal }) =>
    getProjectWorkitems(projectName, params, { signal, ...axiosOptions })

  return { queryKey, queryFn, enabled: !!projectName, staleTime: 3600000, ...queryOptions }
}

export type GetProjectWorkitemsQueryResult = NonNullable<Awaited<ReturnType<typeof getProjectWorkitems>>>
export type GetProjectWorkitemsQueryError = AxiosError<unknown>

export const useGetProjectWorkitems = <
  TData = Awaited<ReturnType<typeof getProjectWorkitems>>,
  TError = AxiosError<unknown>,
>(
  projectName: 'demo' | 'demo_two',
  params?: GetProjectWorkitemsParams,
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof getProjectWorkitems>>, TError, TData>
    axios?: AxiosRequestConfig
  },
): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetProjectWorkitemsQueryOptions(projectName, params, options)

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

/**
 * @summary create workitem tag
 */
export const createWorkitemTag = (
  projectName: 'demo' | 'demo_two',
  restWorkItemTagCreateRequest: RestWorkItemTagCreateRequest,
  options?: AxiosRequestConfig,
): Promise<AxiosResponse<DbWorkItemTag>> => {
  return axios.post(`/project/${projectName}/tag/`, restWorkItemTagCreateRequest, options)
}

export const getCreateWorkitemTagMutationOptions = <TError = AxiosError<unknown>, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof createWorkitemTag>>,
    TError,
    { projectName: 'demo' | 'demo_two'; data: RestWorkItemTagCreateRequest },
    TContext
  >
  axios?: AxiosRequestConfig
}): UseMutationOptions<
  Awaited<ReturnType<typeof createWorkitemTag>>,
  TError,
  { projectName: 'demo' | 'demo_two'; data: RestWorkItemTagCreateRequest },
  TContext
> => {
  const { mutation: mutationOptions, axios: axiosOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof createWorkitemTag>>,
    { projectName: 'demo' | 'demo_two'; data: RestWorkItemTagCreateRequest }
  > = (props) => {
    const { projectName, data } = props ?? {}

    return createWorkitemTag(projectName, data, axiosOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type CreateWorkitemTagMutationResult = NonNullable<Awaited<ReturnType<typeof createWorkitemTag>>>
export type CreateWorkitemTagMutationBody = RestWorkItemTagCreateRequest
export type CreateWorkitemTagMutationError = AxiosError<unknown>

export const useCreateWorkitemTag = <TError = AxiosError<unknown>, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof createWorkitemTag>>,
    TError,
    { projectName: 'demo' | 'demo_two'; data: RestWorkItemTagCreateRequest },
    TContext
  >
  axios?: AxiosRequestConfig
}) => {
  const mutationOptions = getCreateWorkitemTagMutationOptions(options)

  return useMutation(mutationOptions)
}
