import { Alert, List } from '@mantine/core'
import { IconAlertCircle } from '@tabler/icons'

interface ErrorCalloutProps {
  title: string
  errors?: string[]
}

export default function ErrorCallout({ title, errors }: ErrorCalloutProps) {
  if ((!errors || errors.length === 0) && !title) return null

  return (
    <Alert mb={12} icon={<IconAlertCircle size={16} />} title={title} color="red">
      {errors && (
        <List pb={6} spacing="xs" size="sm" center mr={60}>
          {errors.map((error, i) => (
            <List.Item key={i}>{error}</List.Item>
          ))}
        </List>
      )}
    </Alert>
  )
}
