import { Checkbox, Text } from '@mantine/core'
import { LiteralUnion, MRT_ColumnDef, MRT_FilterOption, MRT_RowData } from 'mantine-react-table'
import { EntityFilter } from 'src/config'

export const rangeModes: FilterModeOptions = ['between', 'betweenInclusive', 'inNumberRange']
export const emptyModes: FilterModeOptions = ['empty', 'notEmpty']
export const arrModes: FilterModeOptions = ['arrIncludesSome', 'arrIncludesAll', 'arrIncludes']
export const numberModes: FilterModeOptions = [
  ...rangeModes,
  'equals',
  'greaterThan',
  'greaterThanOrEqualTo',
  'lessThan',
  'lessThanOrEqualTo',
]
export const dateModes = ['between', 'betweenInclusive']
export const textModes: FilterModeOptions = ['contains', 'endsWith', 'equals', 'notEquals', 'startsWith'] //, 'fuzzy'

function filterVariantByType(c: EntityFilter): MRT_ColumnDef<any>['filterVariant'] {
  if (c.type === 'boolean') return 'checkbox'
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
      if (c.type === 'date-time') return <Text size="xs">{val?.toISOString()}</Text>

      return props.renderedCellValue
    },
    ...(c.type === 'date-time' && { size: 230 }),
    ...(c.type === 'boolean' && { size: 160 }),
    ...(c.type === 'boolean' && {
      mantineFilterCheckboxProps: {
        size: 'sm',
        label: 'Filter values',
      },
      enableColumnFilterModes: false,
    }),
  }
}
