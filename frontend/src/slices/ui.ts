import Cookies from 'js-cookie'
import { devtools, persist } from 'zustand/middleware'
import { create } from 'zustand'
import { CONFIG } from 'src/config'

export const LOGIN_COOKIE_KEY = CONFIG.LOGIN_COOKIE_KEY

export const UI_SLICE_PERSIST_KEY = 'ui-slice'

interface UIState {
  accessToken: string
  burgerOpened: boolean
  setBurgerOpened: (opened: boolean) => void
}

const useUISlice = create<UIState>()(
  devtools(
    persist(
      (set) => {
        return {
          accessToken: Cookies.get(LOGIN_COOKIE_KEY) ?? '',
          burgerOpened: false,
          setBurgerOpened: (opened: boolean) => set(setBurgerOpened(opened), false, `setBurgerOpened`),
        }
      },
      {
        version: 2,
        name: UI_SLICE_PERSIST_KEY,
        partialize(state) {
          const { accessToken, ...rest } = state // always refetch access token from storage
          return rest
        },
      },
    ),
    { enabled: true },
  ),
)

export { useUISlice }

type UIAction = (...args: any[]) => Partial<UIState>

function setBurgerOpened(opened: boolean): UIAction {
  return (state: UIState) => {
    return {
      burgerOpened: opened,
    }
  }
}
