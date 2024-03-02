import { useContext } from 'react'
import { DynamicFormContext, DynamicFormContextValue, InputOptions, SelectOptions } from 'src/utils/formGeneration'

export const useDynamicFormContext = (): DynamicFormContextValue => {
  const context = useContext(DynamicFormContext)

  if (!context) {
    throw new Error('useDynamicFormContext must be used within a DynamicFormProvider')
  }

  return context
}

// NOTE: handles select (single return value) and multiselect (array return).
export const selectOptionsBuilder = <Return, V, ReturnElement = Return extends unknown[] ? Return[number] : Return>({
  type,
  values,
  formValueTransformer,
  searchValueTransformer,
  optionTransformer,
  pillTransformer,
  labelColor,
}: SelectOptions<ReturnElement, V>): SelectOptions<ReturnElement, V> => ({
  type,
  values,
  optionTransformer,
  pillTransformer,
  formValueTransformer,
  searchValueTransformer,
  labelColor,
})

export const inputBuilder = <Return, V>({ component }: InputOptions<Return, V>): InputOptions<Return, V> => ({
  component,
})
