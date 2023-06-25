import { css } from '@emotion/react'
import {
  Group,
  TextInput,
  NumberInput,
  Checkbox,
  Button,
  Title,
  Space,
  Divider,
  Text,
  type InputProps,
  ActionIcon,
  Card,
  Container,
  Box,
  Flex,
  Tooltip,
} from '@mantine/core'
import { DateInput, DateTimePicker } from '@mantine/dates'
import { Form, type UseFormReturnType } from '@mantine/form'
import type { UseForm } from '@mantine/form/lib/types'
import { Prism } from '@mantine/prism'
import { useMantineTheme } from '@mantine/styles'
import { Icon123, IconMinus, IconPlus } from '@tabler/icons'
import _ from 'lodash'
import { useState, type ComponentProps } from 'react'
import PageTemplate from 'src/components/PageTemplate'
import type { RestDemoWorkItemCreateRequest } from 'src/gen/model'
import type {
  DeepPartial,
  GenericObject,
  GetKeys,
  RecursiveKeyOf,
  RecursiveKeyOfArray,
  PathType,
} from 'src/types/utils'
import type { SchemaField } from 'src/utils/jsonSchema'
import { entries } from 'src/utils/object'

export type SelectOptionsTypes = 'select' | 'multiselect'

export interface SelectOptions<Return, E = unknown> {
  values: E[]
  type: SelectOptionsTypes
  formValueTransformer: <V extends E>(el: V & E) => Return
  // TODO: via mantine componentValue
  componentTransformer: <V extends E>(el: V & E) => JSX.Element
}

export const selectOptionsBuilder = <Return, V>({
  type,
  values,
  formValueTransformer,
  componentTransformer,
}: SelectOptions<Return, V>): SelectOptions<Return, V> => ({
  type,
  values,
  componentTransformer,
  formValueTransformer,
})

type options<T extends object, U extends string = GetKeys<T>> = {
  // used to populate form inputs if the form field is empty. Applies to all nested fields.
  defaultValues?: Partial<{
    [key in U]: DeepPartial<PathType<T, key>>
  }>
  //  list of options used for Select and MultiSelect
  // TODO: someone had the exact same idea: https://stackoverflow.com/questions/69254779/infer-type-based-on-the-generic-type-of-a-sibling-property-in-typescript
  // more recent version: https://stackoverflow.com/questions/74618270/how-to-make-an-object-property-depend-on-another-one-in-a-generic-type
  // TODO: inputComponent field, e.g. for color picker. if inputComponent === undefined, then switch on schema format as usual
  selectOptions?: Partial<{
    [key in U]: ReturnType<typeof selectOptionsBuilder<PathType<T, key>, unknown>>
  }>
}

type DynamicFormProps<T extends object, U extends string = GetKeys<T>> = {
  form: UseFormReturnType<T, (values: T) => T>
  schemaFields: Record<U, SchemaField>
  options: options<T, U>
  name: string
}

type GenerateComponentProps<U> = {
  form: UseFormReturnType<U>
  fieldType: SchemaField['type']
  props?: {
    input?: any
    container?: any
  }
  formField: string
  removeButton?: JSX.Element
}

function generateComponent<U>({ form, fieldType, props, formField, removeButton }: GenerateComponentProps<U>) {
  // TODO: multiselect and select early check (if found in options.components override)
  const _props = {
    mb: 4,
    ...form.getInputProps(formField),
    ...props?.input,
    ...(removeButton && { rightSection: removeButton, rightSectionWidth: '40px' }),
  }

  // TODO: helpText is `description` prop in mantine.
  // will accecss these via options[field (not formField since its shared for array elements)].<description|label|formValueTransformer|...>

  let el = null
  switch (fieldType) {
    case 'string':
      el = <TextInput {..._props} />
      break
    case 'boolean':
      el = <Checkbox {..._props} />
      break
    case 'date':
      el = <DateInput placeholder="Select date" {..._props} />
      break
    case 'date-time':
      el = <DateTimePicker placeholder="Select date and time" {..._props} />
      break
    case 'integer':
      el = <NumberInput {..._props} />
      break
    default:
      break
  }

  return (
    <Flex align="center" {...props?.container}>
      {el}
    </Flex>
  )
}

function renderTitle(key: string) {
  return (
    <>
      <Title data-testid={`${key}-title`} size={18}>
        {key}
      </Title>
      <Space p={8} />
    </>
  )
}

type GenerateFormInputsProps = {
  parentFieldKey?: string
  index?: number
  parentFormField?: string
  removeButton?: JSX.Element
}

export default function DynamicForm<T extends object, U extends string = GetKeys<T>>({
  name,
  form,
  schemaFields,
  options,
}: DynamicFormProps<T, U>) {
  const theme = useMantineTheme()

  function initialValueByField(field: U) {
    switch (schemaFields[field].type) {
      case 'object':
        return {}
      case 'array':
        return []
      default:
        return undefined
    }
  }

  const addNestedField = (field: U, formField: string) => {
    const initialValue = initialValueByField(field)
    console.log({ addNestedFieldField: field, addNestedFieldFormField: formField, initialValue })

    const newValues = _.cloneDeep(form.values)

    _.set(newValues, formField, [...(_.get(newValues, formField) || []), initialValue])

    form.setValues((currentValues) => newValues)
  }

  function renderRemoveNestedFieldButton(formField: string, index: number) {
    return (
      <Tooltip label="Remove item" position="top-end" withArrow>
        <ActionIcon
          onClick={(e) => {
            console.log({ removeNestedField: `${formField}[${index}]` })
            form.removeListItem(formField, index)
          }}
          // variant="filled"
          css={css`
            background-color: #7c1a1a;
          `}
          size="sm"
          id={`${name}-${formField}-remove-button-${index}`}
        >
          <IconMinus size="1rem" />
        </ActionIcon>
      </Tooltip>
    )
  }

  const generateFormInputs = ({
    parentFieldKey = '',
    parentFormField = '',
    removeButton = null,
  }: GenerateFormInputsProps) => {
    return entries(schemaFields).map(([fieldKey, field]) => {
      function renderNestedHeader() {
        return (
          <div>
            {/* {<Prism language="json">{JSON.stringify({ formField, parentFormField }, null, 4)}</Prism>} */}
            <Flex direction="row">
              {renderTitle(formField)}
              <Button
                size="xs"
                p={4}
                leftIcon={<IconPlus size="1rem" />}
                onClick={() => addNestedField(fieldKey, formField)}
                variant="filled"
                color={'green'}
                id={`${name}-${formField}-add-button`}
              >{`Add ${formField}`}</Button>
            </Flex>
          </div>
        )
      }

      if (
        (parentFieldKey !== '' && !fieldKey.startsWith(parentFieldKey)) ||
        parentFieldKey === fieldKey // fix when parent key has the same name and both are arrays
      ) {
        return null
      }

      const pp = fieldKey.split('.')
      const parentKey = parentFieldKey.replace(/\.*$/, '') || pp.slice(0, pp.length - 1).join('.')

      if (schemaFields[parentKey]?.isArray && parentFieldKey === '') return null

      const formField = constructFormKey(fieldKey, parentFormField)
      if (parentFormField !== '') {
        console.log({ parentFormField })
      }
      // console.log({ formValue: _.get(form.values, formField), formField })

      const containerProps = {
        css: css`
          width: 100%;
        `,
      }

      const inputProps = {
        css: css`
          width: 100%;
        `,
        ...(!field.isArray && { label: formField }),
        required: field.required,
        id: `${name}-${formField}`,
      }

      if (field.isArray && field.type !== 'object') {
        // nested array of nonbjects generation
        return (
          <Card key={fieldKey} mt={24}>
            {renderNestedHeader()}
            {/* existing array fields, if any */}
            {_.get(form.values, formField)?.map((_nestedValue: any, _index: number) => {
              console.log({ _nestedValue, _index })
              return (
                <Flex key={_index}>
                  {generateComponent({
                    form,
                    fieldType: field.type,
                    formField: `${formField}.${_index}`,
                    props: {
                      input: { ...inputProps, id: `${name}-${formField}-${_index}` },
                      container: containerProps,
                    },
                    removeButton: renderRemoveNestedFieldButton(formField, _index),
                  })}
                </Flex>
              )
            })}
          </Card>
        )
      }

      if (field.isArray && field.type === 'object') {
        console.log({ nestedArrayOfObjects: formField })

        // array of objects
        return (
          <Card key={fieldKey} mt={24}>
            {parentFieldKey === '' && <>{renderNestedHeader()}</>}
            {_.get(form.values, formField)?.map((_nestedValue: any, _index: number) => {
              console.log({ nestedArrayOfObjectsIndex: _index })
              return (
                <div key={_index}>
                  <p>{`${fieldKey}[${_index}]`}</p>
                  {renderRemoveNestedFieldButton(formField, _index)}
                  <Group>
                    {generateFormInputs({
                      parentFieldKey: fieldKey,
                      parentFormField: `${formField}.${_index}`,
                      removeButton: null,
                    })}
                  </Group>
                </div>
              )
            })}
          </Card>
        )
      }

      return (
        <Group key={fieldKey} align="center">
          {field.type !== 'object' ? (
            <>
              {removeButton}
              {generateComponent({
                form,
                fieldType: field.type,
                props: { input: inputProps, container: containerProps },
                formField: formField,
                removeButton: null,
              })}
            </>
          ) : (
            <>{renderTitle(formField)}</>
          )}
        </Group>
      )
    })
  }

  // TODO: will also need sorting schemaFields beforehand and then generate normally.
  return (
    <PageTemplate minWidth={1000}>
      <form
        css={css`
          min-width: 100%;
        `}
        id={name}
      >
        {generateFormInputs({})}
      </form>
    </PageTemplate>
  )
}

/**
 * Construct form accessor based on current schema field key and parent form field.
 */
export function constructFormKey(fieldKey: string, parentFormField: string) {
  const currentFieldName = fieldKey.split('.').slice(-1)[0]

  return parentFormField !== '' ? `${parentFormField}.${currentFieldName}` : fieldKey
}
