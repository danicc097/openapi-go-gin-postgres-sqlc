import { EuiLoadingSpinner } from '@elastic/eui'
import ProtectedPage from './ProtectedPage'
import { Navigate } from 'react-router-dom'
import type { Role, Scopes } from 'src/gen/model'
import roles from '@roles'
import { ToastId } from 'src/utils/toasts'
import { useUISlice } from 'src/slices/ui'
import { useAuthenticatedUser } from 'src/hooks/auth/useAuthenticatedUser'
import { useEffect } from 'react'

type ProtectedRouteProps = {
  component: React.ComponentType
  requiredRole?: Role
  requiredScopes?: Scopes
}

/**
 *  Requires an authenticated user and optionally specific role or scopes.
 */
export default function ProtectedRoute({
  component: Component,
  requiredRole = null,
  requiredScopes = null,
  ...props
}: ProtectedRouteProps) {
  const { user } = useAuthenticatedUser()
  const { addToast } = useUISlice()

  useEffect(() => {
    if (!user) {
      addToast({
        id: ToastId.AuthRedirect,
        title: 'Access Denied',
        color: 'warning',
        iconType: 'alert',
        toastLifeTimeMs: 15000,
        text: 'Authenticated users only. Log in here or create a new account to view that page',
      })
    }
  }, [addToast, user])

  const isAuthenticated = true

  if (!user) {
    return <Navigate to="/login" />
  }

  if (!isAuthenticated && user) {
    window.location.replace(`${import.meta.env.VITE_AUTH_SERVER}/login`)
  }

  const element = <Component {...props} />

  const isAuthorized = true // TODO

  return <ProtectedPage element={element} isAuthorized={isAuthorized}></ProtectedPage>
}
