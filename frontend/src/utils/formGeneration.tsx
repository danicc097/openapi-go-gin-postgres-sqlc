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
} from '@mantine/core'
import classes from './form.module.css'
import { DateInput, DateTimePicker } from '@mantine/dates'
import { useFocusWithin } from '@mantine/hooks'
import { CodeHighlight } from '@mantine/code-highlight'
import { rem, useMantineTheme } from '@mantine/core'
import { Icon123, IconMinus, IconPlus, IconTrash } from '@tabler/icons'
import { pluralize, singularize } from 'inflection'
import _, { lowerFirst, memoize } from 'lodash'
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
} from 'react-hook-form'
import { json } from 'react-router-dom'
import { ApiError } from 'src/api/mutator'
import ErrorCallout, { useCalloutErrors } from 'src/components/ErrorCallout/ErrorCallout'
import PageTemplate from 'src/components/PageTemplate'
import type { DemoWorkItemCreateRequest } from 'src/gen/model'
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
import { entries } from 'src/utils/object'
import { nameInitials, sentenceCase } from 'src/utils/strings'

export type SelectOptionsTypes = 'select' | 'multiselect'

export type SelectOptions<Return, E = unknown> = {
  type: SelectOptionsTypes
  values: E[]
  formValueTransformer: <V extends E>(el: V & E) => Return extends unknown[] ? Return[number] : Return
  optionTransformer: <V extends E>(el: V & E) => JSX.Element
  labelTransformer?: <V extends E>(el: V & E) => JSX.Element
  labelColor?: <V extends E>(el: V & E) => string
}

export interface InputOptions<Return, E = unknown> {
  component: JSX.Element
  propsFn?: (registerOnChange: ChangeHandler) => React.ComponentProps<'input'>
}

// NOTE: handles select (single return value) and multiselect (array return).
export const selectOptionsBuilder = <Return, V, ReturnElement = Return extends unknown[] ? Return[number] : Return>({
  type,
  values,
  formValueTransformer,
  optionTransformer,
  labelTransformer,
  labelColor,
}: SelectOptions<ReturnElement, V>): SelectOptions<ReturnElement, V> => ({
  type,
  values,
  optionTransformer,
  labelTransformer,
  formValueTransformer,
  labelColor,
})

export const inputBuilder = <Return, V>({ component }: InputOptions<Return, V>): InputOptions<Return, V> => ({
  component,
})

const comboboxOptionTemplate = (transformer: (...args: any[]) => JSX.Element, option) => {
  return <Box m={2}>{transformer(option)}</Box>
}

interface MultiSelectValueProps extends React.ComponentPropsWithoutRef<'div'> {
  value: string
  onRemove?(): void
}

const valueComponentTemplate =
  (transformer: (...args: any[]) => JSX.Element, colorFn?: (...args: any[]) => string) =>
  ({ value, option, onRemove, ...others }: MultiSelectValueProps & { value: string; option: any }) => {
    let color
    if (colorFn) {
      color = colorFn(option)
    }
    return (
      <div {...others}>
        <Box className={classes['value-component-outer-box']}>
          <Box
            className={classes['value-component-inner-box']}
            color={getContrastYIQ(color) === 'black' ? 'whitesmoke' : 'black'}
          >
            {transformer(option)}
          </Box>
          <CloseButton
            //@ts-ignore
            onMouseDown={onRemove}
            variant="transparent"
            size={22}
            iconSize={14}
            tabIndex={-1}
          />
        </Box>
      </div>
    )
  }

export type DynamicFormOptions<T extends object, ExcludeKeys extends U | null, U extends PropertyKey = GetKeys<T>> = {
  labels: {
    [key in Exclude<U, ExcludeKeys>]: string | null
  }
  renderOrderPriority?: Array<Exclude<keyof T, ExcludeKeys>>
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

type DynamicFormProps<T extends object, ExcludeKeys extends GetKeys<T> | null = null> = {
  schemaFields: Record<Exclude<GetKeys<T>, ExcludeKeys>, SchemaField>
  options: DynamicFormOptions<T, ExcludeKeys, GetKeys<T>>
  formName: string
  onSubmit: React.FormEventHandler<HTMLFormElement>
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

export default function DynamicForm<T extends object, ExcludeKeys extends GetKeys<T> | null = null>({
  formName,
  schemaFields,
  options,
  onSubmit,
}: DynamicFormProps<T, ExcludeKeys>) {
  const theme = useMantineTheme()
  const form = useFormContext()
  const { extractCalloutErrors, setCalloutErrors, calloutErrors } = useCalloutErrors()

  useEffect(() => {
    setCalloutErrors(new ApiError('Remote error message'))
  }, [])

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
        <ErrorCallout title="Custom error title" errors={extractCalloutErrors()} />
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
    const itemName = singularize(options.labels[schemaKey] || '')

    const inputProps = {
      css: css`
        width: 100%;
      `,
      ...(!field.isArray && { label: options.labels[schemaKey] }),
      required: field.required,
      id: `${formName}-${formField}`,
      onKeyPress: (e: React.KeyboardEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        e.key === 'Enter' && e.preventDefault()
      },
    }

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
        <Card mt={12} mb={12} withBorder radius={cardRadius} className={classes['child-card']}>
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
              id: `${formName}-${formField}`,
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
    <Card radius={cardRadius} p={6} withBorder className={classes['array-child-card']}>
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
        <Accordion.Control>{`See form`}</Accordion.Control>
        <Accordion.Panel>
          <CodeHighlight language="json" code={JSON.stringify(myFormData, null, 2)}></CodeHighlight>
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

  const _props = {
    ...registerProps,
    ...props?.input,
    ...(propsOverride && propsOverride),
    ...(!fieldState.isDirty && { defaultValue: convertValueByType(type, formValue) }),
    ...(fieldState.error && { error: sentenceCase(fieldState.error?.message) }),
    required: schemaFields[schemaKey]?.required && type !== 'boolean',
    placeholder: `Enter ${lowerFirst(singularize(options.labels[schemaKey] || ''))}`,
  }

  let el: JSX.Element | null = null
  let customEl: JSX.Element | null = null
  const component = options.input?.[schemaKey]?.component
  // TODO: componentPropsFn must return {}
  const componentPropsFn = options.input?.[schemaKey]?.propsFn
  const selectOptions = options.selectOptions?.[schemaKey]
  const selectRef = useRef<HTMLInputElement | null>(null)
  const [customElMinHeight, setCustomElMinHeight] = useState(34.5)
  const [search, setSearch] = useState('')

  const { ref: focusRef, focused: selectFocused } = useFocusWithin()

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
    el = React.cloneElement(component, {
      ..._props,
      // TODO: this depends on component type, onChange should be customizable in options parameter with registerOnChange as fn param
      // props
      onChange: (e) => registerOnChange({ target: { name: formField, value: e } }),
      ...component.props, // allow user override
      ...(componentPropsFn && componentPropsFn(registerOnChange)), // allow user override
    })
    // TODO: multiSelectOptions: https://codesandbox.io/s/watch-with-usefieldarray-forked-9383hz?file=/src/formGeneration.tsx
    // which do allow custom labels by default and doesnt need workaround
    // use them with tagIDs -> DbWorkItemTag[] -> tag.name
  } else if (selectOptions) {
    switch (selectOptions.type) {
      case 'select':
        {
          const selectedOption = selectOptions.values.find((option) => {
            return selectOptions.formValueTransformer(option) === form.getValues(formField)
          })

          console.log(selectOptions.values)
          const options = selectOptions.values
            .filter((item: any) => JSON.stringify(item).toLowerCase().includes(search.toLowerCase().trim()))
            .map((option) => {
              const value = String(selectOptions.formValueTransformer(option))

              return (
                <Combobox.Option value={value} key={value}>
                  {comboboxOptionTemplate(selectOptions.optionTransformer, option)}
                </Combobox.Option>
              )
            })

          // IMPORTANT: mantine assumes label = value, else it doesn't work: https://github.com/mantinedev/mantine/issues/980
          el = (
            <Box miw={'100%'}>
              <Combobox
                store={combobox}
                withinPortal={true}
                onOptionSubmit={async (value) => {
                  const option = selectOptions.values.find(
                    (option) => String(selectOptions.formValueTransformer(option)) === value,
                  )
                  console.log({ onChangeOption: option })
                  if (!option) return
                  await registerOnChange({
                    target: {
                      name: formField,
                      value: selectOptions.formValueTransformer(option),
                    },
                  })
                  setIsSelectVisible(false)
                  combobox.closeDropdown()
                }}
              >
                <Combobox.Target>
                  <InputBase
                    className={classes['select']}
                    component="button"
                    type="button"
                    pointer
                    rightSection={<Combobox.Chevron />}
                    onClick={() => combobox.toggleDropdown()}
                    rightSectionPointerEvents="none"
                    multiline
                  >
                    {selectedOption ? (
                      comboboxOptionTemplate(selectOptions.optionTransformer, selectedOption)
                    ) : (
                      <Input.Placeholder>Pick value</Input.Placeholder>
                    )}
                  </InputBase>
                </Combobox.Target>

                <Combobox.Search
                  miw={'100%'}
                  value={search}
                  onChange={(event) => setSearch(event.currentTarget.value)}
                  placeholder={`Search items`}
                />

                <Combobox.Dropdown>
                  <Combobox.Options>
                    {options.length > 0 ? options : <Combobox.Empty>Nothing found</Combobox.Empty>}
                  </Combobox.Options>
                </Combobox.Dropdown>
              </Combobox>
            </Box>

            // <Select
            //   onBlur={(e) => setIsSelectVisible(false)}
            //   withinPortal
            //   initiallyopened={selectedOption !== undefined}
            //   searchable
            //   // TODO: need to have typed selectOptions.filter. that way we can filter user.email, username, etc.
            //   // if not set use generic JSON.stringify(item.option).toLowerCase().includes(option.toLowerCase().trim())
            //   // else we need to forcefully use current label/value
            //   filter={(input) => {
            //     if (selectedOption !== '') {
            //       return JSON.stringify(item.option).toLowerCase().includes(selectedOption.toLowerCase().trim())
            //     }

            //     return JSON.stringify(item.option).toLowerCase().includes(selectedOption.toLowerCase().trim())
            //   }}
            //   // IMPORTANT: Select value should always be either string or null as per doc
            //   // (and implicitly label must be equal to value else all is broken unlike with multiselect)
            //   data={selectOptions.values.map((option) => ({
            //     label: String(selectOptions.formValueTransformer(option)),
            //     value: String(selectOptions.formValueTransformer(option)),
            //     option,
            //   }))}
            //   onChange={async (value) => {
            //     const option = selectOptions.values.find(
            //       (option) => String(selectOptions.formValueTransformer(option)) === value,
            //     )
            //     console.log({ onChangeOption: option })
            //     if (!option) return
            //     await registerOnChange({
            //       target: {
            //         name: formField,
            //         value: selectOptions.formValueTransformer(option),
            //       },
            //     })
            //     setIsSelectVisible(false)
            //   }}
            //   value={String(form.getValues(formField))}
            //   {..._props}
            //   ref={(el) => {
            //     selectRef.current = el
            //     focusRef.current = el
            //   }}
            //   placeholder={`Select ${lowerFirst(itemName)}`}
            // />
          )

          // TODO: maybe with v7 this hack isn't needed
          if (!isSelectVisible && selectedOption !== undefined) {
            const { ref, ...customSelectProps } = _props
            customEl = (
              <Input.Wrapper {...customSelectProps} pt={0} pb={0}>
                <Card
                  tabIndex={0}
                  css={css`
                    min-height: ${customElMinHeight}px;
                    display: flex;
                    align-items: center;
                    justify-content: space-between;
                    font-size: ${theme.fontSizes.sm};
                    place-content: center;

                    :focus {
                      border-color: var(--mantine-color-blue-8) !important;
                    }
                  `}
                  onKeyUp={(e) => {
                    if (e.key === 'Enter') setIsSelectVisible(true)
                  }}
                  withBorder
                  //onFocus={toggleVisibility}
                  pl={12}
                  pr={12}
                  pt={0}
                  pb={0}
                  onClick={() => {
                    setIsSelectVisible(true)
                    selectRef.current?.focus()
                  }}
                >
                  {selectOptions.labelTransformer
                    ? selectOptions.labelTransformer(selectedOption)
                    : selectOptions.optionTransformer(selectedOption)}
                </Card>
              </Input.Wrapper>
            )
          }
        }
        break
      case 'multiselect':
        break
        // TODO:
        {
          const data = selectOptions.values.map((option) => ({
            label: selectOptions.formValueTransformer(option),
            value: selectOptions.formValueTransformer(option),
            option,
          }))

          el = (
            <MultiSelect
              styles={{
                input: 'background: light-dark(white, var(--mantine-color-dark-7));',
              }}
              withinPortal
              // TODO: in v7: see https://mantine.dev/combobox/?e=MultiSelectValueRenderer
              // see examples with `custom value` https://mantine.dev/combobox/
              // itemComponent={comboboxOptionTemplate(selectOptions.optionTransformer)}
              valueComponent={valueComponentTemplate(
                selectOptions.labelTransformer ? selectOptions.labelTransformer : selectOptions.optionTransformer,
                selectOptions.labelColor,
              )}
              searchable
              filter={(option, selected, item) => {
                if (option !== '') {
                  return JSON.stringify(item.option).toLowerCase().includes(option.toLowerCase().trim())
                }

                return JSON.stringify(item.option).toLowerCase().includes(option.toLowerCase().trim())
              }}
              data={data}
              onChange={async (values) => {
                console.log({ values, formValues: form.getValues(formField) })
                const options = values.map((value) =>
                  selectOptions.values.find((option) => value === selectOptions.formValueTransformer(option)),
                )
                console.log({ onChangeOptions: options })
                //if (!option) return
                registerOnChange({
                  target: {
                    name: formField,
                    value: options.map((o) => selectOptions.formValueTransformer(o)),
                  },
                })
              }}
              value={form.getValues(formField)}
              {..._props}
              ref={selectRef}
              placeholder={`Select ${lowerFirst(pluralize(itemName))}`}
            />
          )
        }
        break
      default:
        break
    }
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
            {..._props}
            placeholder="Select date"
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
            {..._props}
            placeholder="Select date and time"
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
      {customEl || el}
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
  const theme = useMantineTheme()
  const { formName, options, schemaFields } = useDynamicFormContext()

  return (
    <Tooltip withinPortal label={`Remove ${itemName}`} position="top-end" withArrow>
      <ActionIcon
        onClick={(e) => {
          // NOTE: don't use rhf useFieldArray, way too many edge cases for no gain. if reordering is needed, implement it manually.
          // we could even implement flat array reordering by handling them internally as objects with id prop
          const listItems = form.getValues(formField)
          removeElementByIndex(listItems, index)
          form.unregister(formField) // needs to be before setValue
          form.setValue(formField, listItems as any)
        }}
        // variant="filled"
        className={classes['remove-button']}
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
      <Flex direction="row" align="center">
        {!accordion && options.labels[schemaKey] && renderTitle(formField, options.labels[schemaKey])}
        {options.selectOptions?.[schemaKey]?.type !== 'multiselect' && (
          <Button
            size="xs"
            p={4}
            leftSection={<IconPlus size="1rem" />}
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
        )}
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
