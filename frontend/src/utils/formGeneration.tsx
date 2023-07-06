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
  Tooltip,
  Accordion,
} from '@mantine/core'
import { DateInput, DateTimePicker } from '@mantine/dates'
import { Prism } from '@mantine/prism'
import { useMantineTheme } from '@mantine/styles'
import { Icon123, IconMinus, IconPlus, IconTrash } from '@tabler/icons'
import { singularize } from 'inflection'
import _, { memoize } from 'lodash'
import React, { useState, type ComponentProps, useMemo, type MouseEventHandler, memo } from 'react'
import { useFormContext, type Path, type UseFormReturn, FormProvider, useWatch, useFieldArray } from 'react-hook-form'
import { json } from 'react-router-dom'
import PageTemplate from 'src/components/PageTemplate'
import type { RestDemoWorkItemCreateRequest } from 'src/gen/model'
import useRenders from 'src/hooks/utils/useRenders'
import type {
  DeepPartial,
  GenericObject,
  GetKeys,
  RecursiveKeyOf,
  RecursiveKeyOfArray,
  PathType,
  Branded,
} from 'src/types/utils'
import { removeElementByIndex } from 'src/utils/array'
import type { SchemaField } from 'src/utils/jsonSchema'
import { entries } from 'src/utils/object'
import { sentenceCase } from 'src/utils/strings'
import type { U } from 'vitest/dist/types-b7007192'

export type SelectOptionsTypes = 'select' | 'multiselect'

export interface SelectOptions<Return, E = unknown> {
  values: E[]
  type: SelectOptionsTypes
  formValueTransformer?: <V extends E>(el: V & E) => Return
  // TODO: via mantine componentValue
  componentTransformer?: <V extends E>(el: V & E) => JSX.Element
}

export interface InputOptions<Return, E = unknown> {
  component: JSX.Element
}

export const selectOptionsBuilder = <Return, V>({
  type,
  values,
  formValueTransformer,
  componentTransformer,
}: SelectOptions<Return, V>): SelectOptions<Return, V> => ({
  type,
  values,
  componentTransformer,
  formValueTransformer,
})

export const inputBuilder = <Return, V>({ component }: InputOptions<Return, V>): InputOptions<Return, V> => ({
  component,
})

export type DynamicFormOptions<T extends object, ExcludeKeys extends U | null, U extends PropertyKey = GetKeys<T>> = {
  // FIXME: Exclude<U, ExcludeKeys> breaks indexing type inference - but does exclude
  labels: {
    [key in Exclude<U, ExcludeKeys>]: string | null
  }
  // used to populate form inputs if the form field is empty. Applies to all nested fields.
  defaultValues?: Partial<{
    [key in Exclude<U, ExcludeKeys>]: DeepPartial<
      PathType<
        T,
        // can fix key constraint error with U extends RecursiveKeyOf<T, ''> but not worth it due to cpu usage, just ignore
        //@ts-ignore
        key
      >
    >
  }>
  //  list of options used for Select and MultiSelect
  // TODO: someone had the exact same idea: https://stackoverflow.com/questions/69254779/infer-type-based-on-the-generic-type-of-a-sibling-property-in-typescript
  // more recent version: https://stackoverflow.com/questions/74618270/how-to-make-an-object-property-depend-on-another-one-in-a-generic-type
  // TODO: inputComponent field, e.g. for color picker. if inputComponent === undefined, then switch on schema format as usual
  selectOptions?: Partial<{
    [key in Exclude<U, ExcludeKeys>]: ReturnType<
      typeof selectOptionsBuilder<
        PathType<
          T,
          //@ts-ignore
          key
        >,
        unknown
      >
    >
  }>
  /**
   * override default input component.
   */
  input?: Partial<{
    [key in Exclude<U, ExcludeKeys>]: ReturnType<
      typeof inputBuilder<
        PathType<
          T,
          //@ts-ignore
          key
        >,
        unknown
      >
    >
  }>
  propsOverride?: Partial<{
    [key in Exclude<U, ExcludeKeys>]: {
      label?: string
      description?: string
    }
  }>
  accordion?: Partial<{
    [key in Exclude<U, ExcludeKeys>]: {
      defaultOpen?: boolean
      title?: JSX.Element
    }
  }>
}

type DynamicFormProps<T extends object, U extends PropertyKey = GetKeys<T>, ExcludeKeys extends U | null = null> = {
  schemaFields: Record<Exclude<U, ExcludeKeys>, SchemaField>
  options: DynamicFormOptions<T, ExcludeKeys, U>
  formName: string
}

function renderTitle(key: string) {
  return (
    <>
      <Title data-testid={`${key}-title`} size={18}>
        {key}
      </Title>
      <Space p={8} />
    </>
  )
}

const removeListItem = (form, formField: FormField, index: number) => {
  const listItems = form.getValues(formField)
  removeElementByIndex(listItems, index)
  form.setValue(formField, listItems as any)
  console.log(listItems)
}

export default function DynamicForm<
  T extends object,
  ExcludeKeys extends U | null = null,
  U extends PropertyKey = GetKeys<T>,
>({ formName, schemaFields, options }: DynamicFormProps<T, U, ExcludeKeys>) {
  const theme = useMantineTheme()
  const form = useFormContext()

  // TODO: will also need sorting schemaFields beforehand and then generate normally.
  return (
    <PageTemplate minWidth={1000}>
      <>
        <FormData />
        <form
          onSubmit={(e) => {
            e.preventDefault()
            form.handleSubmit(
              (data) => console.log({ data }),
              (errors) => console.log({ errors }),
            )(e)
          }}
          css={css`
            min-width: 100%;
          `}
          id={formName}
        >
          <button type="submit">submit</button>
          <GeneratedInputs schemaFields={schemaFields} formName={formName} options={options} />
        </form>
      </>
    </PageTemplate>
  )
}

type SchemaKey = Branded<string, 'SchemaKey'>
type FormField = Branded<string, 'FormField'>

type GeneratedInputsProps<T extends object, ExcludeKeys extends U | null, U extends PropertyKey = GetKeys<T>> = {
  parentSchemaKey?: SchemaKey
  index?: number
  parentFormField?: FormField
  removeButton?: JSX.Element | null
  schemaFields: Record<SchemaKey, SchemaField>
  formName: string
  options: DynamicFormOptions<T, null, SchemaKey> // for more performant internal intellisense.
}

function GeneratedInputs<T extends object, ExcludeKeys extends U | null, U extends PropertyKey = GetKeys<T>>({
  parentSchemaKey,
  parentFormField,
  schemaFields,
  formName,
  options,
}: GeneratedInputsProps<T, ExcludeKeys, U>) {
  const form = useFormContext()

  /**
   * Construct form accessor based on current schema field key and parent form field.
   */
  const constructFormField = (schemaKey: SchemaKey, parentFormField?: FormField) => {
    const currentFieldName = schemaKey.split('.').slice(-1)[0]

    return (parentFormField ? `${parentFormField}.${currentFieldName}` : schemaKey) as FormField
  }

  const initialValueByKey = (schemaKey: SchemaKey) => {
    switch (schemaFields[schemaKey]?.type) {
      case 'object':
        return {}
      case 'array':
        return []
      case 'number':
      case 'integer':
        return 0
      case 'boolean':
        return false
      default:
        return ''
    }
  }

  // NOTE: useFieldArray can append empty field just once (prevents user spamming add button)
  const addNestedField = (schemaKey: SchemaKey, formField: FormField) => {
    const initialValue = initialValueByKey(schemaKey)

    const vals = form.getValues(formField) || []

    console.log([...vals, initialValue] as any)

    form.setValue(formField, [...vals, initialValue] as any)
  }

  const children = entries(schemaFields).map(([schemaKey, field]) => {
    const renders = useRenders()

    const NestedHeader = () => {
      return (
        <div>
          {/* {<Prism language="json">{JSON.stringify({ formField, parentFormField }, null, 4)}</Prism>} */}
          <Flex direction="row">
            {!accordion && renderTitle(formField)}
            <Button
              size="xs"
              p={4}
              leftIcon={<IconPlus size="1rem" />}
              onClick={() => addNestedField(schemaKey, formField)}
              variant="filled"
              color={'green'}
              id={`${formName}-${formField}-add-button`}
            >{`Add ${formField}`}</Button>
          </Flex>
        </div>
      )
    }

    if (
      (parentSchemaKey && !schemaKey.startsWith(parentSchemaKey)) ||
      parentSchemaKey === schemaKey || // fix when parent key has the same name and both are arrays
      !options.labels.hasOwnProperty(schemaKey) // labels are mandatory unless form field was excluded
    ) {
      return null
    }

    const pp = schemaKey.split('.')
    const parentKey = parentSchemaKey?.replace(/\.*$/, '') || pp.slice(0, pp.length - 1).join('.')

    if (schemaFields[parentKey]?.isArray && !parentSchemaKey) {
      return null
    }

    const formField = constructFormField(schemaKey, parentFormField)

    const formValue = JSON.stringify(form.getValues(formField))
    // console.log({ formField, formValue })
    const accordion = options.accordion?.[schemaKey]

    const containerProps = {
      css: css`
        width: 100%;
      `,
    }

    const inputProps = {
      css: css`
        width: 100%;
      `,
      ...(!field.isArray && { label: formField }),
      required: field.required,
      id: `${formName}-${formField}`,
    }

    if (field.isArray && field.type !== 'object') {
      // nested array of nonbjects generation
      return (
        <Card key={schemaKey} mt={12} mb={12} withBorder>
          {/* existing array fields, if any */}
          {accordion ? (
            <FormAccordion>
              <NestedHeader />
              <ArrayChildren
                formField={formField}
                schemaKey={schemaKey}
                inputProps={inputProps}
                formName={formName}
                containerProps={containerProps}
                options={options}
                schemaFields={schemaFields}
              />
            </FormAccordion>
          ) : (
            <>
              <NestedHeader />
              <ArrayChildren
                formField={formField}
                schemaKey={schemaKey}
                inputProps={inputProps}
                formName={formName}
                containerProps={containerProps}
                options={options}
                schemaFields={schemaFields}
              />
            </>
          )}
        </Card>
      )
    }

    if (field.isArray && field.type === 'object') {
      // array of objects
      return (
        // TODO: background color based on depth
        <Card key={schemaKey} mt={12} mb={12} withBorder>
          {accordion ? (
            <FormAccordion>
              <NestedHeader />
              <ArrayOfObjectsChildren
                formField={formField}
                formName={formName}
                schemaKey={schemaKey}
                options={options}
                schemaFields={schemaFields}
              />
            </FormAccordion>
          ) : (
            <>
              <NestedHeader />
              <ArrayOfObjectsChildren
                formField={formField}
                formName={formName}
                schemaKey={schemaKey}
                options={options}
                schemaFields={schemaFields}
              />
            </>
          )}
        </Card>
      )
    }

    return (
      <Group key={schemaKey} align="center">
        {field.type !== 'object' ? (
          <>
            <GeneratedInput
              schemaKey={schemaKey}
              formField={formField}
              props={{ input: inputProps, container: containerProps }}
              options={options}
              schemaFields={schemaFields}
              formName={formName}
            />
          </>
        ) : (
          <>{renderTitle(formField)}</>
        )}
      </Group>
    )

    function FormAccordion({ children }): JSX.Element | null {
      if (!accordion) return null

      const value = `${schemaKey}-accordion`

      return (
        <Accordion
          defaultValue={accordion.defaultOpen ? value : null}
          styles={{ control: { padding: 0, maxHeight: '28px' } }}
          {...containerProps}
        >
          <Accordion.Item value={value}>
            <Accordion.Control>{accordion.title ?? `${schemaKey}`}</Accordion.Control>
            <Accordion.Panel>{children}</Accordion.Panel>
          </Accordion.Item>
        </Accordion>
      )
    }
  })

  return <>{children}</>
}

type ArrayOfObjectsChildrenProps<T extends object, ExcludeKeys extends U | null, U extends PropertyKey = GetKeys<T>> = {
  formField: FormField
  schemaKey: SchemaKey
  formName: string
  options: DynamicFormOptions<T, null, SchemaKey> // for more performant internal intellisense.
  schemaFields: Record<SchemaKey, SchemaField>
}

function ArrayOfObjectsChildren<T extends object, ExcludeKeys extends U | null, U extends PropertyKey = GetKeys<T>>({
  formField,
  formName,
  schemaKey,
  options,
  schemaFields,
}: ArrayOfObjectsChildrenProps<T, ExcludeKeys, U>) {
  const form = useFormContext()
  const fieldArray = useFieldArray({
    control: form.control,
    name: formField,
  })
  // form.watch(formField, fieldArray.fields) // inf rerendering
  // useWatch({ name: `${formField}`, control: form.control }) // same errors

  const children = fieldArray.fields.map((item, k) => {
    return (
      <div
        key={item.id}
        css={css`
          min-width: 100%;
        `}
      >
        <Text weight={800}>{`${formField}.${k}`}</Text>
        <Card mt={12} mb={12} withBorder>
          <Tooltip withinPortal label={`Remove ${singularize(options.labels[schemaKey])}`} position="top-end" withArrow>
            <ActionIcon
              onClick={(e) => {
                fieldArray.remove(k)
                // removeListItem(form, formField, k)
              }}
              // variant="filled"
              css={css`
                background-color: #7c1a1a;
              `}
              size="sm"
              id={`${formName}-${formField}-remove-button-${k}`}
            >
              <IconTrash size="1rem" />
            </ActionIcon>
          </Tooltip>
          <Group>
            <GeneratedInputs
              parentSchemaKey={schemaKey}
              parentFormField={`${formField}.${k}` as FormField}
              schemaFields={schemaFields}
              formName={formName}
              options={options}
            />
          </Group>
        </Card>
      </div>
    )
  })

  return (
    <Flex gap={14} align="center" direction="column">
      {children}
    </Flex>
  )
}

type ArrayChildrenProps<T extends object, ExcludeKeys extends U | null, U extends PropertyKey = GetKeys<T>> = {
  formField: FormField
  schemaKey: SchemaKey
  formName: string
  options: DynamicFormOptions<T, null, SchemaKey> // for more performant internal intellisense.
  schemaFields: Record<SchemaKey, SchemaField>
  inputProps: any
  containerProps: any
}

function ArrayChildren<T extends object, ExcludeKeys extends U | null, U extends PropertyKey = GetKeys<T>>({
  formField,
  formName,
  schemaKey,
  inputProps,
  containerProps,
  options,
  schemaFields,
}: ArrayChildrenProps<T, ExcludeKeys, U>) {
  const form = useFormContext()

  // IMPORTANT: https://react-hook-form.com/docs/usefieldarray Does not support flat field array.

  useWatch({ name: `${formField}`, control: form.control }) // same errors

  const children = (form.getValues(formField) || []).map((item, k: number) => {
    return (
      <Flex
        key={k}
        css={css`
          min-width: 100%;
        `}
      >
        <GeneratedInput
          formName={formName}
          schemaKey={schemaKey}
          formField={`${formField}.${k}` as FormField}
          props={{
            input: { ...inputProps, id: `${formName}-${formField}-${k}` },
            container: containerProps,
          }}
          options={options}
          schemaFields={schemaFields}
          index={k}
        />
      </Flex>
    )
  })

  return (
    <Flex gap={14} align="center" direction="column">
      {children}
    </Flex>
  )
}

function FormData() {
  const myFormData = useWatch()

  return (
    <Accordion>
      <Accordion.Item value="form">
        <Accordion.Control>See form</Accordion.Control>
        <Accordion.Panel>
          <Prism language="json">{JSON.stringify(myFormData, null, 2)}</Prism>
        </Accordion.Panel>
      </Accordion.Item>
    </Accordion>
  )
}

type GeneratedInputProps<T extends object, ExcludeKeys extends U | null, U extends PropertyKey = GetKeys<T>> = {
  schemaKey: SchemaKey
  props?: {
    input?: any
    container?: any
  }
  formField: FormField
  withRemoveButton?: boolean
  schemaFields: Record<SchemaKey, SchemaField>
  options: DynamicFormOptions<T, ExcludeKeys, U>
  index?: number
  formName: string
}

const convertValueByType = (type: SchemaField['type'] | undefined, value) => {
  switch (type) {
    case 'date':
    case 'date-time':
      return new Date(value)
    default:
      return value
  }
}

// useMemo with dep list of [JSON.stringify(_.get(form.values, formField)), ...] (will always rerender if its object, but if string only when it changes)
// TODO: just migrate to react-hook-form: https://codesandbox.io/s/dynamic-radio-example-forked-et0wi?file=/src/content/FirstFormSection.tsx
// for builtin support for uncontrolled input
const GeneratedInput = <T extends object, ExcludeKeys extends U | null, U extends PropertyKey = GetKeys<T>>({
  schemaKey,
  props,
  formField,
  withRemoveButton = false,
  options,
  schemaFields,
  index,
  formName,
}: GeneratedInputProps<T, ExcludeKeys, U>) => {
  const form = useFormContext()
  // useWatch({ control: form.control, name: formField }) // completely unnecessary, it's registered...

  const propsOverride = options.propsOverride?.[schemaKey]
  const type = schemaFields[schemaKey]?.type

  const { onChange: registerOnChange, ...registerProps } = form.register(formField, {
    ...(type === 'date' || type === 'date-time'
      ? // TODO: use convertValueByType
        { valueAsDate: true, setValueAs: (v) => (v === '' ? undefined : new Date(v)) }
      : type === 'integer'
      ? { valueAsNumber: true, setValueAs: (v) => (v === '' ? undefined : parseInt(v, 10)) }
      : type === 'number'
      ? { valueAsNumber: true, setValueAs: (v) => (v === '' ? undefined : parseFloat(v)) }
      : type === 'boolean'
      ? { setValueAs: (v) => (v === '' ? undefined : v === 'true') }
      : null),
  })

  const fieldState = form.getFieldState(formField)
  // FIXME: https://stackoverflow.com/questions/75437898/react-hook-form-react-select-cannot-read-properties-of-undefined-reading-n
  // mantine does not alter TextInput onChange but we need to customize onChange for the rest and call rhf onChange manually with
  // value modified back to normal

  const formFieldKeys = formField.split('.')
  // remove last index
  const formFieldArrayPath = formFieldKeys.slice(0, formFieldKeys.length - 1).join('.')

  const formValue = form.getValues(formField)

  if (formValue === null || formValue === undefined) {
    const defaultValue = options.defaultValues?.[schemaKey]
    if (defaultValue) {
      form.setValue(formField, defaultValue)
    }
  }

  // TODO: multiselect and select early check (if found in options.components override)
  const _props = {
    ...registerProps,
    ...props?.input,
    ...(propsOverride && propsOverride),
    ...(!fieldState.isDirty && { defaultValue: convertValueByType(type, formValue) }),
    ...(fieldState.error && { error: sentenceCase(fieldState.error?.message) }),
    required: schemaFields[schemaKey]?.required && type !== 'boolean',
  }

  let el: JSX.Element | null = null
  const component = options.input?.[schemaKey]?.component
  if (component) {
    el = React.cloneElement(component, {
      ..._props,
      ...component.props, // allow user override
      // TODO: this depends on component type, onChange should be customizable in options parameter with registerOnChange as fn param
      onChange: (e) => registerOnChange({ target: { name: formField, value: e } }),
    })
  } else {
    switch (schemaFields[schemaKey]?.type) {
      case 'string':
        el = (
          <TextInput
            onChange={(e) => registerOnChange({ target: { name: formField, value: e.target.value } })}
            {..._props}
          />
        )
        break
      case 'boolean':
        el = (
          <Checkbox
            onChange={(e) => registerOnChange({ target: { name: formField, value: e.target.checked } })}
            pt={10}
            pb={4}
            {..._props}
          />
        )
        break
      case 'date':
        el = (
          <DateInput
            valueFormat="DD/MM/YYYY"
            onChange={(e) =>
              registerOnChange({
                target: { name: formField, value: e },
              })
            }
            placeholder="Select date"
            {..._props}
          />
        )
        break
      case 'date-time':
        el = (
          <DateTimePicker
            onChange={(e) =>
              registerOnChange({
                target: { name: formField, value: e },
              })
            }
            placeholder="Select date and time"
            {..._props}
          />
        )
        break
      case 'integer':
        el = <NumberInput onChange={(e) => registerOnChange({ target: { name: formField, value: e } })} {..._props} />
        break
      case 'number':
        el = (
          <NumberInput
            onChange={(e) => registerOnChange({ target: { name: formField, value: e } })}
            precision={2}
            {..._props}
          />
        )
        break
      default:
        break
    }
  }

  return (
    <Flex align="center" justify={'center'} {...props?.container}>
      {el}
      {index !== undefined && (
        <Tooltip withinPortal label={`Remove ${singularize(options.labels[schemaKey])}`} position="top-end" withArrow>
          <ActionIcon
            onClick={(e) => {
              // fieldArray.remove(index) // doesn't work on flat arrays
              console.log({ formField, index, currentFormValue: form.getValues(formField) })
              removeListItem(form, formField, index)
            }}
            // variant="filled"
            css={css`
              background-color: #7c1a1a;
            `}
            size="sm"
            id={`${formName}-${formField}-remove-button-${index}`}
          >
            <IconMinus size="1rem" />
          </ActionIcon>
        </Tooltip>
      )}
    </Flex>
  )
}
