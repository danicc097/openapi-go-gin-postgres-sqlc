import ProtectedPage from './ProtectedPage'
import { Navigate } from 'react-router-dom'
import type { Role, Scopes } from 'src/gen/model'
import roles from '@roles'
import { ToastId } from 'src/utils/toasts'
import { useUISlice } from 'src/slices/ui'
import useAuthenticatedUser from 'src/hooks/auth/useAuthenticatedUser'
import { useEffect } from 'react'
import { isAuthorized } from 'src/services/authorization'
import { apiPath } from 'src/services/apiPaths'
import { notifications } from '@mantine/notifications'

type ProtectedRouteProps = {
  children: JSX.Element
  requiredRole?: Role
  requiredScopes?: Scopes
}

/**
 *  Requires an authenticated user and optionally specific role or scopes.
 */
export default function ProtectedRoute({ children, requiredRole = null, requiredScopes = null }: ProtectedRouteProps) {
  const { user } = useAuthenticatedUser()

  useEffect(() => {
    if (!user) {
      notifications.show({
        id: ToastId.AuthRedirect,
        title: 'Access Denied',
        color: 'warning',
        variant: 'alert',
        withCloseButton: true,
        autoClose: 15000,
        message: 'Authenticated users only. Log in here or create a new account to view that page',
      })
    }
  }, [user])

  const isAuthenticated = true

  if (!user) {
    return <Navigate to="/login" />
  }

  if (!isAuthenticated && user) {
    window.location.replace(apiPath('/auth/myprovider/login'))
  }

  return <ProtectedPage isAuthorized={isAuthorized({ user, requiredRole, requiredScopes })}>{children}</ProtectedPage>
}
