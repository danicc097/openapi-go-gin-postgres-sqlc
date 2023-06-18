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
import type { UserResponse, UpdateUserAuthRequest, UpdateUserRequest } from '.././model'

type AwaitedInput<T> = PromiseLike<T> | T

type Awaited<O> = O extends AwaitedInput<infer T> ? T : never

/**
 * @summary returns the logged in user
 */
export const getCurrentUser = (options?: AxiosRequestConfig): Promise<AxiosResponse<UserResponse>> => {
  return axios.get(`/user/me`, options)
}

export const getGetCurrentUserQueryKey = () => [`/user/me`] as const

export const getGetCurrentUserInfiniteQueryOptions = <
  TData = Awaited<ReturnType<typeof getCurrentUser>>,
  TError = AxiosError<unknown>,
>(options?: {
  query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getCurrentUser>>, TError, TData>
  axios?: AxiosRequestConfig
}): UseInfiniteQueryOptions<Awaited<ReturnType<typeof getCurrentUser>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetCurrentUserQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getCurrentUser>>> = ({ signal }) =>
    getCurrentUser({ signal, ...axiosOptions })

  return { queryKey, queryFn, staleTime: 3600000, ...queryOptions }
}

export type GetCurrentUserInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getCurrentUser>>>
export type GetCurrentUserInfiniteQueryError = AxiosError<unknown>

export const useGetCurrentUserInfinite = <
  TData = Awaited<ReturnType<typeof getCurrentUser>>,
  TError = AxiosError<unknown>,
>(options?: {
  query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof getCurrentUser>>, TError, TData>
  axios?: AxiosRequestConfig
}): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetCurrentUserInfiniteQueryOptions(options)

  const query = useInfiniteQuery(queryOptions) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

export const getGetCurrentUserQueryOptions = <
  TData = Awaited<ReturnType<typeof getCurrentUser>>,
  TError = AxiosError<unknown>,
>(options?: {
  query?: UseQueryOptions<Awaited<ReturnType<typeof getCurrentUser>>, TError, TData>
  axios?: AxiosRequestConfig
}): UseQueryOptions<Awaited<ReturnType<typeof getCurrentUser>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, axios: axiosOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetCurrentUserQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getCurrentUser>>> = ({ signal }) =>
    getCurrentUser({ signal, ...axiosOptions })

  return { queryKey, queryFn, staleTime: 3600000, ...queryOptions }
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
  const queryOptions = getGetCurrentUserQueryOptions(options)

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

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

export const getUpdateUserAuthorizationMutationOptions = <TError = AxiosError<unknown>, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof updateUserAuthorization>>,
    TError,
    { id: string; data: UpdateUserAuthRequest },
    TContext
  >
  axios?: AxiosRequestConfig
}): UseMutationOptions<
  Awaited<ReturnType<typeof updateUserAuthorization>>,
  TError,
  { id: string; data: UpdateUserAuthRequest },
  TContext
> => {
  const { mutation: mutationOptions, axios: axiosOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof updateUserAuthorization>>,
    { id: string; data: UpdateUserAuthRequest }
  > = (props) => {
    const { id, data } = props ?? {}

    return updateUserAuthorization(id, data, axiosOptions)
  }

  return { mutationFn, ...mutationOptions }
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
  const mutationOptions = getUpdateUserAuthorizationMutationOptions(options)

  return useMutation(mutationOptions)
}
/**
 * @summary deletes the user by id
 */
export const deleteUser = (id: string, options?: AxiosRequestConfig): Promise<AxiosResponse<unknown>> => {
  return axios.delete(`/user/${id}`, options)
}

export const getDeleteUserMutationOptions = <TError = AxiosError<void>, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<Awaited<ReturnType<typeof deleteUser>>, TError, { id: string }, TContext>
  axios?: AxiosRequestConfig
}): UseMutationOptions<Awaited<ReturnType<typeof deleteUser>>, TError, { id: string }, TContext> => {
  const { mutation: mutationOptions, axios: axiosOptions } = options ?? {}

  const mutationFn: MutationFunction<Awaited<ReturnType<typeof deleteUser>>, { id: string }> = (props) => {
    const { id } = props ?? {}

    return deleteUser(id, axiosOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type DeleteUserMutationResult = NonNullable<Awaited<ReturnType<typeof deleteUser>>>

export type DeleteUserMutationError = AxiosError<void>

export const useDeleteUser = <TError = AxiosError<void>, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<Awaited<ReturnType<typeof deleteUser>>, TError, { id: string }, TContext>
  axios?: AxiosRequestConfig
}) => {
  const mutationOptions = getDeleteUserMutationOptions(options)

  return useMutation(mutationOptions)
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

export const getUpdateUserMutationOptions = <TError = AxiosError<unknown>, TContext = unknown>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof updateUser>>,
    TError,
    { id: string; data: UpdateUserRequest },
    TContext
  >
  axios?: AxiosRequestConfig
}): UseMutationOptions<
  Awaited<ReturnType<typeof updateUser>>,
  TError,
  { id: string; data: UpdateUserRequest },
  TContext
> => {
  const { mutation: mutationOptions, axios: axiosOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof updateUser>>,
    { id: string; data: UpdateUserRequest }
  > = (props) => {
    const { id, data } = props ?? {}

    return updateUser(id, data, axiosOptions)
  }

  return { mutationFn, ...mutationOptions }
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
  const mutationOptions = getUpdateUserMutationOptions(options)

  return useMutation(mutationOptions)
}
