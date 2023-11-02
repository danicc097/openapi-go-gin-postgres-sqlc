import { Alert } from '@mantine/core'
import { IconAlertCircle } from '@tabler/icons'
import { AxiosError } from 'axios'
import { useState } from 'react'
import { ApiError } from 'src/api/mutator'
import type { HTTPError } from 'src/gen/model'
import { CalloutError, useFormSlice } from 'src/slices/form'
import type { AppError } from 'src/types/ui'

export default function ErrorCallout({ title, formName }: { title: string; formName: string }) {
  const formSlice = useFormSlice()
  const errors = formSlice.form[formName]?.calloutErrors

  if (!errors) return null

  console.log({ errors })

  return errors?.length > 0 ? (
    <Alert icon={<IconAlertCircle size={16} />} title={title} color="red">
      {errors.map((error, i) => (
        <li key={i}>{renderCalloutError(error)}</li>
      ))}
    </Alert>
  ) : null
}

function renderCalloutError(error: CalloutError) {
  if (error instanceof ApiError || error instanceof AxiosError) {
    return error.message
  }

  return JSON.stringify(error)
}

export const useCalloutErrors = (formName: string) => {
  const formSlice = useFormSlice()
  const calloutErrors = formSlice.form[formName]?.calloutErrors
  const setCalloutErrors = (errors: CalloutError[]) => formSlice.setCalloutErrors(formName, errors)

  const extractCalloutErrors = () => {
    const errors: string[] = []

    if (!calloutErrors) return []

    for (const calloutError of calloutErrors) {
      // TODO: instead construct based on spec HTTPError which internally could have validationError array with loc, etc, see FastAPI template
      // or a regular error with message, title, detail, status...
      // and construct appropriately
      if (calloutError instanceof ApiError) {
        errors.push(calloutError.message)
        continue
      }

      // external call error
      if (calloutError instanceof AxiosError) {
        errors.push(calloutError.message)
        continue
      }

      if (typeof calloutError === 'string') {
        errors.push(calloutError)
        continue
      }

      // client side validation replaced by react hook form ajv resolver
      // error callout is just used for remote errors.
      // however we should also handle locs returned by backend (which have
      // no relation to schema validation). e.g. some field path is invalid because it already exists,
      // then we should set error in its input
      //
      // return calloutErrors?.errors?.map((v, i) => `${v.invalidParams.name}: ${v.invalidParams.reason}`)
    }

    return errors
  }

  const extractCalloutTitle = () => {
    if (!calloutErrors) return ''

    const unknownError = 'An unknown error ocurred'

    for (const calloutError of calloutErrors) {
      if (calloutError instanceof ApiError) {
        if (!calloutError.response?.data) {
          return unknownError
        }
        const error = calloutError.response?.data as HTTPError
        switch (error.type) {
          case 'RequestValidation':
            return error.title
          case 'Unauthenticated':
            return 'Unauthenticated'
          case 'Unauthorized':
            return 'Unauthorized'
          case 'Unknown':
          default:
            return unknownError
        }
      }

      // external call error
      if (calloutErrors instanceof AxiosError) return unknownError
    }

    // errors unrelated to api calls -> validation error
    return 'Validation error'
  }

  return {
    calloutErrors,
    extractCalloutErrors,
    setCalloutErrors,
    extractCalloutTitle,
  }
}
