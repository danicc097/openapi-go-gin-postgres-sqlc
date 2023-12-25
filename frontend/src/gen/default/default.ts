/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import { useInfiniteQuery, useQuery } from '@tanstack/react-query'
import type {
  QueryFunction,
  QueryKey,
  UseInfiniteQueryOptions,
  UseInfiniteQueryResult,
  UseQueryOptions,
  UseQueryResult,
} from '@tanstack/react-query'
import type { HTTPError } from '../model/hTTPError'
import { customInstance } from '../../api/mutator'

// eslint-disable-next-line
type SecondParameter<T extends (...args: any) => any> = T extends (config: any, args: infer P) => any ? P : never

/**
 * @summary Ping pongs
 */
export const ping = (options?: SecondParameter<typeof customInstance>, signal?: AbortSignal) => {
  return customInstance<string>({ url: `/ping`, method: 'GET', signal }, options)
}

export const getPingQueryKey = () => {
  return [`/ping`] as const
}

export const getPingInfiniteQueryOptions = <
  TData = Awaited<ReturnType<typeof ping>>,
  TError = void | HTTPError,
>(options?: {
  query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof ping>>, TError, TData>
  request?: SecondParameter<typeof customInstance>
}) => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getPingQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof ping>>> = ({ signal }) => ping(requestOptions, signal)

  return { queryKey, queryFn, staleTime: 3600000, ...queryOptions } as UseInfiniteQueryOptions<
    Awaited<ReturnType<typeof ping>>,
    TError,
    TData
  > & { queryKey: QueryKey }
}

export type PingInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof ping>>>
export type PingInfiniteQueryError = void | HTTPError

/**
 * @summary Ping pongs
 */
export const usePingInfinite = <TData = Awaited<ReturnType<typeof ping>>, TError = void | HTTPError>(options?: {
  query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof ping>>, TError, TData>
  request?: SecondParameter<typeof customInstance>
}): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getPingInfiniteQueryOptions(options)

  const query = useInfiniteQuery(queryOptions) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

export const getPingQueryOptions = <TData = Awaited<ReturnType<typeof ping>>, TError = void | HTTPError>(options?: {
  query?: UseQueryOptions<Awaited<ReturnType<typeof ping>>, TError, TData>
  request?: SecondParameter<typeof customInstance>
}) => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getPingQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof ping>>> = ({ signal }) => ping(requestOptions, signal)

  return { queryKey, queryFn, staleTime: 3600000, ...queryOptions } as UseQueryOptions<
    Awaited<ReturnType<typeof ping>>,
    TError,
    TData
  > & { queryKey: QueryKey }
}

export type PingQueryResult = NonNullable<Awaited<ReturnType<typeof ping>>>
export type PingQueryError = void | HTTPError

/**
 * @summary Ping pongs
 */
export const usePing = <TData = Awaited<ReturnType<typeof ping>>, TError = void | HTTPError>(options?: {
  query?: UseQueryOptions<Awaited<ReturnType<typeof ping>>, TError, TData>
  request?: SecondParameter<typeof customInstance>
}): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getPingQueryOptions(options)

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

/**
 * @summary Returns this very OpenAPI spec.
 */
export const openapiYamlGet = (options?: SecondParameter<typeof customInstance>, signal?: AbortSignal) => {
  return customInstance<Blob>({ url: `/openapi.yaml`, method: 'GET', responseType: 'blob', signal }, options)
}

export const getOpenapiYamlGetQueryKey = () => {
  return [`/openapi.yaml`] as const
}

export const getOpenapiYamlGetInfiniteQueryOptions = <
  TData = Awaited<ReturnType<typeof openapiYamlGet>>,
  TError = unknown,
>(options?: {
  query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof openapiYamlGet>>, TError, TData>
  request?: SecondParameter<typeof customInstance>
}) => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getOpenapiYamlGetQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof openapiYamlGet>>> = ({ signal }) =>
    openapiYamlGet(requestOptions, signal)

  return { queryKey, queryFn, staleTime: 3600000, ...queryOptions } as UseInfiniteQueryOptions<
    Awaited<ReturnType<typeof openapiYamlGet>>,
    TError,
    TData
  > & { queryKey: QueryKey }
}

export type OpenapiYamlGetInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof openapiYamlGet>>>
export type OpenapiYamlGetInfiniteQueryError = unknown

/**
 * @summary Returns this very OpenAPI spec.
 */
export const useOpenapiYamlGetInfinite = <
  TData = Awaited<ReturnType<typeof openapiYamlGet>>,
  TError = unknown,
>(options?: {
  query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof openapiYamlGet>>, TError, TData>
  request?: SecondParameter<typeof customInstance>
}): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getOpenapiYamlGetInfiniteQueryOptions(options)

  const query = useInfiniteQuery(queryOptions) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

export const getOpenapiYamlGetQueryOptions = <
  TData = Awaited<ReturnType<typeof openapiYamlGet>>,
  TError = unknown,
>(options?: {
  query?: UseQueryOptions<Awaited<ReturnType<typeof openapiYamlGet>>, TError, TData>
  request?: SecondParameter<typeof customInstance>
}) => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getOpenapiYamlGetQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof openapiYamlGet>>> = ({ signal }) =>
    openapiYamlGet(requestOptions, signal)

  return { queryKey, queryFn, staleTime: 3600000, ...queryOptions } as UseQueryOptions<
    Awaited<ReturnType<typeof openapiYamlGet>>,
    TError,
    TData
  > & { queryKey: QueryKey }
}

export type OpenapiYamlGetQueryResult = NonNullable<Awaited<ReturnType<typeof openapiYamlGet>>>
export type OpenapiYamlGetQueryError = unknown

/**
 * @summary Returns this very OpenAPI spec.
 */
export const useOpenapiYamlGet = <TData = Awaited<ReturnType<typeof openapiYamlGet>>, TError = unknown>(options?: {
  query?: UseQueryOptions<Awaited<ReturnType<typeof openapiYamlGet>>, TError, TData>
  request?: SecondParameter<typeof customInstance>
}): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getOpenapiYamlGetQueryOptions(options)

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}
