import Cookies from 'js-cookie'
import create from 'zustand'
import { devtools, persist } from 'zustand/middleware'

export const ACCESS_TOKEN_COOKIE = 'myAppAccessToken'

export const UI_SLICE_PERSIST_KEY = 'ui-slice'

interface UIState {
  twitchToken: string
  setAccessToken: (token: string) => void
  burgerOpened: boolean
  setBurgerOpened: (opened: boolean) => void
}

const useUISlice = create<UIState>()(
  devtools(
    persist(
      (set) => {
        return {
          twitchToken: Cookies.get(ACCESS_TOKEN_COOKIE),
          setAccessToken: (token: string) => set(setAccessToken(token), false, `setAccessToken`),
          burgerOpened: false,
          setBurgerOpened: (opened: boolean) => set(setBurgerOpened(opened), false, `setBurgerOpened`),
        }
      },
      { version: 2, name: UI_SLICE_PERSIST_KEY },
    ),
    { enabled: true },
  ),
)

export { useUISlice }

type UIAction = (...args: any[]) => Partial<UIState>

function setAccessToken(token: string): UIAction {
  return (state: UIState) => {
    Cookies.set(ACCESS_TOKEN_COOKIE, token, {
      expires: 365,
      sameSite: 'none',
      secure: true,
    })
    return {
      twitchToken: token,
    }
  }
}

function setBurgerOpened(opened: boolean): UIAction {
  return (state: UIState) => {
    return {
      burgerOpened: opened,
    }
  }
}
