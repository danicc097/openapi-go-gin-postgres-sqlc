import { DefaultOptions } from '@tanstack/react-query'

// used by orval as defaults
export const reactQueryDefaultAppOptions: DefaultOptions = {
  queries: {
    cacheTime: 1000 * 60 * 5,
    // cacheTime: 0,
    refetchOnWindowFocus: false,
    refetchOnMount: false,
    retryOnMount: false,
    staleTime: Infinity,
    keepPreviousData: true,
  },
  mutations: {
    cacheTime: 1000 * 60 * 5,
  },
}
