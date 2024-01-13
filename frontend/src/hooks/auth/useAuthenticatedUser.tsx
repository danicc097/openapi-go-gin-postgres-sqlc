import { QueryClient, useQueryClient } from '@tanstack/react-query'
import axios from 'axios'
import Cookies from 'js-cookie'
import { useEffect, useRef, useState } from 'react'
import { AXIOS_INSTANCE } from 'src/api/mutator'
import type { User } from 'src/gen/model'
import { useGetCurrentUser } from 'src/gen/user/user'
import useRenders from 'src/hooks/utils/useRenders'
import { persister } from 'src/idb'
import { LOGIN_COOKIE_KEY, UI_SLICE_PERSIST_KEY, useUISlice } from 'src/slices/ui'
import AxiosInterceptors from 'src/utils/axios'
import { useIsFirstRender } from 'usehooks-ts'

export default function useAuthenticatedUser() {
  const mountedRef = useMountedRef()
  const queryClient = useQueryClient()
  const [failedAuthentication, setFailedAuthentication] = useState(false)
  const currentUser = useGetCurrentUser({
    query: {
      retryDelay: 500,
      retry(failureCount, error) {
        console.log(`retry on useAuthenticatedUser: ${failureCount}`)
        const shouldRetry = ui.accessToken !== '' && failureCount < 2 && !failedAuthentication
        if (!shouldRetry) setFailedAuthentication(true)

        return shouldRetry
      },
    },
  })
  const renders = useRenders()
  const isFirstRender = useIsFirstRender()
  const ui = useUISlice()
  const isAuthenticated = !!currentUser.data?.userID
  const isAuthenticating = currentUser.isFetching && ui.accessToken !== ''
  // console.log({ isFirstRender })

  useEffect(() => {
    if (mountedRef.current && isFirstRender) {
      // FIXME: ... one-off logic (in theory, not working)
      // console.log({ renders: renders })
    }
    // console.log({ rendersOutside: renders })

    if (!isAuthenticated && !isAuthenticating) {
      currentUser.refetch()
    }

    AxiosInterceptors.setupAxiosInstance(AXIOS_INSTANCE, ui.accessToken)

    return () => {
      AxiosInterceptors.teardownAxiosInstance(AXIOS_INSTANCE)
    }
  }, [currentUser.data, isFirstRender, isAuthenticated, ui.accessToken])

  const user = currentUser.data

  useEffect(() => {
    if (user) {
      setFailedAuthentication(false)
    }
  }, [user])

  return {
    isAuthenticated,
    isAuthenticating,
    isLoggingOut: ui.isLoggingOut,
    user,
  }
}

// TODO doesnt seem to clear react query
export async function logUserOut(queryClient: QueryClient) {
  await persister.removeClient() // delete indexed db
  Cookies.remove(LOGIN_COOKIE_KEY, {
    expires: 365,
    sameSite: 'none',
    secure: true,
  })
  await queryClient.cancelQueries()
  await queryClient.invalidateQueries()
  queryClient.clear()
  localStorage.removeItem(UI_SLICE_PERSIST_KEY)
  window.location.reload()
}

/**
 * To ensure a useEffect is only called once for shared hooks.
 */
export function useMountedRef() {
  const mounted = useRef(false)

  useEffect(() => {
    mounted.current = true

    return () => {
      mounted.current = false
    }
  }, [])

  return mounted
}
