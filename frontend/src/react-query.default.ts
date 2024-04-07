import { DefaultOptions } from '@tanstack/react-query'

// used by orval as defaults
export const reactQueryDefaultAppOptions: DefaultOptions = {
  queries: {
    cacheTime: 0, //ms
    refetchOnWindowFocus: false,
    refetchOnMount: false,
    retryOnMount: false,
    staleTime: Infinity,
    keepPreviousData: true,
  },
  mutations: {
    cacheTime: 0,
  },
}
