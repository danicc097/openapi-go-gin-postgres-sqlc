import create from 'zustand'
import { devtools, persist } from 'zustand/middleware'
import { Alert } from '@mantine/core'
import { IconAlertCircle } from '@tabler/icons'
import { AxiosError } from 'axios'
import { useState } from 'react'
import { ApiError } from 'src/api/mutator'
import type { HTTPError } from 'src/gen/model'
import type { AppError } from 'src/types/ui'

export const FORM_SLICE_PERSIST_KEY = 'form-slice'

export type CalloutError = AppError | string | null
export type CalloutWarning = string | null

interface FormState {
  callout: {
    [formName: string]: {
      errors: Array<CalloutError>
      warnings: Array<CalloutWarning>
    }
  }
  setCalloutWarning: (formName: string, warning: CalloutWarning) => void
  setCalloutError: (formName: string, error: CalloutError) => void
}

const useFormSlice = create<FormState>()(
  devtools(
    // persist(
    (set) => {
      return {
        callout: {},
        setCalloutWarning: (formName: string, warning: string | null) =>
          set(
            (state) => {
              const form: FormState['callout'][string] = state.callout[formName] || { errors: [], warnings: [] }

              return {
                ...state,
                callout: {
                  ...state.callout,
                  [formName]: {
                    ...form,
                    warnings: warning ? [...form.warnings, warning] : form.warnings,
                  },
                },
              }
            },
            false,
            `setCalloutWarning`,
          ),
        setCalloutError: (formName: string, error: AppError | null) =>
          set(
            (state) => {
              const form: FormState['callout'][string] = state.callout[formName] || { errors: [], warnings: [] }

              return {
                ...state,
                callout: {
                  ...state.callout,
                  [formName]: {
                    ...form,
                    errors: error ? [...form.errors, error] : form.errors,
                  },
                },
              }
            },
            false,
            `setCalloutError`,
          ),
      }
    },
    // { version: 3, name: FORM_SLICE_PERSIST_KEY },
    // ),
    { enabled: true },
  ),
)

export { useFormSlice }
