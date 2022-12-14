import { EuiEmptyPrompt } from '@elastic/eui'
import _ from 'lodash'
import type { ReactNode } from 'react'
import type { Role, Scopes, UserResponse } from 'src/gen/model'

type ProtectedComponentProps = {
  children: JSX.Element
  user: UserResponse
  requiredRole?: Role
  requiredScopes?: Scopes
}

// usage:
// <...
//   <ProtectedComponent requiredRole="manager" ...>
//   <ProtectedComponent requiredRole="admin" ...>
// <...
export default function ProtectedComponent({ children, user, requiredRole, requiredScopes }: ProtectedComponentProps) {
  // TODO isAuthorized

  return children
}
