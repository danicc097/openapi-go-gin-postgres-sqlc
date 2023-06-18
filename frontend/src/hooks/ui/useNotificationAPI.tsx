import { createAvatarImageDataUrl } from 'src/utils/files'
import { ToastId } from 'src/utils/toasts'
import { useUISlice } from 'src/slices/ui'
import { notifications } from '@mantine/notifications'
import { IconForbid } from '@tabler/icons'

export const useNotificationAPI = () => {
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
    notifications.show({
      id: ToastId.NoticationAPIAccessDenied,
      title: `Notification access denied`,
      color: 'danger',
      icon: <IconForbid size="1.2rem" />,
      autoClose: 15000,
      message: `Please enable it via "View site information" at the top bar 🛈 icon`,
    })
  }

  return {
    showTestNotification,
    verifyNotificationPermission,
  }
}
