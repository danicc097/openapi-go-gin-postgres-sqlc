import React from 'react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { Button } from '@mantine/core'
import { apiPath } from 'src/services/apiPaths'
import { faSignIn } from '@fortawesome/free-solid-svg-icons'

export default function LoginTwitchButton() {
  const redirectURI = encodeURIComponent(`${import.meta.env.VITE_URL}/api/auth/twitch/callback/`)

  return (
    <>
      <form>
        <Button
          type="submit"
          style={{
            backgroundColor: '#a970ff',
          }}
          onClick={(e) => {
            e.preventDefault()
            location.href = apiPath('/auth/myprovider/login')
          }}
          leftIcon={<FontAwesomeIcon icon={faSignIn} size="xl" />}
        >
          Login
        </Button>
      </form>
    </>
  )
}
