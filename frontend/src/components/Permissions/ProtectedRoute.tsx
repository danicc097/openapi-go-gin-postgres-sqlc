import ProtectedPage from './ProtectedPage'
import { Navigate, useLocation, useNavigate } from 'react-router-dom'
import type { Role, Scopes } from 'src/gen/model'
import { ToastId } from 'src/utils/toasts'
import { useUISlice } from 'src/slices/ui'
import useAuthenticatedUser from 'src/hooks/auth/useAuthenticatedUser'
import { useEffect, useState } from 'react'
import { isAuthorized, redirectToAuthLogin } from 'src/services/authorization'
import { apiPath } from 'src/services/apiPaths'
import { notifications } from '@mantine/notifications'
import { useMyProviderLogin } from 'src/gen/oidc/oidc'
import { useGetCurrentUser } from 'src/gen/user/user'

type ProtectedRouteProps = {
  children: JSX.Element
  requiredRole?: Role
  requiredScopes?: Scopes
}

/**
 *  Requires an authenticated user and optionally specific role or scopes.
 */
export default function ProtectedRoute({ children, requiredRole, requiredScopes }: ProtectedRouteProps) {
  const { user, isAuthenticated } = useAuthenticatedUser()
  const currentUser = useGetCurrentUser()
  const ui = useUISlice()

  if (!isAuthenticated && !currentUser.isFetching) {
    redirectToAuthLogin()
  }

  if (!user) {
    return null
  }

  return <ProtectedPage isAuthorized={isAuthorized({ user, requiredRole, requiredScopes })}>{children}</ProtectedPage>
}
