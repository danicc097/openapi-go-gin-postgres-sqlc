/* eslint-disable @typescript-eslint/ban-ts-comment */
import { css } from '@emotion/react'
import type { EmotionJSX, ReactJSXElement } from '@emotion/react/types/jsx-namespace'
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
  CheckIcon,
  ComboboxProps,
  InputBaseProps,
  List,
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
  ReactElement,
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
import {
  fieldOptionsBuilder,
  inputBuilder,
  selectOptionsBuilder,
  useDynamicFormContext,
} from 'src/utils/formGeneration.context'
import { useCalloutErrors } from 'src/components/Callout/useCalloutErrors'
import { IconAlertCircle } from '@tabler/icons-react'
import useStopInfiniteRenders from 'src/hooks/utils/useStopInfiniteRenders'
import { joinWithAnd } from 'src/utils/format'

export type SelectOptionsTypes = 'select' | 'multiselect'

export type SelectOptions<Return, E = unknown> = {
  type: SelectOptionsTypes
  values: E[]
  formValueTransformer: <V extends E>(el: V & E) => Return extends unknown[] ? Return[number] : Return
  ariaLabelTransformer?: <V extends E>(el: V & E) => string
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

export type FieldOptions<Return, E = unknown> = {
  /**
   * Returns custom warnings based on the current form value. For arrays, the warning function
   * is executed for each element in form.
   * IMPORTANT: If the warning is called on a parent object whose children
   * may be arrays, initial child values are currently not initialized, ie
   * array entries will be undefined which may lead to runtime errors if
   * the entry is not nullable
   */
  warningFn?: (el: Return extends unknown[] ? Return[number] : Return) => string[]
}

export interface InputOptions<Return, E = unknown> {
  component: JSX.Element
  propsFn?: (registerOnChange: ChangeHandler) => React.ComponentProps<'input'>
}

const comboboxOptionTemplate = (transformer: (...args: any[]) => JSX.Element, option) => {
  return <Box m={2}>{transformer(option)}</Box>
}

interface Props extends React.HTMLProps<HTMLInputElement> {
  [key: string]: any
}

export type DynamicFormOptions<
  T extends object,
  IgnoredFormKeys extends U | null,
  U extends PropertyKey = GetKeys<T>,
> = {
  /**
   * Label mapping for fields.
   * To ignore fields, add leaf keys to IgnoredFormKeys.
   * Ignoring intermediate object/array keys in IgnoredFormKeys just skips rendering its titles.
   */
  labels: {
    [key in Exclude<U, IgnoredFormKeys>]: string
  }
  renderOrderPriority?: Array<Exclude<keyof T, IgnoredFormKeys>>
  /**
   * Used to populate form inputs if the form field is empty. Applies to all nested fields, including array elements
   *  */
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
  /**
   * Builds the given form fields as select or multiselect.
   */
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
   * Returns custom warnings based on the current form value
   */
  fieldOptions?: Partial<{
    [key in Exclude<U, IgnoredFormKeys>]: ReturnType<
      typeof fieldOptionsBuilder<
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
    [key in Exclude<U, IgnoredFormKeys>]: Props
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
  const { setHasClickedSubmit } = useCalloutErrors(formName)

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

  // useStopInfiniteRenders(20)

  return (
    <DynamicFormProvider value={{ formName, options, schemaFields: _schemaFields }}>
      <>
        {/** TODO: usedebounce for all fields  */}
        <FormData />
        {/* TODO: if not visible (large forms), should show a popup arrow on viewport bottom left "Go to errors" to focus on callout */}
        <DynamicFormErrorCallout />

        {/* TODO: can have undo-redo functionality by using RHF reset(...)*/}
        <form
          onSubmit={(e) => {
            setHasClickedSubmit(true)
            onSubmit(e)
          }}
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

function ArrayOfObjectsChildren({ formField, schemaKey }: ArrayOfObjectsChildrenProps) {
  const { formName, options, schemaFields } = useDynamicFormContext()
  const form = useFormContext()
  // form.watch(formField, fieldArray.fields) // inf rerendering
  const theme = useMantineTheme()
  const { colorScheme } = useMantineColorScheme()
  const itemName = singularize(options.labels[schemaKey] || '')

  useWatch({ name: `${formField}`, control: form.control }) // needed

  return (
    <Flex gap={6} align="center" direction="column">
      {(form.getValues(formField) || []).map((item, index) => (
        <ArrayOfObjectsChild
          key={index}
          index={index}
          formField={formField}
          itemName={itemName}
          schemaKey={schemaKey}
        />
      ))}
    </Flex>
  )
}

const ArrayOfObjectsChild = ({ index, formField, itemName, schemaKey }) => {
  const { formName, options, schemaFields } = useDynamicFormContext()
  const form = useFormContext()

  const itemFormField = `${formField}.${index}` as FormField
  const fieldWarnings = useFormSlice((state) => state.form[formName]?.customWarnings[itemFormField])
  const warningFn = options.fieldOptions?.[schemaKey]?.warningFn
  const formFieldWatch = form.watch(itemFormField)
  // TODO: dynamic watching of all nested elements (for warning function recompute only)
  // warningFn does rerender for
  // const formArrayElementWatch = form.watch(`${itemFormField}.items.0`)
  const formSlice = useFormSlice()

  useEffect(() => {
    // FIXME: warnings are not recalculated unless the array is modified.
    // should update useffect deps watching nested values of current array element without triggering inf rerendering
    if (warningFn) {
      const warnings = joinWithAnd(warningFn(formFieldWatch))
      console.log({ warnings, fieldWarnings, formFieldWatch })

      formSlice.setCustomWarning(formName, itemFormField, warnings.length > 0 ? warnings : null)
      form.trigger(itemFormField)
    }
  }, [formFieldWatch, fieldWarnings])

  return (
    <div
      key={index}
      css={css`
        min-width: 100%;
      `}
    >
      <Card mt={12} mb={12} withBorder radius={cardRadius} className={classes.childCard}>
        <Flex justify={'space-between'} direction={'row-reverse'} mb={10}>
          <RemoveButton formField={formField} index={index} itemName={itemName} icon={<IconTrash size="1rem" />} />
          {renderWarningIcon(fieldWarnings)}
        </Flex>
        <Group gap={10}>
          <GeneratedInputs parentSchemaKey={schemaKey} parentFormField={itemFormField} />
        </Group>
      </Card>
    </div>
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

  if (children.length === 0) return <Text size="xs" w={'100%'}>{`No ${lowerFirst(pluralize(itemName))} defined`}</Text>

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

  const hasErrors = hasNonEmptyValue(myFormState.errors)
  if (hasErrors) console.error('form has errors')

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
  const { formName, options, schemaFields } = useDynamicFormContext()

  const [isSelectVisible, setIsSelectVisible] = useState(false)

  const propsOverride = options.propsOverride?.[schemaKey]
  const type = schemaFields[schemaKey]?.type
  const itemName = singularize(options.labels[schemaKey] || '')

  // registerOnChange's type refers to onChange event's type
  const { onChange: registerOnChange, ...registerProps } = form.register(formField, {
    ...(type === 'date' || type === 'date-time'
      ? // TODO: use convertValueByType
        { valueAsDate: true, setValueAs: (v) => (v === '' ? undefined : new Date(v)) }
      : type === 'integer'
      ? { valueAsNumber: true, setValueAs: (v) => (v === '' ? undefined : parseInt(v, 10)) }
      : type === 'number'
      ? { valueAsNumber: true, setValueAs: (v) => (v === '' ? undefined : parseFloat(v)) }
      : null),
  })

  const fieldState = form.getFieldState(formField)

  const parentFormField = getParentFormField(formField)

  const formValue = form.getValues(formField)
  const formSlice = useFormSlice()
  const fieldWarnings = useFormSlice((state) => state.form[formName]?.customWarnings[formField])
  // const warning = formSlice.form[formName]?.customWarnings[formField]
  const formFieldWatch = form.watch(formField)
  const warningFn = options.fieldOptions?.[schemaKey]?.warningFn

  useEffect(() => {
    if (warningFn) {
      const warnings = joinWithAnd(warningFn(formFieldWatch))

      // FIXME: we are doing eerie warning and customerror checks all over the place to prevent inf rerender
      // but should use form field watch properly like right here and then trigger validation
      formSlice.setCustomWarning(formName, formField, warnings.length > 0 ? warnings : null)
      form.trigger(formField)
    }
  }, [formFieldWatch, fieldWarnings])

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
    ...(fieldWarnings && { rightSection: renderWarningIcon([fieldWarnings]) }),
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

  if (_props.error?.includes('match pattern')) {
    _props.error = `${itemName} does not match allowed patterns`
  }

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
        <RemoveButton formField={parentFormField} index={index} itemName={itemName} icon={<IconMinus size="1rem" />} />
      )}
    </Flex>
  )
}

type RemoveButtonProps = {
  // formField without item index
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
      // TODO: should initialize keys which are of type array to []
      //  (just first level, it will update itself recursively if somehow its nested)
      // else runtime operations on unset array throw and cant be catched.
      // see demoWorkItemCreateForm-base.items-add-button
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
  index: number
  formField: FormField
  schemaKey: SchemaKey
  handleValueRemove: (val: string) => void
  props?: React.HTMLProps<HTMLDivElement>
  warnings: Warnings
  setWarnings: React.Dispatch<React.SetStateAction<Warnings>>
}

type CustomMultiselectProps = {
  formField: FormField
  registerOnChange: ChangeHandler
  schemaKey: SchemaKey
  itemName: string
} & InputBaseProps

type Warnings = Record<string, string>

/**
 * Removes the last index from the form field.
 */
function getParentFormField(formField: FormField) {
  const formFieldKeys = formField.split('.')
  const formFieldArrayPath = formFieldKeys.slice(0, formFieldKeys.length - 1).join('.') as FormField
  return formFieldArrayPath
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

  const handleValueRemove = async (val: string) => {
    // index position changed, misleading message. must manually trigger validation for the field via trigger
    formSlice.setCustomError(formName, formField, null)
    form.unregister(formField) // needs to be called before setValue
    form.setValue(
      formField,
      formValues.filter((v) => v !== val),
    )
    await form.trigger(formField, { shouldFocus: true })
  }

  const comboboxOptions = selectOptions.values
    .filter((item: any) => {
      const inSearch = JSON.stringify(
        selectOptions.searchValueTransformer ? selectOptions.searchValueTransformer(item) : item,
      )
        .toLowerCase()
        .includes(search.toLowerCase().trim())

      return inSearch
    })
    .map((option) => {
      const formValue = selectOptions.formValueTransformer(option)
      const selected = formValues.includes(selectOptions.formValueTransformer(option))

      return (
        <Combobox.Option
          value={String(formValue)}
          key={String(formValue)}
          active={selected}
          aria-selected={selected}
          aria-label={
            selectOptions.ariaLabelTransformer ? selectOptions.ariaLabelTransformer(option) : String(formValue)
          }
        >
          <Group align="center" justify="start">
            {selected && <CheckIcon size={12} />}
            {selectOptions.optionTransformer(option)}
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

  const { rightSection, ...props } = inputProps

  const [warnings, setWarnings] = useState<Warnings>({})

  useEffect(() => {
    const msg = Object.values(warnings).join(',')
    if (formSlice.form[formName]?.customWarnings[formField] !== msg) {
      formSlice.setCustomWarning(formName, formField, msg)
    }
  }, [warnings])

  return (
    <Box w={'100%'}>
      <Combobox
        store={combobox}
        onOptionSubmit={(value, props) => {
          const option = selectOptions.values.find(
            (option) => String(selectOptions.formValueTransformer(option)) === value,
          )
          const selected = formValues.includes(selectOptions.formValueTransformer(option))
          if (selected) {
            handleValueRemove(selectOptions.formValueTransformer(option))
            return
          }
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
            {...props}
            // must override input props error
            error={multiselectFirstError}
            rightSection={renderWarningIcon(Object.values(warnings))}
          >
            <Pill.Group>
              {formValues.length > 0 &&
                formValues.map((formValue, i) => (
                  <CustomPill
                    index={i}
                    formField={formField}
                    key={`${formName}-${formField}-${i}-pill`}
                    value={formValue}
                    handleValueRemove={handleValueRemove}
                    schemaKey={schemaKey}
                    warnings={warnings}
                    setWarnings={setWarnings}
                  />
                ))}

              <Combobox.EventsTarget>
                <PillsInput.Field
                  placeholder={`Search ${pluralize(lowerFirst(itemName))}`}
                  onChange={async (event) => {
                    combobox.updateSelectedOptionIndex()
                    setSearch(event.currentTarget.value)
                  }}
                  data-testid={`${formName}-search--${formField}`}
                  value={search}
                  onFocus={() => combobox.openDropdown()}
                  onBlur={() => combobox.closeDropdown()}
                  onKeyDown={async (event) => {
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
} & InputBaseProps

function CustomSelect({ formField, registerOnChange, schemaKey, itemName, ...inputProps }: CustomSelectProps) {
  const form = useFormContext()
  const { formName, options, schemaFields } = useDynamicFormContext()
  const formSlice = useFormSlice()

  const selectOptions = options.selectOptions![schemaKey]!

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

  const formValue = form.getValues(formField)
  const selectedOption = selectOptions.values.find((option) => {
    return selectOptions.formValueTransformer(option) === formValue
  })

  // TODO: test.
  if (formValue !== null && formValue !== undefined && selectedOption === undefined) {
    const message = `${itemName} "${formValue}" does not exist`
    // else inf loop. we could unregister - direct form changes work - but state updates are lost (e.g. api call updates possible values) so its a no go
    if (formSlice.form[formName]?.customWarnings[formField] !== message) {
      formSlice.setCustomWarning(formName, formField, message)
    }
  } else {
    // once we receive api calls with correct data, etc. invalidate old warnings
    formSlice.setCustomWarning(formName, formField, null)
  }

  const comboboxOptions = selectOptions.values
    .filter((item: any) =>
      JSON.stringify(selectOptions.searchValueTransformer ? selectOptions.searchValueTransformer(item) : item)
        .toLowerCase()
        .includes(search.toLowerCase().trim()),
    )
    .map((option) => {
      const formValue = selectOptions.formValueTransformer(option)
      const selected = selectedOption ? selectOptions.formValueTransformer(selectedOption) === formValue : false

      return (
        <Combobox.Option
          value={String(formValue)}
          key={String(formValue)}
          aria-selected={selected}
          aria-label={
            selectOptions.ariaLabelTransformer ? selectOptions.ariaLabelTransformer(option) : String(formValue)
          }
        >
          <Group align="center" justify="start">
            {selected && <CheckIcon size={12} />}
            {comboboxOptionTemplate(selectOptions.optionTransformer, option)}
          </Group>
        </Combobox.Option>
      )
    })

  const parentSchemaKey = schemaKey.split('.').slice(0, -1).join('.') as SchemaKey

  const { rightSection, ...props } = inputProps

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
        <Combobox.Target withAriaAttributes={true}>
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
            name={formField}
            role="combobox"
            {...props}
          >
            {selectedOption ? (
              comboboxOptionTemplate(selectOptions.optionTransformer, selectedOption)
            ) : (
              <Input.Placeholder>
                <Flex direction="row" justify="space-between" align="center">
                  <Text size="sm">{`Pick ${singularize(lowerFirst(itemName))}`}</Text>
                  {rightSection}
                </Flex>
              </Input.Placeholder>
            )}
          </InputBase>
        </Combobox.Target>

        <Combobox.Dropdown>
          <Combobox.Search
            miw={'100%'}
            value={search}
            onChange={(event) => setSearch(event.currentTarget.value)}
            placeholder={`Search ${lowerFirst(itemName)}`}
            data-testid={`${formName}-search--${formField}`}
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

function renderWarningIcon(warnings?: string[] | string | null): React.ReactNode {
  const _warnings = (!isArray(warnings) ? [warnings] : warnings).filter((v) => !!v)

  return (
    _warnings &&
    _warnings?.length > 0 && (
      <Tooltip
        withinPortal
        label={
          <>
            {_warnings?.map((w, i) => (
              <Text key={i} size="sm">
                {w}
              </Text>
            ))}
          </>
        }
        position="top-end"
        withArrow
      >
        <IconAlertCircle size={20} color="orange" />
      </Tooltip>
    )
  )
}

function CustomPill({
  value,
  schemaKey,
  handleValueRemove,
  formField,
  warnings,
  setWarnings,
  index,
  ...props
}: CustomPillProps): JSX.Element | null {
  const { formName, options, schemaFields } = useDynamicFormContext()
  const selectOptions = options.selectOptions![schemaKey]!
  const itemName = singularize(options.labels[schemaKey] || '')
  // const formSlice = useFormSlice()
  // const warning = formSlice.form[formName]?.customWarnings[formField]

  let invalidValue = null

  const option = selectOptions.values.find((option) => selectOptions.formValueTransformer(option) === value)
  if (!option) {
    // for multiselects, explicitly set wrong values so that error positions make sense and the user knows there is a wrong form value beforehand
    // instead of us deleting it implicitly
    // TODO: should have multiselect and select option to allow creating values on the fly.
    // if no option found in search, a button option to `Create ${itemName}` shows a modal
    // with the form that creates that entity, e.g. tag, and once created we refetch selectOptions.values
    // see https://mantine.dev/combobox/?e=SelectCreatable
    invalidValue = value
    if (invalidValue !== null && invalidValue !== undefined) {
      if (!warnings[index]) {
        setWarnings((v) => ({ ...v, [index]: `${itemName} "${value}" does not exist` }))
      }
    }
  }

  let color = '#bbbbbb'
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
        onMouseDown={() => {
          if (invalidValue !== null) {
            setWarnings({})
            // formSlice.setCustomWarning(formName, formField, null)
          }
          handleValueRemove(value)
        }}
        variant="transparent"
        size={22}
        iconSize={14}
        tabIndex={-1}
        data-testid={`${formName}-${formField}-remove--${value}`}
        aria-label={`Remove ${itemName} ${value}`}
      />
    </Box>
  )
}
