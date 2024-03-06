import { Alert } from '@mantine/core'
import { IconAlertCircle } from '@tabler/icons'

interface WarningCalloutProps {
  title: string
  warnings?: string[]
}

export default function WarningCallout({ title, warnings }: WarningCalloutProps) {
  if (!warnings || warnings.length === 0) return null

  return (
    <Alert mb={12} icon={<IconAlertCircle size={16} />} title={title} color="yellow">
      {warnings.map((warning, i) => (
        <div role="alert" key={i}>
          {warning}
        </div>
      ))}
    </Alert>
  )
}
