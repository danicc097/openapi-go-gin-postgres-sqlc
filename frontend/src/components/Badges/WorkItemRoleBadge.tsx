import { Badge } from '@mantine/core'
import { capitalize } from 'lodash'
import { memo } from 'react'
import type { Role, WorkItemRole } from 'src/gen/model'
import { workItemRoleColor, getContrastYIQ } from 'src/utils/colors'

function _workItemRoleBadge({ role }: { role: WorkItemRole }) {
  const color = workItemRoleColor(role)
  const name = capitalize(role.replace(/([A-Z])/g, ' $1').trim())

  return (
    <Badge
      size="sm"
      radius="sm"
      style={{ backgroundColor: color, color: getContrastYIQ(color) === 'black' ? 'whitesmoke' : '#131313' }}
    >
      {name}
    </Badge>
  )
}

const WorkItemRoleBadge = memo(_workItemRoleBadge)

export default WorkItemRoleBadge
