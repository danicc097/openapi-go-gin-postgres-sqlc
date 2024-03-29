import { useMantineTheme, Menu, Button, rem, Text, ActionIcon, Tooltip, Badge, Flex, TextInput } from '@mantine/core'
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
} from '@tabler/icons'
import { IconRestore } from '@tabler/icons-react'
import { MRT_Column, MRT_Header, MRT_TableInstance } from 'mantine-react-table'
import { ComponentProps, useEffect, useRef, useState } from 'react'
import { EntityFilter, EntityFilterType } from 'src/config'
import { useMantineReactTableFilters } from 'src/hooks/ui/useMantineReactTableFilters'
import { emptyModes } from 'src/utils/mantine-react-table'
import classes from './mantine-react-table.module.css'
import { DateInput } from '@mantine/dates'
import { sentenceCase } from 'src/utils/strings'
import { c } from 'vitest/dist/reporters-MmQN-57K'
import { lowerCase } from 'lodash'

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
  type: EntityFilterType
  tableName: string
  columnProps: GenericColumnProps
}

export function CustomMRTFilter({ columnProps, nullable, type, tableName }: CustomMRTFilterProps) {
  const { dynamicConfig, removeFilterMode, setFilterMode } = useMantineReactTableFilters(tableName)
  const filterMode = dynamicConfig?.filterModes[columnProps.column.id] ?? ''
  if (emptyModes.includes(filterMode)) {
    if (columnProps.rangeFilterIndex === 1) {
      return null
    }
    return (
      <Badge className={'date-filter-badge'} size="sm">
        {sentenceCase(filterMode)}
      </Badge>
    )
  }
  if (type === 'date-time') {
    return <MRTDateInput columnProps={columnProps} />
  }

  return (
    <Flex gap={4} direction={'column'} pt={20}>
      <MRTTextInput columnProps={columnProps} />
      {filterMode && (
        <Text c="var(--mantine-color-placeholder)" size="xs" className="filter-mode-custom-label">
          Filter mode: {sentenceCase(filterMode)}
        </Text>
      )}
    </Flex>
  )
}

type MRTDateInputProps = {
  columnProps: GenericColumnProps
  props?: ComponentProps<typeof DateInput>
}

export function MRTDateInput({ columnProps: { column, rangeFilterIndex }, ...props }: MRTDateInputProps) {
  const { dynamicConfig, removeFilterMode, setFilterMode } = useMantineReactTableFilters('demoTable')
  const filterMode = dynamicConfig?.filterModes[column.id]

  return (
    <Flex gap={4} direction={'row'} pt={20} align="flex-start" justify="center">
      <DateInput
        placeholder={`${rangeFilterIndex === 0 ? 'Start' : 'End'} date`}
        size="xs"
        valueFormat="DD/MM/YYYY"
        classNames={{
          root: classes.root,
          input: classes.input,
          label: classes.label,
        }}
        miw={100}
        rightSection={
          /* TODO: may be cleaner to append nodes above via bare javascript below mrt-table-head-cell-content
      ideally mrt should allow rendering extra nodes below filters
      */
          filterMode && (
            <Text size="xs" fw={800}>
              {filterMode === 'between' ? '⇿' : '⬌'}
            </Text>
          )
        }
      />
    </Flex>
  )
}

type MRTTextInputProps = {
  columnProps: GenericColumnProps
  props?: ComponentProps<typeof TextInput>
}

export function MRTTextInput({ columnProps: { column }, ...props }: MRTTextInputProps) {
  const columnFilterValue = (column.getFilterValue() as string) ?? ''
  const [filterValue, setFilterValue] = useState<any>(() => columnFilterValue)
  const [debouncedFilterValue] = useDebouncedValue(filterValue, 400)
  const { dynamicConfig, removeFilterMode, setFilterMode } = useMantineReactTableFilters('demoTable')

  const isMounted = useRef(false)

  useEffect(() => {
    if (!isMounted.current) return
    column.setFilterValue(debouncedFilterValue ?? undefined)
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
      handleClear()
    } else {
      setFilterValue(tableFilterValue ?? '')
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [columnFilterValue])

  const filterMode = dynamicConfig?.filterModes[column.id]

  return (
    <TextInput
      {...props}
      value={filterValue}
      size="xs"
      onChange={(event) => {
        setFilterValue(event.currentTarget.value)
        if (!filterMode) setFilterMode(column.id, 'contains')
      }}
      rightSection={
        filterMode ? (
          <ActionIcon aria-label={'Clear search'} color="gray" onClick={handleClear} size="xs" variant="transparent">
            <Tooltip label={'Clear search'} withinPortal>
              <IconX />
            </Tooltip>
          </ActionIcon>
        ) : null
      }
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
