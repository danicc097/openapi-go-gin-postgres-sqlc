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
  field: string
}

// IMPORTANT: field dot notation requires indexes for arrays. e.g. `members.0.role`.
function generateComponent<T>({ form, fieldType, props, field }: GenerateComponentProps<T>) {
  // TODO: multiselect and select early check (if found in options.components override)
  const _props = {
    ...form.getInputProps(field),
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
        // null,
        // 0,
        {}, // mantine form  does not support array of nonobjects validation
        // { role: 'preparer', userID: 'rsfsese' }, // should have initial object generated based on path if type === object, else it will attempt setting on null
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
      const parentKey = parentPathPrefix.replace(/\.*$/, '')

      let indexSuffix = ''
      if (index !== undefined) {
        indexSuffix = `.${index}`
      }
      const fieldKey = String(key)
      const formKey = String(key) + indexSuffix // TODO: pass to everything that uses `form` instead
      console.log({ formValue: _.get(form.values, formKey), formKey })

      const componentProps = {
        css: css`
          min-width: 100%;
        `,
        label: key,
        required: field.required,
      }

      if (field.isArray && field.type !== 'object') {
        // FIXME: should not generate, same as with members.*
        // array of primitives

        // TODO: form.getInputProps instead.
        // form.getInputProps('base.<nested>', {type: "checkbox | input"})

        return (
          <Card key={fieldKey} mt={24}>
            {JSON.stringify(form.values[fieldKey])}
            <div>
              <ActionIcon onClick={() => addNestedField(fieldKey)} variant="filled" color={'green'}>
                <IconPlus size="1rem" />
              </ActionIcon>
              <Title size={18}>{key}</Title>
            </div>
            {/* existing array fields, if any */}
            {form.values[fieldKey]?.map((_nestedValue: any, index: number) => {
              return (
                <div key={index} style={{ display: 'flex', marginBottom: theme.spacing.xs }}>
                  {JSON.stringify({ [fieldKey]: index })}
                  {generateComponent({
                    form,
                    fieldType: field.type,
                    field: fieldKey,
                    props: componentProps,
                  })}
                  <ActionIcon onClick={() => removeNestedField(fieldKey, index)} variant="filled" color={'green'}>
                    <IconMinus size="1rem" />
                  </ActionIcon>
                </div>
              )
            })}
          </Card>
        )
      }

      if (field.isArray && field.type === 'object') {
        if (schemaFields[parentKey]?.isArray) {
          return null
        }
        // array of objects
        return (
          <Card key={fieldKey} mt={24}>
            {JSON.stringify({ [fieldKey]: form.values[fieldKey] })}
            <ActionIcon onClick={() => addNestedField(fieldKey)} variant="filled" color={'green'}>
              <IconPlus size="1rem" />
            </ActionIcon>

            {renderNestedArrayOfObjects()}
          </Card>
        )
      }

      const paths = key.split('.')
      const parent = paths.slice(0, paths.length - 1).join('.')

      return (
        <Group key={fieldKey} align="center">
          {field.type !== 'object' ? (
            <>{generateComponent({ form, fieldType: field.type, props: componentProps, field: fieldKey })}</>
          ) : (
            <>
              <Title size={18}>{key}</Title>
              <Space p={8} />
            </>
          )}
        </Group>
      )

      function renderNestedArrayOfObjects() {
        /* // TODO: IF LEN form.values[fieldKey] instead need to get all nested objects values with keys that start with fieldKey,
          e.g. "members" ---> startswith "members." will return schema for those map. DO NOT MAP FROM FORM ITSELF.
          */
        try {
          return form.values[fieldKey]?.map((_nestedValue: any, index: number) => {
            return (
              <div key={index} style={{ marginBottom: theme.spacing.sm }}>
                <p>{`${fieldKey}[${index}]`}</p>
                {/**
                 * TODO: handlers should be shared for nested paths, simply have index opt and if not null construct index access on last path
                 * generateFormFields needs index as well (convert to options {parentPathPrefix string, index: number}), to handle changes as form.values.<path>.<index>.<...> as per https://mantine.dev/form/nested/
                 * we can use form.removeListItem for handleRemove
                 * */}
                {JSON.stringify({ [fieldKey]: schemaFields[fieldKey] })}
                <Group>{generateFormFields({ parentPathPrefix: fieldKey, index })}</Group>
                <ActionIcon onClick={() => removeNestedField(fieldKey, index)} variant="filled" color={'red'}>
                  <IconMinus size="1rem" />
                </ActionIcon>
              </div>
            )
          })
        } catch (error) {
          console.log(`renderNestedArrayOfObjects error: ${error}`)
          return null
        }
      }
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
