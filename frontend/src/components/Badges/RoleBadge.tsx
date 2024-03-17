import { Badge } from '@mantine/core'
import { capitalize } from 'lodash'
import { memo } from 'react'
import type { Role } from 'src/gen/model'
import { roleColor, getContrastYIQ } from 'src/utils/colors'

const RoleBadge = memo(function ({ role }: { role: Role }) {
  const color = roleColor(role)
  const name = capitalize(role.replace(/([A-Z])/g, ' $1').trim())

  return (
    <Badge
      size="sm"
      radius="md"
      style={{ backgroundColor: color, color: getContrastYIQ(color) === 'black' ? 'whitesmoke' : '#131313' }}
    >
      {name}
    </Badge>
  )
})

export default RoleBadge
