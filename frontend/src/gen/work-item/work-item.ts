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
  CreateWorkItemRequest
} from '../model/createWorkItemRequest'
import type {
  DbWorkItem
} from '../model/dbWorkItem'
import { customInstance } from '../../api/mutator';


// eslint-disable-next-line
  type SecondParameter<T extends (...args: any) => any> = T extends (
  config: any,
  args: infer P,
) => any
  ? P
  : never;


/**
 * @summary create workitem
 */
export const createWorkitem = (
    createWorkItemRequest: CreateWorkItemRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<DbWorkItem>(
      {url: `/work-item/`, method: 'POST',
      headers: {'Content-Type': 'application/json', },
      data: createWorkItemRequest
    },
      options);
    }
  


export const getCreateWorkitemMutationOptions = <TError = unknown,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof createWorkitem>>, TError,{data: CreateWorkItemRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof createWorkitem>>, TError,{data: CreateWorkItemRequest}, TContext> => {
 const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof createWorkitem>>, {data: CreateWorkItemRequest}> = (props) => {
          const {data} = props ?? {};

          return  createWorkitem(data,requestOptions)
        }

        


   return  { mutationFn, ...mutationOptions }}

    export type CreateWorkitemMutationResult = NonNullable<Awaited<ReturnType<typeof createWorkitem>>>
    export type CreateWorkitemMutationBody = CreateWorkItemRequest
    export type CreateWorkitemMutationError = unknown

    /**
 * @summary create workitem
 */
export const useCreateWorkitem = <TError = unknown,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof createWorkitem>>, TError,{data: CreateWorkItemRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {

      const mutationOptions = getCreateWorkitemMutationOptions(options);

      return useMutation(mutationOptions);
    }
    /**
 * @summary get workitem
 */
export const getWorkItem = (
    workItemID: EntityIDs.WorkItemID,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      
      
      return customInstance<DbWorkItem>(
      {url: `/work-item/${workItemID}/`, method: 'GET', signal
    },
      options);
    }
  

export const getGetWorkItemQueryKey = (workItemID: EntityIDs.WorkItemID,) => {
    return [`/work-item/${workItemID}/`] as const;
    }

    
export const getGetWorkItemInfiniteQueryOptions = <TData = Awaited<ReturnType<typeof getWorkItem>>, TError = unknown>(workItemID: EntityIDs.WorkItemID, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getWorkItem>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetWorkItemQueryKey(workItemID);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getWorkItem>>> = ({ signal }) => getWorkItem(workItemID, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(workItemID),  cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true, retry: function(failureCount, error) {
      return failureCount < 3;
    },  ...queryOptions} as UseInfiniteQueryOptions<Awaited<ReturnType<typeof getWorkItem>>, TError, TData> & { queryKey: QueryKey }
}

export type GetWorkItemInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getWorkItem>>>
export type GetWorkItemInfiniteQueryError = unknown

/**
 * @summary get workitem
 */
export const useGetWorkItemInfinite = <TData = Awaited<ReturnType<typeof getWorkItem>>, TError = unknown>(
 workItemID: EntityIDs.WorkItemID, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getWorkItem>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetWorkItemInfiniteQueryOptions(workItemID,options)

  const query = useInfiniteQuery(queryOptions) as  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



export const getGetWorkItemQueryOptions = <TData = Awaited<ReturnType<typeof getWorkItem>>, TError = unknown>(workItemID: EntityIDs.WorkItemID, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getWorkItem>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetWorkItemQueryKey(workItemID);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getWorkItem>>> = ({ signal }) => getWorkItem(workItemID, requestOptions, signal);

      

      

   return  { queryKey, queryFn, enabled: !!(workItemID),  cacheTime: 300000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true, retry: function(failureCount, error) {
      return failureCount < 3;
    },  ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getWorkItem>>, TError, TData> & { queryKey: QueryKey }
}

export type GetWorkItemQueryResult = NonNullable<Awaited<ReturnType<typeof getWorkItem>>>
export type GetWorkItemQueryError = unknown

/**
 * @summary get workitem
 */
export const useGetWorkItem = <TData = Awaited<ReturnType<typeof getWorkItem>>, TError = unknown>(
 workItemID: EntityIDs.WorkItemID, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getWorkItem>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetWorkItemQueryOptions(workItemID,options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



/**
 * @summary update workitem
 */
export const updateWorkitem = (
    workItemID: EntityIDs.WorkItemID,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<DbWorkItem>(
      {url: `/work-item/${workItemID}/`, method: 'PATCH'
    },
      options);
    }
  


export const getUpdateWorkitemMutationOptions = <TError = unknown,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updateWorkitem>>, TError,{workItemID: EntityIDs.WorkItemID}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof updateWorkitem>>, TError,{workItemID: EntityIDs.WorkItemID}, TContext> => {
 const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof updateWorkitem>>, {workItemID: EntityIDs.WorkItemID}> = (props) => {
          const {workItemID} = props ?? {};

          return  updateWorkitem(workItemID,requestOptions)
        }

        


   return  { mutationFn, ...mutationOptions }}

    export type UpdateWorkitemMutationResult = NonNullable<Awaited<ReturnType<typeof updateWorkitem>>>
    
    export type UpdateWorkitemMutationError = unknown

    /**
 * @summary update workitem
 */
export const useUpdateWorkitem = <TError = unknown,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updateWorkitem>>, TError,{workItemID: EntityIDs.WorkItemID}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {

      const mutationOptions = getUpdateWorkitemMutationOptions(options);

      return useMutation(mutationOptions);
    }
    /**
 * @summary delete workitem
 */
export const deleteWorkitem = (
    workItemID: EntityIDs.WorkItemID,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<void>(
      {url: `/work-item/${workItemID}/`, method: 'DELETE'
    },
      options);
    }
  


export const getDeleteWorkitemMutationOptions = <TError = unknown,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deleteWorkitem>>, TError,{workItemID: EntityIDs.WorkItemID}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof deleteWorkitem>>, TError,{workItemID: EntityIDs.WorkItemID}, TContext> => {
 const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof deleteWorkitem>>, {workItemID: EntityIDs.WorkItemID}> = (props) => {
          const {workItemID} = props ?? {};

          return  deleteWorkitem(workItemID,requestOptions)
        }

        


   return  { mutationFn, ...mutationOptions }}

    export type DeleteWorkitemMutationResult = NonNullable<Awaited<ReturnType<typeof deleteWorkitem>>>
    
    export type DeleteWorkitemMutationError = unknown

    /**
 * @summary delete workitem
 */
export const useDeleteWorkitem = <TError = unknown,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deleteWorkitem>>, TError,{workItemID: EntityIDs.WorkItemID}, TContext>, request?: SecondParameter<typeof customInstance>}
) => {

      const mutationOptions = getDeleteWorkitemMutationOptions(options);

      return useMutation(mutationOptions);
    }
    