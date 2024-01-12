import { Flex, Group, Text } from '@mantine/core'
import _ from 'lodash'
import React from 'react'
import { ErrorPage } from 'src/components/ErrorPage/ErrorPage'
import InfiniteLoader from 'src/components/Loading/InfiniteLoader'
import type { Role, Scopes } from 'src/gen/model'
import useAuthenticatedUser from 'src/hooks/auth/useAuthenticatedUser'
import HttpStatus from 'src/utils/httpStatus'

type ProtectedPageProps = {
  children: JSX.Element
  isAuthorized: boolean
  unauthorizedMessage?: string
}

export default function ProtectedPage({ children, isAuthorized, unauthorizedMessage }: ProtectedPageProps) {
  const { isAuthenticating } = useAuthenticatedUser()

  if (isAuthenticating) {
    return (
      <Flex direction={'column'} justify="center" align="center">
        <InfiniteLoader />
        <Text size="lg">Loading user...</Text>
      </Flex>
    )
  }

  if (!isAuthorized) {
    return <ErrorPage status={HttpStatus.FORBIDDEN_403} unauthorizedMessage={unauthorizedMessage} />
  }

  return children
}
