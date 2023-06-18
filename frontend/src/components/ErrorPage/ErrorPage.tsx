import { createStyles, Title, Text, Button, Container, Group } from '@mantine/core'
import { useNavigate } from 'react-router-dom'
import { FOOTER_HEIGHT } from 'src/components/Footer'
import { HEADER_HEIGHT } from 'src/components/Header'
import HttpStatus from 'src/utils/httpStatus'

const useStyles = createStyles((theme) => ({
  root: {
    paddingBottom: 80,
  },

  label: {
    textAlign: 'center',
    fontWeight: 900,
    fontSize: 220,
    lineHeight: 1,
    marginBottom: theme.spacing.xl,
    color: theme.colorScheme === 'dark' ? theme.colors.dark[4] : theme.colors.gray[2],

    [theme.fn.smallerThan('sm')]: {
      fontSize: 120,
    },
  },

  title: {
    textAlign: 'center',
    fontWeight: 900,
    fontSize: 38,

    [theme.fn.smallerThan('sm')]: {
      fontSize: 32,
    },
  },

  description: {
    maxWidth: 500,
    margin: 'auto',
    marginTop: theme.spacing.xl,
    marginBottom: theme.spacing.xl,
  },
}))

interface ErrorPageProps {
  status: number
}

export function ErrorPage({ status }: ErrorPageProps) {
  const { classes, theme } = useStyles()

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
      mih={`calc(100vh - ${HEADER_HEIGHT}px - ${FOOTER_HEIGHT}px)`}
      bg={theme.colorScheme === 'dark' ? theme.colors.dark[7] : theme.white}
    >
      <div className={classes.label}>{status}</div>
      <Title className={classes.title}>You have found a secret place.</Title>
      <Text color="dimmed" size="lg" align="center" className={classes.description}>
        {text}
      </Text>
      <Group position="center">
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
