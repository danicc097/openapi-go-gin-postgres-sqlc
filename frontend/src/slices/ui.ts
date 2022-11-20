import type { Toast } from '@elastic/eui/src/components/toast/global_toast_list'
import create from 'zustand'
import { devtools, persist } from 'zustand/middleware'

type Theme = 'dark' | 'light'

interface UIState {
  theme: Theme
  toastList: Toast[]
  addToast: (toast: Toast) => void
  removeToastByID: (toastID: string) => void
  dismissToast: (toast: Toast) => void
  switchTheme: () => void
}

const useUISlice = create<UIState>()(
  devtools(
    // persist(
    // TODO only for theme
    (set) => ({
      theme: localStorage.getItem('theme') as Theme,
      toastList: [],
      addToast: (toast: Toast) => set(addToast(toast)),
      removeToastByID: (toastID: string) => set(removeToastByID(toastID)),
      dismissToast: (toast: Toast) => set(dismissToast(toast.id)),
      switchTheme: () => set(switchTheme()),
    }),
    //   { version: 2, name: 'ui-slice' },
    // ),
  ),
)

export { useUISlice }

type UIAction = (...args: any[]) => Partial<UIState>

function switchTheme(): UIAction {
  return (state: UIState) => {
    const newTheme = state.theme === 'dark' ? 'light' : 'dark'
    localStorage.setItem('theme', newTheme)
    return {
      theme: newTheme,
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
