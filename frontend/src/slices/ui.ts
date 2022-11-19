import type { Toast } from '@elastic/eui/src/components/toast/global_toast_list'
import create from 'zustand'
import { devtools, persist } from 'zustand/middleware'

type Theme = 'dark' | 'light'

interface UIState {
  theme: Theme
  toastList: Toast[]
  addToast: (toast: Toast) => void
  removeToast: (toast: Toast) => void
  switchTheme: () => void
}

const useUISlice = create<UIState>()(
  devtools(
    persist(
      (set) => ({
        theme: 'dark', // TODO zustand middleware for persisting to LS
        toastList: [],
        addToast: (toast: Toast) => set(addToast(toast)),
        removeToast: (toast: Toast) => set(removeToast(toast.id)),
        switchTheme: () => set(switchTheme()),
      }),
      { version: 1, name: 'persist-name' },
    ),
  ),
)

export { useUISlice }

type UIAction = (...args: any[]) => Partial<UIState>

function switchTheme(): UIAction {
  return (state: UIState) => {
    return {
      theme: state.theme === 'dark' ? 'light' : 'dark',
    }
  }
}

function removeToast(id: string): UIAction {
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
