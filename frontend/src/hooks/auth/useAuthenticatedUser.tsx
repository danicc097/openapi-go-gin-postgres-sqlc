import { QueryClient, useQueryClient } from '@tanstack/react-query'
import axios from 'axios'
import Cookies from 'js-cookie'
import { useEffect, useRef } from 'react'
import { AXIOS_INSTANCE } from 'src/api/mutator'
import type { User } from 'src/gen/model'
import { useGetCurrentUser } from 'src/gen/user/user'
import { persister } from 'src/idb'
import { LOGIN_COOKIE_KEY, UI_SLICE_PERSIST_KEY, useUISlice } from 'src/slices/ui'
import AxiosInterceptors from 'src/utils/axios'
import { useIsFirstRender } from 'usehooks-ts'

export default function useAuthenticatedUser() {
  const mountedRef = useMountedRef()
  const queryClient = useQueryClient()
  const currentUser = useGetCurrentUser({
    query: {
      retry(failureCount, error) {
        return ui.accessToken !== '' && failureCount < 3
      },
    },
  })
  const isFirstRender = useIsFirstRender()
  const ui = useUISlice()
  const isAuthenticated = !!currentUser.data?.userID

  useEffect(() => {
    if (mountedRef.current && isFirstRender) {
      // FIXME: ... one-off logic (in theory, not working)
    }

    if (!isAuthenticated && !currentUser.isFetching && ui.accessToken !== '') {
      currentUser.refetch()
    }

    AxiosInterceptors.setupAxiosInstance(AXIOS_INSTANCE, ui.accessToken)

    return () => {
      AxiosInterceptors.teardownAxiosInstance(AXIOS_INSTANCE)
    }
  }, [currentUser.data, isFirstRender, isAuthenticated, ui.accessToken])

  const user = currentUser.data!

  return {
    isAuthenticated,
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
