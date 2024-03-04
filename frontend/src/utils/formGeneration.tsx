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
  Input,
  Code,
  MultiSelect,
  CloseButton,
  useMantineColorScheme,
  getThemeColor,
  Combobox,
  useCombobox,
  InputBase,
  PillsInput,
  Pill,
  ScrollArea,
} from '@mantine/core'
import classes from './form.module.css'
import { DateInput, DateTimePicker } from '@mantine/dates'
import { useFocusWithin } from '@mantine/hooks'
import { CodeHighlight } from '@mantine/code-highlight'
import { rem, useMantineTheme } from '@mantine/core'
import { Icon123, IconMinus, IconPlus, IconTrash } from '@tabler/icons'
import { pluralize, singularize } from 'inflection'
import _, { concat, flatten, isArray, lowerCase, lowerFirst, memoize, upperFirst } from 'lodash'
import { Virtuoso } from 'react-virtuoso'

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
  useRef,
  useEffect,
  ReactNode,
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
  type UseFormRegisterReturn,
  type ChangeHandler,
  type UseFormSetError,
} from 'react-hook-form'
import { json } from 'react-router-dom'
import { ApiError } from 'src/api/mutator'
import DynamicFormErrorCallout from 'src/components/Callout/DynamicFormErrorCallout'
import PageTemplate from 'src/components/PageTemplate'
import useRenders from 'src/hooks/utils/useRenders'
import type {
  DeepPartial,
  GenericObject,
  GetKeys,
  RecursiveKeyOf,
  RecursiveKeyOfArray,
  PathType,
  Branded,
  UniqueArray,
  Callable,
} from 'src/types/utils'
import { removeElementByIndex } from 'src/utils/array'
import { getContrastYIQ } from 'src/utils/colors'
import type { SchemaField } from 'src/utils/jsonSchema'
import { entries, hasNonEmptyValue, isObject, keys } from 'src/utils/object'
import { nameInitials, sentenceCase } from 'src/utils/strings'
import { useFormSlice } from 'src/slices/form'
import RandExp, { randexp } from 'randexp'
import type { FormField, SchemaKey } from 'src/utils/form'
import { useCalloutErrors } from 'src/components/Callout/useCalloutErrors'
import { inputBuilder, selectOptionsBuilder, useDynamicFormContext } from 'src/utils/formGeneration.context'

export type SelectOptionsTypes = 'select' | 'multiselect'

export type SelectOptions<Return, E = unknown> = {
  type: SelectOptionsTypes
  values: E[]
  formValueTransformer: <V extends E>(el: V & E) => Return extends unknown[] ? Return[number] : Return
  /** Modify search behavior, e.g. matching against `${el.<field_1>} ${el.<field_2>} ${el.field_3}`.
   * It searches in the whole stringified object by default.
   */
  searchValueTransformer?: <V extends E>(el: V & E) => string
  /**
   * Overrides combobox option components.
   */
  optionTransformer: <V extends E>(el: V & E) => JSX.Element
  /**
   * Overrides combobox selected item pill components.
   */
  pillTransformer?: <V extends E>(el: V & E) => JSX.Element
  /**
   * Overrides default combobox item label color.
   */
  labelColor?: <V extends E>(el: V & E) => string
}

export interface InputOptions<Return, E = unknown> {
  component: JSX.Element
  propsFn?: (registerOnChange: ChangeHandler) => React.ComponentProps<'input'>
}

const comboboxOptionTemplate = (transformer: (...args: any[]) => JSX.Element, option) => {
  return <Box m={2}>{transformer(option)}</Box>
}

export type DynamicFormOptions<
  T extends object,
  IgnoredFormKeys extends U | null,
  U extends PropertyKey = GetKeys<T>,
> = {
  /**
   * Label mapping for fields. Use null to skip rendering field entirely.
   */
  labels: {
    [key in Exclude<U, IgnoredFormKeys>]: string | null
  }
  renderOrderPriority?: Array<Exclude<keyof T, IgnoredFormKeys>>
  // used to populate form inputs if the form field is empty. Applies to all nested fields.
  defaultValues?: Partial<{
    [key in Exclude<U, IgnoredFormKeys>]: DeepPartial<
      PathType<
        T,
        // can fix key constraint error with U extends RecursiveKeyOf<T, ''> but not worth it due to cpu usage, just ignore
        //@ts-ignore
        key
      >
    >
  }>
  selectOptions?: Partial<{
    [key in Exclude<U, IgnoredFormKeys>]: ReturnType<
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
    [key in Exclude<U, IgnoredFormKeys>]: ReturnType<
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
    [key in Exclude<U, IgnoredFormKeys>]: {
      description?: string
      disabled?: boolean
    }
  }>
  accordion?: Partial<{
    [key in Exclude<U, IgnoredFormKeys>]: {
      defaultOpen?: boolean
      title?: JSX.Element
    }
  }>
}

export type DynamicFormContextValue = {
  formName: string
  schemaFields: Record<SchemaKey, SchemaField>
  options: DynamicFormOptions<any, null, SchemaKey> // for more performant internal intellisense. for user it will be typed
}

export const DynamicFormContext = createContext<DynamicFormContextValue | undefined>(undefined)

type DynamicFormProviderProps = {
  value: DynamicFormContextValue
  children: React.ReactNode
}

const DynamicFormProvider = ({ value, children }: DynamicFormProviderProps) => {
  return <DynamicFormContext.Provider value={value}>{children}</DynamicFormContext.Provider>
}

type DynamicFormProps<T extends object, IgnoredFormKeys extends GetKeys<T> | null = null> = {
  schemaFields: Record<Exclude<GetKeys<T>, IgnoredFormKeys>, SchemaField>
  options: DynamicFormOptions<T, IgnoredFormKeys, GetKeys<T>>
  formName: string
  onSubmit: React.FormEventHandler<HTMLFormElement>
}

function renderTitle(key: FormField, formName: string, title: ReactNode) {
  return (
    <>
      <Title data-testid={`${formName}-${key}-title`} size={18}>
        {title}
      </Title>
      <Space p={8} />
    </>
  )
}

const cardRadius = 6

/**
 *
 * NOTE: arrays of arrays not supported.
 */
export default function DynamicForm<Form extends object, IgnoredFormKeys extends GetKeys<Form> | null = null>({
  formName,
  schemaFields,
  options,
  onSubmit,
}: DynamicFormProps<Form, IgnoredFormKeys>) {
  const theme = useMantineTheme()
  const form = useFormContext()
  const formSlice = useFormSlice()
  const { extractCalloutErrors, setCalloutErrors, calloutErrors, extractCalloutTitle } = useCalloutErrors(formName)
  console.log({ formboolis: form.getValues('demoProject.reopened') })
  let _schemaFields: DynamicFormContextValue['schemaFields'] = schemaFields
  if (options.renderOrderPriority) {
    const _schemaKeys: SchemaKey[] = []
    _.uniq(options.renderOrderPriority).forEach((k, i) => {
      entries(schemaFields).forEach(([sk, v], i) => {
        if (!String(sk).startsWith(String(k))) return
        _schemaKeys.push(sk as SchemaKey)
      })
    })
    entries(schemaFields).forEach(([k, v], i) => {
      if (_schemaKeys.includes(k as SchemaKey)) return
      _schemaKeys.push(k as SchemaKey)
    })

    _schemaFields = _schemaKeys.reduce((acc, key) => {
      acc[key] = schemaFields[key]
      return acc
    }, {})
  }

  return (
    <DynamicFormProvider value={{ formName, options, schemaFields: _schemaFields }}>
      <>
        <FormData />
        <DynamicFormErrorCallout />
        <form
          onSubmit={onSubmit}
          css={css`
            min-width: 100%;
          `}
          data-testid={formName}
        >
          <Button type="submit">Submit</Button>
          <GeneratedInputs />
        </form>
      </>
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
    const itemName = singularize(options.labels[schemaKey] || '')

    const inputProps = {
      css: css`
        width: 100%;
      `,
      ...(!field.isArray && { label: options.labels[schemaKey] }),
      required: field.required,
      'data-testid': `${formName}-${formField}`,
      onKeyPress: (e: React.KeyboardEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        if (e.key !== 'Enter') {
          return
        }
        if (document.activeElement?.tagName.toLowerCase() === 'textarea') {
          console.log('Enter key pressed in textarea')
        } else {
          e.preventDefault()
        }
      },
    }

    if (field.isArray && field.type !== 'object') {
      // nested array of nonbjects generation
      return accordion ? (
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
          {
            <FormAccordion schemaKey={schemaKey}>
              <NestedHeader formField={formField} schemaKey={schemaKey} itemName={itemName} />
              <Space p={10} />
              <ArrayChildren formField={formField} schemaKey={schemaKey} inputProps={inputProps} />
            </FormAccordion>
          }
        </Card>
      ) : (
        <>
          {/*
          FIXME: NestedHeader should only render if schemaKey field is not a multiselect, else we add and remove via buttons
          which does need header for title and `+ add $item`
           */}{' '}
          <NestedHeader formField={formField} schemaKey={schemaKey} itemName={itemName} />
          <ArrayChildren formField={formField} schemaKey={schemaKey} inputProps={inputProps} />
        </>
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
          <>{options.labels[schemaKey] && renderTitle(formField, formName, options.labels[schemaKey])}</>
        )}
      </Group>
    )
  })

  const renderCount = useRenders()

  return (
    <>
      {/* <Code c={'red'}>Renders: {renderCount}</Code> */}
      {children}
    </>
  )
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
  const { colorScheme } = useMantineColorScheme()
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
        <Card mt={12} mb={12} withBorder radius={cardRadius} className={classes.childCard}>
          <Flex justify={'end'} mb={10}>
            <RemoveButton formField={formField} index={k} itemName={itemName} icon={<IconTrash size="1rem" />} />
          </Flex>
          <Group gap={10}>
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
  const itemName = singularize(options.labels[schemaKey] || '')

  useWatch({ name: `${formField}`, control: form.control }) // needed

  if (options.selectOptions?.[schemaKey]?.type === 'multiselect') {
    return (
      <Flex
        css={css`
          width: 100%;
        `}
      >
        <GeneratedInput
          schemaKey={schemaKey}
          formField={formField as FormField}
          props={{
            input: {
              ...inputProps,
              'data-testid': `${formName}-${formField}`,
            },
            container: {
              ...containerProps,
            },
          }}
        />
      </Flex>
    )
  }

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
            input: { ...inputProps, 'data-testid': `${formName}-${formField}-${k}` },
            container: containerProps,
          }}
          index={k}
        />
      </Flex>
    )
  })

  if (children.length === 0) return null

  return (
    <Flex
      gap={6}
      align="left"
      direction="column"
      css={css`
        width: 100%;
      `}
    >
      <Card radius={cardRadius} p={6} withBorder className={classes.arrayChildCard}>
        {children}
      </Card>
    </Flex>
  )
}

function FormData() {
  const myFormData = useWatch()
  const form = useFormContext()
  const myFormState = useFormState({ control: form.control })

  console.log(`form has errors: ${hasNonEmptyValue(myFormState.errors)}`)

  let code = ''
  try {
    code = JSON.stringify(myFormData, null, 2)
  } catch (error) {
    console.error(error)
  }
  return (
    <Accordion>
      <Accordion.Item value="form">
        <Accordion.Control>{`See form`}</Accordion.Control>
        <Accordion.Panel>
          <CodeHighlight language="json" code={code}></CodeHighlight>
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

const GeneratedInput = ({ schemaKey, props, formField, index }: GeneratedInputProps) => {
  const form = useFormContext()
  const theme = useMantineTheme()
  // useWatch({ control: form.control, name: formField }) // completely unnecessary, it's registered...
  const { formName, options, schemaFields } = useDynamicFormContext()

  const [isSelectVisible, setIsSelectVisible] = useState(false)

  const propsOverride = options.propsOverride?.[schemaKey]
  const type = schemaFields[schemaKey]?.type
  const itemName = singularize(options.labels[schemaKey] || '')

  // registerOnChange's type refers to onChange event's type
  const { onChange: registerOnChange, ...registerProps } = form.register(formField, {
    // IMPORTANT: this is the type set in registerOnChange!
    ...(type === 'date' || type === 'date-time'
      ? // TODO: use convertValueByType
        { valueAsDate: true, setValueAs: (v) => (v === '' ? undefined : new Date(v)) }
      : type === 'integer'
      ? { valueAsNumber: true, setValueAs: (v) => (v === '' ? undefined : parseInt(v, 10)) }
      : type === 'number'
      ? { valueAsNumber: true, setValueAs: (v) => (v === '' ? undefined : parseFloat(v)) }
      : type === 'boolean'
      ? {
          setValueAs: (v) => {
            console.log({ settingValueBoolean: v })
            return v === '' ? false : v === 'true'
          },
        }
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

  const _propsWithoutRegister = {
    ...props?.input,
    ...propsOverride,
    ...(!fieldState.isDirty && { defaultValue: convertValueByType(type, formValue) }),
    ...(fieldState.error && { error: sentenceCase(fieldState.error?.message) }),
    required: schemaFields[schemaKey]?.required && type !== 'boolean',
    placeholder: `Enter ${lowerFirst(singularize(options.labels[schemaKey] || ''))}`,
  }

  const _props = {
    ...registerProps,
    ..._propsWithoutRegister,
  }

  let formFieldComponent: JSX.Element | null = null
  const component = options.input?.[schemaKey]?.component
  // TODO: componentPropsFn must return {}
  const componentPropsFn = options.input?.[schemaKey]?.propsFn
  const selectOptions = options.selectOptions?.[schemaKey]
  const selectRef = useRef<HTMLInputElement | null>(null)
  const [customElMinHeight, setCustomElMinHeight] = useState(34.5)

  const { ref: focusRef, focused: selectFocused } = useFocusWithin()

  useEffect(() => {
    if (isSelectVisible) {
      setCustomElMinHeight(selectRef.current?.clientHeight ?? 34.5)
      selectRef.current?.focus()
    }
  }, [isSelectVisible])

  useEffect(() => {
    if (!selectFocused) setIsSelectVisible(false)
  }, [selectFocused])

  if (component) {
    // explicit component given
    formFieldComponent = React.cloneElement(component, {
      ..._props,
      // IMPORTANT: some mantine components require e, others e.target.value
      onChange: (e) => registerOnChange({ target: { name: formField, value: e.target?.value ?? e } }),
      ...component.props, // allow user override
      ...(componentPropsFn && componentPropsFn(registerOnChange)), // allow user override
    })
  } else if (selectOptions) {
    const props = {
      ..._propsWithoutRegister,
      ...(componentPropsFn && componentPropsFn(registerOnChange)), // allow user override
    }

    switch (selectOptions.type) {
      case 'select':
        formFieldComponent = (
          <CustomSelect
            formField={formField}
            registerOnChange={registerOnChange}
            schemaKey={schemaKey}
            itemName={itemName}
            {...props}
          />
        )
        break
      case 'multiselect':
        formFieldComponent = (
          <CustomMultiselect
            formField={formField}
            registerOnChange={registerOnChange}
            schemaKey={schemaKey}
            itemName={itemName}
            {...props}
          />
        )

        break
      default:
        break
    }
  } else {
    switch (schemaFields[schemaKey]?.type) {
      case 'string':
        formFieldComponent = (
          <TextInput
            onChange={(e) => registerOnChange({ target: { name: formField, value: e.target.value } })}
            {..._props}
          />
        )
        break
      case 'boolean':
        // cannot set both defaultValue and have defaultValues in form for checkboxes. see discussions in gh
        const { defaultValue: ____, ...checkboxProps } = _props
        formFieldComponent = <Checkbox pt={10} pb={4} {...{ ...checkboxProps, ...form.register(formField) }} />
        break
      case 'date':
        formFieldComponent = (
          <DateInput
            valueFormat="DD/MM/YYYY"
            onChange={(e) =>
              registerOnChange({
                target: { name: formField, value: e },
              })
            }
            {..._props}
            placeholder="Select date"
          />
        )
        break
      case 'date-time':
        formFieldComponent = (
          <DateTimePicker
            onChange={(e) =>
              registerOnChange({
                target: { name: formField, value: e },
              })
            }
            {..._props}
            placeholder="Select date and time"
          />
        )
        break
      case 'integer':
        formFieldComponent = (
          <NumberInput
            onChange={(e) =>
              registerOnChange({
                target: {
                  name: formField,
                  value: Number(e), // else broken validation onChange
                },
              })
            }
            {..._props}
          />
        )
        break
      case 'number':
        formFieldComponent = (
          <NumberInput
            onChange={(e) =>
              registerOnChange({
                target: {
                  name: formField,
                  value: Number(e), // else broken validation onChange
                },
              })
            }
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
      {formFieldComponent}
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
}

// needs to be own component to trigger rerender on delete, can't have conditional useWatch
const RemoveButton = ({ formField, index, itemName, icon }: RemoveButtonProps) => {
  const form = useFormContext()
  const { colorScheme } = useMantineColorScheme()
  const { formName, options, schemaFields } = useDynamicFormContext()

  return (
    <Tooltip withinPortal label={`Remove ${lowerFirst(itemName)}`} position="top-end" withArrow>
      <ActionIcon
        onClick={(e) => {
          // NOTE: don't use rhf useFieldArray, way too many edge cases for no gain. if reordering is needed, implement it manually.
          // we could even implement flat array reordering by handling them internally as objects with id prop
          const listItems = form.getValues(formField)
          removeElementByIndex(listItems, index)
          form.unregister(formField) // needs to be called before setValue
          form.setValue(formField, listItems as any)
        }}
        color={'#bd3535'}
        size="sm"
        data-testid={`${formName}-${formField}-remove-button-${index}`}
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

  return options.selectOptions?.[schemaKey]?.type !== 'multiselect' ? (
    <div>
      <Flex direction="row" align="center">
        {!accordion && options.labels[schemaKey] && renderTitle(formField, formName, options.labels[schemaKey])}
        {
          <Button
            size="xs"
            p={4}
            leftSection={<IconPlus size="1rem" />}
            onClick={() => {
              const initialValue = initialValueByType(schemaFields[schemaKey]?.type)
              const vals = form.getValues(formField) || []
              console.log([...vals, initialValue] as any)

              form.unregister(formField) // needs to be called before setValue
              form.setValue(formField, [...vals, initialValue] as any)
            }}
            variant="filled"
            color={'green'}
            data-testid={`${formName}-${formField}-add-button`}
          >{`Add ${lowerFirst(itemName)}`}</Button>
        }
      </Flex>
    </div>
  ) : null
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

type CustomPillProps = {
  value: any
  schemaKey: SchemaKey
  handleValueRemove: (val: string) => void
  props?: React.HTMLProps<HTMLDivElement>
}

type CustomMultiselectProps = {
  formField: FormField
  registerOnChange: ChangeHandler
  schemaKey: SchemaKey
  itemName: string
}

function CustomMultiselect({
  formField,
  registerOnChange,
  schemaKey,
  itemName,
  ...inputProps
}: CustomMultiselectProps) {
  const form = useFormContext()
  const { formName, options, schemaFields } = useDynamicFormContext()

  const selectOptions = options.selectOptions![schemaKey]!
  const formValues = (form.getValues(formField) as any[]) || []

  const [search, setSearch] = useState('')

  const combobox = useCombobox({
    onDropdownClose: () => combobox.resetSelectedOption(),
    onDropdownOpen: () => combobox.updateSelectedOptionIndex('active'),
  })

  const handleValueRemove = (val: string) => {
    console.log({ val, formValues, a: formValues.filter((v) => v !== val) })
    formSlice.setCustomError(formName, formField, null) // index position changed, misleading message
    form.unregister(formField) // needs to be called before setValue
    form.setValue(
      formField,
      formValues.filter((v) => v !== val),
    )
  }
  const comboboxOptions = selectOptions.values
    .filter((item: any) => {
      const inSearch = JSON.stringify(
        selectOptions.searchValueTransformer ? selectOptions.searchValueTransformer(item) : item,
      )
        .toLowerCase()
        .includes(search.toLowerCase().trim())

      const notSelected = !formValues.includes(selectOptions.formValueTransformer(item))

      return inSearch && notSelected
    })
    .map((option) => {
      const value = String(selectOptions.formValueTransformer(option))
      return (
        <Combobox.Option value={value} key={value} active={formValues.includes(value)}>
          <Group align="stretch" justify="space-between">
            {selectOptions.optionTransformer(option)}
            <CloseButton
              onMouseDown={() => handleValueRemove(selectOptions.formValueTransformer(option))}
              variant="transparent"
              color="gray"
              size={22}
              iconSize={14}
              tabIndex={-1}
            />
          </Group>
        </Combobox.Option>
      )
    })

  const formState = useFormState({ control: form.control })

  const formSlice = useFormSlice()
  const multiselectFirstError = formSlice.form[formName]?.customErrors[formField]

  useEffect(() => {
    const formFieldErrors = _.get(formState.errors, formField)
    if (isArray(formFieldErrors) && !multiselectFirstError) {
      formFieldErrors.forEach((error, index) => {
        if (!!error) {
          const message = `${itemName} number ${index + 1} ${error.message}`
          formSlice.setCustomError(formName, formField, message)
        }
      })
    }
  }, [formState])

  return (
    <Box w={'100%'}>
      <Combobox
        store={combobox}
        onOptionSubmit={(value, props) => {
          const option = selectOptions.values.find(
            (option) => String(selectOptions.formValueTransformer(option)) === value,
          )
          formSlice.setCustomError(formName, formField, null)
          registerOnChange({
            target: {
              name: formField,
              value: [...formValues, selectOptions.formValueTransformer(option)],
            },
          })
        }}
        withinPortal
      >
        <Combobox.DropdownTarget>
          <PillsInput
            styles={{
              error: {},
            }}
            label={pluralize(upperFirst(itemName))}
            onClick={() => combobox.openDropdown()}
            {...inputProps}
            // must override input props error
            error={multiselectFirstError}
          >
            <Pill.Group>
              {formValues.length > 0 &&
                formValues.map((formValue, i) => (
                  <CustomPill
                    key={`${formField}-${i}-pill`}
                    value={formValue}
                    handleValueRemove={handleValueRemove}
                    schemaKey={schemaKey}
                  />
                ))}

              <Combobox.EventsTarget>
                <PillsInput.Field
                  placeholder={`Search ${pluralize(lowerFirst(itemName))}`}
                  onChange={(event) => {
                    combobox.updateSelectedOptionIndex()
                    setSearch(event.currentTarget.value)
                  }}
                  value={search}
                  onFocus={() => combobox.openDropdown()}
                  onBlur={() => combobox.closeDropdown()}
                  onKeyDown={(event) => {
                    if (event.key === 'Backspace' && search.length === 0) {
                      event.preventDefault()
                      formSlice.setCustomError(formName, formField, null)
                      form.unregister(formField) // needs to be called before setValue
                      form.setValue(formField, formValues)
                    }
                  }}
                />
              </Combobox.EventsTarget>
            </Pill.Group>
          </PillsInput>
        </Combobox.DropdownTarget>

        <Combobox.Dropdown>
          <Combobox.Options
            mah={200} // scrollable
            style={{ overflowY: 'auto' }}
          >
            <ScrollArea.Autosize mah={200} type="scroll">
              <Virtuoso
                style={{ height: '200px' }} // match height with autosize
                totalCount={comboboxOptions.length}
                itemContent={(index) => comboboxOptions[index]}
              />
            </ScrollArea.Autosize>
          </Combobox.Options>
        </Combobox.Dropdown>
      </Combobox>
    </Box>
  )
}

type CustomSelectProps = {
  formField: FormField
  registerOnChange: ChangeHandler
  schemaKey: SchemaKey
  itemName: string
}

function CustomSelect({ formField, registerOnChange, schemaKey, itemName, ...inputProps }: CustomSelectProps) {
  const form = useFormContext()
  const { formName, options, schemaFields } = useDynamicFormContext()
  const formSlice = useFormSlice()

  const selectOptions = options.selectOptions![schemaKey]!
  const formValues = (form.getValues(formField) as any[]) || []

  const [search, setSearch] = useState('')

  const combobox = useCombobox({
    onDropdownClose: () => {
      combobox.resetSelectedOption()
      combobox.focusTarget()
      setSearch('')
    },

    onDropdownOpen: () => {
      combobox.focusSearchInput()
    },
  })

  const selectedOption = selectOptions.values.find((option) => {
    return selectOptions.formValueTransformer(option) === form.getValues(formField)
  })

  const comboboxOptions = selectOptions.values
    .filter((item: any) =>
      JSON.stringify(selectOptions.searchValueTransformer ? selectOptions.searchValueTransformer(item) : item)
        .toLowerCase()
        .includes(search.toLowerCase().trim()),
    )
    .map((option) => {
      const value = String(selectOptions.formValueTransformer(option))

      return (
        <Combobox.Option value={value} key={value}>
          {comboboxOptionTemplate(selectOptions.optionTransformer, option)}
        </Combobox.Option>
      )
    })
  const { extractCalloutErrors, setCalloutErrors, calloutErrors, extractCalloutTitle } = useCalloutErrors(formName)

  const parentSchemaKey = schemaKey.split('.').slice(0, -1).join('.') as SchemaKey

  return (
    <Box w={'100%'}>
      <Combobox
        store={combobox}
        withinPortal={true}
        position="bottom-start"
        withArrow
        onOptionSubmit={async (value) => {
          const option = selectOptions.values.find(
            (option) => String(selectOptions.formValueTransformer(option)) === value,
          )
          console.log({ onChangeOption: option })
          if (!option) {
            // in form gen we do want to concatenate errors, to be shown upon submit clicked.
            // react hook form will show input errors as well on registered components
            formSlice.setCustomError(formName, formField, `${value} is not a valid ${itemName}`)
            return
          }
          await registerOnChange({
            target: {
              name: formField,
              value: selectOptions.formValueTransformer(option),
            },
          })
          combobox.closeDropdown()
        }}
      >
        <Combobox.Target withAriaAttributes={false}>
          <InputBase
            label={!schemaFields[parentSchemaKey]?.isArray ? singularize(upperFirst(itemName)) : null}
            className={classes.select}
            component="button"
            type="button"
            pointer
            rightSection={<Combobox.Chevron />}
            onClick={() => combobox.toggleDropdown()}
            rightSectionPointerEvents="none"
            multiline
            {...inputProps}
          >
            {selectedOption ? (
              comboboxOptionTemplate(selectOptions.optionTransformer, selectedOption)
            ) : (
              <Input.Placeholder>{`Pick ${singularize(lowerFirst(itemName))}`}</Input.Placeholder>
            )}
          </InputBase>
        </Combobox.Target>

        <Combobox.Dropdown>
          <Combobox.Search
            miw={'100%'}
            value={search}
            onChange={(event) => setSearch(event.currentTarget.value)}
            placeholder={`Search ${lowerFirst(itemName)}`}
          />
          <Combobox.Options
            mah={200} // scrollable
            style={{ overflowY: 'auto' }}
          >
            {comboboxOptions.length > 0 ? (
              <ScrollArea.Autosize mah={200} type="scroll">
                <Virtuoso
                  style={{ height: '200px' }} // match height with autosize
                  totalCount={comboboxOptions.length}
                  itemContent={(index) => comboboxOptions[index]}
                />
              </ScrollArea.Autosize>
            ) : (
              <Combobox.Empty>Nothing found</Combobox.Empty>
            )}
          </Combobox.Options>
        </Combobox.Dropdown>
      </Combobox>
    </Box>
  )
}

function CustomPill({ value, schemaKey, handleValueRemove, ...props }: CustomPillProps): JSX.Element | null {
  const { formName, options, schemaFields } = useDynamicFormContext()
  const { extractCalloutErrors, setCalloutErrors, calloutErrors, extractCalloutTitle } = useCalloutErrors(formName)
  const selectOptions = options.selectOptions![schemaKey]!
  const itemName = singularize(options.labels[schemaKey] || '')

  let invalidValue = null

  const option = selectOptions.values.find((option) => selectOptions.formValueTransformer(option) === value)
  if (!option) {
    console.log(`${value} is not a valid ${singularize(lowerCase(itemName))}`)

    // explicitly set wrong values so that error positions make sense and the user knows there is a wrong form value beforehand
    // instead of us deleting it implicitly
    invalidValue = value
  }

  let color = '#bbbbbb' // for invalid values
  if (selectOptions?.labelColor && !invalidValue) {
    color = selectOptions?.labelColor(option)
  }

  const transformer = selectOptions.pillTransformer ? selectOptions.pillTransformer : selectOptions.optionTransformer

  return (
    <Box
      className={classes.valueComponentOuterBox}
      css={css`
        background-color: ${color};
        * {
          color: ${getContrastYIQ(color) === 'black' ? 'whitesmoke' : '#131313'};
        }
      `}
      {...props}
    >
      <Box className={classes.valueComponentInnerBox}>{invalidValue || transformer(option)}</Box>
      <CloseButton
        onMouseDown={() => handleValueRemove(value)}
        variant="transparent"
        size={22}
        iconSize={14}
        tabIndex={-1}
      />
    </Box>
  )
}
