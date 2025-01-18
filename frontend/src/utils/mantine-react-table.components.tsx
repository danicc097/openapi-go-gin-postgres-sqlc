import {
  useMantineTheme,
  Menu,
  Button,
  rem,
  Text,
  ActionIcon,
  Tooltip,
  Badge,
  Flex,
  TextInput,
  Box,
  MenuItem,
  Checkbox,
  NumberInput,
} from '@mantine/core'
import { useDebouncedValue, useDisclosure } from '@mantine/hooks'
import {
  IconDots,
  IconPackage,
  IconSquareCheck,
  IconUsers,
  IconCalendar,
  IconEdit,
  IconTrash,
  IconSend,
  IconX,
  IconClearAll,
} from '@tabler/icons'
import { IconRestore, IconRowRemove } from '@tabler/icons-react'
import {
  MRT_Column,
  MRT_Header,
  MRT_InternalFilterOption,
  MRT_TableInstance,
  mrtFilterOptions,
} from 'mantine-react-table'
import { ComponentProps, RefObject, createElement, forwardRef, memo, useEffect, useRef, useState } from 'react'
import { EntityFilter, EntityFieldType } from 'src/config'
import { useMantineReactTableFilters } from 'src/hooks/ui/useMantineReactTableFilters'
import { emptyModes, indexOneModes, indexZeroModes, rangeModes } from 'src/utils/mantine-react-table'
import classes from './mantine-react-table.module.css'
import { DateInput } from '@mantine/dates'
import { sentenceCase } from 'src/utils/strings'
import _, { lowerCase } from 'lodash'
import { MRT_Localization_EN } from 'mantine-react-table/locales/en/index.esm.mjs'
import dayjs from 'dayjs'
import { render } from 'react-dom'
import { createRoot } from 'react-dom/client'

const FILTER_OPTIONS: MRT_InternalFilterOption[] = [
  ...mrtFilterOptions(MRT_Localization_EN),
  // can be extended with custom filter modes if required
  {
    label: 'Extra mode',
    divider: false,
    option: 'extraMode',
    symbol: '🆕',
  },
]

interface RowActionsMenuProps {
  canRestore: boolean
}

type GenericColumnProps = {
  column: MRT_Column<any, any>
  header: MRT_Header<any>
  rangeFilterIndex?: number
  table: MRT_TableInstance<any>
}

interface CustomMRTFilterProps {
  nullable: boolean
  type: EntityFieldType
  tableName: string
  columnProps: GenericColumnProps
}

export function CustomMRTFilter({ columnProps, nullable, type, tableName }: CustomMRTFilterProps) {
  const { dynamicConfig, removeFilterMode, setFilterMode } = useMantineReactTableFilters(tableName)
  const filterMode = dynamicConfig?.filterModes[columnProps.column.id] ?? ''

  const inputRef = useRef<HTMLInputElement>(null)

  const findInputFilterModeLabels = () =>
    inputRef.current?.closest('.mantine-Table-th')?.querySelectorAll('.filter-mode-helptext')

  // Effect to manage appending/removing the label
  useEffect(() => {
    const existingLabel = findInputFilterModeLabels()

    existingLabel?.forEach((e) => e.remove())
    if (columnProps.rangeFilterIndex === 1 && rangeModes.includes(filterMode)) return

    if (filterMode && inputRef.current) {
      const container = document.createElement('div')
      const root = createRoot(container)
      root.render(
        <p className={`${classes.filterMode} filter-mode-helptext`}>{`Filter mode: ${sentenceCase(filterMode)}`}</p>,
      )
      inputRef.current?.closest('.mantine-Table-th')?.appendChild(container)
    }
  }, [filterMode])

  if (emptyModes.includes(filterMode)) {
    const existingLabel = findInputFilterModeLabels()

    existingLabel?.forEach((e) => e.remove())
    if (columnProps.rangeFilterIndex === 1) {
      return null
    }
    return (
      <Badge className={'date-filter-badge'} size="sm">
        {sentenceCase(filterMode)}
      </Badge>
    )
  }

  if (
    (columnProps.rangeFilterIndex === 1 && indexZeroModes.includes(filterMode)) ||
    (columnProps.rangeFilterIndex === 0 && indexOneModes.includes(filterMode))
  ) {
    return (
      <Badge className={'date-filter-badge'} size="sm">
        Empty
      </Badge>
    )
  }

  if (type === 'number' || type === 'integer') {
    return (
      <MRTNumberInput
        ref={inputRef}
        columnProps={columnProps}
        type={type}
        props={
          {
            // rightSection: FILTER_OPTIONS.find((o) => o.option === filterMode)?.symbol
          }
        }
      />
    )
  }
  if (type === 'date-time') {
    return (
      <MRTDateInput
        ref={inputRef}
        columnProps={columnProps}
        props={
          {
            // rightSection: FILTER_OPTIONS.find((o) => o.option === filterMode)?.symbol
          }
        }
      />
    )
  }
  if (type === 'boolean') {
    return <MRTCheckboxInput ref={inputRef} columnProps={columnProps} />
  }

  return <MRTTextInput ref={inputRef} columnProps={columnProps} />
}

type MRTNumberInputProps = {
  columnProps: GenericColumnProps
  props?: ComponentProps<typeof NumberInput>
  type: EntityFieldType
}

export const MRTNumberInput = forwardRef(function MRTNumberInput(
  { columnProps: { column, rangeFilterIndex = 0 }, type, props: { ...props } = {} }: MRTNumberInputProps,
  ref,
) {
  const { dynamicConfig, removeFilterMode, setFilterMode } = useMantineReactTableFilters('demoTable')
  const filterMode = dynamicConfig?.filterModes[column.id]
  const columnRangeValue = (column.getFilterValue() as (string | undefined)[]) ?? ['', '']
  const columnFilterValue = columnRangeValue[rangeFilterIndex]
  const [filterValue, setFilterValue] = useState<any>(() => columnFilterValue)
  const [debouncedFilterValue] = useDebouncedValue<string>(filterValue, 250)

  const isMounted = useRef(false)
  // see https://github.com/KevinVandy/mantine-react-table/blob/v2/packages/mantine-react-table/src/components/inputs/MRT_FilterTextInput.tsx#L47
  // debounced doing weird things when being cleared
  useEffect(() => {
    if (!isMounted.current) return
    column.setFilterValue((old: [string, string]) => {
      const newFilterValues = Array.isArray(old) ? old : ['', '']
      newFilterValues[rangeFilterIndex] = debouncedFilterValue
      return newFilterValues
    })
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [debouncedFilterValue])

  const handleClear = () => {
    setFilterValue('')
    // dynamicConfig?.filterModes[column.id] && removeFilterMode(column.id)
    columnRangeValue[rangeFilterIndex] = ''
    column.setFilterValue(columnRangeValue)
  }

  useEffect(() => {
    if (((column.getFilterValue() as string[]) ?? []).every((i) => i === null || i === undefined || i === '')) {
      removeFilterMode(column.id)
    }
  }, [debouncedFilterValue])

  // one-off fire
  useEffect(() => {
    if (!isMounted.current) {
      isMounted.current = true
      return
    }
    const tableFilterValue = column.getFilterValue() as (string | undefined)[]

    if (_.isEqual(tableFilterValue, ['', '']) || tableFilterValue === undefined) {
      handleClear()
    } else {
      setFilterValue(tableFilterValue?.[rangeFilterIndex] ?? '')
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [column.getFilterValue()]) // don't use columnFilterValue

  return (
    <Flex ref={ref as any} gap={4} direction={'row'} pt={20} align="flex-start" justify="center">
      <NumberInput
        {...props}
        placeholder={rangeFilterIndex === 0 ? 'Min' : 'Max'}
        data-testid={`input-filter--${column.id}-${rangeFilterIndex === 0 ? 'min' : 'max'}`}
        value={filterValue}
        allowDecimal={type === 'number'}
        onChange={(event) => {
          setFilterValue(event !== '' ? event : null)
          if (!filterMode) setFilterMode(column.id, 'between')
        }}
        size="xs"
        classNames={{
          root: classes.root,
          input: classes.input,
          label: classes.label,
        }}
        miw={40}
        rightSection={filterValue ? renderClearSearchButton(handleClear) : <></>}
      />
    </Flex>
  )
})

type MRTDateInputProps = {
  columnProps: GenericColumnProps
  props?: ComponentProps<typeof DateInput>
}

export const MRTDateInput = forwardRef(function MRTDateInput(
  { columnProps: { column, rangeFilterIndex = 0 }, props: { ...props } = {} }: MRTDateInputProps,
  ref,
) {
  const { dynamicConfig, removeFilterMode, setFilterMode } = useMantineReactTableFilters('demoTable')
  const filterMode = dynamicConfig?.filterModes[column.id]
  const columnRangeValue = (column.getFilterValue() as (string | undefined)[]) ?? [undefined, undefined]
  const columnFilterValue = columnRangeValue[rangeFilterIndex]
  const [filterValue, setFilterValue] = useState<any>(() => columnFilterValue)
  const [debouncedFilterValue] = useDebouncedValue(filterValue, 250)

  const isMounted = useRef(false)

  useEffect(() => {
    if (!isMounted.current) return
    columnRangeValue[rangeFilterIndex] = debouncedFilterValue ?? undefined
    column.setFilterValue(columnRangeValue)
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [debouncedFilterValue])

  const handleClear = () => {
    setFilterValue(undefined)
    // dynamicConfig?.filterModes[column.id] && removeFilterMode(column.id)
    columnRangeValue[rangeFilterIndex] = undefined
    column.setFilterValue(columnRangeValue)
  }

  useEffect(() => {
    if (((column.getFilterValue() as string[]) ?? []).every((i) => i === null || i === undefined || i === '')) {
      removeFilterMode(column.id)
    }
  }, [debouncedFilterValue])

  useEffect(() => {
    if (!isMounted.current) {
      isMounted.current = true
      return
    }
    const tableFilterValue = column.getFilterValue() as (string | undefined)[]
    console.log({ tableFilterValue })
    if (_.isEqual(tableFilterValue, [undefined, undefined])) {
      handleClear()
    } else {
      setFilterValue(tableFilterValue?.[rangeFilterIndex] ?? undefined)
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [column.getFilterValue()]) // don't use columnFilterValue

  return (
    <Flex ref={ref as any} gap={4} direction={'row'} pt={20} align="flex-start" justify="center">
      <DateInput
        {...props}
        data-testid={`input-filter--${column.id}-${rangeFilterIndex === 0 ? 'min' : 'max'}`}
        placeholder={`${rangeFilterIndex === 0 ? 'Min' : 'Max'} date`}
        value={filterValue ? dayjs(filterValue).toDate() : null}
        onChange={(event) => {
          setFilterValue(event ? dayjs(event.toDateString()) : null)
          if (!filterMode) setFilterMode(column.id, 'between')
        }}
        size="xs"
        valueFormat="DD/MM/YYYY"
        classNames={{
          root: classes.root,
          input: classes.input,
          label: classes.label,
        }}
        miw={60}
        rightSection={filterValue ? renderClearSearchButton(handleClear) : null}
      />
    </Flex>
  )
})

type MRTTextInputProps = {
  columnProps: GenericColumnProps
  props?: ComponentProps<typeof TextInput>
}

interface CustomColumnFilterModeMenuItemsProps {
  modeOptions?: string[] | null
  column: GenericColumnProps['column']
}

export const CustomColumnFilterModeMenuItems = memo(({ modeOptions, column }: CustomColumnFilterModeMenuItemsProps) => {
  const { dynamicConfig, removeFilterMode, setFilterMode } = useMantineReactTableFilters('demoTable')
  return modeOptions && modeOptions.length > 0 ? (
    <>
      {modeOptions.map((option) => {
        const fopt = FILTER_OPTIONS.find((v) => v.option === option)
        if (!fopt) return

        return (
          <MenuItem
            key={fopt.option}
            onClick={() => {
              column.setFilterValue(null)
              setFilterMode(column.id, fopt.option)
            }}
          >
            <Flex
              gap={10}
              justify="flex-start"
              align="center"
              style={{
                color:
                  dynamicConfig?.filterModes[column.id ?? ''] === fopt.option
                    ? 'var(--mantine-primary-color-5)'
                    : 'inherit',
              }}
            >
              <Box miw={20} style={{ alignSelf: 'center', textAlign: 'center' }}>
                {fopt.symbol}
              </Box>
              <Text size="sm">{sentenceCase(fopt.label)}</Text>
            </Flex>
          </MenuItem>
        )
      })}
      <Menu.Divider />
      <MenuItem
        key={'clearFilter'}
        onClick={() => {
          column.setFilterValue(undefined)
          removeFilterMode(column.id)
        }}
      >
        <Flex gap={10} justify="flex-start" align="center">
          <IconClearAll stroke={1} size={18} />
          <Text size="sm">Clear filters</Text>
        </Flex>
      </MenuItem>
    </>
  ) : (
    <Text size="xs" p={8}>
      No options available
    </Text>
  )
})

type MRTCheckboxInputProps = {
  columnProps: GenericColumnProps
  props?: ComponentProps<typeof Checkbox>
}

export const MRTCheckboxInput = forwardRef(function MRTCheckboxInput(
  { columnProps: { column }, props: { ...props } = {} }: MRTCheckboxInputProps,
  ref,
) {
  const { dynamicConfig, removeFilterMode, setFilterMode } = useMantineReactTableFilters('demoTable')

  const filterMode = dynamicConfig?.filterModes[column.id]
  const value = column.getFilterValue()

  return (
    <Checkbox
      {...props}
      data-testid={`input-filter--${column.id}`}
      ref={ref as any}
      checked={value === 'true'}
      data-checked={value === 'true'}
      {...(value === undefined && { indeterminate: true, 'data-indeterminate': true })}
      size="xs"
      onChange={(event) => {
        const newValue =
          column.getFilterValue() === undefined ? 'true' : column.getFilterValue() === 'true' ? 'false' : undefined
        column.setFilterValue(newValue)
        if (!filterMode) setFilterMode(column.id, 'equals')
        if (newValue === undefined) removeFilterMode(column.id)
      }}
      label={`Filter values`}
      // labelProps={{ 'data-floating': floating }}
      classNames={{
        root: classes.checkBox,
        label: classes.checkboxLabel,
      }}
    />
  )
})

export const MRTTextInput = forwardRef(function MRTTextInput(
  { columnProps: { column }, props: { ...props } = {} }: MRTTextInputProps,
  ref,
) {
  const columnFilterValue = (column.getFilterValue() as string) ?? ''
  const [filterValue, setFilterValue] = useState<any>(() => columnFilterValue)
  const [debouncedFilterValue] = useDebouncedValue(filterValue, 250)
  const { dynamicConfig, removeFilterMode, setFilterMode } = useMantineReactTableFilters('demoTable')

  const isMounted = useRef(false)

  useEffect(() => {
    if (!isMounted.current) return
    column.setFilterValue(debouncedFilterValue ?? undefined)
    if (!debouncedFilterValue) {
      removeFilterMode(column.id)
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [debouncedFilterValue])

  const [focused, setFocused] = useState(false)
  const floating = focused || filterValue?.length > 0 || undefined

  const handleClear = () => {
    setFilterValue('')
    removeFilterMode(column.id)
    column.setFilterValue(undefined)
  }

  //receive table filter value and set it to local state
  useEffect(() => {
    if (!isMounted.current) {
      isMounted.current = true
      return
    }
    const tableFilterValue = column.getFilterValue()
    if (tableFilterValue === undefined) {
      setFilterValue('')
    } else {
      setFilterValue(tableFilterValue ?? '')
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [column.getFilterValue()]) // don't use columnFilterValue

  const filterMode = dynamicConfig?.filterModes[column.id]

  return (
    <TextInput
      {...props}
      data-testid={`input-filter--${column.id}`}
      ref={ref as any}
      value={filterValue}
      size="xs"
      onChange={(event) => {
        setFilterValue(event.currentTarget.value)
        if (!filterMode) setFilterMode(column.id, 'contains')
      }}
      rightSection={filterValue ? renderClearSearchButton(handleClear) : null}
      placeholder={`Filter by ${lowerCase(column.id)}`}
      // labelProps={{ 'data-floating': floating }}
      classNames={{
        root: classes.root,
        input: classes.input,
        label: classes.label,
      }}
      onFocus={() => setFocused(true)}
      onBlur={() => setFocused(false)}
    />
  )
})

function renderClearSearchButton(handleClear: () => void) {
  return (
    <Tooltip label={'Clear search'} withinPortal>
      <Box>
        <ActionIcon aria-label={'Clear search'} color="gray" onClick={handleClear} size="xs" variant="transparent">
          <IconX size={14} />
        </ActionIcon>
      </Box>
    </Tooltip>
  )
}

export function RowActionsMenu({ canRestore: canBeRestored }: RowActionsMenuProps) {
  const theme = useMantineTheme()
  const [loading, setLoading] = useState(false)

  return (
    <Menu transitionProps={{ transition: 'pop-bottom-left' }} width={220} withinPortal withArrow>
      <Menu.Target>
        <Tooltip label="Show actions" withArrow>
          <ActionIcon>
            <IconDots style={{ height: rem(18) }} stroke={1.5} />
          </ActionIcon>
        </Tooltip>
      </Menu.Target>
      <Menu.Dropdown>
        <Menu.Item leftSection={<IconEdit style={{ height: rem(16) }} color={theme.colors.blue[6]} stroke={1.5} />}>
          Edit
        </Menu.Item>
        <Menu.Item leftSection={<IconTrash style={{ height: rem(16) }} color={theme.colors.red[6]} stroke={1.5} />}>
          Delete
        </Menu.Item>
        {canBeRestored && (
          <Menu.Item
            leftSection={<IconRestore style={{ height: rem(16) }} color={theme.colors.yellow[6]} stroke={1.5} />}
          >
            Restore
          </Menu.Item>
        )}
        <Menu.Divider />
        <Menu.Item leftSection={<IconSend style={{ height: rem(16) }} color={theme.colors.cyan[6]} stroke={1.5} />}>
          Message members
        </Menu.Item>
      </Menu.Dropdown>
    </Menu>
  )
}
