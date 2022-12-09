import React, { Dispatch, useEffect, useMemo, useState } from 'react'

import {
  EuiAvatar,
  EuiIcon,
  EuiHeaderSection,
  EuiHeaderSectionItem,
  EuiHeaderSectionItemButton,
  EuiHeaderLinks,
  EuiHeaderLink,
  EuiPopover,
  EuiFlexGroup,
  EuiFlexItem,
  EuiLink,
  EuiHorizontalRule,
  htmlIdGenerator,
  EuiButton,
  EuiSpacer,
  EuiButtonEmpty,
  useEuiTour,
  EuiTourState,
  EuiTourStep,
} from '@elastic/eui'
import { Link, useNavigate } from 'react-router-dom'
import loginIcon from 'src/assets/img/loginIcon.svg'
import UserAvatar from '../UserAvatar/UserAvatar'
import CollapsibleNav from './CollapsibleNav/CollapsibleNav'
import { useNotificationAPI } from 'src/hooks/ui/useNotificationAPI'
import logoDark from 'src/assets/logo/two-white-clouds.svg'
import logoLight from 'src/assets/logo/two-black-clouds.svg'
import { AvatarMenu, StyledEuiHeader, LogoSection } from './Navbar.styles'
import { useUISlice } from 'src/slices/ui'
import { useAuthenticatedUser } from 'src/hooks/auth/useAuthenticatedUser'
import { ThemeSwitcher } from 'src/ThemeSwitcher/ThemeSwitcher'
import _ from 'lodash'
import config from '@config'

export default function Navbar() {
  const [avatarMenuOpen, setAvatarMenuOpen] = useState<boolean>(false)
  const { user, logUserOut, avatarColor } = useAuthenticatedUser()
  const navigate = useNavigate()
  const { showTestNotification } = useNotificationAPI()
  const theme = useUISlice((state) => state.theme)
  const [notify, setNotify] = useState<boolean>(false)
  const [logo, setLogo] = useState<string>(getLogo(theme))

  useEffect(() => {
    if (user && notify) {
      showTestNotification(user.email)
      setNotify(false)
    }
  }, [user, showTestNotification, notify])

  useEffect(() => {
    setLogo(getLogo(theme))
  }, [theme])

  function getLogo(theme: string) {
    return theme === 'dark' ? logoDark : logoLight
  }

  const toggleAvatarMenu = () => setAvatarMenuOpen(!avatarMenuOpen)
  const closeAvatarMenu = () => setAvatarMenuOpen(false)
  const handleLogout = () => {
    closeAvatarMenu()
    // logUserOut()
    navigate('/')
  }

  const avatarButton = (
    <EuiHeaderSectionItemButton
      aria-label="User avatar"
      data-test-subj="avatar"
      onClick={() => user?.email && toggleAvatarMenu()}
    >
      {user?.email ? (
        <UserAvatar size="l" user={user} color={avatarColor} initialsLength={2} />
      ) : (
        <Link to="/login">
          <EuiAvatar size="l" color="#1E90FF" name="user" imageUrl={loginIcon} />
        </Link>
      )}
    </EuiHeaderSectionItemButton>
  )

  const renderAvatarMenu = () => {
    if (!user?.email) return null
    return (
      <AvatarMenu>
        <EuiFlexGroup
          gutterSize="xs"
          direction="column"
          alignItems="flexEnd"
          justifyContent="spaceAround"
          className="avatar-dropdown"
          style={{ alignItems: 'center' }}
        >
          <EuiFlexItem style={{ alignItems: 'center', flexGrow: 1 }}>
            <EuiFlexGroup direction="row" alignItems="center">
              <EuiFlexItem>
                <UserAvatar size="l" user={user} color={avatarColor} initialsLength={2} />
              </EuiFlexItem>
              <EuiFlexGroup
                direction="column"
                justifyContent="spaceAround"
                // alignItems="center"
              >
                <EuiFlexItem>
                  <strong>{user?.username}</strong>
                </EuiFlexItem>
                <EuiFlexItem>{_.truncate(user?.email, { length: 25 })}</EuiFlexItem>
              </EuiFlexGroup>
            </EuiFlexGroup>
          </EuiFlexItem>

          <EuiHorizontalRule margin="m" />

          <ThemeSwitcher />

          <EuiHorizontalRule margin="m" />

          <EuiFlexGroup
            direction="row"
            alignItems="center"
            className="avatar-dropdown-actions"
            style={{ alignSelf: 'flex-start' }}
          >
            <EuiFlexItem grow={1}>
              <EuiIcon type="user" size="m" />
            </EuiFlexItem>
            <EuiFlexItem grow={8}>
              <EuiLink href={`${config.AUTH_SERVER_UI_PROFILE}`}>Profile</EuiLink>
            </EuiFlexItem>
          </EuiFlexGroup>

          <EuiSpacer size="s" />

          <EuiFlexGroup
            direction="row"
            alignItems="center"
            className="avatar-dropdown-actions"
            style={{ alignSelf: 'flex-start' }}
          >
            <EuiFlexItem grow={1}>
              <EuiIcon type="push" size="m" />
            </EuiFlexItem>
            <EuiFlexItem grow={3}>
              <a href="#" onClick={() => setNotify(true)}>
                Notification
              </a>
            </EuiFlexItem>
          </EuiFlexGroup>

          <EuiHorizontalRule margin="m" />

          <EuiFlexGroup
            direction="row"
            alignItems="center"
            justifyContent="flexStart"
            className="avatar-dropdown-actions"
            style={{ alignSelf: 'flex-start' }}
          >
            <EuiFlexItem grow={1}>
              <EuiIcon type="exit" size="m" />
            </EuiFlexItem>
            <EuiFlexItem grow={8}>
              <EuiLink onClick={handleLogout} color="danger" data-test-subj="logout">
                Log out
              </EuiLink>
            </EuiFlexItem>
          </EuiFlexGroup>
        </EuiFlexGroup>
      </AvatarMenu>
    )
  }

  return (
    <StyledEuiHeader
      sections={[
        {
          items: [
            user?.userID !== '' ? <CollapsibleNav user={user} /> : null,

            <LogoSection href="/" key={0}>
              <EuiIcon type={logo} size="l" />
            </LogoSection>,

            <EuiHeaderLinks aria-label="app navigation links" key={0}>
              <EuiHeaderLink
                iconType="documentation"
                target="_blank"
                href={import.meta.env.VITE_WIKI_URL}
                data-test-subj="wiki"
              >
                Wiki
              </EuiHeaderLink>
              <EuiHeaderLink
                iconType="help"
                onClick={() => {
                  navigate('/help')
                }}
                className="help-header-link"
              >
                Help
              </EuiHeaderLink>
            </EuiHeaderLinks>,
          ],
          borders: 'right',
        },
        {
          items: [
            <EuiPopover
              id="avatar-menu"
              key={'765'}
              isOpen={avatarMenuOpen}
              closePopover={closeAvatarMenu}
              anchorPosition="downRight"
              button={avatarButton}
              panelPaddingSize="m"
            >
              {renderAvatarMenu()}
            </EuiPopover>,
          ],
        },
      ]}
    ></StyledEuiHeader>
  )
}
