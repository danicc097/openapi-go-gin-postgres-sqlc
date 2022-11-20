import React, { useState } from 'react'
import find from 'lodash/find'
import findIndex from 'lodash/findIndex'
import roles from '@roles'
import {
  EuiCollapsibleNavGroup,
  EuiHeaderSectionItemButton,
  EuiHeaderLogo,
  EuiHeader,
  EuiIcon,
  EuiButton,
  EuiButtonEmpty,
  EuiPageTemplate,
  EuiPinnableListGroup,
  EuiPinnableListGroupItemProps,
  EuiFlexItem,
  EuiHorizontalRule,
  EuiImage,
  EuiListGroup,
  useGeneratedHtmlId,
  EuiCollapsibleNav,
} from '@elastic/eui'
import { useNavigate } from 'react-router-dom'
import type { User } from 'src/redux/slices/gen/internalApi'

type CollapsibleNavProps = {
  user: User
}

const CollapsibleNav = ({ user }: CollapsibleNavProps) => {
  const navigate = useNavigate()

  const TopLinks: EuiPinnableListGroupItemProps[] = [
    {
      label: 'Home',
      iconType: 'home',
      isActive: true,
      'aria-current': true,
      onClick: () => {
        navigate('/')
      },
      pinnable: false,
    },
  ]

  const LearnLinks: EuiPinnableListGroupItemProps[] = [
    {
      label: 'Docs',
      onClick: () => {
        null
      },
    },
    {
      label: 'Blogs',
      onClick: () => {
        null
      },
    },
    {
      label: 'Webinars',
      onClick: () => {
        null
      },
    },
    { label: 'Elastic.co', href: 'https://elastic.co' },
  ]

  const SkillsLinks: EuiPinnableListGroupItemProps[] = [
    {
      label: 'Knowledge levels',
      onClick: () => {
        navigate('/knowledge-levels')
      },
    },
    {
      label: 'Knowledge map',
      onClick: () => {
        null
      },
    },
    {
      label: 'Progress',
      onClick: () => {
        null
      },
    },
  ]

  const AdminLinks: EuiPinnableListGroupItemProps[] = [
    {
      label: 'User verification',
      onClick: () => {
        navigate('/admin/unverified-users')
      },
    },
    {
      label: 'User password reset',
      onClick: () => {
        navigate('/admin/password-reset')
      },
    },
    {
      label: 'User password reset requests',
      onClick: () => {
        navigate('/admin/password-reset-requests')
      },
    },
    {
      label: 'User permissions management',
      onClick: () => {
        navigate('/admin/user-permissions-management')
      },
    },
  ]

  const [navIsOpen, setNavIsOpen] = useState<boolean>(
    JSON.parse(String(localStorage.getItem('euiCollapsibleNavExample--isOpen'))) || false,
  )
  const [navIsDocked, setNavIsDocked] = useState<boolean>(
    JSON.parse(String(localStorage.getItem('euiCollapsibleNavExample--isDocked'))) || false,
  )

  const adminGroup = 'Admin'
  const learnGroup = 'Learn'
  const skillsGroup = 'Skills'

  const [openGroups, setOpenGroups] = useState(
    JSON.parse(String(localStorage.getItem('openNavGroups'))) || ['Learn', 'Skills', 'Admin'],
  )

  const toggleAccordion = (isOpen: boolean, title?: string) => {
    if (!title) return
    const itExists = openGroups.includes(title)
    if (isOpen) {
      if (itExists) return
      openGroups.push(title)
    } else {
      const index = openGroups.indexOf(title)
      if (index > -1) {
        openGroups.splice(index, 1)
      }
    }
    setOpenGroups([...openGroups])
    localStorage.setItem('openNavGroups', JSON.stringify(openGroups))
  }

  const [pinnedItems, setPinnedItems] = useState<EuiPinnableListGroupItemProps[]>(
    JSON.parse(String(localStorage.getItem('pinnedItems'))) || [],
  )

  const addPin = (item: any) => {
    if (!item || find(pinnedItems, { label: item.label })) {
      return
    }
    item.pinned = true
    const newPinnedItems = pinnedItems ? pinnedItems.concat(item) : [item]
    setPinnedItems(newPinnedItems)
    localStorage.setItem('pinnedItems', JSON.stringify(newPinnedItems))
  }

  const removePin = (item: any) => {
    const pinIndex = findIndex(pinnedItems, { label: item.label })
    if (pinIndex > -1) {
      item.pinned = false
      const newPinnedItems = pinnedItems
      newPinnedItems.splice(pinIndex, 1)
      setPinnedItems([...newPinnedItems])
      localStorage.setItem('pinnedItems', JSON.stringify(newPinnedItems))
    }
  }

  function alterLinksWithCurrentState(
    links: EuiPinnableListGroupItemProps[],
    showPinned = false,
  ): EuiPinnableListGroupItemProps[] {
    return links.map((link) => {
      const { pinned, ...rest } = link
      return {
        pinned: showPinned ? pinned : false,
        ...rest,
      }
    })
  }

  function addLinkNameToPinTitle(listItem: EuiPinnableListGroupItemProps) {
    return `Pin ${listItem.label} to top`
  }

  function addLinkNameToUnpinTitle(listItem: EuiPinnableListGroupItemProps) {
    return `Unpin ${listItem.label}`
  }

  const collapsibleNavId = useGeneratedHtmlId({ prefix: 'collapsibleNav' })

  return (
    <EuiCollapsibleNav
      // className="eui-yScroll" // breaks right close button
      id={collapsibleNavId}
      aria-label="Main navigation"
      isOpen={navIsOpen}
      isDocked={navIsDocked}
      button={
        <EuiHeaderSectionItemButton
          aria-label="Toggle main navigation"
          onClick={() => {
            setNavIsOpen(!navIsOpen)
            localStorage.setItem('euiCollapsibleNavExample--isOpen', JSON.stringify(!navIsOpen))
          }}
        >
          <EuiIcon type={'menu'} size="m" aria-hidden="true" />
        </EuiHeaderSectionItemButton>
      }
      paddingSize="none"
      onClose={() => {
        setNavIsOpen(false)
        localStorage.setItem('euiCollapsibleNavExample--isOpen', JSON.stringify(false))
      }}
      maskProps={{ headerZindexLocation: 'below' }}
    >
      <EuiFlexItem grow={false} style={{ flexShrink: 0 }}>
        <EuiCollapsibleNavGroup background="light" style={{ maxHeight: '40vh' }}>
          <EuiPinnableListGroup
            aria-label="Pinned links"
            listItems={alterLinksWithCurrentState(TopLinks).concat(alterLinksWithCurrentState(pinnedItems, true))}
            unpinTitle={addLinkNameToUnpinTitle}
            onPinClick={removePin}
            maxWidth="none"
            color="text"
            gutterSize="none"
            size="s"
          />
        </EuiCollapsibleNavGroup>
      </EuiFlexItem>

      <EuiFlexItem grow={false} style={{ flexShrink: 0 }}>
        {/* <EuiCollapsibleNavGroup isCollapsible={false} background="dark">
          <EuiListGroup
            color="ghost"
            maxWidth="none"
            gutterSize="none"
            size="s"
            listItems={[
              {
                label: 'Manage deployment',
                href: '#',
                iconType: 'logoCloud',
                iconProps: {
                  color: 'ghost',
                },
              },
            ]}
          />
        </EuiCollapsibleNavGroup> */}
      </EuiFlexItem>

      <EuiHorizontalRule margin="none" />

      {/* <EuiFlexItem>
        <EuiCollapsibleNavGroup
          title={
            <a href="#/navigation/collapsible-nav" onClick={(e) => e.stopPropagation()}>
              Training
            </a>
          }
          iconType="training"
          isCollapsible
          initialIsOpen={openGroups.includes(learnGroup)}
          onToggle={(isOpen: boolean) => toggleAccordion(isOpen, learnGroup)}
        >
          <EuiPinnableListGroup
            aria-label={learnGroup}
            listItems={alterLinksWithCurrentState(LearnLinks)}
            pinTitle={addLinkNameToPinTitle}
            onPinClick={addPin}
            maxWidth="none"
            color="subdued"
            gutterSize="none"
            size="s"
          />
        </EuiCollapsibleNavGroup>
      </EuiFlexItem> */}

      <EuiFlexItem
      // className="eui-yScroll"
      >
        {/* <EuiCollapsibleNavGroup
          title={
            <a
              onClick={(e) => {
                e.stopPropagation()
              }}
            >
              My skills
            </a>
          }
          iconType="canvasApp"
          isCollapsible
          initialIsOpen={openGroups.includes(skillsGroup)}
          onToggle={(isOpen: boolean) => toggleAccordion(isOpen, skillsGroup)}
        >
          <EuiPinnableListGroup
            aria-label={skillsGroup}
            listItems={alterLinksWithCurrentState(SkillsLinks)}
            pinTitle={addLinkNameToPinTitle}
            onPinClick={addPin}
            maxWidth="none"
            color="subdued"
            gutterSize="none"
            size="s"
          />
        </EuiCollapsibleNavGroup> */}

        {user?.role_rank > roles.admin.rank ? (
          <EuiCollapsibleNavGroup
            title={
              <a
                onClick={(e) => {
                  e.stopPropagation()
                }}
              >
                Admin panel
              </a>
            }
            iconType="securityApp"
            isCollapsible
            initialIsOpen={openGroups.includes(adminGroup)}
            onToggle={(isOpen: boolean) => toggleAccordion(isOpen, adminGroup)}
          >
            <EuiPinnableListGroup
              aria-label={adminGroup}
              listItems={alterLinksWithCurrentState(AdminLinks)}
              pinTitle={addLinkNameToPinTitle}
              onPinClick={addPin}
              maxWidth="none"
              color="subdued"
              gutterSize="none"
              size="s"
            />
          </EuiCollapsibleNavGroup>
        ) : null}
      </EuiFlexItem>

      <EuiFlexItem grow={false}>
        <span />
        <EuiCollapsibleNavGroup>
          <EuiButton
            fullWidth
            onClick={() => {
              setNavIsDocked(!navIsDocked)
              localStorage.setItem('euiCollapsibleNavExample--isDocked', JSON.stringify(!navIsDocked))
            }}
          >
            {navIsDocked ? 'Undock sidebar' : 'Dock sidebar'}
          </EuiButton>
        </EuiCollapsibleNavGroup>
      </EuiFlexItem>
    </EuiCollapsibleNav>
  )
}

export default CollapsibleNav
