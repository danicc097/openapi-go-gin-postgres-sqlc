import { Alert } from '@mantine/core'
import { IconAlertCircle } from '@tabler/icons'
import { AxiosError } from 'axios'
import { concat, lowerFirst } from 'lodash'
import { randexp } from 'randexp'
import { useState } from 'react'
import { useFormContext, useFormState } from 'react-hook-form'
import { ApiError } from 'src/api/mutator'
import ErrorCallout from 'src/components/Callout/ErrorCallout'
import WarningCallout from 'src/components/Callout/WarningCallout'
import { useCalloutErrors } from 'src/components/Callout/useCalloutErrors'
import type { HTTPError } from 'src/gen/model'
import { type CalloutError, useFormSlice } from 'src/slices/form'
import type { AppError } from 'src/types/ui'
import { flattenRHFError, type SchemaKey } from 'src/utils/form'
import { useDynamicFormContext } from 'src/utils/formGeneration.context'
import { entries } from 'src/utils/object'

/**
 * Shows errors and warnings of the current context dynamic form
 */
export default function DynamicFormErrorCallout() {
  const formSlice = useFormSlice()
  const form = useFormContext()
  const { formName, options, schemaFields } = useDynamicFormContext()
  const {
    extractCalloutErrors,
    calloutWarnings,
    setCalloutErrors,
    calloutErrors,
    extractCalloutTitle,
    hasClickedSubmit,
  } = useCalloutErrors(formName)
  const formState = useFormState({ control: form.control })

  const rhfErrors = flattenRHFError({
    obj: formState.errors,
  })

  const title = extractCalloutTitle()

  const warnings = calloutWarnings ? entries(calloutWarnings).map(([schemaKey, warning], idx) => warning) : []
  const errors = concat(
    extractCalloutErrors(),
    entries(rhfErrors).map(([schemaKey, error], _) => {
      let message = lowerFirst(error.message) // lowerCase breaks regexes
      schemaKey = schemaKey.replace(/\.\d+$/, '') as SchemaKey // TODO: in flattener instead

      const itemName = options.labels[schemaKey] || ''

      if (error.index) {
        message = `item ${error.index + 1} ${message}`
      }

      const match = /match pattern "(.*?)"/g.exec(message)
      if (match) {
        message = `${message} (example: ${randexp(match[1] || '')})`
      }

      return `${itemName}: ${message}`
    }),
  )

  return (
    <>
      {errors?.length > 0 ? <ErrorCallout title={title} errors={errors} /> : null}
      {warnings?.length > 0 && !hasClickedSubmit ? <WarningCallout title={'Warning'} warnings={warnings} /> : null}
    </>
  )
}
