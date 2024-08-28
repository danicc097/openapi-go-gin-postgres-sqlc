import type * as EntityIDs from 'src/gen/entity-ids'
/**
 * Generated by orval v6.30.2 🍺
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
  UseMutationResult,
  UseQueryOptions,
  UseQueryResult
} from '@tanstack/react-query'
import type {
  ActivityResponse,
  CreateActivityRequest,
  HTTPError,
  UpdateActivityRequest
} from '.././model'
import { customInstance } from '../../api/mutator';
import type { ErrorType } from '../../api/mutator';


type SecondParameter<T extends (...args: any) => any> = Parameters<T>[1];


/**
 * @summary create activity.
 */
export const createActivity = (
    projectName: 'demo' | 'demo_two',
    createActivityRequest: CreateActivityRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<ActivityResponse>(
      {url: `/project/${projectName}/activity/`, method: 'POST',
      headers: {'Content-Type': 'application/json', },
      data: createActivityRequest
    },
      options);
    }
  


export const getCreateActivityMutationOptions = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof createActivity>>, TError,{projectName: 'demo' | 'demo_two';data: CreateActivityRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof createActivity>>, TError,{projectName: 'demo' | 'demo_two';data: CreateActivityRequest}, TContext> => {
const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof createActivity>>, {projectName: 'demo' | 'demo_two';data: CreateActivityRequest}> = (props) => {
          const {projectName,data} = props ?? {};

          return  createActivity(projectName,data,requestOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type CreateActivityMutationResult = NonNullable<Awaited<ReturnType<typeof createActivity>>>
    export type CreateActivityMutationBody = CreateActivityRequest
    export type CreateActivityMutationError = ErrorType<void | HTTPError>

    /**
 * @summary create activity.
 */
export const useCreateActivity = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof createActivity>>, TError,{projectName: 'demo' | 'demo_two';data: CreateActivityRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationResult<
        Awaited<ReturnType<typeof createActivity>>,
        TError,
        {projectName: 'demo' | 'demo_two';data: CreateActivityRequest},
        TContext
      > => {

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
      
      
      return customInstance<ActivityResponse>(
      {url: `/activity/${activityID}`, method: 'GET', signal
    },
      options);
    }
  

export const getGetActivityQueryKey = (activityID: EntityIDs.ActivityID,) => {
    return [`/activity/${activityID}`] as const;
    }

    
export const getGetActivityInfiniteQueryOptions = <TData = Awaited<ReturnType<typeof getActivity>>, TError = ErrorType<void | HTTPError>>(activityID: EntityIDs.ActivityID, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getActivity>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetActivityQueryKey(activityID);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getActivity>>> = ({ signal }) => getActivity(activityID, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(activityID),  cacheTime: 2000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true,  ...queryOptions} as UseInfiniteQueryOptions<Awaited<ReturnType<typeof getActivity>>, TError, TData> & { queryKey: QueryKey }
}

export type GetActivityInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getActivity>>>
export type GetActivityInfiniteQueryError = ErrorType<void | HTTPError>

/**
 * @summary get activity.
 */
export const useGetActivityInfinite = <TData = Awaited<ReturnType<typeof getActivity>>, TError = ErrorType<void | HTTPError>>(
 activityID: EntityIDs.ActivityID, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getActivity>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetActivityInfiniteQueryOptions(activityID,options)

  const query = useInfiniteQuery(queryOptions) as  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



export const getGetActivityQueryOptions = <TData = Awaited<ReturnType<typeof getActivity>>, TError = ErrorType<void | HTTPError>>(activityID: EntityIDs.ActivityID, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getActivity>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetActivityQueryKey(activityID);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getActivity>>> = ({ signal }) => getActivity(activityID, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(activityID),  cacheTime: 2000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true,  ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getActivity>>, TError, TData> & { queryKey: QueryKey }
}

export type GetActivityQueryResult = NonNullable<Awaited<ReturnType<typeof getActivity>>>
export type GetActivityQueryError = ErrorType<void | HTTPError>

/**
 * @summary get activity.
 */
export const useGetActivity = <TData = Awaited<ReturnType<typeof getActivity>>, TError = ErrorType<void | HTTPError>>(
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
    updateActivityRequest: UpdateActivityRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<ActivityResponse>(
      {url: `/activity/${activityID}`, method: 'PATCH',
      headers: {'Content-Type': 'application/json', },
      data: updateActivityRequest
    },
      options);
    }
  


export const getUpdateActivityMutationOptions = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updateActivity>>, TError,{activityID: EntityIDs.ActivityID;data: UpdateActivityRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof updateActivity>>, TError,{activityID: EntityIDs.ActivityID;data: UpdateActivityRequest}, TContext> => {
const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof updateActivity>>, {activityID: EntityIDs.ActivityID;data: UpdateActivityRequest}> = (props) => {
          const {activityID,data} = props ?? {};

          return  updateActivity(activityID,data,requestOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type UpdateActivityMutationResult = NonNullable<Awaited<ReturnType<typeof updateActivity>>>
    export type UpdateActivityMutationBody = UpdateActivityRequest
    export type UpdateActivityMutationError = ErrorType<void | HTTPError>

    /**
 * @summary update activity.
 */
export const useUpdateActivity = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updateActivity>>, TError,{activityID: EntityIDs.ActivityID;data: UpdateActivityRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationResult<
        Awaited<ReturnType<typeof updateActivity>>,
        TError,
        {activityID: EntityIDs.ActivityID;data: UpdateActivityRequest},
        TContext
      > => {

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
  


export const getDeleteActivityMutationOptions = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deleteActivity>>, TError,{activityID: EntityIDs.ActivityID}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof deleteActivity>>, TError,{activityID: EntityIDs.ActivityID}, TContext> => {
const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof deleteActivity>>, {activityID: EntityIDs.ActivityID}> = (props) => {
          const {activityID} = props ?? {};

          return  deleteActivity(activityID,requestOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type DeleteActivityMutationResult = NonNullable<Awaited<ReturnType<typeof deleteActivity>>>
    
    export type DeleteActivityMutationError = ErrorType<void | HTTPError>

    /**
 * @summary delete activity.
 */
export const useDeleteActivity = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deleteActivity>>, TError,{activityID: EntityIDs.ActivityID}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationResult<
        Awaited<ReturnType<typeof deleteActivity>>,
        TError,
        {activityID: EntityIDs.ActivityID},
        TContext
      > => {

      const mutationOptions = getDeleteActivityMutationOptions(options);

      return useMutation(mutationOptions);
    }
    