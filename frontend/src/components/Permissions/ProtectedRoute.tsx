import { EuiLoadingSpinner } from '@elastic/eui'
import { useProtectedRoute } from 'src/hooks/auth/useProtectedRoute'
import ProtectedPage from './ProtectedPage'
import { Navigate } from 'react-router-dom'
import type { Role, Scopes } from 'src/gen/model'
import roles from '@roles'

type ProtectedRouteProps = {
  component: React.ComponentType
  requiredRole?: Role
  requiredScopes?: Scopes
}

export default function ProtectedRoute({
  component: Component,
  requiredRole = null,
  requiredScopes = null,
  ...props
}: ProtectedRouteProps) {
  const { user } = useProtectedRoute()

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
