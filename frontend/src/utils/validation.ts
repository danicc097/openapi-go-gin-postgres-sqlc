import type { Decoder } from 'src/client-validator/gen/helpers'
import type { ValidationErrors } from 'src/client-validator/validate'

export const validateField = (decoder: Decoder<any>, key: string, values: unknown): string => {
  try {
    decoder.decode(values)
    return null
  } catch (error) {
    const vErrors: ValidationErrors = error.validationErrors
    let errMsg = null
    // with elastic ui instead of validateField we should instead
    // generate a formErrors object so that we can have the reason null | string indexed by key name:
    // <EuiFormRow label="Title" isInvalid={Boolean(formErrors.title)} error={formErrors.title}>
    //    <EuiFieldText isInvalid={Boolean(formErrors.title)} ... /> // simply colours it
    // </EuiFormRow>
    vErrors?.errors?.forEach((v) => {
      if (v.invalidParams.name === key) {
        // TODO meaningful error per form row instead of a callout
        console.log(v.invalidParams.reason)
        errMsg = ' ' // trigger error with empty message for the key being validated
      }
    })

    return errMsg
  }
}
