import type { AxiosError } from 'axios'
import type { ApiError } from 'src/api/mutator'
import type { ValidationErrors } from 'src/client-validator/validate'

type FormErrors<Form> = Partial<
  {
    [key in keyof Form]: string | boolean
  } & {
    form: string
  }
>

type AppError = ValidationErrors | ApiError | AxiosError
