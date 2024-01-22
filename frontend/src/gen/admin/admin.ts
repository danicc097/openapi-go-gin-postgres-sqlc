import type * as EntityIDs from 'src/gen/entity-ids'
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
export const adminPing = (options?: SecondParameter<typeof customInstance>, signal?: AbortSignal) => {
  return customInstance<string>({ url: `/admin/ping`, method: 'GET', signal }, options)
}

export const getAdminPingQueryKey = () => {
  return [`/admin/ping`] as const
}

export const getAdminPingInfiniteQueryOptions = <
  TData = Awaited<ReturnType<typeof adminPing>>,
  TError = void | HTTPError,
>(options?: {
  query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof adminPing>>, TError, TData>
  request?: SecondParameter<typeof customInstance>
}) => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getAdminPingQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof adminPing>>> = ({ signal }) =>
    adminPing(requestOptions, signal)

  return {
    queryKey,
    queryFn,
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
  } as UseInfiniteQueryOptions<Awaited<ReturnType<typeof adminPing>>, TError, TData> & { queryKey: QueryKey }
}

export type AdminPingInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof adminPing>>>
export type AdminPingInfiniteQueryError = void | HTTPError

/**
 * @summary Ping pongs
 */
export const useAdminPingInfinite = <
  TData = Awaited<ReturnType<typeof adminPing>>,
  TError = void | HTTPError,
>(options?: {
  query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof adminPing>>, TError, TData>
  request?: SecondParameter<typeof customInstance>
}): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getAdminPingInfiniteQueryOptions(options)

  const query = useInfiniteQuery(queryOptions) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

export const getAdminPingQueryOptions = <
  TData = Awaited<ReturnType<typeof adminPing>>,
  TError = void | HTTPError,
>(options?: {
  query?: UseQueryOptions<Awaited<ReturnType<typeof adminPing>>, TError, TData>
  request?: SecondParameter<typeof customInstance>
}) => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getAdminPingQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof adminPing>>> = ({ signal }) =>
    adminPing(requestOptions, signal)

  return {
    queryKey,
    queryFn,
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
  } as UseQueryOptions<Awaited<ReturnType<typeof adminPing>>, TError, TData> & { queryKey: QueryKey }
}

export type AdminPingQueryResult = NonNullable<Awaited<ReturnType<typeof adminPing>>>
export type AdminPingQueryError = void | HTTPError

/**
 * @summary Ping pongs
 */
export const useAdminPing = <TData = Awaited<ReturnType<typeof adminPing>>, TError = void | HTTPError>(options?: {
  query?: UseQueryOptions<Awaited<ReturnType<typeof adminPing>>, TError, TData>
  request?: SecondParameter<typeof customInstance>
}): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getAdminPingQueryOptions(options)

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}
