import { css } from '@emotion/react'
import {
  Title,
  Text,
  Button,
  Container,
  Group,
  useMantineTheme,
  useMantineColorScheme,
  Flex,
  Space,
} from '@mantine/core'
import { useNavigate } from 'react-router-dom'
import HttpStatus from 'src/utils/httpStatus'
import classes from './ErrorPage.module.css'

interface ErrorPageProps {
  status: number
  unauthorizedMessage?: string
}

export function ErrorPage({ status, unauthorizedMessage }: ErrorPageProps) {
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
    <Flex direction={'column'} align={'center'} className={classes.root}>
      <div className={classes.label}>{status}</div>
      <Title className={classes.title}>You have found a secret place.</Title>
      <Text color="dimmed" size="m" ta="center" className={classes.description}>
        {text}
      </Text>
      {unauthorizedMessage && (
        <>
          <Text p={30} color="dimmed" size="m" ta="center" className={classes.description}>
            {unauthorizedMessage}
          </Text>
        </>
      )}
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
    </Flex>
  )
}
