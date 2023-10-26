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
  Box,
  Tooltip,
  Badge,
  useMantineTheme,
  useMantineColorScheme,
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
import logoDark from 'src/assets/logo/two-white-clouds.svg'
import logoLight from 'src/assets/logo/two-black-clouds.svg'
import classes from './Header.module.css'
import cx from 'clsx'

interface HeaderProps {
  tabs: string[]
}

export const HEADER_HEIGHT = 60

export default function Header({ tabs }: HeaderProps) {
  const queryClient = useQueryClient()
  const navigate = useNavigate()
  const theme = useMantineTheme()
  // const [opened, { toggle }] = useDisclosure(false)
  const [userMenuOpened, setUserMenuOpened] = useState(false)
  const { user } = useAuthenticatedUser()
  const [loginOut, setLoginOut] = useState(false)
  const { colorScheme } = useMantineColorScheme() // TODO: app logo useffect

  const items = tabs.map((tab) => (
    <Tabs.Tab value={tab} key={tab}>
      {tab}
    </Tabs.Tab>
  ))
  const { burgerOpened, setBurgerOpened } = useUISlice()

  const [notify, setNotify] = useState<boolean>(false)
  const { showTestNotification } = useNotificationAPI()
  const [logo, setLogo] = useState<string>(colorScheme === 'dark' ? logoDark : logoLight)

  useEffect(() => {
    setLogo(colorScheme === 'dark' ? logoDark : logoLight)
  }, [theme])

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
    if (loginOut)
      return (
        <Group gap={'md'} align="center">
          <Loader size={'sm'} variant="dots"></Loader>
          Logging out...
        </Group>
      )

    return user ? (
      <UnstyledButton className={cx(classes.user, { [classes.userActive]: userMenuOpened })}>
        <Group gap={'md'} m={4}>
          <Avatar alt={user.username} radius="xl" size={35} mt={6} mb={6} />
          <Text
            className="display-name"
            css={`
              font-weight: 500;
            `}
          >
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
          z-index: 100;
        `}
      >
        {/* TODO: v7 is https://mantine.dev/core/app-shell/ */}
        <MantineHeader height={HEADER_HEIGHT} px="md" sx={{ height: '100%' }} className={classes.header}>
          <Group
            m={6}
            css={css`
              align-self: center;
              position: apart;
            `}
          >
            <a href="/">
              <img src={logo} height={HEADER_HEIGHT * 0.5} css={css``}></img>
            </a>
            <Menu
              width={220}
              position="bottom-end"
              onClose={() => setUserMenuOpened(false)}
              onOpen={() => {
                if (user) setUserMenuOpened(true)
              }}
              disabled={!user}
            >
              <Menu.Target>{renderAvatarMenu()}</Menu.Target>
              <Menu.Dropdown
                css={css`
                  p {
                    margin: 0px;
                  }
                `}
              >
                <Menu.Item onClick={() => setNotify(true)} leftSection={<IconHeart size={20} />}>
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
                  leftSection={<FontAwesomeIcon icon={faUser} size="xl" />}
                >
                  <Text fz="s">Profile</Text>
                </Menu.Item>
                <Menu.Divider />
                <ThemeSwitcher />
                <Menu.Divider />
                <Menu.Label>Settings</Menu.Label>
                <Menu.Item leftSection={<IconSettings size={14} stroke={1.5} />}>Account settings</Menu.Item>
                <Menu.Divider />
                <Menu.Item leftSection={<IconLogout size={14} stroke={1.5} />} onClick={onLogout}>
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
                tabSection: classes.tabsList,
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
