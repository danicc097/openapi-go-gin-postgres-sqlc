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
} from '@mantine/core'
import { DateInput } from '@mantine/dates'
import { Form, type UseFormReturnType } from '@mantine/form'
import type { UseForm } from '@mantine/form/lib/types'
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

  console.log(formField)

  // TODO: helpText is `description` prop in mantine

  switch (fieldType) {
    case 'string':
      return <TextInput {..._props} />
    case 'boolean':
      return <Checkbox {..._props} />
    case 'date-time':
      return <DateInput placeholder="Select date" {..._props} />
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
      default:
        return undefined
    }
  }

  const addNestedField = (field: T) => {
    const initialValue = initialValueByField(field)
    console.log({ addNestedField: field, initialValue })

    const newValues = _.cloneDeep(form.values)

    _.set(newValues, field, [...(_.get(newValues, field, []) || []), initialValue])

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

  const generateFormInputs = ({ parentPathPrefix = '', index }: { parentPathPrefix?: string; index?: number }) => {
    return entries(schemaFields).map(([fieldKey, field]) => {
      if (parentPathPrefix !== '' && !fieldKey.startsWith(parentPathPrefix)) {
        return null
      }

      const pp = fieldKey.split('.')
      const parentKey = parentPathPrefix.replace(/\.*$/, '') || pp.slice(0, pp.length - 1).join('.')

      if (schemaFields[parentKey]?.isArray && parentPathPrefix === '') return null

      const formField = constructFormKey(fieldKey, index)
      // console.log({ formValue: _.get(form.values, formField), formField })

      const componentProps = {
        css: css`
          min-width: 100%;
        `,
        label: fieldKey,
        required: field.required,
      }

      if (field.isArray && field.type !== 'object') {
        // nested array of nonbjects generation
        return (
          <Card key={fieldKey} mt={24}>
            {JSON.stringify(_.get(form.values, formField))}
            <div>
              {renderTitle(fieldKey)}
              <ActionIcon onClick={() => addNestedField(fieldKey)} variant="filled" color={'green'}>
                <IconPlus size="1rem" />
              </ActionIcon>
            </div>
            {/* existing array fields, if any */}
            {console.log({ nestedArray: formField, formValue: _.get(form.values, formField) })}
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

      if (field.isArray && field.type === 'object') {
        console.log({ nestedArrayOfObjects: formField })

        /**
         *
         * FIXME: is not using index to access form (cause of toFixed not a function)
         */

        // array of objects
        return (
          <Card key={fieldKey} mt={24}>
            {parentPathPrefix === '' && (
              <>
                {renderTitle(fieldKey)}
                <ActionIcon onClick={() => addNestedField(fieldKey)} variant="filled" color={'green'}>
                  <IconPlus size="1rem" />
                </ActionIcon>
              </>
            )}
            {/* FIXME: bad gen array is nested - removenested and form inputs wrong. (base.metadata vs tagIDs working fine) */}
            {_.get(form.values, formField)?.map((_nestedValue: any, _index: number) => {
              console.log({ nestedArrayOfObjectsIndex: _index })
              return (
                <div key={_index} style={{ marginBottom: theme.spacing.sm }}>
                  <p>{`${fieldKey}[${_index}]`}</p>
                  <Group>{generateFormInputs({ parentPathPrefix: fieldKey, index: _index })}</Group>
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
            <>{renderTitle(fieldKey)}</>
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
 * Construct form accessor based on dot notation path and index (in case of array element).
 */
function constructFormKey(key: string, index?: number) {
  const formPaths = key.split('.')
  const formField =
    index !== undefined
      ? [...formPaths.slice(0, formPaths.length - 1), index, formPaths[formPaths.length - 1]].join('.')
      : String(key)
  return formField
}
