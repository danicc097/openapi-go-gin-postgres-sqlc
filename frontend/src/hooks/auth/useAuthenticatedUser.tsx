import { shallowEqual } from 'react-redux'
import { useAppDispatch, useAppSelector } from 'src/redux/hooks'
import { capitalize } from 'lodash'
import type { User } from 'src/redux/slices/gen/internalApi'
import { useMemo } from 'react'
import { COLOR_BLIND_PALETTE } from 'src/utils/colors'

export const useAuthenticatedUser = () => {
  const dispatch = useAppDispatch()

  // const user = useAppSelector((state) => state.auth.user, shallowEqual)
  const user: User = useMemo(
    () => ({
      email: 'test@mail.com',
      first_name: 'John',
      last_name: 'Doe',
      role_rank: 5,
      full_name: 'John Doe',
      username: 'john.doe',
      scopes: ['test-scope'],
    }),
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
