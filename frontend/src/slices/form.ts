import { devtools, persist } from 'zustand/middleware'
import { Alert } from '@mantine/core'
import { IconAlertCircle } from '@tabler/icons'
import { AxiosError } from 'axios'
import { useState } from 'react'
import { ApiError } from 'src/api/mutator'
import type { HTTPError } from 'src/gen/model'
import type { AppError } from 'src/types/ui'
import { create } from 'zustand'

export const FORM_SLICE_PERSIST_KEY = 'form-slice'

export type CalloutError = AppError | string
export type CalloutWarning = string

interface Form {
  calloutErrors: CalloutError[]
  customWarnings: Record<string, string | null>
  // indexed by formField. Used for errors that aren't registered in react hook form
  customErrors: Record<string, string | null>
}

interface FormState {
  form: {
    [formName: string]: Form
  }
  setCalloutErrors: (formName: string, error: CalloutError[]) => void
  setCustomWarning: (formName: string, formField: string, warning: string | null) => void
  resetCustomWarnings: (formName: string) => void
  setCustomError: (formName: string, formField: string, error: string | null) => void
  resetCustomErrors: (formName: string) => void
}

const initialForm: Form = { calloutErrors: [], customWarnings: {}, customErrors: {} }

const useFormSlice = create<FormState>()(
  devtools(
    // persist(
    (set) => {
      return {
        form: {},
        setCalloutErrors: (formName: string, errors: CalloutError[]) =>
          set((state) => {
            const form = state.form[formName] || initialForm

            return {
              ...state,
              form: {
                ...state.form,
                [formName]: {
                  ...form,
                  calloutErrors: errors,
                },
              },
            }
          }),
        setCustomError: (formName: string, formField: string, error: string | null) =>
          set((state) => {
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
          }),
        resetCustomErrors: (formName: string) =>
          set((state) => {
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
          }),

        setCustomWarning: (formName: string, formField: string, warning: string | null) =>
          set((state) => {
            const form = state.form[formName] || initialForm

            if (warning === null) {
              delete state.form[formName]?.customWarnings?.[formField]
              return state
            }

            return {
              ...state,
              form: {
                ...state.form,
                [formName]: {
                  ...form,
                  customWarnings: {
                    ...form.customWarnings,
                    [formField]: warning,
                  },
                },
              },
            }
          }),
        resetCustomWarnings: (formName: string) =>
          set((state) => {
            const form = state.form[formName] || initialForm

            return {
              ...state,
              form: {
                ...state.form,
                [formName]: {
                  ...form,
                  customWarnings: {},
                },
              },
            }
          }),
      }
    },
    // { version: 3, name: FORM_SLICE_PERSIST_KEY },
    // ),
    { enabled: true },
  ),
)

export { useFormSlice }
