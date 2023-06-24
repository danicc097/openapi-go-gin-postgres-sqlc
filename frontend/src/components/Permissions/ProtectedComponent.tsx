import _ from 'lodash'
import type { ReactNode } from 'react'
import type { Role, Scopes, User } from 'src/gen/model'
import useAuthenticatedUser from 'src/hooks/auth/useAuthenticatedUser'
import { isAuthorized } from 'src/services/authorization'

type ProtectedComponentProps = {
  children: JSX.Element
  requiredRole?: Role
  requiredScopes?: Scopes
}

export default function ProtectedComponent({ children, requiredRole, requiredScopes }: ProtectedComponentProps) {
  const { user } = useAuthenticatedUser()

  if (!isAuthorized({ user, requiredRole, requiredScopes })) return null

  return children
}
