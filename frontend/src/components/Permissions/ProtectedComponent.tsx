import _ from 'lodash'
import type { ReactNode } from 'react'
import { OperationAuth } from 'src/config'
import type { Role, Scopes } from 'src/gen/model'
import useAuthenticatedUser from 'src/hooks/auth/useAuthenticatedUser'
import { useOperationAuth as useIsAuthorizedForOp } from 'src/hooks/auth/useOperationAuth'
import { checkAuthorization } from 'src/services/authorization'

type ProtectedComponentProps = {
  children: JSX.Element
  operationAuth: OperationAuth
}

export default function ProtectedComponent({ children, operationAuth }: ProtectedComponentProps) {
  if (!useIsAuthorizedForOp(operationAuth)) return null

  return children
}
