import { QueryClient, useQueryClient } from '@tanstack/react-query'
import Cookies from 'js-cookie'
import { useEffect, useRef } from 'react'
import { persister } from 'src/App'
import type { UserResponse } from 'src/gen/model'
import { useGetCurrentUser } from 'src/gen/user/user'
import { ACCESS_TOKEN_COOKIE, UI_SLICE_PERSIST_KEY } from 'src/slices/ui'
import { useIsFirstRender } from 'usehooks-ts'

export default function useAuthenticatedUser() {
  const mountedRef = useMountedRef()
  const queryClient = useQueryClient()
  const currentUser = useGetCurrentUser()
  const isFirstRender = useIsFirstRender()

  const isAuthenticated = !!currentUser.data?.data?.userID

  useEffect(() => {
    if (mountedRef.current && isFirstRender) {
      console.log('would have triggered useAuthenticatedUser useEffect')
      // if (!twitchValidateToken.isLoading) twitchValidateToken.refetch()
    }
  }, [currentUser.data, isFirstRender])

  const user: UserResponse = {
    userID: 'c7fd2433-dbb7-4612-ab13-ddb0d3404728',
    username: 'user_2',
    email: 'user_2@email.com',
    firstName: 'Name 2',
    lastName: 'Surname 2',
    fullName: 'Name 2 Surname 2',
    hasPersonalNotifications: false,
    hasGlobalNotifications: true,
    createdAt: new Date('2023-04-01T06:24:22.390699Z'),
    deletedAt: null,

    role: 'user',
    scopes: ['users:read'],

    apiKey: null,
    teams: null,
    projects: null,
  }

  return {
    isAuthenticated,
    user,
  }
}

// TODO doesnt seem to clear react query
export async function logUserOut(queryClient: QueryClient) {
  await persister.removeClient() // delete indexed db
  await queryClient.cancelQueries()
  await queryClient.invalidateQueries()
  queryClient.clear()
  Cookies.remove(ACCESS_TOKEN_COOKIE, {
    expires: 365,
    sameSite: 'none',
    secure: true,
  })
  localStorage.removeItem(UI_SLICE_PERSIST_KEY)
  window.location.reload()
}

/**
 * To ensure a useEffect is only called once for shared hooks.
 */
const useMountedRef = () => {
  const mountedRef = useRef(false)

  useEffect(() => {
    setTimeout(() => {
      mountedRef.current = true
    })

    return () => (mountedRef.current = null)
  }, [])

  return mountedRef
}
