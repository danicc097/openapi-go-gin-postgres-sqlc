// Or from '@reduxjs/toolkit/query' if not using the auto-generated hooks
import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'
import Config from 'config.json'

// initialize an empty api service that we'll inject endpoints into later as needed
const PORT = import.meta.env.PROD ? Config.API_PORT : ''

export const emptyInternalApi = createApi({
  baseQuery: fetchBaseQuery({
    baseUrl: `https://${Config.DOMAIN}${PORT}${Config.API_PREFIX}${Config.API_VERSION}`,
  }),
  endpoints: () => ({}),
  reducerPath: 'internalApi',
  // TODO prepareHeaders for auth
})
