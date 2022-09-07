import { shallowEqual } from 'react-redux'
import { useAppDispatch, useAppSelector } from 'src/redux'

export const useUI = () => {
  const dispatch = useAppDispatch()

  const toasts = useAppSelector((state) => state.ui.toastList, shallowEqual)
  const theme = useAppSelector((state) => state.ui.theme)

  return { toasts, theme }
}
