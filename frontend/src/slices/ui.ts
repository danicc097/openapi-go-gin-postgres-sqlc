import Cookies from 'js-cookie'
import { devtools, persist } from 'zustand/middleware'
import { create } from 'zustand'
import { CONFIG } from 'src/config'
import { ProjectName } from 'src/gen/model'

export const LOGIN_COOKIE_KEY = CONFIG.LOGIN_COOKIE_KEY

export const UI_SLICE_PERSIST_KEY = 'ui-slice'

interface UIState {
  isLoggingOut: boolean
  setIsLoggingOut: (v: boolean) => void
  accessToken: string
  burgerOpened: boolean
  setBurgerOpened: (v: boolean) => void
  project: ProjectName
  setProject: (p: ProjectName) => void
  team: string | null
  setTeam: (p: string | null) => void
}

const useUISlice = create<UIState>()(
  devtools(
    persist(
      (set) => {
        return {
          isLoggingOut: false,
          setIsLoggingOut: (v: boolean) => set((state) => ({ isLoggingOut: v })),
          accessToken: Cookies.get(LOGIN_COOKIE_KEY) ?? '',
          burgerOpened: false,
          setBurgerOpened: (v: boolean) => set((state) => ({ burgerOpened: v })),
          project: 'demo',
          setProject: (v: ProjectName) => set((state) => ({ project: v })),
          team: '',
          setTeam: (v: string) => set((state) => ({ team: v })),
        }
      },
      {
        version: 2,
        name: UI_SLICE_PERSIST_KEY,
        partialize(state) {
          const { accessToken, isLoggingOut, ...rest } = state // always get access token from cookie
          return rest
        },
      },
    ),
    { enabled: import.meta.env.TESTING !== 'true' },
  ),
)

export { useUISlice }

type UIAction = (...args: any[]) => Partial<UIState>
