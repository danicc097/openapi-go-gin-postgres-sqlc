import _ from 'lodash'
import React from 'react'
import { ErrorPage } from 'src/components/ErrorPage/ErrorPage'
import type { Role, Scopes } from 'src/gen/model'
import HttpStatus from 'src/utils/httpStatus'

type ProtectedPageProps = {
  children: JSX.Element
  isAuthorized: boolean
}

export default function ProtectedPage({ children, isAuthorized }: ProtectedPageProps) {
  if (!isAuthorized) {
    return <ErrorPage status={HttpStatus.FORBIDDEN_403} />
  }

  return children
}
