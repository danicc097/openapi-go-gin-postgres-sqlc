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
} from '@mantine/core'
import { DateInput, DateTimePicker } from '@mantine/dates'
import { Form, type UseFormReturnType } from '@mantine/form'
import type { UseForm } from '@mantine/form/lib/types'
import { Prism } from '@mantine/prism'
import { useMantineTheme } from '@mantine/styles'
import { IconMinus, IconPlus } from '@tabler/icons'
import _ from 'lodash'
import { useState, type ComponentProps } from 'react'
import type { FieldPath } from 'react-hook-form'
import PageTemplate from 'src/components/PageTemplate'
import type { RestDemoWorkItemCreateRequest } from 'src/gen/model'
import type { GenericObject, RecursiveKeyOf, RecursiveKeyOfArray, TypeOf } from 'src/types/utils'
import type { SchemaField } from 'src/utils/jsonSchema'
import { entries } from 'src/utils/object'

type options<T extends string, U extends GenericObject> = {
  defaultValue: Partial<{
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    //@ts-ignore
    [key in T]: TypeOf<U, key>
  }>
}

type DynamicFormProps<T extends string, U extends GenericObject> = {
  form: UseFormReturnType<U, (values: U) => U>
  schemaFields: Record<T, SchemaField>
  options: options<T, U>
}

type GenerateComponentProps<T> = {
  form: UseFormReturnType<T>
  fieldType: SchemaField['type']
  props: any
  formField: string
}

// IMPORTANT: field dot notation requires indexes for arrays. e.g. `members.0.role`.
function generateComponent<T>({ form, fieldType, props, formField }: GenerateComponentProps<T>) {
  // TODO: multiselect and select early check (if found in options.components override)
  const _props = {
    mb: 4,
    ...form.getInputProps(formField),
    ...props,
  }

  // TODO: helpText is `description` prop in mantine

  switch (fieldType) {
    case 'string':
      return <TextInput {..._props} />
    case 'boolean':
      return <Checkbox {..._props} />
    case 'date':
      return <DateInput placeholder="Select date" {..._props} />
    case 'date-time':
      return <DateTimePicker placeholder="Select date and time" {..._props} />
    case 'integer':
      return <NumberInput {..._props} />
    default:
      return null
  }
}

function renderTitle(key: string) {
  return (
    <>
      <Title size={18}>{key}</Title>
      <Space p={8} />
    </>
  )
}

type GenerateFormInputsProps = {
  parentFieldKey?: string
  index?: number
  parentFormField?: string
}

export const DynamicForm = <T extends string, U extends GenericObject>({
  form,
  schemaFields,
  options,
}: DynamicFormProps<T, U>) => {
  const theme = useMantineTheme()

  function initialValueByField(field: T) {
    switch (schemaFields[field].type) {
      case 'object':
        return {}
      case 'array':
        return []
      default:
        return undefined
    }
  }

  const addNestedField = (field: T, formField: string) => {
    const initialValue = initialValueByField(field)
    console.log({ addNestedFieldField: field, addNestedFieldFormField: formField, initialValue })

    const newValues = _.cloneDeep(form.values)

    _.set(newValues, formField, [...(_.get(newValues, field, []) || []), initialValue])

    form.setValues((currentValues) => newValues)
  }

  function renderRemoveNestedFieldButton(formField: string, index: number) {
    return (
      <ActionIcon
        onClick={(e) => {
          console.log({ removeNestedField: `${formField}[${index}]` })
          form.removeListItem(formField, index)
        }}
        variant="filled"
        color={'red'}
      >
        <IconMinus size="1rem" />
      </ActionIcon>
    )
  }

  const generateFormInputs = ({ parentFieldKey = '', parentFormField = '' }: GenerateFormInputsProps) => {
    return entries(schemaFields).map(([fieldKey, field]) => {
      function renderNestedHeader() {
        return (
          <div>
            {<Prism language="json">{JSON.stringify({ formField, parentFormField }, null, 4)}</Prism>}
            <Flex direction="row">
              {renderTitle(formField)}
              <Button
                size="xs"
                p={4}
                leftIcon={<IconPlus size="1rem" />}
                onClick={() => addNestedField(fieldKey, formField)}
                variant="filled"
                color={'green'}
              >{`Add ${formField}`}</Button>
            </Flex>
          </div>
        )
      }

      if (parentFieldKey !== '' && !fieldKey.startsWith(parentFieldKey)) {
        return null
      }

      const pp = fieldKey.split('.')
      const parentKey = parentFieldKey.replace(/\.*$/, '') || pp.slice(0, pp.length - 1).join('.')

      if (schemaFields[parentKey]?.isArray && parentFieldKey === '') return null

      const formField = constructFormKey(fieldKey, parentFormField)
      if (parentFormField !== '') {
        console.log({ parentFormField })
      }
      console.log({ formValue: _.get(form.values, formField), formField })

      const componentProps = {
        css: css`
          min-width: 100%;
        `,
        label: formField,
        required: field.required,
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
                <div key={_index}>
                  {generateComponent({
                    form,
                    fieldType: field.type,
                    formField: `${formField}.${_index}`,
                    props: componentProps,
                  })}
                  {renderRemoveNestedFieldButton(formField, _index)}
                </div>
              )
            })}
          </Card>
        )
      }

      // fix deeply nested
      // FIXME: need to check if there are isArray in existing field, e.g. base.title.items -->
      // need to check if base or base.title is already arrayofobject or array! and we need to pass what we will call parentFormField option, e.g. "base.title.2.items",
      // parentFormField keeps accumulating form field access with index when we do `generateFormInputs`:
      // generateFormInputs({ parentFieldKey: fieldKey, index: _index, parentFormField: `${formField}.${index}` }
      // apart from just "base.title.items" to construct index access on deeply nested generation doing some string wrangling
      // (same reasoning as constructFormKey)

      if (field.isArray && field.type === 'object') {
        console.log({ nestedArrayOfObjects: formField })

        // array of objects
        return (
          <Card key={fieldKey} mt={24}>
            {parentFieldKey === '' && <>{renderNestedHeader()}</>}
            {/* FIXME: bad gen array is nested - removenested and form inputs wrong. (base.metadata vs tagIDs working fine) */}
            {_.get(form.values, formField)?.map((_nestedValue: any, _index: number) => {
              console.log({ nestedArrayOfObjectsIndex: _index })
              return (
                <div key={_index} style={{ marginBottom: theme.spacing.sm }}>
                  <p>{`${fieldKey}[${_index}]`}</p>
                  <Group>
                    {generateFormInputs({
                      parentFieldKey: fieldKey,
                      parentFormField: `${formField}.${_index}`,
                    })}
                  </Group>
                  {renderRemoveNestedFieldButton(formField, _index)}
                </div>
              )
            })}
          </Card>
        )
      }

      return (
        <Group key={fieldKey} align="center">
          {field.type !== 'object' ? (
            <>{generateComponent({ form, fieldType: field.type, props: componentProps, formField: formField })}</>
          ) : (
            <>{renderTitle(formField)}</>
          )}
        </Group>
      )
    })
  }

  return (
    <PageTemplate minWidth={800}>
      <form
        css={css`
          min-width: 100%;
        `}
      >
        {generateFormInputs({})}
      </form>
    </PageTemplate>
  )
}

/**
 * Construct form accessor based on current schema field key and parent form field.
 */
function constructFormKey(key: string, parentFormField: string) {
  const currentFieldName = key.split('.').slice(-1)[0]

  return parentFormField !== '' ? `${parentFormField}.${currentFieldName}` : key
}
