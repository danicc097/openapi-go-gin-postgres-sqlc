import _ from 'lodash'
import { useState } from 'react'
import { FilterModes, useTableConfigSlice } from 'src/slices/tableConfig'

export function useMantineReactTableFilters(tableName: string) {
  const tableConfig = useTableConfigSlice()

  const dynamicConfig = useTableConfigSlice((state) => tableConfig.dynamicConfig[tableName])
  const staticConfig = useTableConfigSlice((state) => tableConfig.staticConfig[tableName])

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

  return { dynamicConfig, staticConfig, setFilterMode, removeFilterMode }
}
