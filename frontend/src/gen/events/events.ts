/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import { useQuery, useInfiniteQuery } from '@tanstack/react-query'
import type {
  UseQueryOptions,
  UseInfiniteQueryOptions,
  QueryFunction,
  UseQueryResult,
  UseInfiniteQueryResult,
  QueryKey,
} from '@tanstack/react-query'
import type { EventsParams } from '.././model'
import { customInstance } from '../../api/mutator'

type AwaitedInput<T> = PromiseLike<T> | T

type Awaited<O> = O extends AwaitedInput<infer T> ? T : never

// eslint-disable-next-line
type SecondParameter<T extends (...args: any) => any> = T extends (config: any, args: infer P) => any ? P : never

export const events = (
  params: EventsParams,
  options?: SecondParameter<typeof customInstance>,
  signal?: AbortSignal,
) => {
  return customInstance<string>({ url: `/events`, method: 'get', params, signal }, options)
}

export const getEventsQueryKey = (params: EventsParams) => [`/events`, ...(params ? [params] : [])] as const

export const getEventsInfiniteQueryOptions = <TData = Awaited<ReturnType<typeof events>>, TError = unknown>(
  params: EventsParams,
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof events>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseInfiniteQueryOptions<Awaited<ReturnType<typeof events>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getEventsQueryKey(params)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof events>>> = ({ signal }) =>
    events(params, requestOptions, signal)

  return { queryKey, queryFn, staleTime: 3600000, ...queryOptions }
}

export type EventsInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof events>>>
export type EventsInfiniteQueryError = unknown

export const useEventsInfinite = <TData = Awaited<ReturnType<typeof events>>, TError = unknown>(
  params: EventsParams,
  options?: {
    query?: UseInfiniteQueryOptions<Awaited<ReturnType<typeof events>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getEventsInfiniteQueryOptions(params, options)

  const query = useInfiniteQuery(queryOptions) as UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}

export const getEventsQueryOptions = <TData = Awaited<ReturnType<typeof events>>, TError = unknown>(
  params: EventsParams,
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof events>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseQueryOptions<Awaited<ReturnType<typeof events>>, TError, TData> & { queryKey: QueryKey } => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getEventsQueryKey(params)

  const queryFn: QueryFunction<Awaited<ReturnType<typeof events>>> = ({ signal }) =>
    events(params, requestOptions, signal)

  return { queryKey, queryFn, staleTime: 3600000, ...queryOptions }
}

export type EventsQueryResult = NonNullable<Awaited<ReturnType<typeof events>>>
export type EventsQueryError = unknown

export const useEvents = <TData = Awaited<ReturnType<typeof events>>, TError = unknown>(
  params: EventsParams,
  options?: {
    query?: UseQueryOptions<Awaited<ReturnType<typeof events>>, TError, TData>
    request?: SecondParameter<typeof customInstance>
  },
): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getEventsQueryOptions(params, options)

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & { queryKey: QueryKey }

  query.queryKey = queryOptions.queryKey

  return query
}