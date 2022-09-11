import { useEffect, useState } from 'react'
import type { FormErrors } from 'src/hooks/utils/form'
import { useCreateUserMutation } from 'src/redux/slices/gen/internalApi'
import type { schemas } from 'src/types/schema'

export const useCreateUserForm = () => {
  const defaultForm: schemas['CreateUserRequest'] & { passwordConfirm: string } = {
    username: '',
    email: '',
    password: '',
    passwordConfirm: '',
  }
  const [form, setForm] = useState(defaultForm)

  const [formErrors, setFormErrors] = useState<FormErrors<typeof form>>({})
  const [_, createUserResult] = useCreateUserMutation()
  // use api slice's isUninitialized
  // const [hasSubmitted, setHasSubmitted] = useState(false)

  const handleChange = (value: string) => {
    setFormErrors((formErrors) => ({
      ...formErrors,
      passwordConfirm: form.password !== value ? 'Passwords do not match' : null,
    }))

    setForm((form) => ({ ...form, passwordConfirm: value }))
  }

  // const handlePasswordConfirmChange = (value: string) => {
  //   setFormErrors((formErrors) => ({
  //     ...formErrors,
  //     passwordConfirm: form.password !== value ? 'Passwords do not match' : null,
  //   }))

  //   setForm((form) => ({ ...form, passwordConfirm: value }))
  // }

  return {
    form,
    setForm,
    formErrors,
    setFormErrors,
    handleChange,
    defaultForm,
  }
}
