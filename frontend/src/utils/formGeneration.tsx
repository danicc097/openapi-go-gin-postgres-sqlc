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

const DynamicForm = <T extends string, U extends GenericObject>({
  schemaFields,
  optionsOverride,
}: DynamicFormProps<T, U>) => {
  const theme = useMantineTheme()
  const [formData, setFormData] = useState<any>({})

  optionsOverride.defaultValue['members.role']

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

  const generateFormFields = (fields: DynamicFormProps<T, U>['schemaFields'], prefix = '') => {
    return entries(fields).map(([key, field]) => {
      const fieldKey = prefix ? `${prefix}.${key}` : key
      const value = formData[fieldKey] || optionsOverride[fieldKey].defaultValue || ''

      if (field.isArray && field.type !== 'object') {
        return (
          <Group key={fieldKey}>
            <div style={{ display: 'flex', marginBottom: theme.spacing.xs }}>
              {field.type === 'string' && (
                <TextInput
                  required={field.required}
                  value={value}
                  onChange={(event) => handleChange(event.currentTarget.value, fieldKey)}
                />
              )}
              {field.type === 'boolean' && (
                <Checkbox
                  required={field.required}
                  checked={value}
                  onChange={(event) => handleChange(event.currentTarget.checked, fieldKey)}
                />
              )}
              {field.type === 'date-time' && (
                <DateInput required={field.required} value={value} onChange={(date) => handleChange(date, fieldKey)} />
              )}
              {field.type === 'integer' && (
                <NumberInput
                  required={field.required}
                  value={value}
                  onChange={(event) => handleChange(Number(event), fieldKey)}
                />
              )}
              <Button
                onClick={() => handleAddNestedField(fieldKey)}
                style={{ marginLeft: theme.spacing.sm }}
                leftIcon={<IconPlus />}
              ></Button>
            </div>
            {formData[fieldKey]?.map((_nestedValue: any, index: number) => (
              <div key={index} style={{ display: 'flex', marginBottom: theme.spacing.xs }}>
                {field.type === 'string' && (
                  <TextInput
                    required={field.required}
                    value={formData[fieldKey]?.[index] || ''}
                    onChange={(event) => handleNestedChange(event.currentTarget.value, fieldKey, index)}
                  />
                )}
                {field.type === 'boolean' && (
                  <Checkbox
                    required={field.required}
                    checked={formData[fieldKey]?.[index] || false}
                    onChange={(event) => handleNestedChange(event.currentTarget.checked, fieldKey, index)}
                  />
                )}
                {field.type === 'date-time' && (
                  <DateInput
                    required={field.required}
                    value={formData[fieldKey]?.[index] || ''}
                    onChange={(date) => handleNestedChange(date, fieldKey, index)}
                  />
                )}
                {field.type === 'integer' && (
                  <NumberInput
                    required={field.required}
                    value={formData[fieldKey]?.[index] || ''}
                    onChange={(event) => handleNestedChange(Number(event), fieldKey, index)}
                  />
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

DynamicForm<RestDemoWorkItemCreateRequestFormField, RestDemoWorkItemCreateRequest>({
  schemaFields: {
    base: { isArray: false, required: true, type: 'object' },
    'base.closed': { type: 'date-time', required: true, isArray: false },
    'base.description': { type: 'string', required: true, isArray: false },
    'base.kanbanStepID': { type: 'integer', required: true, isArray: false },
    'base.metadata': { type: 'integer', required: true, isArray: true },
    'base.targetDate': { type: 'date-time', required: true, isArray: false },
    'base.teamID': { type: 'integer', required: true, isArray: false },
    'base.title': { type: 'string', required: true, isArray: false },
    'base.workItemTypeID': { type: 'integer', required: true, isArray: false },
    demoProject: { isArray: false, required: true, type: 'object' },
    'demoProject.lastMessageAt': { type: 'date-time', required: true, isArray: false },
    'demoProject.line': { type: 'string', required: true, isArray: false },
    'demoProject.ref': { type: 'string', required: true, isArray: false },
    'demoProject.reopened': { type: 'boolean', required: true, isArray: false },
    'demoProject.workItemID': { type: 'integer', required: true, isArray: false },
    members: { type: 'object', required: true, isArray: true },
    'members.role': { type: 'string', required: true, isArray: false },
    'members.userID': { type: 'string', required: true, isArray: false },
    tagIDs: { type: 'integer', required: true, isArray: true },
  },
  optionsOverride: {
    defaultValue: {
      'demoProject.line': '534543523', // should fail due to TypeOf
      members: [{ role: 'preparer', userID: 'c446259c-1083-4212-98fe-bd080c41e7d7' }],
    },
  },
})
