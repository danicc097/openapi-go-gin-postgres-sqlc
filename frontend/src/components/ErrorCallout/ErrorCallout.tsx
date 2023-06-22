import { Alert } from '@mantine/core'
import { IconAlertCircle } from '@tabler/icons'

export default function ErrorCallout({ title, errors }: { title: string; errors: string[] }) {
  return (
    errors?.length > 0 && (
      <Alert icon={<IconAlertCircle size={16} />} title={title} color="red">
        {errors.map((error, i) => (
          <li key={i}>{error}</li>
        ))}
      </Alert>
    )
  )
}
