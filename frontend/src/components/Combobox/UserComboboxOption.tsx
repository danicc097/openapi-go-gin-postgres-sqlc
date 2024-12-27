import { Avatar, Group, Space } from '@mantine/core'
import RoleBadge from 'src/components/Badges/RoleBadge'
import type { UserResponse } from 'src/gen/model'
import { nameInitials } from 'src/utils/strings'

interface UserComboboxOptionProps {
  user: UserResponse
}

export default function UserComboboxOption({ user }: UserComboboxOptionProps) {
  return (
    <Group align="center">
      <div style={{ display: 'flex', alignItems: 'center', maxHeight: 1 }}>
        <Avatar size={28} radius="xl" data-test-id="header-profile-avatar" alt={user?.username}>
          {nameInitials(user.fullName || '')}
        </Avatar>
        <Space p={5} />
        <RoleBadge role={user.role} />
      </div>

      <div style={{ marginLeft: 'auto' }}>{user?.email}</div>
    </Group>
  )
}
