type FormErrors<Form> = Partial<
  {
    [key in keyof Form]: string | boolean
  } & {
    form: string
  }
>

type AppError = ValidationErrors | ApiError | AxiosError
