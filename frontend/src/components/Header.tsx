import { useEffect, useState } from 'react'
import {
  createStyles,
  Container,
  Avatar,
  UnstyledButton,
  Group,
  Text,
  Menu,
  Tabs,
  Burger,
  Loader,
  Header as MantineHeader,
  Box,
  Tooltip,
  Badge,
  useMantineTheme,
} from '@mantine/core'
import { IconLogout, IconHeart, IconSettings, IconChevronDown } from '@tabler/icons'
import LoginButton from './LoginButton'
import { ThemeSwitcher } from './ThemeSwitcher'
import useAuthenticatedUser, { logUserOut } from 'src/hooks/auth/useAuthenticatedUser'
import { css } from '@emotion/react'
import { useQueryClient } from '@tanstack/react-query'
import { useNavigate } from 'react-router-dom'
import { useUISlice } from 'src/slices/ui'
import { useGetCurrentUser } from 'src/gen/user/user'
import { useNotificationAPI } from 'src/hooks/ui/useNotificationAPI'
import CONFIG from 'src/config'
import { faUser } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'

const useStyles = createStyles((theme) => ({
  header: {
    backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[6] : theme.colors.gray[0],
    borderBottom: `1px solid ${theme.colorScheme === 'dark' ? 'transparent' : theme.colors.gray[2]}`,
    padding: '30px',
    display: 'grid',
    alignContent: 'center',
  },

  user: {
    color: theme.colorScheme === 'dark' ? theme.colors.dark[0] : theme.black,
    padding: `${theme.spacing.xs}px ${theme.spacing.sm}px`,
    borderRadius: theme.radius.sm,
    transition: 'background-color 100ms ease',

    '&:hover': {
      backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[8] : theme.white,
    },

    [theme.fn.smallerThan('xs')]: {
      '.display-name': {
        display: 'none',
      },
    },
  },

  userActive: {
    backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[8] : theme.white,
  },

  tabs: {
    [theme.fn.smallerThan('sm')]: {
      display: 'none',
    },
  },

  burger: {
    [theme.fn.largerThan(1200)]: {
      display: 'none',
    },
  },

  tabsList: {
    borderBottom: '0 !important',
  },

  tab: {
    fontWeight: 500,
    height: 38,
    backgroundColor: 'transparent',

    '&:hover': {
      backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[5] : theme.colors.gray[1],
    },

    '&[data-active]': {
      backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[7] : theme.white,
      borderColor: theme.colorScheme === 'dark' ? theme.colors.dark[7] : theme.colors.gray[2],
    },
  },
}))

interface HeaderProps {
  tabs: string[]
}

export const HEADER_HEIGHT = 60

export default function Header({ tabs }: HeaderProps) {
  const queryClient = useQueryClient()
  const navigate = useNavigate()
  const { classes, theme, cx } = useStyles()
  // const [opened, { toggle }] = useDisclosure(false)
  const [userMenuOpened, setUserMenuOpened] = useState(false)
  const { user } = useAuthenticatedUser()
  const [loginOut, setLoginOut] = useState(false)

  const items = tabs.map((tab) => (
    <Tabs.Tab value={tab} key={tab}>
      {tab}
    </Tabs.Tab>
  ))
  const { burgerOpened, setBurgerOpened } = useUISlice()

  const [notify, setNotify] = useState<boolean>(false)
  const { showTestNotification } = useNotificationAPI()

  useEffect(() => {
    if (user && notify) {
      showTestNotification(user.email)
      setNotify(false)
    }
  }, [user, showTestNotification, notify])

  const onLogout = async () => {
    setLoginOut(true)
    await logUserOut(queryClient)
  }

  function renderAvatarMenu() {
    return loginOut ? (
      <Group spacing={7} align="center">
        <Loader size={'sm'} variant="dots"></Loader>
        {loginOut ? 'Logging out...' : 'Logging in...'}
      </Group>
    ) : user ? (
      <UnstyledButton className={cx(classes.user, { [classes.userActive]: userMenuOpened })}>
        <Group spacing={7}>
          <Avatar alt={user.username} radius="xl" size={35} mt={6} mb={6} />
          <Text className="display-name" weight={500}>
            {user.username}
          </Text>
          <IconChevronDown size={12} stroke={1.5} />
        </Group>
      </UnstyledButton>
    ) : (
      <LoginButton />
    )
  }

  return (
    <>
      <Box
        css={css`
          position: sticky;
          top: 0;
          z-index: 10000;
        `}
      >
        <MantineHeader height={HEADER_HEIGHT} px="md" sx={{ height: '100%' }} className={classes.header}>
          <Group
            position="apart"
            css={css`
              align-self: center;
            `}
          >
            <Menu
              width={220}
              position="bottom-end"
              onClose={() => setUserMenuOpened(false)}
              onOpen={() => {
                if (user) setUserMenuOpened(true)
              }}
            >
              <Menu.Target>{renderAvatarMenu()}</Menu.Target>
              <Menu.Dropdown
                css={css`
                  p {
                    margin: 0px;
                  }
                `}
              >
                <Menu.Item onClick={() => setNotify(true)} icon={<IconHeart size={20} />}>
                  <Text fz="s">Test notification</Text>
                </Menu.Item>
                <Menu.Divider />
                <Menu.Item
                  onClick={() =>
                    Object.assign(document.createElement('a'), {
                      target: '_blank',
                      rel: 'noopener noreferrer',
                      href: CONFIG.AUTH_SERVER_UI_PROFILE,
                    }).click()
                  }
                  icon={<FontAwesomeIcon icon={faUser} size="xl" />}
                >
                  <Text fz="s">Profile</Text>
                </Menu.Item>
                <Menu.Divider />
                <ThemeSwitcher />
                <Menu.Divider />
                <Menu.Label>Settings</Menu.Label>
                <Menu.Item icon={<IconSettings size={14} stroke={1.5} />}>Account settings</Menu.Item>
                <Menu.Divider />
                <Menu.Item icon={<IconLogout size={14} stroke={1.5} />} onClick={onLogout}>
                  Logout
                </Menu.Item>
              </Menu.Dropdown>
            </Menu>
          </Group>
          <Container>
            <Tabs
              defaultValue="Home"
              variant="outline"
              classNames={{
                root: classes.tabs,
                tabsList: classes.tabsList,
                tab: classes.tab,
              }}
            >
              <Tabs.List>{items}</Tabs.List>
            </Tabs>
          </Container>
        </MantineHeader>
      </Box>
    </>
  )
}
