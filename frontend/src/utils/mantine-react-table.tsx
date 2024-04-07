import { Checkbox, Popover, Text } from '@mantine/core'
import { LiteralUnion, MRT_ColumnDef, MRT_FilterOption, MRT_RowData } from 'mantine-react-table'
import { EntityFilter } from 'src/config'
import classes from './mantine-react-table.module.css'
import dayjs from 'dayjs'
import { ReactNode, cloneElement, isValidElement } from 'react'

export const rangeModes: FilterModeOptions = ['between', 'betweenInclusive']
export const emptyModes: FilterModeOptions = ['empty', 'notEmpty']
export const arrModes: FilterModeOptions = ['arrIncludesSome', 'arrIncludesAll', 'arrIncludes']
export const indexZeroModes = ['equals', 'greaterThan', 'greaterThanOrEqualTo']
export const indexOneModes = ['lessThan', 'lessThanOrEqualTo']
export const numberModes: FilterModeOptions = [...rangeModes, ...indexZeroModes, ...indexOneModes]
export const dateModes = rangeModes
export const textModes: FilterModeOptions = ['contains', 'endsWith', 'equals', 'notEquals', 'startsWith'] //, 'fuzzy'

function filterVariantByType(c: EntityFilter): MRT_ColumnDef<any>['filterVariant'] {
  if (c.type === 'boolean') return 'text' // 'checkbox' will ignore the given Filter fn
  if (c.type === 'number') return 'range'
  if (c.type === 'integer') return 'range'
  if (c.type === 'date-time') return 'date-range'

  return 'text'
}

type FilterModeOptions = Array<LiteralUnion<string & MRT_FilterOption>>

function columnFilterModeOptionsByType(c: EntityFilter): FilterModeOptions {
  const modes: FilterModeOptions = []
  switch (c.type) {
    case 'number':
      modes.push(...numberModes)
      break
    case 'integer':
      modes.push(...numberModes)
      break
    case 'date-time':
      modes.push(...dateModes)
      break
    default:
      modes.push(...textModes)
      break
  }

  if (c.nullable) {
    modes.push(...emptyModes)
  }

  return modes
}

export function columnPropsByType<T extends MRT_RowData>(id: string, c: EntityFilter): Partial<MRT_ColumnDef<T>> {
  return {
    filterVariant: filterVariantByType(c),
    columnFilterModeOptions: columnFilterModeOptionsByType(c),
    Cell(props) {
      const val = props.row.original?.[id]
      if (c.type === 'boolean') return <Checkbox size="xs" readOnly checked={val}></Checkbox>
      if (c.type === 'date-time')
        return (
          <Popover withArrow>
            <Popover.Target>
              <Text size="xs">{val?.toISOString()}</Text>
            </Popover.Target>
            <Popover.Dropdown>
              <Text size="xs">{dayjs().to(dayjs(val))}</Text>
            </Popover.Dropdown>
          </Popover>
        )

      return props.renderedCellValue
    },
    ...(c.type === 'date-time' && { size: 230 }),
    ...(c.type === 'boolean' && { size: 160 }),
    ...(c.type === 'boolean' && {
      mantineFilterCheckboxProps: {
        size: 'xs',
        label: 'Filter values',
        classNames: {
          label: classes.checkboxLabel,
        },
      },
      enableColumnFilterModes: false,
    }),
    // enableColumnActions: false,
    renderColumnActionsMenuItems(props) {
      const menuItems = removeNodesWithTextContent(props.internalColumnMenuItems, 'Clear filter') // TODO: array to remove, and regex
      return menuItems
    },
  }
}

const removeNodesWithTextContent = (node: ReactNode, textContentToRemove: string): ReactNode => {
  if (typeof node === 'string') {
    return node
  }

  if (Array.isArray(node)) {
    return node.map((n, index) => {
      const modifiedNode = removeNodesWithTextContent(n, textContentToRemove)
      // Preserve the key prop for arrays
      if (isValidElement(n) && isValidElement(modifiedNode)) {
        return cloneElement(modifiedNode, { key: n.key ?? index })
      }
      return modifiedNode
    })
  }

  if (isValidElement(node)) {
    const element = node as React.ReactElement
    if (element?.props?.children?.toString().includes(textContentToRemove)) {
      return null
    }

    const props = element.props || {}
    const children = props.children

    if (!children) {
      return cloneElement(element, props)
    }

    return cloneElement(element, {
      ...props,
      children: removeNodesWithTextContent(children, textContentToRemove),
    })
  }

  return node
}
