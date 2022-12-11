type FormErrors<Form> = Partial<
  {
    [key in keyof Form]: string | boolean
  } & {
    form: string
  }
>
