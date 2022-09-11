import type { Decoder } from 'src/client-validator/gen/helpers'
import type { ValidationErrors } from 'src/client-validator/validate'

export const validateField = (decoder: Decoder<any>, key: string, values: unknown): string => {
  try {
    decoder.decode(values)
    return null
  } catch (error) {
    const vErrors: ValidationErrors = error.validationErrors
    let errMsg = null
    vErrors?.errors?.forEach((v) => {
      if (v.invalidParams.name === key) {
        errMsg = ' '
      }
    })

    return errMsg
  }
}
