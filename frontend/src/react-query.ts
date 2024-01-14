import { DefaultOptions, QueryClient } from '@tanstack/react-query'

export const reactQueryDefaultAppOptions: DefaultOptions = {
  queries: {
    cacheTime: 1000 * 60 * 5,
    // cacheTime: 0,
    refetchOnWindowFocus: false,
    refetchOnMount: false,
    retryOnMount: false,
    staleTime: Infinity,
    keepPreviousData: true,
    retry: function (failureCount, error) {
      return failureCount < 3
    },
  },
  mutations: {
    cacheTime: 1000 * 60 * 5,
  },
}
export const queryClient = new QueryClient({
  defaultOptions: reactQueryDefaultAppOptions,
  // queryCache,
})
