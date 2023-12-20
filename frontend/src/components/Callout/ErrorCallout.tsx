import { Alert } from '@mantine/core'
import { IconAlertCircle } from '@tabler/icons'
import { AxiosError } from 'axios'
import { useState } from 'react'
import { ApiError } from 'src/api/mutator'
import type { HTTPError } from 'src/gen/model'
import { type CalloutError, useFormSlice } from 'src/slices/form'
import type { AppError } from 'src/types/ui'
import { entries } from 'src/utils/object'
interface ErrorCalloutProps {
  title: string
  errors?: string[]
}

export default function ErrorCallout({ title, errors }: ErrorCalloutProps) {
  if (!errors || errors.length === 0) return null

  return errors?.length > 0 ? (
    <Alert icon={<IconAlertCircle size={16} />} title={title} color="red">
      {errors.map((error, i) => (
        <li key={i}>{error}</li>
      ))}
    </Alert>
  ) : null
}

export const useCalloutErrors = (formName: string) => {
  const formSlice = useFormSlice()
  const calloutErrors = formSlice.form[formName]?.calloutErrors
  const customErrors = formSlice.form[formName]?.customErrors
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
    console.log(formSlice.form)
    if (formSlice.form[formName]?.customErrors) {
      return 'Validation error'
    }
    const unknownError = 'An unknown error ocurred'

    if (!calloutErrors) {
      return unknownError
    }

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

    // errors unrelated to api calls and validation
    return unknownError
  }

  return {
    calloutErrors,
    customErrors,
    extractCalloutErrors,
    setCalloutErrors,
    extractCalloutTitle,
  }
}
