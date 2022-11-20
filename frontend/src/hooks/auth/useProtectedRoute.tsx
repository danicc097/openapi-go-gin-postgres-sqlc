import { useEffect } from 'react'
import { useUISlice } from 'src/slices/ui'
import { ToastId } from 'src/utils/toasts'
import { useAuthenticatedUser } from './useAuthenticatedUser'

export const useProtectedRoute = (
  redirectTitle = 'Access Denied',
  redirectMessage = 'Authenticated users only. Login here or create a new account to view that page',
) => {
  // const { isAuthenticated, isAdmin, isVerifiedUser, role, isLoading } = useAuthenticatedUser()
  // const { addToast } = useUISlice()
  // useEffect(() => {
  //   if (!isAuthenticated) {
  //     addToast({
  //       id: ToastId.AuthRedirect,
  //       title: redirectTitle,
  //       color: 'warning',
  //       iconType: 'alert',
  //       toastLifeTimeMs: 15000,
  //       text: redirectMessage,
  //     })
  //   }
  // }, [isAuthenticated, redirectTitle, redirectMessage, addToast])
  // return { isAuthenticated, isAdmin, isVerifiedUser, role, isLoading }
}
