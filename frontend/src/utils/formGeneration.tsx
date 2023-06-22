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

export const DynamicForm = <T extends string, U extends GenericObject>({
  form,
  schemaFields,
  options,
}: DynamicFormProps<T, U>) => {
  const theme = useMantineTheme()

  const addNestedField = (field: string) => {
    console.log({ addNestedField: field })
    form.setValues((currentValues) => ({
      ...currentValues,
      [field]: [
        ...(currentValues[field] || []), // can't use insertListItem directly if not initialized so just do it at once
        // TODO: maybe will need initialValue based on type from schema
        {}, //  mantine form  does not support array of nonobjects validation
      ],
    }))
  }

  const removeNestedField = (field: string, index: number) => {
    console.log({ removeNestedField: `${field}[${index}]` })

    form.removeListItem(field, index)
  }

  const generateFormFields = ({ parentPathPrefix = '', index }: { parentPathPrefix?: string; index?: number }) => {
    return entries(schemaFields).map(([key, field]) => {
      if (parentPathPrefix !== '' && !key.startsWith(parentPathPrefix)) {
        return
      }
      const pp = key.split('.')
      const parentKey = parentPathPrefix.replace(/\.*$/, '') || pp.slice(0, pp.length - 1).join('.')

      const fieldKey = String(key)
      const formKey = constructFormKey(fieldKey, index)
      // console.log({ formValue: _.get(form.values, formKey), formKey })

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
            {JSON.stringify(_.get(form.values, formKey))}
            <div>
              <ActionIcon onClick={() => addNestedField(fieldKey)} variant="filled" color={'green'}>
                <IconPlus size="1rem" />
              </ActionIcon>
              <Title size={18}>{key}</Title>
            </div>
            {/* existing array fields, if any */}
            {_.get(form.values, formKey)?.map((_nestedValue: any, index: number) => {
              return (
                <div key={index}>
                  {JSON.stringify({ [fieldKey]: index })}
                  {generateComponent({
                    form,
                    fieldType: field.type,
                    formField: formKey,
                    props: componentProps,
                  })}
                  <ActionIcon onClick={() => removeNestedField(fieldKey, index)} variant="filled" color={'red'}>
                    <IconMinus size="1rem" />
                  </ActionIcon>
                </div>
              )
            })}
          </Card>
        )
      }

      if (field.isArray && field.type === 'object') {
        // FIXME: should not generate nested, same as with members.*
        // array of primitives
        console.log({ parentKey, parentPathPrefix })
        if (schemaFields[parentKey]?.isArray) {
          return null
        }
        // array of objects
        return (
          <Card key={fieldKey} mt={24}>
            <ActionIcon onClick={() => addNestedField(fieldKey)} variant="filled" color={'green'}>
              <IconPlus size="1rem" />
            </ActionIcon>
            {_.get(form.values, formKey)?.map((_nestedValue: any, index: number) => {
              return (
                <div key={index} style={{ marginBottom: theme.spacing.sm }}>
                  <p>{`${fieldKey}[${index}]`}</p>
                  <Group>{generateFormFields({ parentPathPrefix: fieldKey, index })}</Group>
                  <ActionIcon onClick={() => removeNestedField(fieldKey, index)} variant="filled" color={'red'}>
                    <IconMinus size="1rem" />
                  </ActionIcon>
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
            <>{generateComponent({ form, fieldType: field.type, props: componentProps, formField: formKey })}</>
          ) : (
            <>
              <Title size={18}>{key}</Title>
              <Space p={8} />
            </>
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
        {generateFormFields({})}
      </form>
    </PageTemplate>
  )
}

/**
 * Construct form accessor based on dot notation path and index (in case of array element).
 */
function constructFormKey(key: string, index?: number) {
  const formPaths = key.split('.')
  const formKey =
    index !== undefined
      ? [...formPaths.slice(0, formPaths.length - 1), index, formPaths[formPaths.length - 1]].join('.')
      : String(key)
  return formKey
}
