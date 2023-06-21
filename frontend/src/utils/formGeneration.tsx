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
      [field]: [...(currentValues[field] || []), ''],
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
      // TODO: check if parent is isArray, in which case return early and do nothing, since
      // children have already been generated.
      // console.log(prefix)
      // // Skip generating fields for array items
      // if (fields[prefix]?.isArray) {
      //   return null
      // }

      const fieldKey = prefix ? `${prefix}.${key}` : key
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
            <div style={{ display: 'flex', marginBottom: theme.spacing.xs }}>
              {generateComponent(field.type, componentProps, fieldKey)}
              <Button
                onClick={() => handleAddNestedField(fieldKey)}
                style={{ marginLeft: theme.spacing.sm }}
                leftIcon={<IconPlus />}
              ></Button>
            </div>
            {form.values[fieldKey]?.map((_nestedValue: any, index: number) => (
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
                <Button
                  onClick={() => handleRemoveNestedField(fieldKey, index)}
                  style={{ marginLeft: theme.spacing.sm }}
                  leftIcon={<IconMinus />}
                ></Button>
              </div>
            ))}
          </Group>
        )
      }

      if (field.isArray && field.type === 'object') {
        // array of objects
        return (
          <Group key={fieldKey}>
            <div style={{ display: 'flex', marginBottom: theme.spacing.xs }}>
              <button onClick={() => handleAddNestedField(fieldKey)} style={{ marginLeft: theme.spacing.sm }}>
                +
              </button>
            </div>
            {form.values[fieldKey]?.map((_nestedValue: any, index: number) => (
              <div key={index} style={{ marginBottom: theme.spacing.sm }}>
                <Group>{generateFormFields(fields[key] as any, fieldKey)}</Group>
                <button
                  onClick={() => handleRemoveNestedField(fieldKey, index)}
                  style={{ marginLeft: theme.spacing.sm }}
                >
                  -
                </button>
              </div>
            ))}
          </Group>
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
