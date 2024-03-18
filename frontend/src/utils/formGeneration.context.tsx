import { useContext } from 'react'
import { AllKeysMandatory } from 'src/types/utils'
import {
  DynamicFormContext,
  DynamicFormContextValue,
  FieldOptions as FieldOptionsBuilder,
  InputOptions as InputBuilder,
  SelectOptions as SelectOptionsBuilder,
} from 'src/utils/formGeneration'

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
}: SelectOptionsBuilder<ReturnElement, V>): AllKeysMandatory<SelectOptionsBuilder<ReturnElement, V>> => ({
  type,
  values,
  optionTransformer,
  formValueTransformer,
  // workaround to never forget adding new fields
  pillTransformer: pillTransformer!,
  ariaLabelTransformer: ariaLabelTransformer!,
  searchValueTransformer: searchValueTransformer!,
  labelColor: labelColor!,
})

export const fieldOptionsBuilder = <Return, V, ReturnElement = Return extends unknown[] ? Return[number] : Return>({
  warningFn,
}: FieldOptionsBuilder<ReturnElement, V>): AllKeysMandatory<FieldOptionsBuilder<ReturnElement, V>> => ({
  // workaround to never forget adding new fields
  warningFn: warningFn!,
})

export const inputBuilder = <Return, V>({ component }: InputBuilder<Return, V>): InputBuilder<Return, V> => ({
  component,
})
