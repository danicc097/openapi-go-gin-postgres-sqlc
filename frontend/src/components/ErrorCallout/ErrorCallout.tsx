import { Alert } from '@mantine/core'
import { IconAlertCircle } from '@tabler/icons'
import { AxiosError } from 'axios'
import { useState } from 'react'
import { ApiError } from 'src/api/mutator'
import type { AppError } from 'src/types/ui'

export default function ErrorCallout({ title, errors }: { title: string; errors: string[] }) {
  return errors?.length > 0 ? (
    <Alert icon={<IconAlertCircle size={16} />} title={title} color="red">
      {errors.map((error, i) => (
        <li key={i}>{error}</li>
      ))}
    </Alert>
  ) : null
}

export const useCalloutErrors = () => {
  const [calloutErrors, setCalloutErrors] = useState<AppError | null>(null)

  const extractCalloutErrors = () => {
    if (!calloutErrors) return []

    // TODO: instead construct based on spec HTTPError which internally could have validationError array with loc, etc, see FastAPI template
    // or a regular error with message, title, detail, status...
    // and construct appropriately
    if (calloutErrors instanceof ApiError) return [calloutErrors.message]

    // external call error
    if (calloutErrors instanceof AxiosError) return [calloutErrors.message]

    return []

    // client side validation replaced by react hook form ajv resolver
    // error callout is just used for remote errors.
    // however we should also handle locs returned by backend (which have
    // no relation to schema validation). e.g. some field path is invalid because it already exists,
    // then we should set error in its input
    //
    // return calloutErrors?.errors?.map((v, i) => `${v.invalidParams.name}: ${v.invalidParams.reason}`)
  }

  return {
    calloutErrors,
    extractCalloutErrors,
    setCalloutErrors,
  }
}
