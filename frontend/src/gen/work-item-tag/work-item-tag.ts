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
  CreateWorkItemTagRequest,
  HTTPError,
  UpdateWorkItemTagRequest,
  WorkItemTagResponse
} from '.././model'
import { customInstance } from '../../api/mutator';
import type { ErrorType } from '../../api/mutator';


type SecondParameter<T extends (...args: any) => any> = Parameters<T>[1];


/**
 * @summary create workitemtag.
 */
export const createWorkItemTag = (
    projectName: 'demo' | 'demo_two',
    createWorkItemTagRequest: CreateWorkItemTagRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<WorkItemTagResponse>(
      {url: `/project/${projectName}/work-item-tag/`, method: 'POST',
      headers: {'Content-Type': 'application/json', },
      data: createWorkItemTagRequest
    },
      options);
    }
  


export const getCreateWorkItemTagMutationOptions = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof createWorkItemTag>>, TError,{projectName: 'demo' | 'demo_two';data: CreateWorkItemTagRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof createWorkItemTag>>, TError,{projectName: 'demo' | 'demo_two';data: CreateWorkItemTagRequest}, TContext> => {
const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof createWorkItemTag>>, {projectName: 'demo' | 'demo_two';data: CreateWorkItemTagRequest}> = (props) => {
          const {projectName,data} = props ?? {};

          return  createWorkItemTag(projectName,data,requestOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type CreateWorkItemTagMutationResult = NonNullable<Awaited<ReturnType<typeof createWorkItemTag>>>
    export type CreateWorkItemTagMutationBody = CreateWorkItemTagRequest
    export type CreateWorkItemTagMutationError = ErrorType<void | HTTPError>

    /**
 * @summary create workitemtag.
 */
export const useCreateWorkItemTag = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof createWorkItemTag>>, TError,{projectName: 'demo' | 'demo_two';data: CreateWorkItemTagRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationResult<
        Awaited<ReturnType<typeof createWorkItemTag>>,
        TError,
        {projectName: 'demo' | 'demo_two';data: CreateWorkItemTagRequest},
        TContext
      > => {

      const mutationOptions = getCreateWorkItemTagMutationOptions(options);

      return useMutation(mutationOptions);
    }
    /**
 * @summary get workitemtag.
 */
export const getWorkItemTag = (
    workItemTagID: EntityIDs.WorkItemTagID,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      
      
      return customInstance<WorkItemTagResponse>(
      {url: `/work-item-tag/${workItemTagID}`, method: 'GET', signal
    },
      options);
    }
  

export const getGetWorkItemTagQueryKey = (workItemTagID: EntityIDs.WorkItemTagID,) => {
    return [`/work-item-tag/${workItemTagID}`] as const;
    }

    
export const getGetWorkItemTagInfiniteQueryOptions = <TData = Awaited<ReturnType<typeof getWorkItemTag>>, TError = ErrorType<void | HTTPError>>(workItemTagID: EntityIDs.WorkItemTagID, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getWorkItemTag>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetWorkItemTagQueryKey(workItemTagID);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getWorkItemTag>>> = ({ signal }) => getWorkItemTag(workItemTagID, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(workItemTagID),  cacheTime: 2000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true,  ...queryOptions} as UseInfiniteQueryOptions<Awaited<ReturnType<typeof getWorkItemTag>>, TError, TData> & { queryKey: QueryKey }
}

export type GetWorkItemTagInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getWorkItemTag>>>
export type GetWorkItemTagInfiniteQueryError = ErrorType<void | HTTPError>

/**
 * @summary get workitemtag.
 */
export const useGetWorkItemTagInfinite = <TData = Awaited<ReturnType<typeof getWorkItemTag>>, TError = ErrorType<void | HTTPError>>(
 workItemTagID: EntityIDs.WorkItemTagID, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getWorkItemTag>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetWorkItemTagInfiniteQueryOptions(workItemTagID,options)

  const query = useInfiniteQuery(queryOptions) as  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



export const getGetWorkItemTagQueryOptions = <TData = Awaited<ReturnType<typeof getWorkItemTag>>, TError = ErrorType<void | HTTPError>>(workItemTagID: EntityIDs.WorkItemTagID, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getWorkItemTag>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetWorkItemTagQueryKey(workItemTagID);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getWorkItemTag>>> = ({ signal }) => getWorkItemTag(workItemTagID, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(workItemTagID),  cacheTime: 2000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true,  ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getWorkItemTag>>, TError, TData> & { queryKey: QueryKey }
}

export type GetWorkItemTagQueryResult = NonNullable<Awaited<ReturnType<typeof getWorkItemTag>>>
export type GetWorkItemTagQueryError = ErrorType<void | HTTPError>

/**
 * @summary get workitemtag.
 */
export const useGetWorkItemTag = <TData = Awaited<ReturnType<typeof getWorkItemTag>>, TError = ErrorType<void | HTTPError>>(
 workItemTagID: EntityIDs.WorkItemTagID, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getWorkItemTag>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetWorkItemTagQueryOptions(workItemTagID,options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



/**
 * @summary update workitemtag.
 */
export const updateWorkItemTag = (
    workItemTagID: EntityIDs.WorkItemTagID,
    updateWorkItemTagRequest: UpdateWorkItemTagRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<WorkItemTagResponse>(
      {url: `/work-item-tag/${workItemTagID}`, method: 'PATCH',
      headers: {'Content-Type': 'application/json', },
      data: updateWorkItemTagRequest
    },
      options);
    }
  


export const getUpdateWorkItemTagMutationOptions = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updateWorkItemTag>>, TError,{workItemTagID: EntityIDs.WorkItemTagID;data: UpdateWorkItemTagRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof updateWorkItemTag>>, TError,{workItemTagID: EntityIDs.WorkItemTagID;data: UpdateWorkItemTagRequest}, TContext> => {
const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof updateWorkItemTag>>, {workItemTagID: EntityIDs.WorkItemTagID;data: UpdateWorkItemTagRequest}> = (props) => {
          const {workItemTagID,data} = props ?? {};

          return  updateWorkItemTag(workItemTagID,data,requestOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type UpdateWorkItemTagMutationResult = NonNullable<Awaited<ReturnType<typeof updateWorkItemTag>>>
    export type UpdateWorkItemTagMutationBody = UpdateWorkItemTagRequest
    export type UpdateWorkItemTagMutationError = ErrorType<void | HTTPError>

    /**
 * @summary update workitemtag.
 */
export const useUpdateWorkItemTag = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updateWorkItemTag>>, TError,{workItemTagID: EntityIDs.WorkItemTagID;data: UpdateWorkItemTagRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationResult<
        Awaited<ReturnType<typeof updateWorkItemTag>>,
        TError,
        {workItemTagID: EntityIDs.WorkItemTagID;data: UpdateWorkItemTagRequest},
        TContext
      > => {

      const mutationOptions = getUpdateWorkItemTagMutationOptions(options);

      return useMutation(mutationOptions);
    }
    /**
 * @summary delete workitemtag.
 */
export const deleteWorkItemTag = (
    workItemTagID: EntityIDs.WorkItemTagID,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<void>(
      {url: `/work-item-tag/${workItemTagID}`, method: 'DELETE'
    },
      options);
    }
  


export const getDeleteWorkItemTagMutationOptions = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deleteWorkItemTag>>, TError,{workItemTagID: EntityIDs.WorkItemTagID}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof deleteWorkItemTag>>, TError,{workItemTagID: EntityIDs.WorkItemTagID}, TContext> => {
const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof deleteWorkItemTag>>, {workItemTagID: EntityIDs.WorkItemTagID}> = (props) => {
          const {workItemTagID} = props ?? {};

          return  deleteWorkItemTag(workItemTagID,requestOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type DeleteWorkItemTagMutationResult = NonNullable<Awaited<ReturnType<typeof deleteWorkItemTag>>>
    
    export type DeleteWorkItemTagMutationError = ErrorType<void | HTTPError>

    /**
 * @summary delete workitemtag.
 */
export const useDeleteWorkItemTag = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deleteWorkItemTag>>, TError,{workItemTagID: EntityIDs.WorkItemTagID}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationResult<
        Awaited<ReturnType<typeof deleteWorkItemTag>>,
        TError,
        {workItemTagID: EntityIDs.WorkItemTagID},
        TContext
      > => {

      const mutationOptions = getDeleteWorkItemTagMutationOptions(options);

      return useMutation(mutationOptions);
    }
    