import ProtectedPage from './ProtectedPage'
import { Navigate, useLocation, useNavigate } from 'react-router-dom'
import type { Role, Scopes } from 'src/gen/model'
import { ToastId } from 'src/utils/toasts'
import { useUISlice } from 'src/slices/ui'
import useAuthenticatedUser from 'src/hooks/auth/useAuthenticatedUser'
import { useEffect, useState } from 'react'
import { Authorization, checkAuthorization, redirectToAuthLogin } from 'src/services/authorization'
import { apiPath } from 'src/services/apiPaths'
import { notifications } from '@mantine/notifications'
import { useMyProviderLogin } from 'src/gen/oidc/oidc'
import { useGetCurrentUser } from 'src/gen/user/user'
import { joinWithAnd } from 'src/utils/format'

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
  const ui = useUISlice()

  const authResult = checkAuthorization({ user, requiredRole, requiredScopes })

  // if (!isAuthenticated && !currentUser.isFetching) {
  //   redirectToAuthLogin();
  // }

  return <ProtectedPage authResult={authResult}>{children}</ProtectedPage>
}
