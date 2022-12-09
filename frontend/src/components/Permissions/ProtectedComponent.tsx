import { EuiEmptyPrompt } from '@elastic/eui'
import _ from 'lodash'
import React from 'react'
import type { Role, Scopes, UserResponse } from 'src/gen/model'

type ProtectedComponentProps = {
  element: JSX.Element
  requiredRole?: Role
  requiredScopes?: Scopes
}

export default function ProtectedComponent({ element, requiredRole, requiredScopes }: ProtectedComponentProps) {
  // TODO isAuthorized

  return element
}
