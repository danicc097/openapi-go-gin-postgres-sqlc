import { Alert } from '@mantine/core'
import { IconAlertCircle } from '@tabler/icons'

interface ErrorCalloutProps {
  title: string
  errors?: string[]
}

export default function ErrorCallout({ title, errors }: ErrorCalloutProps) {
  if (!errors || errors.length === 0) return null

  return errors?.length > 0 ? (
    <Alert icon={<IconAlertCircle size={16} />} title={title} color="red">
      {errors.map((error, i) => (
        <div role="alert" key={i}>
          {error}
        </div>
      ))}
    </Alert>
  ) : null
}
