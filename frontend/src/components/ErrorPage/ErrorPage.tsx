import { css } from '@emotion/react'
import { Title, Text, Button, Container, Group, useMantineTheme, useMantineColorScheme } from '@mantine/core'
import { useNavigate } from 'react-router-dom'
import { HEADER_HEIGHT } from 'src/components/Header'
import HttpStatus from 'src/utils/httpStatus'
import classes from './ErrorPage.module.css'

interface ErrorPageProps {
  status: number
}

export function ErrorPage({ status }: ErrorPageProps) {
  const { colorScheme } = useMantineColorScheme()
  const navigate = useNavigate()

  let text = 'An unknown error ocurred.'
  switch (status) {
    case HttpStatus.NOT_FOUND_404:
      text = 'You may have mistyped the address, or the page has been moved to another URL.'
      break
    case HttpStatus.FORBIDDEN_403:
      text = "You don't have the required permissions to access this content."
      break
    case HttpStatus.UNAUTHORIZED_401:
      text = 'You need to log in before accessing this content.'
    default:
      break
  }

  return (
    <Container
      className={classes.root}
      miw={'100vw'}
      css={css`
        min-height: calc(100vh - var(--footer-height) - var(--header-height));
        background: light-dark(--mantine-color-white, --mantine-color-dark-7);
      `}
    >
      <div className={classes.label}>{status}</div>
      <Title className={classes.title}>You have found a secret place.</Title>
      <Text color="dimmed" size="lg" ta="center" className={classes.description}>
        {text}
      </Text>
      <Group align="center">
        <Button
          size="md"
          color="teal"
          onClick={() => {
            navigate('/')
          }}
        >
          Take me back to the home page
        </Button>
        <Button
          size="md"
          onClick={() => {
            navigate(-1)
          }}
        >
          Take me back to the previous page
        </Button>
      </Group>
    </Container>
  )
}
