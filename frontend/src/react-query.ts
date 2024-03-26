import { DefaultOptions, QueryClient } from '@tanstack/react-query'
import { AxiosError } from 'axios'
import { ApiError } from 'src/api/mutator'
import { reactQueryDefaultAppOptions } from 'src/react-query.default'
import HttpStatus from 'src/utils/httpStatus'

export const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      ...reactQueryDefaultAppOptions.queries,
      retry: function (failureCount, error: AxiosError | ApiError) {
        const status = error.response?.status
        if (status && status >= 400 && status < 500 && status !== HttpStatus.UNAUTHORIZED_401) {
          return false
        }
        return failureCount < 3
      },
    },
    mutations: {
      ...reactQueryDefaultAppOptions.mutations,
      retry: function (failureCount, error: AxiosError | ApiError) {
        const status = error.response?.status
        if (status && status >= 400 && status < 500) {
          return false
        }
        return failureCount < 2
      },
    },
  },
  // queryCache,
})
