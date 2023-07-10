/* eslint-disable @typescript-eslint/ban-ts-comment */
import { css } from '@emotion/react'
import type { EmotionJSX } from '@emotion/react/types/jsx-namespace'
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
  Select,
  Avatar,
} from '@mantine/core'
import { DateInput, DateTimePicker } from '@mantine/dates'
import { Prism } from '@mantine/prism'
import { useMantineTheme } from '@mantine/styles'
import { Icon123, IconMinus, IconPlus, IconTrash } from '@tabler/icons'
import { singularize } from 'inflection'
import _, { memoize } from 'lodash'
import React, {
  useState,
  type ComponentProps,
  useMemo,
  type MouseEventHandler,
  memo,
  createContext,
  useContext,
  type PropsWithChildren,
  forwardRef,
} from 'react'
import {
  useFormContext,
  type Path,
  type UseFormReturn,
  FormProvider,
  useWatch,
  useFieldArray,
  type UseFieldArrayReturn,
  useFormState,
} from 'react-hook-form'
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
import { nameInitials, sentenceCase } from 'src/utils/strings'
import type { U } from 'vitest/dist/types-b7007192'

export type SelectOptionsTypes = 'select' | 'multiselect'

export interface SelectOptions<Return, E = unknown> {
  values: E[]
  type: SelectOptionsTypes
  formValueTransformer: <V extends E>(el: V & E) => Return
  optionTransformer: <V extends E>(el: V & E) => JSX.Element
  labelTransformer: <V extends E>(el: V & E) => string
}

export interface InputOptions<Return, E = unknown> {
  component: JSX.Element
}

export const selectOptionsBuilder = <Return, V>({
  type,
  values,
  formValueTransformer,
  optionTransformer,
  labelTransformer,
}: SelectOptions<Return, V>): SelectOptions<Return, V> => ({
  type,
  values,
  optionTransformer,
  labelTransformer,
  formValueTransformer,
})

export const inputBuilder = <Return, V>({ component }: InputOptions<Return, V>): InputOptions<Return, V> => ({
  component,
})

const itemComponentTemplate = (transformer: (...args: any[]) => JSX.Element) =>
  forwardRef<HTMLDivElement, any>(({ value, option, ...others }, ref) => {
    return (
      <div ref={ref} {...others}>
        {transformer(option)}
      </div>
    )
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

type DynamicFormContextValue = {
  formName: string
  schemaFields: Record<SchemaKey, SchemaField>
  options: DynamicFormOptions<any, null, SchemaKey> // for more performant internal intellisense. for user it will be typed
}

const DynamicFormContext = createContext<DynamicFormContextValue | undefined>(undefined)

type DynamicFormProviderProps = {
  value: DynamicFormContextValue
  children: React.ReactNode
}

const DynamicFormProvider = ({ value, children }: DynamicFormProviderProps) => {
  return <DynamicFormContext.Provider value={value}>{children}</DynamicFormContext.Provider>
}

const useDynamicFormContext = (): DynamicFormContextValue => {
  const context = useContext(DynamicFormContext)

  if (!context) {
    throw new Error('useDynamicFormContext must be used within a DynamicFormProvider')
  }

  return context
}

type DynamicFormProps<T extends object, U extends PropertyKey = GetKeys<T>, ExcludeKeys extends U | null = null> = {
  schemaFields: Record<Exclude<U, ExcludeKeys>, SchemaField>
  options: DynamicFormOptions<T, ExcludeKeys, U>
  formName: string
}

function renderTitle(key: FormField, title) {
  return (
    <>
      <Title data-testid={`${key}-title`} size={18}>
        {title}
      </Title>
      <Space p={8} />
    </>
  )
}

const cardRadius = 6

export default function DynamicForm<
  T extends object,
  ExcludeKeys extends U | null = null,
  U extends PropertyKey = GetKeys<T>,
>({ formName, schemaFields, options }: DynamicFormProps<T, U, ExcludeKeys>) {
  const theme = useMantineTheme()
  const form = useFormContext()

  const { isDirty, isSubmitting, submitCount } = form.formState

  // TODO: will also need sorting schemaFields beforehand and then generate normally.
  return (
    <DynamicFormProvider value={{ formName, options, schemaFields }}>
      <PageTemplate minWidth={800}>
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
            data-testid={formName}
          >
            <button type="submit">submit</button>
            <GeneratedInputs />
          </form>
        </>
      </PageTemplate>
    </DynamicFormProvider>
  )
}

/**
 * Construct form accessor based on current schema field key and parent form field.
 */
const constructFormField = (schemaKey: SchemaKey, parentFormField?: FormField) => {
  const currentFieldName = schemaKey.split('.').slice(-1)[0]

  return (parentFormField ? `${parentFormField}.${currentFieldName}` : schemaKey) as FormField
}

type SchemaKey = Branded<string, 'SchemaKey'>
type FormField = Branded<string, 'FormField'>

type GeneratedInputsProps = {
  parentSchemaKey?: SchemaKey
  index?: number
  parentFormField?: FormField
  removeButton?: JSX.Element | null
}

const containerProps = {
  css: css`
    width: 100%;
  `,
}

function GeneratedInputs({ parentSchemaKey, parentFormField }: GeneratedInputsProps) {
  const { formName, options, schemaFields } = useDynamicFormContext()

  const children = entries(schemaFields).map(([schemaKey, field]) => {
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

    const accordion = options.accordion?.[schemaKey]

    const inputProps = {
      css: css`
        width: 100%;
      `,
      ...(!field.isArray && { label: formField }),
      required: field.required,
      id: `${formName}-${formField}`,
    }
    const itemName = singularize(options.labels[schemaKey] || '')

    if (field.isArray && field.type !== 'object') {
      // nested array of nonbjects generation
      return (
        <Card
          radius={cardRadius}
          key={schemaKey}
          mt={12}
          mb={12}
          withBorder
          css={css`
            width: 100%;
          `}
        >
          {/* existing array fields, if any */}
          {accordion ? (
            <FormAccordion schemaKey={schemaKey}>
              <NestedHeader formField={formField} schemaKey={schemaKey} itemName={itemName} />
              <Space p={10} />
              <ArrayChildren formField={formField} schemaKey={schemaKey} inputProps={inputProps} />
            </FormAccordion>
          ) : (
            <>
              <NestedHeader formField={formField} schemaKey={schemaKey} itemName={itemName} />
              <Space p={6} />
              <ArrayChildren formField={formField} schemaKey={schemaKey} inputProps={inputProps} />
            </>
          )}
        </Card>
      )
    }

    if (field.isArray && field.type === 'object') {
      // array of objects
      return (
        <Card radius={cardRadius} key={schemaKey} mt={12} mb={12} withBorder>
          {accordion ? (
            <FormAccordion schemaKey={schemaKey}>
              <NestedHeader formField={formField} schemaKey={schemaKey} itemName={itemName} />
              <ArrayOfObjectsChildren formField={formField} schemaKey={schemaKey} />
            </FormAccordion>
          ) : (
            <>
              <NestedHeader formField={formField} schemaKey={schemaKey} itemName={itemName} />
              <ArrayOfObjectsChildren formField={formField} schemaKey={schemaKey} />
            </>
          )}
        </Card>
      )
    }

    return (
      <Group
        key={schemaKey}
        align="center"
        css={css`
          width: 100%;
        `}
      >
        {field.type !== 'object' ? (
          <>
            <GeneratedInput
              schemaKey={schemaKey}
              formField={formField}
              props={{ input: inputProps, container: containerProps }}
            />
          </>
        ) : (
          <>{options.labels[schemaKey] && renderTitle(formField, options.labels[schemaKey])}</>
        )}
      </Group>
    )
  })

  return <>{children}</>
}

type FormAccordionProps = {
  schemaKey: SchemaKey
  children: React.ReactNode
}

function FormAccordion({ children, schemaKey }: FormAccordionProps): JSX.Element | null {
  const { formName, options, schemaFields } = useDynamicFormContext()

  const accordion = options.accordion?.[schemaKey]

  if (!accordion) return null

  const value = `${schemaKey}-accordion`

  return (
    <Accordion
      defaultValue={accordion.defaultOpen ? value : null}
      styles={{
        control: { padding: 0, maxHeight: '28px' },
        content: { paddingRight: 0, paddingLeft: 0 },
      }}
      {...containerProps}
    >
      <Accordion.Item value={value}>
        <Accordion.Control>{accordion.title ?? `${schemaKey}`}</Accordion.Control>
        <Accordion.Panel>{children}</Accordion.Panel>
      </Accordion.Item>
    </Accordion>
  )
}

type ArrayOfObjectsChildrenProps = {
  formField: FormField
  schemaKey: SchemaKey
}

function ArrayOfObjectsChildren({
  formField,

  schemaKey,
}: ArrayOfObjectsChildrenProps) {
  const { formName, options, schemaFields } = useDynamicFormContext()
  const form = useFormContext()
  // form.watch(formField, fieldArray.fields) // inf rerendering
  const theme = useMantineTheme()
  const itemName = singularize(options.labels[schemaKey] || '')

  useWatch({ name: `${formField}`, control: form.control }) // needed

  const children = (form.getValues(formField) || []).map((item, k: number) => {
    // input focus loss on rerender when defining component inside another function scope
    return (
      <div
        // reodering: https://codesandbox.io/s/watch-usewatch-calc-forked-5vrcsk?file=/src/fieldArray.js
        key={k}
        css={css`
          min-width: 100%;
        `}
      >
        <Card
          mt={12}
          mb={12}
          withBorder
          radius={cardRadius}
          bg={theme.colorScheme === 'dark' ? theme.colors.dark[7] : theme.colors.gray[2]}
        >
          <Flex justify={'end'}>
            <RemoveButton formField={formField} index={k} itemName={itemName} icon={<IconTrash size="1rem" />} />
          </Flex>
          <Group>
            <GeneratedInputs parentSchemaKey={schemaKey} parentFormField={`${formField}.${k}` as FormField} />
          </Group>
        </Card>
      </div>
    )
  })

  return (
    <Flex gap={6} align="center" direction="column">
      {children}
    </Flex>
  )
}

type ArrayChildrenProps = {
  formField: FormField
  schemaKey: SchemaKey
  inputProps: any
}

function ArrayChildren({ formField, schemaKey, inputProps }: ArrayChildrenProps) {
  const form = useFormContext()
  const theme = useMantineTheme()
  const { formName, options, schemaFields } = useDynamicFormContext()

  useWatch({ name: `${formField}`, control: form.control }) // needed

  const children = (form.getValues(formField) || []).map((item, k: number) => {
    // input focus loss on rerender when defining component inside another function scope
    return (
      <Flex
        // IMPORTANT: https://react-hook-form.com/docs/usefieldarray Does not support flat field array.
        // if reordering needed change spec to use object. since it's all generated its the easiest way to not mess up validation, etc.
        key={k}
        css={css`
          width: 100%;
        `}
      >
        <GeneratedInput
          schemaKey={schemaKey}
          formField={`${formField}.${k}` as FormField}
          props={{
            input: { ...inputProps, id: `${formName}-${formField}-${k}` },
            container: containerProps,
          }}
          index={k}
        />
      </Flex>
    )
  })

  if (children.length === 0) return null

  return (
    <Card
      radius={cardRadius}
      p={6}
      bg={theme.colorScheme === 'dark' ? theme.colors.dark[7] : theme.colors.gray[2]}
      withBorder
      css={css`
        width: 100%;
      `}
    >
      <Flex
        gap={6}
        align="center"
        direction="column"
        css={css`
          width: 100%;
        `}
      >
        {children}
      </Flex>
    </Card>
  )
}

function FormData() {
  const myFormData = useWatch()
  const myFormState = useFormState()

  // console.log(JSON.stringify(myFormData.base.items, null, 2))

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

type GeneratedInputProps = {
  schemaKey: SchemaKey
  props?: {
    input?: PropsWithChildren<any>
    container?: PropsWithChildren<any>
  }
  formField: FormField
  withRemoveButton?: boolean
  index?: number
}

const convertValueByType = (type: SchemaField['type'] | undefined, value) => {
  switch (type) {
    case 'date':
    case 'date-time':
      return value ? new Date(value) : undefined
    default:
      return value
  }
}

// useMemo with dep list of [JSON.stringify(_.get(form.values, formField)), ...] (will always rerender if its object, but if string only when it changes)
// TODO: just migrate to react-hook-form: https://codesandbox.io/s/dynamic-radio-example-forked-et0wi?file=/src/content/FirstFormSection.tsx
// for builtin support for uncontrolled input
const GeneratedInput = ({ schemaKey, props, formField, index }: GeneratedInputProps) => {
  const form = useFormContext()
  // useWatch({ control: form.control, name: formField }) // completely unnecessary, it's registered...
  const { formName, options, schemaFields } = useDynamicFormContext()

  const propsOverride = options.propsOverride?.[schemaKey]
  const type = schemaFields[schemaKey]?.type
  const itemName = singularize(options.labels[schemaKey] || '')

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

  const formFieldKeys = formField.split('.')
  // remove last index
  const formFieldArrayPath = formFieldKeys.slice(0, formFieldKeys.length - 1).join('.') as FormField

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
  const selectOptions = options.selectOptions?.[schemaKey]

  if (component) {
    el = React.cloneElement(component, {
      ..._props,
      ...component.props, // allow user override
      // TODO: this depends on component type, onChange should be customizable in options parameter with registerOnChange as fn param
      onChange: (e) => registerOnChange({ target: { name: formField, value: e } }),
    })
  } else if (selectOptions) {
    // TODO: on watch, should set current Select value by filtering
    // selectOptions.values.find((option) => selectOptions.formValueTransformer(option) === value)
    // else value lost on rerendering

    el = (
      <Select
        withinPortal
        label="Select user to update"
        itemComponent={itemComponentTemplate(selectOptions.optionTransformer)}
        data-test-subj="updateUserAuthForm__selectable"
        searchable
        filter={(option, item) => {
          if (option !== '') {
            return item.label?.toLowerCase().includes(option.toLowerCase().trim())
          }
          return (
            item.value.toLowerCase().includes(option.toLowerCase().trim()) ||
            (item.description && String(item.description)?.toLowerCase().includes(option.toLowerCase().trim()))
          )
        }}
        data={selectOptions.values.map((option) => ({
          label: selectOptions.labelTransformer(option),
          value: selectOptions.formValueTransformer(option),
          option,
        }))}
        onChange={(value) => {
          const option = selectOptions.values.find((option) => selectOptions.formValueTransformer(option) === value)
          console.log({ onChangeOption: option })
          if (!option) return
          registerOnChange({
            target: {
              name: formField,
              value: selectOptions.formValueTransformer(option),
            },
          })
        }}
        {..._props}
      />
    )
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
    <Flex align="center" gap={6} justify={'center'} {...props?.container}>
      {el}
      {index !== undefined && (
        <RemoveButton
          formField={formFieldArrayPath}
          index={index}
          itemName={itemName}
          icon={<IconMinus size="1rem" />}
        />
      )}
    </Flex>
  )
}

type RemoveButtonProps = {
  formField: FormField
  index: number
  itemName: string
  icon: React.ReactNode
  withFieldArray?: boolean
}

// needs to be own component to trigger rerender on delete, can't have conditional useWatch
const RemoveButton = ({ formField, index, itemName, icon, withFieldArray = false }: RemoveButtonProps) => {
  const form = useFormContext()
  const { colorScheme } = useMantineTheme()
  const { formName, options, schemaFields } = useDynamicFormContext()

  useWatch({ name: `${formField}`, control: form.control })

  const fieldArray = useFieldArray({
    control: form.control,
    name: formField,
  })

  return (
    <Tooltip withinPortal label={`Remove ${itemName}`} position="top-end" withArrow>
      <ActionIcon
        onClick={(e) => {
          // NOTE: don't use rhf useFieldArray, way too many edge cases for no gain. if reordering is needed, implement it manually.
          if (withFieldArray) {
            fieldArray.remove(index)
          } else {
            const listItems = form.getValues(formField)
            removeElementByIndex(listItems, index)
            form.unregister(formField) // needs to be before setValue
            form.setValue(formField, listItems as any)
          }
        }}
        // variant="filled"
        css={css`
          background-color: ${colorScheme === 'dark' ? '#7c1a1a' : '#b03434'};
          color: white;
          :hover {
            background-color: gray;
          }
        `}
        size="sm"
        id={`${formName}-${formField}-remove-button-${index}`}
      >
        {icon}
      </ActionIcon>
    </Tooltip>
  )
}

type NestedHeaderProps = {
  schemaKey: SchemaKey
  formField: FormField
  itemName: string
}

const NestedHeader = ({ formField, schemaKey, itemName }: NestedHeaderProps) => {
  const { formName, options, schemaFields } = useDynamicFormContext()

  const form = useFormContext()
  const accordion = options.accordion?.[schemaKey]

  return (
    <div>
      {/* {<Prism language="json">{JSON.stringify({ formField, parentFormField }, null, 4)}</Prism>} */}
      <Flex direction="row" align="center">
        {!accordion && options.labels[schemaKey] && renderTitle(formField, options.labels[schemaKey])}
        <Button
          size="xs"
          p={4}
          leftIcon={<IconPlus size="1rem" />}
          onClick={() => {
            const initialValue = initialValueByType(schemaFields[schemaKey]?.type)
            const vals = form.getValues(formField) || []
            console.log([...vals, initialValue] as any)

            form.setValue(formField, [...vals, initialValue] as any)
          }}
          variant="filled"
          color={'green'}
          id={`${formName}-${formField}-add-button`}
        >{`Add ${itemName}`}</Button>
      </Flex>
    </div>
  )
}

const initialValueByType = (type?: SchemaField['type']) => {
  switch (type) {
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
      console.log(`unknown type: ${type}`)
      return ''
  }
}
