import { createAvatarImageDataUrl } from 'src/utils/files'
import { ToastId } from 'src/utils/toasts'
import { useUISlice } from 'src/slices/ui'
import { EuiIcon, EuiText } from '@elastic/eui'

export const useNotificationAPI = () => {
  const { addToast } = useUISlice()

  const createTestNotification = (email: string) => {
    new Notification('Hello world!', {
      body: 'Push notification.\n\nUse this to test the notification system.',
      // image: './notification_icon.png',
      icon: createAvatarImageDataUrl(email),
      data: {
        test: 'test',
      },
    })
  }

  const showTestNotification = (email: string) => {
    if ('Notification' in window && Notification.permission === 'granted') {
      createTestNotification(email)
    } else if (Notification.permission !== 'denied') {
      Notification.requestPermission().then((permission) => {
        if (permission === 'granted') {
          createTestNotification(email)
        }
      })
    } else {
      addNotificationAccessDeniedToast()
    }
  }

  const verifyNotificationPermission = () => {
    console.log('Verifying notification API access')
    if ('Notification' in window && Notification.permission === 'granted') {
      return
    } else if (Notification.permission !== 'denied') {
      Notification.requestPermission().then((permission) => {
        if (permission !== 'granted') {
          addNotificationAccessDeniedToast()
        }
      })
    } else {
      addNotificationAccessDeniedToast()
    }
  }

  function addNotificationAccessDeniedToast() {
    addToast({
      id: ToastId.NoticationAPIAccessDenied,
      title: `Notification access denied`,
      color: 'danger',
      iconType: 'alert',
      toastLifeTimeMs: 15000,
      text: (
        <>
          <EuiText>{`Please enable it via "View site information" at the top bar 🛈 icon`}</EuiText>
        </>
      ),
    })
  }

  return {
    showTestNotification,
    verifyNotificationPermission,
  }
}
