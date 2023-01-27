import { createFormContext } from '@mantine/form'
import type { UpdateUserAuthRequest } from 'src/gen/model'

// TODO
export const [UpdateUserAuthFormProvider, useUpdateUserAuthFormContext, useUpdateUserAuthForm] =
  createFormContext<UpdateUserAuthRequest>()
