import { Flex, Text } from '@mantine/core'
import React, { useEffect, useState } from 'react'
import InfiniteLoader from 'src/components/Loading/InfiniteLoader'
import useAuthenticatedUser from 'src/hooks/auth/useAuthenticatedUser'
import { ErrorPage } from 'src/components/ErrorPage/ErrorPage'
import { Authorization } from 'src/services/authorization'
import HttpStatus from 'src/utils/httpStatus'

type ProtectedPageProps = {
  children: JSX.Element
  authResult: Authorization
}

export default function ProtectedPage({ children, authResult }: ProtectedPageProps) {
  const { isAuthenticating } = useAuthenticatedUser()
  const [dotCount, setDotCount] = useState(1)

  useEffect(() => {
    const intervalId = setInterval(() => {
      setDotCount((prevCount) => (prevCount % 3) + 1)
    }, 500)

    return () => clearInterval(intervalId)
  }, [])

  if (isAuthenticating) {
    return (
      <Flex p={50} direction={'column'} justify="center" align="center">
        <InfiniteLoader />
        <Text pt={20} size="lg">
          Authenticating user{'.'.repeat(dotCount)}
        </Text>
      </Flex>
    )
  }

  if (!authResult.authorized) {
    return <ErrorPage status={HttpStatus.FORBIDDEN_403} authResult={authResult} />
  }

  return children
}
