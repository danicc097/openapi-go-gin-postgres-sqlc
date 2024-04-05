import _ from 'lodash'
import { useState } from 'react'
import { FilterModes, initialDynamicConfig, initialStaticConfig, useTableConfigSlice } from 'src/slices/tableConfig'

export function useMantineReactTableFilters(tableName: string) {
  const tableConfig = useTableConfigSlice()

  const dynamicConfig = useTableConfigSlice((state) => state.dynamicConfig[tableName] ?? initialDynamicConfig)
  const staticConfig = useTableConfigSlice((state) => state.staticConfig[tableName] ?? initialStaticConfig)

  // const dynamicConfig = tableConfig.dynamicConfig[tableName] // won't trigger rerender
  // const staticConfig = tableConfig.staticConfig[tableName] // won't trigger rerender

  function setFilterMode(id: string, filterMode: string) {
    const modes = _.cloneDeep(dynamicConfig?.filterModes) ?? {}
    modes[id] = filterMode
    tableConfig.setFilterModes(tableName, modes)
  }

  function removeFilterMode(id: string) {
    const modes = _.cloneDeep(dynamicConfig?.filterModes) ?? {}
    delete modes[id]
    tableConfig.setFilterModes(tableName, modes)
  }

  function setHiddenColumns(columns: Record<string, boolean>) {
    console.log({ columns })
    tableConfig.setStaticConfig(tableName, {
      ...staticConfig,
      hiddenColumns: columns,
    })
  }

  function setColumnOrder(columns: string[]) {
    tableConfig.setStaticConfig(tableName, {
      ...staticConfig,
      columnOrder: columns,
    })
  }

  return { dynamicConfig, staticConfig, setFilterMode, removeFilterMode, setColumnOrder, setHiddenColumns }
}
