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
  GetPaginatedUsersParams,
  HTTPError,
  PaginatedUsersResponse,
  UpdateUserAuthRequest,
  UpdateUserRequest,
  UserResponse
} from '.././model'
import { customInstance } from '../../api/mutator';
import type { ErrorType } from '../../api/mutator';


type SecondParameter<T extends (...args: any) => any> = Parameters<T>[1];


/**
 * @summary Get paginated users
 */
export const getPaginatedUsers = (
    params: GetPaginatedUsersParams,
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      
      
      return customInstance<PaginatedUsersResponse>(
      {url: `/user/page`, method: 'GET',
        params, signal
    },
      options);
    }
  

export const getGetPaginatedUsersQueryKey = (params: GetPaginatedUsersParams,) => {
    return [`/user/page`, ...(params ? [params]: [])] as const;
    }

    
export const getGetPaginatedUsersInfiniteQueryOptions = <TData = Awaited<ReturnType<typeof getPaginatedUsers>>, TError = ErrorType<void | HTTPError>>(params: GetPaginatedUsersParams, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getPaginatedUsers>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetPaginatedUsersQueryKey(params);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getPaginatedUsers>>> = ({ signal, pageParam }) => getPaginatedUsers({...params, cursor: pageParam || params?.['cursor']}, requestOptions, signal);

      

      

   return  { queryKey, queryFn,   cacheTime: 2000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true,  ...queryOptions} as UseInfiniteQueryOptions<Awaited<ReturnType<typeof getPaginatedUsers>>, TError, TData> & { queryKey: QueryKey }
}

export type GetPaginatedUsersInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getPaginatedUsers>>>
export type GetPaginatedUsersInfiniteQueryError = ErrorType<void | HTTPError>

/**
 * @summary Get paginated users
 */
export const useGetPaginatedUsersInfinite = <TData = Awaited<ReturnType<typeof getPaginatedUsers>>, TError = ErrorType<void | HTTPError>>(
 params: GetPaginatedUsersParams, options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getPaginatedUsers>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetPaginatedUsersInfiniteQueryOptions(params,options)

  const query = useInfiniteQuery(queryOptions) as  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



export const getGetPaginatedUsersQueryOptions = <TData = Awaited<ReturnType<typeof getPaginatedUsers>>, TError = ErrorType<void | HTTPError>>(params: GetPaginatedUsersParams, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getPaginatedUsers>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetPaginatedUsersQueryKey(params);

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getPaginatedUsers>>> = ({ signal }) => getPaginatedUsers(params, requestOptions, signal);

      

      

   return  { queryKey, queryFn,   cacheTime: 2000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true,  ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getPaginatedUsers>>, TError, TData> & { queryKey: QueryKey }
}

export type GetPaginatedUsersQueryResult = NonNullable<Awaited<ReturnType<typeof getPaginatedUsers>>>
export type GetPaginatedUsersQueryError = ErrorType<void | HTTPError>

/**
 * @summary Get paginated users
 */
export const useGetPaginatedUsers = <TData = Awaited<ReturnType<typeof getPaginatedUsers>>, TError = ErrorType<void | HTTPError>>(
 params: GetPaginatedUsersParams, options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getPaginatedUsers>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetPaginatedUsersQueryOptions(params,options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



/**
 * @summary returns the logged in user
 */
export const getCurrentUser = (
    
 options?: SecondParameter<typeof customInstance>,signal?: AbortSignal
) => {
      
      
      return customInstance<UserResponse>(
      {url: `/user/me`, method: 'GET', signal
    },
      options);
    }
  

export const getGetCurrentUserQueryKey = () => {
    return [`/user/me`] as const;
    }

    
export const getGetCurrentUserInfiniteQueryOptions = <TData = Awaited<ReturnType<typeof getCurrentUser>>, TError = ErrorType<unknown>>( options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getCurrentUser>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetCurrentUserQueryKey();

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getCurrentUser>>> = ({ signal }) => getCurrentUser(requestOptions, signal);

      

      

   return  { queryKey, queryFn,   cacheTime: 2000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true,  ...queryOptions} as UseInfiniteQueryOptions<Awaited<ReturnType<typeof getCurrentUser>>, TError, TData> & { queryKey: QueryKey }
}

export type GetCurrentUserInfiniteQueryResult = NonNullable<Awaited<ReturnType<typeof getCurrentUser>>>
export type GetCurrentUserInfiniteQueryError = ErrorType<unknown>

/**
 * @summary returns the logged in user
 */
export const useGetCurrentUserInfinite = <TData = Awaited<ReturnType<typeof getCurrentUser>>, TError = ErrorType<unknown>>(
  options?: { query?:UseInfiniteQueryOptions<Awaited<ReturnType<typeof getCurrentUser>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetCurrentUserInfiniteQueryOptions(options)

  const query = useInfiniteQuery(queryOptions) as  UseInfiniteQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



export const getGetCurrentUserQueryOptions = <TData = Awaited<ReturnType<typeof getCurrentUser>>, TError = ErrorType<unknown>>( options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getCurrentUser>>, TError, TData>, request?: SecondParameter<typeof customInstance>}
) => {

const {query: queryOptions, request: requestOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetCurrentUserQueryKey();

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getCurrentUser>>> = ({ signal }) => getCurrentUser(requestOptions, signal);

      

      

   return  { queryKey, queryFn,   cacheTime: 2000, refetchOnWindowFocus: false, refetchOnMount: false, retryOnMount: false, staleTime: Infinity, keepPreviousData: true,  ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getCurrentUser>>, TError, TData> & { queryKey: QueryKey }
}

export type GetCurrentUserQueryResult = NonNullable<Awaited<ReturnType<typeof getCurrentUser>>>
export type GetCurrentUserQueryError = ErrorType<unknown>

/**
 * @summary returns the logged in user
 */
export const useGetCurrentUser = <TData = Awaited<ReturnType<typeof getCurrentUser>>, TError = ErrorType<unknown>>(
  options?: { query?:UseQueryOptions<Awaited<ReturnType<typeof getCurrentUser>>, TError, TData>, request?: SecondParameter<typeof customInstance>}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetCurrentUserQueryOptions(options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



/**
 * @summary updates user role and scopes by id
 */
export const updateUserAuthorization = (
    id: string,
    updateUserAuthRequest: UpdateUserAuthRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<void>(
      {url: `/user/${id}/authorization`, method: 'PATCH',
      headers: {'Content-Type': 'application/json', },
      data: updateUserAuthRequest
    },
      options);
    }
  


export const getUpdateUserAuthorizationMutationOptions = <TError = ErrorType<unknown>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updateUserAuthorization>>, TError,{id: string;data: UpdateUserAuthRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof updateUserAuthorization>>, TError,{id: string;data: UpdateUserAuthRequest}, TContext> => {
const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof updateUserAuthorization>>, {id: string;data: UpdateUserAuthRequest}> = (props) => {
          const {id,data} = props ?? {};

          return  updateUserAuthorization(id,data,requestOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type UpdateUserAuthorizationMutationResult = NonNullable<Awaited<ReturnType<typeof updateUserAuthorization>>>
    export type UpdateUserAuthorizationMutationBody = UpdateUserAuthRequest
    export type UpdateUserAuthorizationMutationError = ErrorType<unknown>

    /**
 * @summary updates user role and scopes by id
 */
export const useUpdateUserAuthorization = <TError = ErrorType<unknown>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updateUserAuthorization>>, TError,{id: string;data: UpdateUserAuthRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationResult<
        Awaited<ReturnType<typeof updateUserAuthorization>>,
        TError,
        {id: string;data: UpdateUserAuthRequest},
        TContext
      > => {

      const mutationOptions = getUpdateUserAuthorizationMutationOptions(options);

      return useMutation(mutationOptions);
    }
    /**
 * @summary deletes the user by id
 */
export const deleteUser = (
    id: string,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<void>(
      {url: `/user/${id}`, method: 'DELETE'
    },
      options);
    }
  


export const getDeleteUserMutationOptions = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deleteUser>>, TError,{id: string}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof deleteUser>>, TError,{id: string}, TContext> => {
const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof deleteUser>>, {id: string}> = (props) => {
          const {id} = props ?? {};

          return  deleteUser(id,requestOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type DeleteUserMutationResult = NonNullable<Awaited<ReturnType<typeof deleteUser>>>
    
    export type DeleteUserMutationError = ErrorType<void | HTTPError>

    /**
 * @summary deletes the user by id
 */
export const useDeleteUser = <TError = ErrorType<void | HTTPError>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deleteUser>>, TError,{id: string}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationResult<
        Awaited<ReturnType<typeof deleteUser>>,
        TError,
        {id: string},
        TContext
      > => {

      const mutationOptions = getDeleteUserMutationOptions(options);

      return useMutation(mutationOptions);
    }
    /**
 * @summary updates the user by id
 */
export const updateUser = (
    id: string,
    updateUserRequest: UpdateUserRequest,
 options?: SecondParameter<typeof customInstance>,) => {
      
      
      return customInstance<UserResponse>(
      {url: `/user/${id}`, method: 'PATCH',
      headers: {'Content-Type': 'application/json', },
      data: updateUserRequest
    },
      options);
    }
  


export const getUpdateUserMutationOptions = <TError = ErrorType<unknown>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updateUser>>, TError,{id: string;data: UpdateUserRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationOptions<Awaited<ReturnType<typeof updateUser>>, TError,{id: string;data: UpdateUserRequest}, TContext> => {
const {mutation: mutationOptions, request: requestOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof updateUser>>, {id: string;data: UpdateUserRequest}> = (props) => {
          const {id,data} = props ?? {};

          return  updateUser(id,data,requestOptions)
        }

        


  return  { mutationFn, ...mutationOptions }}

    export type UpdateUserMutationResult = NonNullable<Awaited<ReturnType<typeof updateUser>>>
    export type UpdateUserMutationBody = UpdateUserRequest
    export type UpdateUserMutationError = ErrorType<unknown>

    /**
 * @summary updates the user by id
 */
export const useUpdateUser = <TError = ErrorType<unknown>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof updateUser>>, TError,{id: string;data: UpdateUserRequest}, TContext>, request?: SecondParameter<typeof customInstance>}
): UseMutationResult<
        Awaited<ReturnType<typeof updateUser>>,
        TError,
        {id: string;data: UpdateUserRequest},
        TContext
      > => {

      const mutationOptions = getUpdateUserMutationOptions(options);

      return useMutation(mutationOptions);
    }
    