import { EuiEmptyPrompt } from '@elastic/eui'
import _ from 'lodash'
import React from 'react'
import type { Role, Scopes } from 'src/gen/model'

type ProtectedPageProps = {
  children: JSX.Element
  isAuthorized: boolean
}

export default function ProtectedPage({ children, isAuthorized }: ProtectedPageProps) {
  if (!isAuthorized) {
    return (
      <EuiEmptyPrompt
        iconType="securityApp"
        iconColor={null}
        title={<h2 className="eui-textInheritColor">Access Denied</h2>}
        body={<p>{`You don't have the required permissions to access this content.`}</p>}
      />
    )
  }

  return children
}
