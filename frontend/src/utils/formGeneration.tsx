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

  const handleFieldChange = (value: any, field: string, index?: number) => {
    const paths = field.split('.')
    const path =
      index !== undefined ? [...paths.slice(0, paths.length - 2), index, paths[paths.length - 1]].join('.') : field
    console.log({ path, value })
    form.setFieldValue(path, value)
  }

  const addNestedField = (field: string) => {
    console.log({ addNestedField: field })
    form.setValues((currentValues) => ({
      ...currentValues,
      [field]: [
        ...(currentValues[field] || []),
        // TODO: maybe will need initialValue based on type from schema
        // null,
        // 0,
        { role: 'preparer', userID: 'rsfsese' }, // should have initial object generated based on path if type === object, else it will attempt setting on null
      ],
    }))
  }

  const removeNestedField = (field: string, index: number) => {
    console.log({ removeNestedField: `${field}[${index}]` })

    form.removeListItem(field, index)
  }

  const generateFormFields = (
    fields: DynamicFormProps<T, U>['schemaFields'],
    { prefix = '', index }: { prefix?: string; index?: number },
  ) => {
    const generateComponent = (fieldType: SchemaField['type'], props: any, field: string, index?: number) => {
      const paths = field.split('.')
      const parent = paths.slice(0, paths.length - 1).join('.')
      if (fields[parent]?.isArray) {
        return null
      }

      const _field = index !== undefined ? [parent, index, paths[paths.length - 1]].join('.') : field

      console.log({ val: form.getInputProps(_field).value, onchange: form.getInputProps(_field).onChange })

      switch (fieldType) {
        // FIXME: wrong form.getInputProps, bad fieldKey when indexes involved
        case 'string':
          return (
            <TextInput
              {...{
                ...form.getInputProps(_field),
                ...props,
              }}
            />
          )
        case 'boolean':
          return (
            <Checkbox
              {...{
                ...form.getInputProps(_field),
                ...props,
              }}
            />
          )
        case 'date-time':
          return <DateInput placeholder="Date input" {...{ ...form.getInputProps(_field), ...props }} />
        case 'integer':
          return (
            <NumberInput
              {...{
                ...form.getInputProps(_field),
                ...props,
              }}
            />
          )
        default:
          return null
      }
    }

    return entries(fields).map(([key, field]) => {
      if (prefix !== '' && !key.startsWith(prefix)) {
        return
      }

      let fieldKey = prefix !== '' ? `${prefix}.${key}` : key
      if (index !== undefined) {
        fieldKey = `${fieldKey}.${index}`
      }
      const value = form.values[fieldKey] || options[fieldKey]?.defaultValue || ''
      console.log({ value, fieldKey })

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
          <Group key={fieldKey}>
            {JSON.stringify(form.values[fieldKey])}
            <div style={{ display: 'flex', marginBottom: theme.spacing.xs }}>
              <ActionIcon onClick={() => addNestedField(fieldKey)} variant="filled" color={'green'}>
                <IconPlus size="1rem" />
              </ActionIcon>
              {generateComponent(field.type, componentProps, fieldKey, index)}
            </div>
            {/* existing array fields, if any */}
            {form.values[fieldKey]?.map((_nestedValue: any, index: number) => {
              return (
                <div key={index} style={{ display: 'flex', marginBottom: theme.spacing.xs }}>
                  {JSON.stringify({ [fieldKey]: index })}
                  {generateComponent(
                    field.type,
                    {
                      ...componentProps,
                      value: form.values[fieldKey]?.[index] || '',
                      onChange:
                        field.type === 'integer' || field.type === 'date-time'
                          ? (val: any) => handleFieldChange(val, fieldKey, index)
                          : (event: any) => handleFieldChange(event.currentTarget.value, fieldKey, index),
                    },
                    fieldKey,
                    index,
                  )}
                  <ActionIcon onClick={() => removeNestedField(fieldKey, index)} variant="filled" color={'green'}>
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
            <ActionIcon onClick={() => addNestedField(fieldKey)} variant="filled" color={'green'}>
              <IconPlus size="1rem" />
            </ActionIcon>

            {renderNestedArrayOfObjects()}
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
                 * generateFormFields needs index as well (convert to options {prefix string, index: number}), to handle changes as form.values.<path>.<index>.<...> as per https://mantine.dev/form/nested/
                 * we can use form.removeListItem for handleRemove
                 * */}
                {JSON.stringify({ [fieldKey]: fields[fieldKey] })}
                <Group>{generateFormFields(fields, { prefix: fieldKey, index })}</Group>
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
        {generateFormFields(schemaFields, {})}
      </form>
    </PageTemplate>
  )
}
