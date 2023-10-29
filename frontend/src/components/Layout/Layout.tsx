import { css, jsx } from '@emotion/react'
import { useEffect, useRef, useState } from 'react'
import { Helmet } from 'react-helmet'
// import Navbar from '../Navbar/Navbar'
import { Fragment } from 'react'
import shallow from 'zustand/shallow'
import {
  ActionIcon,
  ActionIconGroup,
  AppShell,
  Avatar,
  Drawer,
  Flex,
  Group,
  Loader,
  Menu,
  Skeleton,
  Tabs,
  Text,
  Tooltip,
  UnstyledButton,
  useMantineColorScheme,
  useMantineTheme,
  Container,
} from '@mantine/core'
import classes from './Layout.module.css'
import {
  IconLogout,
  IconHeart,
  IconSettings,
  IconChevronDown,
  IconBrandTwitter,
  IconBrandYoutube,
  IconBrandInstagram,
} from '@tabler/icons'
import useAuthenticatedUser, { logUserOut } from 'src/hooks/auth/useAuthenticatedUser'
import { useQueryClient } from '@tanstack/react-query'
import { useNavigate } from 'react-router-dom'
import { useGetCurrentUser } from 'src/gen/user/user'
import { useNotificationAPI } from 'src/hooks/ui/useNotificationAPI'
import CONFIG from 'src/config'
import { faUser } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import logoDark from 'src/assets/logo/two-white-clouds.svg'
import logoLight from 'src/assets/logo/two-black-clouds.svg'
import { useUISlice } from 'src/slices/ui'
import cx from 'clsx'
import LoginButton from 'src/components/LoginButton'
import { useDisclosure } from '@mantine/hooks'
import { ThemeSwitcher } from 'src/components/ThemeSwitcher'
import TestMantineV7 from 'src/components/Layout/TestMantineV7'

type LayoutProps = {
  children: React.ReactElement
}

export default function Layout({ children }: LayoutProps) {
  const queryClient = useQueryClient()
  const navigate = useNavigate()
  const theme = useMantineTheme()
  const [opened, { toggle }] = useDisclosure(false)
  const [userMenuOpened, setUserMenuOpened] = useState(false)
  const { user } = useAuthenticatedUser()
  const [loginOut, setLoginOut] = useState(false)
  const { colorScheme } = useMantineColorScheme() // TODO: app logo useffect
  const { burgerOpened, setBurgerOpened } = useUISlice()

  const tabs = []
  const items = tabs.map((tab) => (
    <Tabs.Tab value={tab} key={tab}>
      {tab}
    </Tabs.Tab>
  ))

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
      <UnstyledButton className={cx(classes.user, { [classes.userActive as string]: userMenuOpened })}>
        <Group gap={'md'} m={4}>
          <Avatar alt={user.username} radius="xl" size={35} mt={6} mb={6} />
          <Text className={classes.displayName} fw={500}>
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
    <Fragment>
      <Helmet>
        <meta charSet="utf-8" />
        <title>My APP</title>
        <link rel="canonical" href="#" />
      </Helmet>
      <AppShell
        className={classes.appShell}
        header={{ height: 60 }}
        footer={{ height: 60 }}
        navbar={{ width: 300, breakpoint: 'sm', collapsed: { mobile: !opened } }}
        // aside={{ width: 300, breakpoint: 'md', collapsed: { desktop: false, mobile: true } }}
        padding="md"
      >
        <AppShell.Header>
          <Group
            m={6}
            css={css`
              align-self: center;
              position: apart;
            `}
          >
            <a href="/">
              <img src={logo} css={css`var(--header-height) * 0.5`}></img>
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
        </AppShell.Header>
        <AppShell.Navbar p="md">
          Navbar
          {Array(15)
            .fill(0)
            .map((_, index) => (
              <Skeleton key={index} h={28} mt="sm" animate={false} />
            ))}
        </AppShell.Navbar>
        <AppShell.Main>
          {/* <TestMantineV7 /> */}
          {children}
        </AppShell.Main>
        {/* <AppShell.Aside p="md">Aside</AppShell.Aside> */}
        <AppShell.Footer p="md">
          <div className={classes.footer}>
            <Container className={classes.inner}>
              <Group align="left">
                <Text size={'xs'}>Copyright © {new Date().getFullYear()}</Text>
                <Text size={'xs'}>Build version: {CONFIG.BUILD_VERSION ?? 'DEVELOPMENT'}</Text>
              </Group>
              <Group gap={15} className={classes.links} align="right">
                <Tooltip label={`Follow us on Twitter`}>
                  <ActionIcon size="sm" variant="subtle">
                    <a href="#" target="_blank" rel="noopener noreferrer">
                      <IconBrandTwitter size={18} stroke={1.5} color="#2d8bb3" />
                    </a>
                  </ActionIcon>
                </Tooltip>
                <Tooltip label={`Follow us on YouTube`}>
                  <ActionIcon size="sm" variant="subtle">
                    <a href="#" target="_blank" rel="noopener noreferrer">
                      <IconBrandYoutube size={18} stroke={1.5} color="#d63808" />
                    </a>
                  </ActionIcon>
                </Tooltip>
                <Tooltip label={`Follow us on Instagram`}>
                  <ActionIcon size="sm" variant="subtle">
                    <a href="#" target="_blank" rel="noopener noreferrer">
                      <IconBrandInstagram size={18} stroke={1.5} color="#e15d16" />
                    </a>
                  </ActionIcon>
                </Tooltip>
              </Group>
            </Container>
          </div>
        </AppShell.Footer>
      </AppShell>

      {/* </ThemeProvider> */}
    </Fragment>
  )
}
