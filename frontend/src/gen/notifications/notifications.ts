import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * OpenAPI openapi-go-gin-postgres-sqlc
 * openapi-go-gin-postgres-sqlc
 * OpenAPI spec version: 2.0.0
 */
import {
  useInfiniteQuery,
  useQuery
} from '@tanstack/react-query'
import type {
  QueryFunction,
  QueryKey,
  UseInfiniteQueryOptions,
  UseInfiniteQueryResult,
  UseQueryOptions,
  UseQueryResult
} from '@tanstack/react-query'
import type {
  GetPaginatedNotificationsParams
} from '../model/getPaginatedNotificationsParams'
import type {
  HTTPError
} from '../model/hTTPError'
import type {
  RestPaginatedNotificationsResponse
} from '../model/restPaginatedNotificationsResponse'
import { customInstance } from '../../api/mutator';


// eslint-disable-next-line
  type SecondParameter<T extends (...args: any) => any> = T extends (
  config: any,
  args: infer P,
) => any
  ? P
  : never;


/**
 * @summary Get paginated user notifications
 */
export const getPaginatedNotifications = (
    params: GetPaginatedNotificationsParams,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      
      
      return customInstance<RestPaginatedNotificationsResponse>(
      {url: `/notifications/user/page`, method: 'GET',
        params, signal
    },
      options);
    }
  

export const getGetPaginatedNotificationsQueryKey = (params: GetPaginatedNotificationsParams,) => {
    return [`/notifications/user/page`, ...(params ? [params]: [])] as const;
    }

    
export const getGetPaginatedNotificationsInfiniteQueryOptions = <TData = Awaited<ReturnType<typeof getPaginatedNotifications>>, TError = void | HTTPError>(params: GetPaginatedNotificationsParams, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getPaginatedNotifications>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetPaginatedNotificationsQueryKey(params);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getPaginatedNotifications>>> = ({ signal }) => getPaginatedNotifications(params, requestOptions, signal);

      

      

   return  { queryKey, queryFn,   cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true, retry: function(failureCount, error) {
      return failureCount < 3;
    },  ...queryOptions} as UseInfiniteQueryOptions<Awaited<ReturnType<typeof getPaginatedNotifications>>, TError, TData> & { queryKey: QueryKey }
}

export type GetPaginatedNotificationsInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getPaginatedNotifications>>>
export type GetPaginatedNotificationsInfiniteQueryError = void | HTTPError

/**
 * @summary Get paginated user notifications
 */
export const useGetPaginatedNotificationsInfinite = <TData = Awaited<ReturnType<typeof getPaginatedNotifications>>, TError = void | HTTPError>(
 params: GetPaginatedNotificationsParams, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getPaginatedNotifications>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetPaginatedNotificationsInfiniteQueryOptions(params,options)

  const query = useInfiniteQuery(queryOptions) as  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



export const getGetPaginatedNotificationsQueryOptions = <TData = Awaited<ReturnType<typeof getPaginatedNotifications>>, TError = void | HTTPError>(params: GetPaginatedNotificationsParams, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getPaginatedNotifications>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetPaginatedNotificationsQueryKey(params);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getPaginatedNotifications>>> = ({ signal }) => getPaginatedNotifications(params, requestOptions, signal);

      

      

   return  { queryKey, queryFn,   cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true, retry: function(failureCount, error) {
      return failureCount < 3;
    },  ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getPaginatedNotifications>>, TError, TData> & { queryKey: QueryKey }
}

export type GetPaginatedNotificationsQueryResult = NonNullable<Awaited<ReturnType<typeof getPaginatedNotifications>>>
export type GetPaginatedNotificationsQueryError = void | HTTPError

/**
 * @summary Get paginated user notifications
 */
export const useGetPaginatedNotifications = <TData = Awaited<ReturnType<typeof getPaginatedNotifications>>, TError = void | HTTPError>(
 params: GetPaginatedNotificationsParams, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getPaginatedNotifications>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetPaginatedNotificationsQueryOptions(params,options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



