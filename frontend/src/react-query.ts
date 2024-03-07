import { DefaultOptions, QueryClient } from '@tanstack/react-query'
import { AxiosError } from 'axios'
import { ApiError } from 'src/api/mutator'

export const reactQueryDefaultAppOptions: DefaultOptions = {
  queries: {
    cacheTime: 1000 * 60 * 5,
    // cacheTime: 0,
    refetchOnWindowFocus: false,
    refetchOnMount: false,
    retryOnMount: false,
    staleTime: Infinity,
    keepPreviousData: true,
    retry: function (failureCount, error: AxiosError | ApiError) {
      const status = error.response?.status
      if (status && status >= 400 && status < 500) {
        return false
      }
      return failureCount < 3
    },
  },
  mutations: {
    cacheTime: 1000 * 60 * 5,
    retry: function (failureCount, error: AxiosError | ApiError) {
      const status = error.response?.status
      if (status && status >= 400 && status < 500) {
        return false
      }
      return failureCount < 2
    },
  },
}

export const queryClient = new QueryClient({
  defaultOptions: reactQueryDefaultAppOptions,
  // queryCache,
})
