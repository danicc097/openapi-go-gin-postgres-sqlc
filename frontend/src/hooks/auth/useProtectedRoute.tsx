import { useEffect } from 'react'
import { useUISlice } from 'src/slices/ui'
import { ToastId } from 'src/utils/toasts'
import { useAuthenticatedUser } from './useAuthenticatedUser'

export const useProtectedRoute = (
  redirectTitle = 'Access Denied',
  redirectMessage = 'Authenticated users only. Login here or create a new account to view that page',
) => {
  const { user } = useAuthenticatedUser()
  const { addToast } = useUISlice()

  useEffect(() => {
    if (!user) {
      addToast({
        id: ToastId.AuthRedirect,
        title: redirectTitle,
        color: 'warning',
        iconType: 'alert',
        toastLifeTimeMs: 15000,
        text: redirectMessage,
      })
    }
  }, [redirectTitle, redirectMessage, addToast, user])

  return { user }
}
