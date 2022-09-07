import { shallowEqual } from 'react-redux'
import { useAppDispatch, useAppSelector } from 'src/redux'
import uiSlice from 'src/redux/slices/ui'

export const useUI = () => {
  const dispatch = useAppDispatch()

  const toasts = useAppSelector((state) => state.ui.toastList, shallowEqual)
  const theme = useAppSelector((state) => state.ui.theme)

  const switchTheme = () => {
    dispatch(uiSlice.actions.switchTheme())
  }

  const addToast = (payload: unknown) => {
    dispatch(uiSlice.actions.addToast(payload))
  }

  return { toasts, theme, addToast, switchTheme }
}
