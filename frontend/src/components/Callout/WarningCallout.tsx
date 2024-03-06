import { css } from '@emotion/react'
import { Alert, List } from '@mantine/core'
import { IconAlertCircle } from '@tabler/icons'

interface WarningCalloutProps {
  title: string
  warnings?: string[]
}

export default function WarningCallout({ title, warnings }: WarningCalloutProps) {
  if (!warnings || warnings.length === 0) return null

  return (
    <Alert mb={12} icon={<IconAlertCircle size={16} />} title={title} color="yellow">
      <List spacing="xs" size="sm" center mr={60}>
        {warnings.map((warning, i) => (
          <List.Item key={i}>{warning}</List.Item>
        ))}
      </List>
    </Alert>
  )
}
