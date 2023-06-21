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
} from '@mantine/core'
import { DateInput } from '@mantine/dates'
import { Form, type UseFormReturnType } from '@mantine/form'
import { useMantineTheme } from '@mantine/styles'
import { IconMinus, IconPlus } from '@tabler/icons'
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

export const DynamicForm = <T extends string, U extends GenericObject>({
  form,
  schemaFields,
  options,
}: DynamicFormProps<T, U>) => {
  const theme = useMantineTheme()

  const handleChange = (value: any, field: string) => {
    form.setValues((currentValues) => ({
      ...currentValues,
      [field]: value,
    }))
  }

  const handleNestedChange = (value: any, field: string, index: number) => {
    form.setValues((currentValues) => ({
      ...currentValues,
      [field]: currentValues[field].map((item: any, i: number) => (i === index ? value : item)),
    }))
  }

  const handleAddNestedField = (field: string) => {
    form.setValues((currentValues) => ({
      ...currentValues,
      [field]: [
        ...(currentValues[field] || []),
        null,
        // { role: 'preparer', userID: 'rsfsese' }, // should have initial object generated based on path if type === object, else it will attempt setting on null
      ],
    }))
  }

  const handleRemoveNestedField = (field: string, index: number) => {
    form.setValues((currentValues) => ({
      ...currentValues,
      [field]: currentValues[field].filter((_item: any, i: number) => i !== index),
    }))
  }

  const generateFormFields = (fields: DynamicFormProps<T, U>['schemaFields'], prefix = '') => {
    const generateComponent = (fieldType: SchemaField['type'], props: any, field: string) => {
      const paths = field.split('.')
      const parent = paths.slice(0, paths.length - 1).join('.')
      if (fields[parent]?.isArray) {
        return null
      }

      switch (fieldType) {
        case 'string':
          return <TextInput {...{ ...props, ...form.getInputProps(field) }} />
        case 'boolean':
          return <Checkbox {...{ ...props, ...form.getInputProps(field) }} />
        case 'date-time':
          return <DateInput placeholder="Date input" {...{ ...props, ...form.getInputProps(field) }} />
        case 'integer':
          return <NumberInput {...{ ...props, ...form.getInputProps(field) }} />
        default:
          return null
      }
    }

    return entries(fields).map(([key, field]) => {
      if (prefix !== '' && !key.startsWith(prefix)) {
        return
      }
      // TODO: check if parent is isArray, in which case return early and do nothing, since
      // children have already been generated.
      // console.log(prefix)
      // // Skip generating fields for array items
      // if (fields[prefix]?.isArray) {
      //   return null
      // }

      const fieldKey = prefix !== '' ? `${prefix}.${key}` : key
      const value = form.values[fieldKey] || options[fieldKey]?.defaultValue || ''

      const componentProps = {
        css: css`
          min-width: 100%;
        `,
        label: key,
        required: field.required,
        value: value,
        onChange:
          field.type === 'integer' || field.type === 'date-time'
            ? (val: any) => handleChange(val, fieldKey)
            : (event: any) => handleChange(event.currentTarget.value, fieldKey),
      }

      if (field.isArray && field.type !== 'object') {
        // array of primitives

        // TODO: form.getInputProps instead.
        // form.getInputProps('base.<nested>', {type: "checkbox | input"})

        return (
          <Group key={fieldKey}>
            {JSON.stringify(form.values[fieldKey])}
            <div style={{ display: 'flex', marginBottom: theme.spacing.xs }}>
              <ActionIcon onClick={() => handleAddNestedField(fieldKey)} variant="filled" color={'green'}>
                <IconPlus size="1rem" />
              </ActionIcon>
              {generateComponent(field.type, componentProps, fieldKey)}
            </div>
            {/* existing array fields, if any */}
            {form.values[fieldKey]?.map((_nestedValue: any, index: number) => {
              return (
                <div key={index} style={{ display: 'flex', marginBottom: theme.spacing.xs }}>
                  {generateComponent(
                    field.type,
                    {
                      ...componentProps,
                      value: form.values[fieldKey]?.[index] || '',
                      onChange:
                        field.type === 'integer' || field.type === 'date-time'
                          ? (val: any) => handleNestedChange(val, fieldKey, index)
                          : (event: any) => handleNestedChange(event.currentTarget.value, fieldKey, index),
                    },
                    fieldKey,
                  )}
                  <ActionIcon onClick={() => handleRemoveNestedField(fieldKey, index)} variant="filled" color={'green'}>
                    <IconMinus size="1rem" />
                  </ActionIcon>
                </div>
              )
            })}
          </Group>
        )
      }

      if (field.isArray && field.type === 'object') {
        if (fields[prefix]?.isArray) {
          return null
        }
        // array of objects
        return (
          <Card key={fieldKey} mt={24}>
            {JSON.stringify({ [fieldKey]: form.values[fieldKey] })}
            <ActionIcon onClick={() => handleAddNestedField(fieldKey)} variant="filled" color={'green'}>
              <IconPlus size="1rem" />
            </ActionIcon>
            {form.values[fieldKey]?.map((_nestedValue: any, index: number) => {
              return (
                <div key={index} style={{ marginBottom: theme.spacing.sm }}>
                  <p>{`${fieldKey}[${index}]`}</p>
                  <Group>{generateFormFields(fields, fieldKey)}</Group>
                  <ActionIcon onClick={() => handleRemoveNestedField(fieldKey, index)} variant="filled" color={'red'}>
                    <IconMinus size="1rem" />
                  </ActionIcon>
                </div>
              )
            })}
          </Card>
        )
      }

      return (
        <Group key={fieldKey} align="center">
          {field.type !== 'object' ? (
            <>{generateComponent(field.type, componentProps, fieldKey)}</>
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
        {generateFormFields(schemaFields)}
      </form>
    </PageTemplate>
  )
}
