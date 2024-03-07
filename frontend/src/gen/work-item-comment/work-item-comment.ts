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
  CreateWorkItemCommentRequest
} from '../model/createWorkItemCommentRequest'
import type {
  HTTPError
} from '../model/hTTPError'
import type {
  UpdateWorkItemCommentRequest
} from '../model/updateWorkItemCommentRequest'
import type {
  WorkItemComment
} from '../model/workItemComment'
import { customInstance } from '../../api/mutator';
import type { ErrorType } from '../../api/mutator';


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
    createWorkItemCommentRequest: CreateWorkItemCommentRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<WorkItemComment>(
      {url: `/work-item/${workItemID}/comment/`, method: 'POST',
      headers: {'Content-Type': 'application/json', },
      data: createWorkItemCommentRequest
    },
      options);
    }
  


export const getCreateWorkItemCommentMutationOptions = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof createWorkItemComment>>, TError,{workItemID: EntityIDs.WorkItemID;data: CreateWorkItemCommentRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof createWorkItemComment>>, TError,{workItemID: EntityIDs.WorkItemID;data: CreateWorkItemCommentRequest}, TContext> => {
 const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof createWorkItemComment>>, {workItemID: EntityIDs.WorkItemID;data: CreateWorkItemCommentRequest}> = (props) => {
          const {workItemID,data} = props ?? {};

          return  createWorkItemComment(workItemID,data,requestOptions)
        }

        


   return  { mutationFn, ...mutationOptions }}

    export type CreateWorkItemCommentMutationResult = NonNullable<Awaited<ReturnType<typeof createWorkItemComment>>>
    export type CreateWorkItemCommentMutationBody = CreateWorkItemCommentRequest
    export type CreateWorkItemCommentMutationError = ErrorType<void | HTTPError>

    /**
 * @summary create work item comment.
 */
export const useCreateWorkItemComment = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof createWorkItemComment>>, TError,{workItemID: EntityIDs.WorkItemID;data: CreateWorkItemCommentRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
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
      
      
      return customInstance<WorkItemComment>(
      {url: `/work-item/${workItemID}/comment/${workItemCommentID}`, method: 'GET', signal
    },
      options);
    }
  

export const getGetWorkItemCommentQueryKey = (workItemID: EntityIDs.WorkItemID,
    workItemCommentID: EntityIDs.WorkItemCommentID,) => {
    return [`/work-item/${workItemID}/comment/${workItemCommentID}`] as const;
    }

    
export const getGetWorkItemCommentInfiniteQueryOptions = <TData = Awaited<ReturnType<typeof getWorkItemComment>>, TError = ErrorType<void | HTTPError>>(workItemID: EntityIDs.WorkItemID,
    workItemCommentID: EntityIDs.WorkItemCommentID, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getWorkItemComment>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetWorkItemCommentQueryKey(workItemID,workItemCommentID);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getWorkItemComment>>> = ({ signal }) => getWorkItemComment(workItemID,workItemCommentID, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(workItemID && workItemCommentID),  cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true,  ...queryOptions} as UseInfiniteQueryOptions<Awaited<ReturnType<typeof getWorkItemComment>>, TError, TData> & { queryKey: QueryKey }
}

export type GetWorkItemCommentInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getWorkItemComment>>>
export type GetWorkItemCommentInfiniteQueryError = ErrorType<void | HTTPError>

/**
 * @summary get work item comment.
 */
export const useGetWorkItemCommentInfinite = <TData = Awaited<ReturnType<typeof getWorkItemComment>>, TError = ErrorType<void | HTTPError>>(
 workItemID: EntityIDs.WorkItemID,
    workItemCommentID: EntityIDs.WorkItemCommentID, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getWorkItemComment>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetWorkItemCommentInfiniteQueryOptions(workItemID,workItemCommentID,options)

  const query = useInfiniteQuery(queryOptions) as  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



export const getGetWorkItemCommentQueryOptions = <TData = Awaited<ReturnType<typeof getWorkItemComment>>, TError = ErrorType<void | HTTPError>>(workItemID: EntityIDs.WorkItemID,
    workItemCommentID: EntityIDs.WorkItemCommentID, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getWorkItemComment>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetWorkItemCommentQueryKey(workItemID,workItemCommentID);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getWorkItemComment>>> = ({ signal }) => getWorkItemComment(workItemID,workItemCommentID, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(workItemID && workItemCommentID),  cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true,  ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getWorkItemComment>>, TError, TData> & { queryKey: QueryKey }
}

export type GetWorkItemCommentQueryResult = NonNullable<Awaited<ReturnType<typeof getWorkItemComment>>>
export type GetWorkItemCommentQueryError = ErrorType<void | HTTPError>

/**
 * @summary get work item comment.
 */
export const useGetWorkItemComment = <TData = Awaited<ReturnType<typeof getWorkItemComment>>, TError = ErrorType<void | HTTPError>>(
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
    updateWorkItemCommentRequest: UpdateWorkItemCommentRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<WorkItemComment>(
      {url: `/work-item/${workItemID}/comment/${workItemCommentID}`, method: 'PATCH',
      headers: {'Content-Type': 'application/json', },
      data: updateWorkItemCommentRequest
    },
      options);
    }
  


export const getUpdateWorkItemCommentMutationOptions = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updateWorkItemComment>>, TError,{workItemID: EntityIDs.WorkItemID;workItemCommentID: EntityIDs.WorkItemCommentID;data: UpdateWorkItemCommentRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof updateWorkItemComment>>, TError,{workItemID: EntityIDs.WorkItemID;workItemCommentID: EntityIDs.WorkItemCommentID;data: UpdateWorkItemCommentRequest}, TContext> => {
 const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof updateWorkItemComment>>, {workItemID: EntityIDs.WorkItemID;workItemCommentID: EntityIDs.WorkItemCommentID;data: UpdateWorkItemCommentRequest}> = (props) => {
          const {workItemID,workItemCommentID,data} = props ?? {};

          return  updateWorkItemComment(workItemID,workItemCommentID,data,requestOptions)
        }

        


   return  { mutationFn, ...mutationOptions }}

    export type UpdateWorkItemCommentMutationResult = NonNullable<Awaited<ReturnType<typeof updateWorkItemComment>>>
    export type UpdateWorkItemCommentMutationBody = UpdateWorkItemCommentRequest
    export type UpdateWorkItemCommentMutationError = ErrorType<void | HTTPError>

    /**
 * @summary update work item comment.
 */
export const useUpdateWorkItemComment = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updateWorkItemComment>>, TError,{workItemID: EntityIDs.WorkItemID;workItemCommentID: EntityIDs.WorkItemCommentID;data: UpdateWorkItemCommentRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
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
  


export const getDeleteWorkItemCommentMutationOptions = <TError = ErrorType<HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deleteWorkItemComment>>, TError,{workItemID: EntityIDs.WorkItemID;workItemCommentID: EntityIDs.WorkItemCommentID}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof deleteWorkItemComment>>, TError,{workItemID: EntityIDs.WorkItemID;workItemCommentID: EntityIDs.WorkItemCommentID}, TContext> => {
 const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof deleteWorkItemComment>>, {workItemID: EntityIDs.WorkItemID;workItemCommentID: EntityIDs.WorkItemCommentID}> = (props) => {
          const {workItemID,workItemCommentID} = props ?? {};

          return  deleteWorkItemComment(workItemID,workItemCommentID,requestOptions)
        }

        


   return  { mutationFn, ...mutationOptions }}

    export type DeleteWorkItemCommentMutationResult = NonNullable<Awaited<ReturnType<typeof deleteWorkItemComment>>>
    
    export type DeleteWorkItemCommentMutationError = ErrorType<HTTPError>

    /**
 * @summary delete .
 */
export const useDeleteWorkItemComment = <TError = ErrorType<HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deleteWorkItemComment>>, TError,{workItemID: EntityIDs.WorkItemID;workItemCommentID: EntityIDs.WorkItemCommentID}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {

      const mutationOptions = getDeleteWorkItemCommentMutationOptions(options);

      return useMutation(mutationOptions);
    }
    