import type { SetStateAction } from 'react'
import ROLES from 'src/roles'

export function validateRole(role: string): boolean {
  return Object.keys(ROLES).includes(role)
}

export function _getFormErrors(
  form: any,
  errors: FormErrors<typeof form>,
  hasSubmitted: boolean,
  ...errorLists: Array<Array<unknown>>
) {
  const formErrors = []

  if (errors?.form) {
    formErrors.push(errors.form)
  }

  if (hasSubmitted && errorLists.some((list) => list.length)) {
    return formErrors.concat(errorLists.flat())
  }

  return formErrors
}

type validateFormBeforeSubmitParams = {
  form: any
  setErrors: SetStateAction<any>
  optionalFields?: Array<string>
}

export const validateFormBeforeSubmit = ({
  form,
  optionalFields,
  setErrors,
}: validateFormBeforeSubmitParams): boolean => {
  setErrors({})

  const _optionalFields = optionalFields || []

  Object.entries(form).forEach(([k, v]) => {
    if (!v && !_optionalFields.includes(k)) {
      setErrors((errors) => ({ ...errors, [k]: `${k} is required` }))
    }
    if (typeof v === 'string') form[k] = v.trim()
  })

  if (!Object.entries(form).every(([k, v]) => _optionalFields.includes(k) || !!v)) {
    setErrors((errors) => ({ ...errors, form: 'You must fill out all required fields' }))
    return false
  }
  return true
}
