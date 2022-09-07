// Or from '@reduxjs/toolkit/query' if not using the auto-generated hooks
import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'

// initialize an empty api service that we'll inject endpoints into later as needed
export const emptyInternalApi = createApi({
  baseQuery: fetchBaseQuery({ baseUrl: 'https://localhost:8090/v2/' }),
  endpoints: () => ({}),
  reducerPath: 'internalApi',
  // TODO prepareHeaders for auth
})
