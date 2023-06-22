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

  const addNestedField = (field: string, initialValue: any) => {
    console.log({ addNestedField: field })
    form.setValues((currentValues) => ({
      ...currentValues,
      [field]: [
        ...(currentValues[field] || []), // can't use insertListItem directly if not initialized so just do it at once
        // TODO: maybe will need initialValue based on type from schema
        initialValue, //  mantine form  does not support array of nonobjects validation
      ],
    }))
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
    return entries(schemaFields).map(([key, field]) => {
      if (parentPathPrefix !== '' && !key.startsWith(parentPathPrefix)) {
        return
      }
      const pp = key.split('.')
      const parentKey = parentPathPrefix.replace(/\.*$/, '') || pp.slice(0, pp.length - 1).join('.')

      const fieldKey = String(key)
      const formField = constructFormKey(fieldKey, index)
      // console.log({ formValue: _.get(form.values, formField), formField })

      const componentProps = {
        css: css`
          min-width: 100%;
        `,
        label: key,
        required: field.required,
      }

      if (field.isArray && field.type !== 'object') {
        // nested array objects generation
        return (
          <Card key={fieldKey} mt={24}>
            {JSON.stringify(_.get(form.values, formField))}
            <div>
              {renderTitle(key)}
              <ActionIcon onClick={() => addNestedField(fieldKey, null)} variant="filled" color={'green'}>
                <IconPlus size="1rem" />
              </ActionIcon>
            </div>
            {/* existing array fields, if any */}
            {console.log({ nestedArray: formField })}
            {_.get(form.values, formField)?.map((_nestedValue: any, index: number) => {
              return (
                <div key={index}>
                  {JSON.stringify({ [fieldKey]: index })}
                  {generateComponent({
                    form,
                    fieldType: field.type,
                    formField: formField,
                    props: componentProps,
                  })}
                  {renderRemoveNestedFieldButton(formField, index)}
                </div>
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
            {renderTitle(key)}
            <ActionIcon onClick={() => addNestedField(fieldKey, {})} variant="filled" color={'green'}>
              <IconPlus size="1rem" />
            </ActionIcon>
            {_.get(form.values, formField)?.map((_nestedValue: any, index: number) => {
              return (
                <div key={index} style={{ marginBottom: theme.spacing.sm }}>
                  <p>{`${fieldKey}[${index}]`}</p>
                  <Group>{generateFormInputs({ parentPathPrefix: fieldKey, index })}</Group>
                  {renderRemoveNestedFieldButton(formField, index)}
                </div>
              )
            })}
          </Card>
        )
      }

      if (schemaFields[parentKey]?.isArray && parentPathPrefix === '') return null

      return (
        <Group key={fieldKey} align="center">
          {field.type !== 'object' ? (
            <>{generateComponent({ form, fieldType: field.type, props: componentProps, formField: formField })}</>
          ) : (
            <>{renderTitle(key)}</>
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
