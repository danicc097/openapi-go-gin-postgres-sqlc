import { EuiFormRow, EuiSuperSelect, EuiSuperSelectOption } from '@elastic/eui'
import type { UseFormReturnType } from '@mantine/form'
import { capitalize } from 'lodash'

export function createLabel(name: string, isRequired: boolean): React.ReactNode {
  return (
    <>
      {capitalize(name)}
      {isRequired && <b css={{ color: 'red' }}> *</b>}
    </>
  )
}

type RenderSuperSelectParams<U extends keyof Form, Form> = {
  formKey: U
  form: UseFormReturnType<Form, (values: Form) => Form>
  options: EuiSuperSelectOption<Form[U]>[]
  requiredFormKeys: Partial<Form>
  onSuperSelectChange: (value: any) => void
}

export function renderSuperSelect<Form, U extends keyof Form>({
  formKey,
  form,
  options,
  requiredFormKeys,
  onSuperSelectChange,
}: RenderSuperSelectParams<U, Form>) {
  return (
    <EuiFormRow
      label={createLabel(String(formKey), formKey in requiredFormKeys)}
      helpText={`Select a ${String(formKey)}.`}
      isInvalid={Boolean(form.getInputProps(formKey).error)}
      {...form.getInputProps(formKey)}
      error={capitalize(form.getInputProps(formKey).error)}
      fullWidth
    >
      <EuiSuperSelect
        name={String(formKey)}
        options={options}
        valueOfSelected={form.values[formKey]}
        onChange={onSuperSelectChange}
        itemLayoutAlign="top"
        hasDividers
        fullWidth
        isInvalid={Boolean(form.getInputProps(formKey).error)}
      />
    </EuiFormRow>
  )
}
