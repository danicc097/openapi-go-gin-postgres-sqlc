import _ from 'lodash'
import type { ReactNode } from 'react'
import { OperationAuth } from 'src/config'
import type { Role, Scopes, User } from 'src/gen/model'
import useAuthenticatedUser from 'src/hooks/auth/useAuthenticatedUser'
import { isAuthorized } from 'src/services/authorization'

type ProtectedComponentProps = {
  children: JSX.Element
  operationAuth: OperationAuth
}

export default function ProtectedComponent({ children, operationAuth }: ProtectedComponentProps) {
  const { user } = useAuthenticatedUser()

  if (!isAuthorized({ user, requiredRole: operationAuth.role, requiredScopes: operationAuth.scopes })) return null

  return children
}
