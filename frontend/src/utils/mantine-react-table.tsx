import { Checkbox, Text } from '@mantine/core'
import { MRT_ColumnDef, MRT_RowData } from 'mantine-react-table'
import { EntityFilter } from 'src/config'

function filterVariantByType(c: EntityFilter): MRT_ColumnDef<any>['filterVariant'] {
  if (c.type === 'boolean') return 'checkbox'
  if (c.type === 'number') return 'range'
  if (c.type === 'integer') return 'range'
  if (c.type === 'date-time') return 'date-range'

  return 'text'
}

export function columnPropsByType<T extends MRT_RowData>(id: string, c: EntityFilter): Partial<MRT_ColumnDef<T>> {
  return {
    filterVariant: filterVariantByType(c),
    Cell(props) {
      const val = props.row.original?.[id]
      if (c.type === 'boolean') return <Checkbox readOnly checked={val}></Checkbox>
      if (c.type === 'date-time') return <Text size="xs">{val?.toISOString()}</Text>

      return props.renderedCellValue
    },
    ...(c.type === 'boolean' && {
      mantineFilterCheckboxProps: {
        size: 'sm',
        label: 'Filter values',
      },
      enableColumnFilterModes: false,
    }),
  }
}
