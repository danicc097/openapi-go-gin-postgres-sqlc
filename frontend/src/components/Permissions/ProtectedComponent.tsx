import { EuiEmptyPrompt } from '@elastic/eui'
import _ from 'lodash'
import React from 'react'
import type { Role, Scopes, UserResponse } from 'src/gen/model'

type ProtectedComponentProps = {
  element: JSX.Element
  user: UserResponse
  requiredRole?: Role
  requiredScopes?: Scopes
}

// usage:
// <...
//   <ProtectedComponent requiredRole="manager" ...>
//   <ProtectedComponent requiredRole="admin" ...>
// <...
export default function ProtectedComponent({ element, user, requiredRole, requiredScopes }: ProtectedComponentProps) {
  // TODO isAuthorized

  return element
}
