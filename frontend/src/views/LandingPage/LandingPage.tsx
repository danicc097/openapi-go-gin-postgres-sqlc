import { notifications } from '@mantine/notifications'
import { IconAlertCircle } from '@tabler/icons'
import { useEffect } from 'react'
import KanbanBoard from 'src/components/KanbanBoard/KanbanBoard'
import { useUISlice } from 'src/slices/ui'
import { ToastId } from 'src/utils/toasts'

export default function LandingPage() {
  useEffect(() => {
    null
  }, [])

  return <KanbanBoard></KanbanBoard>
}
