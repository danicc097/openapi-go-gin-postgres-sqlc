import { useEffect, useState } from 'react'
import type { FormErrors } from 'src/hooks/utils/form'
import type { schemas } from 'src/types/schema'

export const useCreateUserForm = () => {
  const defaultForm: schemas['CreateUserRequest'] & { passwordConfirm: string } = {
    username: '',
    email: '',
    password: '',
    passwordConfirm: '',
  }
  const [form, setForm] = useState(defaultForm)

  const [errors, setErrors] = useState<FormErrors<typeof form>>({})
  const [hasSubmitted, setHasSubmitted] = useState(false)

  const handlePasswordConfirmChange = (value: string) => {
    setErrors((errors) => ({
      ...errors,
      passwordConfirm: form.password !== value ? 'Passwords do not match' : null,
    }))

    setForm((form) => ({ ...form, passwordConfirm: value }))
  }

  return {
    form,
    setForm,
    errors,
    setErrors,
    hasSubmitted,
    setHasSubmitted,
    handlePasswordConfirmChange,
    defaultForm,
  }
}
