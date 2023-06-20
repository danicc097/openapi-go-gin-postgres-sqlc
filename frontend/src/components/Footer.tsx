import { createStyles, Container, Group, ActionIcon, Image, Text, Tooltip, Avatar } from '@mantine/core'
import { IconBrandTwitter, IconBrandYoutube, IconBrandInstagram, IconBrandTwitch } from '@tabler/icons'
import { Dropdown } from 'mantine-design-system'
import CONFIG from 'src/config'

export const FOOTER_HEIGHT = 55

const useStyles = createStyles((theme) => ({
  footer: {
    borderTop: `1px solid ${theme.colorScheme === 'dark' ? theme.colors.dark[5] : theme.colors.gray[2]}`,
    backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[8] : theme.white,
    position: 'fixed',
    bottom: 0,
    left: 0,
    width: '100%',
    zIndex: 999,
  },

  inner: {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    minWidth: '95vw',
    minHeight: FOOTER_HEIGHT,

    [theme.fn.smallerThan('xs')]: {
      flexDirection: 'column',
    },
  },

  links: {
    [theme.fn.smallerThan('xs')]: {
      marginTop: theme.spacing.md,
    },
  },
}))

export default function Footer() {
  const { classes } = useStyles()

  return (
    <div className={classes.footer}>
      <Container className={classes.inner}>
        <Text fz="xs">
          <Group position="left" noWrap>
            <p>Copyright © {new Date().getFullYear()}</p>
            <p>Build version: {CONFIG.BUILD_NUMBER ?? 'DEVELOPMENT'}</p>
          </Group>
        </Text>
        <Dropdown
          position="top"
          control={
            <ActionIcon variant="transparent">
              <Avatar size={35} radius="xl" data-test-id="header-profile-avatar" />
            </ActionIcon>
          }
        >
          <Dropdown.Item key="user" onClick={(e) => console.log('user 1')}>
            user 1
          </Dropdown.Item>
          <Dropdown.Item key="user2" onClick={(e) => console.log('user 2')}>
            user 2
          </Dropdown.Item>
        </Dropdown>
        <Group spacing={0} className={classes.links} position="right" noWrap>
          <Tooltip label={`Follow us on Twitter`}>
            <ActionIcon size="lg">
              <a href="#" target="_blank" rel="noopener noreferrer">
                <IconBrandTwitter size={18} stroke={1.5} color="#2d8bb3" />
              </a>
            </ActionIcon>
          </Tooltip>
          <Tooltip label={`Follow us on YouTube`}>
            <ActionIcon size="lg">
              <a href="#" target="_blank" rel="noopener noreferrer">
                <IconBrandYoutube size={18} stroke={1.5} color="#d63808" />
              </a>
            </ActionIcon>
          </Tooltip>
          <Tooltip label={`Follow us on Instagram`}>
            <ActionIcon size="lg">
              <a href="#" target="_blank" rel="noopener noreferrer">
                <IconBrandInstagram size={18} stroke={1.5} color="#e15d16" />
              </a>
            </ActionIcon>
          </Tooltip>
          <Tooltip label={`Follow us on Twitch`}>
            <ActionIcon size="lg">
              <a href="#" target="_blank" rel="noopener noreferrer">
                <IconBrandTwitch size={18} stroke={1.5} color="#a970ff" />
              </a>
            </ActionIcon>
          </Tooltip>
        </Group>
      </Container>
    </div>
  )
}
