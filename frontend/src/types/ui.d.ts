import type { AxiosError } from 'axios'
import type { ApiError } from 'src/api/mutator'
import type { ValidationErrors } from 'src/client-validator/validate'

type AppError = ApiError | AxiosError // TODO: react hook form errors instead of validationerrors
