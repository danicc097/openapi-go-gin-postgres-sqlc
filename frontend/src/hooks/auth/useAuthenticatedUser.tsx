import { shallowEqual } from 'react-redux'
import { useAppDispatch, useAppSelector } from 'src/redux/hooks'
import { capitalize } from 'lodash'
import type { UserResponse } from 'src/redux/slices/gen/internalApi'
import { useMemo } from 'react'
import { COLOR_BLIND_PALETTE } from 'src/utils/colors'
import roles from '@roles'
import scopes from '@scopes'
export const useAuthenticatedUser = () => {
  const dispatch = useAppDispatch()

  // const user = useAppSelector((state) => state.auth.user, shallowEqual)
  const user: UserResponse = useMemo(
    () =>
      ({
        email: 'test@mail.com',
        first_name: 'John',
        last_name: 'Doe',
        role_rank: roles.user.rank,
        full_name: 'John Doe',
        username: 'john.doe',
        scopes: ['users:read', 'test-scope', 'scopes:write'],
        createdAt: Date.now(),
        deletedAt: null,
      } as UserResponse),
    [],
  )
  const avatarColor =
    COLOR_BLIND_PALETTE[capitalize(user?.email).charCodeAt(0) % COLOR_BLIND_PALETTE.length] || '#1060e0'

  const logUserOut = () => {
    null
  }

  const registerNewUser = ({ username, email, password }) => {
    null
  }
  const requestUserLogin = ({ email, password }) => {
    null
  }

  return {
    registerNewUser,
    requestUserLogin,
    user,
    avatarColor,
    logUserOut,
  }
}
