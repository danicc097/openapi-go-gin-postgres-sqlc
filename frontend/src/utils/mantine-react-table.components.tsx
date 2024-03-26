import { useMantineTheme, Menu, Button, rem, Text, ActionIcon, Tooltip } from '@mantine/core'
import { useDisclosure } from '@mantine/hooks'
import {
  IconDots,
  IconPackage,
  IconSquareCheck,
  IconUsers,
  IconCalendar,
  IconEdit,
  IconTrash,
  IconSend,
} from '@tabler/icons'
import { IconRestore } from '@tabler/icons-react'
import { useState } from 'react'

interface RowActionsMenuProps {
  canRestore: boolean
}

export function RowActionsMenu({ canRestore: canBeRestored }: RowActionsMenuProps) {
  const theme = useMantineTheme()
  const [loading, setLoading] = useState(false)

  return (
    <Menu transitionProps={{ transition: 'pop-bottom-left' }} width={220} withinPortal withArrow>
      <Menu.Target>
        <Tooltip label="Show actions" withArrow>
          <ActionIcon>
            <IconDots style={{ height: rem(18) }} stroke={1.5} />
          </ActionIcon>
        </Tooltip>
      </Menu.Target>
      <Menu.Dropdown>
        <Menu.Item leftSection={<IconEdit style={{ height: rem(16) }} color={theme.colors.blue[6]} stroke={1.5} />}>
          Edit
        </Menu.Item>
        <Menu.Item leftSection={<IconTrash style={{ height: rem(16) }} color={theme.colors.red[6]} stroke={1.5} />}>
          Delete
        </Menu.Item>
        {canBeRestored && (
          <Menu.Item
            leftSection={<IconRestore style={{ height: rem(16) }} color={theme.colors.yellow[6]} stroke={1.5} />}
          >
            Restore
          </Menu.Item>
        )}
        <Menu.Divider />
        <Menu.Item leftSection={<IconSend style={{ height: rem(16) }} color={theme.colors.cyan[6]} stroke={1.5} />}>
          Message members
        </Menu.Item>
      </Menu.Dropdown>
    </Menu>
  )
}
