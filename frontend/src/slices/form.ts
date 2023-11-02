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

export type CalloutError = AppError | string
export type CalloutWarning = string

interface Form {
  errors: CalloutError[]
  warnings: CalloutWarning[]
  // indexed by formField. Used for errors that aren't registered in react hook form
  customErrors: Record<string, string>
}

interface FormState {
  form: {
    [formName: string]: Form
  }
  setCalloutWarning: (formName: string, warning: CalloutWarning[]) => void
  setCalloutErrors: (formName: string, error: CalloutError[]) => void
}

const initialForm: Form = { errors: [], warnings: [], customErrors: {} }

const useFormSlice = create<FormState>()(
  devtools(
    // persist(
    (set) => {
      return {
        form: {},
        setCalloutWarning: (formName: string, warnings: CalloutWarning[]) =>
          set(
            (state) => {
              const form = state.form[formName] || initialForm

              return {
                ...state,
                form: {
                  ...state.form,
                  [formName]: {
                    ...form,
                    warnings: warnings,
                  },
                },
              }
            },
            false,
            `setCalloutWarning`,
          ),
        setCalloutErrors: (formName: string, errors: CalloutError[]) =>
          set(
            (state) => {
              const form = state.form[formName] || initialForm

              return {
                ...state,
                form: {
                  ...state.form,
                  [formName]: {
                    ...form,
                    errors: errors,
                  },
                },
              }
            },
            false,
            `setCalloutErrors`,
          ),
        setCustomError: (formName: string, formField: string, error: string) =>
          set(
            (state) => {
              const form = state.form[formName] || initialForm

              return {
                ...state,
                form: {
                  ...state.form,
                  [formName]: {
                    ...form,
                    customErrors: {
                      ...form.customErrors,
                      [formField]: error,
                    },
                  },
                },
              }
            },
            false,
            `setCalloutErrors`,
          ),
        resetCustomErrors: (formName: string) =>
          set(
            (state) => {
              const form = state.form[formName] || initialForm

              return {
                ...state,
                form: {
                  ...state.form,
                  [formName]: {
                    ...form,
                    customErrors: {},
                  },
                },
              }
            },
            false,
            `setCalloutErrors`,
          ),
      }
    },
    // { version: 3, name: FORM_SLICE_PERSIST_KEY },
    // ),
    { enabled: true },
  ),
)

export { useFormSlice }
