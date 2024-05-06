import Cookies from 'js-cookie'
import { devtools, persist } from 'zustand/middleware'
import { create } from 'zustand'
import { CONFIG } from 'src/config'

export type FilterModes = Record<string, string>

export const LOGIN_COOKIE_KEY = CONFIG.LOGIN_COOKIE_KEY

export const CONFIG_SLICE_PERSIST_KEY = 'config-slice'

export type DynamicConfig = {
  filterModes: FilterModes
}

export type StaticConfig = {
  hiddenColumns?: Record<string, boolean> // since they will change on app updates, just store hidden ones
  columnOrder?: string[]
}

interface TableConfigState {
  dynamicConfig: {
    [tableName: string]: DynamicConfig
  }
  staticConfig: {
    [tableName: string]: StaticConfig
  }
  setFilterModes(tableName: string, filterModes: FilterModes): void
  setStaticConfig(tableName: string, config: StaticConfig): void
}

const initialDynamicConfig: DynamicConfig = {
  filterModes: {},
}
const initialStaticConfig: StaticConfig = {}
const useTableConfigSlice = create<TableConfigState>()(
  devtools(
    persist(
      (set) => {
        return {
          dynamicConfig: {},
          staticConfig: {},
          setFilterModes: (tableName, filterModes) => {
            set((state) => {
              const dynamicConfig = state.dynamicConfig[tableName] || initialDynamicConfig

              return {
                ...state,
                dynamicConfig: {
                  ...state.dynamicConfig,
                  [tableName]: {
                    ...dynamicConfig,
                    filterModes: filterModes,
                  },
                },
              }
            })
          },
          setStaticConfig(tableName, config) {
            return set((state) => {
              return {
                ...state,
                staticConfig: {
                  ...state.staticConfig,
                  [tableName]: config,
                },
              }
            })
          },
        }
      },
      {
        version: 2,
        name: CONFIG_SLICE_PERSIST_KEY,
        partialize(state) {
          const { dynamicConfig, ...rest } = state // just want to persist visualization settings
          return rest
        },
      },
    ),
    { enabled: import.meta.env.TESTING !== 'true' },
  ),
)

export { useTableConfigSlice, initialDynamicConfig, initialStaticConfig }
