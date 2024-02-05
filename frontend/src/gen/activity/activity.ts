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
  useMutation,
  useQuery
} from '@tanstack/react-query'
import type {
  MutationFunction,
  QueryFunction,
  QueryKey,
  UseInfiniteQueryOptions,
  UseInfiniteQueryResult,
  UseMutationOptions,
  UseQueryOptions,
  UseQueryResult
} from '@tanstack/react-query'
import type {
  HTTPError
} from '../model/hTTPError'
import type {
  RestActivity
} from '../model/restActivity'
import type {
  RestCreateActivityRequest
} from '../model/restCreateActivityRequest'
import type {
  RestUpdateActivityRequest
} from '../model/restUpdateActivityRequest'
import { customInstance } from '../../api/mutator';


// eslint-disable-next-line
  type SecondParameter<T extends (...args: any) => any> = T extends (
  config: any,
  args: infer P,
) => any
  ? P
  : never;


/**
 * @summary create activity.
 */
export const createActivity = (
    projectName: 'demo' | 'demo_two',
    restCreateActivityRequest: RestCreateActivityRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<RestActivity>(
      {url: `/project/${projectName}/activity/`, method: 'POST',
      headers: {'Content-Type': 'application/json', },
      data: restCreateActivityRequest
    },
      options);
    }
  


export const getCreateActivityMutationOptions = <TError = void | HTTPError,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof createActivity>>, TError,{projectName: 'demo' | 'demo_two';data: RestCreateActivityRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof createActivity>>, TError,{projectName: 'demo' | 'demo_two';data: RestCreateActivityRequest}, TContext> => {
 const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof createActivity>>, {projectName: 'demo' | 'demo_two';data: RestCreateActivityRequest}> = (props) => {
          const {projectName,data} = props ?? {};

          return  createActivity(projectName,data,requestOptions)
        }

        


   return  { mutationFn, ...mutationOptions }}

    export type CreateActivityMutationResult = NonNullable<Awaited<ReturnType<typeof createActivity>>>
    export type CreateActivityMutationBody = RestCreateActivityRequest
    export type CreateActivityMutationError = void | HTTPError

    /**
 * @summary create activity.
 */
export const useCreateActivity = <TError = void | HTTPError,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof createActivity>>, TError,{projectName: 'demo' | 'demo_two';data: RestCreateActivityRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {

      const mutationOptions = getCreateActivityMutationOptions(options);

      return useMutation(mutationOptions);
    }
    /**
 * @summary get activity.
 */
export const getActivity = (
    activityID: EntityIDs.ActivityID,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      
      
      return customInstance<RestActivity>(
      {url: `/activity/${activityID}`, method: 'GET', signal
    },
      options);
    }
  

export const getGetActivityQueryKey = (activityID: EntityIDs.ActivityID,) => {
    return [`/activity/${activityID}`] as const;
    }

    
export const getGetActivityInfiniteQueryOptions = <TData = Awaited<ReturnType<typeof getActivity>>, TError = void | HTTPError>(activityID: EntityIDs.ActivityID, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getActivity>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetActivityQueryKey(activityID);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getActivity>>> = ({ signal }) => getActivity(activityID, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(activityID),  cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true, retry: function(failureCount, error) {
      return failureCount < 3;
    },  ...queryOptions} as UseInfiniteQueryOptions<Awaited<ReturnType<typeof getActivity>>, TError, TData> & { queryKey: QueryKey }
}

export type GetActivityInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getActivity>>>
export type GetActivityInfiniteQueryError = void | HTTPError

/**
 * @summary get activity.
 */
export const useGetActivityInfinite = <TData = Awaited<ReturnType<typeof getActivity>>, TError = void | HTTPError>(
 activityID: EntityIDs.ActivityID, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getActivity>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetActivityInfiniteQueryOptions(activityID,options)

  const query = useInfiniteQuery(queryOptions) as  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



export const getGetActivityQueryOptions = <TData = Awaited<ReturnType<typeof getActivity>>, TError = void | HTTPError>(activityID: EntityIDs.ActivityID, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getActivity>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetActivityQueryKey(activityID);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getActivity>>> = ({ signal }) => getActivity(activityID, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(activityID),  cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true, retry: function(failureCount, error) {
      return failureCount < 3;
    },  ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getActivity>>, TError, TData> & { queryKey: QueryKey }
}

export type GetActivityQueryResult = NonNullable<Awaited<ReturnType<typeof getActivity>>>
export type GetActivityQueryError = void | HTTPError

/**
 * @summary get activity.
 */
export const useGetActivity = <TData = Awaited<ReturnType<typeof getActivity>>, TError = void | HTTPError>(
 activityID: EntityIDs.ActivityID, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getActivity>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetActivityQueryOptions(activityID,options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



/**
 * @summary update activity.
 */
export const updateActivity = (
    activityID: EntityIDs.ActivityID,
    restUpdateActivityRequest: RestUpdateActivityRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<RestActivity>(
      {url: `/activity/${activityID}`, method: 'PATCH',
      headers: {'Content-Type': 'application/json', },
      data: restUpdateActivityRequest
    },
      options);
    }
  


export const getUpdateActivityMutationOptions = <TError = void | HTTPError,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updateActivity>>, TError,{activityID: EntityIDs.ActivityID;data: RestUpdateActivityRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof updateActivity>>, TError,{activityID: EntityIDs.ActivityID;data: RestUpdateActivityRequest}, TContext> => {
 const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof updateActivity>>, {activityID: EntityIDs.ActivityID;data: RestUpdateActivityRequest}> = (props) => {
          const {activityID,data} = props ?? {};

          return  updateActivity(activityID,data,requestOptions)
        }

        


   return  { mutationFn, ...mutationOptions }}

    export type UpdateActivityMutationResult = NonNullable<Awaited<ReturnType<typeof updateActivity>>>
    export type UpdateActivityMutationBody = RestUpdateActivityRequest
    export type UpdateActivityMutationError = void | HTTPError

    /**
 * @summary update activity.
 */
export const useUpdateActivity = <TError = void | HTTPError,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updateActivity>>, TError,{activityID: EntityIDs.ActivityID;data: RestUpdateActivityRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {

      const mutationOptions = getUpdateActivityMutationOptions(options);

      return useMutation(mutationOptions);
    }
    /**
 * @summary delete activity.
 */
export const deleteActivity = (
    activityID: EntityIDs.ActivityID,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<void>(
      {url: `/activity/${activityID}`, method: 'DELETE'
    },
      options);
    }
  


export const getDeleteActivityMutationOptions = <TError = HTTPError,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deleteActivity>>, TError,{activityID: EntityIDs.ActivityID}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof deleteActivity>>, TError,{activityID: EntityIDs.ActivityID}, TContext> => {
 const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof deleteActivity>>, {activityID: EntityIDs.ActivityID}> = (props) => {
          const {activityID} = props ?? {};

          return  deleteActivity(activityID,requestOptions)
        }

        


   return  { mutationFn, ...mutationOptions }}

    export type DeleteActivityMutationResult = NonNullable<Awaited<ReturnType<typeof deleteActivity>>>
    
    export type DeleteActivityMutationError = HTTPError

    /**
 * @summary delete activity.
 */
export const useDeleteActivity = <TError = HTTPError,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deleteActivity>>, TError,{activityID: EntityIDs.ActivityID}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {

      const mutationOptions = getDeleteActivityMutationOptions(options);

      return useMutation(mutationOptions);
    }
    