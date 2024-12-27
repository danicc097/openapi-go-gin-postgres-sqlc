import type { HTTPValidationError, ValidationError } from 'src/gen/model'

export const isSerializable = (obj: any) => {
  try {
    JSON.stringify(obj)
    return true
  } catch (e) {
    return false
  }
}

export const errorFieldToMessageMapping: {
  [key: string]: string
} = {
  email: 'email',
  username: 'username',
  password: 'password',
  sender_: 'sender',
  receiver_role: 'receiver role',
  old_password: 'old password',
  title: 'title',
  body: 'body',
  label: 'label',
  link: 'link',
}

export const parseErrorDetail = (errorDetail: ValidationError): string => {
  let errorMessage = ''

  if (Array.isArray(errorDetail?.loc)) {
    if (errorDetail.loc[0] === 'path') return errorMessage

    if (errorDetail.loc[0] === 'query') return `Invalid ${errorDetail.loc[1]}: ${errorDetail.msg}`

    if (errorDetail.loc[0] === 'body') {
      const invalidField = errorDetail.loc[2] || errorDetail.loc[1]
      if (!invalidField) return errorMessage
      errorMessage = `Invalid ${errorFieldToMessageMapping[invalidField] ?? invalidField}: ${errorDetail?.msg}`
    }
  } else {
    errorMessage = 'Something unknown went wrong. Contact support.\n'
  }

  return errorMessage
}

export const extractErrorMessages = (error: HTTPValidationError): unknown[] => {
  const errorList: unknown[] = []

  if (typeof error === 'string') errorList.push(error)

  if (typeof error?.detail === 'string') {
    errorList.push(error.detail === 'Not Found' ? 'Internal Server Error' : error.detail)
  }

  if (Array.isArray(error?.detail)) {
    error.detail.forEach((errorDetail) => {
      const errorMessage = parseErrorDetail(errorDetail)
      errorList.push(errorMessage)
    })
  }
  return errorList
}
