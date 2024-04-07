import React from 'react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { Button, useMantineTheme } from '@mantine/core'
import { apiPath } from 'src/services/apiPaths'
import { faSignIn } from '@fortawesome/free-solid-svg-icons'
import { redirectToAuthLogin } from 'src/services/authorization'

export default function LoginButton() {
  const { colors } = useMantineTheme()

  return (
    <Button
      style={{
        backgroundColor: colors.blue[9],
      }}
      onClick={(e) => {
        e.preventDefault()
        redirectToAuthLogin()
      }}
      leftSection={<FontAwesomeIcon icon={faSignIn} size="sm" />}
    >
      Login
    </Button>
  )
}
