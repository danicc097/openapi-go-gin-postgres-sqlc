import { Group, TextInput, NumberInput, Checkbox, Button } from '@mantine/core'
import { DateInput } from '@mantine/dates'
import { Form } from '@mantine/form'
import { useMantineTheme } from '@mantine/styles'
import { IconMinus, IconPlus } from '@tabler/icons'
import { useState } from 'react'
import type { FieldPath } from 'react-hook-form'
import type { RestDemoWorkItemCreateRequest } from 'src/gen/model'
import type { GenericObject, RecursiveKeyOf, RecursiveKeyOfArray, TypeOf } from 'src/types/utils'
import type { SchemaField } from 'src/utils/jsonSchema'
import { entries } from 'src/utils/object'

type RestDemoWorkItemCreateRequestFormField =
  // hack to use 'members.role' instead of 'members.??.role'
  FieldPath<RestDemoWorkItemCreateRequest> | RecursiveKeyOfArray<RestDemoWorkItemCreateRequest['members'], 'members'>

type OptionsOverride<T extends string, U extends GenericObject> = {
  defaultValue: Partial<{
    [key in T]: TypeOf<U, key>
  }>
}

type DynamicFormProps<T extends string, U extends GenericObject> = {
  schemaFields: Record<T, SchemaField>
  optionsOverride: OptionsOverride<T, U>
}

export const DynamicForm = <T extends string, U extends GenericObject>({
  schemaFields,
  optionsOverride,
}: DynamicFormProps<T, U>) => {
  const theme = useMantineTheme()
  const [formData, setFormData] = useState<any>({})

  const handleChange = (value: any, field: string) => {
    setFormData((prevData) => ({ ...prevData, [field]: value }))
  }

  const handleNestedChange = (value: any, field: string, index: number) => {
    setFormData((prevData) => ({
      ...prevData,
      [field]: prevData[field].map((item: any, i: number) => (i === index ? value : item)),
    }))
  }

  const handleAddNestedField = (field: string) => {
    setFormData((prevData) => ({
      ...prevData,
      [field]: [...(prevData[field] || []), ''],
    }))
  }

  const handleRemoveNestedField = (field: string, index: number) => {
    setFormData((prevData) => ({
      ...prevData,
      [field]: prevData[field].filter((_item: any, i: number) => i !== index),
    }))
  }

  const generateComponent = (fieldType: SchemaField['type'], props: any) => {
    switch (fieldType) {
      case 'string':
        return <TextInput {...props} />
      case 'boolean':
        return <Checkbox {...props} />
      case 'date-time':
        return <DateInput {...props} />
      case 'integer':
        return <NumberInput {...props} />
      default:
        return null
    }
  }

  const generateFormFields = (fields: DynamicFormProps<T, U>['schemaFields'], prefix = '') => {
    return entries(fields).map(([key, field]) => {
      const fieldKey = prefix ? `${prefix}.${key}` : key
      const value = formData[fieldKey] || optionsOverride[fieldKey]?.defaultValue || ''

      if (field.isArray && field.type !== 'object') {
        const componentProps = {
          required: field.required,
          value: value,
          onChange: (event: any) => handleChange(event.currentTarget.value, fieldKey),
        }

        return (
          <Group key={fieldKey}>
            <div style={{ display: 'flex', marginBottom: theme.spacing.xs }}>
              {generateComponent(field.type, componentProps)}
              <Button
                onClick={() => handleAddNestedField(fieldKey)}
                style={{ marginLeft: theme.spacing.sm }}
                leftIcon={<IconPlus />}
              ></Button>
            </div>
            {formData[fieldKey]?.map((_nestedValue: any, index: number) => (
              <div key={index} style={{ display: 'flex', marginBottom: theme.spacing.xs }}>
                {generateComponent(field.type, {
                  ...componentProps,
                  value: formData[fieldKey]?.[index] || '',
                  onChange: (event: any) => handleNestedChange(event.currentTarget.value, fieldKey, index),
                })}
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
        return (
          <Group key={fieldKey}>
            <div style={{ display: 'flex', marginBottom: theme.spacing.xs }}>
              <button onClick={() => handleAddNestedField(fieldKey)} style={{ marginLeft: theme.spacing.sm }}>
                +
              </button>
            </div>
            {formData[fieldKey]?.map((_nestedValue: any, index: number) => (
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
        <TextInput
          key={fieldKey}
          required={field.required}
          value={value}
          onChange={(event) => handleChange(event.currentTarget.value, fieldKey)}
        />
      )
    })
  }

  return <form>{generateFormFields(schemaFields)}</form>
}