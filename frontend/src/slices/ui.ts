import type { Toast } from '@elastic/eui/src/components/toast/global_toast_list'
import create from 'zustand'
import { devtools, persist } from 'zustand/middleware'

export type Theme = 'dark' | 'light'

interface UIState {
  theme: Theme
  toastList: Toast[]
  addToast: (toast: Toast) => void
  removeToastByID: (toastID: string) => void
  dismissToast: (toast: Toast) => void
  setTheme: (theme: Theme) => void
}

const useUISlice = create<UIState>()(
  devtools(
    // persist(
    (set) => {
      const theme = (localStorage.getItem('theme') ?? 'light') as Theme
      return {
        theme: theme,
        toastList: [],
        addToast: (toast: Toast) => set(addToast(toast), false, `addToast-${toast.id}`),
        removeToastByID: (toastID: string) => set(removeToastByID(toastID), false, `removeToastByID-${toastID}`),
        dismissToast: (toast: Toast) => set(dismissToast(toast.id), false, `dismissToast-${toast.id}`),
        setTheme: (theme: Theme) => set(setTheme(theme), false, `setTheme-${theme}`),
      }
    },
    //   { version: 2, name: 'ui-slice' },
    // ),
    { enabled: true },
  ),
)

export { useUISlice }

type UIAction = (...args: any[]) => Partial<UIState>

function setTheme(theme: Theme): UIAction {
  return (state: UIState) => {
    localStorage.setItem('theme', theme)
    return {
      theme: theme,
    }
  }
}

function removeToastByID(toastID: string): UIAction {
  return (state: UIState) => {
    return {
      toastList: state.toastList.filter((toast) => toast.id !== toastID),
    }
  }
}

function dismissToast(id: string): UIAction {
  return (state: UIState) => {
    return {
      toastList: state.toastList.filter((toast) => toast.id !== id),
    }
  }
}
function addToast(toast: Toast): UIAction {
  return (state: UIState) => {
    state.toastList.push(toast)
    return {
      toastList: state.toastList,
    }
  }
}
