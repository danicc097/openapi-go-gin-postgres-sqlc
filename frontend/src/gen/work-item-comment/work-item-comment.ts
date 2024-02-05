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
  RestCreateWorkItemCommentRequest
} from '../model/restCreateWorkItemCommentRequest'
import type {
  RestUpdateWorkItemCommentRequest
} from '../model/restUpdateWorkItemCommentRequest'
import type {
  RestWorkItemComment
} from '../model/restWorkItemComment'
import { customInstance } from '../../api/mutator';


// eslint-disable-next-line
  type SecondParameter<T extends (...args: any) => any> = T extends (
  config: any,
  args: infer P,
) => any
  ? P
  : never;


/**
 * @summary create work item comment.
 */
export const createWorkItemComment = (
    workItemID: EntityIDs.WorkItemID,
    restCreateWorkItemCommentRequest: RestCreateWorkItemCommentRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<RestWorkItemComment>(
      {url: `/work-item/${workItemID}/comment/`, method: 'POST',
      headers: {'Content-Type': 'application/json', },
      data: restCreateWorkItemCommentRequest
    },
      options);
    }
  


export const getCreateWorkItemCommentMutationOptions = <TError = void | HTTPError,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof createWorkItemComment>>, TError,{workItemID: EntityIDs.WorkItemID;data: RestCreateWorkItemCommentRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof createWorkItemComment>>, TError,{workItemID: EntityIDs.WorkItemID;data: RestCreateWorkItemCommentRequest}, TContext> => {
 const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof createWorkItemComment>>, {workItemID: EntityIDs.WorkItemID;data: RestCreateWorkItemCommentRequest}> = (props) => {
          const {workItemID,data} = props ?? {};

          return  createWorkItemComment(workItemID,data,requestOptions)
        }

        


   return  { mutationFn, ...mutationOptions }}

    export type CreateWorkItemCommentMutationResult = NonNullable<Awaited<ReturnType<typeof createWorkItemComment>>>
    export type CreateWorkItemCommentMutationBody = RestCreateWorkItemCommentRequest
    export type CreateWorkItemCommentMutationError = void | HTTPError

    /**
 * @summary create work item comment.
 */
export const useCreateWorkItemComment = <TError = void | HTTPError,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof createWorkItemComment>>, TError,{workItemID: EntityIDs.WorkItemID;data: RestCreateWorkItemCommentRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {

      const mutationOptions = getCreateWorkItemCommentMutationOptions(options);

      return useMutation(mutationOptions);
    }
    /**
 * @summary get work item comment.
 */
export const getWorkItemComment = (
    workItemID: EntityIDs.WorkItemID,
    workItemCommentID: EntityIDs.WorkItemCommentID,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      
      
      return customInstance<RestWorkItemComment>(
      {url: `/work-item/${workItemID}/comment/${workItemCommentID}`, method: 'GET', signal
    },
      options);
    }
  

export const getGetWorkItemCommentQueryKey = (workItemID: EntityIDs.WorkItemID,
    workItemCommentID: EntityIDs.WorkItemCommentID,) => {
    return [`/work-item/${workItemID}/comment/${workItemCommentID}`] as const;
    }

    
export const getGetWorkItemCommentInfiniteQueryOptions = <TData = Awaited<ReturnType<typeof getWorkItemComment>>, TError = void | HTTPError>(workItemID: EntityIDs.WorkItemID,
    workItemCommentID: EntityIDs.WorkItemCommentID, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getWorkItemComment>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetWorkItemCommentQueryKey(workItemID,workItemCommentID);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getWorkItemComment>>> = ({ signal }) => getWorkItemComment(workItemID,workItemCommentID, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(workItemID && workItemCommentID),  cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true, retry: function(failureCount, error) {
      return failureCount < 3;
    },  ...queryOptions} as UseInfiniteQueryOptions<Awaited<ReturnType<typeof getWorkItemComment>>, TError, TData> & { queryKey: QueryKey }
}

export type GetWorkItemCommentInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getWorkItemComment>>>
export type GetWorkItemCommentInfiniteQueryError = void | HTTPError

/**
 * @summary get work item comment.
 */
export const useGetWorkItemCommentInfinite = <TData = Awaited<ReturnType<typeof getWorkItemComment>>, TError = void | HTTPError>(
 workItemID: EntityIDs.WorkItemID,
    workItemCommentID: EntityIDs.WorkItemCommentID, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getWorkItemComment>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetWorkItemCommentInfiniteQueryOptions(workItemID,workItemCommentID,options)

  const query = useInfiniteQuery(queryOptions) as  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



export const getGetWorkItemCommentQueryOptions = <TData = Awaited<ReturnType<typeof getWorkItemComment>>, TError = void | HTTPError>(workItemID: EntityIDs.WorkItemID,
    workItemCommentID: EntityIDs.WorkItemCommentID, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getWorkItemComment>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetWorkItemCommentQueryKey(workItemID,workItemCommentID);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getWorkItemComment>>> = ({ signal }) => getWorkItemComment(workItemID,workItemCommentID, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(workItemID && workItemCommentID),  cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true, retry: function(failureCount, error) {
      return failureCount < 3;
    },  ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getWorkItemComment>>, TError, TData> & { queryKey: QueryKey }
}

export type GetWorkItemCommentQueryResult = NonNullable<Awaited<ReturnType<typeof getWorkItemComment>>>
export type GetWorkItemCommentQueryError = void | HTTPError

/**
 * @summary get work item comment.
 */
export const useGetWorkItemComment = <TData = Awaited<ReturnType<typeof getWorkItemComment>>, TError = void | HTTPError>(
 workItemID: EntityIDs.WorkItemID,
    workItemCommentID: EntityIDs.WorkItemCommentID, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getWorkItemComment>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetWorkItemCommentQueryOptions(workItemID,workItemCommentID,options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



/**
 * @summary update work item comment.
 */
export const updateWorkItemComment = (
    workItemID: EntityIDs.WorkItemID,
    workItemCommentID: EntityIDs.WorkItemCommentID,
    restUpdateWorkItemCommentRequest: RestUpdateWorkItemCommentRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<RestWorkItemComment>(
      {url: `/work-item/${workItemID}/comment/${workItemCommentID}`, method: 'PATCH',
      headers: {'Content-Type': 'application/json', },
      data: restUpdateWorkItemCommentRequest
    },
      options);
    }
  


export const getUpdateWorkItemCommentMutationOptions = <TError = void | HTTPError,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updateWorkItemComment>>, TError,{workItemID: EntityIDs.WorkItemID;workItemCommentID: EntityIDs.WorkItemCommentID;data: RestUpdateWorkItemCommentRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof updateWorkItemComment>>, TError,{workItemID: EntityIDs.WorkItemID;workItemCommentID: EntityIDs.WorkItemCommentID;data: RestUpdateWorkItemCommentRequest}, TContext> => {
 const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof updateWorkItemComment>>, {workItemID: EntityIDs.WorkItemID;workItemCommentID: EntityIDs.WorkItemCommentID;data: RestUpdateWorkItemCommentRequest}> = (props) => {
          const {workItemID,workItemCommentID,data} = props ?? {};

          return  updateWorkItemComment(workItemID,workItemCommentID,data,requestOptions)
        }

        


   return  { mutationFn, ...mutationOptions }}

    export type UpdateWorkItemCommentMutationResult = NonNullable<Awaited<ReturnType<typeof updateWorkItemComment>>>
    export type UpdateWorkItemCommentMutationBody = RestUpdateWorkItemCommentRequest
    export type UpdateWorkItemCommentMutationError = void | HTTPError

    /**
 * @summary update work item comment.
 */
export const useUpdateWorkItemComment = <TError = void | HTTPError,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updateWorkItemComment>>, TError,{workItemID: EntityIDs.WorkItemID;workItemCommentID: EntityIDs.WorkItemCommentID;data: RestUpdateWorkItemCommentRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {

      const mutationOptions = getUpdateWorkItemCommentMutationOptions(options);

      return useMutation(mutationOptions);
    }
    /**
 * @summary delete .
 */
export const deleteWorkItemComment = (
    workItemID: EntityIDs.WorkItemID,
    workItemCommentID: EntityIDs.WorkItemCommentID,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<void>(
      {url: `/work-item/${workItemID}/comment/${workItemCommentID}`, method: 'DELETE'
    },
      options);
    }
  


export const getDeleteWorkItemCommentMutationOptions = <TError = HTTPError,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deleteWorkItemComment>>, TError,{workItemID: EntityIDs.WorkItemID;workItemCommentID: EntityIDs.WorkItemCommentID}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof deleteWorkItemComment>>, TError,{workItemID: EntityIDs.WorkItemID;workItemCommentID: EntityIDs.WorkItemCommentID}, TContext> => {
 const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof deleteWorkItemComment>>, {workItemID: EntityIDs.WorkItemID;workItemCommentID: EntityIDs.WorkItemCommentID}> = (props) => {
          const {workItemID,workItemCommentID} = props ?? {};

          return  deleteWorkItemComment(workItemID,workItemCommentID,requestOptions)
        }

        


   return  { mutationFn, ...mutationOptions }}

    export type DeleteWorkItemCommentMutationResult = NonNullable<Awaited<ReturnType<typeof deleteWorkItemComment>>>
    
    export type DeleteWorkItemCommentMutationError = HTTPError

    /**
 * @summary delete .
 */
export const useDeleteWorkItemComment = <TError = HTTPError,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deleteWorkItemComment>>, TError,{workItemID: EntityIDs.WorkItemID;workItemCommentID: EntityIDs.WorkItemCommentID}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {

      const mutationOptions = getDeleteWorkItemCommentMutationOptions(options);

      return useMutation(mutationOptions);
    }
    