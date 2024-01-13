import { QueryClient } from '@tanstack/react-query'

export const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      cacheTime: 1000 * 60 * 5, // 5 min
      // cacheTime: 0,
      refetchOnWindowFocus: false,
      refetchOnMount: false,
      retryOnMount: false,
      staleTime: Infinity,
      keepPreviousData: true,
      retry(failureCount, error) {
        return failureCount < 3
      },
    },
    mutations: {
      cacheTime: 1000 * 60 * 5, // 5 minutes
    },
  },
  // queryCache,
})
