import { useContext } from 'react'
import { AllKeysMandatory } from 'src/types/utils'
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
  ariaLabelTransformer,
  searchValueTransformer,
  optionTransformer,
  pillTransformer,
  labelColor,
}: SelectOptions<ReturnElement, V>): AllKeysMandatory<SelectOptions<ReturnElement, V>> => ({
  type,
  values,
  optionTransformer,
  formValueTransformer,
  // workaround to never forget adding new fields to selectOptionsBuilder
  pillTransformer: pillTransformer!,
  ariaLabelTransformer: ariaLabelTransformer!,
  searchValueTransformer: searchValueTransformer!,
  labelColor: labelColor!,
})

export const inputBuilder = <Return, V>({ component }: InputOptions<Return, V>): InputOptions<Return, V> => ({
  component,
})
